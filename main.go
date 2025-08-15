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
			fmt.Println("Planting feature coming soon!")
		case "w", "water":
			fmt.Println("Watering feature coming soon!")
		case "h", "harvest":
			fmt.Println("Harvesting feature coming soon!")
		case "s", "shop":
			fmt.Println("Shop feature coming soon!")
		case "sleep":
			fmt.Println("Sleep feature coming soon!")
		default:
			fmt.Println("Invalid action. Try again!")
		}
	}
}