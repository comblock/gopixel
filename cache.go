package gopixel

import "github.com/comblock/gopixel/structs"

type Cache struct {
	Players                map[string]*structs.Player
	SkyblockActiveAuctions *structs.SkyblockActiveAuctions
	Bazaar                 *structs.Bazaar
	SkyblockProfile        map[string]*structs.SkyblockProfile
	SkyblockProfiles       map[string]*structs.SkyblockProfiles
	SkyblockNews           *structs.SkyblockNews
	Leaderboards           *structs.Leaderboards
	SkyblockItems          *structs.SkyblockItems
}

func newCache() *Cache {
	return &Cache{
		Players:                make(map[string]*structs.Player),
		SkyblockActiveAuctions: new(structs.SkyblockActiveAuctions),
		Bazaar:                 new(structs.Bazaar),
		SkyblockProfile:        make(map[string]*structs.SkyblockProfile),
		SkyblockProfiles:       make(map[string]*structs.SkyblockProfiles),
		SkyblockNews:           new(structs.SkyblockNews),
		Leaderboards:           new(structs.Leaderboards),
		SkyblockItems:          new(structs.SkyblockItems),
	}
}
