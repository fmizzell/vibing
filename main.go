package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type CropType string

const (
	CarrotSeed CropType = "carrot"
	TomatoSeed CropType = "tomato"
	CornSeed   CropType = "corn"
)

type CropInfo struct {
	Name        string
	SeedPrice   int
	SellPrice   int
	GrowthDays  int
	DisplayChar string
}

var Crops = map[CropType]CropInfo{
	CarrotSeed: {"Carrots", 5, 12, 1, "C"},
	TomatoSeed: {"Tomatoes", 8, 20, 2, "T"},
	CornSeed:   {"Corn", 12, 35, 3, "O"},
}

type Plot struct {
	Crop      CropType
	DaysLeft  int
	Watered   bool
	IsPlanted bool
}

type Player struct {
	Money  int
	Energy int
	Day    int
	Seeds  map[CropType]int
}

type Game struct {
	Player Player
	Farm   [5][5]Plot
}

func NewGame() *Game {
	return &Game{
		Player: Player{
			Money:  50,
			Energy: 10,
			Day:    1,
			Seeds:  make(map[CropType]int),
		},
		Farm: [5][5]Plot{},
	}
}

func (g *Game) DisplayFarm() {
	fmt.Printf("\n=== Farm Vibes ===\n")
	fmt.Printf("Day %d | Money: $%d | Energy: %d/10\n\n", g.Player.Day, g.Player.Money, g.Player.Energy)

	fmt.Println("Your Farm:")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			plot := g.Farm[i][j]
			if !plot.IsPlanted {
				fmt.Print(". ")
			} else if plot.DaysLeft > 0 {
				fmt.Print("^ ")
			} else {
				cropInfo := Crops[plot.Crop]
				fmt.Printf("%s ", cropInfo.DisplayChar)
			}
		}
		fmt.Println()
	}
}

func (g *Game) ShowActions() {
	fmt.Println("\n[P]lant [W]ater [H]arvest [S]hop [Sleep] [Q]uit")
	fmt.Print("What would you like to do? ")
}

func (g *Game) BuySeeds(cropType CropType, quantity int) error {
	cropInfo := Crops[cropType]
	totalCost := cropInfo.SeedPrice * quantity

	if g.Player.Money < totalCost {
		return errors.New("insufficient funds")
	}

	g.Player.Money -= totalCost
	g.Player.Seeds[cropType] += quantity
	return nil
}

func (g *Game) PlantSeed(row, col int, cropType CropType) error {
	if row < 0 || row >= 5 || col < 0 || col >= 5 {
		return errors.New("invalid position")
	}

	if g.Farm[row][col].IsPlanted {
		return errors.New("plot already planted")
	}

	if g.Player.Seeds[cropType] <= 0 {
		return errors.New("no seeds available")
	}

	if g.Player.Energy <= 0 {
		return errors.New("insufficient energy")
	}

	cropInfo := Crops[cropType]
	g.Farm[row][col] = Plot{
		Crop:      cropType,
		DaysLeft:  cropInfo.GrowthDays,
		Watered:   false,
		IsPlanted: true,
	}

	g.Player.Seeds[cropType]--
	g.Player.Energy--

	return nil
}

func (g *Game) WaterPlot(row, col int) error {
	if row < 0 || row >= 5 || col < 0 || col >= 5 {
		return errors.New("invalid position")
	}

	if !g.Farm[row][col].IsPlanted {
		return errors.New("no crop planted")
	}

	if g.Player.Energy <= 0 {
		return errors.New("insufficient energy")
	}

	g.Farm[row][col].Watered = true
	g.Player.Energy--

	return nil
}

func (g *Game) Sleep() {
	g.Player.Day++
	g.Player.Energy = 10

	// Advance crop growth
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if g.Farm[i][j].IsPlanted && g.Farm[i][j].DaysLeft > 0 {
				g.Farm[i][j].DaysLeft--
			}
		}
	}
}

func (g *Game) HarvestPlot(row, col int) (int, error) {
	if row < 0 || row >= 5 || col < 0 || col >= 5 {
		return 0, errors.New("invalid position")
	}

	plot := &g.Farm[row][col]
	if !plot.IsPlanted {
		return 0, errors.New("no crop planted")
	}

	if plot.DaysLeft > 0 {
		return 0, errors.New("crop not ready")
	}

	cropInfo := Crops[plot.Crop]
	earnings := cropInfo.SellPrice

	g.Player.Money += earnings

	// Clear the plot
	*plot = Plot{}

	return earnings, nil
}

func main() {
	game := NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Farm Vibes! 🌱")

	for {
		game.DisplayFarm()
		game.ShowActions()

		if !scanner.Scan() {
			break
		}

		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch input {
		case "q", "quit":
			fmt.Println("Thanks for playing Farm Vibes! 🌾")
			return
		case "p", "plant":
			fmt.Print("Enter row (0-4): ")
			if !scanner.Scan() { break }
			row := scanner.Text()
			fmt.Print("Enter col (0-4): ")
			if !scanner.Scan() { break }
			col := scanner.Text()
			fmt.Print("Enter crop (carrot/tomato/corn): ")
			if !scanner.Scan() { break }
			crop := CropType(strings.ToLower(scanner.Text()))
			
			var r, c int
			fmt.Sscanf(row, "%d", &r)
			fmt.Sscanf(col, "%d", &c)
			
			if err := game.PlantSeed(r, c, crop); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Planted successfully!")
			}
		case "w", "water":
			fmt.Print("Enter row (0-4): ")
			if !scanner.Scan() { break }
			row := scanner.Text()
			fmt.Print("Enter col (0-4): ")
			if !scanner.Scan() { break }
			col := scanner.Text()
			
			var r, c int
			fmt.Sscanf(row, "%d", &r)
			fmt.Sscanf(col, "%d", &c)
			
			if err := game.WaterPlot(r, c); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Watered successfully!")
			}
		case "h", "harvest":
			fmt.Print("Enter row (0-4): ")
			if !scanner.Scan() { break }
			row := scanner.Text()
			fmt.Print("Enter col (0-4): ")
			if !scanner.Scan() { break }
			col := scanner.Text()
			
			var r, c int
			fmt.Sscanf(row, "%d", &r)
			fmt.Sscanf(col, "%d", &c)
			
			if earnings, err := game.HarvestPlot(r, c); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Harvested! Earned $%d\n", earnings)
			}
		case "s", "shop":
			fmt.Println("\n=== Seed Shop ===")
			fmt.Println("1. Carrot seeds - $5 each (1 day, sells for $12)")
			fmt.Println("2. Tomato seeds - $8 each (2 days, sells for $20)")
			fmt.Println("3. Corn seeds - $12 each (3 days, sells for $35)")
			fmt.Print("Enter crop type (carrot/tomato/corn) or 'back': ")
			if !scanner.Scan() { break }
			shopInput := strings.ToLower(scanner.Text())
			
			if shopInput == "back" {
				continue
			}
			
			crop := CropType(shopInput)
			fmt.Print("Enter quantity: ")
			if !scanner.Scan() { break }
			var quantity int
			fmt.Sscanf(scanner.Text(), "%d", &quantity)
			
			if err := game.BuySeeds(crop, quantity); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Bought %d %s seeds!\n", quantity, crop)
			}
		case "sleep":
			game.Sleep()
			fmt.Printf("You slept well! It's now day %d and your energy is restored.\n", game.Player.Day)
		default:
			fmt.Println("Invalid action. Try again!")
		}
	}
}
