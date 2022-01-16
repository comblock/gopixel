package gopixel

import (
	"fmt"

	json "github.com/mailru/easyjson"

	"errors"

	"github.com/comblock/gopixel/structs"
)

// Method to get the global achievements
func (client *Client) Achievements() (*structs.Achievements, error) {
	var achievements *structs.Achievements = new(structs.Achievements)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/achievements?key=%v", client.Key))

	if err != nil {
		return achievements, err
	}

	err = json.Unmarshal(data, achievements)

	if !achievements.Success {
		err = errors.New(achievements.Cause)
	}

	return achievements, err
}

// Method to get a list of the games
func (client *Client) Games() (*structs.Games, error) {
	var games *structs.Games = new(structs.Games)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/games?key=%v", client.Key))
	if err != nil {
		return games, err
	}

	err = json.Unmarshal(data, games)

	return games, err
}

// Method to get the active boosters
func (client *Client) Boosters() (*structs.Boosters, error) {
	var boosters *structs.Boosters = new(structs.Boosters)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/boosters?key=%v", client.Key))
	if err != nil {
		return boosters, err
	}

	err = json.Unmarshal(data, boosters)

	return boosters, err
}

// Method to get a list of all challenges
func (client *Client) Challenges() (*structs.Challenges, error) {
	var challenges *structs.Challenges = new(structs.Challenges)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/challenges?challenges=%v", client.Key))
	if err != nil {
		return challenges, err
	}

	err = json.Unmarshal(data, challenges)

	return challenges, err
}

// Method to check if a key is valid and
func (client *Client) KeyData() (*structs.Key, error) {
	var key *structs.Key = new(structs.Key)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/key?key=%v", client.Key))
	if err != nil {
		return key, err
	}

	err = json.Unmarshal(data, key)

	if !key.Success {
		err = errors.New(key.Cause)
	}

	return key, err
}

// Method to get the current player counts
func (client *Client) PlayerCounts() (*structs.PlayerCounts, error) {
	var counts *structs.PlayerCounts = new(structs.PlayerCounts)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/counts?key=%v", client.Key))
	if err != nil {
		return counts, err
	}

	err = json.Unmarshal(data, counts)

	return counts, err
}

// Method to get the current leaderboards
func (client *Client) Leaderboards() (*structs.Leaderboards, error) {
	var leaderboards *structs.Leaderboards = new(structs.Leaderboards)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/leaderboards?key=%v", client.Key))
	if err != nil {
		return leaderboards, err
	}

	err = json.Unmarshal(data, leaderboards)

	return leaderboards, err
}

// Method to get the punishment statistics
func (client *Client) PunishmentStats() (*structs.PunishmentStats, error) {
	var stats *structs.PunishmentStats = new(structs.PunishmentStats)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/punishmentstats?key=%v", client.Key))
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(data, stats)

	return stats, err
}

// Method to get the quests
func (client *Client) Quests() (*structs.Quests, error) {
	var quests *structs.Quests = new(structs.Quests)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/quests?key=%v", client.Key))
	if err != nil {
		return quests, err
	}

	err = json.Unmarshal(data, quests)

	return quests, err
}

// Method to get the ranked skywars data of a player. Will return an error if no data is found
func (client *Client) RankedSkywars(name string) (*structs.RankedSkywars, error) {
	var rankedSkywars *structs.RankedSkywars = new(structs.RankedSkywars)

	uuid, err := client.Uuid(name)

	if err != nil {
		return rankedSkywars, err
	}

	data, err := client.get(fmt.Sprintf("api.hypixel.net/player/ranked/skywars?uuid=%v&key=%v", uuid, client.Key))

	if err != nil {
		return rankedSkywars, err
	}

	err = json.Unmarshal(data, rankedSkywars)

	if !rankedSkywars.Success {
		err = errors.New(rankedSkywars.Cause)
	}

	return rankedSkywars, err
}

func (client *Client) VanityPets() (*structs.Vanity, error) {
	var VanityPets *structs.Vanity = new(structs.Vanity)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/vanity/pets?key=%v", client.Key))
	if err != nil {
		return VanityPets, err
	}

	err = json.Unmarshal(data, VanityPets)

	if !VanityPets.Success {
		err = errors.New(VanityPets.Cause)
	}

	return VanityPets, err
}

func (client *Client) VanityCompanions() (*structs.Vanity, error) {
	var VanityCompanions *structs.Vanity = new(structs.Vanity)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/vanity/companions?key=%v", client.Key))
	if err != nil {
		return VanityCompanions, err
	}

	err = json.Unmarshal(data, VanityCompanions)

	if !VanityCompanions.Success {
		err = errors.New(VanityCompanions.Cause)
	}

	return VanityCompanions, err
}
