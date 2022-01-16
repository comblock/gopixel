package gopixel

import (
	"fmt"

	json "github.com/mailru/easyjson"

	"errors"

	"github.com/comblock/gopixel/structs"
)

// Function to convert a player name to uuid using the mojang api
func (client *Client) Uuid(name string) (string, error) {
	data, err := client.get(fmt.Sprintf("api.mojang.com/users/profiles/minecraft/%v", name))

	var mojangPlayer structs.MojangPlayer

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, &mojangPlayer)

	return mojangPlayer.ID, err

}

// Method to get the friends of a player
func (client *Client) Friends(name string) (*structs.Friends, error) {
	var friends *structs.Friends = new(structs.Friends)

	uuid, err := client.Uuid(name)

	if err != nil {
		return friends, err
	}
	data, err := client.get(fmt.Sprintf("api.hypixel.net/friends?uuid=%v&key=%v", uuid, client.Key))

	if err != nil {
		return friends, err
	}

	err = json.Unmarshal(data, friends)

	if !friends.Success {
		err = errors.New(friends.Cause)
	}

	return friends, err
}

// Method to get a player's status
func (client *Client) PlayerStatus(name string) (*structs.PlayerStatus, error) {
	var playerStatus *structs.PlayerStatus = new(structs.PlayerStatus)

	uuid, err := client.Uuid(name)

	if err != nil {
		return playerStatus, err
	}
	data, err := client.get(fmt.Sprintf("api.hypixel.net/status?uuid=%v&key=%v", uuid, client.Key))

	if err != nil {
		return playerStatus, err
	}

	err = json.Unmarshal(data, playerStatus)

	if !playerStatus.Success {
		err = errors.New(playerStatus.Cause)
	}

	return playerStatus, err
}

// Method to get the data of a player
func (client *Client) PlayerData(name string) (*structs.Player, error) {
	var player *structs.Player = new(structs.Player)

	uuid, err := client.Uuid(name)

	if err != nil {
		return player, err
	}
	data, err := client.get(fmt.Sprintf("api.hypixel.net/player?uuid=%v&key=%v", uuid, client.Key))

	if err != nil {
		return player, err
	}

	err = json.Unmarshal(data, player)

	if !player.Success {
		err = errors.New(player.Cause)
	}

	return player, err
}

// Method to get the recently played games of a player
func (client *Client) RecentGames(name string) (*structs.RecentGames, error) {
	var recentGames *structs.RecentGames = new(structs.RecentGames)

	uuid, err := client.Uuid(name)
	if err != nil {
		return recentGames, err
	}

	data, err := client.get(fmt.Sprintf("api.hypixel.net/recentgames?uuid=%v&key=%v", uuid, client.Key))
	if err != nil {
		return recentGames, err
	}

	err = json.Unmarshal(data, recentGames)

	return recentGames, err
}
