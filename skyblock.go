package gopixel

import (
	"fmt"
	"sync"

	json "github.com/mailru/easyjson"

	"errors"

	"github.com/comblock/gopixel/structs"
)

// Method to get the bazaar data
func (client *Client) Bazaar() (*structs.Bazaar, error) {
	var bazaar *structs.Bazaar = new(structs.Bazaar)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/bazaar?key=%v", client.Key))
	if err != nil {
		return bazaar, err
	}

	err = json.Unmarshal(data, bazaar)

	if !bazaar.Success {
		err = errors.New(bazaar.Cause)
	}

	client.Cache.Bazaar = bazaar

	return bazaar, err
}

//  Method to get the active skyblock auctions
// ! DO NOT USE THIS WITHOUT CACHING, it will send out a lot of requests (50 or so) and this can rate limit your api key very quickly
func (client *Client) SkyblockActiveAuctions() (*structs.SkyblockActiveAuctions, error) {
	var auctions *structs.SkyblockActiveAuctions = new(structs.SkyblockActiveAuctions)
	var auctionsLock sync.Mutex
	var wg sync.WaitGroup
	var firstPage *structs.SkyblockActiveAuctionsPage = new(structs.SkyblockActiveAuctionsPage)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/auctions?key=%v", client.Key))
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, firstPage)

	auctions.Auctions = append(auctions.Auctions, firstPage.Auctions...)

	for i := firstPage.Page + 1; i < firstPage.TotalPages; i++ {
		wg.Add(1)
		go func(i int, auctions *structs.SkyblockActiveAuctions, auctionsLock *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()

			var page structs.SkyblockActiveAuctionsPage

			data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/auctions?key=%v&page=%v", client.Key, i))
			if err != nil {
				return
			}
			if err := json.Unmarshal(data, &page); err != nil {
				return
			}
			auctionsLock.Lock()
			auctions.Auctions = append(auctions.Auctions, page.Auctions...)
			auctionsLock.Unlock()
		}(i, auctions, &auctionsLock, &wg)
	}

	wg.Wait()

	client.Cache.SkyblockActiveAuctions = auctions

	return auctions, err
}

func (client *Client) SkyblockActiveAuctionsPage(page int) (*structs.SkyblockActiveAuctionsPage, error) {
	var auctions *structs.SkyblockActiveAuctionsPage = new(structs.SkyblockActiveAuctionsPage)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/auctions?key=%v&page=%v", client.Key, page))
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, auctions)

	return auctions, err
}

// Method to get the auctions of a player
func (client *Client) SkyblockAuctionByPlayer(player string) (*structs.SkyblockAuctions, error) {
	var auctions *structs.SkyblockAuctions = new(structs.SkyblockAuctions)

	uuid, err := client.Uuid(player)
	if err != nil {
		return auctions, err
	}

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/auction?player=%v&key=%v", uuid, client.Key))
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, auctions)

	return auctions, err
}

// Method to get an auction by its uuid
func (client *Client) SkyblockAuctionByUuid(uuid string) (*structs.SkyblockAuctions, error) {
	var auctions *structs.SkyblockAuctions = new(structs.SkyblockAuctions)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/auction?uuid=%v&key=%v", uuid, client.Key))
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, auctions)

	return auctions, err
}

// Method to get an auction by a profile uuid
func (client *Client) SkyblockAuctionByProfileUuid(uuid string) (*structs.SkyblockAuctions, error) {
	var auctions *structs.SkyblockAuctions = new(structs.SkyblockAuctions)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/auction?profile=%v&key=%v", uuid, client.Key))
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, auctions)

	return auctions, err
}

// Method to get a list of the skyblock collections
func (client *Client) SkyblockCollections() (*structs.SkyblockCollections, error) {
	var skyblockCollections *structs.SkyblockCollections = new(structs.SkyblockCollections)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/skyblock/collections?key=%v", client.Key))
	if err != nil {
		return skyblockCollections, err
	}

	err = json.Unmarshal(data, skyblockCollections)

	return skyblockCollections, err
}

// Method to get the ended auctions
func (client *Client) SkyblockEndedAuctions() (*structs.SkyblockEndedAuctions, error) {
	var auctions *structs.SkyblockEndedAuctions = new(structs.SkyblockEndedAuctions)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/auctions_ended?key=%v", client.Key))
	if err != nil {
		return auctions, err
	}

	err = json.Unmarshal(data, auctions)

	return auctions, err
}

// Method to get a list of the skyblock items
func (client *Client) SkyblockItems() (*structs.SkyblockItems, error) {
	var skyblockItems *structs.SkyblockItems = new(structs.SkyblockItems)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/skyblock/items?key=%v", client.Key))
	if err != nil {
		return skyblockItems, err
	}

	err = json.Unmarshal(data, skyblockItems)

	return skyblockItems, err
}

// Method to get the skyblock news
func (client *Client) SkyblockNews() (*structs.SkyblockNews, error) {
	var news *structs.SkyblockNews = new(structs.SkyblockNews)

	data, err := client.get("api.hypixel.net/resources/guilds/permissions?key=" + client.Key)
	if err != nil {
		return news, err
	}

	err = json.Unmarshal(data, news)

	client.Cache.SkyblockNews = news

	return news, err
}

// Method to get a skyblock profile by its uuid
func (client *Client) SkyblockProfile(profile string) (*structs.SkyblockProfile, error) {
	var skyblockProfile *structs.SkyblockProfile = new(structs.SkyblockProfile)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/profile?profile=%v&key=%v", profile, client.Key))

	if err != nil {
		return skyblockProfile, err
	}

	err = json.Unmarshal(data, skyblockProfile)

	if !skyblockProfile.Success {
		err = errors.New(skyblockProfile.Cause)
	}

	client.Cache.SkyblockProfile[profile] = skyblockProfile

	return skyblockProfile, err
}

// Method to get the skyblock profiles of a player
func (client *Client) SkyblockProfiles(name string) (*structs.SkyblockProfiles, error) {
	var skyblockProfile *structs.SkyblockProfiles = new(structs.SkyblockProfiles)

	uuid, err := client.Uuid(name)

	if err != nil {
		return skyblockProfile, err
	}

	data, err := client.get(fmt.Sprintf("api.hypixel.net/skyblock/profiles?uuid=%v&key=%v", uuid, client.Key))

	if err != nil {
		return skyblockProfile, err
	}

	err = json.Unmarshal(data, skyblockProfile)

	if !skyblockProfile.Success {
		err = errors.New(skyblockProfile.Cause)
	}

	client.Cache.SkyblockProfiles[uuid] = skyblockProfile

	return skyblockProfile, err
}

// Method to get a list of the skills in skyblock
func (client *Client) SkyblockSkills() (*structs.SkyblockSkills, error) {
	var skyblockSkills *structs.SkyblockSkills = new(structs.SkyblockSkills)

	data, err := client.get(fmt.Sprintf("api.hypixel.net/resources/skyblock/skills?key=%v", client.Key))
	if err != nil {
		return skyblockSkills, err
	}

	err = json.Unmarshal(data, skyblockSkills)

	return skyblockSkills, err
}
