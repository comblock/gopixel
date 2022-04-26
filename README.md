# Gopixel
[![GoDoc](https://img.shields.io/badge/Godoc-Reference-%2300ADD8?style=for-the-badge)](https://godoc.org/github.com/comblock/gopixel)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/comblock/gopixel?style=for-the-badge)
[![Go Report Card](https://img.shields.io/badge/go%20report-A-red.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/comblock/gopixel)
![GitHub](https://img.shields.io/github/license/comblock/gopixel?style=for-the-badge)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/comblock/gopixel?color=yellow&style=for-the-badge)

## Gopixel is now archived and will not receive future updates 

A simple and complete[^1] wrapper for the hypixel API

## Key features 
- Full API coverage[^1]
- Autocomplete for fields
- Near complete structs
- MIT license

## Installation
```
go get github.com/comblock/gopixel
```

## Usage
```go
import "github.com/comblock/gopixel"

var client *gopixel.Client = gopixel.NewClient(<your API key>, <Amount of retries (0 is default value)>)

func main() {
	// example
	bazaar := client.Bazaar()
  	fmt.Println(bazaar.Products.BrownMushroom.QuickStatus.SellPrice)

```

### Supported endpoints
- https://api.hypixel.net/player
- https://api.hypixel.net/friends
- https://api.hypixel.net/recentgames
- https://api.hypixel.net/status
- https://api.hypixel.net/guild
- https://api.hypixel.net/player/ranked/skywars
- https://api.hypixel.net/resources/games
- https://api.hypixel.net/resources/achievements
- https://api.hypixel.net/resources/challenges
- https://api.hypixel.net/resources/quests
- https://api.hypixel.net/resources/guilds/achievements
- https://api.hypixel.net/resources/guilds/permissions
- https://api.hypixel.net/resources/skyblock/collections
- https://api.hypixel.net/resources/skyblock/skills
- https://api.hypixel.net/resources/skyblock/items
- https://api.hypixel.net/skyblock/news
- https://api.hypixel.net/skyblock/auction
- https://api.hypixel.net/skyblock/auctions
- https://api.hypixel.net/skyblock/auctions_ended
- https://api.hypixel.net/skyblock/bazaar
- https://api.hypixel.net/skyblock/profile
- https://api.hypixel.net/skyblock/profiles
- https://api.hypixel.net/boosters
- https://api.hypixel.net/counts
- https://api.hypixel.net/leaderboards
- https://api.hypixel.net/punishmentstats


If you have any questions, DM me on discord (comblock#5184)
[^1]: as of 30/10/2021
