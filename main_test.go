package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFarmingJourney(t *testing.T) {
	// === Day 1: A farmer's journey begins ===
	game := NewGame()
	assert.Equal(t, 1, game.Player.Day)
	assert.Equal(t, 50, game.Player.Money)
	assert.Equal(t, 10, game.Player.Energy)
	assert.Empty(t, game.Player.Seeds)

	// Try to buy expensive corn seeds - should fail
	err := game.BuySeeds(CornSeed, 10) // $12 each = $120, but only have $50
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "insufficient funds")
	assert.Equal(t, 50, game.Player.Money) // Money unchanged

	// Buy affordable seeds instead
	err = game.BuySeeds(CarrotSeed, 3) // $5 each = $15
	assert.NoError(t, err)
	err = game.BuySeeds(TomatoSeed, 2) // $8 each = $16
	assert.NoError(t, err)
	
	// Should have $19 left and seeds in inventory
	assert.Equal(t, 19, game.Player.Money)
	assert.Equal(t, 3, game.Player.Seeds[CarrotSeed])
	assert.Equal(t, 2, game.Player.Seeds[TomatoSeed])

	// Plant some crops on the farm (costs energy)
	err = game.PlantSeed(0, 0, CarrotSeed) // Plant at position (0,0)
	assert.NoError(t, err)
	err = game.PlantSeed(0, 1, CarrotSeed)
	assert.NoError(t, err)
	err = game.PlantSeed(1, 0, TomatoSeed)
	assert.NoError(t, err)
	
	// Should have used seeds from inventory and spent energy
	assert.Equal(t, 1, game.Player.Seeds[CarrotSeed]) // 3 - 2 = 1
	assert.Equal(t, 1, game.Player.Seeds[TomatoSeed]) // 2 - 1 = 1
	assert.Equal(t, 7, game.Player.Energy) // 10 - 3 = 7 (planting costs energy)
	
	// Check farm state
	assert.True(t, game.Farm[0][0].IsPlanted)
	assert.Equal(t, CarrotSeed, game.Farm[0][0].Crop)
	assert.Equal(t, 1, game.Farm[0][0].DaysLeft) // Carrots take 1 day

	// Water the crops (costs energy)
	err = game.WaterPlot(0, 0)
	assert.NoError(t, err)
	err = game.WaterPlot(0, 1)
	assert.NoError(t, err)
	err = game.WaterPlot(1, 0)
	assert.NoError(t, err)
	
	assert.Equal(t, 4, game.Player.Energy) // 7 - 3 = 4
	assert.True(t, game.Farm[0][0].Watered)

	// Sleep to advance to next day and restore energy
	game.Sleep()
	assert.Equal(t, 2, game.Player.Day)
	assert.Equal(t, 10, game.Player.Energy) // Restored
	
	// Carrots should be ready (1 day growth), tomatoes need 1 more day
	assert.Equal(t, 0, game.Farm[0][0].DaysLeft) // Carrot ready!
	assert.Equal(t, 0, game.Farm[0][1].DaysLeft) // Carrot ready!
	assert.Equal(t, 1, game.Farm[1][0].DaysLeft) // Tomato needs 1 more day

	// Harvest the ready carrots
	money, err := game.HarvestPlot(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, 12, money) // Carrot sells for $12
	
	money, err = game.HarvestPlot(0, 1)
	assert.NoError(t, err)
	assert.Equal(t, 12, money)
	
	// Should have earned money
	assert.Equal(t, 43, game.Player.Money) // 19 + 12 + 12 = 43
	
	// Plots should be empty now
	assert.False(t, game.Farm[0][0].IsPlanted)
	assert.False(t, game.Farm[0][1].IsPlanted)

	// Sleep again to get tomatoes ready
	game.Sleep()
	assert.Equal(t, 3, game.Player.Day)
	assert.Equal(t, 0, game.Farm[1][0].DaysLeft) // Tomato ready!
	
	// Harvest tomato for big profit
	money, err = game.HarvestPlot(1, 0)
	assert.NoError(t, err)
	assert.Equal(t, 20, money) // Tomato sells for $20
	assert.Equal(t, 63, game.Player.Money) // 43 + 20 = 63
	
	// Successful farming journey! Made $13 profit (63 - 50 = 13)
}