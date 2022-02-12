# Testing

In order to test the package you need a config.json file that looks like this:
```jsonc
{
	"key": "",			// Your API key
	"retries": 0,			// The amount of times you want the program to retry after failing a get request
	"player": "",			// The player for most methods that require a player
	"skyblock_profile": "",		// The profile ID that is used to test SkyblockProfile
	"ranked_skywars_player": "",	// This is to test the RankedSkywars method, it's seperate because not all players will have data
	"guild_id": "",			// Guild ID to test GuildById
	"guild_name": "",		// Guild name to test GuildByName
	"output_file": "",		// The name of the output file, will default to test.log if not set
}
```
If a test returns a struct filled with nil values there is probably a bug or the response was empty.

Test by using this command `go test -timeout 0` to prevent the test from timing out 