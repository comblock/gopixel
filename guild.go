package gopixel

import (
	"fmt"

	json "github.com/mailru/easyjson"

	"errors"

	"github.com/comblock/gopixel/structs"
)

// Method to get a list of guild achievements
func (client *Client) GuildAchievements() (structs.GuildAchievements, error) {
	var guildAchievements structs.GuildAchievements

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/guilds/achievements?key=%v", client.Key))
	if err != nil {
		return guildAchievements, err
	}

	err = json.Unmarshal(data, &guildAchievements)

	return guildAchievements, err
}

// Method to get a guild by its id
func (client *Client) GuildById(id string) (structs.Guild, error) {
	var guild structs.Guild

	data, err := client.get(fmt.Sprintf("api.hypixel.net/guild?id=%v&key=%v", id, client.Key))
	if err != nil {
		return guild, err
	}

	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}

	return guild, err
}

// Method to get a guild by its name
func (client *Client) GuildByName(name string) (structs.Guild, error) {
	var guild structs.Guild

	data, err := client.get(fmt.Sprintf("api.hypixel.net/guild?name=%v&key=%v", name, client.Key))
	if err != nil {
		return guild, err
	}

	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}
	return guild, err
}

// Method to get a guild by a player
func (client *Client) GuildByPlayer(player string) (structs.Guild, error) {
	var guild structs.Guild

	uuid, err := client.Uuid(player)
	if err != nil {
		return guild, err
	}

	data, err := client.get(fmt.Sprintf("api.hypixel.net/guild?player=%v&key=%v", uuid, client.Key))
	if err != nil {
		return guild, err
	}
	err = json.Unmarshal(data, &guild)

	if !guild.Success {
		err = errors.New(guild.Cause)
	}

	return guild, err
}
