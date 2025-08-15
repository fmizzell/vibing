# Farm Vibes 🌱

A cozy CLI farming simulator built with Go and TDD.

## Current Status

- ✅ Basic game structure with 5x5 farm grid
- ✅ Crop system (Carrots, Tomatoes, Corn) with prices and growth times
- ✅ Player inventory and shop system with `BuySeeds()`
- 🚧 **WIP**: Planting, watering, harvesting mechanics
- 🚧 **WIP**: Sleep/time progression system

## TDD Journey Test

Our main test (`TestFarmingJourney`) tells the complete story:
1. Try to buy expensive seeds → insufficient funds error
2. Buy affordable seeds → success
3. Plant seeds → uses inventory and energy
4. Water crops → costs energy
5. Sleep → advance time, restore energy, grow crops
6. Harvest → earn money, clear plots

## Next Steps

Run failing tests to see what needs implementing:
```bash
go test -v
```

Currently missing: `PlantSeed()`, `WaterPlot()`, `Sleep()`, `HarvestPlot()`

## Run Game

```bash
go run main.go
```