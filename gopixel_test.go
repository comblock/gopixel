package gopixel

import (
	"encoding/json"

	"testing"

	"os"

	"github.com/davecgh/go-spew/spew"
)

// The config struct with all of the required variables
type Config struct {
	Key                 string `json:"key"`
	Retries             uint   `json:"retries"`
	Player              string `json:"player"`
	SkyblockProfile     string `json:"skyblock_profile"`
	RankedSkywarsPlayer string `json:"ranked_skywars_player"`
	GuildId             string `json:"guild_id"`
	GuildName           string `json:"guild_name"`
	OutputFile          string `json:"output_file"`
}

func getConfig() *Config {
	var config *Config
	bytes, err := os.ReadFile("config.json")

	if err != nil {
		bytes = []byte(os.Getenv("CONFIG"))
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}
	return config
}

var output *os.File
var config *Config = getConfig()
var client *Client = NewClient(config.Key, config.Retries)

func init() {
	// Opens the output file
	if config.OutputFile == "" {
		config.OutputFile = "test.log"
	}

	var err error
	output, err = os.Create(config.OutputFile)

	if err != nil {
		panic(err)
	}
}

func TestKeyData(t *testing.T) {
	key, err := client.KeyData()
	if err != nil {
		t.Error((err))
	}
	spew.Fdump(output, key)
}

func TestPlayerData(t *testing.T) {
	player, err := client.PlayerData(config.Player)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, player)
}

func TestAchievements(t *testing.T) {
	achievements, err := client.Achievements()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, achievements)
}

func TestBazaar(t *testing.T) {
	bazaar, err := client.Bazaar()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, bazaar)
}

func TestFriends(t *testing.T) {
	friends, err := client.Friends(config.Player)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, friends)
}

func TestQuests(t *testing.T) {
	quests, err := client.Quests()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, quests)
}

func TestChallenges(t *testing.T) {
	challenges, err := client.Challenges()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, challenges)
}

func TestGames(t *testing.T) {
	games, err := client.Games()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, games)
}

func TestGuildByPlayer(t *testing.T) {
	guild, err := client.GuildByPlayer(config.Player)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, guild)
}

func TestGuildById(t *testing.T) {
	guild, err := client.GuildById(config.GuildId)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, guild)
}

func TestGuildByName(t *testing.T) {
	guild, err := client.GuildByName(config.GuildName)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, guild)
}

func TestGuildAchievements(t *testing.T) {
	guildAchievements, err := client.GuildAchievements()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, guildAchievements)
}

func TestGuilPermissions(t *testing.T) {
	guildPermissions, err := client.GuildAchievements()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, guildPermissions)
}

func TestSkyblockCollections(t *testing.T) {
	skyblockCollections, err := client.SkyblockCollections()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, skyblockCollections)
}

func TestSkyblockProfiles(t *testing.T) {
	profiles, err := client.SkyblockProfiles(config.Player)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, profiles)
}

func TestSkyblockProfile(t *testing.T) {
	profiles, err := client.SkyblockProfile(config.SkyblockProfile)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, profiles)
}

func TestSkyblockItems(t *testing.T) {
	items, err := client.SkyblockItems()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, items)
}

func TestRecentGames(t *testing.T) {
	games, err := client.RecentGames(config.Player)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, games)
}

func TestPlayerStatus(t *testing.T) {
	status, err := client.PlayerStatus(config.Player)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, status)
}

func TestSkyblockSkills(t *testing.T) {
	skills, err := client.SkyblockSkills()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, skills)
}

// If this test fails with the error message "No result was found" it means that the player doesn't have any data and NOT that the function is broken
func TestRankedSkywars(t *testing.T) {
	data, err := client.RankedSkywars(config.RankedSkywarsPlayer)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, data)
}

func TestSkyblockNews(t *testing.T) {
	news, err := client.SkyblockNews()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, news)
}

func TestSkyblockActiveAuctions(t *testing.T) {
	auctions, err := client.SkyblockActiveAuctions()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, auctions)
}

func TestSkyblockActiveAuctionsPage(t *testing.T) {
	page, err := client.SkyblockActiveAuctionsPage(0)
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, page)
}

func TestSkyblockEndedAuctions(t *testing.T) {
	auctions, err := client.SkyblockActiveAuctions()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, auctions)
}

func TestBoosters(t *testing.T) {
	boosters, err := client.Boosters()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, boosters)
}

func TestPlayerCounts(t *testing.T) {
	counts, err := client.PlayerCounts()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, counts)
}

func TestLeaderboards(t *testing.T) {
	leaderboards, err := client.Leaderboards()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, leaderboards)
}

func TestPunishmentStats(t *testing.T) {
	stats, err := client.PunishmentStats()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, stats)
}

func TestVanityPets(t *testing.T) {
	pets, err := client.VanityPets()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, pets)
}

func TestVanityCompanions(t *testing.T) {
	companions, err := client.VanityCompanions()
	if err != nil {
		t.Error(err)
	}
	spew.Fdump(output, companions)
}

func TestCache(t *testing.T) {
	spew.Fdump(output, client.Cache)
}
