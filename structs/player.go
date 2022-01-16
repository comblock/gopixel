package structs

type Parkour []struct {
	TimeStart float64 `json:"timeStart"`
	TimeTook  float64 `json:"timeTook"`
}

type MojangPlayer struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Player struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	Data    struct {
		ID           string `json:"_id"`
		Achievements struct {
			ArenaBossed                     float64 `json:"arena_bossed"`
			ArenaClimbTheRanks              float64 `json:"arena_climb_the_ranks"`
			ArenaGladiator                  float64 `json:"arena_gladiator"`
			ArenaGottaWearEmAll             float64 `json:"arena_gotta_wear_em_all"`
			BlitzCoins                      float64 `json:"blitz_coins"`
			BlitzKills                      float64 `json:"blitz_kills"`
			BlitzWins                       float64 `json:"blitz_wins"`
			GeneralCoins                    float64 `json:"general_coins"`
			GeneralWins                     float64 `json:"general_wins"`
			PaintballCoins                  float64 `json:"paintball_coins"`
			PaintballKills                  float64 `json:"paintball_kills"`
			PaintballWins                   float64 `json:"paintball_wins"`
			QuakeKillingSprees              float64 `json:"quake_killing_sprees"`
			QuakeKills                      float64 `json:"quake_kills"`
			QuakeWins                       float64 `json:"quake_wins"`
			TntgamesBowSpleefWins           float64 `json:"tntgames_bow_spleef_wins"`
			TntgamesTntRunWins              float64 `json:"tntgames_tnt_run_wins"`
			TntgamesWizardsWins             float64 `json:"tntgames_wizards_wins"`
			VampirezCoins                   float64 `json:"vampirez_coins"`
			VampirezKillSurvivors           float64 `json:"vampirez_kill_survivors"`
			VampirezKillVampires            float64 `json:"vampirez_kill_vampires"`
			VampirezSurvivorWins            float64 `json:"vampirez_survivor_wins"`
			Walls3Coins                     float64 `json:"walls3_coins"`
			Walls3Kills                     float64 `json:"walls3_kills"`
			Walls3Wins                      float64 `json:"walls3_wins"`
			WallsCoins                      float64 `json:"walls_coins"`
			WallsKills                      float64 `json:"walls_kills"`
			WallsWins                       float64 `json:"walls_wins"`
			WarlordsWarriorLevel            float64 `json:"warlords_warrior_level"`
			WarlordsMageLevel               float64 `json:"warlords_mage_level"`
			WarlordsPaladinLevel            float64 `json:"warlords_paladin_level"`
			WarlordsKills                   float64 `json:"warlords_kills"`
			WarlordsShamanLevel             float64 `json:"warlords_shaman_level"`
			WarlordsAssist                  float64 `json:"warlords_assist"`
			WarlordsRepairWeapons           float64 `json:"warlords_repair_weapons"`
			WarlordsCoins                   float64 `json:"warlords_coins"`
			SkywarsKitsTeam                 float64 `json:"skywars_kits_team"`
			SkywarsWinsTeam                 float64 `json:"skywars_wins_team"`
			SkywarsKillsSolo                float64 `json:"skywars_kills_solo"`
			SkywarsKillsTeam                float64 `json:"skywars_kills_team"`
			SkywarsCages                    float64 `json:"skywars_cages"`
			SkywarsWinsSolo                 float64 `json:"skywars_wins_solo"`
			SkywarsKitsSolo                 float64 `json:"skywars_kits_solo"`
			UhcChampion                     float64 `json:"uhc_champion"`
			UhcHunter                       float64 `json:"uhc_hunter"`
			UhcMovingUp                     float64 `json:"uhc_moving_up"`
			UhcBounty                       float64 `json:"uhc_bounty"`
			WarlordsCtfWins                 float64 `json:"warlords_ctf_wins"`
			WarlordsDomWins                 float64 `json:"warlords_dom_wins"`
			WarlordsTdmWins                 float64 `json:"warlords_tdm_wins"`
			Walls3Guardian                  float64 `json:"walls3_guardian"`
			TruecombatSoloWinner            float64 `json:"truecombat_solo_winner"`
			TruecombatSoloKiller            float64 `json:"truecombat_solo_killer"`
			TruecombatKitHoarderSolo        float64 `json:"truecombat_kit_hoarder_solo"`
			TruecombatKitHoarderTeam        float64 `json:"truecombat_kit_hoarder_team"`
			Walls3Rusher                    float64 `json:"walls3_rusher"`
			SupersmashSmashWinner           float64 `json:"supersmash_smash_winner"`
			SupersmashHeroSlayer            float64 `json:"supersmash_hero_slayer"`
			SupersmashSmashChampion         float64 `json:"supersmash_smash_champion"`
			CopsandcrimsSerialKiller        float64 `json:"copsandcrims_serial_killer"`
			CopsandcrimsHeroTerrorist       float64 `json:"copsandcrims_hero_terrorist"`
			CopsandcrimsBombSpecialist      float64 `json:"copsandcrims_bomb_specialist"`
			CopsandcrimsCacBanker           float64 `json:"copsandcrims_cac_banker"`
			SkywarsKitsMega                 float64 `json:"skywars_kits_mega"`
			SkywarsWinsMega                 float64 `json:"skywars_wins_mega"`
			SkywarsKillsMega                float64 `json:"skywars_kills_mega"`
			GeneralQuestMaster              float64 `json:"general_quest_master"`
			GeneralChallenger               float64 `json:"general_challenger"`
			SpeeduhcHunter                  float64 `json:"speeduhc_hunter"`
			SpeeduhcUhcMaster               float64 `json:"speeduhc_uhc_master"`
			SpeeduhcCollector               float64 `json:"speeduhc_collector"`
			SpeeduhcSalty                   float64 `json:"speeduhc_salty"`
			SkyclashCardsUnlocked           float64 `json:"skyclash_cards_unlocked"`
			SkyclashPacksOpened             float64 `json:"skyclash_packs_opened"`
			SkyclashKills                   float64 `json:"skyclash_kills"`
			SkyclashWins                    float64 `json:"skyclash_wins"`
			ArcadeArcadeWinner              float64 `json:"arcade_arcade_winner"`
			ArcadeArcadeBanker              float64 `json:"arcade_arcade_banker"`
			ArcadeFootballPro               float64 `json:"arcade_football_pro"`
			ArcadeFarmhuntDominator         float64 `json:"arcade_farmhunt_dominator"`
			ArcadeBountyHunter              float64 `json:"arcade_bounty_hunter"`
			ArcadeTeamWork                  float64 `json:"arcade_team_work"`
			ArcadeZombieKiller              float64 `json:"arcade_zombie_killer"`
			ArcadeMiniwallsWinner           float64 `json:"arcade_miniwalls_winner"`
			TruecombatTeamWinner            float64 `json:"truecombat_team_winner"`
			TruecombatTeamKiller            float64 `json:"truecombat_team_killer"`
			QuakeCoins                      float64 `json:"quake_coins"`
			GingerbreadRacer                float64 `json:"gingerbread_racer"`
			GingerbreadBanker               float64 `json:"gingerbread_banker"`
			GingerbreadWinner               float64 `json:"gingerbread_winner"`
			VampirezZombieKiller            float64 `json:"vampirez_zombie_killer"`
			TntgamesPvpRunWins              float64 `json:"tntgames_pvp_run_wins"`
			TntgamesPvpRunKiller            float64 `json:"tntgames_pvp_run_killer"`
			TntgamesTntTagWins              float64 `json:"tntgames_tnt_tag_wins"`
			TntgamesTntWizardsKills         float64 `json:"tntgames_tnt_wizards_kills"`
			TntgamesTntTriathlon            float64 `json:"tntgames_tnt_triathlon"`
			TntgamesTntBanker               float64 `json:"tntgames_tnt_banker"`
			TntgamesTntWizardsCaps          float64 `json:"tntgames_tnt_wizards_caps"`
			QuakeHeadshots                  float64 `json:"quake_headshots"`
			QuakeWeaponArsenal              float64 `json:"quake_weapon_arsenal"`
			BlitzWinsTeams                  float64 `json:"blitz_wins_teams"`
			BlitzWarVeteran                 float64 `json:"blitz_war_veteran"`
			BlitzLooter                     float64 `json:"blitz_looter"`
			BedwarsWins                     float64 `json:"bedwars_wins"`
			BedwarsBeds                     float64 `json:"bedwars_beds"`
			BedwarsLevel                    float64 `json:"bedwars_level"`
			BedwarsLootBox                  float64 `json:"bedwars_loot_box"`
			MurdermysteryWinsAsSurvivor     float64 `json:"murdermystery_wins_as_survivor"`
			MurdermysteryKillsAsMurderer    float64 `json:"murdermystery_kills_as_murderer"`
			MurdermysteryWinsAsMurderer     float64 `json:"murdermystery_wins_as_murderer"`
			TruecombatKing                  float64 `json:"truecombat_king"`
			Halloween2017Pumpkinator        float64 `json:"halloween2017_pumpkinator"`
			Halloween2017PumpkinSmasher     float64 `json:"halloween2017_pumpkin_smasher"`
			Christmas2017NoChristmas        float64 `json:"christmas2017_no_christmas"`
			Christmas2017SantaSaysRounds    float64 `json:"christmas2017_santa_says_rounds"`
			Christmas2017Advent             float64 `json:"christmas2017_advent"`
			BuildbattleBuildbattlePoints    float64 `json:"buildbattle_buildbattle_points"`
			DuelsDuelsWinner                float64 `json:"duels_duels_winner"`
			BuildbattleGuessTheBuildGuesses float64 `json:"buildbattle_guess_the_build_guesses"`
			BuildbattleGuessTheBuildWinner  float64 `json:"buildbattle_guess_the_build_winner"`
			BuildbattleBuildBattleScore     float64 `json:"buildbattle_build_battle_score"`
			BuildbattleBuildBattleVoter     float64 `json:"buildbattle_build_battle_voter"`
			BuildbattleBuildBattlePoints    float64 `json:"buildbattle_build_battle_points"`
			DuelsDuelsWinStreak             float64 `json:"duels_duels_win_streak"`
			QuakeGodlikes                   float64 `json:"quake_godlikes"`
			BedwarsBedwarsKiller            float64 `json:"bedwars_bedwars_killer"`
			BedwarsCollectorsEdition        float64 `json:"bedwars_collectors_edition"`
			WarlordsCtfObjective            float64 `json:"warlords_ctf_objective"`
			SkywarsWinsLab                  float64 `json:"skywars_wins_lab"`
			MurdermysteryHoarder            float64 `json:"murdermystery_hoarder"`
			PaintballKillStreaks            float64 `json:"paintball_kill_streaks"`
			GingerbreadMystery              float64 `json:"gingerbread_mystery"`
			PaintballHatCollector           float64 `json:"paintball_hat_collector"`
			ArenaMagicalBox                 float64 `json:"arena_magical_box"`
			Christmas2017Advent2018         float64 `json:"christmas2017_advent_2018"`
			CopsandcrimsHeadshotKills       float64 `json:"copsandcrims_headshot_kills"`
			DuelsGoals                      float64 `json:"duels_goals"`
			DuelsBridgeDuelsWins            float64 `json:"duels_bridge_duels_wins"`
			DuelsBridgeWins                 float64 `json:"duels_bridge_wins"`
			DuelsBridgeDoublesWins          float64 `json:"duels_bridge_doubles_wins"`
			DuelsDuelsTraveller             float64 `json:"duels_duels_traveller"`
			DuelsBridgeWinStreak            float64 `json:"duels_bridge_win_streak"`
			DuelsUniqueMapWins              float64 `json:"duels_unique_map_wins"`
			UhcConsumer                     float64 `json:"uhc_consumer"`
			BlitzKitCollector               float64 `json:"blitz_kit_collector"`
			SkywarsHeads                    float64 `json:"skywars_heads"`
			ArcadeZombiesNiceShot           float64 `json:"arcade_zombies_nice_shot"`
			ArcadeZombiesRoundProgression   float64 `json:"arcade_zombies_round_progression"`
			ArcadeZombiesHighRound          float64 `json:"arcade_zombies_high_round"`
			EasterThrowEggs                 float64 `json:"easter_throw_eggs"`
			SkyclashTreasureHunter          float64 `json:"skyclash_treasure_hunter"`
			SkyclashMobBeheading            float64 `json:"skyclash_mob_beheading"`
			WarlordsDomObjective            float64 `json:"warlords_dom_objective"`
			BlitzKitExpert                  float64 `json:"blitz_kit_expert"`
			BlitzMobMaster                  float64 `json:"blitz_mob_master"`
			BlitzFightingExpert             float64 `json:"blitz_fighting_expert"`
			BlitzTreasureSeeker             float64 `json:"blitz_treasure_seeker"`
			BlitzKitExperienceCollector     float64 `json:"blitz_kit_experience_collector"`
			TntgamesBlockRunner             float64 `json:"tntgames_block_runner"`
			SkyblockMinionLover             float64 `json:"skyblock_minion_lover"`
			SkyblockExcavator               float64 `json:"skyblock_excavator"`
			SkyblockCombat                  float64 `json:"skyblock_combat"`
			SkyblockGatherer                float64 `json:"skyblock_gatherer"`
			SkyblockAngler                  float64 `json:"skyblock_angler"`
			SkyblockHarvester               float64 `json:"skyblock_harvester"`
			SkyblockTreasury                float64 `json:"skyblock_treasury"`
			SkyblockConcoctor               float64 `json:"skyblock_concoctor"`
			SkyblockAugmentation            float64 `json:"skyblock_augmentation"`
			TntgamesClinic                  float64 `json:"tntgames_clinic"`
			SpeeduhcPromotion               float64 `json:"speeduhc_promotion"`
			TruecombatFeelsLucky            float64 `json:"truecombat_feels_lucky"`
			Christmas2017Advent2019         float64 `json:"christmas2017_advent_2019"`
			Christmas2017PresentCollector   float64 `json:"christmas2017_present_collector"`
			SkyblockUniqueGifts             float64 `json:"skyblock_unique_gifts"`
			Christmas2017SecretSanta        float64 `json:"christmas2017_secret_santa"`
			EasterEggFinder                 float64 `json:"easter_egg_finder"`
			EasterMasterTracker             float64 `json:"easter_master_tracker"`
			PitContracts                    float64 `json:"pit_contracts"`
			PitGold                         float64 `json:"pit_gold"`
			PitKills                        float64 `json:"pit_kills"`
			PitPrestiges                    float64 `json:"pit_prestiges"`
			PitRenown                       float64 `json:"pit_renown"`
			SkyblockDomesticator            float64 `json:"skyblock_domesticator"`
			SkyblockSlayer                  float64 `json:"skyblock_slayer"`
			SummerTreasureHoarder           float64 `json:"summer_treasure_hoarder"`
			SummerTreasureMaster            float64 `json:"summer_treasure_master"`
			SkywarsNewDayNewChallenge       float64 `json:"skywars_new_day_new_challenge"`
			SkywarsYouReAStar               float64 `json:"skywars_you_re_a_star"`
			DuelsDuelsDivision              float64 `json:"duels_duels_division"`
			ArcadeCtwOhSheep                float64 `json:"arcade_ctw_oh_sheep"`
			ArcadeCtwSlayer                 float64 `json:"arcade_ctw_slayer"`
			SkyblockDungeoneer              float64 `json:"skyblock_dungeoneer"`
			SkyblockTreasureHunter          float64 `json:"skyblock_treasure_hunter"`
			SkyblockHardWorkingMiner        float64 `json:"skyblock_hard_working_miner"`
			SkyblockGoblinKiller            float64 `json:"skyblock_goblin_killer"`
		} `json:"achievements"`
		AchievementsOneTime []string `json:"achievementsOneTime"`
		AutoSpawnPet        bool     `json:"auto_spawn_pet"`
		Channel             string   `json:"channel"`
		Chat                bool     `json:"chat"`
		Disguise            string   `json:"disguise"`
		Displayname         string   `json:"displayname"`
		EulaCoins           bool     `json:"eulaCoins"`
		FireworkStorage     []struct {
			FlightDuration float64 `json:"flight_duration"`
			Shape          string  `json:"shape"`
			Trail          bool    `json:"trail"`
			Twinkle        bool    `json:"twinkle"`
			Colors         string  `json:"colors"`
			FadeColors     string  `json:"fade_colors"`
			Selected       bool    `json:"selected"`
		} `json:"fireworkStorage"`
		FirstLogin                 float64       `json:"firstLogin"`
		FriendRequests             []interface{} `json:"friendRequests"`
		FriendRequestsUUID         []interface{} `json:"friendRequestsUuid"`
		Gadget                     string        `json:"gadget"`
		Karma                      float64       `json:"karma"`
		KnownAliases               []string      `json:"knownAliases"`
		KnownAliasesLower          []string      `json:"knownAliasesLower"`
		LastLogin                  float64       `json:"lastLogin"`
		MostRecentMinecraftVersion float64       `json:"mostRecentMinecraftVersion"`
		MostRecentlyThanked        string        `json:"mostRecentlyThanked"`
		MostRecentlyTipped         string        `json:"mostRecentlyTipped"`
		MostRecentlyTippedUUID     string        `json:"mostRecentlyTippedUuid"`
		NetworkExp                 float64       `json:"networkExp"`
		Notifications              bool          `json:"notifications"`
		PackageRank                string        `json:"packageRank"`
		ParkourCompletions         struct {
			ArcadeGames    Parkour `json:"ArcadeGames"`
			BlitzLobby     Parkour `json:"BlitzLobby"`
			MainLobby      Parkour `json:"MainLobby"`
			QuakeCraft     Parkour `json:"QuakeCraft"`
			Tnt            Parkour `json:"TNT"`
			Christmas      Parkour `json:"christmas"`
			Main           Parkour `json:"main"`
			Turbo          Parkour `json:"Turbo"`
			Skywars        Parkour `json:"Skywars"`
			Paintball      Parkour `json:"Paintball"`
			CopsnCrims     Parkour `json:"CopsnCrims"`
			Vampirez       Parkour `json:"vampirez"`
			TruePVPParkour Parkour `json:"TruePVPParkour"`
			TheWallsLobby  Parkour `json:"TheWallsLobby"`
			MegaWalls      Parkour `json:"MegaWalls"`
			Uhc            Parkour `json:"uhc"`
			NewMainLobby   Parkour `json:"NewMainLobby"`
			Prototype      Parkour `json:"Prototype"`
			Bedwars        Parkour `json:"Bedwars"`
			Warlords       Parkour `json:"Warlords"`
			SpeedUHC       Parkour `json:"SpeedUHC"`
			SkywarsAug2017 Parkour `json:"SkywarsAug2017"`
		} `json:"ParkourCompletions"`
		Playername string `json:"playername"`
		Quests     struct {
			Blitzerk struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"blitzerk"`
			ExplosiveGames struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"explosive_games"`
			Gladiator struct {
				Active struct {
					Objectives struct {
						FourV4 float64 `json:"4v4"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"gladiator"`
			Megawaller struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"megawaller"`
			NuggetWarriors struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Vampirez float64 `json:"vampirez"`
						Quake    float64 `json:"quake"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"nugget_warriors"`
			PaintballExpert struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Kill float64 `json:"kill"`
						Play float64 `json:"play"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"paintball_expert"`
			SerialKiller struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Quake     float64 `json:"quake"`
						Megawalls float64 `json:"megawalls"`
						Blitz     float64 `json:"blitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"serial_killer"`
			SpaceMission struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"space_mission"`
			TntAddict struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_addict"`
			Waller struct {
				Active struct {
					Objectives struct {
						Win float64 `json:"win"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"waller"`
			WarlordsCtf struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_ctf"`
			WarlordsDedication struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						WarlordsWeeklyDedi float64 `json:"warlords_weekly_dedi"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_dedication"`
			WarlordsDomination struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_domination"`
			WarlordsWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_win"`
			WarriorsJourney struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Vampirezkillvamps float64 `json:"vampirezkillvamps"`
						Quake25Kill       float64 `json:"quake25kill"`
						Megawallswin      float64 `json:"megawallswin"`
						Blitzkill         float64 `json:"blitzkill"`
						Vampirezkillhuman float64 `json:"vampirezkillhuman"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warriors_journey"`
			WelcomeToHell struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Megawalls float64 `json:"megawalls"`
						Blitz     float64 `json:"blitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"welcome_to_hell"`
			GingerbreadMastery struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_mastery"`
			GingerbreadRacer struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_racer"`
			GingerbreadMaps struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_maps"`
			GingerbreadBlingBling struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_bling_bling"`
			SkywarsWeeklyKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						SkywarsWeeklyKills float64 `json:"skywars_weekly_kills"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_weekly_kills"`
			SkywarsTeamKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_team_kills"`
			SkywarsTeamWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_team_win"`
			SkywarsSoloKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						SkywarsSoloKills float64 `json:"skywars_solo_kills"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_solo_kills"`
			SkywarsSoloWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_solo_win"`
			UhcWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						UhcKills float64 `json:"uhc_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_weekly"`
			UhcDaily struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"uhc_daily"`
			BlitzWeeklyMaster struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Killblitz10      float64 `json:"killblitz10"`
						BlitzGamesPlayed float64 `json:"blitz_games_played"`
						Winblitz         float64 `json:"winblitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_weekly_master"`
			BlitzKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Killblitz10 float64 `json:"killblitz10"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_kills"`
			BlitzWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_win"`
			BlitzGameOfTheDay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_game_of_the_day"`
			WarlordsTdm struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_tdm"`
			MegaWallsWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						MegaWallsKillWeekly float64 `json:"mega_walls_kill_weekly"`
						MegaWallsPlayWeekly float64 `json:"mega_walls_play_weekly"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_weekly"`
			MegaWallsKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						MegaWallsKillDaily float64 `json:"mega_walls_kill_daily"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_kill"`
			MegaWallsWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_win"`
			MegaWallsPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_play"`
			QuakeWeeklyPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"quake_weekly_play"`
			QuakeDailyKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_kill"`
			QuakeDailyPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						QuakeDailyPlay float64 `json:"quake_daily_play"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_play"`
			CrazyWallsDailyPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						CrazyWallsDailyPlay float64 `json:"crazy_walls_daily_play"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"crazy_walls_daily_play"`
			CrazyWallsDailyWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"crazy_walls_daily_win"`
			CrazyWallsDailyKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						CrazyWallsDailyKill float64 `json:"crazy_walls_daily_kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"crazy_walls_daily_kill"`
			CrazyWallsWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"crazy_walls_weekly"`
			VampirezDailyWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_win"`
			VampirezWeeklyKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						VampirezWeeklyKillVampire float64 `json:"vampirez_weekly_kill_vampire"`
						VampirezWeeklyKillZombie  float64 `json:"vampirez_weekly_kill_zombie"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_weekly_kill"`
			VampirezWeeklyWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						VampirezWeeklyWinSurvivor float64 `json:"vampirez_weekly_win_survivor"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"vampirez_weekly_win"`
			VampirezDailyKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						VampirezDailyKillVampire float64 `json:"vampirez_daily_kill_vampire"`
						VampirezDailyKillZombie  float64 `json:"vampirez_daily_kill_zombie"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_kill"`
			VampirezDailyPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_play"`
			TntWeeklyPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntWeeklyPlay float64 `json:"tnt_weekly_play"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_weekly_play"`
			TntDailyWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_daily_win"`
			TntDailyPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_daily_play"`
			SupersmashSoloWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_solo_win"`
			SupersmashWeeklyKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						SupersmashWeeklyKills float64 `json:"supersmash_weekly_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_weekly_kills"`
			SupersmashTeamKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"supersmash_team_kills"`
			SupersmashSoloKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_solo_kills"`
			SupersmashTeamWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"supersmash_team_win"`
			WallsWeekly struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						WallsWeeklyPlay  float64 `json:"walls_weekly_play"`
						WallsWeeklyKills float64 `json:"walls_weekly_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_weekly"`
			WallsDailyWin struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_win"`
			WallsDailyKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_kill"`
			WallsDailyPlay struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_play"`
			ArcadeSpecialist struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play float64 `json:"play"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"arcade_specialist"`
			ArcadeWinner struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"arcade_winner"`
			ArcadeGamer struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play float64 `json:"play"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"arcade_gamer"`
			Paintballer struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"paintballer"`
			PaintballKiller struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Kill float64 `json:"kill"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"paintball_killer"`
			ArenaDailyPlay struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_daily_play"`
			ArenaDailyWins struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_daily_wins"`
			ArenaWeeklyPlay struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_weekly_play"`
			ArenaDailyKills struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"arena_daily_kills"`
			NormalBrawler struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"normal_brawler"`
			UhcMadness struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Kill float64 `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_madness"`
			UhcAddict struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_addict"`
			InsaneBrawler struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"insane_brawler"`
			HuntingSeason struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"hunting_season"`
			CvcKillDailyNormal struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_kill_daily_normal"`
			CvcKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_kill"`
			CvcWinDailyDeathmatch struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_win_daily_deathmatch"`
			CvcWinDailyNormal struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_win_daily_normal"`
			CvcKillWeekly struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						CvcPlayWeekly  float64 `json:"cvc_play_weekly"`
						CvcPlayWeekly2 float64 `json:"cvc_play_weekly_2"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"cvc_kill_weekly"`
			SkyclashVoid struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_void"`
			SkyclashKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Kill float64 `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_kills"`
			SkyclashWeeklyKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Kill float64 `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_weekly_kills"`
			SkyclashPlayGames struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_play_games"`
			SkyclashPlayPoints struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skyclash_play_points"`
			UhcTeam struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						UhcGames float64 `json:"uhc_games"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_team"`
			UhcDm struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_dm"`
			UhcSolo struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_solo"`
			QuakeDailyWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_win"`
			TntPvprunWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntPvprunWeekly float64 `json:"tnt_pvprun_weekly"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_pvprun_weekly"`
			TntPvprunDaily struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntPvprunDaily float64 `json:"tnt_pvprun_daily"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_pvprun_daily"`
			TntTntrunDaily struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tntrun_daily"`
			TntBowspleefWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntBowspleefWeekly float64 `json:"tnt_bowspleef_weekly"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_bowspleef_weekly"`
			TntWizardsDaily struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_wizards_daily"`
			TntTnttagDaily struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tnttag_daily"`
			TntBowspleefDaily struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntBowspleefDaily float64 `json:"tnt_bowspleef_daily"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_bowspleef_daily"`
			TntTntrunWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntTntrunWeekly float64 `json:"tnt_tntrun_weekly"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tntrun_weekly"`
			TntWizardsWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntWizardsWeeklyKills float64 `json:"tnt_wizards_weekly_kills"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_wizards_weekly"`
			TntTnttagWeekly struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						TntTnttagWeekly float64 `json:"tnt_tnttag_weekly"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_tnttag_weekly"`
			BedwarsDailyOneMore struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsDailyPlayed float64 `json:"bedwars_daily_played"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_one_more"`
			BedwarsDailyWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_win"`
			BedwarsWeeklyBedElims struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						BedwarsBedElims float64 `json:"bedwars_bed_elims"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_weekly_bed_elims"`
			SkywarsWeeklyArcadeWinAll struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						SkywarsArcadeWeeklyWin float64 `json:"skywars_arcade_weekly_win"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_weekly_arcade_win_all"`
			SkywarsArcadeWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_arcade_win"`
			MmWeeklyWins struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						MmWeeklyWin float64 `json:"mm_weekly_win"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"mm_weekly_wins"`
			MmDailyTargetKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mm_daily_target_kill"`
			MmDailyWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"mm_daily_win"`
			MmWeeklyMurdererKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"mm_weekly_murderer_kills"`
			MmDailyPowerPlay struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"mm_daily_power_play"`
			BedwarsWeeklyPumpkinator struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_weekly_pumpkinator"`
			BedwarsWeeklySanta struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						BedwarsSpecialWeeklySanta float64 `json:"bedwars_special_weekly_santa"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_weekly_santa"`
			SkywarsSpecialNorthPole struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_special_north_pole"`
			BlitzSpecialDailyNorthPole struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_special_daily_north_pole"`
			UhcWeeklySpecialCookie struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"uhc_weekly_special_cookie"`
			MmSpecialWeeklySanta struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mm_special_weekly_santa"`
			SkywarsWeeklyFreeLootChest struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
			} `json:"skywars_weekly_free_loot_chest"`
			SkywarsDailyTokens struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_daily_tokens"`
			SkywarsWeeklyHardChest struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						SkywarsWeeklyHardLootChest float64 `json:"skywars_weekly_hard_loot_chest"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_weekly_hard_chest"`
			DuelsPlayer struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play float64 `json:"play"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"duels_player"`
			DuelsWinner struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"duels_winner"`
			DuelsKiller struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"duels_killer"`
			DuelsWeeklyWins struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"duels_weekly_wins"`
			DuelsWeeklyKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Kill float64 `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"duels_weekly_kills"`
			BuildBattleWeekly struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Play float64 `json:"play"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"build_battle_weekly"`
			BuildBattleWinner struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_winner"`
			BuildBattlePlayer struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play float64 `json:"play"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_player"`
			PrototypePitWeeklyGold struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						PrototypePitWeeklyGold float64 `json:"prototype_pit_weekly_gold"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_weekly_gold"`
			PrototypePitDailyContract struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_daily_contract"`
			PrototypePitDailyKills struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Kill float64 `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"prototype_pit_daily_kills"`
			SkywarsMegaDoublesWins struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_mega_doubles_wins"`
			BedwarsWeeklyDreamWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsDreamWins float64 `json:"bedwars_dream_wins"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_weekly_dream_win"`
			MegaWallsFaithful struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_faithful"`
			WarlordsAllStar struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						WarlordsWeeklyDamage float64 `json:"warlords_weekly_damage"`
						WarlordsWeeklyHeal   float64 `json:"warlords_weekly_heal"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_all_star"`
			WarlordsObjectives struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_objectives"`
			WarlordsVictorious struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_victorious"`
			BedwarsDailyNightmares struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_daily_nightmares"`
			SkywarsHalloweenHarvest struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						SkywarsHalloweenKills float64 `json:"skywars_halloween_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_halloween_harvest"`
			BedwarsDailyGifts struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsDailySpecialChristmasGifts float64 `json:"bedwars_daily_special_christmas_gifts"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_gifts"`
			SkywarsCorruptWin struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_corrupt_win"`
			BuildBattleChristmas struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"build_battle_christmas"`
			BlitzWeeklyChaosMaster struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_weekly_chaos_master"`
			BlitzWinChaos struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_win_chaos"`
			VampirezWeeklyHumanKill struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						VampirezWeeklyKillSurvivor float64 `json:"vampirez_weekly_kill_survivor"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"vampirez_weekly_human_kill"`
			VampirezDailyHumanKill struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						VampirezDailyKillHuman float64 `json:"vampirez_daily_kill_human"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_human_kill"`
			BlitzLootChestDaily struct {
				Completions []struct {
					Time float64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Lootchestblitz float64 `json:"lootchestblitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_loot_chest_daily"`
			BlitzLootChestWeekly struct {
				Active struct {
					Started    float64 `json:"started"`
					Objectives struct {
						Lootchestblitz  float64 `json:"lootchestblitz"`
						Dealdamageblitz float64 `json:"dealdamageblitz"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitz_loot_chest_weekly"`
			SoloBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"solo_brawler"`
			TeamBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"team_brawler"`
			SkywarsHalloweenHarvest2019 struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_halloween_harvest_2019"`
			BuildBattleHalloween struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_halloween"`
			MmSpecialWeeklyKillerInstinct2019 struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"mm_special_weekly_killer_instinct_2019"`
			TntWeeklySpecial struct {
				Active struct {
					Objectives struct {
						TntWeeklyCandyCanes float64 `json:"tnt_weekly_candy_canes"`
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_weekly_special"`
			BuildBattleChristmasWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_christmas_weekly"`
			TntLunarQuest struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_lunar_quest"`
			SkywarsLunarQuest struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started float64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_lunar_quest"`
		} `json:"quests"`
		SeeRequests bool `json:"seeRequests"`
		Stats       struct {
			Arcade struct {
				BountyKillsOneinthequiver                     float64 `json:"bounty_kills_oneinthequiver"`
				Coins                                         float64 `json:"coins"`
				DeathsOneinthequiver                          float64 `json:"deaths_oneinthequiver"`
				DeathsThrowOut                                float64 `json:"deaths_throw_out"`
				Flash                                         bool    `json:"flash"`
				HeadshotsDayone                               float64 `json:"headshots_dayone"`
				KillsDayone                                   float64 `json:"kills_dayone"`
				KillsDragonwars2                              float64 `json:"kills_dragonwars2"`
				KillsOneinthequiver                           float64 `json:"kills_oneinthequiver"`
				KillsThrowOut                                 float64 `json:"kills_throw_out"`
				MaxWave                                       float64 `json:"max_wave"`
				PoopCollected                                 float64 `json:"poop_collected"`
				StampLevel                                    float64 `json:"stamp_level"`
				TimeStamp                                     float64 `json:"time_stamp"`
				WinsDayone                                    float64 `json:"wins_dayone"`
				WinsDragonwars2                               float64 `json:"wins_dragonwars2"`
				WinsFarmHunt                                  float64 `json:"wins_farm_hunt"`
				WinsOneinthequiver                            float64 `json:"wins_oneinthequiver"`
				WinsParty                                     float64 `json:"wins_party"`
				WinsParty2                                    float64 `json:"wins_party_2"`
				SwKills                                       float64 `json:"sw_kills"`
				SwShotsFired                                  float64 `json:"sw_shots_fired"`
				SwWeeklyKillsB                                float64 `json:"sw_weekly_kills_b"`
				SwRebelKills                                  float64 `json:"sw_rebel_kills"`
				SwDeaths                                      float64 `json:"sw_deaths"`
				SwMonthlyKillsA                               float64 `json:"sw_monthly_kills_a"`
				WinsThrowOut                                  float64 `json:"wins_throw_out"`
				WinsEnder                                     float64 `json:"wins_ender"`
				WinsParty3                                    float64 `json:"wins_party_3"`
				HitwRecordQ                                   float64 `json:"hitw_record_q"`
				HitwRecordF                                   float64 `json:"hitw_record_f"`
				RoundsHoleInTheWall                           float64 `json:"rounds_hole_in_the_wall"`
				WinsHoleInTheWall                             float64 `json:"wins_hole_in_the_wall"`
				SwWeeklyKillsA                                float64 `json:"sw_weekly_kills_a"`
				SwEmpireKills                                 float64 `json:"sw_empire_kills"`
				SwGameWins                                    float64 `json:"sw_game_wins"`
				WinsSantaSays                                 float64 `json:"wins_santa_says"`
				RoundsSantaSays                               float64 `json:"rounds_santa_says"`
				RoundsSimonSays                               float64 `json:"rounds_simon_says"`
				WeeklyCoinsB                                  float64 `json:"weekly_coins_b"`
				WinsSimonSays                                 float64 `json:"wins_simon_says"`
				MonthlyCoinsA                                 float64 `json:"monthly_coins_a"`
				WeeklyCoinsA                                  float64 `json:"weekly_coins_a"`
				MonthlyCoinsB                                 float64 `json:"monthly_coins_b"`
				KillsMiniWalls                                float64 `json:"kills_mini_walls"`
				DeathsMiniWalls                               float64 `json:"deaths_mini_walls"`
				WitherKillsMiniWalls                          float64 `json:"wither_kills_mini_walls"`
				FinalKillsMiniWalls                           float64 `json:"final_kills_mini_walls"`
				PowerkicksSoccer                              float64 `json:"powerkicks_soccer"`
				GoalsSoccer                                   float64 `json:"goals_soccer"`
				WinsSoccer                                    float64 `json:"wins_soccer"`
				MiniwallsActiveKit                            string  `json:"miniwalls_activeKit"`
				ArrowsHitMiniWalls                            float64 `json:"arrows_hit_mini_walls"`
				WinsMiniWalls                                 float64 `json:"wins_mini_walls"`
				ArrowsShotMiniWalls                           float64 `json:"arrows_shot_mini_walls"`
				WitherDamageMiniWalls                         float64 `json:"wither_damage_mini_walls"`
				Dec2016Achievements                           bool    `json:"dec2016_achievements"`
				Dec2016Achievements2                          bool    `json:"dec2016_achievements2"`
				WinsDrawTheirThing                            float64 `json:"wins_draw_their_thing"`
				Hints                                         bool    `json:"hints"`
				KicksSoccer                                   float64 `json:"kicks_soccer"`
				WinsHypixelSports                             float64 `json:"wins_hypixel_sports"`
				FastestTime20ZombiesAlienarcadiumNormal       float64 `json:"fastest_time_20_zombies_alienarcadium_normal"`
				BestRoundZombiesAlienarcadium                 float64 `json:"best_round_zombies_alienarcadium"`
				BestRoundZombiesAlienarcadiumNormal           float64 `json:"best_round_zombies_alienarcadium_normal"`
				FastestTime20Zombies                          float64 `json:"fastest_time_20_zombies"`
				BestRoundZombies                              float64 `json:"best_round_zombies"`
				FastestTime10ZombiesAlienarcadiumNormal       float64 `json:"fastest_time_10_zombies_alienarcadium_normal"`
				FastestTime10Zombies                          float64 `json:"fastest_time_10_zombies"`
				RainbowZombieKillsZombies                     float64 `json:"rainbow_zombie_kills_zombies"`
				GiantZombieKillsZombies                       float64 `json:"giant_zombie_kills_zombies"`
				SpaceBlasterZombieKillsZombies                float64 `json:"space_blaster_zombie_kills_zombies"`
				ChglugluZombieKillsZombies                    float64 `json:"chgluglu_zombie_kills_zombies"`
				TimesKnockedDownZombiesAlienarcadium          float64 `json:"times_knocked_down_zombies_alienarcadium"`
				WindowsRepairedZombiesAlienarcadium           float64 `json:"windows_repaired_zombies_alienarcadium"`
				TimesKnockedDownZombies                       float64 `json:"times_knocked_down_zombies"`
				ZombieKillsZombiesAlienarcadiumNormal         float64 `json:"zombie_kills_zombies_alienarcadium_normal"`
				ZombieKillsZombies                            float64 `json:"zombie_kills_zombies"`
				GhastZombieKillsZombies                       float64 `json:"ghast_zombie_kills_zombies"`
				WormSmallZombieKillsZombies                   float64 `json:"worm_small_zombie_kills_zombies"`
				DeathsZombiesAlienarcadium                    float64 `json:"deaths_zombies_alienarcadium"`
				BulletsShotZombies                            float64 `json:"bullets_shot_zombies"`
				BasicZombieKillsZombies                       float64 `json:"basic_zombie_kills_zombies"`
				MegaMagmaZombieKillsZombies                   float64 `json:"mega_magma_zombie_kills_zombies"`
				WindowsRepairedZombiesAlienarcadiumNormal     float64 `json:"windows_repaired_zombies_alienarcadium_normal"`
				BulletsHitZombies                             float64 `json:"bullets_hit_zombies"`
				WormZombieKillsZombies                        float64 `json:"worm_zombie_kills_zombies"`
				IronGolemZombieKillsZombies                   float64 `json:"iron_golem_zombie_kills_zombies"`
				HeadshotsZombies                              float64 `json:"headshots_zombies"`
				PlayersRevivedZombies                         float64 `json:"players_revived_zombies"`
				TotalRoundsSurvivedZombiesAlienarcadiumNormal float64 `json:"total_rounds_survived_zombies_alienarcadium_normal"`
				DeathsZombies                                 float64 `json:"deaths_zombies"`
				TimesKnockedDownZombiesAlienarcadiumNormal    float64 `json:"times_knocked_down_zombies_alienarcadium_normal"`
				TotalRoundsSurvivedZombiesAlienarcadium       float64 `json:"total_rounds_survived_zombies_alienarcadium"`
				WindowsRepairedZombies                        float64 `json:"windows_repaired_zombies"`
				TotalRoundsSurvivedZombies                    float64 `json:"total_rounds_survived_zombies"`
				PlayersRevivedZombiesAlienarcadium            float64 `json:"players_revived_zombies_alienarcadium"`
				SentinelZombieKillsZombies                    float64 `json:"sentinel_zombie_kills_zombies"`
				ZombieKillsZombiesAlienarcadium               float64 `json:"zombie_kills_zombies_alienarcadium"`
				PigZombieZombieKillsZombies                   float64 `json:"pig_zombie_zombie_kills_zombies"`
				DeathsZombiesAlienarcadiumNormal              float64 `json:"deaths_zombies_alienarcadium_normal"`
				SpaceGruntZombieKillsZombies                  float64 `json:"space_grunt_zombie_kills_zombies"`
				SkeletonZombieKillsZombies                    float64 `json:"skeleton_zombie_kills_zombies"`
				BlobZombieKillsZombies                        float64 `json:"blob_zombie_kills_zombies"`
				ClownZombieKillsZombies                       float64 `json:"clown_zombie_kills_zombies"`
				PlayersRevivedZombiesAlienarcadiumNormal      float64 `json:"players_revived_zombies_alienarcadium_normal"`
				MegaBlobZombieKillsZombies                    float64 `json:"mega_blob_zombie_kills_zombies"`
				DeliveredSantaSimulator                       float64 `json:"delivered_santa_simulator"`
				SpottedSantaSimulator                         float64 `json:"spotted_santa_simulator"`
				EggsFoundEasterSimulator                      float64 `json:"eggs_found_easter_simulator"`
				WinsEasterSimulator                           float64 `json:"wins_easter_simulator"`
				ItemsFoundScubaSimulator                      float64 `json:"items_found_scuba_simulator"`
				TotalPointsScubaSimulator                     float64 `json:"total_points_scuba_simulator"`
				WinsScubaSimulator                            float64 `json:"wins_scuba_simulator"`
				WinsGrinchSimulatorV2                         float64 `json:"wins_grinch_simulator_v2"`
				GiftsGrinchSimulatorV2                        float64 `json:"gifts_grinch_simulator_v2"`
				LastTourneyAd                                 float64 `json:"lastTourneyAd"`
			} `json:"Arcade"`
			Arena struct {
				ActiveRune      string   `json:"active_rune"`
				Coins           float64  `json:"coins"`
				CoinsSpent      float64  `json:"coins_spent"`
				Damage2V2       float64  `json:"damage_2v2"`
				Damage4V4       float64  `json:"damage_4v4"`
				DamageFfa       float64  `json:"damage_ffa"`
				Deaths2V2       float64  `json:"deaths_2v2"`
				Deaths4V4       float64  `json:"deaths_4v4"`
				Games2V2        float64  `json:"games_2v2"`
				Games4V4        float64  `json:"games_4v4"`
				GamesFfa        float64  `json:"games_ffa"`
				Hat             string   `json:"hat"`
				Healed2V2       float64  `json:"healed_2v2"`
				Healed4V4       float64  `json:"healed_4v4"`
				HealedFfa       float64  `json:"healed_ffa"`
				Keys            float64  `json:"keys"`
				Kills2V2        float64  `json:"kills_2v2"`
				Kills4V4        float64  `json:"kills_4v4"`
				KillsFfa        float64  `json:"kills_ffa"`
				Losses2V2       float64  `json:"losses_2v2"`
				Losses4V4       float64  `json:"losses_4v4"`
				LvlCooldown     float64  `json:"lvl_cooldown"`
				LvlDamage       float64  `json:"lvl_damage"`
				LvlEnergy       float64  `json:"lvl_energy"`
				LvlHealth       float64  `json:"lvl_health"`
				MagicalChest    float64  `json:"magical_chest"`
				Offensive       string   `json:"offensive"`
				Packages        []string `json:"packages"`
				Rating          float64  `json:"rating"`
				Support         string   `json:"support"`
				Utility         string   `json:"utility"`
				WinStreaks2V2   float64  `json:"win_streaks_2v2"`
				WinStreaks4V4   float64  `json:"win_streaks_4v4"`
				WinStreaksFfa   float64  `json:"win_streaks_ffa"`
				Wins2V2         float64  `json:"wins_2v2"`
				Wins4V4         float64  `json:"wins_4v4"`
				WinsFfa         float64  `json:"wins_ffa"`
				RuneLevelEnergy float64  `json:"rune_level_energy"`
				SelectedSword   string   `json:"selected_sword"`
				Ultimate        string   `json:"ultimate"`
				RuneLevelDamage float64  `json:"rune_level_damage"`
				Wins            float64  `json:"wins"`
			} `json:"Arena"`
			Battleground struct {
				Assists                  float64  `json:"assists"`
				BerserkerPlays           float64  `json:"berserker_plays"`
				BrokenInventory          float64  `json:"broken_inventory"`
				ChosenClass              string   `json:"chosen_class"`
				Coins                    float64  `json:"coins"`
				CurrentWeapon            float64  `json:"current_weapon"`
				Damage                   float64  `json:"damage"`
				DamageBerserker          float64  `json:"damage_berserker"`
				DamageDefender           float64  `json:"damage_defender"`
				DamagePrevented          float64  `json:"damage_prevented"`
				DamagePreventedBerserker float64  `json:"damage_prevented_berserker"`
				DamagePreventedDefender  float64  `json:"damage_prevented_defender"`
				DamagePreventedWarrior   float64  `json:"damage_prevented_warrior"`
				DamageTaken              float64  `json:"damage_taken"`
				DamageWarrior            float64  `json:"damage_warrior"`
				Deaths                   float64  `json:"deaths"`
				DefenderPlays            float64  `json:"defender_plays"`
				FlagConquerSelf          float64  `json:"flag_conquer_self"`
				FlagConquerTeam          float64  `json:"flag_conquer_team"`
				Heal                     float64  `json:"heal"`
				HealBerserker            float64  `json:"heal_berserker"`
				HealDefender             float64  `json:"heal_defender"`
				HealWarrior              float64  `json:"heal_warrior"`
				Kills                    float64  `json:"kills"`
				LifeLeeched              float64  `json:"life_leeched"`
				LifeLeechedBerserker     float64  `json:"life_leeched_berserker"`
				LifeLeechedWarrior       float64  `json:"life_leeched_warrior"`
				Losses                   float64  `json:"losses"`
				LossesDefender           float64  `json:"losses_defender"`
				LossesWarrior            float64  `json:"losses_warrior"`
				MageArmorSelection       float64  `json:"mage_armor_selection"`
				MageSpec                 string   `json:"mage_spec"`
				MagicDust                float64  `json:"magic_dust"`
				Packages                 []string `json:"packages"`
				PaladinSpec              string   `json:"paladin_spec"`
				Repaired                 float64  `json:"repaired"`
				RepairedCommon           float64  `json:"repaired_common"`
				RepairedEpic             float64  `json:"repaired_epic"`
				RepairedRare             float64  `json:"repaired_rare"`
				SelectedMount            string   `json:"selected_mount"`
				UpgradeRandom            float64  `json:"upgrade_random"`
				UpgradeRandomEpic        float64  `json:"upgrade_random_epic"`
				VoidShards               float64  `json:"void_shards"`
				WarriorArmorSelection    float64  `json:"warrior_armor_selection"`
				WarriorCooldown          float64  `json:"warrior_cooldown"`
				WarriorCritchance        float64  `json:"warrior_critchance"`
				WarriorCritmultiplier    float64  `json:"warrior_critmultiplier"`
				WarriorEnergy            float64  `json:"warrior_energy"`
				WarriorHealth            float64  `json:"warrior_health"`
				WarriorHelmetSelection   float64  `json:"warrior_helmet_selection"`
				WarriorPlays             float64  `json:"warrior_plays"`
				WarriorSkill1            float64  `json:"warrior_skill1"`
				WarriorSkill2            float64  `json:"warrior_skill2"`
				WarriorSkill3            float64  `json:"warrior_skill3"`
				WarriorSkill4            float64  `json:"warrior_skill4"`
				WarriorSkill5            float64  `json:"warrior_skill5"`
				WarriorSpec              string   `json:"warrior_spec"`
				WeaponInventory          []struct {
					Spec struct {
						Spec        float64 `json:"spec"`
						PlayerClass float64 `json:"playerClass"`
					} `json:"spec"`
					Ability      float64 `json:"ability"`
					AbilityBoost float64 `json:"abilityBoost"`
					Damage       float64 `json:"damage"`
					Energy       float64 `json:"energy"`
					Chance       float64 `json:"chance"`
					Multiplier   float64 `json:"multiplier"`
					Health       float64 `json:"health"`
					Cooldown     float64 `json:"cooldown"`
					Movement     float64 `json:"movement"`
					Material     string  `json:"material"`
					ID           float64 `json:"id"`
					Category     string  `json:"category"`
					Crafted      bool    `json:"crafted"`
					UpgradeMax   float64 `json:"upgradeMax"`
					UpgradeTimes float64 `json:"upgradeTimes"`
					PlayStreak   bool    `json:"playStreak,omitempty"`
				} `json:"weapon_inventory"`
				WinStreak                  float64 `json:"win_streak"`
				Wins                       float64 `json:"wins"`
				WinsBerserker              float64 `json:"wins_berserker"`
				WinsBlu                    float64 `json:"wins_blu"`
				WinsCapturetheflag         float64 `json:"wins_capturetheflag"`
				WinsCapturetheflagBlu      float64 `json:"wins_capturetheflag_blu"`
				WinsCapturetheflagRed      float64 `json:"wins_capturetheflag_red"`
				WinsDefender               float64 `json:"wins_defender"`
				WinsDomination             float64 `json:"wins_domination"`
				WinsDominationBlu          float64 `json:"wins_domination_blu"`
				WinsDominationRed          float64 `json:"wins_domination_red"`
				WinsRed                    float64 `json:"wins_red"`
				WinsWarrior                float64 `json:"wins_warrior"`
				SalvagedDustReward         float64 `json:"salvaged_dust_reward"`
				SalvagedWeaponsRare        float64 `json:"salvaged_weapons_rare"`
				SalvagedWeapons            float64 `json:"salvaged_weapons"`
				SalvagedShardsReward       float64 `json:"salvaged_shards_reward"`
				SalvagedWeaponsCommon      float64 `json:"salvaged_weapons_common"`
				HealMage                   float64 `json:"heal_mage"`
				DamagePyromancer           float64 `json:"damage_pyromancer"`
				DamagePreventedMage        float64 `json:"damage_prevented_mage"`
				HealPyromancer             float64 `json:"heal_pyromancer"`
				DamagePreventedPyromancer  float64 `json:"damage_prevented_pyromancer"`
				PyromancerPlays            float64 `json:"pyromancer_plays"`
				DamageMage                 float64 `json:"damage_mage"`
				MagePlays                  float64 `json:"mage_plays"`
				SalvagedWeaponsEpic        float64 `json:"salvaged_weapons_epic"`
				LossesBerserker            float64 `json:"losses_berserker"`
				Crafted                    float64 `json:"crafted"`
				CraftedEpic                float64 `json:"crafted_epic"`
				ShamanSpec                 string  `json:"shaman_spec"`
				PlayStreak                 float64 `json:"play_streak"`
				DamagePreventedCryomancer  float64 `json:"damage_prevented_cryomancer"`
				HealCryomancer             float64 `json:"heal_cryomancer"`
				WinsMage                   float64 `json:"wins_mage"`
				DamageCryomancer           float64 `json:"damage_cryomancer"`
				CryomancerPlays            float64 `json:"cryomancer_plays"`
				WinsCryomancer             float64 `json:"wins_cryomancer"`
				WinsTeamdeathmatch         float64 `json:"wins_teamdeathmatch"`
				WinsTeamdeathmatchRed      float64 `json:"wins_teamdeathmatch_red"`
				WinsTeamdeathmatchA        float64 `json:"wins_teamdeathmatch_a"`
				LossesCryomancer           float64 `json:"losses_cryomancer"`
				LossesMage                 float64 `json:"losses_mage"`
				WinsTeamdeathmatchBlu      float64 `json:"wins_teamdeathmatch_blu"`
				WinsTeamdeathmatchB        float64 `json:"wins_teamdeathmatch_b"`
				RepairedLegendary          float64 `json:"repaired_legendary"`
				SalvagedWeaponsLegendary   float64 `json:"salvaged_weapons_legendary"`
				WinsPyromancer             float64 `json:"wins_pyromancer"`
				LossesPyromancer           float64 `json:"losses_pyromancer"`
				EarthwardenPlays           float64 `json:"earthwarden_plays"`
				HealEarthwarden            float64 `json:"heal_earthwarden"`
				HealShaman                 float64 `json:"heal_shaman"`
				WinsShaman                 float64 `json:"wins_shaman"`
				DamagePreventedEarthwarden float64 `json:"damage_prevented_earthwarden"`
				WinsEarthwarden            float64 `json:"wins_earthwarden"`
				DamageShaman               float64 `json:"damage_shaman"`
				DamageEarthwarden          float64 `json:"damage_earthwarden"`
				DamagePreventedShaman      float64 `json:"damage_prevented_shaman"`
				ShamanPlays                float64 `json:"shaman_plays"`
				Penalty                    float64 `json:"penalty"`
				AfkWarned                  float64 `json:"afk_warned"`
				MageHelmetSelection        float64 `json:"mage_helmet_selection"`
				PaladinArmorSelection      float64 `json:"paladin_armor_selection"`
				PaladinHelmetSelection     float64 `json:"paladin_helmet_selection"`
				ShamanArmorSelection       float64 `json:"shaman_armor_selection"`
				ShamanHelmetSelection      float64 `json:"shaman_helmet_selection"`
				DamageAquamancer           float64 `json:"damage_aquamancer"`
				HealAquamancer             float64 `json:"heal_aquamancer"`
				AquamancerPlays            float64 `json:"aquamancer_plays"`
				DamagePreventedAquamancer  float64 `json:"damage_prevented_aquamancer"`
				LossesAquamancer           float64 `json:"losses_aquamancer"`
				WinsAquamancer             float64 `json:"wins_aquamancer"`
				DamagePreventedProtector   float64 `json:"damage_prevented_protector"`
				ProtectorPlays             float64 `json:"protector_plays"`
				DamagePreventedPaladin     float64 `json:"damage_prevented_paladin"`
				WinsProtector              float64 `json:"wins_protector"`
				HealProtector              float64 `json:"heal_protector"`
				WinsPaladin                float64 `json:"wins_paladin"`
				PaladinPlays               float64 `json:"paladin_plays"`
				HealPaladin                float64 `json:"heal_paladin"`
				DamageProtector            float64 `json:"damage_protector"`
				DamagePaladin              float64 `json:"damage_paladin"`
				AvengerPlays               float64 `json:"avenger_plays"`
				WinsAvenger                float64 `json:"wins_avenger"`
				HealAvenger                float64 `json:"heal_avenger"`
				DamageAvenger              float64 `json:"damage_avenger"`
				DamagePreventedAvenger     float64 `json:"damage_prevented_avenger"`
				LossesAvenger              float64 `json:"losses_avenger"`
				LossesPaladin              float64 `json:"losses_paladin"`
				CraftedRare                float64 `json:"crafted_rare"`
				CrusaderPlays              float64 `json:"crusader_plays"`
				HealCrusader               float64 `json:"heal_crusader"`
				WinsCrusader               float64 `json:"wins_crusader"`
				DamageCrusader             float64 `json:"damage_crusader"`
				DamagePreventedCrusader    float64 `json:"damage_prevented_crusader"`
				WinsDominationB            float64 `json:"wins_domination_b"`
				WinsCapturetheflagB        float64 `json:"wins_capturetheflag_b"`
				MageSkill5                 float64 `json:"mage_skill5"`
				MageSkill4                 float64 `json:"mage_skill4"`
				MageSkill3                 float64 `json:"mage_skill3"`
				MageSkill2                 float64 `json:"mage_skill2"`
				MageSkill1                 float64 `json:"mage_skill1"`
				MageCritmultiplier         float64 `json:"mage_critmultiplier"`
				MageCritchance             float64 `json:"mage_critchance"`
				MageCooldown               float64 `json:"mage_cooldown"`
				MageEnergy                 float64 `json:"mage_energy"`
				MageHealth                 float64 `json:"mage_health"`
				CraftedLegendary           float64 `json:"crafted_legendary"`
				RerollLegendary            float64 `json:"reroll_legendary"`
				Reroll                     float64 `json:"reroll"`
				LossesCrusader             float64 `json:"losses_crusader"`
				DamagePreventedThunderlord float64 `json:"damage_prevented_thunderlord"`
				HealThunderlord            float64 `json:"heal_thunderlord"`
				WinsThunderlord            float64 `json:"wins_thunderlord"`
				ThunderlordPlays           float64 `json:"thunderlord_plays"`
				DamageThunderlord          float64 `json:"damage_thunderlord"`
				UpgradeCraftedLegendary    float64 `json:"upgrade_crafted_legendary"`
				UpgradeCrafted             float64 `json:"upgrade_crafted"`
				WinsCapturetheflagA        float64 `json:"wins_capturetheflag_a"`
				WinsDominationA            float64 `json:"wins_domination_a"`
				RewardInventory            float64 `json:"reward_inventory"`
				ShamanSkill5               float64 `json:"shaman_skill5"`
				ShamanSkill4               float64 `json:"shaman_skill4"`
				ShamanSkill3               float64 `json:"shaman_skill3"`
				ShamanSkill2               float64 `json:"shaman_skill2"`
				ShamanSkill1               float64 `json:"shaman_skill1"`
				ShamanCritmultiplier       float64 `json:"shaman_critmultiplier"`
				ShamanCooldown             float64 `json:"shaman_cooldown"`
				ShamanEnergy               float64 `json:"shaman_energy"`
				ShamanHealth               float64 `json:"shaman_health"`
				ShamanCritchance           float64 `json:"shaman_critchance"`
				PaladinSkill5              float64 `json:"paladin_skill5"`
				PaladinSkill4              float64 `json:"paladin_skill4"`
				PaladinSkill3              float64 `json:"paladin_skill3"`
				PaladinSkill2              float64 `json:"paladin_skill2"`
				PaladinSkill1              float64 `json:"paladin_skill1"`
				PaladinCritmultiplier      float64 `json:"paladin_critmultiplier"`
				PaladinCritchance          float64 `json:"paladin_critchance"`
				PaladinCooldown            float64 `json:"paladin_cooldown"`
				PaladinEnergy              float64 `json:"paladin_energy"`
				PaladinHealth              float64 `json:"paladin_health"`
				KillsDomination            float64 `json:"kills_domination"`
				DomPointCaptures           float64 `json:"dom_point_captures"`
				TotalDominationScore       float64 `json:"total_domination_score"`
				PowerupsCollected          float64 `json:"powerups_collected"`
				DomPointDefends            float64 `json:"dom_point_defends"`
				KillsTeamdeathmatch        float64 `json:"kills_teamdeathmatch"`
				KillsCapturetheflag        float64 `json:"kills_capturetheflag"`
				FlagReturns                float64 `json:"flag_returns"`
				MvpCount                   float64 `json:"mvp_count"`
				LossesThunderlord          float64 `json:"losses_thunderlord"`
				LossesShaman               float64 `json:"losses_shaman"`
				Autostrikemode             bool    `json:"autostrikemode"`
				EnergyPowerups             bool    `json:"energyPowerups"`
				Hints                      bool    `json:"hints"`
				BoundWeapon                struct {
					Warrior struct {
					} `json:"warrior"`
				} `json:"bound_weapon"`
				Simplifiedresourcepack bool `json:"simplifiedresourcepack"`
			} `json:"Battleground"`
			HungerGames struct {
				Arachnologist      float64  `json:"arachnologist"`
				Archer             float64  `json:"archer"`
				Armorer            float64  `json:"armorer"`
				Astronaut          float64  `json:"astronaut"`
				Aura               string   `json:"aura"`
				Baker              float64  `json:"baker"`
				Blaze              float64  `json:"blaze"`
				Blood              bool     `json:"blood"`
				ChosenTaunt        string   `json:"chosen_taunt"`
				Coins              float64  `json:"coins"`
				Creepertamer       float64  `json:"creepertamer"`
				Deaths             float64  `json:"deaths"`
				Fisherman          float64  `json:"fisherman"`
				Horsetamer         float64  `json:"horsetamer"`
				Kills              float64  `json:"kills"`
				Knight             float64  `json:"knight"`
				Meatmaster         float64  `json:"meatmaster"`
				Necromancer        float64  `json:"necromancer"`
				Packages           []string `json:"packages"`
				Reddragon          float64  `json:"reddragon"`
				Rogue              float64  `json:"rogue"`
				Scout              float64  `json:"scout"`
				Slimeyslime        float64  `json:"slimeyslime"`
				Snowman            float64  `json:"snowman"`
				Speleologist       float64  `json:"speleologist"`
				Tim                float64  `json:"tim"`
				Toxicologist       float64  `json:"toxicologist"`
				Troll              float64  `json:"troll"`
				Wins               float64  `json:"wins"`
				Wolftamer          float64  `json:"wolftamer"`
				WinsTeams          float64  `json:"wins_teams"`
				WeeklyKillsA       float64  `json:"weekly_kills_a"`
				MonthlyKillsB      float64  `json:"monthly_kills_b"`
				RamboWin           float64  `json:"rambo_win"`
				Paladin            float64  `json:"paladin"`
				WeeklyKillsB       float64  `json:"weekly_kills_b"`
				MonthlyKillsA      float64  `json:"monthly_kills_a"`
				VotesMiradorBasin  float64  `json:"votes_Mirador Basin"`
				Hunter             float64  `json:"hunter"`
				RamboWins          float64  `json:"rambo_wins"`
				ChosenVictorydance string   `json:"chosen_victorydance"`
				HunterInventory    struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"HunterInventory"`
				ArmorerInventory struct {
					Num8 string `json:"8"`
				} `json:"ArmorerInventory"`
				AstronautInventory struct {
					Num0 string `json:"0"`
					Num2 string `json:"2"`
					Num8 string `json:"8"`
				} `json:"AstronautInventory"`
				Defaultkit      string `json:"defaultkit"`
				ArcherInventory struct {
					Num1 string `json:"1"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"ArcherInventory"`
				HorsetamerInventory struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num4 string `json:"4"`
					Num8 string `json:"8"`
				} `json:"HorsetamerInventory"`
				ChosenFinisher     string  `json:"chosen_finisher"`
				HypeTrain          float64 `json:"hype train"`
				HypeTrainInventory struct {
					Num0  string `json:"0"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num13 string `json:"13"`
					Num14 string `json:"14"`
					Num15 string `json:"15"`
					Num16 string `json:"16"`
				} `json:"Hype TrainInventory"`
				Afterkill              string  `json:"afterkill"`
				KitItemRenameHunterBow string  `json:"kit_item_rename_Hunter_bow"`
				LastTourneyAd          float64 `json:"lastTourneyAd"`
				WinsBackup             float64 `json:"wins_backup"`
				WinsTeamsNormal        float64 `json:"wins_teams_normal"`
				WinsSoloNormal         float64 `json:"wins_solo_normal"`
				Autoarmor              bool    `json:"autoarmor"`
				NecromancerInventory   struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num10 string `json:"10"`
					Num11 string `json:"11"`
					Num12 string `json:"12"`
				} `json:"NecromancerInventory"`
				ArrowsFiredNecromancer  float64 `json:"arrows_fired_necromancer"`
				ChestsOpenedNecromancer float64 `json:"chests_opened_necromancer"`
				Damage                  float64 `json:"damage"`
				ExpNecromancer          float64 `json:"exp_necromancer"`
				MobsSpawnedNecromancer  float64 `json:"mobs_spawned_necromancer"`
				ArrowsHit               float64 `json:"arrows_hit"`
				DamageNecromancer       float64 `json:"damage_necromancer"`
				GamesPlayedNecromancer  float64 `json:"games_played_necromancer"`
				TimePlayedNecromancer   float64 `json:"time_played_necromancer"`
				ArrowsHitNecromancer    float64 `json:"arrows_hit_necromancer"`
				KillsSoloNormal         float64 `json:"kills_solo_normal"`
				PotionsDrunk            float64 `json:"potions_drunk"`
				ArrowsFired             float64 `json:"arrows_fired"`
				BlitzUses               float64 `json:"blitz_uses"`
				MobsSpawned             float64 `json:"mobs_spawned"`
				TimePlayed              float64 `json:"time_played"`
				ChestsOpened            float64 `json:"chests_opened"`
				DamageTakenNecromancer  float64 `json:"damage_taken_necromancer"`
				DamageTaken             float64 `json:"damage_taken"`
				PotionsDrunkNecromancer float64 `json:"potions_drunk_necromancer"`
				KillsNecromancer        float64 `json:"kills_necromancer"`
				GamesPlayed             float64 `json:"games_played"`
				DamageTakenHunter       float64 `json:"damage_taken_hunter"`
				PotionsThrown           float64 `json:"potions_thrown"`
				KillsHunter             float64 `json:"kills_hunter"`
				DamageHunter            float64 `json:"damage_hunter"`
				GamesPlayedHunter       float64 `json:"games_played_hunter"`
				TimePlayedHunter        float64 `json:"time_played_hunter"`
				ChestsOpenedHunter      float64 `json:"chests_opened_hunter"`
				ExpHunter               float64 `json:"exp_hunter"`
				PotionsThrownHunter     float64 `json:"potions_thrown_hunter"`
				PotionsDrunkHunter      float64 `json:"potions_drunk_hunter"`
				ArrowsFiredHunter       float64 `json:"arrows_fired_hunter"`
				MobsSpawnedHunter       float64 `json:"mobs_spawned_hunter"`
				KillsTeamsNormal        float64 `json:"kills_teams_normal"`
			} `json:"HungerGames"`
			Mcgo struct {
				BombsDefused                float64  `json:"bombs_defused"`
				BombsPlanted                float64  `json:"bombs_planted"`
				CarbineCostReduction        float64  `json:"carbine_cost_reduction"`
				CarbineDamageIncrease       float64  `json:"carbine_damage_increase"`
				CarbineRecoilReduction      float64  `json:"carbine_recoil_reduction"`
				CarbineReloadSpeedReduction float64  `json:"carbine_reload_speed_reduction"`
				Coins                       float64  `json:"coins"`
				CopKills                    float64  `json:"cop_kills"`
				CriminalKills               float64  `json:"criminal_kills"`
				GameWins                    float64  `json:"game_wins"`
				HeadshotKills               float64  `json:"headshot_kills"`
				Kills                       float64  `json:"kills"`
				Kills102014                 float64  `json:"kills_10_2014"`
				Kills3102014                float64  `json:"kills_3_10_2014"`
				KnifeDamageIncrease         float64  `json:"knife_damage_increase"`
				MagnumCostReduction         float64  `json:"magnum_cost_reduction"`
				MagnumReloadSpeedReduction  float64  `json:"magnum_reload_speed_reduction"`
				Packages                    []string `json:"packages"`
				PistolDamageIncrease        float64  `json:"pistol_damage_increase"`
				RifleCostReduction          float64  `json:"rifle_cost_reduction"`
				RifleDamageIncrease         float64  `json:"rifle_damage_increase"`
				RifleRecoilReduction        float64  `json:"rifle_recoil_reduction"`
				RifleReloadSpeedReduction   float64  `json:"rifle_reload_speed_reduction"`
				RoundWins                   float64  `json:"round_wins"`
				SelectedCarbineDev          string   `json:"selectedCarbineDev"`
				SelectedMagnumDev           string   `json:"selectedMagnumDev"`
				SelectedOcelotChestplateDev string   `json:"selectedOcelotChestplateDev"`
				SelectedOcelotHelmetDev     string   `json:"selectedOcelotHelmetDev"`
				SelectedSmgDev              string   `json:"selectedSmgDev"`
				ShotgunDamageIncrease       float64  `json:"shotgun_damage_increase"`
				ShotgunRecoilReduction      float64  `json:"shotgun_recoil_reduction"`
				ShotsFired                  float64  `json:"shots_fired"`
				Deaths                      float64  `json:"deaths"`
				BountyHunter                float64  `json:"bounty_hunter"`
				SelectedRifleDev            string   `json:"selectedRifleDev"`
				MagnumRecoilReduction       float64  `json:"magnum_recoil_reduction"`
				PistolRecoilReduction       float64  `json:"pistol_recoil_reduction"`
				Mcgo                        struct {
					Points float64 `json:"points"`
				} `json:"mcgo"`
				WeeklyKillsB                 float64 `json:"weekly_kills_b"`
				MonthlyKillsA                float64 `json:"monthly_kills_a"`
				KnifeAttackDelay             float64 `json:"knife_attack_delay"`
				SelectedKnifeDev             string  `json:"selectedKnifeDev"`
				StrengthTraining             float64 `json:"strength_training"`
				PistolReloadSpeedReduction   float64 `json:"pistol_reload_speed_reduction"`
				BodyArmorCost                float64 `json:"body_armor_cost"`
				PocketChange                 float64 `json:"pocket_change"`
				GameWinsAlleyway             float64 `json:"game_wins_alleyway"`
				MagnumDamageIncrease         float64 `json:"magnum_damage_increase"`
				GameWinsDeathmatch           float64 `json:"game_wins_deathmatch"`
				KillsDeathmatch              float64 `json:"kills_deathmatch"`
				DeathsDeathmatch             float64 `json:"deaths_deathmatch"`
				SmgCostReduction             float64 `json:"smg_cost_reduction"`
				SmgReloadSpeedReduction      float64 `json:"smg_reload_speed_reduction"`
				ShotgunCostReduction         float64 `json:"shotgun_cost_reduction"`
				ShotgunReloadSpeedReduction  float64 `json:"shotgun_reload_speed_reduction"`
				SelectedCreeperHelmetDev     string  `json:"selectedCreeperHelmetDev"`
				SelectedCreeperChestplateDev string  `json:"selectedCreeperChestplateDev"`
				GameWinsTemple               float64 `json:"game_wins_temple"`
				GameWinsReserve              float64 `json:"game_wins_reserve"`
				GameWinsCarrier              float64 `json:"game_wins_carrier"`
				CriminalKillsDeathmatch      float64 `json:"criminal_kills_deathmatch"`
				CopKillsDeathmatch           float64 `json:"cop_kills_deathmatch"`
				GrenadeKills                 float64 `json:"grenade_kills"`
				LastTourneyAd                float64 `json:"lastTourneyAd"`
			} `json:"MCGO"`
			Paintball struct {
				Coins          float64  `json:"coins"`
				Deaths         float64  `json:"deaths"`
				Fortune        float64  `json:"fortune"`
				Godfather      float64  `json:"godfather"`
				Hat            string   `json:"hat"`
				Kills          float64  `json:"kills"`
				Packages       []string `json:"packages"`
				ShotsFired     float64  `json:"shots_fired"`
				Superluck      float64  `json:"superluck"`
				Wins           float64  `json:"wins"`
				Killstreaks    float64  `json:"killstreaks"`
				Adrenaline     float64  `json:"adrenaline"`
				Endurance      float64  `json:"endurance"`
				Transfusion    float64  `json:"transfusion"`
				WeeklyKillsB   float64  `json:"weekly_kills_b"`
				MonthlyKillsB  float64  `json:"monthly_kills_b"`
				WeeklyKillsA   float64  `json:"weekly_kills_a"`
				MonthlyKillsA  float64  `json:"monthly_kills_a"`
				ShowKillPrefix bool     `json:"showKillPrefix"`
			} `json:"Paintball"`
			Quake struct {
				Barrel                             string   `json:"barrel"`
				Coins                              float64  `json:"coins"`
				Deaths                             float64  `json:"deaths"`
				Hat                                string   `json:"hat"`
				Kills                              float64  `json:"kills"`
				Killstreaks                        float64  `json:"killstreaks"`
				Packages                           []string `json:"packages"`
				Trigger                            string   `json:"trigger"`
				Wins                               float64  `json:"wins"`
				Case                               string   `json:"case"`
				KillsTeams                         float64  `json:"kills_teams"`
				KillstreaksTeams                   float64  `json:"killstreaks_teams"`
				DeathsTeams                        float64  `json:"deaths_teams"`
				WinsTeams                          float64  `json:"wins_teams"`
				Muzzle                             string   `json:"muzzle"`
				Sight                              string   `json:"sight"`
				InstantRespawn                     bool     `json:"instantRespawn"`
				Armor                              string   `json:"armor"`
				Null                               string   `json:"null"`
				Killsound                          string   `json:"killsound"`
				MonthlyKillsA                      float64  `json:"monthly_kills_a"`
				WeeklyKillsA                       float64  `json:"weekly_kills_a"`
				WeeklyKillsB                       float64  `json:"weekly_kills_b"`
				MonthlyKillsB                      float64  `json:"monthly_kills_b"`
				DashPower                          string   `json:"dash_power"`
				DashCooldown                       string   `json:"dash_cooldown"`
				Leggings                           string   `json:"leggings"`
				Boots                              string   `json:"boots"`
				AlternativeGunCooldownIndicator    bool     `json:"alternative_gun_cooldown_indicator"`
				CompassSelected                    bool     `json:"compass_selected"`
				EnableSound                        bool     `json:"enable_sound"`
				ShowDashCooldown                   bool     `json:"showDashCooldown"`
				HighestKillstreak                  float64  `json:"highest_killstreak"`
				DistanceTravelledTeams             float64  `json:"distance_travelled_teams"`
				KillsSinceUpdateFeb2017Teams       float64  `json:"kills_since_update_feb_2017_teams"`
				ShotsFiredTeams                    float64  `json:"shots_fired_teams"`
				HeadshotsTeams                     float64  `json:"headshots_teams"`
				ShowKillPrefix                     bool     `json:"showKillPrefix"`
				KillsDm                            float64  `json:"kills_dm"`
				KillsDmTeams                       float64  `json:"kills_dm_teams"`
				KillsTimeattack                    float64  `json:"kills_timeattack"`
				MessageOthersKillsDeaths           bool     `json:"messageOthers\u0027 Kills/deaths"`
				MessageYourKills                   bool     `json:"messageYour Kills"`
				MessageCoinMessages                bool     `json:"messageCoin Messages"`
				MessageMultiKills                  bool     `json:"messageMulti-kills"`
				MessageKillstreaks                 bool     `json:"messageKillstreaks"`
				MessagePowerupCollections          bool     `json:"messagePowerup Collections"`
				MessageYourDeaths                  bool     `json:"messageYour Deaths"`
				Headshots                          float64  `json:"headshots"`
				DistanceTravelled                  float64  `json:"distance_travelled"`
				ShotsFired                         float64  `json:"shots_fired"`
				KillsSinceUpdateFeb2017            float64  `json:"kills_since_update_feb_2017"`
				KillstreaksSoloTourney             float64  `json:"killstreaks_solo_tourney"`
				HeadshotsSoloTourney               float64  `json:"headshots_solo_tourney"`
				WinsSoloTourney                    float64  `json:"wins_solo_tourney"`
				KillsSoloTourney                   float64  `json:"kills_solo_tourney"`
				ShotsFiredSoloTourney              float64  `json:"shots_fired_solo_tourney"`
				KillsSinceUpdateFeb2017SoloTourney float64  `json:"kills_since_update_feb_2017_solo_tourney"`
				DistanceTravelledSoloTourney       float64  `json:"distance_travelled_solo_tourney"`
				DeathsSoloTourney                  float64  `json:"deaths_solo_tourney"`
				MessageCoin                        bool     `json:"messageCoin"`
			} `json:"Quake"`
			TNTGames struct {
				BloodwizardExplode            float64  `json:"bloodwizard_explode"`
				BloodwizardRegen              float64  `json:"bloodwizard_regen"`
				Coins                         float64  `json:"coins"`
				DeathsBowspleef               float64  `json:"deaths_bowspleef"`
				DeathsCapture                 float64  `json:"deaths_capture"`
				DoublejumpTntrun              float64  `json:"doublejump_tntrun"`
				FirewizardExplode             float64  `json:"firewizard_explode"`
				FirewizardRegen               float64  `json:"firewizard_regen"`
				KillsCapture                  float64  `json:"kills_capture"`
				KineticwizardExplode          float64  `json:"kineticwizard_explode"`
				KineticwizardRegen            float64  `json:"kineticwizard_regen"`
				Packages                      []string `json:"packages"`
				SpleefDoublejump              float64  `json:"spleef_doublejump"`
				SpleefRepulse                 float64  `json:"spleef_repulse"`
				SpleefTriple                  float64  `json:"spleef_triple"`
				TagsBowspleef                 float64  `json:"tags_bowspleef"`
				WinsBowspleef                 float64  `json:"wins_bowspleef"`
				WinsCapture                   float64  `json:"wins_capture"`
				WinsTntag                     float64  `json:"wins_tntag"`
				WinsTntrun                    float64  `json:"wins_tntrun"`
				WitherwizardExplode           float64  `json:"witherwizard_explode"`
				WitherwizardRegen             float64  `json:"witherwizard_regen"`
				SelectedHat                   string   `json:"selected_hat"`
				RecordPvprun                  float64  `json:"record_pvprun"`
				KillsPvprun                   float64  `json:"kills_pvprun"`
				RecordTntrun                  float64  `json:"record_tntrun"`
				TagSpeed                      float64  `json:"tag_speed"`
				CaptureClass                  string   `json:"capture_class"`
				AssistsCapture                float64  `json:"assists_capture"`
				WinsPvprun                    float64  `json:"wins_pvprun"`
				VotesTNTRunA                  float64  `json:"votes_TNT Run (A)"`
				ActiveParticleEffect          string   `json:"active_particle_effect"`
				IcewizardRegen                float64  `json:"icewizard_regen"`
				IcewizardExplode              float64  `json:"icewizard_explode"`
				ActiveDeathEffect             string   `json:"active_death_effect"`
				KillsTntag                    float64  `json:"kills_tntag"`
				NewWitherwizardRegen          float64  `json:"new_witherwizard_regen"`
				NewSpleefTripleshot           float64  `json:"new_spleef_tripleshot"`
				NewIcewizardExplode           float64  `json:"new_icewizard_explode"`
				NewSpleefRepulsor             float64  `json:"new_spleef_repulsor"`
				NewKineticwizardRegen         float64  `json:"new_kineticwizard_regen"`
				NewSpleefDoubleJumps          float64  `json:"new_spleef_double_jumps"`
				NewFirewizardExplode          float64  `json:"new_firewizard_explode"`
				NewTntagSpeedy                float64  `json:"new_tntag_speedy"`
				NewIcewizardRegen             float64  `json:"new_icewizard_regen"`
				NewWitherwizardExplode        float64  `json:"new_witherwizard_explode"`
				NewBloodwizardRegen           float64  `json:"new_bloodwizard_regen"`
				NewPvprunDoubleJumps          float64  `json:"new_pvprun_double_jumps"`
				NewBloodwizardExplode         float64  `json:"new_bloodwizard_explode"`
				NewTntrunDoubleJumps          float64  `json:"new_tntrun_double_jumps"`
				NewFirewizardRegen            float64  `json:"new_firewizard_regen"`
				NewKineticwizardExplode       float64  `json:"new_kineticwizard_explode"`
				NewToxicwizardRegen           float64  `json:"new_toxicwizard_regen"`
				NewToxicwizardExplode         float64  `json:"new_toxicwizard_explode"`
				NewTntrunSlownessPotions      float64  `json:"new_tntrun_slowness_potions"`
				NewTntrunSpeedPotions         float64  `json:"new_tntrun_speed_potions"`
				TagSpeeditup                  float64  `json:"tag_speeditup"`
				TagBlastprotection            float64  `json:"tag_blastprotection"`
				TagSlowitdown                 float64  `json:"tag_slowitdown"`
				NewPvprunFortitude            float64  `json:"new_pvprun_fortitude"`
				NewPvprunRegeneration         float64  `json:"new_pvprun_regeneration"`
				NewPvprunNotoriety            float64  `json:"new_pvprun_notoriety"`
				RunPotionsSplashedOnPlayers   float64  `json:"run_potions_splashed_on_players"`
				Winstreak                     float64  `json:"winstreak"`
				WizardsSelectedClass          string   `json:"wizards_selected_class"`
				NewWitherwizardAssists        float64  `json:"new_witherwizard_assists"`
				NewWitherwizardKills          float64  `json:"new_witherwizard_kills"`
				NewWitherwizardDeaths         float64  `json:"new_witherwizard_deaths"`
				DeathsPvprun                  float64  `json:"deaths_pvprun"`
				DeathsTntrun                  float64  `json:"deaths_tntrun"`
				NewKineticwizardKills         float64  `json:"new_kineticwizard_kills"`
				NewKineticwizardAssists       float64  `json:"new_kineticwizard_assists"`
				NewToxicwizardDeaths          float64  `json:"new_toxicwizard_deaths"`
				NewToxicwizardKills           float64  `json:"new_toxicwizard_kills"`
				NewKineticwizardDeaths        float64  `json:"new_kineticwizard_deaths"`
				NewToxicwizardAssists         float64  `json:"new_toxicwizard_assists"`
				NewActiveParticleEffect       string   `json:"new_active_particle_effect"`
				NewFirewizardAssists          float64  `json:"new_firewizard_assists"`
				NewFirewizardKills            float64  `json:"new_firewizard_kills"`
				NewFirewizardDeaths           float64  `json:"new_firewizard_deaths"`
				Wins                          float64  `json:"wins"`
				NewBloodwizardAssists         float64  `json:"new_bloodwizard_assists"`
				NewBloodwizardKills           float64  `json:"new_bloodwizard_kills"`
				NewBloodwizardDeaths          float64  `json:"new_bloodwizard_deaths"`
				NewToxicwizardRegenLegacy     float64  `json:"new_toxicwizard_regen_legacy"`
				NewIcewizardExplodeLegacy     float64  `json:"new_icewizard_explode_legacy"`
				NewFirewizardRegenLegacy      float64  `json:"new_firewizard_regen_legacy"`
				NewSpleefTripleshotLegacy     float64  `json:"new_spleef_tripleshot_legacy"`
				NewSpleefDoubleJumpsLegacy    float64  `json:"new_spleef_double_jumps_legacy"`
				NewKineticwizardExplodeLegacy float64  `json:"new_kineticwizard_explode_legacy"`
				NewIcewizardRegenLegacy       float64  `json:"new_icewizard_regen_legacy"`
				NewBloodwizardExplodeLegacy   float64  `json:"new_bloodwizard_explode_legacy"`
				NewWitherwizardRegenLegacy    float64  `json:"new_witherwizard_regen_legacy"`
				NewTntrunDoubleJumpsLegacy    float64  `json:"new_tntrun_double_jumps_legacy"`
				NewKineticwizardRegenLegacy   float64  `json:"new_kineticwizard_regen_legacy"`
				NewSpleefRepulsorLegacy       float64  `json:"new_spleef_repulsor_legacy"`
				NewBloodwizardRegenLegacy     float64  `json:"new_bloodwizard_regen_legacy"`
				NewWitherwizardExplodeLegacy  float64  `json:"new_witherwizard_explode_legacy"`
				NewToxicwizardExplodeLegacy   float64  `json:"new_toxicwizard_explode_legacy"`
				NewFirewizardExplodeLegacy    float64  `json:"new_firewizard_explode_legacy"`
				NewPvprunDoubleJumpsLegacy    float64  `json:"new_pvprun_double_jumps_legacy"`
				PointsCapture                 float64  `json:"points_capture"`
				NewKineticwizardDamageTaken   float64  `json:"new_kineticwizard_damage_taken"`
				NewFirewizardDamageTaken      float64  `json:"new_firewizard_damage_taken"`
				NewKineticwizardHealing       float64  `json:"new_kineticwizard_healing"`
				KineticHealingCapture         float64  `json:"kinetic_healing_capture"`
				NewFirewizardHealing          float64  `json:"new_firewizard_healing"`
				AirTimeCapture                float64  `json:"air_time_capture"`
				Privategames                  struct {
					LowGravity       bool   `json:"low_gravity"`
					Speed            string `json:"speed"`
					TntRunSnowballs  bool   `json:"tnt_run_snowballs"`
					TntTagDeathmatch bool   `json:"tnt_tag_deathmatch"`
					PvpRunSwordType  string `json:"pvp_run_sword_type"`
					PvpRunArmorType  string `json:"pvp_run_armor_type"`
					MaxedPerks       bool   `json:"maxed_perks"`
					TntTagNoPowerups bool   `json:"tnt_tag_no_powerups"`
					PvpRunKnockback  bool   `json:"pvp_run_knockback"`
				} `json:"privategames"`
				Flags struct {
					EnableExplosiveDash              bool `json:"enable_explosive_dash"`
					GiveDjFeather                    bool `json:"give_dj_feather"`
					ShowTipHolograms                 bool `json:"show_tip_holograms"`
					ShowTntrunActionbarInfo          bool `json:"show_tntrun_actionbar_info"`
					ShowTnttagActionbarInfo          bool `json:"show_tnttag_actionbar_info"`
					ShowWizardsActionbarInfo         bool `json:"show_wizards_actionbar_info"`
					ShowWizardsCooldownNotifications bool `json:"show_wizards_cooldown_notifications"`
				} `json:"flags"`
				LastTourneyAd         float64 `json:"lastTourneyAd"`
				NewStormwizardExplode float64 `json:"new_stormwizard_explode"`
				NewStormwizardRegen   float64 `json:"new_stormwizard_regen"`
			} `json:"TNTGames"`
			Uhc struct {
				Coins                   float64  `json:"coins"`
				Deaths                  float64  `json:"deaths"`
				EquippedKit             string   `json:"equippedKit"`
				HeadsEaten              float64  `json:"heads_eaten"`
				Kills                   float64  `json:"kills"`
				Packages                []string `json:"packages"`
				PerkAlchemyLineA        float64  `json:"perk_alchemy_line_a"`
				PerkAlchemyLineB        float64  `json:"perk_alchemy_line_b"`
				PerkArmorsmithLineA     float64  `json:"perk_armorsmith_line_a"`
				PerkArmorsmithLineB     float64  `json:"perk_armorsmith_line_b"`
				PerkArmorsmithLineC     float64  `json:"perk_armorsmith_line_c"`
				PerkEnchantingLineA     float64  `json:"perk_enchanting_line_a"`
				PerkEnchantingLineB     float64  `json:"perk_enchanting_line_b"`
				PerkEnchantingLineC     float64  `json:"perk_enchanting_line_c"`
				PerkEngineeringLineA    float64  `json:"perk_engineering_line_a"`
				PerkEngineeringLineB    float64  `json:"perk_engineering_line_b"`
				PerkEngineeringLineC    float64  `json:"perk_engineering_line_c"`
				PerkWeaponsmithLineA    float64  `json:"perk_weaponsmith_line_a"`
				PerkWeaponsmithLineB    float64  `json:"perk_weaponsmith_line_b"`
				PerkWeaponsmithLineC    float64  `json:"perk_weaponsmith_line_c"`
				Score                   float64  `json:"score"`
				Wins                    float64  `json:"wins"`
				PerkBloodcraftLineA     float64  `json:"perk_bloodcraft_line_a"`
				PerkBloodcraftLineB     float64  `json:"perk_bloodcraft_line_b"`
				PerkHunterLineA         float64  `json:"perk_hunter_line_a"`
				PerkHunterLineB         float64  `json:"perk_hunter_line_b"`
				PerkHunterLineC         float64  `json:"perk_hunter_line_c"`
				KitARCHERYTOOLS         float64  `json:"kit_ARCHERY_TOOLS"`
				PerkCookingLineA        float64  `json:"perk_cooking_line_a"`
				PerkCookingLineC        float64  `json:"perk_cooking_line_c"`
				PerkCookingLineB        float64  `json:"perk_cooking_line_b"`
				MonthlyKillsA           float64  `json:"monthly_kills_a"`
				MonthlyWinsA            float64  `json:"monthly_wins_a"`
				PerkCookingPrestige     float64  `json:"perk_cooking_prestige"`
				PerkBloodcraftLineC     float64  `json:"perk_bloodcraft_line_c"`
				PerkBloodcraftPrestige  float64  `json:"perk_bloodcraft_prestige"`
				DeathsSolo              float64  `json:"deaths_solo"`
				Cache3                  bool     `json:"cache3"`
				MonthlyKillsB           float64  `json:"monthly_kills__b"`
				ClearupAchievement      bool     `json:"clearup_achievement"`
				KillsSolo               float64  `json:"kills_solo"`
				HeadsEatenSolo          float64  `json:"heads_eaten_solo"`
				PerkSurvivalismLineA    float64  `json:"perk_survivalism_line_a"`
				PerkSurvivalismLineC    float64  `json:"perk_survivalism_line_c"`
				PerkSurvivalismLineB    float64  `json:"perk_survivalism_line_b"`
				KillsSolo2              float64  `json:"kills_solo2"`
				SavedStats              bool     `json:"saved_stats"`
				Kills2                  float64  `json:"kills2"`
				Wins2                   float64  `json:"wins2"`
				WinsSolo                float64  `json:"wins_solo"`
				PerkWeaponsmithPrestige float64  `json:"perk_weaponsmith_prestige"`
			} `json:"UHC"`
			VampireZ struct {
				BabyHater            float64  `json:"baby_hater"`
				Blood                bool     `json:"blood"`
				BloodBooster         float64  `json:"blood_booster"`
				Coins                float64  `json:"coins"`
				DrainPunch           float64  `json:"drain_punch"`
				ExpertSwag           float64  `json:"expert_swag"`
				ExplosiveKiller      float64  `json:"explosive_killer"`
				FinalBreath          float64  `json:"final_breath"`
				Fireproofing         float64  `json:"fireproofing"`
				GoldBooster          float64  `json:"gold_booster"`
				GoldStarter          float64  `json:"gold_starter"`
				HumanDeaths          float64  `json:"human_deaths"`
				HumanKills           float64  `json:"human_kills"`
				HumanWins            float64  `json:"human_wins"`
				LootDrops            float64  `json:"loot_drops"`
				MostVampireKills     float64  `json:"most_vampire_kills"`
				Packages             []string `json:"packages"`
				VampireDeaths        float64  `json:"vampire_deaths"`
				VampireDoubler       float64  `json:"vampire_doubler"`
				VampireKills         float64  `json:"vampire_kills"`
				VampiricScream       float64  `json:"vampiric_scream"`
				WaveBooster          float64  `json:"wave_booster"`
				ZombieDoubler        float64  `json:"zombie_doubler"`
				ZombieKills          float64  `json:"zombie_kills"`
				Hellborn             float64  `json:"hellborn"`
				Renfield             float64  `json:"renfield"`
				BloodDrinker         float64  `json:"blood_drinker"`
				TerrorLevel          float64  `json:"terror_level"`
				KillBooster          float64  `json:"kill_booster"`
				Transfusion          float64  `json:"transfusion"`
				BasicSwag            float64  `json:"basic_swag"`
				AdvancedSwag         float64  `json:"advanced_swag"`
				Constitution         float64  `json:"constitution"`
				Foresight            float64  `json:"foresight"`
				VotesDarkValley      float64  `json:"votes_Dark Valley"`
				GoldBought           float64  `json:"gold_bought"`
				MonthlyHumanWinsB    float64  `json:"monthly_human_wins_b"`
				WeeklyHumanWinsA     float64  `json:"weekly_human_wins_a"`
				MonthlyVampireWinsB  float64  `json:"monthly_vampire_wins_b"`
				WeeklyVampireWinsA   float64  `json:"weekly_vampire_wins_a"`
				WeeklyHumanWinsB     float64  `json:"weekly_human_wins_b"`
				MonthlyHumanWinsA    float64  `json:"monthly_human_wins_a"`
				WeeklyVampireWinsB   float64  `json:"weekly_vampire_wins_b"`
				MonthlyVampireWinsA  float64  `json:"monthly_vampire_wins_a"`
				VampiricMinion       float64  `json:"vampiric_minion"`
				FrankensteinsMonster float64  `json:"frankensteins_monster"`
				UpdatedStats         bool     `json:"updated_stats"`
				VampireWins          float64  `json:"vampire_wins"`
				MostVampireKillsNew  float64  `json:"most_vampire_kills_new"`
			} `json:"VampireZ"`
			Walls struct {
				Berserk       float64  `json:"berserk"`
				Blood         bool     `json:"blood"`
				BossDigger    float64  `json:"boss_digger"`
				BossSkills    float64  `json:"boss_skills"`
				Coins         float64  `json:"coins"`
				Deaths        float64  `json:"deaths"`
				Fortune       float64  `json:"fortune"`
				InsaneFarmer  float64  `json:"insane_farmer"`
				Kills         float64  `json:"kills"`
				Losses        float64  `json:"losses"`
				Packages      []string `json:"packages"`
				Scotsman      float64  `json:"scotsman"`
				Swift         float64  `json:"swift"`
				Wins          float64  `json:"wins"`
				Chef          float64  `json:"chef"`
				Hunter        float64  `json:"hunter"`
				Sage          float64  `json:"sage"`
				SmartBoy      float64  `json:"smart_boy"`
				Bacon         float64  `json:"bacon"`
				Guitarist     float64  `json:"guitarist"`
				SnackLover    float64  `json:"snack_lover"`
				Blacksmith    float64  `json:"blacksmith"`
				FinalForm     float64  `json:"final_form"`
				SkybaseKing   float64  `json:"skybase_king"`
				Einstein      float64  `json:"einstein"`
				Canadian      float64  `json:"canadian"`
				Tenacity      float64  `json:"tenacity"`
				Vampirism     float64  `json:"vampirism"`
				ExpertMiner   float64  `json:"expert_miner"`
				Adrenaline    float64  `json:"adrenaline"`
				Lazyman       float64  `json:"lazyman"`
				Step          float64  `json:"step"`
				Bomberman     float64  `json:"bomberman"`
				Ready         float64  `json:"ready"`
				Artisan       float64  `json:"artisan"`
				Fisherman     float64  `json:"fisherman"`
				GoldRush      float64  `json:"gold_rush"`
				LeatherWorker float64  `json:"leather_worker"`
				SoupDrinker   float64  `json:"soup_drinker"`
				BossGuardian  float64  `json:"boss_guardian"`
				Opportunity   float64  `json:"opportunity"`
				Haste         float64  `json:"haste"`
				MonthlyKillsB float64  `json:"monthly_kills_b"`
				WeeklyKillsA  float64  `json:"weekly_kills_a"`
				Inventory     struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num10 string `json:"10"`
					Num11 string `json:"11"`
					Num12 string `json:"12"`
					Num13 string `json:"13"`
					Num14 string `json:"14"`
					Num15 string `json:"15"`
					Num16 string `json:"16"`
					Num18 string `json:"18"`
					Num19 string `json:"19"`
					Num20 string `json:"20"`
					Num21 string `json:"21"`
					Num22 string `json:"22"`
					Num23 string `json:"23"`
				} `json:"Inventory"`
				BurnBabyBurn     float64 `json:"burn_baby_burn"`
				Vitality         float64 `json:"vitality"`
				WeeklyWinsA      float64 `json:"weekly_wins_a"`
				MonthlyWinsB     float64 `json:"monthly_wins_b"`
				CreeperEgg       float64 `json:"creeper_egg"`
				Chainkiller      float64 `json:"chainkiller"`
				MasterTroll      float64 `json:"master_troll"`
				Fireproof        float64 `json:"fireproof"`
				InventoryLayout2 struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num10 string `json:"10"`
					Num11 string `json:"11"`
					Num12 string `json:"12"`
					Num13 string `json:"13"`
					Num14 string `json:"14"`
					Num15 string `json:"15"`
					Num16 string `json:"16"`
					Num17 string `json:"17"`
					Num18 string `json:"18"`
					Num19 string `json:"19"`
					Num20 string `json:"20"`
					Num21 string `json:"21"`
					Num22 string `json:"22"`
					Num23 string `json:"23"`
					Num24 string `json:"24"`
					Num25 string `json:"25"`
					Num26 string `json:"26"`
					Num27 string `json:"27"`
				} `json:"InventoryLayout2"`
				BlacksmithStarter float64 `json:"blacksmith_starter"`
			} `json:"Walls"`
			Walls3 struct {
				Assists                                  float64  `json:"assists"`
				AssistsSkeleton                          float64  `json:"assists_Skeleton"`
				BlazeA                                   float64  `json:"blaze_a"`
				BlazeB                                   float64  `json:"blaze_b"`
				BlazeC                                   float64  `json:"blaze_c"`
				BlazeD                                   float64  `json:"blaze_d"`
				BlazeG                                   float64  `json:"blaze_g"`
				Blood                                    bool     `json:"blood"`
				ChosenClass                              string   `json:"chosen_class"`
				Coins                                    float64  `json:"coins"`
				CreeperA                                 float64  `json:"creeper_a"`
				CreeperB                                 float64  `json:"creeper_b"`
				CreeperC                                 float64  `json:"creeper_c"`
				CreeperD                                 float64  `json:"creeper_d"`
				CreeperG                                 float64  `json:"creeper_g"`
				Deaths                                   float64  `json:"deaths"`
				DeathsBlaze                              float64  `json:"deaths_Blaze"`
				DeathsCreeper                            float64  `json:"deaths_Creeper"`
				DeathsEnderman                           float64  `json:"deaths_Enderman"`
				DeathsGolem                              float64  `json:"deaths_Golem"`
				DeathsHerobrine                          float64  `json:"deaths_Herobrine"`
				DeathsPigman                             float64  `json:"deaths_Pigman"`
				DeathsShaman                             float64  `json:"deaths_Shaman"`
				DeathsSkeleton                           float64  `json:"deaths_Skeleton"`
				DeathsSpider                             float64  `json:"deaths_Spider"`
				DeathsZombie                             float64  `json:"deaths_Zombie"`
				EndermanA                                float64  `json:"enderman_a"`
				EndermanB                                float64  `json:"enderman_b"`
				EndermanC                                float64  `json:"enderman_c"`
				EndermanD                                float64  `json:"enderman_d"`
				FinalKillsBlaze                          float64  `json:"finalKills_Blaze"`
				FinalKillsCreeper                        float64  `json:"finalKills_Creeper"`
				FinalKillsEnderman                       float64  `json:"finalKills_Enderman"`
				FinalKillsGolem                          float64  `json:"finalKills_Golem"`
				FinalKillsHerobrine                      float64  `json:"finalKills_Herobrine"`
				FinalKillsPigman                         float64  `json:"finalKills_Pigman"`
				FinalKillsShaman                         float64  `json:"finalKills_Shaman"`
				FinalKillsSkeleton                       float64  `json:"finalKills_Skeleton"`
				FinalKillsSpider                         float64  `json:"finalKills_Spider"`
				FinalKillsZombie                         float64  `json:"finalKills_Zombie"`
				GolemD                                   float64  `json:"golem_d"`
				HerobrineA                               float64  `json:"herobrine_a"`
				HerobrineB                               float64  `json:"herobrine_b"`
				HerobrineC                               float64  `json:"herobrine_c"`
				HerobrineD                               float64  `json:"herobrine_d"`
				HerobrineG                               float64  `json:"herobrine_g"`
				Kills                                    float64  `json:"kills"`
				KillsBlaze                               float64  `json:"kills_Blaze"`
				KillsCreeper                             float64  `json:"kills_Creeper"`
				KillsEnderman                            float64  `json:"kills_Enderman"`
				KillsGolem                               float64  `json:"kills_Golem"`
				KillsHerobrine                           float64  `json:"kills_Herobrine"`
				KillsPigman                              float64  `json:"kills_Pigman"`
				KillsShaman                              float64  `json:"kills_Shaman"`
				KillsSkeleton                            float64  `json:"kills_Skeleton"`
				KillsSpider                              float64  `json:"kills_Spider"`
				KillsZombie                              float64  `json:"kills_Zombie"`
				Losses                                   float64  `json:"losses"`
				LossesBlaze                              float64  `json:"losses_Blaze"`
				LossesCreeper                            float64  `json:"losses_Creeper"`
				LossesEnderman                           float64  `json:"losses_Enderman"`
				LossesGolem                              float64  `json:"losses_Golem"`
				LossesHerobrine                          float64  `json:"losses_Herobrine"`
				LossesPigman                             float64  `json:"losses_Pigman"`
				LossesShaman                             float64  `json:"losses_Shaman"`
				LossesSkeleton                           float64  `json:"losses_Skeleton"`
				LossesSpider                             float64  `json:"losses_Spider"`
				LossesZombie                             float64  `json:"losses_Zombie"`
				MonthlyFinalKillsHerobrineA              float64  `json:"monthly_finalKills_Herobrine_a"`
				MonthlyFinalKillsPigmanA                 float64  `json:"monthly_finalKills_Pigman_a"`
				MonthlyFinalKillsPigmanB                 float64  `json:"monthly_finalKills_Pigman_b"`
				MonthlyFinalKillsSkeletonA               float64  `json:"monthly_finalKills_Skeleton_a"`
				MonthlyFinalKillsSkeletonB               float64  `json:"monthly_finalKills_Skeleton_b"`
				MonthlyFinalKillsSpiderA                 float64  `json:"monthly_finalKills_Spider_a"`
				MonthlyFinalKillsSpiderB                 float64  `json:"monthly_finalKills_Spider_b"`
				MonthlyFinalKillsZombieA                 float64  `json:"monthly_finalKills_Zombie_a"`
				MutationsVisibility                      bool     `json:"mutations_visibility"`
				Packages                                 []string `json:"packages"`
				PigmanA                                  float64  `json:"pigman_a"`
				PigmanB                                  float64  `json:"pigman_b"`
				PigmanC                                  float64  `json:"pigman_c"`
				PigmanD                                  float64  `json:"pigman_d"`
				ShamanA                                  float64  `json:"shaman_a"`
				ShamanB                                  float64  `json:"shaman_b"`
				ShamanC                                  float64  `json:"shaman_c"`
				ShamanD                                  float64  `json:"shaman_d"`
				ShamanG                                  float64  `json:"shaman_g"`
				SkeletonA                                float64  `json:"skeleton_a"`
				SkeletonB                                float64  `json:"skeleton_b"`
				SkeletonC                                float64  `json:"skeleton_c"`
				SkeletonD                                float64  `json:"skeleton_d"`
				SkeletonG                                float64  `json:"skeleton_g"`
				SpiderA                                  float64  `json:"spider_a"`
				SpiderB                                  float64  `json:"spider_b"`
				SpiderD                                  float64  `json:"spider_d"`
				WarCry                                   string   `json:"war_cry"`
				WeeklyDeaths                             float64  `json:"weeklyDeaths"`
				WeeklyDeathsBlaze                        float64  `json:"weeklyDeaths_Blaze"`
				WeeklyDeathsGolem                        float64  `json:"weeklyDeaths_Golem"`
				WeeklyDeathsHerobrine                    float64  `json:"weeklyDeaths_Herobrine"`
				WeeklyDeathsPigman                       float64  `json:"weeklyDeaths_Pigman"`
				WeeklyDeathsSkeleton                     float64  `json:"weeklyDeaths_Skeleton"`
				WeeklyDeathsSpider                       float64  `json:"weeklyDeaths_Spider"`
				WeeklyDeathsZombie                       float64  `json:"weeklyDeaths_Zombie"`
				WeeklyFinalKillsBlaze                    float64  `json:"weeklyFinalKills_Blaze"`
				WeeklyFinalKillsGolem                    float64  `json:"weeklyFinalKills_Golem"`
				WeeklyFinalKillsPigman                   float64  `json:"weeklyFinalKills_Pigman"`
				WeeklyKills                              float64  `json:"weeklyKills"`
				WeeklyKillsBlaze                         float64  `json:"weeklyKills_Blaze"`
				WeeklyKillsGolem                         float64  `json:"weeklyKills_Golem"`
				WeeklyKillsHerobrine                     float64  `json:"weeklyKills_Herobrine"`
				WeeklyKillsPigman                        float64  `json:"weeklyKills_Pigman"`
				WeeklyKillsSkeleton                      float64  `json:"weeklyKills_Skeleton"`
				WeeklyKillsSpider                        float64  `json:"weeklyKills_Spider"`
				WeeklyKillsZombie                        float64  `json:"weeklyKills_Zombie"`
				WeeklyLosses                             float64  `json:"weeklyLosses"`
				WeeklyLossesBlaze                        float64  `json:"weeklyLosses_Blaze"`
				WeeklyLossesGolem                        float64  `json:"weeklyLosses_Golem"`
				WeeklyLossesHerobrine                    float64  `json:"weeklyLosses_Herobrine"`
				WeeklyLossesPigman                       float64  `json:"weeklyLosses_Pigman"`
				WeeklyLossesSkeleton                     float64  `json:"weeklyLosses_Skeleton"`
				WeeklyLossesSpider                       float64  `json:"weeklyLosses_Spider"`
				WeeklyLossesZombie                       float64  `json:"weeklyLosses_Zombie"`
				WeeklyWins                               float64  `json:"weeklyWins"`
				WeeklyWinsBlaze                          float64  `json:"weeklyWins_Blaze"`
				WeeklyWinsGolem                          float64  `json:"weeklyWins_Golem"`
				WeeklyWinsPigman                         float64  `json:"weeklyWins_Pigman"`
				WeeklyWinsSkeleton                       float64  `json:"weeklyWins_Skeleton"`
				WeeklyWinsSpider                         float64  `json:"weeklyWins_Spider"`
				WeeklyWinsZombie                         float64  `json:"weeklyWins_Zombie"`
				WeeklyFinalKillsHerobrineB               float64  `json:"weekly_finalKills_Herobrine_b"`
				WeeklyFinalKillsPigmanA                  float64  `json:"weekly_finalKills_Pigman_a"`
				WeeklyFinalKillsPigmanB                  float64  `json:"weekly_finalKills_Pigman_b"`
				WeeklyFinalKillsSkeletonA                float64  `json:"weekly_finalKills_Skeleton_a"`
				WeeklyFinalKillsSkeletonB                float64  `json:"weekly_finalKills_Skeleton_b"`
				WeeklyFinalKillsSpiderA                  float64  `json:"weekly_finalKills_Spider_a"`
				WeeklyFinalKillsSpiderB                  float64  `json:"weekly_finalKills_Spider_b"`
				WeeklyFinalKillsZombieB                  float64  `json:"weekly_finalKills_Zombie_b"`
				Wins                                     float64  `json:"wins"`
				WinsBlaze                                float64  `json:"wins_Blaze"`
				WinsCreeper                              float64  `json:"wins_Creeper"`
				WinsEnderman                             float64  `json:"wins_Enderman"`
				WinsGolem                                float64  `json:"wins_Golem"`
				WinsHerobrine                            float64  `json:"wins_Herobrine"`
				WinsPigman                               float64  `json:"wins_Pigman"`
				WinsShaman                               float64  `json:"wins_Shaman"`
				WinsSkeleton                             float64  `json:"wins_Skeleton"`
				WinsSpider                               float64  `json:"wins_Spider"`
				WinsZombie                               float64  `json:"wins_Zombie"`
				ZombieA                                  float64  `json:"zombie_a"`
				ZombieB                                  float64  `json:"zombie_b"`
				ZombieC                                  float64  `json:"zombie_c"`
				ZombieD                                  float64  `json:"zombie_d"`
				ZombieG                                  float64  `json:"zombie_g"`
				LossesArcanist                           float64  `json:"losses_Arcanist"`
				WeeklyLossesArcanist                     float64  `json:"weeklyLosses_Arcanist"`
				MonthlyFinalKillsGolemA                  float64  `json:"monthly_finalKills_Golem_a"`
				WeeklyFinalKillsGolemA                   float64  `json:"weekly_finalKills_Golem_a"`
				AssistsGolem                             float64  `json:"assists_Golem"`
				FinalAssistsGolem                        float64  `json:"finalAssists_Golem"`
				AssistsHerobrine                         float64  `json:"assists_Herobrine"`
				WeeklyWinsHerobrine                      float64  `json:"weeklyWins_Herobrine"`
				FinalAssistsHerobrine                    float64  `json:"finalAssists_Herobrine"`
				FinalAssistsSpider                       float64  `json:"finalAssists_Spider"`
				PlaysStandard                            float64  `json:"plays_standard"`
				MonthlyFinalKillsHerobrineB              float64  `json:"monthly_finalKills_Herobrine_b"`
				WeeklyFinalKillsHerobrineA               float64  `json:"weekly_finalKills_Herobrine_a"`
				MonthlyFinalKillsGolemB                  float64  `json:"monthly_finalKills_Golem_b"`
				FinalAssistsSkeleton                     float64  `json:"finalAssists_Skeleton"`
				PlaysPractice                            float64  `json:"plays_practice"`
				WeeklyWinsPracticeSkeleton               float64  `json:"weeklyWins_practice_Skeleton"`
				WinsPractice                             float64  `json:"wins_practice"`
				WeeklyWinsPractice                       float64  `json:"weeklyWins_practice"`
				WinsPracticeSkeleton                     float64  `json:"wins_practice_Skeleton"`
				FinalKillsPirate                         float64  `json:"finalKills_Pirate"`
				FinalAssistsPirate                       float64  `json:"finalAssists_Pirate"`
				WeeklyLossesPirate                       float64  `json:"weeklyLosses_Pirate"`
				AssistsPirate                            float64  `json:"assists_Pirate"`
				DeathsPirate                             float64  `json:"deaths_Pirate"`
				WeeklyKillsPirate                        float64  `json:"weeklyKills_Pirate"`
				WeeklyFinalKillsPirateA                  float64  `json:"weekly_finalKills_Pirate_a"`
				LossesPirate                             float64  `json:"losses_Pirate"`
				WeeklyDeathsPirate                       float64  `json:"weeklyDeaths_Pirate"`
				KillsPirate                              float64  `json:"kills_Pirate"`
				MonthlyFinalKillsPirateB                 float64  `json:"monthly_finalKills_Pirate_b"`
				FinalKillsHunter                         float64  `json:"finalKills_Hunter"`
				WinsHunter                               float64  `json:"wins_Hunter"`
				WeeklyWinsHunter                         float64  `json:"weeklyWins_Hunter"`
				WeeklyKillsHunter                        float64  `json:"weeklyKills_Hunter"`
				AssistsHunter                            float64  `json:"assists_Hunter"`
				KillsHunter                              float64  `json:"kills_Hunter"`
				FinalAssistsHunter                       float64  `json:"finalAssists_Hunter"`
				WeeklyFinalKillsHunterA                  float64  `json:"weekly_finalKills_Hunter_a"`
				MonthlyFinalKillsHunterB                 float64  `json:"monthly_finalKills_Hunter_b"`
				LossesHunter                             float64  `json:"losses_Hunter"`
				WeeklyLossesHunter                       float64  `json:"weeklyLosses_Hunter"`
				WeeklyLossesFaceOffHunter                float64  `json:"weeklyLosses_face_off_Hunter"`
				PlaysFaceOff                             float64  `json:"plays_face_off"`
				LossesFaceOff                            float64  `json:"losses_face_off"`
				LossesFaceOffHunter                      float64  `json:"losses_face_off_Hunter"`
				WeeklyLossesFaceOff                      float64  `json:"weeklyLosses_face_off"`
				WinsFaceOffHunter                        float64  `json:"wins_face_off_Hunter"`
				WinsFaceOff                              float64  `json:"wins_face_off"`
				WeeklyWinsFaceOffHunter                  float64  `json:"weeklyWins_face_off_Hunter"`
				WeeklyDeathsHunter                       float64  `json:"weeklyDeaths_Hunter"`
				WeeklyWinsFaceOff                        float64  `json:"weeklyWins_face_off"`
				DeathsHunter                             float64  `json:"deaths_Hunter"`
				WeeklyFinalKillsHunterB                  float64  `json:"weekly_finalKills_Hunter_b"`
				MonthlyFinalKillsHunterA                 float64  `json:"monthly_finalKills_Hunter_a"`
				WeeklyLossesFaceOffBlaze                 float64  `json:"weeklyLosses_face_off_Blaze"`
				LossesFaceOffBlaze                       float64  `json:"losses_face_off_Blaze"`
				AssistsSpider                            float64  `json:"assists_Spider"`
				WitherCoins                              float64  `json:"witherCoins"`
				SmileyKills                              string   `json:"smiley_kills"`
				MonthlyFinalKillsB                       float64  `json:"monthly_finalKills_b"`
				WeeklyWinsFaceOffHerobrine               float64  `json:"weeklyWins_face_off_Herobrine"`
				WinsFaceOffHerobrine                     float64  `json:"wins_face_off_Herobrine"`
				WeeklyFinalKillsA                        float64  `json:"weekly_finalKills_a"`
				MonthlyFinalKillsA                       float64  `json:"monthly_finalKills_a"`
				WeeklyFinalKillsB                        float64  `json:"weekly_finalKills_b"`
				WeeklyWinsArcanist                       float64  `json:"weeklyWins_Arcanist"`
				FinalKillsArcanist                       float64  `json:"finalKills_Arcanist"`
				WinsFaceOffArcanist                      float64  `json:"wins_face_off_Arcanist"`
				WeeklyFinalKillsArcanistA                float64  `json:"weekly_finalKills_Arcanist_a"`
				WeeklyDeathsArcanist                     float64  `json:"weeklyDeaths_Arcanist"`
				WeeklyKillsArcanist                      float64  `json:"weeklyKills_Arcanist"`
				MonthlyFinalKillsArcanistB               float64  `json:"monthly_finalKills_Arcanist_b"`
				KillsArcanist                            float64  `json:"kills_Arcanist"`
				DeathsArcanist                           float64  `json:"deaths_Arcanist"`
				WinsArcanist                             float64  `json:"wins_Arcanist"`
				FinalAssistsArcanist                     float64  `json:"finalAssists_Arcanist"`
				WeeklyWinsFaceOffArcanist                float64  `json:"weeklyWins_face_off_Arcanist"`
				AssistsArcanist                          float64  `json:"assists_Arcanist"`
				WeeklyFinalKillsBlazeB                   float64  `json:"weekly_finalKills_Blaze_b"`
				MonthlyFinalKillsBlazeB                  float64  `json:"monthly_finalKills_Blaze_b"`
				FinalAssistsBlaze                        float64  `json:"finalAssists_Blaze"`
				AssistsBlaze                             float64  `json:"assists_Blaze"`
				WeeklyWinsFaceOffBlaze                   float64  `json:"weeklyWins_face_off_Blaze"`
				WinsFaceOffBlaze                         float64  `json:"wins_face_off_Blaze"`
				SpiderEnderchestLevel                    float64  `json:"spider_enderchest_level"`
				HunterFinalKills                         float64  `json:"hunter_final_kills"`
				HunterTotalFinalKills                    float64  `json:"hunter_total_final_kills"`
				ZombieKills                              float64  `json:"zombie_kills"`
				PigmanTotalFinalKills                    float64  `json:"pigman_total_final_kills"`
				ZombieTotalFinalKillsStandard            float64  `json:"zombie_total_final_kills_standard"`
				EndermanFinalKillsStandard               float64  `json:"enderman_final_kills_standard"`
				ArcanistTotalFinalKills                  float64  `json:"arcanist_total_final_kills"`
				SkeletonFinalKills                       float64  `json:"skeleton_final_kills"`
				HerobrineEnderchestLevel                 float64  `json:"herobrine_enderchest_level"`
				BlazeFinalKillsStandard                  float64  `json:"blaze_final_kills_standard"`
				BlazeTotalFinalKillsStandard             float64  `json:"blaze_total_final_kills_standard"`
				HunterTotalFinalKillsStandard            float64  `json:"hunter_total_final_kills_standard"`
				EndermanTotalFinalKillsStandard          float64  `json:"enderman_total_final_kills_standard"`
				PirateTotalFinalKills                    float64  `json:"pirate_total_final_kills"`
				SkeletonFinalKillsStandard               float64  `json:"skeleton_final_kills_standard"`
				CreeperTotalFinalKills                   float64  `json:"creeper_total_final_kills"`
				ZombieWins                               float64  `json:"zombie_wins"`
				ArcanistWins                             float64  `json:"arcanist_wins"`
				CreeperFinalKillsStandard                float64  `json:"creeper_final_kills_standard"`
				SkeletonTotalFinalKillsStandard          float64  `json:"skeleton_total_final_kills_standard"`
				CreeperLosses                            float64  `json:"creeper_losses"`
				HunterDeaths                             float64  `json:"hunter_deaths"`
				ArcanistKills                            float64  `json:"arcanist_kills"`
				HerobrineTotalFinalKillsStandard         float64  `json:"herobrine_total_final_kills_standard"`
				BlazeWins                                float64  `json:"blaze_wins"`
				CreeperFinalKills                        float64  `json:"creeper_final_kills"`
				ArcanistDeaths                           float64  `json:"arcanist_deaths"`
				ArcanistLosses                           float64  `json:"arcanist_losses"`
				GolemTotalFinalKillsStandard             float64  `json:"golem_total_final_kills_standard"`
				HerobrineKills                           float64  `json:"herobrine_kills"`
				ZombieDeaths                             float64  `json:"zombie_deaths"`
				PirateTotalFinalKillsStandard            float64  `json:"pirate_total_final_kills_standard"`
				ShamanWins                               float64  `json:"shaman_wins"`
				SpiderTotalFinalKills                    float64  `json:"spider_total_final_kills"`
				ArcanistFinalKills                       float64  `json:"arcanist_final_kills"`
				ShamanTotalFinalKillsStandard            float64  `json:"shaman_total_final_kills_standard"`
				HunterWins                               float64  `json:"hunter_wins"`
				ShamanFinalKills                         float64  `json:"shaman_final_kills"`
				BlazeFinalKills                          float64  `json:"blaze_final_kills"`
				GolemKills                               float64  `json:"golem_kills"`
				GolemWins                                float64  `json:"golem_wins"`
				CreeperKills                             float64  `json:"creeper_kills"`
				EndermanWins                             float64  `json:"enderman_wins"`
				GolemDeaths                              float64  `json:"golem_deaths"`
				ZombieTotalFinalKills                    float64  `json:"zombie_total_final_kills"`
				ShamanKills                              float64  `json:"shaman_kills"`
				GolemTotalFinalKills                     float64  `json:"golem_total_final_kills"`
				ArcanistTotalFinalKillsStandard          float64  `json:"arcanist_total_final_kills_standard"`
				PirateDeaths                             float64  `json:"pirate_deaths"`
				PigmanTotalFinalKillsStandard            float64  `json:"pigman_total_final_kills_standard"`
				CreeperWins                              float64  `json:"creeper_wins"`
				SpiderFinalKillsStandard                 float64  `json:"spider_final_kills_standard"`
				BlazeKills                               float64  `json:"blaze_kills"`
				HerobrineTotalFinalKills                 float64  `json:"herobrine_total_final_kills"`
				ArcanistFinalKillsStandard               float64  `json:"arcanist_final_kills_standard"`
				HerobrineFinalKillsStandard              float64  `json:"herobrine_final_kills_standard"`
				PigmanWins                               float64  `json:"pigman_wins"`
				HunterKills                              float64  `json:"hunter_kills"`
				EndermanKills                            float64  `json:"enderman_kills"`
				PigmanDeaths                             float64  `json:"pigman_deaths"`
				PigmanLosses                             float64  `json:"pigman_losses"`
				ZombieFinalKills                         float64  `json:"zombie_final_kills"`
				PigmanFinalKills                         float64  `json:"pigman_final_kills"`
				SpiderKills                              float64  `json:"spider_kills"`
				ShamanLosses                             float64  `json:"shaman_losses"`
				EndermanFinalKills                       float64  `json:"enderman_final_kills"`
				PirateLosses                             float64  `json:"pirate_losses"`
				SpiderWins                               float64  `json:"spider_wins"`
				PirateKills                              float64  `json:"pirate_kills"`
				SkeletonWins                             float64  `json:"skeleton_wins"`
				PirateFinalKills                         float64  `json:"pirate_final_kills"`
				SpiderLosses                             float64  `json:"spider_losses"`
				CreeperTotalFinalKillsStandard           float64  `json:"creeper_total_final_kills_standard"`
				SkeletonDeaths                           float64  `json:"skeleton_deaths"`
				ZombieLosses                             float64  `json:"zombie_losses"`
				BlazeTotalFinalKills                     float64  `json:"blaze_total_final_kills"`
				SpiderFinalKills                         float64  `json:"spider_final_kills"`
				SkeletonKills                            float64  `json:"skeleton_kills"`
				HunterLosses                             float64  `json:"hunter_losses"`
				HerobrineLosses                          float64  `json:"herobrine_losses"`
				ShamanTotalFinalKills                    float64  `json:"shaman_total_final_kills"`
				CreeperDeaths                            float64  `json:"creeper_deaths"`
				EndermanTotalFinalKills                  float64  `json:"enderman_total_final_kills"`
				GolemFinalKills                          float64  `json:"golem_final_kills"`
				PirateFinalKillsStandard                 float64  `json:"pirate_final_kills_standard"`
				EndermanLosses                           float64  `json:"enderman_losses"`
				ShamanFinalKillsStandard                 float64  `json:"shaman_final_kills_standard"`
				PigmanFinalKillsStandard                 float64  `json:"pigman_final_kills_standard"`
				SkeletonLosses                           float64  `json:"skeleton_losses"`
				ShamanDeaths                             float64  `json:"shaman_deaths"`
				HunterFinalKillsStandard                 float64  `json:"hunter_final_kills_standard"`
				PigmanKills                              float64  `json:"pigman_kills"`
				SkeletonTotalFinalKills                  float64  `json:"skeleton_total_final_kills"`
				GolemEnderchestLevel                     float64  `json:"golem_enderchest_level"`
				HerobrineWins                            float64  `json:"herobrine_wins"`
				ZombieFinalKillsStandard                 float64  `json:"zombie_final_kills_standard"`
				BlazeLosses                              float64  `json:"blaze_losses"`
				BlazeDeaths                              float64  `json:"blaze_deaths"`
				SpiderDeaths                             float64  `json:"spider_deaths"`
				GolemLosses                              float64  `json:"golem_losses"`
				HerobrineDeaths                          float64  `json:"herobrine_deaths"`
				SpiderTotalFinalKillsStandard            float64  `json:"spider_total_final_kills_standard"`
				GolemFinalKillsStandard                  float64  `json:"golem_final_kills_standard"`
				HerobrineFinalKills                      float64  `json:"herobrine_final_kills"`
				EndermanDeaths                           float64  `json:"enderman_deaths"`
				HerobrineAKillsStandard                  float64  `json:"herobrine_a_kills_standard"`
				HerobrineADefenderAssists                float64  `json:"herobrine_a_defender_assists"`
				HerobrineATotalKillsStandard             float64  `json:"herobrine_a_total_kills_standard"`
				HerobrineAssists                         float64  `json:"herobrine_assists"`
				HerobrineAActivationsStandard            float64  `json:"herobrine_a_activations_standard"`
				HerobrineKillsMeleeBehindStandard        float64  `json:"herobrine_kills_melee_behind_standard"`
				DefenderAssistsStandard                  float64  `json:"defender_assists_standard"`
				HerobrineActivationsDeathmatchStandard   float64  `json:"herobrine_activations_deathmatch_standard"`
				HerobrineBlocksBrokenStandard            float64  `json:"herobrine_blocks_broken_standard"`
				HerobrineMetersFallenStandard            float64  `json:"herobrine_meters_fallen_standard"`
				HerobrineGamesPlayed                     float64  `json:"herobrine_games_played"`
				HerobrineIronOreBroken                   float64  `json:"herobrine_iron_ore_broken"`
				TotalDeaths                              float64  `json:"total_deaths"`
				HerobrineKillsStandard                   float64  `json:"herobrine_kills_standard"`
				HerobrineBlocksBroken                    float64  `json:"herobrine_blocks_broken"`
				LossesStandard                           float64  `json:"losses_standard"`
				HerobrineKillsMeleeBehind                float64  `json:"herobrine_kills_melee_behind"`
				HerobrineFinalDeathsStandard             float64  `json:"herobrine_final_deaths_standard"`
				TotalKills                               float64  `json:"total_kills"`
				ActivationsStandard                      float64  `json:"activations_standard"`
				HerobrineADefenderKillsStandard          float64  `json:"herobrine_a_defender_kills_standard"`
				HerobrineArrowsFiredStandard             float64  `json:"herobrine_arrows_fired_standard"`
				TreasuresFound                           float64  `json:"treasures_found"`
				HerobrineAKills                          float64  `json:"herobrine_a_kills"`
				ActivationsDeathmatchStandard            float64  `json:"activations_deathmatch_standard"`
				ArrowsFiredStandard                      float64  `json:"arrows_fired_standard"`
				HerobrineADamageDealt                    float64  `json:"herobrine_a_damage_dealt"`
				HerobrineDamageDealtStandard             float64  `json:"herobrine_damage_dealt_standard"`
				HerobrineDefenderKills                   float64  `json:"herobrine_defender_kills"`
				HerobrineBlocksPlaced                    float64  `json:"herobrine_blocks_placed"`
				TimePlayed                               float64  `json:"time_played"`
				DefenderKillsStandard                    float64  `json:"defender_kills_standard"`
				HerobrineBlocksPlacedPreparation         float64  `json:"herobrine_blocks_placed_preparation"`
				HerobrineLossesStandard                  float64  `json:"herobrine_losses_standard"`
				HerobrineATotalKills                     float64  `json:"herobrine_a_total_kills"`
				HerobrineDefenderAssists                 float64  `json:"herobrine_defender_assists"`
				FinalDeaths                              float64  `json:"final_deaths"`
				GamesPlayed                              float64  `json:"games_played"`
				HerobrineKillsMelee                      float64  `json:"herobrine_kills_melee"`
				BlocksPlacedStandard                     float64  `json:"blocks_placed_standard"`
				HerobrineTotalDeaths                     float64  `json:"herobrine_total_deaths"`
				BlocksPlacedPreparation                  float64  `json:"blocks_placed_preparation"`
				HerobrineFinalDeaths                     float64  `json:"herobrine_final_deaths"`
				HerobrineADefenderKills                  float64  `json:"herobrine_a_defender_kills"`
				HerobrineADefenderAssistsStandard        float64  `json:"herobrine_a_defender_assists_standard"`
				BlocksPlacedPreparationStandard          float64  `json:"blocks_placed_preparation_standard"`
				HerobrineBlocksPlacedStandard            float64  `json:"herobrine_blocks_placed_standard"`
				DamageDealt                              float64  `json:"damage_dealt"`
				HerobrineActivations                     float64  `json:"herobrine_activations"`
				TotalKillsStandard                       float64  `json:"total_kills_standard"`
				HerobrineAActivationsDeathmatchStandard  float64  `json:"herobrine_a_activations_deathmatch_standard"`
				TimePlayedStandard                       float64  `json:"time_played_standard"`
				HerobrineAAssists                        float64  `json:"herobrine_a_assists"`
				KillsMeleeStandard                       float64  `json:"kills_melee_standard"`
				HerobrineTreasuresFoundStandard          float64  `json:"herobrine_treasures_found_standard"`
				HerobrineIronOreBrokenStandard           float64  `json:"herobrine_iron_ore_broken_standard"`
				KillsStandard                            float64  `json:"kills_standard"`
				IronOreBrokenStandard                    float64  `json:"iron_ore_broken_standard"`
				TreasuresFoundStandard                   float64  `json:"treasures_found_standard"`
				HerobrineDefenderAssistsStandard         float64  `json:"herobrine_defender_assists_standard"`
				BlocksBrokenStandard                     float64  `json:"blocks_broken_standard"`
				HerobrineMetersWalkedSpeed               float64  `json:"herobrine_meters_walked_speed"`
				DamageDealtStandard                      float64  `json:"damage_dealt_standard"`
				MetersWalkedSpeedStandard                float64  `json:"meters_walked_speed_standard"`
				HerobrineAAssistsStandard                float64  `json:"herobrine_a_assists_standard"`
				HerobrineArrowsFired                     float64  `json:"herobrine_arrows_fired"`
				HerobrineMetersWalked                    float64  `json:"herobrine_meters_walked"`
				FinalDeathsStandard                      float64  `json:"final_deaths_standard"`
				KillsMelee                               float64  `json:"kills_melee"`
				HerobrineTreasuresFound                  float64  `json:"herobrine_treasures_found"`
				DefenderKills                            float64  `json:"defender_kills"`
				HerobrineDamageDealt                     float64  `json:"herobrine_damage_dealt"`
				HerobrineBlocksPlacedPreparationStandard float64  `json:"herobrine_blocks_placed_preparation_standard"`
				HerobrineTimePlayed                      float64  `json:"herobrine_time_played"`
				MetersWalked                             float64  `json:"meters_walked"`
				HerobrineActivationsStandard             float64  `json:"herobrine_activations_standard"`
				MetersFallenStandard                     float64  `json:"meters_fallen_standard"`
				KillsMeleeBehind                         float64  `json:"kills_melee_behind"`
				HerobrineMetersWalkedSpeedStandard       float64  `json:"herobrine_meters_walked_speed_standard"`
				BlocksBroken                             float64  `json:"blocks_broken"`
				Activations                              float64  `json:"activations"`
				AssistsStandard                          float64  `json:"assists_standard"`
				BlocksPlaced                             float64  `json:"blocks_placed"`
				HerobrineTotalDeathsStandard             float64  `json:"herobrine_total_deaths_standard"`
				HerobrineAActivationsDeathmatch          float64  `json:"herobrine_a_activations_deathmatch"`
				MetersWalkedSpeed                        float64  `json:"meters_walked_speed"`
				ArrowsFired                              float64  `json:"arrows_fired"`
				IronOreBroken                            float64  `json:"iron_ore_broken"`
				TotalDeathsStandard                      float64  `json:"total_deaths_standard"`
				HerobrineDefenderKillsStandard           float64  `json:"herobrine_defender_kills_standard"`
				HerobrineTotalKillsStandard              float64  `json:"herobrine_total_kills_standard"`
				HerobrineKillsMeleeStandard              float64  `json:"herobrine_kills_melee_standard"`
				KillsMeleeBehindStandard                 float64  `json:"kills_melee_behind_standard"`
				HerobrineGamesPlayedStandard             float64  `json:"herobrine_games_played_standard"`
				HerobrineMetersFallen                    float64  `json:"herobrine_meters_fallen"`
				HerobrineMetersWalkedStandard            float64  `json:"herobrine_meters_walked_standard"`
				DefenderAssists                          float64  `json:"defender_assists"`
				MetersFallen                             float64  `json:"meters_fallen"`
				HerobrineAKillsMeleeStandard             float64  `json:"herobrine_a_kills_melee_standard"`
				HerobrineADamageDealtStandard            float64  `json:"herobrine_a_damage_dealt_standard"`
				HerobrineAssistsStandard                 float64  `json:"herobrine_assists_standard"`
				HerobrineTimePlayedStandard              float64  `json:"herobrine_time_played_standard"`
				MetersWalkedStandard                     float64  `json:"meters_walked_standard"`
				HerobrineActivationsDeathmatch           float64  `json:"herobrine_activations_deathmatch"`
				HerobrineAKillsMelee                     float64  `json:"herobrine_a_kills_melee"`
				ActivationsDeathmatch                    float64  `json:"activations_deathmatch"`
				GamesPlayedStandard                      float64  `json:"games_played_standard"`
				HerobrineAActivations                    float64  `json:"herobrine_a_activations"`
				HerobrineTotalKills                      float64  `json:"herobrine_total_kills"`
				PhoenixInventory                         struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num6 string `json:"6"`
					Num7 string `json:"7"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"phoenixInventory"`
				PhoenixMetersWalked                     float64 `json:"phoenix_meters_walked"`
				PhoenixBlocksPlacedPreparation          float64 `json:"phoenix_blocks_placed_preparation"`
				BlocksPlacedFaceOff                     float64 `json:"blocks_placed_face_off"`
				GamesPlayedFaceOff                      float64 `json:"games_played_face_off"`
				PhoenixMetersFallen                     float64 `json:"phoenix_meters_fallen"`
				PhoenixGamesPlayedFaceOff               float64 `json:"phoenix_games_played_face_off"`
				PhoenixBlocksPlacedPreparationFaceOff   float64 `json:"phoenix_blocks_placed_preparation_face_off"`
				PhoenixBlocksPlacedFaceOff              float64 `json:"phoenix_blocks_placed_face_off"`
				MetersFallenFaceOff                     float64 `json:"meters_fallen_face_off"`
				PhoenixWinsFaceOff                      float64 `json:"phoenix_wins_face_off"`
				PhoenixMetersWalkedFaceOff              float64 `json:"phoenix_meters_walked_face_off"`
				BlocksPlacedPreparationFaceOff          float64 `json:"blocks_placed_preparation_face_off"`
				PhoenixGamesPlayed                      float64 `json:"phoenix_games_played"`
				PhoenixMetersFallenFaceOff              float64 `json:"phoenix_meters_fallen_face_off"`
				PhoenixWins                             float64 `json:"phoenix_wins"`
				PhoenixBlocksPlaced                     float64 `json:"phoenix_blocks_placed"`
				MetersWalkedFaceOff                     float64 `json:"meters_walked_face_off"`
				PhoenixLossesStandard                   float64 `json:"phoenix_losses_standard"`
				PhoenixMetersWalkedStandard             float64 `json:"phoenix_meters_walked_standard"`
				PhoenixArrowsFired                      float64 `json:"phoenix_arrows_fired"`
				PhoenixMetersFallenStandard             float64 `json:"phoenix_meters_fallen_standard"`
				PhoenixLosses                           float64 `json:"phoenix_losses"`
				PhoenixArrowsFiredStandard              float64 `json:"phoenix_arrows_fired_standard"`
				PhoenixGamesPlayedStandard              float64 `json:"phoenix_games_played_standard"`
				PhoenixWinsStandard                     float64 `json:"phoenix_wins_standard"`
				FinalKills                              float64 `json:"final_kills"`
				SkeletonWinsStandard                    float64 `json:"skeleton_wins_standard"`
				EndermanWinsStandard                    float64 `json:"enderman_wins_standard"`
				ZombieWinsStandard                      float64 `json:"zombie_wins_standard"`
				BlazeWinsStandard                       float64 `json:"blaze_wins_standard"`
				ArcanistWinsStandard                    float64 `json:"arcanist_wins_standard"`
				HunterWinsStandard                      float64 `json:"hunter_wins_standard"`
				HerobrineWinsStandard                   float64 `json:"herobrine_wins_standard"`
				ShamanWinsStandard                      float64 `json:"shaman_wins_standard"`
				FinalKillsStandard                      float64 `json:"final_kills_standard"`
				CreeperWinsStandard                     float64 `json:"creeper_wins_standard"`
				TotalFinalKillsStandard                 float64 `json:"total_final_kills_standard"`
				SpiderWinsStandard                      float64 `json:"spider_wins_standard"`
				PigmanWinsStandard                      float64 `json:"pigman_wins_standard"`
				GolemWinsStandard                       float64 `json:"golem_wins_standard"`
				TotalFinalKills                         float64 `json:"total_final_kills"`
				WerewolfBlocksPlacedPreparationStandard float64 `json:"werewolf_blocks_placed_preparation_standard"`
				WerewolfPotionsDrunkStandard            float64 `json:"werewolf_potions_drunk_standard"`
				WerewolfMetersWalkedSpeed               float64 `json:"werewolf_meters_walked_speed"`
				WerewolfGamesPlayed                     float64 `json:"werewolf_games_played"`
				PotionsDrunkStandard                    float64 `json:"potions_drunk_standard"`
				WerewolfBlocksPlacedStandard            float64 `json:"werewolf_blocks_placed_standard"`
				WerewolfLossesStandard                  float64 `json:"werewolf_losses_standard"`
				WerewolfLosses                          float64 `json:"werewolf_losses"`
				WerewolfMetersWalkedStandard            float64 `json:"werewolf_meters_walked_standard"`
				WerewolfTreasuresFound                  float64 `json:"werewolf_treasures_found"`
				WerewolfTimePlayed                      float64 `json:"werewolf_time_played"`
				WerewolfIronOreBrokenStandard           float64 `json:"werewolf_iron_ore_broken_standard"`
				WerewolfMetersFallenStandard            float64 `json:"werewolf_meters_fallen_standard"`
				WerewolfBlocksBrokenStandard            float64 `json:"werewolf_blocks_broken_standard"`
				WerewolfTreasuresFoundStandard          float64 `json:"werewolf_treasures_found_standard"`
				WerewolfMetersWalkedSpeedStandard       float64 `json:"werewolf_meters_walked_speed_standard"`
				WerewolfTimePlayedStandard              float64 `json:"werewolf_time_played_standard"`
				WerewolfPotionsDrunk                    float64 `json:"werewolf_potions_drunk"`
				WerewolfGamesPlayedStandard             float64 `json:"werewolf_games_played_standard"`
				WerewolfMetersFallen                    float64 `json:"werewolf_meters_fallen"`
				WerewolfMetersWalked                    float64 `json:"werewolf_meters_walked"`
				WerewolfBlocksBroken                    float64 `json:"werewolf_blocks_broken"`
				PotionsDrunk                            float64 `json:"potions_drunk"`
				WerewolfIronOreBroken                   float64 `json:"werewolf_iron_ore_broken"`
				WerewolfBlocksPlaced                    float64 `json:"werewolf_blocks_placed"`
				WerewolfBlocksPlacedPreparation         float64 `json:"werewolf_blocks_placed_preparation"`
				WerewolfDefenderKillsStandard           float64 `json:"werewolf_defender_kills_standard"`
				WerewolfTotalDeathsStandard             float64 `json:"werewolf_total_deaths_standard"`
				DiamondOreBrokenStandard                float64 `json:"diamond_ore_broken_standard"`
				WerewolfTotalDeaths                     float64 `json:"werewolf_total_deaths"`
				WerewolfAAssistsStandard                float64 `json:"werewolf_a_assists_standard"`
				WerewolfKills                           float64 `json:"werewolf_kills"`
				WerewolfBAssists                        float64 `json:"werewolf_b_assists"`
				WoodChopped                             float64 `json:"wood_chopped"`
				WerewolfDefenderAssistsStandard         float64 `json:"werewolf_defender_assists_standard"`
				WerewolfAActivationsStandard            float64 `json:"werewolf_a_activations_standard"`
				WoodChoppedStandard                     float64 `json:"wood_chopped_standard"`
				WerewolfAssists                         float64 `json:"werewolf_assists"`
				WerewolfADamageDealt                    float64 `json:"werewolf_a_damage_dealt"`
				WerewolfWoodChopped                     float64 `json:"werewolf_wood_chopped"`
				WerewolfDeathsStandard                  float64 `json:"werewolf_deaths_standard"`
				WerewolfDamageDealtStandard             float64 `json:"werewolf_damage_dealt_standard"`
				WerewolfDiamondOreBrokenStandard        float64 `json:"werewolf_diamond_ore_broken_standard"`
				WerewolfWoodChoppedStandard             float64 `json:"werewolf_wood_chopped_standard"`
				WerewolfDefenderAssists                 float64 `json:"werewolf_defender_assists"`
				WerewolfDamageDealt                     float64 `json:"werewolf_damage_dealt"`
				WerewolfGActivationsStandard            float64 `json:"werewolf_g_activations_standard"`
				WerewolfTotalKills                      float64 `json:"werewolf_total_kills"`
				DiamondOreBroken                        float64 `json:"diamond_ore_broken"`
				WerewolfATotalKills                     float64 `json:"werewolf_a_total_kills"`
				WerewolfBTotalKills                     float64 `json:"werewolf_b_total_kills"`
				WerewolfADefenderAssists                float64 `json:"werewolf_a_defender_assists"`
				WerewolfBAssistsStandard                float64 `json:"werewolf_b_assists_standard"`
				WerewolfKillsMeleeStandard              float64 `json:"werewolf_kills_melee_standard"`
				WerewolfADamageDealtStandard            float64 `json:"werewolf_a_damage_dealt_standard"`
				WerewolfDeaths                          float64 `json:"werewolf_deaths"`
				WerewolfKillsStandard                   float64 `json:"werewolf_kills_standard"`
				WerewolfAssistsStandard                 float64 `json:"werewolf_assists_standard"`
				WerewolfDiamondOreBroken                float64 `json:"werewolf_diamond_ore_broken"`
				DeathsStandard                          float64 `json:"deaths_standard"`
				WerewolfATotalKillsStandard             float64 `json:"werewolf_a_total_kills_standard"`
				WerewolfDefenderKills                   float64 `json:"werewolf_defender_kills"`
				WerewolfActivationsStandard             float64 `json:"werewolf_activations_standard"`
				WerewolfGActivations                    float64 `json:"werewolf_g_activations"`
				WerewolfTotalKillsStandard              float64 `json:"werewolf_total_kills_standard"`
				WerewolfBTotalKillsStandard             float64 `json:"werewolf_b_total_kills_standard"`
				WerewolfKillsMelee                      float64 `json:"werewolf_kills_melee"`
				WerewolfAAssists                        float64 `json:"werewolf_a_assists"`
				WerewolfAActivations                    float64 `json:"werewolf_a_activations"`
				WerewolfADefenderAssistsStandard        float64 `json:"werewolf_a_defender_assists_standard"`
				WerewolfActivations                     float64 `json:"werewolf_activations"`
				ArcanistFinalAssists                    float64 `json:"arcanist_final_assists"`
				SkeletonFinalAssists                    float64 `json:"skeleton_final_assists"`
				HunterFinalAssists                      float64 `json:"hunter_final_assists"`
				PirateFinalAssists                      float64 `json:"pirate_final_assists"`
				GolemFinalAssists                       float64 `json:"golem_final_assists"`
				BlazeFinalAssists                       float64 `json:"blaze_final_assists"`
				SpiderFinalAssists                      float64 `json:"spider_final_assists"`
				HerobrineFinalAssists                   float64 `json:"herobrine_final_assists"`
				ActiveChallengeMap                      string  `json:"activeChallengeMap"`
				WinsStandard                            float64 `json:"wins_standard"`
				WerewolfWinsStandard                    float64 `json:"werewolf_wins_standard"`
				WerewolfWins                            float64 `json:"werewolf_wins"`
				HunterDeathsStandard                    float64 `json:"hunter_deaths_standard"`
				HunterBlocksPlacedPreparation           float64 `json:"hunter_blocks_placed_preparation"`
				HunterArrowsFiredStandard               float64 `json:"hunter_arrows_fired_standard"`
				HunterTimePlayed                        float64 `json:"hunter_time_played"`
				HunterMetersFallenStandard              float64 `json:"hunter_meters_fallen_standard"`
				HunterMetersWalkedStandard              float64 `json:"hunter_meters_walked_standard"`
				HunterBlocksPlaced                      float64 `json:"hunter_blocks_placed"`
				HunterBlocksPlacedPreparationStandard   float64 `json:"hunter_blocks_placed_preparation_standard"`
				HunterGamesPlayedStandard               float64 `json:"hunter_games_played_standard"`
				HunterTimePlayedStandard                float64 `json:"hunter_time_played_standard"`
				HunterMetersWalked                      float64 `json:"hunter_meters_walked"`
				HunterLossesStandard                    float64 `json:"hunter_losses_standard"`
				HunterTotalDeathsStandard               float64 `json:"hunter_total_deaths_standard"`
				HunterArrowsFired                       float64 `json:"hunter_arrows_fired"`
				HunterMetersFallen                      float64 `json:"hunter_meters_fallen"`
				HunterBlocksPlacedStandard              float64 `json:"hunter_blocks_placed_standard"`
				HunterGamesPlayed                       float64 `json:"hunter_games_played"`
				HunterTotalDeaths                       float64 `json:"hunter_total_deaths"`
				HunterActivationsDeathmatch             float64 `json:"hunter_activations_deathmatch"`
				HunterKillsMeleeBehind                  float64 `json:"hunter_kills_melee_behind"`
				HunterFinalKillsMeleeBehindStandard     float64 `json:"hunter_final_kills_melee_behind_standard"`
				HunterTotalKills                        float64 `json:"hunter_total_kills"`
				HunterBlocksBroken                      float64 `json:"hunter_blocks_broken"`
				HunterAActivationsDeathmatchStandard    float64 `json:"hunter_a_activations_deathmatch_standard"`
				FinalKillsMelee                         float64 `json:"final_kills_melee"`
				BreadEatenStandard                      float64 `json:"bread_eaten_standard"`
				HunterAssistsStandard                   float64 `json:"hunter_assists_standard"`
				HunterBlocksBrokenStandard              float64 `json:"hunter_blocks_broken_standard"`
				HunterDefenderAssistsStandard           float64 `json:"hunter_defender_assists_standard"`
				FoodEatenStandard                       float64 `json:"food_eaten_standard"`
				GoldenApplesEaten                       float64 `json:"golden_apples_eaten"`
				HunterAActivationsDeathmatch            float64 `json:"hunter_a_activations_deathmatch"`
				HunterBreadEaten                        float64 `json:"hunter_bread_eaten"`
				HunterKillsStandard                     float64 `json:"hunter_kills_standard"`
				FinalAssists                            float64 `json:"final_assists"`
				HunterArrowsHitStandard                 float64 `json:"hunter_arrows_hit_standard"`
				FinalKillsMeleeBehind                   float64 `json:"final_kills_melee_behind"`
				HunterAActivationsStandard              float64 `json:"hunter_a_activations_standard"`
				HunterDefenderKills                     float64 `json:"hunter_defender_kills"`
				FinalKillsMeleeBehindStandard           float64 `json:"final_kills_melee_behind_standard"`
				HunterBreadEatenStandard                float64 `json:"hunter_bread_eaten_standard"`
				FoodEaten                               float64 `json:"food_eaten"`
				HunterMetersWalkedSpeedStandard         float64 `json:"hunter_meters_walked_speed_standard"`
				HunterFoodEaten                         float64 `json:"hunter_food_eaten"`
				HunterSteaksEaten                       float64 `json:"hunter_steaks_eaten"`
				HunterGoldenApplesEatenStandard         float64 `json:"hunter_golden_apples_eaten_standard"`
				HunterKillsMeleeStandard                float64 `json:"hunter_kills_melee_standard"`
				HunterTotalKillsStandard                float64 `json:"hunter_total_kills_standard"`
				HunterActivationsStandard               float64 `json:"hunter_activations_standard"`
				HunterFinalAssistsStandard              float64 `json:"hunter_final_assists_standard"`
				FinalAssistsStandard                    float64 `json:"final_assists_standard"`
				HunterActivations                       float64 `json:"hunter_activations"`
				HunterArrowsHit                         float64 `json:"hunter_arrows_hit"`
				HunterMetersWalkedSpeed                 float64 `json:"hunter_meters_walked_speed"`
				BreadEaten                              float64 `json:"bread_eaten"`
				SteaksEaten                             float64 `json:"steaks_eaten"`
				SteaksEatenStandard                     float64 `json:"steaks_eaten_standard"`
				HunterKillsMeleeBehindStandard          float64 `json:"hunter_kills_melee_behind_standard"`
				HunterAActivations                      float64 `json:"hunter_a_activations"`
				HunterActivationsDeathmatchStandard     float64 `json:"hunter_activations_deathmatch_standard"`
				HunterDefenderKillsStandard             float64 `json:"hunter_defender_kills_standard"`
				FinalKillsMeleeStandard                 float64 `json:"final_kills_melee_standard"`
				HunterKillsMelee                        float64 `json:"hunter_kills_melee"`
				HunterFinalKillsMelee                   float64 `json:"hunter_final_kills_melee"`
				ArrowsHitStandard                       float64 `json:"arrows_hit_standard"`
				HunterDefenderAssists                   float64 `json:"hunter_defender_assists"`
				HunterFinalKillsMeleeStandard           float64 `json:"hunter_final_kills_melee_standard"`
				ArrowsHit                               float64 `json:"arrows_hit"`
				HunterGoldenApplesEaten                 float64 `json:"hunter_golden_apples_eaten"`
				HunterAssists                           float64 `json:"hunter_assists"`
				HunterFinalKillsMeleeBehind             float64 `json:"hunter_final_kills_melee_behind"`
				GoldenApplesEatenStandard               float64 `json:"golden_apples_eaten_standard"`
				HunterFoodEatenStandard                 float64 `json:"hunter_food_eaten_standard"`
				HunterSteaksEatenStandard               float64 `json:"hunter_steaks_eaten_standard"`
				DreadlordInventory                      struct {
					Num0 string `json:"0"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num7 string `json:"7"`
					Num8 string `json:"8"`
					Num9 string `json:"9"`
				} `json:"dreadlordInventory"`
				DreadlordDamageDealtStandard             float64 `json:"dreadlord_damage_dealt_standard"`
				DreadlordTimePlayed                      float64 `json:"dreadlord_time_played"`
				DreadlordAssistsStandard                 float64 `json:"dreadlord_assists_standard"`
				DreadlordActivationsDeathmatchStandard   float64 `json:"dreadlord_activations_deathmatch_standard"`
				DreadlordTreasuresFound                  float64 `json:"dreadlord_treasures_found"`
				DreadlordAActivationsDeathmatchStandard  float64 `json:"dreadlord_a_activations_deathmatch_standard"`
				DreadlordMetersWalkedStandard            float64 `json:"dreadlord_meters_walked_standard"`
				DreadlordAFinalAssistsStandard           float64 `json:"dreadlord_a_final_assists_standard"`
				DreadlordATotalFinalKillsStandard        float64 `json:"dreadlord_a_total_final_kills_standard"`
				DreadlordTotalFinalKillsStandard         float64 `json:"dreadlord_total_final_kills_standard"`
				DreadlordAFinalKillsMelee                float64 `json:"dreadlord_a_final_kills_melee"`
				DreadlordActivationsDeathmatch           float64 `json:"dreadlord_activations_deathmatch"`
				DreadlordFinalDeaths                     float64 `json:"dreadlord_final_deaths"`
				DreadlordTotalKills                      float64 `json:"dreadlord_total_kills"`
				DreadlordKills                           float64 `json:"dreadlord_kills"`
				DreadlordDefenderKillsStandard           float64 `json:"dreadlord_defender_kills_standard"`
				DreadlordFinalDeathsStandard             float64 `json:"dreadlord_final_deaths_standard"`
				DreadlordADamageDealtStandard            float64 `json:"dreadlord_a_damage_dealt_standard"`
				DreadlordTimePlayedStandard              float64 `json:"dreadlord_time_played_standard"`
				DreadlordKillsMeleeStandard              float64 `json:"dreadlord_kills_melee_standard"`
				DreadlordATotalFinalKills                float64 `json:"dreadlord_a_total_final_kills"`
				DreadlordFinalKillsMelee                 float64 `json:"dreadlord_final_kills_melee"`
				DreadlordTotalKillsStandard              float64 `json:"dreadlord_total_kills_standard"`
				DreadlordMetersWalked                    float64 `json:"dreadlord_meters_walked"`
				WitherDamageStandard                     float64 `json:"wither_damage_standard"`
				DreadlordAAssistsStandard                float64 `json:"dreadlord_a_assists_standard"`
				DreadlordMetersFallenStandard            float64 `json:"dreadlord_meters_fallen_standard"`
				DreadlordADamageDealt                    float64 `json:"dreadlord_a_damage_dealt"`
				DreadlordKillsMelee                      float64 `json:"dreadlord_kills_melee"`
				DreadlordAFinalKillsStandard             float64 `json:"dreadlord_a_final_kills_standard"`
				DreadlordGamesPlayedStandard             float64 `json:"dreadlord_games_played_standard"`
				DreadlordDeaths                          float64 `json:"dreadlord_deaths"`
				DreadlordAFinalAssists                   float64 `json:"dreadlord_a_final_assists"`
				DreadlordFinalAssists                    float64 `json:"dreadlord_final_assists"`
				DreadlordBlocksPlacedPreparation         float64 `json:"dreadlord_blocks_placed_preparation"`
				DreadlordActivationsStandard             float64 `json:"dreadlord_activations_standard"`
				DreadlordATotalKills                     float64 `json:"dreadlord_a_total_kills"`
				DreadlordTreasuresFoundStandard          float64 `json:"dreadlord_treasures_found_standard"`
				DreadlordWitherDamage                    float64 `json:"dreadlord_wither_damage"`
				DreadlordBlocksBroken                    float64 `json:"dreadlord_blocks_broken"`
				DreadlordWitherDamageStandard            float64 `json:"dreadlord_wither_damage_standard"`
				DreadlordKillsMeleeBehind                float64 `json:"dreadlord_kills_melee_behind"`
				DreadlordMetersWalkedSpeedStandard       float64 `json:"dreadlord_meters_walked_speed_standard"`
				DreadlordKillsStandard                   float64 `json:"dreadlord_kills_standard"`
				DreadlordMetersFallen                    float64 `json:"dreadlord_meters_fallen"`
				DreadlordTotalDeaths                     float64 `json:"dreadlord_total_deaths"`
				DreadlordAActivationsDeathmatch          float64 `json:"dreadlord_a_activations_deathmatch"`
				DreadlordFinalKillsMeleeBehind           float64 `json:"dreadlord_final_kills_melee_behind"`
				DreadlordTotalFinalKills                 float64 `json:"dreadlord_total_final_kills"`
				DreadlordAFinalKills                     float64 `json:"dreadlord_a_final_kills"`
				DreadlordFinalKillsMeleeStandard         float64 `json:"dreadlord_final_kills_melee_standard"`
				DreadlordFinalAssistsStandard            float64 `json:"dreadlord_final_assists_standard"`
				DreadlordDeathsStandard                  float64 `json:"dreadlord_deaths_standard"`
				DreadlordADefenderAssistsStandard        float64 `json:"dreadlord_a_defender_assists_standard"`
				DreadlordBlocksPlaced                    float64 `json:"dreadlord_blocks_placed"`
				DreadlordFinalKills                      float64 `json:"dreadlord_final_kills"`
				DreadlordActivations                     float64 `json:"dreadlord_activations"`
				DreadlordBlocksPlacedStandard            float64 `json:"dreadlord_blocks_placed_standard"`
				DreadlordDefenderAssists                 float64 `json:"dreadlord_defender_assists"`
				DreadlordDefenderKills                   float64 `json:"dreadlord_defender_kills"`
				DreadlordATotalKillsStandard             float64 `json:"dreadlord_a_total_kills_standard"`
				DreadlordAssists                         float64 `json:"dreadlord_assists"`
				DreadlordIronOreBroken                   float64 `json:"dreadlord_iron_ore_broken"`
				DreadlordLosses                          float64 `json:"dreadlord_losses"`
				DreadlordBlocksPlacedPreparationStandard float64 `json:"dreadlord_blocks_placed_preparation_standard"`
				DreadlordTotalDeathsStandard             float64 `json:"dreadlord_total_deaths_standard"`
				DreadlordADefenderAssists                float64 `json:"dreadlord_a_defender_assists"`
				DreadlordPotionsDrunk                    float64 `json:"dreadlord_potions_drunk"`
				DreadlordMetersWalkedSpeed               float64 `json:"dreadlord_meters_walked_speed"`
				DreadlordBlocksBrokenStandard            float64 `json:"dreadlord_blocks_broken_standard"`
				DreadlordGamesPlayed                     float64 `json:"dreadlord_games_played"`
				DreadlordAFinalKillsMeleeBehindStandard  float64 `json:"dreadlord_a_final_kills_melee_behind_standard"`
				DreadlordPotionsDrunkStandard            float64 `json:"dreadlord_potions_drunk_standard"`
				WitherDamage                             float64 `json:"wither_damage"`
				DreadlordAFinalKillsMeleeStandard        float64 `json:"dreadlord_a_final_kills_melee_standard"`
				DreadlordFinalKillsStandard              float64 `json:"dreadlord_final_kills_standard"`
				DreadlordDamageDealt                     float64 `json:"dreadlord_damage_dealt"`
				DreadlordAFinalKillsMeleeBehind          float64 `json:"dreadlord_a_final_kills_melee_behind"`
				DreadlordLossesStandard                  float64 `json:"dreadlord_losses_standard"`
				DreadlordAActivations                    float64 `json:"dreadlord_a_activations"`
				DreadlordKillsMeleeBehindStandard        float64 `json:"dreadlord_kills_melee_behind_standard"`
				DreadlordAAssists                        float64 `json:"dreadlord_a_assists"`
				DreadlordAActivationsStandard            float64 `json:"dreadlord_a_activations_standard"`
				DreadlordFinalKillsMeleeBehindStandard   float64 `json:"dreadlord_final_kills_melee_behind_standard"`
				DreadlordDefenderAssistsStandard         float64 `json:"dreadlord_defender_assists_standard"`
				DreadlordIronOreBrokenStandard           float64 `json:"dreadlord_iron_ore_broken_standard"`
				DreadlordFoodEatenStandard               float64 `json:"dreadlord_food_eaten_standard"`
				IronArmorGiftedStandard                  float64 `json:"iron_armor_gifted_standard"`
				DreadlordBreadCraftedStandard            float64 `json:"dreadlord_bread_crafted_standard"`
				DreadlordIronArmorGifted                 float64 `json:"dreadlord_iron_armor_gifted"`
				DreadlordSteaksEaten                     float64 `json:"dreadlord_steaks_eaten"`
				BreadCraftedStandard                     float64 `json:"bread_crafted_standard"`
				DreadlordBreadCrafted                    float64 `json:"dreadlord_bread_crafted"`
				DreadlordWitherKillsStandard             float64 `json:"dreadlord_wither_kills_standard"`
				DreadlordWitherKills                     float64 `json:"dreadlord_wither_kills"`
				DreadlordSteaksEatenStandard             float64 `json:"dreadlord_steaks_eaten_standard"`
				IronArmorGifted                          float64 `json:"iron_armor_gifted"`
				BreadCrafted                             float64 `json:"bread_crafted"`
				DreadlordAKillsMeleeStandard             float64 `json:"dreadlord_a_kills_melee_standard"`
				WitherKillsStandard                      float64 `json:"wither_kills_standard"`
				DreadlordAKillsMelee                     float64 `json:"dreadlord_a_kills_melee"`
				WitherKills                              float64 `json:"wither_kills"`
				DreadlordIronArmorGiftedStandard         float64 `json:"dreadlord_iron_armor_gifted_standard"`
				DreadlordFoodEaten                       float64 `json:"dreadlord_food_eaten"`
				DreadlordAKillsStandard                  float64 `json:"dreadlord_a_kills_standard"`
				DreadlordAKills                          float64 `json:"dreadlord_a_kills"`
				DreadlordArrowsHitStandard               float64 `json:"dreadlord_arrows_hit_standard"`
				ApplesEaten                              float64 `json:"apples_eaten"`
				DreadlordArrowsFired                     float64 `json:"dreadlord_arrows_fired"`
				ApplesEatenStandard                      float64 `json:"apples_eaten_standard"`
				DreadlordBreadEaten                      float64 `json:"dreadlord_bread_eaten"`
				DreadlordADefenderKillsStandard          float64 `json:"dreadlord_a_defender_kills_standard"`
				DreadlordArrowsHit                       float64 `json:"dreadlord_arrows_hit"`
				DreadlordArrowsFiredStandard             float64 `json:"dreadlord_arrows_fired_standard"`
				DreadlordADefenderKills                  float64 `json:"dreadlord_a_defender_kills"`
				DreadlordGoldenApplesEaten               float64 `json:"dreadlord_golden_apples_eaten"`
				DreadlordApplesEatenStandard             float64 `json:"dreadlord_apples_eaten_standard"`
				DreadlordBreadEatenStandard              float64 `json:"dreadlord_bread_eaten_standard"`
				DreadlordApplesEaten                     float64 `json:"dreadlord_apples_eaten"`
				DreadlordGoldenApplesEatenStandard       float64 `json:"dreadlord_golden_apples_eaten_standard"`
				MolemanWitherDamageStandard              float64 `json:"moleman_wither_damage_standard"`
				MolemanDefenderKills                     float64 `json:"moleman_defender_kills"`
				MolemanAAssists                          float64 `json:"moleman_a_assists"`
				MolemanAActivationsDeathmatchStandard    float64 `json:"moleman_a_activations_deathmatch_standard"`
				MolemanMetersWalked                      float64 `json:"moleman_meters_walked"`
				MolemanBlocksPlacedPreparationStandard   float64 `json:"moleman_blocks_placed_preparation_standard"`
				MolemanAKillsMeleeStandard               float64 `json:"moleman_a_kills_melee_standard"`
				MolemanATotalFinalKillsStandard          float64 `json:"moleman_a_total_final_kills_standard"`
				MolemanActivationsDeathmatchStandard     float64 `json:"moleman_activations_deathmatch_standard"`
				MolemanBlocksPlacedStandard              float64 `json:"moleman_blocks_placed_standard"`
				MolemanFinalAssists                      float64 `json:"moleman_final_assists"`
				MolemanAssists                           float64 `json:"moleman_assists"`
				MolemanTotalDeathsStandard               float64 `json:"moleman_total_deaths_standard"`
				MolemanAActivations                      float64 `json:"moleman_a_activations"`
				MolemanTotalFinalKillsStandard           float64 `json:"moleman_total_final_kills_standard"`
				MolemanBActivations                      float64 `json:"moleman_b_activations"`
				MolemanFinalKills                        float64 `json:"moleman_final_kills"`
				MolemanAKillsStandard                    float64 `json:"moleman_a_kills_standard"`
				MolemanFoodEaten                         float64 `json:"moleman_food_eaten"`
				MolemanPotionsDrunkStandard              float64 `json:"moleman_potions_drunk_standard"`
				MolemanATotalFinalKills                  float64 `json:"moleman_a_total_final_kills"`
				MolemanFinalAssistsStandard              float64 `json:"moleman_final_assists_standard"`
				MolemanAKillsMelee                       float64 `json:"moleman_a_kills_melee"`
				MolemanBActivationsStandard              float64 `json:"moleman_b_activations_standard"`
				DefenderFinalAssists                     float64 `json:"defender_final_assists"`
				MolemanBlocksBroken                      float64 `json:"moleman_blocks_broken"`
				MolemanFinalKillsMelee                   float64 `json:"moleman_final_kills_melee"`
				MolemanGamesPlayed                       float64 `json:"moleman_games_played"`
				MolemanWitherDamage                      float64 `json:"moleman_wither_damage"`
				MolemanMetersFallen                      float64 `json:"moleman_meters_fallen"`
				MolemanBreadEaten                        float64 `json:"moleman_bread_eaten"`
				MolemanFinalKillsMeleeBehindStandard     float64 `json:"moleman_final_kills_melee_behind_standard"`
				MolemanATotalKillsStandard               float64 `json:"moleman_a_total_kills_standard"`
				MolemanIronOreBroken                     float64 `json:"moleman_iron_ore_broken"`
				MolemanFinalKillsMeleeBehind             float64 `json:"moleman_final_kills_melee_behind"`
				MolemanAFinalAssists                     float64 `json:"moleman_a_final_assists"`
				MolemanGamesPlayedStandard               float64 `json:"moleman_games_played_standard"`
				MolemanDefenderKillsStandard             float64 `json:"moleman_defender_kills_standard"`
				MolemanArrowsHit                         float64 `json:"moleman_arrows_hit"`
				MolemanDeaths                            float64 `json:"moleman_deaths"`
				MolemanAActivationsDeathmatch            float64 `json:"moleman_a_activations_deathmatch"`
				MolemanArrowsFired                       float64 `json:"moleman_arrows_fired"`
				MolemanLosses                            float64 `json:"moleman_losses"`
				MolemanBreadCrafted                      float64 `json:"moleman_bread_crafted"`
				MolemanActivationsDeathmatch             float64 `json:"moleman_activations_deathmatch"`
				MolemanBlocksPlaced                      float64 `json:"moleman_blocks_placed"`
				MolemanADamageDealtStandard              float64 `json:"moleman_a_damage_dealt_standard"`
				MolemanTotalKills                        float64 `json:"moleman_total_kills"`
				MolemanFoodEatenStandard                 float64 `json:"moleman_food_eaten_standard"`
				MolemanSteaksEatenStandard               float64 `json:"moleman_steaks_eaten_standard"`
				MolemanAssistsStandard                   float64 `json:"moleman_assists_standard"`
				MolemanTreasuresFoundStandard            float64 `json:"moleman_treasures_found_standard"`
				MolemanMetersWalkedStandard              float64 `json:"moleman_meters_walked_standard"`
				MolemanTotalFinalKills                   float64 `json:"moleman_total_final_kills"`
				MolemanTotalDeaths                       float64 `json:"moleman_total_deaths"`
				MolemanFinalDeathsStandard               float64 `json:"moleman_final_deaths_standard"`
				MolemanBreadCraftedStandard              float64 `json:"moleman_bread_crafted_standard"`
				MolemanFinalKillsMeleeStandard           float64 `json:"moleman_final_kills_melee_standard"`
				MolemanDefenderFinalAssists              float64 `json:"moleman_defender_final_assists"`
				MolemanDefenderFinalAssistsStandard      float64 `json:"moleman_defender_final_assists_standard"`
				MolemanABlocksBroken                     float64 `json:"moleman_a_blocks_broken"`
				MolemanDamageDealtStandard               float64 `json:"moleman_damage_dealt_standard"`
				MolemanADamageDealt                      float64 `json:"moleman_a_damage_dealt"`
				MolemanTimePlayedStandard                float64 `json:"moleman_time_played_standard"`
				MolemanActivationsStandard               float64 `json:"moleman_activations_standard"`
				MolemanTotalKillsStandard                float64 `json:"moleman_total_kills_standard"`
				MolemanFinalKillsStandard                float64 `json:"moleman_final_kills_standard"`
				MolemanSteaksEaten                       float64 `json:"moleman_steaks_eaten"`
				MolemanADefenderFinalAssists             float64 `json:"moleman_a_defender_final_assists"`
				MolemanADefenderFinalAssistsStandard     float64 `json:"moleman_a_defender_final_assists_standard"`
				MolemanMetersWalkedSpeed                 float64 `json:"moleman_meters_walked_speed"`
				MolemanMetersWalkedSpeedStandard         float64 `json:"moleman_meters_walked_speed_standard"`
				MolemanMetersFallenStandard              float64 `json:"moleman_meters_fallen_standard"`
				MolemanFinalDeaths                       float64 `json:"moleman_final_deaths"`
				MolemanBlocksPlacedPreparation           float64 `json:"moleman_blocks_placed_preparation"`
				MolemanPotionsDrunk                      float64 `json:"moleman_potions_drunk"`
				MolemanBlocksBrokenStandard              float64 `json:"moleman_blocks_broken_standard"`
				MolemanTimePlayed                        float64 `json:"moleman_time_played"`
				DefenderFinalAssistsStandard             float64 `json:"defender_final_assists_standard"`
				MolemanAKills                            float64 `json:"moleman_a_kills"`
				MolemanTreasuresFound                    float64 `json:"moleman_treasures_found"`
				MolemanDamageDealt                       float64 `json:"moleman_damage_dealt"`
				MolemanKills                             float64 `json:"moleman_kills"`
				MolemanKillsMeleeStandard                float64 `json:"moleman_kills_melee_standard"`
				MolemanATotalKills                       float64 `json:"moleman_a_total_kills"`
				MolemanAActivationsStandard              float64 `json:"moleman_a_activations_standard"`
				MolemanArrowsFiredStandard               float64 `json:"moleman_arrows_fired_standard"`
				MolemanAFinalAssistsStandard             float64 `json:"moleman_a_final_assists_standard"`
				MolemanAAssistsStandard                  float64 `json:"moleman_a_assists_standard"`
				MolemanLossesStandard                    float64 `json:"moleman_losses_standard"`
				MolemanBreadEatenStandard                float64 `json:"moleman_bread_eaten_standard"`
				MolemanABlocksBrokenStandard             float64 `json:"moleman_a_blocks_broken_standard"`
				MolemanKillsMelee                        float64 `json:"moleman_kills_melee"`
				MolemanKillsStandard                     float64 `json:"moleman_kills_standard"`
				MolemanIronOreBrokenStandard             float64 `json:"moleman_iron_ore_broken_standard"`
				MolemanArrowsHitStandard                 float64 `json:"moleman_arrows_hit_standard"`
				MolemanActivations                       float64 `json:"moleman_activations"`
				MolemanDeathsStandard                    float64 `json:"moleman_deaths_standard"`
				MolemanBTotalFinalKills                  float64 `json:"moleman_b_total_final_kills"`
				MolemanAbsorptionPotionsDrunkStandard    float64 `json:"moleman_absorption_potions_drunk_standard"`
				MolemanBFinalAssists                     float64 `json:"moleman_b_final_assists"`
				AbsorptionPotionsDrunkStandard           float64 `json:"absorption_potions_drunk_standard"`
				MolemanAbsorptionPotionsDrunk            float64 `json:"moleman_absorption_potions_drunk"`
				AbsorptionPotionsDrunk                   float64 `json:"absorption_potions_drunk"`
				MolemanBTotalFinalKillsStandard          float64 `json:"moleman_b_total_final_kills_standard"`
				MolemanBFinalAssistsStandard             float64 `json:"moleman_b_final_assists_standard"`
				SpiderGamesPlayedStandard                float64 `json:"spider_games_played_standard"`
				SpiderLossesStandard                     float64 `json:"spider_losses_standard"`
				SpiderGamesPlayed                        float64 `json:"spider_games_played"`
				SpiderTotalDeathsStandard                float64 `json:"spider_total_deaths_standard"`
				SpiderTotalDeaths                        float64 `json:"spider_total_deaths"`
				SpiderDeathsStandard                     float64 `json:"spider_deaths_standard"`
				SpiderTimePlayed                         float64 `json:"spider_time_played"`
				SpiderTimePlayedStandard                 float64 `json:"spider_time_played_standard"`
				SpiderMetersWalkedStandard               float64 `json:"spider_meters_walked_standard"`
				SpiderFinalDeathsStandard                float64 `json:"spider_final_deaths_standard"`
				SpiderMetersWalked                       float64 `json:"spider_meters_walked"`
				SpiderFinalDeaths                        float64 `json:"spider_final_deaths"`
				SpiderMetersFallenStandard               float64 `json:"spider_meters_fallen_standard"`
				SpiderBlocksPlaced                       float64 `json:"spider_blocks_placed"`
				SpiderBlocksPlacedPreparationStandard    float64 `json:"spider_blocks_placed_preparation_standard"`
				SpiderBlocksPlacedStandard               float64 `json:"spider_blocks_placed_standard"`
				SpiderIronOreBroken                      float64 `json:"spider_iron_ore_broken"`
				SpiderMetersFallen                       float64 `json:"spider_meters_fallen"`
				SpiderBlocksBrokenStandard               float64 `json:"spider_blocks_broken_standard"`
				SpiderBlocksPlacedPreparation            float64 `json:"spider_blocks_placed_preparation"`
				SpiderTreasuresFound                     float64 `json:"spider_treasures_found"`
				SpiderIronOreBrokenStandard              float64 `json:"spider_iron_ore_broken_standard"`
				SpiderTreasuresFoundStandard             float64 `json:"spider_treasures_found_standard"`
				SpiderBlocksBroken                       float64 `json:"spider_blocks_broken"`
				SpiderAActivations                       float64 `json:"spider_a_activations"`
				SpiderActivationsStandard                float64 `json:"spider_activations_standard"`
				SpiderActivations                        float64 `json:"spider_activations"`
				SpiderAActivationsStandard               float64 `json:"spider_a_activations_standard"`
				PigmanArrowsHitFaceOff                   float64 `json:"pigman_arrows_hit_face_off"`
				PigmanDamageDealtFaceOff                 float64 `json:"pigman_damage_dealt_face_off"`
				PigmanFoodEatenFaceOff                   float64 `json:"pigman_food_eaten_face_off"`
				PigmanFinalAssistsFaceOff                float64 `json:"pigman_final_assists_face_off"`
				PigmanPotionsDrunkFaceOff                float64 `json:"pigman_potions_drunk_face_off"`
				PigmanIronOreBrokenFaceOff               float64 `json:"pigman_iron_ore_broken_face_off"`
				PigmanTreasuresFoundFaceOff              float64 `json:"pigman_treasures_found_face_off"`
				PigmanAAssists                           float64 `json:"pigman_a_assists"`
				PigmanActivations                        float64 `json:"pigman_activations"`
				PigmanBActivationsFaceOff                float64 `json:"pigman_b_activations_face_off"`
				PigmanBlocksPlacedPreparationFaceOff     float64 `json:"pigman_blocks_placed_preparation_face_off"`
				PigmanAKillsMeleeFaceOff                 float64 `json:"pigman_a_kills_melee_face_off"`
				KillsFaceOff                             float64 `json:"kills_face_off"`
				PigmanActivationsDeathmatch              float64 `json:"pigman_activations_deathmatch"`
				ArrowsHitFaceOff                         float64 `json:"arrows_hit_face_off"`
				PigmanGamesPlayed                        float64 `json:"pigman_games_played"`
				PigmanAssistsFaceOff                     float64 `json:"pigman_assists_face_off"`
				PigmanATotalKillsFaceOff                 float64 `json:"pigman_a_total_kills_face_off"`
				FoodEatenFaceOff                         float64 `json:"food_eaten_face_off"`
				PigmanAActivationsFaceOff                float64 `json:"pigman_a_activations_face_off"`
				PigmanATotalKills                        float64 `json:"pigman_a_total_kills"`
				BlocksBrokenFaceOff                      float64 `json:"blocks_broken_face_off"`
				PigmanIronArmorGiftedFaceOff             float64 `json:"pigman_iron_armor_gifted_face_off"`
				SteaksEatenFaceOff                       float64 `json:"steaks_eaten_face_off"`
				PigmanBlocksBrokenFaceOff                float64 `json:"pigman_blocks_broken_face_off"`
				PigmanDefenderKills                      float64 `json:"pigman_defender_kills"`
				PigmanADefenderAssistsFaceOff            float64 `json:"pigman_a_defender_assists_face_off"`
				PigmanKillsMeleeFaceOff                  float64 `json:"pigman_kills_melee_face_off"`
				PigmanKillsMeleeBehindFaceOff            float64 `json:"pigman_kills_melee_behind_face_off"`
				PigmanGamesPlayedFaceOff                 float64 `json:"pigman_games_played_face_off"`
				PigmanMetersWalkedSpeedFaceOff           float64 `json:"pigman_meters_walked_speed_face_off"`
				PigmanIronOreBroken                      float64 `json:"pigman_iron_ore_broken"`
				PigmanDefenderAssists                    float64 `json:"pigman_defender_assists"`
				PigmanADefenderKillsFaceOff              float64 `json:"pigman_a_defender_kills_face_off"`
				PigmanTotalKillsFaceOff                  float64 `json:"pigman_total_kills_face_off"`
				PigmanADamageDealt                       float64 `json:"pigman_a_damage_dealt"`
				PigmanWinsFaceOff                        float64 `json:"pigman_wins_face_off"`
				KillsMeleeFaceOff                        float64 `json:"kills_melee_face_off"`
				PigmanTimePlayed                         float64 `json:"pigman_time_played"`
				PigmanBActivationsDeathmatchFaceOff      float64 `json:"pigman_b_activations_deathmatch_face_off"`
				PigmanActivationsFaceOff                 float64 `json:"pigman_activations_face_off"`
				PigmanTotalDeaths                        float64 `json:"pigman_total_deaths"`
				TotalFinalKillsFaceOff                   float64 `json:"total_final_kills_face_off"`
				PigmanADefenderKills                     float64 `json:"pigman_a_defender_kills"`
				KillsMeleeBehindFaceOff                  float64 `json:"kills_melee_behind_face_off"`
				MetersWalkedSpeedFaceOff                 float64 `json:"meters_walked_speed_face_off"`
				PigmanDeathsFaceOff                      float64 `json:"pigman_deaths_face_off"`
				PigmanTreasuresFound                     float64 `json:"pigman_treasures_found"`
				PigmanPotionsDrunk                       float64 `json:"pigman_potions_drunk"`
				WoodChoppedFaceOff                       float64 `json:"wood_chopped_face_off"`
				PigmanBreadCraftedFaceOff                float64 `json:"pigman_bread_crafted_face_off"`
				PigmanSteaksEatenFaceOff                 float64 `json:"pigman_steaks_eaten_face_off"`
				PigmanWoodChopped                        float64 `json:"pigman_wood_chopped"`
				TreasuresFoundFaceOff                    float64 `json:"treasures_found_face_off"`
				PigmanAActivationsDeathmatchFaceOff      float64 `json:"pigman_a_activations_deathmatch_face_off"`
				ActivationsDeathmatchFaceOff             float64 `json:"activations_deathmatch_face_off"`
				PigmanAKillsFaceOff                      float64 `json:"pigman_a_kills_face_off"`
				ActivationsFaceOff                       float64 `json:"activations_face_off"`
				DefenderKillsFaceOff                     float64 `json:"defender_kills_face_off"`
				PigmanKillsMelee                         float64 `json:"pigman_kills_melee"`
				PigmanATotalFinalKills                   float64 `json:"pigman_a_total_final_kills"`
				AssistsFaceOff                           float64 `json:"assists_face_off"`
				PigmanDefenderAssistsFaceOff             float64 `json:"pigman_defender_assists_face_off"`
				DamageDealtFaceOff                       float64 `json:"damage_dealt_face_off"`
				PigmanFinalKillsMeleeBehind              float64 `json:"pigman_final_kills_melee_behind"`
				DefenderAssistsFaceOff                   float64 `json:"defender_assists_face_off"`
				PigmanFinalKillsMelee                    float64 `json:"pigman_final_kills_melee"`
				PigmanGActivationsFaceOff                float64 `json:"pigman_g_activations_face_off"`
				PigmanBActivationsDeathmatch             float64 `json:"pigman_b_activations_deathmatch"`
				PigmanMetersWalkedFaceOff                float64 `json:"pigman_meters_walked_face_off"`
				PigmanAKills                             float64 `json:"pigman_a_kills"`
				TimePlayedFaceOff                        float64 `json:"time_played_face_off"`
				PigmanFinalKillsFaceOff                  float64 `json:"pigman_final_kills_face_off"`
				PigmanTotalKills                         float64 `json:"pigman_total_kills"`
				PigmanADefenderAssists                   float64 `json:"pigman_a_defender_assists"`
				PigmanBlocksPlaced                       float64 `json:"pigman_blocks_placed"`
				PigmanATotalFinalKillsFaceOff            float64 `json:"pigman_a_total_final_kills_face_off"`
				PigmanArrowsFired                        float64 `json:"pigman_arrows_fired"`
				PigmanAKillsMelee                        float64 `json:"pigman_a_kills_melee"`
				PigmanAFinalAssistsFaceOff               float64 `json:"pigman_a_final_assists_face_off"`
				PigmanTimePlayedFaceOff                  float64 `json:"pigman_time_played_face_off"`
				TotalKillsFaceOff                        float64 `json:"total_kills_face_off"`
				PigmanKillsFaceOff                       float64 `json:"pigman_kills_face_off"`
				PigmanBActivations                       float64 `json:"pigman_b_activations"`
				PigmanAActivations                       float64 `json:"pigman_a_activations"`
				PigmanMetersWalked                       float64 `json:"pigman_meters_walked"`
				PigmanFinalAssists                       float64 `json:"pigman_final_assists"`
				PigmanAActivationsDeathmatch             float64 `json:"pigman_a_activations_deathmatch"`
				PigmanArrowsHit                          float64 `json:"pigman_arrows_hit"`
				PigmanDamageDealt                        float64 `json:"pigman_damage_dealt"`
				PigmanActivationsDeathmatchFaceOff       float64 `json:"pigman_activations_deathmatch_face_off"`
				PigmanFinalKillsMeleeBehindFaceOff       float64 `json:"pigman_final_kills_melee_behind_face_off"`
				PigmanKillsMeleeBehind                   float64 `json:"pigman_kills_melee_behind"`
				IronArmorGiftedFaceOff                   float64 `json:"iron_armor_gifted_face_off"`
				PigmanTotalFinalKillsFaceOff             float64 `json:"pigman_total_final_kills_face_off"`
				PigmanWoodChoppedFaceOff                 float64 `json:"pigman_wood_chopped_face_off"`
				PigmanGActivations                       float64 `json:"pigman_g_activations"`
				PigmanBlocksPlacedPreparation            float64 `json:"pigman_blocks_placed_preparation"`
				ArrowsFiredFaceOff                       float64 `json:"arrows_fired_face_off"`
				PigmanIronArmorGifted                    float64 `json:"pigman_iron_armor_gifted"`
				PigmanBlocksPlacedFaceOff                float64 `json:"pigman_blocks_placed_face_off"`
				PigmanAAssistsFaceOff                    float64 `json:"pigman_a_assists_face_off"`
				PigmanTotalDeathsFaceOff                 float64 `json:"pigman_total_deaths_face_off"`
				PigmanDefenderKillsFaceOff               float64 `json:"pigman_defender_kills_face_off"`
				PigmanFoodEaten                          float64 `json:"pigman_food_eaten"`
				PigmanAFinalAssists                      float64 `json:"pigman_a_final_assists"`
				FinalKillsMeleeFaceOff                   float64 `json:"final_kills_melee_face_off"`
				PigmanMetersFallenFaceOff                float64 `json:"pigman_meters_fallen_face_off"`
				PigmanADamageDealtFaceOff                float64 `json:"pigman_a_damage_dealt_face_off"`
				PigmanBreadEatenFaceOff                  float64 `json:"pigman_bread_eaten_face_off"`
				PigmanBreadCrafted                       float64 `json:"pigman_bread_crafted"`
				BreadEatenFaceOff                        float64 `json:"bread_eaten_face_off"`
				PigmanArrowsFiredFaceOff                 float64 `json:"pigman_arrows_fired_face_off"`
				PigmanFinalKillsMeleeFaceOff             float64 `json:"pigman_final_kills_melee_face_off"`
				IronOreBrokenFaceOff                     float64 `json:"iron_ore_broken_face_off"`
				PigmanAssists                            float64 `json:"pigman_assists"`
				BreadCraftedFaceOff                      float64 `json:"bread_crafted_face_off"`
				PigmanMetersFallen                       float64 `json:"pigman_meters_fallen"`
				FinalKillsMeleeBehindFaceOff             float64 `json:"final_kills_melee_behind_face_off"`
				PigmanBlocksBroken                       float64 `json:"pigman_blocks_broken"`
				PigmanMetersWalkedSpeed                  float64 `json:"pigman_meters_walked_speed"`
				TotalDeathsFaceOff                       float64 `json:"total_deaths_face_off"`
				FinalKillsFaceOff                        float64 `json:"final_kills_face_off"`
				PotionsDrunkFaceOff                      float64 `json:"potions_drunk_face_off"`
				FinalAssistsFaceOff                      float64 `json:"final_assists_face_off"`
				PigmanBreadEaten                         float64 `json:"pigman_bread_eaten"`
				DeathsFaceOff                            float64 `json:"deaths_face_off"`
				PigmanSteaksEaten                        float64 `json:"pigman_steaks_eaten"`
				SkeletonKillsMeleeBehind                 float64 `json:"skeleton_kills_melee_behind"`
				SkeletonGamesPlayedFaceOff               float64 `json:"skeleton_games_played_face_off"`
				SkeletonAssists                          float64 `json:"skeleton_assists"`
				SkeletonTreasuresFound                   float64 `json:"skeleton_treasures_found"`
				SkeletonDefenderAssists                  float64 `json:"skeleton_defender_assists"`
				SkeletonBlocksPlacedPreparation          float64 `json:"skeleton_blocks_placed_preparation"`
				SkeletonBlocksPlacedFaceOff              float64 `json:"skeleton_blocks_placed_face_off"`
				SkeletonMetersWalkedSpeed                float64 `json:"skeleton_meters_walked_speed"`
				SkeletonTotalKills                       float64 `json:"skeleton_total_kills"`
				SkeletonLossesFaceOff                    float64 `json:"skeleton_losses_face_off"`
				SkeletonDefenderKills                    float64 `json:"skeleton_defender_kills"`
				SkeletonBlocksPlaced                     float64 `json:"skeleton_blocks_placed"`
				SkeletonBlocksBrokenFaceOff              float64 `json:"skeleton_blocks_broken_face_off"`
				SkeletonActivationsFaceOff               float64 `json:"skeleton_activations_face_off"`
				SkeletonIronOreBrokenFaceOff             float64 `json:"skeleton_iron_ore_broken_face_off"`
				SkeletonTreasuresFoundFaceOff            float64 `json:"skeleton_treasures_found_face_off"`
				SkeletonKillsFaceOff                     float64 `json:"skeleton_kills_face_off"`
				SkeletonBlocksPlacedPreparationFaceOff   float64 `json:"skeleton_blocks_placed_preparation_face_off"`
				SkeletonPotionsDrunk                     float64 `json:"skeleton_potions_drunk"`
				SkeletonMetersWalkedSpeedFaceOff         float64 `json:"skeleton_meters_walked_speed_face_off"`
				SkeletonMetersWalked                     float64 `json:"skeleton_meters_walked"`
				SkeletonTimePlayed                       float64 `json:"skeleton_time_played"`
				SkeletonMetersFallen                     float64 `json:"skeleton_meters_fallen"`
				SkeletonKillsMeleeFaceOff                float64 `json:"skeleton_kills_melee_face_off"`
				SkeletonMetersWalkedFaceOff              float64 `json:"skeleton_meters_walked_face_off"`
				SkeletonAActivations                     float64 `json:"skeleton_a_activations"`
				SkeletonPotionsDrunkFaceOff              float64 `json:"skeleton_potions_drunk_face_off"`
				SkeletonArrowsFired                      float64 `json:"skeleton_arrows_fired"`
				SkeletonDefenderAssistsFaceOff           float64 `json:"skeleton_defender_assists_face_off"`
				SkeletonAActivationsFaceOff              float64 `json:"skeleton_a_activations_face_off"`
				SkeletonArrowsHitFaceOff                 float64 `json:"skeleton_arrows_hit_face_off"`
				SkeletonBlocksBroken                     float64 `json:"skeleton_blocks_broken"`
				SkeletonAssistsFaceOff                   float64 `json:"skeleton_assists_face_off"`
				SkeletonArrowsHit                        float64 `json:"skeleton_arrows_hit"`
				SkeletonIronOreBroken                    float64 `json:"skeleton_iron_ore_broken"`
				SkeletonActivations                      float64 `json:"skeleton_activations"`
				SkeletonKillsMelee                       float64 `json:"skeleton_kills_melee"`
				SkeletonGamesPlayed                      float64 `json:"skeleton_games_played"`
				SkeletonKillsMeleeBehindFaceOff          float64 `json:"skeleton_kills_melee_behind_face_off"`
				SkeletonDefenderKillsFaceOff             float64 `json:"skeleton_defender_kills_face_off"`
				SkeletonMetersFallenFaceOff              float64 `json:"skeleton_meters_fallen_face_off"`
				SkeletonTotalKillsFaceOff                float64 `json:"skeleton_total_kills_face_off"`
				SkeletonTimePlayedFaceOff                float64 `json:"skeleton_time_played_face_off"`
				SkeletonArrowsFiredFaceOff               float64 `json:"skeleton_arrows_fired_face_off"`
			} `json:"Walls3"`
			GingerBread struct {
				EngineActive                                  string   `json:"engine_active"`
				FrameActive                                   string   `json:"frame_active"`
				BoosterActive                                 string   `json:"booster_active"`
				HelmetActive                                  string   `json:"helmet_active"`
				JacketActive                                  string   `json:"jacket_active"`
				PantsActive                                   string   `json:"pants_active"`
				SkinActive                                    string   `json:"skin_active"`
				ShoesActive                                   string   `json:"shoes_active"`
				Packages                                      []string `json:"packages"`
				Horn                                          string   `json:"horn"`
				Coins                                         float64  `json:"coins"`
				GoldTrophy                                    float64  `json:"gold_trophy"`
				LapsCompleted                                 float64  `json:"laps_completed"`
				CoinsPickedUp                                 float64  `json:"coins_picked_up"`
				BlueTorpedoHit                                float64  `json:"blue_torpedo_hit"`
				BoxPickups                                    float64  `json:"box_pickups"`
				BananaHitsSent                                float64  `json:"banana_hits_sent"`
				Wins                                          float64  `json:"wins"`
				Parts                                         string   `json:"parts"`
				BananaHitsReceived                            float64  `json:"banana_hits_received"`
				SilverTrophy                                  float64  `json:"silver_trophy"`
				BronzeTrophy                                  float64  `json:"bronze_trophy"`
				ParticleTrail                                 string   `json:"particle_trail"`
				GoldTrophyOlympus                             float64  `json:"gold_trophy_olympus"`
				GoldTrophyJunglerush                          float64  `json:"gold_trophy_junglerush"`
				GoldTrophyHypixelgp                           float64  `json:"gold_trophy_hypixelgp"`
				GoldTrophyRetro                               float64  `json:"gold_trophy_retro"`
				GoldTrophyCanyon                              float64  `json:"gold_trophy_canyon"`
				SilverTrophyWeeklyB                           float64  `json:"silver_trophy_weekly_b"`
				SilverTrophyMonthlyB                          float64  `json:"silver_trophy_monthly_b"`
				SilverTrophyOlympus                           float64  `json:"silver_trophy_olympus"`
				SilverTrophyJunglerush                        float64  `json:"silver_trophy_junglerush"`
				SilverTrophyRetro                             float64  `json:"silver_trophy_retro"`
				SilverTrophyCanyon                            float64  `json:"silver_trophy_canyon"`
				BronzeTrophyWeeklyB                           float64  `json:"bronze_trophy_weekly_b"`
				BronzeTrophyMonthlyB                          float64  `json:"bronze_trophy_monthly_b"`
				BronzeTrophyOlympus                           float64  `json:"bronze_trophy_olympus"`
				BronzeTrophyJunglerush                        float64  `json:"bronze_trophy_junglerush"`
				BronzeTrophyHypixelgp                         float64  `json:"bronze_trophy_hypixelgp"`
				BronzeTrophyRetro                             float64  `json:"bronze_trophy_retro"`
				BronzeTrophyCanyon                            float64  `json:"bronze_trophy_canyon"`
				BoxPickupsWeeklyB                             float64  `json:"box_pickups_weekly_b"`
				BoxPickupsMonthlyB                            float64  `json:"box_pickups_monthly_b"`
				BoxPickupsOlympus                             float64  `json:"box_pickups_olympus"`
				BoxPickupsJunglerush                          float64  `json:"box_pickups_junglerush"`
				BoxPickupsHypixelgp                           float64  `json:"box_pickups_hypixelgp"`
				BoxPickupsRetro                               float64  `json:"box_pickups_retro"`
				BoxPickupsCanyon                              float64  `json:"box_pickups_canyon"`
				OlympusPlays                                  float64  `json:"olympus_plays"`
				JunglerushPlays                               float64  `json:"junglerush_plays"`
				HypixelgpPlays                                float64  `json:"hypixelgp_plays"`
				RetroPlays                                    float64  `json:"retro_plays"`
				CanyonPlays                                   float64  `json:"canyon_plays"`
				BoxPickupsWeeklyA                             float64  `json:"box_pickups_weekly_a"`
				BoxPickupsMonthlyA                            float64  `json:"box_pickups_monthly_a"`
				GoldTrophyMonthlyB                            float64  `json:"gold_trophy_monthly_b"`
				GoldTrophyWeeklyB                             float64  `json:"gold_trophy_weekly_b"`
				GoldTrophyMonthlyA                            float64  `json:"gold_trophy_monthly_a"`
				SilverTrophyWeeklyA                           float64  `json:"silver_trophy_weekly_a"`
				SilverTrophyMonthlyA                          float64  `json:"silver_trophy_monthly_a"`
				GoldTrophyWeeklyA                             float64  `json:"gold_trophy_weekly_a"`
				BronzeTrophyMonthlyA                          float64  `json:"bronze_trophy_monthly_a"`
				BronzeTrophyWeeklyA                           float64  `json:"bronze_trophy_weekly_a"`
				LastTourneyAd                                 float64  `json:"lastTourneyAd"`
				TourneyGingerbreadSolo0BananaHitsSent         float64  `json:"tourney_gingerbread_solo_0_banana_hits_sent"`
				TourneyGingerbreadSolo0BlueTorpedoHit         float64  `json:"tourney_gingerbread_solo_0_blue_torpedo_hit"`
				TourneyGingerbreadSolo0BoxPickups             float64  `json:"tourney_gingerbread_solo_0_box_pickups"`
				TourneyGingerbreadSolo0BoxPickupsJunglerush   float64  `json:"tourney_gingerbread_solo_0_box_pickups_junglerush"`
				TourneyGingerbreadSolo0BoxPickupsMonthlyA     float64  `json:"tourney_gingerbread_solo_0_box_pickups_monthly_a"`
				TourneyGingerbreadSolo0BoxPickupsWeeklyA      float64  `json:"tourney_gingerbread_solo_0_box_pickups_weekly_a"`
				TourneyGingerbreadSolo0CoinsPickedUp          float64  `json:"tourney_gingerbread_solo_0_coins_picked_up"`
				TourneyGingerbreadSolo0JunglerushPlays        float64  `json:"tourney_gingerbread_solo_0_junglerush_plays"`
				TourneyGingerbreadSolo0LapsCompleted          float64  `json:"tourney_gingerbread_solo_0_laps_completed"`
				TourneyGingerbreadSolo0SilverTrophy           float64  `json:"tourney_gingerbread_solo_0_silver_trophy"`
				TourneyGingerbreadSolo0SilverTrophyJunglerush float64  `json:"tourney_gingerbread_solo_0_silver_trophy_junglerush"`
				TourneyGingerbreadSolo0SilverTrophyMonthlyA   float64  `json:"tourney_gingerbread_solo_0_silver_trophy_monthly_a"`
				TourneyGingerbreadSolo0SilverTrophyWeeklyA    float64  `json:"tourney_gingerbread_solo_0_silver_trophy_weekly_a"`
				TourneyGingerbreadSolo0Wins                   float64  `json:"tourney_gingerbread_solo_0_wins"`
				TourneyGingerbreadSolo0BoxPickupsOlympus      float64  `json:"tourney_gingerbread_solo_0_box_pickups_olympus"`
				TourneyGingerbreadSolo0GoldTrophy             float64  `json:"tourney_gingerbread_solo_0_gold_trophy"`
				TourneyGingerbreadSolo0GoldTrophyMonthlyA     float64  `json:"tourney_gingerbread_solo_0_gold_trophy_monthly_a"`
				TourneyGingerbreadSolo0GoldTrophyOlympus      float64  `json:"tourney_gingerbread_solo_0_gold_trophy_olympus"`
				TourneyGingerbreadSolo0GoldTrophyWeeklyA      float64  `json:"tourney_gingerbread_solo_0_gold_trophy_weekly_a"`
				TourneyGingerbreadSolo0OlympusPlays           float64  `json:"tourney_gingerbread_solo_0_olympus_plays"`
				TourneyGingerbreadSolo0BoxPickupsWeeklyB      float64  `json:"tourney_gingerbread_solo_0_box_pickups_weekly_b"`
				TourneyGingerbreadSolo0GoldTrophyWeeklyB      float64  `json:"tourney_gingerbread_solo_0_gold_trophy_weekly_b"`
				TourneyGingerbreadSolo0BananaHitsReceived     float64  `json:"tourney_gingerbread_solo_0_banana_hits_received"`
				TourneyGingerbreadSolo0BoxPickupsRetro        float64  `json:"tourney_gingerbread_solo_0_box_pickups_retro"`
				TourneyGingerbreadSolo0RetroPlays             float64  `json:"tourney_gingerbread_solo_0_retro_plays"`
				TourneyGingerbreadSolo0SilverTrophyOlympus    float64  `json:"tourney_gingerbread_solo_0_silver_trophy_olympus"`
				TourneyGingerbreadSolo0SilverTrophyWeeklyB    float64  `json:"tourney_gingerbread_solo_0_silver_trophy_weekly_b"`
				TourneyGingerbreadSolo0BoxPickupsHypixelgp    float64  `json:"tourney_gingerbread_solo_0_box_pickups_hypixelgp"`
				TourneyGingerbreadSolo0HypixelgpPlays         float64  `json:"tourney_gingerbread_solo_0_hypixelgp_plays"`
				TourneyGingerbreadSolo0SilverTrophyHypixelgp  float64  `json:"tourney_gingerbread_solo_0_silver_trophy_hypixelgp"`
				TourneyGingerbreadSolo0BoxPickupsCanyon       float64  `json:"tourney_gingerbread_solo_0_box_pickups_canyon"`
				TourneyGingerbreadSolo0BronzeTrophy           float64  `json:"tourney_gingerbread_solo_0_bronze_trophy"`
				TourneyGingerbreadSolo0BronzeTrophyCanyon     float64  `json:"tourney_gingerbread_solo_0_bronze_trophy_canyon"`
				TourneyGingerbreadSolo0BronzeTrophyMonthlyA   float64  `json:"tourney_gingerbread_solo_0_bronze_trophy_monthly_a"`
				TourneyGingerbreadSolo0BronzeTrophyWeeklyB    float64  `json:"tourney_gingerbread_solo_0_bronze_trophy_weekly_b"`
				TourneyGingerbreadSolo0CanyonPlays            float64  `json:"tourney_gingerbread_solo_0_canyon_plays"`
				TourneyGingerbreadSolo0SilverTrophyRetro      float64  `json:"tourney_gingerbread_solo_0_silver_trophy_retro"`
				TourneyGingerbreadSolo0SilverTrophyCanyon     float64  `json:"tourney_gingerbread_solo_0_silver_trophy_canyon"`
				TourneyGingerbreadSolo0BronzeTrophyJunglerush float64  `json:"tourney_gingerbread_solo_0_bronze_trophy_junglerush"`
				TourneyGingerbreadSolo0GoldTrophyJunglerush   float64  `json:"tourney_gingerbread_solo_0_gold_trophy_junglerush"`
				TourneyGingerbreadSolo0BronzeTrophyHypixelgp  float64  `json:"tourney_gingerbread_solo_0_bronze_trophy_hypixelgp"`
				TourneyGingerbreadSolo0GoldTrophyCanyon       float64  `json:"tourney_gingerbread_solo_0_gold_trophy_canyon"`
				TourneyGingerbreadSolo0BronzeTrophyOlympus    float64  `json:"tourney_gingerbread_solo_0_bronze_trophy_olympus"`
				TourneyGingerbreadSolo0GoldTrophyRetro        float64  `json:"tourney_gingerbread_solo_0_gold_trophy_retro"`
				TourneyGingerbreadSolo0GoldTrophyHypixelgp    float64  `json:"tourney_gingerbread_solo_0_gold_trophy_hypixelgp"`
			} `json:"GingerBread"`
			SkyWars struct {
				WinStreak                                     float64  `json:"win_streak"`
				SurvivedPlayers                               float64  `json:"survived_players"`
				Losses                                        float64  `json:"losses"`
				LossesSolo                                    float64  `json:"losses_solo"`
				SurvivedPlayersKitBasicSoloDefault            float64  `json:"survived_players_kit_basic_solo_default"`
				BlocksBroken                                  float64  `json:"blocks_broken"`
				DeathsKitBasicSoloDefault                     float64  `json:"deaths_kit_basic_solo_default"`
				LossesSoloNormal                              float64  `json:"losses_solo_normal"`
				DeathsSolo                                    float64  `json:"deaths_solo"`
				Deaths                                        float64  `json:"deaths"`
				Quits                                         float64  `json:"quits"`
				DeathsSoloNormal                              float64  `json:"deaths_solo_normal"`
				SurvivedPlayersSolo                           float64  `json:"survived_players_solo"`
				LossesKitBasicSoloDefault                     float64  `json:"losses_kit_basic_solo_default"`
				Coins                                         float64  `json:"coins"`
				Souls                                         float64  `json:"souls"`
				ArrowsHit                                     float64  `json:"arrows_hit"`
				BlocksPlaced                                  float64  `json:"blocks_placed"`
				KillsKitBasicSoloDefault                      float64  `json:"kills_kit_basic_solo_default"`
				Games                                         float64  `json:"games"`
				Kills                                         float64  `json:"kills"`
				ArrowsShot                                    float64  `json:"arrows_shot"`
				WinsSoloInsane                                float64  `json:"wins_solo_insane"`
				GamesSolo                                     float64  `json:"games_solo"`
				KillsSoloInsane                               float64  `json:"kills_solo_insane"`
				WinsSolo                                      float64  `json:"wins_solo"`
				Wins                                          float64  `json:"wins"`
				WinsKitBasicSoloDefault                       float64  `json:"wins_kit_basic_solo_default"`
				SoulsGathered                                 float64  `json:"souls_gathered"`
				ItemsEnchanted                                float64  `json:"items_enchanted"`
				KillsSolo                                     float64  `json:"kills_solo"`
				GamesKitBasicSoloDefault                      float64  `json:"games_kit_basic_solo_default"`
				SoulWell                                      float64  `json:"soul_well"`
				UsedSoulWell                                  bool     `json:"usedSoulWell"`
				Packages                                      []string `json:"packages"`
				DeathsSoloInsane                              float64  `json:"deaths_solo_insane"`
				LossesSoloInsane                              float64  `json:"losses_solo_insane"`
				XezbethLuck                                   float64  `json:"xezbeth_luck"`
				AssistsSolo                                   float64  `json:"assists_solo"`
				AssistsKitBasicSoloDefault                    float64  `json:"assists_kit_basic_solo_default"`
				EnderpearlsThrown                             float64  `json:"enderpearls_thrown"`
				Assists                                       float64  `json:"assists"`
				EggThrown                                     float64  `json:"egg_thrown"`
				ActiveKitSOLO                                 string   `json:"activeKit_SOLO"`
				LossesKitBasicSoloEcologist                   float64  `json:"losses_kit_basic_solo_ecologist"`
				DeathsKitBasicSoloEcologist                   float64  `json:"deaths_kit_basic_solo_ecologist"`
				KillsKitBasicSoloEcologist                    float64  `json:"kills_kit_basic_solo_ecologist"`
				SurvivedPlayersKitBasicSoloEcologist          float64  `json:"survived_players_kit_basic_solo_ecologist"`
				GamesKitBasicSoloEcologist                    float64  `json:"games_kit_basic_solo_ecologist"`
				WinsKitBasicSoloEcologist                     float64  `json:"wins_kit_basic_solo_ecologist"`
				WinsSoloNormal                                float64  `json:"wins_solo_normal"`
				KillsSoloNormal                               float64  `json:"kills_solo_normal"`
				KillsKitMiningTeamDefault                     float64  `json:"kills_kit_mining_team_default"`
				SurvivedPlayersTeam                           float64  `json:"survived_players_team"`
				LossesKitMiningTeamDefault                    float64  `json:"losses_kit_mining_team_default"`
				LossesTeamInsane                              float64  `json:"losses_team_insane"`
				DeathsKitMiningTeamDefault                    float64  `json:"deaths_kit_mining_team_default"`
				DeathsTeam                                    float64  `json:"deaths_team"`
				KillsTeam                                     float64  `json:"kills_team"`
				DeathsTeamInsane                              float64  `json:"deaths_team_insane"`
				KillsTeamInsane                               float64  `json:"kills_team_insane"`
				SurvivedPlayersKitMiningTeamDefault           float64  `json:"survived_players_kit_mining_team_default"`
				LossesTeam                                    float64  `json:"losses_team"`
				TeamResistanceBoost                           float64  `json:"team_resistance_boost"`
				SoulWellLegendaries                           float64  `json:"soul_well_legendaries"`
				RefillChestDestroy                            float64  `json:"refill_chest_destroy"`
				PaidSouls                                     float64  `json:"paid_souls"`
				SoulWellRares                                 float64  `json:"soul_well_rares"`
				SoloResistanceBoost                           float64  `json:"solo_resistance_boost"`
				ActiveKitTEAM                                 string   `json:"activeKit_TEAM"`
				LossesKitSupportingTeamHealer                 float64  `json:"losses_kit_supporting_team_healer"`
				SurvivedPlayersKitSupportingTeamHealer        float64  `json:"survived_players_kit_supporting_team_healer"`
				DeathsKitSupportingTeamHealer                 float64  `json:"deaths_kit_supporting_team_healer"`
				WinsTeamInsane                                float64  `json:"wins_team_insane"`
				GamesKitSupportingTeamHealer                  float64  `json:"games_kit_supporting_team_healer"`
				KillsKitSupportingTeamHealer                  float64  `json:"kills_kit_supporting_team_healer"`
				WinsKitSupportingTeamHealer                   float64  `json:"wins_kit_supporting_team_healer"`
				WinsTeam                                      float64  `json:"wins_team"`
				GamesTeam                                     float64  `json:"games_team"`
				VotesTemple                                   float64  `json:"votes_Temple"`
				AssistsKitSupportingTeamHealer                float64  `json:"assists_kit_supporting_team_healer"`
				AssistsTeam                                   float64  `json:"assists_team"`
				AssistsKitBasicSoloEcologist                  float64  `json:"assists_kit_basic_solo_ecologist"`
				VotesShattered                                float64  `json:"votes_Shattered"`
				VotesShire                                    float64  `json:"votes_Shire"`
				TeamEnderMastery                              float64  `json:"team_ender_mastery"`
				TeamMiningExpertise                           float64  `json:"team_mining_expertise"`
				SoloMiningExpertise                           float64  `json:"solo_mining_expertise"`
				VotesDwarven                                  float64  `json:"votes_Dwarven"`
				TeamInstantSmelting                           float64  `json:"team_instant_smelting"`
				GamesKitAttackingTeamHunter                   float64  `json:"games_kit_attacking_team_hunter"`
				KillsKitAttackingTeamHunter                   float64  `json:"kills_kit_attacking_team_hunter"`
				SurvivedPlayersKitAttackingTeamHunter         float64  `json:"survived_players_kit_attacking_team_hunter"`
				WinsKitAttackingTeamHunter                    float64  `json:"wins_kit_attacking_team_hunter"`
				DeathsKitAttackingTeamHunter                  float64  `json:"deaths_kit_attacking_team_hunter"`
				LossesKitAttackingTeamHunter                  float64  `json:"losses_kit_attacking_team_hunter"`
				VotesAtuin                                    float64  `json:"votes_Atuin"`
				SurvivedPlayersKitAttackingTeamKnight         float64  `json:"survived_players_kit_attacking_team_knight"`
				GamesKitAttackingTeamKnight                   float64  `json:"games_kit_attacking_team_knight"`
				WinsKitAttackingTeamKnight                    float64  `json:"wins_kit_attacking_team_knight"`
				KillsKitAttackingTeamKnight                   float64  `json:"kills_kit_attacking_team_knight"`
				KillsTeamNormal                               float64  `json:"kills_team_normal"`
				WinsTeamNormal                                float64  `json:"wins_team_normal"`
				DeathsTeamNormal                              float64  `json:"deaths_team_normal"`
				LossesTeamNormal                              float64  `json:"losses_team_normal"`
				LossesKitAttackingTeamKnight                  float64  `json:"losses_kit_attacking_team_knight"`
				DeathsKitAttackingTeamKnight                  float64  `json:"deaths_kit_attacking_team_knight"`
				AssistsKitAttackingTeamKnight                 float64  `json:"assists_kit_attacking_team_knight"`
				SoloBlazingArrows                             float64  `json:"solo_blazing_arrows"`
				VotesCongo                                    float64  `json:"votes_Congo"`
				DeathsKitAttackingTeamScout                   float64  `json:"deaths_kit_attacking_team_scout"`
				WinsKitAttackingTeamScout                     float64  `json:"wins_kit_attacking_team_scout"`
				SurvivedPlayersKitAttackingTeamScout          float64  `json:"survived_players_kit_attacking_team_scout"`
				KillsKitAttackingTeamScout                    float64  `json:"kills_kit_attacking_team_scout"`
				GamesKitAttackingTeamScout                    float64  `json:"games_kit_attacking_team_scout"`
				AssistsKitAttackingTeamScout                  float64  `json:"assists_kit_attacking_team_scout"`
				LossesKitAttackingTeamScout                   float64  `json:"losses_kit_attacking_team_scout"`
				VotesLongIsland                               float64  `json:"votes_LongIsland"`
				SurvivedPlayersKitSupportingTeamArmorsmith    float64  `json:"survived_players_kit_supporting_team_armorsmith"`
				GamesKitSupportingTeamArmorsmith              float64  `json:"games_kit_supporting_team_armorsmith"`
				WinsKitSupportingTeamArmorsmith               float64  `json:"wins_kit_supporting_team_armorsmith"`
				KillsKitSupportingTeamArmorsmith              float64  `json:"kills_kit_supporting_team_armorsmith"`
				VotesToadstool                                float64  `json:"votes_Toadstool"`
				DeathsKitSupportingTeamArmorsmith             float64  `json:"deaths_kit_supporting_team_armorsmith"`
				LossesKitSupportingTeamArmorsmith             float64  `json:"losses_kit_supporting_team_armorsmith"`
				AssistsKitSupportingTeamArmorsmith            float64  `json:"assists_kit_supporting_team_armorsmith"`
				VotesFrostbite                                float64  `json:"votes_Frostbite"`
				WinsKitMiningTeamSpeleologist                 float64  `json:"wins_kit_mining_team_speleologist"`
				GamesKitMiningTeamSpeleologist                float64  `json:"games_kit_mining_team_speleologist"`
				KillsKitMiningTeamSpeleologist                float64  `json:"kills_kit_mining_team_speleologist"`
				AssistsKitMiningTeamSpeleologist              float64  `json:"assists_kit_mining_team_speleologist"`
				SurvivedPlayersKitMiningTeamSpeleologist      float64  `json:"survived_players_kit_mining_team_speleologist"`
				DeathsKitMiningTeamSpeleologist               float64  `json:"deaths_kit_mining_team_speleologist"`
				LossesKitMiningTeamSpeleologist               float64  `json:"losses_kit_mining_team_speleologist"`
				VotesDragonice                                float64  `json:"votes_Dragonice"`
				SoloInstantSmelting                           float64  `json:"solo_instant_smelting"`
				GamesKitDefendingTeamArmorer                  float64  `json:"games_kit_defending_team_armorer"`
				SurvivedPlayersKitDefendingTeamArmorer        float64  `json:"survived_players_kit_defending_team_armorer"`
				WinsKitDefendingTeamArmorer                   float64  `json:"wins_kit_defending_team_armorer"`
				KillsKitDefendingTeamArmorer                  float64  `json:"kills_kit_defending_team_armorer"`
				AssistsKitDefendingTeamArmorer                float64  `json:"assists_kit_defending_team_armorer"`
				VotesSiege                                    float64  `json:"votes_Siege"`
				WinsKitBasicSoloArmorsmith                    float64  `json:"wins_kit_basic_solo_armorsmith"`
				SurvivedPlayersKitBasicSoloArmorsmith         float64  `json:"survived_players_kit_basic_solo_armorsmith"`
				KillsKitBasicSoloArmorsmith                   float64  `json:"kills_kit_basic_solo_armorsmith"`
				GamesKitBasicSoloArmorsmith                   float64  `json:"games_kit_basic_solo_armorsmith"`
				VotesElven                                    float64  `json:"votes_Elven"`
				VotesOnionring                                float64  `json:"votes_Onionring"`
				VotesOverfall                                 float64  `json:"votes_Overfall"`
				TeamArrowRecovery                             float64  `json:"team_arrow_recovery"`
				SoloMarksmanship                              float64  `json:"solo_marksmanship"`
				SoloEnderMastery                              float64  `json:"solo_ender_mastery"`
				TeamSpeedBoost                                float64  `json:"team_speed_boost"`
				TeamBlazingArrows                             float64  `json:"team_blazing_arrows"`
				DeathsKitAdvancedSoloCannoneer                float64  `json:"deaths_kit_advanced_solo_cannoneer"`
				LossesKitAdvancedSoloCannoneer                float64  `json:"losses_kit_advanced_solo_cannoneer"`
				SurvivedPlayersKitAdvancedSoloCannoneer       float64  `json:"survived_players_kit_advanced_solo_cannoneer"`
				KillsKitAdvancedSoloCannoneer                 float64  `json:"kills_kit_advanced_solo_cannoneer"`
				WinsKitAdvancedSoloCannoneer                  float64  `json:"wins_kit_advanced_solo_cannoneer"`
				GamesKitAdvancedSoloCannoneer                 float64  `json:"games_kit_advanced_solo_cannoneer"`
				LossesKitBasicSoloArmorsmith                  float64  `json:"losses_kit_basic_solo_armorsmith"`
				DeathsKitBasicSoloArmorsmith                  float64  `json:"deaths_kit_basic_solo_armorsmith"`
				SurvivedPlayersKitAdvancedSoloPyro            float64  `json:"survived_players_kit_advanced_solo_pyro"`
				DeathsKitAdvancedSoloPyro                     float64  `json:"deaths_kit_advanced_solo_pyro"`
				LossesKitAdvancedSoloPyro                     float64  `json:"losses_kit_advanced_solo_pyro"`
				GamesKitAdvancedSoloPyro                      float64  `json:"games_kit_advanced_solo_pyro"`
				WinsKitAdvancedSoloPyro                       float64  `json:"wins_kit_advanced_solo_pyro"`
				KillsKitAdvancedSoloPyro                      float64  `json:"kills_kit_advanced_solo_pyro"`
				AssistsKitAdvancedSoloPyro                    float64  `json:"assists_kit_advanced_solo_pyro"`
				DeathsKitAttackingTeamSnowman                 float64  `json:"deaths_kit_attacking_team_snowman"`
				SurvivedPlayersKitAttackingTeamSnowman        float64  `json:"survived_players_kit_attacking_team_snowman"`
				LossesKitAttackingTeamSnowman                 float64  `json:"losses_kit_attacking_team_snowman"`
				GamesKitAttackingTeamSnowman                  float64  `json:"games_kit_attacking_team_snowman"`
				KillsKitAttackingTeamSnowman                  float64  `json:"kills_kit_attacking_team_snowman"`
				WinsKitDefendingTeamGuardian                  float64  `json:"wins_kit_defending_team_guardian"`
				SurvivedPlayersKitDefendingTeamGuardian       float64  `json:"survived_players_kit_defending_team_guardian"`
				GamesKitDefendingTeamGuardian                 float64  `json:"games_kit_defending_team_guardian"`
				KillsKitDefendingTeamGuardian                 float64  `json:"kills_kit_defending_team_guardian"`
				LossesKitDefendingTeamGuardian                float64  `json:"losses_kit_defending_team_guardian"`
				DeathsKitDefendingTeamGuardian                float64  `json:"deaths_kit_defending_team_guardian"`
				AssistsKitDefendingTeamGuardian               float64  `json:"assists_kit_defending_team_guardian"`
				AssistsKitAttackingTeamSnowman                float64  `json:"assists_kit_attacking_team_snowman"`
				WinsKitAttackingTeamSnowman                   float64  `json:"wins_kit_attacking_team_snowman"`
				AssistsKitAdvancedSoloCannoneer               float64  `json:"assists_kit_advanced_solo_cannoneer"`
				ActiveKitMEGA                                 string   `json:"activeKit_MEGA"`
				DeathsKitMegaMegaDefault                      float64  `json:"deaths_kit_mega_mega_default"`
				SurvivedPlayersMega                           float64  `json:"survived_players_mega"`
				DeathsMegaNormal                              float64  `json:"deaths_mega_normal"`
				LossesMegaNormal                              float64  `json:"losses_mega_normal"`
				DeathsMega                                    float64  `json:"deaths_mega"`
				SurvivedPlayersKitMegaMegaDefault             float64  `json:"survived_players_kit_mega_mega_default"`
				LossesMega                                    float64  `json:"losses_mega"`
				LossesKitMegaMegaDefault                      float64  `json:"losses_kit_mega_mega_default"`
				KillsMega                                     float64  `json:"kills_mega"`
				KillsKitMegaMegaDefault                       float64  `json:"kills_kit_mega_mega_default"`
				WinsMega                                      float64  `json:"wins_mega"`
				GamesKitMegaMegaDefault                       float64  `json:"games_kit_mega_mega_default"`
				GamesMega                                     float64  `json:"games_mega"`
				WinsKitMegaMegaDefault                        float64  `json:"wins_kit_mega_mega_default"`
				KillsMegaNormal                               float64  `json:"kills_mega_normal"`
				WinsMegaNormal                                float64  `json:"wins_mega_normal"`
				AssistsKitMegaMegaDefault                     float64  `json:"assists_kit_mega_mega_default"`
				AssistsMega                                   float64  `json:"assists_mega"`
				SurvivedPlayersKitAdvancedSoloArmorer         float64  `json:"survived_players_kit_advanced_solo_armorer"`
				GamesKitAdvancedSoloArmorer                   float64  `json:"games_kit_advanced_solo_armorer"`
				WinsKitAdvancedSoloArmorer                    float64  `json:"wins_kit_advanced_solo_armorer"`
				KillsKitAdvancedSoloArmorer                   float64  `json:"kills_kit_advanced_solo_armorer"`
				LossesKitDefendingTeamArmorer                 float64  `json:"losses_kit_defending_team_armorer"`
				DeathsKitDefendingTeamArmorer                 float64  `json:"deaths_kit_defending_team_armorer"`
				KitMegaMegaPaladin                            float64  `json:"kit_mega_mega_paladin"`
				KitMegaMegaArmorer                            float64  `json:"kit_mega_mega_armorer"`
				KillsKitMegaMegaPaladin                       float64  `json:"kills_kit_mega_mega_paladin"`
				SurvivedPlayersKitMegaMegaPaladin             float64  `json:"survived_players_kit_mega_mega_paladin"`
				WinsKitMegaMegaPaladin                        float64  `json:"wins_kit_mega_mega_paladin"`
				GamesKitMegaMegaPaladin                       float64  `json:"games_kit_mega_mega_paladin"`
				DeathsKitMegaMegaPaladin                      float64  `json:"deaths_kit_mega_mega_paladin"`
				LossesKitMegaMegaPaladin                      float64  `json:"losses_kit_mega_mega_paladin"`
				MegaRusher                                    float64  `json:"mega_rusher"`
				AssistsKitMegaMegaPaladin                     float64  `json:"assists_kit_mega_mega_paladin"`
				MegaEnderMastery                              float64  `json:"mega_ender_mastery"`
				LossesKitMegaMegaArmorer                      float64  `json:"losses_kit_mega_mega_armorer"`
				AssistsKitMegaMegaArmorer                     float64  `json:"assists_kit_mega_mega_armorer"`
				SurvivedPlayersKitMegaMegaArmorer             float64  `json:"survived_players_kit_mega_mega_armorer"`
				DeathsKitMegaMegaArmorer                      float64  `json:"deaths_kit_mega_mega_armorer"`
				KillsKitMegaMegaArmorer                       float64  `json:"kills_kit_mega_mega_armorer"`
				WinsKitMegaMegaArmorer                        float64  `json:"wins_kit_mega_mega_armorer"`
				GamesKitMegaMegaArmorer                       float64  `json:"games_kit_mega_mega_armorer"`
				VotesDwarfFortress                            float64  `json:"votes_Dwarf Fortress"`
				MegaTank                                      float64  `json:"mega_tank"`
				AssistsKitBasicSoloArmorsmith                 float64  `json:"assists_kit_basic_solo_armorsmith"`
				KitMegaMegaBaseballPlayer                     float64  `json:"kit_mega_mega_baseball-player"`
				KillsKitMegaMegaBaseballPlayer                float64  `json:"kills_kit_mega_mega_baseball-player"`
				AssistsKitMegaMegaBaseballPlayer              float64  `json:"assists_kit_mega_mega_baseball-player"`
				WinsKitMegaMegaBaseballPlayer                 float64  `json:"wins_kit_mega_mega_baseball-player"`
				SurvivedPlayersKitMegaMegaBaseballPlayer      float64  `json:"survived_players_kit_mega_mega_baseball-player"`
				GamesKitMegaMegaBaseballPlayer                float64  `json:"games_kit_mega_mega_baseball-player"`
				LossesKitMegaMegaBaseballPlayer               float64  `json:"losses_kit_mega_mega_baseball-player"`
				DeathsKitMegaMegaBaseballPlayer               float64  `json:"deaths_kit_mega_mega_baseball-player"`
				MegaMiningExpertise                           float64  `json:"mega_mining_expertise"`
				VotesSteampunk                                float64  `json:"votes_Steampunk"`
				MegaJuggernaut                                float64  `json:"mega_juggernaut"`
				MegaArrowRecovery                             float64  `json:"mega_arrow_recovery"`
				TeamJuggernaut                                float64  `json:"team_juggernaut"`
				SoloJuggernaut                                float64  `json:"solo_juggernaut"`
				KitMegaMegaCannoneer                          float64  `json:"kit_mega_mega_cannoneer"`
				SurvivedPlayersKitMegaMegaCannoneer           float64  `json:"survived_players_kit_mega_mega_cannoneer"`
				AssistsKitMegaMegaCannoneer                   float64  `json:"assists_kit_mega_mega_cannoneer"`
				WinsKitMegaMegaCannoneer                      float64  `json:"wins_kit_mega_mega_cannoneer"`
				GamesKitMegaMegaCannoneer                     float64  `json:"games_kit_mega_mega_cannoneer"`
				KillsKitMegaMegaCannoneer                     float64  `json:"kills_kit_mega_mega_cannoneer"`
				LossesKitMegaMegaCannoneer                    float64  `json:"losses_kit_mega_mega_cannoneer"`
				DeathsKitMegaMegaCannoneer                    float64  `json:"deaths_kit_mega_mega_cannoneer"`
				SoloArrowRecovery                             float64  `json:"solo_arrow_recovery"`
				SoloBulldozer                                 float64  `json:"solo_bulldozer"`
				LossesKitAdvancedSoloArmorer                  float64  `json:"losses_kit_advanced_solo_armorer"`
				DeathsKitAdvancedSoloArmorer                  float64  `json:"deaths_kit_advanced_solo_armorer"`
				AssistsKitAdvancedSoloArmorer                 float64  `json:"assists_kit_advanced_solo_armorer"`
				SoloSpeedBoost                                float64  `json:"solo_speed_boost"`
				VotesStrata                                   float64  `json:"votes_Strata"`
				VotesJinzhou                                  float64  `json:"votes_Jinzhou"`
				VotesCanopy                                   float64  `json:"votes_Canopy"`
				VotesTribal                                   float64  `json:"votes_Tribal"`
				VotesEntangled                                float64  `json:"votes_Entangled"`
				VotesSkychurch                                float64  `json:"votes_Skychurch"`
				KitMegaMegaWitch                              float64  `json:"kit_mega_mega_witch"`
				LossesKitMegaMegaWitch                        float64  `json:"losses_kit_mega_mega_witch"`
				AssistsKitMegaMegaWitch                       float64  `json:"assists_kit_mega_mega_witch"`
				DeathsKitMegaMegaWitch                        float64  `json:"deaths_kit_mega_mega_witch"`
				KillsKitMegaMegaWitch                         float64  `json:"kills_kit_mega_mega_witch"`
				SurvivedPlayersKitMegaMegaWitch               float64  `json:"survived_players_kit_mega_mega_witch"`
				WinsKitAdvancedSoloEnchanter                  float64  `json:"wins_kit_advanced_solo_enchanter"`
				GamesKitAdvancedSoloEnchanter                 float64  `json:"games_kit_advanced_solo_enchanter"`
				KillsKitAdvancedSoloEnchanter                 float64  `json:"kills_kit_advanced_solo_enchanter"`
				SurvivedPlayersKitAdvancedSoloEnchanter       float64  `json:"survived_players_kit_advanced_solo_enchanter"`
				LossesKitAdvancedSoloEnchanter                float64  `json:"losses_kit_advanced_solo_enchanter"`
				AssistsKitAdvancedSoloEnchanter               float64  `json:"assists_kit_advanced_solo_enchanter"`
				DeathsKitAdvancedSoloEnchanter                float64  `json:"deaths_kit_advanced_solo_enchanter"`
				GamesKitMegaMegaWitch                         float64  `json:"games_kit_mega_mega_witch"`
				WinsKitMegaMegaWitch                          float64  `json:"wins_kit_mega_mega_witch"`
				KitMegaMegaArmorsmith                         float64  `json:"kit_mega_mega_armorsmith"`
				KillsKitMegaMegaArmorsmith                    float64  `json:"kills_kit_mega_mega_armorsmith"`
				AssistsKitMegaMegaArmorsmith                  float64  `json:"assists_kit_mega_mega_armorsmith"`
				WinsKitMegaMegaArmorsmith                     float64  `json:"wins_kit_mega_mega_armorsmith"`
				GamesKitMegaMegaArmorsmith                    float64  `json:"games_kit_mega_mega_armorsmith"`
				SurvivedPlayersKitMegaMegaArmorsmith          float64  `json:"survived_players_kit_mega_mega_armorsmith"`
				GamesKitAdvancedSoloFarmer                    float64  `json:"games_kit_advanced_solo_farmer"`
				WinsKitAdvancedSoloFarmer                     float64  `json:"wins_kit_advanced_solo_farmer"`
				SurvivedPlayersKitAdvancedSoloFarmer          float64  `json:"survived_players_kit_advanced_solo_farmer"`
				KillsKitAdvancedSoloFarmer                    float64  `json:"kills_kit_advanced_solo_farmer"`
				MegaBlazingArrows                             float64  `json:"mega_blazing_arrows"`
				LossesKitMegaMegaArmorsmith                   float64  `json:"losses_kit_mega_mega_armorsmith"`
				DeathsKitMegaMegaArmorsmith                   float64  `json:"deaths_kit_mega_mega_armorsmith"`
				VotesTowers                                   float64  `json:"votes_Towers"`
				KitMegaMegaKnight                             float64  `json:"kit_mega_mega_knight"`
				AssistsKitMegaMegaKnight                      float64  `json:"assists_kit_mega_mega_knight"`
				KillsKitMegaMegaKnight                        float64  `json:"kills_kit_mega_mega_knight"`
				WinsKitMegaMegaKnight                         float64  `json:"wins_kit_mega_mega_knight"`
				DeathsKitMegaMegaKnight                       float64  `json:"deaths_kit_mega_mega_knight"`
				SurvivedPlayersKitMegaMegaKnight              float64  `json:"survived_players_kit_mega_mega_knight"`
				KitMegaMegaScout                              float64  `json:"kit_mega_mega_scout"`
				KitMegaMegaSkeletor                           float64  `json:"kit_mega_mega_skeletor"`
				KitMegaMegaHunter                             float64  `json:"kit_mega_mega_hunter"`
				KitMegaMegaHealer                             float64  `json:"kit_mega_mega_healer"`
				WinsKitDefendingTeamBaseballPlayer            float64  `json:"wins_kit_defending_team_baseball-player"`
				SurvivedPlayersKitDefendingTeamBaseballPlayer float64  `json:"survived_players_kit_defending_team_baseball-player"`
				KillsKitDefendingTeamBaseballPlayer           float64  `json:"kills_kit_defending_team_baseball-player"`
				GamesKitDefendingTeamBaseballPlayer           float64  `json:"games_kit_defending_team_baseball-player"`
				VotesAegis                                    float64  `json:"votes_Aegis"`
				KillsWeeklyB                                  float64  `json:"kills_weekly_b"`
				KillsMonthlyB                                 float64  `json:"kills_monthly_b"`
				GamesKitMegaMegaKnight                        float64  `json:"games_kit_mega_mega_knight"`
				LossesKitMegaMegaKnight                       float64  `json:"losses_kit_mega_mega_knight"`
				KillsWeeklyA                                  float64  `json:"kills_weekly_a"`
				HarvestingSeason                              float64  `json:"harvesting_season"`
				GamesKitMiningTeamCannoneer                   float64  `json:"games_kit_mining_team_cannoneer"`
				SurvivedPlayersKitMiningTeamCannoneer         float64  `json:"survived_players_kit_mining_team_cannoneer"`
				WinsKitMiningTeamCannoneer                    float64  `json:"wins_kit_mining_team_cannoneer"`
				KillsKitMiningTeamCannoneer                   float64  `json:"kills_kit_mining_team_cannoneer"`
				AssistsKitMiningTeamCannoneer                 float64  `json:"assists_kit_mining_team_cannoneer"`
				AssistsKitDefendingTeamBaseballPlayer         float64  `json:"assists_kit_defending_team_baseball-player"`
				LossesKitDefendingTeamBaseballPlayer          float64  `json:"losses_kit_defending_team_baseball-player"`
				DeathsKitDefendingTeamBaseballPlayer          float64  `json:"deaths_kit_defending_team_baseball-player"`
				LossesKitMiningTeamCannoneer                  float64  `json:"losses_kit_mining_team_cannoneer"`
				DeathsKitMiningTeamCannoneer                  float64  `json:"deaths_kit_mining_team_cannoneer"`
				KillsMonthlyA                                 float64  `json:"kills_monthly_a"`
				AssistsKitMegaMegaScout                       float64  `json:"assists_kit_mega_mega_scout"`
				KillsKitMegaMegaScout                         float64  `json:"kills_kit_mega_mega_scout"`
				DeathsKitMegaMegaScout                        float64  `json:"deaths_kit_mega_mega_scout"`
				SurvivedPlayersKitMegaMegaScout               float64  `json:"survived_players_kit_mega_mega_scout"`
				LossesKitMegaMegaScout                        float64  `json:"losses_kit_mega_mega_scout"`
				WinsKitMegaMegaScout                          float64  `json:"wins_kit_mega_mega_scout"`
				GamesKitMegaMegaScout                         float64  `json:"games_kit_mega_mega_scout"`
				LossesKitAdvancedSoloFarmer                   float64  `json:"losses_kit_advanced_solo_farmer"`
				DeathsKitAdvancedSoloFarmer                   float64  `json:"deaths_kit_advanced_solo_farmer"`
				KillsKitMegaMegaSkeletor                      float64  `json:"kills_kit_mega_mega_skeletor"`
				SurvivedPlayersKitMegaMegaSkeletor            float64  `json:"survived_players_kit_mega_mega_skeletor"`
				DeathsKitMegaMegaSkeletor                     float64  `json:"deaths_kit_mega_mega_skeletor"`
				LossesKitMegaMegaSkeletor                     float64  `json:"losses_kit_mega_mega_skeletor"`
				ActiveKitRANKED                               string   `json:"activeKit_RANKED"`
				WinstreakRanked                               float64  `json:"winstreak_ranked"`
				AssistsRanked                                 float64  `json:"assists_ranked"`
				Killstreak                                    float64  `json:"killstreak"`
				GamesRanked                                   float64  `json:"games_ranked"`
				WinsRankedNormal                              float64  `json:"wins_ranked_normal"`
				KillsKitRankedRankedDefault                   float64  `json:"kills_kit_ranked_ranked_default"`
				KillstreakRanked                              float64  `json:"killstreak_ranked"`
				KillsRankedNormal                             float64  `json:"kills_ranked_normal"`
				GamesKitRankedRankedDefault                   float64  `json:"games_kit_ranked_ranked_default"`
				SurvivedPlayersRanked                         float64  `json:"survived_players_ranked"`
				KillsRanked                                   float64  `json:"kills_ranked"`
				KillstreakKitRankedRankedDefault              float64  `json:"killstreak_kit_ranked_ranked_default"`
				Winstreak                                     float64  `json:"winstreak"`
				WinstreakKitRankedRankedDefault               float64  `json:"winstreak_kit_ranked_ranked_default"`
				AssistsKitRankedRankedDefault                 float64  `json:"assists_kit_ranked_ranked_default"`
				SurvivedPlayersKitRankedRankedDefault         float64  `json:"survived_players_kit_ranked_ranked_default"`
				WinsKitRankedRankedDefault                    float64  `json:"wins_kit_ranked_ranked_default"`
				WinsRanked                                    float64  `json:"wins_ranked"`
				HighestKillstreak                             float64  `json:"highestKillstreak"`
				HighestWinstreak                              float64  `json:"highestWinstreak"`
				LossesKitRankedRankedChampion                 float64  `json:"losses_kit_ranked_ranked_champion"`
				DeathsKitRankedRankedChampion                 float64  `json:"deaths_kit_ranked_ranked_champion"`
				AssistsKitRankedRankedChampion                float64  `json:"assists_kit_ranked_ranked_champion"`
				LossesRankedNormal                            float64  `json:"losses_ranked_normal"`
				DeathsRanked                                  float64  `json:"deaths_ranked"`
				LossesRanked                                  float64  `json:"losses_ranked"`
				DeathsRankedNormal                            float64  `json:"deaths_ranked_normal"`
				GamesKitRankedRankedChampion                  float64  `json:"games_kit_ranked_ranked_champion"`
				SurvivedPlayersKitRankedRankedChampion        float64  `json:"survived_players_kit_ranked_ranked_champion"`
				KillstreakKitRankedRankedChampion             float64  `json:"killstreak_kit_ranked_ranked_champion"`
				WinsKitRankedRankedChampion                   float64  `json:"wins_kit_ranked_ranked_champion"`
				KillsKitRankedRankedChampion                  float64  `json:"kills_kit_ranked_ranked_champion"`
				WinstreakKitRankedRankedChampion              float64  `json:"winstreak_kit_ranked_ranked_champion"`
				MegaInstantSmelting                           float64  `json:"mega_instant_smelting"`
				MegaNourishment                               float64  `json:"mega_nourishment"`
				TeamNourishment                               float64  `json:"team_nourishment"`
				SoloNourishment                               float64  `json:"solo_nourishment"`
				MegaNotoriety                                 float64  `json:"mega_notoriety"`
				TeamKnowledge                                 float64  `json:"team_knowledge"`
				RankedInstantSmelting                         float64  `json:"ranked_instant_smelting"`
				SoloKnowledge                                 float64  `json:"solo_knowledge"`
				TeamSavior                                    float64  `json:"team_savior"`
				TeamMarksmanship                              float64  `json:"team_marksmanship"`
				KillstreakKitBasicSoloArmorsmith              float64  `json:"killstreak_kit_basic_solo_armorsmith"`
				WinstreakKitBasicSoloArmorsmith               float64  `json:"winstreak_kit_basic_solo_armorsmith"`
				KillstreakSolo                                float64  `json:"killstreak_solo"`
				WinstreakSolo                                 float64  `json:"winstreak_solo"`
				WinstreakKitEnderchestSoloEnderchest          float64  `json:"winstreak_kit_enderchest_solo_enderchest"`
				GamesKitEnderchestSoloEnderchest              float64  `json:"games_kit_enderchest_solo_enderchest"`
				WinsKitEnderchestSoloEnderchest               float64  `json:"wins_kit_enderchest_solo_enderchest"`
				KillsKitEnderchestSoloEnderchest              float64  `json:"kills_kit_enderchest_solo_enderchest"`
				SurvivedPlayersKitEnderchestSoloEnderchest    float64  `json:"survived_players_kit_enderchest_solo_enderchest"`
				KillstreakKitEnderchestSoloEnderchest         float64  `json:"killstreak_kit_enderchest_solo_enderchest"`
				LossesKitEnderchestSoloEnderchest             float64  `json:"losses_kit_enderchest_solo_enderchest"`
				DeathsKitEnderchestSoloEnderchest             float64  `json:"deaths_kit_enderchest_solo_enderchest"`
				KillstreakKitMegaMegaKnight                   float64  `json:"killstreak_kit_mega_mega_knight"`
				WinstreakKitMegaMegaKnight                    float64  `json:"winstreak_kit_mega_mega_knight"`
				WinstreakMega                                 float64  `json:"winstreak_mega"`
				KillstreakMega                                float64  `json:"killstreak_mega"`
				KillstreakKitAdvancedSoloFarmer               float64  `json:"killstreak_kit_advanced_solo_farmer"`
				WinstreakKitAdvancedSoloFarmer                float64  `json:"winstreak_kit_advanced_solo_farmer"`
				WinstreakKitAdvancedSoloPyro                  float64  `json:"winstreak_kit_advanced_solo_pyro"`
				KillstreakKitAdvancedSoloPyro                 float64  `json:"killstreak_kit_advanced_solo_pyro"`
				SurvivedPlayersKitAdvancedSoloHunter          float64  `json:"survived_players_kit_advanced_solo_hunter"`
				DeathsKitAdvancedSoloHunter                   float64  `json:"deaths_kit_advanced_solo_hunter"`
				LossesKitAdvancedSoloHunter                   float64  `json:"losses_kit_advanced_solo_hunter"`
				SoloRevenge                                   float64  `json:"solo_revenge"`
				GamesKitBasicSoloFisherman                    float64  `json:"games_kit_basic_solo_fisherman"`
				SurvivedPlayersKitBasicSoloFisherman          float64  `json:"survived_players_kit_basic_solo_fisherman"`
				WinsKitBasicSoloFisherman                     float64  `json:"wins_kit_basic_solo_fisherman"`
				KillstreakKitBasicSoloFisherman               float64  `json:"killstreak_kit_basic_solo_fisherman"`
				WinstreakKitBasicSoloFisherman                float64  `json:"winstreak_kit_basic_solo_fisherman"`
				KillsKitBasicSoloFisherman                    float64  `json:"kills_kit_basic_solo_fisherman"`
				WinstreakKitAdvancedSoloCannoneer             float64  `json:"winstreak_kit_advanced_solo_cannoneer"`
				KillstreakKitAdvancedSoloCannoneer            float64  `json:"killstreak_kit_advanced_solo_cannoneer"`
				KillstreakKitMegaMegaSkeletor                 float64  `json:"killstreak_kit_mega_mega_skeletor"`
				WinsKitMegaMegaSkeletor                       float64  `json:"wins_kit_mega_mega_skeletor"`
				GamesKitMegaMegaSkeletor                      float64  `json:"games_kit_mega_mega_skeletor"`
				WinstreakKitMegaMegaSkeletor                  float64  `json:"winstreak_kit_mega_mega_skeletor"`
				AssistsKitMegaMegaSkeletor                    float64  `json:"assists_kit_mega_mega_skeletor"`
				TeamAnnoyOMite                                float64  `json:"team_annoy-o-mite"`
				TeamFat                                       float64  `json:"team_fat"`
				WinstreakKitAdvancedSoloEnchanter             float64  `json:"winstreak_kit_advanced_solo_enchanter"`
				KillstreakKitAdvancedSoloEnchanter            float64  `json:"killstreak_kit_advanced_solo_enchanter"`
				SoloAnnoyOMite                                float64  `json:"solo_annoy-o-mite"`
				WinsKitRankedRankedArmorer                    float64  `json:"wins_kit_ranked_ranked_armorer"`
				WinstreakKitRankedRankedArmorer               float64  `json:"winstreak_kit_ranked_ranked_armorer"`
				GamesKitRankedRankedArmorer                   float64  `json:"games_kit_ranked_ranked_armorer"`
				SurvivedPlayersKitRankedRankedArmorer         float64  `json:"survived_players_kit_ranked_ranked_armorer"`
				KillsKitRankedRankedArmorer                   float64  `json:"kills_kit_ranked_ranked_armorer"`
				KillstreakKitRankedRankedArmorer              float64  `json:"killstreak_kit_ranked_ranked_armorer"`
				AssistsKitRankedRankedArmorer                 float64  `json:"assists_kit_ranked_ranked_armorer"`
				LossesKitRankedRankedArmorer                  float64  `json:"losses_kit_ranked_ranked_armorer"`
				DeathsKitRankedRankedArmorer                  float64  `json:"deaths_kit_ranked_ranked_armorer"`
				WinstreakKitAdvancedSoloEnderman              float64  `json:"winstreak_kit_advanced_solo_enderman"`
				WinsKitAdvancedSoloEnderman                   float64  `json:"wins_kit_advanced_solo_enderman"`
				KillsKitAdvancedSoloEnderman                  float64  `json:"kills_kit_advanced_solo_enderman"`
				SurvivedPlayersKitAdvancedSoloEnderman        float64  `json:"survived_players_kit_advanced_solo_enderman"`
				GamesKitAdvancedSoloEnderman                  float64  `json:"games_kit_advanced_solo_enderman"`
				KillstreakKitAdvancedSoloEnderman             float64  `json:"killstreak_kit_advanced_solo_enderman"`
				AssistsKitAdvancedSoloFarmer                  float64  `json:"assists_kit_advanced_solo_farmer"`
				WinsKitBasicSoloSnowman                       float64  `json:"wins_kit_basic_solo_snowman"`
				WinstreakKitBasicSoloSnowman                  float64  `json:"winstreak_kit_basic_solo_snowman"`
				GamesKitBasicSoloSnowman                      float64  `json:"games_kit_basic_solo_snowman"`
				SurvivedPlayersKitBasicSoloSnowman            float64  `json:"survived_players_kit_basic_solo_snowman"`
				KillsKitBasicSoloSnowman                      float64  `json:"kills_kit_basic_solo_snowman"`
				KillstreakKitBasicSoloSnowman                 float64  `json:"killstreak_kit_basic_solo_snowman"`
				WinstreakKitBasicSoloSpeleologist             float64  `json:"winstreak_kit_basic_solo_speleologist"`
				WinsKitBasicSoloSpeleologist                  float64  `json:"wins_kit_basic_solo_speleologist"`
				KillsKitBasicSoloSpeleologist                 float64  `json:"kills_kit_basic_solo_speleologist"`
				GamesKitBasicSoloSpeleologist                 float64  `json:"games_kit_basic_solo_speleologist"`
				KillstreakKitBasicSoloSpeleologist            float64  `json:"killstreak_kit_basic_solo_speleologist"`
				SurvivedPlayersKitBasicSoloSpeleologist       float64  `json:"survived_players_kit_basic_solo_speleologist"`
				DeathsKitBasicSoloFisherman                   float64  `json:"deaths_kit_basic_solo_fisherman"`
				LossesKitBasicSoloFisherman                   float64  `json:"losses_kit_basic_solo_fisherman"`
				MegaMarksmanship                              float64  `json:"mega_marksmanship"`
				DeathsKitBasicSoloSnowman                     float64  `json:"deaths_kit_basic_solo_snowman"`
				LossesKitBasicSoloSnowman                     float64  `json:"losses_kit_basic_solo_snowman"`
				ActiveKillEffect                              string   `json:"activeKillEffect"`
				ActiveProjectileTrail                         string   `json:"activeProjectileTrail"`
				AssistsKitBasicSoloTroll                      float64  `json:"assists_kit_basic_solo_troll"`
				WinstreakKitBasicSoloTroll                    float64  `json:"winstreak_kit_basic_solo_troll"`
				KillsKitBasicSoloTroll                        float64  `json:"kills_kit_basic_solo_troll"`
				KillstreakKitBasicSoloTroll                   float64  `json:"killstreak_kit_basic_solo_troll"`
				SurvivedPlayersKitBasicSoloTroll              float64  `json:"survived_players_kit_basic_solo_troll"`
				WinsKitBasicSoloTroll                         float64  `json:"wins_kit_basic_solo_troll"`
				GamesKitBasicSoloTroll                        float64  `json:"games_kit_basic_solo_troll"`
				SoloFat                                       float64  `json:"solo_fat"`
				DeathsKitBasicSoloTroll                       float64  `json:"deaths_kit_basic_solo_troll"`
				LossesKitBasicSoloTroll                       float64  `json:"losses_kit_basic_solo_troll"`
				KillstreakKitSupportingTeamPharaoh            float64  `json:"killstreak_kit_supporting_team_pharaoh"`
				SurvivedPlayersKitSupportingTeamPharaoh       float64  `json:"survived_players_kit_supporting_team_pharaoh"`
				WinstreakKitSupportingTeamPharaoh             float64  `json:"winstreak_kit_supporting_team_pharaoh"`
				WinsKitSupportingTeamPharaoh                  float64  `json:"wins_kit_supporting_team_pharaoh"`
				KillstreakTeam                                float64  `json:"killstreak_team"`
				WinstreakTeam                                 float64  `json:"winstreak_team"`
				KillsKitSupportingTeamPharaoh                 float64  `json:"kills_kit_supporting_team_pharaoh"`
				GamesKitSupportingTeamPharaoh                 float64  `json:"games_kit_supporting_team_pharaoh"`
				AssistsKitBasicSoloFisherman                  float64  `json:"assists_kit_basic_solo_fisherman"`
				DeathsKitMegaMegaHunter                       float64  `json:"deaths_kit_mega_mega_hunter"`
				KillsKitMegaMegaHunter                        float64  `json:"kills_kit_mega_mega_hunter"`
				SurvivedPlayersKitMegaMegaHunter              float64  `json:"survived_players_kit_mega_mega_hunter"`
				LossesKitMegaMegaHunter                       float64  `json:"losses_kit_mega_mega_hunter"`
				ActiveVictoryDance                            string   `json:"activeVictoryDance"`
				KillstreakKitRankedRankedPyromancer           float64  `json:"killstreak_kit_ranked_ranked_pyromancer"`
				GamesKitRankedRankedPyromancer                float64  `json:"games_kit_ranked_ranked_pyromancer"`
				WinstreakKitRankedRankedPyromancer            float64  `json:"winstreak_kit_ranked_ranked_pyromancer"`
				KillsKitRankedRankedPyromancer                float64  `json:"kills_kit_ranked_ranked_pyromancer"`
				SurvivedPlayersKitRankedRankedPyromancer      float64  `json:"survived_players_kit_ranked_ranked_pyromancer"`
				WinsKitRankedRankedPyromancer                 float64  `json:"wins_kit_ranked_ranked_pyromancer"`
				DeathsKitRankedRankedPyromancer               float64  `json:"deaths_kit_ranked_ranked_pyromancer"`
				LossesKitRankedRankedPyromancer               float64  `json:"losses_kit_ranked_ranked_pyromancer"`
				AssistsKitRankedRankedPyromancer              float64  `json:"assists_kit_ranked_ranked_pyromancer"`
				KillstreakKitMegaMegaPaladin                  float64  `json:"killstreak_kit_mega_mega_paladin"`
				WinstreakKitMegaMegaPaladin                   float64  `json:"winstreak_kit_mega_mega_paladin"`
				KillstreakKitDefendingTeamArmorer             float64  `json:"killstreak_kit_defending_team_armorer"`
				WinstreakKitDefendingTeamArmorer              float64  `json:"winstreak_kit_defending_team_armorer"`
				KillsKitBasicSoloPharaoh                      float64  `json:"kills_kit_basic_solo_pharaoh"`
				WinsKitBasicSoloPharaoh                       float64  `json:"wins_kit_basic_solo_pharaoh"`
				GamesKitBasicSoloPharaoh                      float64  `json:"games_kit_basic_solo_pharaoh"`
				WinstreakKitBasicSoloPharaoh                  float64  `json:"winstreak_kit_basic_solo_pharaoh"`
				SurvivedPlayersKitBasicSoloPharaoh            float64  `json:"survived_players_kit_basic_solo_pharaoh"`
				KillstreakKitBasicSoloPharaoh                 float64  `json:"killstreak_kit_basic_solo_pharaoh"`
				LossesKitBasicSoloPharaoh                     float64  `json:"losses_kit_basic_solo_pharaoh"`
				DeathsKitBasicSoloPharaoh                     float64  `json:"deaths_kit_basic_solo_pharaoh"`
				AssistsKitBasicSoloPharaoh                    float64  `json:"assists_kit_basic_solo_pharaoh"`
				KitMegaMegaHellhound                          float64  `json:"kit_mega_mega_hellhound"`
				KillstreakKitMegaMegaArmorer                  float64  `json:"killstreak_kit_mega_mega_armorer"`
				WinstreakKitMegaMegaArmorer                   float64  `json:"winstreak_kit_mega_mega_armorer"`
				WinstreakKitAdvancedSoloArmorer               float64  `json:"winstreak_kit_advanced_solo_armorer"`
				KillstreakKitAdvancedSoloArmorer              float64  `json:"killstreak_kit_advanced_solo_armorer"`
				WinstreakKitAdvancedSoloHunter                float64  `json:"winstreak_kit_advanced_solo_hunter"`
				KillsKitAdvancedSoloHunter                    float64  `json:"kills_kit_advanced_solo_hunter"`
				GamesKitAdvancedSoloHunter                    float64  `json:"games_kit_advanced_solo_hunter"`
				WinsKitAdvancedSoloHunter                     float64  `json:"wins_kit_advanced_solo_hunter"`
				KillstreakKitAdvancedSoloHunter               float64  `json:"killstreak_kit_advanced_solo_hunter"`
				LossesKitAdvancedSoloEnderman                 float64  `json:"losses_kit_advanced_solo_enderman"`
				DeathsKitAdvancedSoloEnderman                 float64  `json:"deaths_kit_advanced_solo_enderman"`
				WinstreakKitMiningTeamCannoneer               float64  `json:"winstreak_kit_mining_team_cannoneer"`
				KillstreakKitMiningTeamCannoneer              float64  `json:"killstreak_kit_mining_team_cannoneer"`
				SoloBridger                                   float64  `json:"solo_bridger"`
				SoloEnvironmentalExpert                       float64  `json:"solo_environmental_expert"`
				SoloLuckyCharm                                float64  `json:"solo_lucky_charm"`
				MeleeKillsSolo                                float64  `json:"melee_kills_solo"`
				MeleeKillsKitBasicSoloGrenade                 float64  `json:"melee_kills_kit_basic_solo_grenade"`
				TimePlayedKitBasicSoloGrenade                 float64  `json:"time_played_kit_basic_solo_grenade"`
				ChestsOpened                                  float64  `json:"chests_opened"`
				MostKillsGameKitBasicSoloGrenade              float64  `json:"most_kills_game_kit_basic_solo_grenade"`
				MeleeKills                                    float64  `json:"melee_kills"`
				KillsKitBasicSoloGrenade                      float64  `json:"kills_kit_basic_solo_grenade"`
				GamesKitBasicSoloGrenade                      float64  `json:"games_kit_basic_solo_grenade"`
				SurvivedPlayersKitBasicSoloGrenade            float64  `json:"survived_players_kit_basic_solo_grenade"`
				MostKillsGameSolo                             float64  `json:"most_kills_game_solo"`
				DeathsKitBasicSoloGrenade                     float64  `json:"deaths_kit_basic_solo_grenade"`
				ChestsOpenedSolo                              float64  `json:"chests_opened_solo"`
				LossesKitBasicSoloGrenade                     float64  `json:"losses_kit_basic_solo_grenade"`
				TimePlayedSolo                                float64  `json:"time_played_solo"`
				TimePlayed                                    float64  `json:"time_played"`
				MostKillsGame                                 float64  `json:"most_kills_game"`
				ChestsOpenedKitBasicSoloGrenade               float64  `json:"chests_opened_kit_basic_solo_grenade"`
				LongestBowShotKitBasicSoloGrenade             float64  `json:"longest_bow_shot_kit_basic_solo_grenade"`
				LongestBowShot                                float64  `json:"longest_bow_shot"`
				FastestWin                                    float64  `json:"fastest_win"`
				FastestWinSolo                                float64  `json:"fastest_win_solo"`
				FastestWinKitBasicSoloGrenade                 float64  `json:"fastest_win_kit_basic_solo_grenade"`
				LongestBowShotSolo                            float64  `json:"longest_bow_shot_solo"`
				VoidKills                                     float64  `json:"void_kills"`
				ArrowsShotKitBasicSoloGrenade                 float64  `json:"arrows_shot_kit_basic_solo_grenade"`
				WinstreakKitBasicSoloGrenade                  float64  `json:"winstreak_kit_basic_solo_grenade"`
				VoidKillsKitBasicSoloGrenade                  float64  `json:"void_kills_kit_basic_solo_grenade"`
				ArrowsHitSolo                                 float64  `json:"arrows_hit_solo"`
				WinsKitBasicSoloGrenade                       float64  `json:"wins_kit_basic_solo_grenade"`
				KillstreakKitBasicSoloGrenade                 float64  `json:"killstreak_kit_basic_solo_grenade"`
				ArrowsHitKitBasicSoloGrenade                  float64  `json:"arrows_hit_kit_basic_solo_grenade"`
				VoidKillsSolo                                 float64  `json:"void_kills_solo"`
				ArrowsShotSolo                                float64  `json:"arrows_shot_solo"`
				AssistsKitMegaMegaHellhound                   float64  `json:"assists_kit_mega_mega_hellhound"`
				LossesKitMegaMegaHellhound                    float64  `json:"losses_kit_mega_mega_hellhound"`
				DeathsKitMegaMegaHellhound                    float64  `json:"deaths_kit_mega_mega_hellhound"`
				ChestsOpenedKitMegaMegaHellhound              float64  `json:"chests_opened_kit_mega_mega_hellhound"`
				ChestsOpenedMega                              float64  `json:"chests_opened_mega"`
				MostKillsGameMega                             float64  `json:"most_kills_game_mega"`
				MeleeKillsMega                                float64  `json:"melee_kills_mega"`
				MostKillsGameKitMegaMegaHellhound             float64  `json:"most_kills_game_kit_mega_mega_hellhound"`
				MeleeKillsKitMegaMegaHellhound                float64  `json:"melee_kills_kit_mega_mega_hellhound"`
				KillsKitMegaMegaHellhound                     float64  `json:"kills_kit_mega_mega_hellhound"`
				TimePlayedKitMegaMegaHellhound                float64  `json:"time_played_kit_mega_mega_hellhound"`
				SurvivedPlayersKitMegaMegaHellhound           float64  `json:"survived_players_kit_mega_mega_hellhound"`
				TimePlayedMega                                float64  `json:"time_played_mega"`
				AssistsKitBasicSoloGrenade                    float64  `json:"assists_kit_basic_solo_grenade"`
				MegaLuckyCharm                                float64  `json:"mega_lucky_charm"`
				MegaBridger                                   float64  `json:"mega_bridger"`
				DeathsKitBasicSoloFrog                        float64  `json:"deaths_kit_basic_solo_frog"`
				ChestsOpenedKitBasicSoloFrog                  float64  `json:"chests_opened_kit_basic_solo_frog"`
				TimePlayedKitBasicSoloFrog                    float64  `json:"time_played_kit_basic_solo_frog"`
				LossesKitBasicSoloFrog                        float64  `json:"losses_kit_basic_solo_frog"`
				FastestWinKitBasicSoloFrog                    float64  `json:"fastest_win_kit_basic_solo_frog"`
				SurvivedPlayersKitBasicSoloFrog               float64  `json:"survived_players_kit_basic_solo_frog"`
				VoidKillsKitBasicSoloFrog                     float64  `json:"void_kills_kit_basic_solo_frog"`
				WinstreakKitBasicSoloFrog                     float64  `json:"winstreak_kit_basic_solo_frog"`
				GamesKitBasicSoloFrog                         float64  `json:"games_kit_basic_solo_frog"`
				MostKillsGameKitBasicSoloFrog                 float64  `json:"most_kills_game_kit_basic_solo_frog"`
				KillstreakKitBasicSoloFrog                    float64  `json:"killstreak_kit_basic_solo_frog"`
				WinsKitBasicSoloFrog                          float64  `json:"wins_kit_basic_solo_frog"`
				MeleeKillsKitBasicSoloFrog                    float64  `json:"melee_kills_kit_basic_solo_frog"`
				KillsKitBasicSoloFrog                         float64  `json:"kills_kit_basic_solo_frog"`
				LongestBowShotKitBasicSoloFrog                float64  `json:"longest_bow_shot_kit_basic_solo_frog"`
				ArrowsShotKitBasicSoloFrog                    float64  `json:"arrows_shot_kit_basic_solo_frog"`
				ArrowsHitKitBasicSoloFrog                     float64  `json:"arrows_hit_kit_basic_solo_frog"`
				AssistsKitBasicSoloFrog                       float64  `json:"assists_kit_basic_solo_frog"`
				LongestBowShotKitBasicSoloPrincess            float64  `json:"longest_bow_shot_kit_basic_solo_princess"`
				FastestWinKitBasicSoloPrincess                float64  `json:"fastest_win_kit_basic_solo_princess"`
				VoidKillsKitBasicSoloPrincess                 float64  `json:"void_kills_kit_basic_solo_princess"`
				KillsKitBasicSoloPrincess                     float64  `json:"kills_kit_basic_solo_princess"`
				WinstreakKitBasicSoloPrincess                 float64  `json:"winstreak_kit_basic_solo_princess"`
				ArrowsShotKitBasicSoloPrincess                float64  `json:"arrows_shot_kit_basic_solo_princess"`
				GamesKitBasicSoloPrincess                     float64  `json:"games_kit_basic_solo_princess"`
				TimePlayedKitBasicSoloPrincess                float64  `json:"time_played_kit_basic_solo_princess"`
				SurvivedPlayersKitBasicSoloPrincess           float64  `json:"survived_players_kit_basic_solo_princess"`
				WinsKitBasicSoloPrincess                      float64  `json:"wins_kit_basic_solo_princess"`
				ChestsOpenedKitBasicSoloPrincess              float64  `json:"chests_opened_kit_basic_solo_princess"`
				ArrowsHitKitBasicSoloPrincess                 float64  `json:"arrows_hit_kit_basic_solo_princess"`
				KillstreakKitBasicSoloPrincess                float64  `json:"killstreak_kit_basic_solo_princess"`
				MostKillsGameKitBasicSoloPrincess             float64  `json:"most_kills_game_kit_basic_solo_princess"`
				MeleeKillsKitBasicSoloPrincess                float64  `json:"melee_kills_kit_basic_solo_princess"`
				SurvivedPlayersKitBasicSoloDisco              float64  `json:"survived_players_kit_basic_solo_disco"`
				VoidKillsKitBasicSoloDisco                    float64  `json:"void_kills_kit_basic_solo_disco"`
				MostKillsGameKitBasicSoloDisco                float64  `json:"most_kills_game_kit_basic_solo_disco"`
				DeathsKitBasicSoloDisco                       float64  `json:"deaths_kit_basic_solo_disco"`
				KillsKitBasicSoloDisco                        float64  `json:"kills_kit_basic_solo_disco"`
				LossesKitBasicSoloDisco                       float64  `json:"losses_kit_basic_solo_disco"`
				AssistsKitBasicSoloDisco                      float64  `json:"assists_kit_basic_solo_disco"`
				TimePlayedKitBasicSoloDisco                   float64  `json:"time_played_kit_basic_solo_disco"`
				GamesKitBasicSoloDisco                        float64  `json:"games_kit_basic_solo_disco"`
				FastestWinKitBasicSoloDisco                   float64  `json:"fastest_win_kit_basic_solo_disco"`
				WinsKitBasicSoloDisco                         float64  `json:"wins_kit_basic_solo_disco"`
				WinstreakKitBasicSoloDisco                    float64  `json:"winstreak_kit_basic_solo_disco"`
				KillstreakKitBasicSoloDisco                   float64  `json:"killstreak_kit_basic_solo_disco"`
				ChestsOpenedKitBasicSoloDisco                 float64  `json:"chests_opened_kit_basic_solo_disco"`
				MeleeKillsKitBasicSoloDisco                   float64  `json:"melee_kills_kit_basic_solo_disco"`
				LongestBowShotKitBasicSoloDisco               float64  `json:"longest_bow_shot_kit_basic_solo_disco"`
				ArrowsHitKitBasicSoloDisco                    float64  `json:"arrows_hit_kit_basic_solo_disco"`
				ArrowsShotKitBasicSoloDisco                   float64  `json:"arrows_shot_kit_basic_solo_disco"`
				ChestsOpenedKitBasicSoloBatguy                float64  `json:"chests_opened_kit_basic_solo_batguy"`
				DeathsKitBasicSoloBatguy                      float64  `json:"deaths_kit_basic_solo_batguy"`
				TimePlayedKitBasicSoloBatguy                  float64  `json:"time_played_kit_basic_solo_batguy"`
				LossesKitBasicSoloBatguy                      float64  `json:"losses_kit_basic_solo_batguy"`
				SurvivedPlayersKitBasicSoloBatguy             float64  `json:"survived_players_kit_basic_solo_batguy"`
				LongestBowShotKitBasicSoloBatguy              float64  `json:"longest_bow_shot_kit_basic_solo_batguy"`
				FastestWinKitBasicSoloBatguy                  float64  `json:"fastest_win_kit_basic_solo_batguy"`
				ArrowsShotKitBasicSoloBatguy                  float64  `json:"arrows_shot_kit_basic_solo_batguy"`
				VoidKillsKitBasicSoloBatguy                   float64  `json:"void_kills_kit_basic_solo_batguy"`
				MostKillsGameKitBasicSoloBatguy               float64  `json:"most_kills_game_kit_basic_solo_batguy"`
				GamesKitBasicSoloBatguy                       float64  `json:"games_kit_basic_solo_batguy"`
				KillsKitBasicSoloBatguy                       float64  `json:"kills_kit_basic_solo_batguy"`
				ArrowsHitKitBasicSoloBatguy                   float64  `json:"arrows_hit_kit_basic_solo_batguy"`
				MeleeKillsKitBasicSoloBatguy                  float64  `json:"melee_kills_kit_basic_solo_batguy"`
				WinsKitBasicSoloBatguy                        float64  `json:"wins_kit_basic_solo_batguy"`
				WinstreakKitBasicSoloBatguy                   float64  `json:"winstreak_kit_basic_solo_batguy"`
				KillstreakKitBasicSoloBatguy                  float64  `json:"killstreak_kit_basic_solo_batguy"`
				LossesKitBasicSoloPrincess                    float64  `json:"losses_kit_basic_solo_princess"`
				DeathsKitBasicSoloPrincess                    float64  `json:"deaths_kit_basic_solo_princess"`
				AssistsKitBasicSoloPrincess                   float64  `json:"assists_kit_basic_solo_princess"`
				MobKills                                      float64  `json:"mob_kills"`
				MobKillsKitBasicSoloPrincess                  float64  `json:"mob_kills_kit_basic_solo_princess"`
				MobKillsSolo                                  float64  `json:"mob_kills_solo"`
				VoidKillsKitMegaMegaArmorsmith                float64  `json:"void_kills_kit_mega_mega_armorsmith"`
				ChestsOpenedKitMegaMegaArmorsmith             float64  `json:"chests_opened_kit_mega_mega_armorsmith"`
				MostKillsGameKitMegaMegaArmorsmith            float64  `json:"most_kills_game_kit_mega_mega_armorsmith"`
				VoidKillsMega                                 float64  `json:"void_kills_mega"`
				MeleeKillsKitMegaMegaArmorsmith               float64  `json:"melee_kills_kit_mega_mega_armorsmith"`
				TimePlayedKitMegaMegaArmorsmith               float64  `json:"time_played_kit_mega_mega_armorsmith"`
				FastestWinKitAdvancedSoloCannoneer            float64  `json:"fastest_win_kit_advanced_solo_cannoneer"`
				TimePlayedKitAdvancedSoloCannoneer            float64  `json:"time_played_kit_advanced_solo_cannoneer"`
				VoidKillsKitAdvancedSoloCannoneer             float64  `json:"void_kills_kit_advanced_solo_cannoneer"`
				ChestsOpenedKitAdvancedSoloCannoneer          float64  `json:"chests_opened_kit_advanced_solo_cannoneer"`
				MostKillsGameKitAdvancedSoloCannoneer         float64  `json:"most_kills_game_kit_advanced_solo_cannoneer"`
				FastestWinKitBasicSoloArmorsmith              float64  `json:"fastest_win_kit_basic_solo_armorsmith"`
				LongestBowShotKitBasicSoloArmorsmith          float64  `json:"longest_bow_shot_kit_basic_solo_armorsmith"`
				ArrowsHitKitBasicSoloArmorsmith               float64  `json:"arrows_hit_kit_basic_solo_armorsmith"`
				ChestsOpenedKitBasicSoloArmorsmith            float64  `json:"chests_opened_kit_basic_solo_armorsmith"`
				MostKillsGameKitBasicSoloArmorsmith           float64  `json:"most_kills_game_kit_basic_solo_armorsmith"`
				ArrowsShotKitBasicSoloArmorsmith              float64  `json:"arrows_shot_kit_basic_solo_armorsmith"`
				VoidKillsKitBasicSoloArmorsmith               float64  `json:"void_kills_kit_basic_solo_armorsmith"`
				TimePlayedKitBasicSoloArmorsmith              float64  `json:"time_played_kit_basic_solo_armorsmith"`
				MeleeKillsKitBasicSoloArmorsmith              float64  `json:"melee_kills_kit_basic_solo_armorsmith"`
				ChestsOpenedKitAdvancedSoloPyro               float64  `json:"chests_opened_kit_advanced_solo_pyro"`
				MostKillsGameKitAdvancedSoloPyro              float64  `json:"most_kills_game_kit_advanced_solo_pyro"`
				MeleeKillsKitAdvancedSoloPyro                 float64  `json:"melee_kills_kit_advanced_solo_pyro"`
				TimePlayedKitAdvancedSoloPyro                 float64  `json:"time_played_kit_advanced_solo_pyro"`
				LongestBowShotKitAdvancedSoloPyro             float64  `json:"longest_bow_shot_kit_advanced_solo_pyro"`
				FastestWinKitAdvancedSoloPyro                 float64  `json:"fastest_win_kit_advanced_solo_pyro"`
				ArrowsShotKitAdvancedSoloPyro                 float64  `json:"arrows_shot_kit_advanced_solo_pyro"`
				VoidKillsKitAdvancedSoloPyro                  float64  `json:"void_kills_kit_advanced_solo_pyro"`
				ArrowsHitKitAdvancedSoloPyro                  float64  `json:"arrows_hit_kit_advanced_solo_pyro"`
				VotesFestiveTribute                           float64  `json:"votes_Festive Tribute"`
				TimePlayedKitEnderchestSoloEnderchest         float64  `json:"time_played_kit_enderchest_solo_enderchest"`
				ChestsOpenedKitEnderchestSoloEnderchest       float64  `json:"chests_opened_kit_enderchest_solo_enderchest"`
				FastestWinTeam                                float64  `json:"fastest_win_team"`
				FastestWinKitDefendingTeamArmorer             float64  `json:"fastest_win_kit_defending_team_armorer"`
				MeleeKillsKitDefendingTeamArmorer             float64  `json:"melee_kills_kit_defending_team_armorer"`
				MeleeKillsTeam                                float64  `json:"melee_kills_team"`
				TimePlayedKitDefendingTeamArmorer             float64  `json:"time_played_kit_defending_team_armorer"`
				MostKillsGameTeam                             float64  `json:"most_kills_game_team"`
				ChestsOpenedTeam                              float64  `json:"chests_opened_team"`
				TimePlayedTeam                                float64  `json:"time_played_team"`
				MostKillsGameKitDefendingTeamArmorer          float64  `json:"most_kills_game_kit_defending_team_armorer"`
				ChestsOpenedKitDefendingTeamArmorer           float64  `json:"chests_opened_kit_defending_team_armorer"`
				VoidKillsTeam                                 float64  `json:"void_kills_team"`
				VoidKillsKitDefendingTeamArmorer              float64  `json:"void_kills_kit_defending_team_armorer"`
				LongestBowShotTeam                            float64  `json:"longest_bow_shot_team"`
				LongestBowShotKitDefendingTeamArmorer         float64  `json:"longest_bow_shot_kit_defending_team_armorer"`
				ArrowsShotTeam                                float64  `json:"arrows_shot_team"`
				ArrowsHitTeam                                 float64  `json:"arrows_hit_team"`
				ArrowsHitKitDefendingTeamArmorer              float64  `json:"arrows_hit_kit_defending_team_armorer"`
				ArrowsShotKitDefendingTeamArmorer             float64  `json:"arrows_shot_kit_defending_team_armorer"`
				FastestWinKitBasicSoloSnowman                 float64  `json:"fastest_win_kit_basic_solo_snowman"`
				VoidKillsKitBasicSoloSnowman                  float64  `json:"void_kills_kit_basic_solo_snowman"`
				MeleeKillsKitBasicSoloSnowman                 float64  `json:"melee_kills_kit_basic_solo_snowman"`
				ChestsOpenedKitBasicSoloSnowman               float64  `json:"chests_opened_kit_basic_solo_snowman"`
				MostKillsGameKitBasicSoloSnowman              float64  `json:"most_kills_game_kit_basic_solo_snowman"`
				TimePlayedKitBasicSoloSnowman                 float64  `json:"time_played_kit_basic_solo_snowman"`
				FastestWinKitAdvancedSoloEnderman             float64  `json:"fastest_win_kit_advanced_solo_enderman"`
				MostKillsGameKitAdvancedSoloEnderman          float64  `json:"most_kills_game_kit_advanced_solo_enderman"`
				TimePlayedKitAdvancedSoloEnderman             float64  `json:"time_played_kit_advanced_solo_enderman"`
				VoidKillsKitAdvancedSoloEnderman              float64  `json:"void_kills_kit_advanced_solo_enderman"`
				ChestsOpenedKitAdvancedSoloEnderman           float64  `json:"chests_opened_kit_advanced_solo_enderman"`
				MeleeKillsKitAdvancedSoloEnderman             float64  `json:"melee_kills_kit_advanced_solo_enderman"`
				SurvivedPlayersKitBasicSoloScout              float64  `json:"survived_players_kit_basic_solo_scout"`
				LossesKitBasicSoloScout                       float64  `json:"losses_kit_basic_solo_scout"`
				DeathsKitBasicSoloScout                       float64  `json:"deaths_kit_basic_solo_scout"`
				TimePlayedKitBasicSoloScout                   float64  `json:"time_played_kit_basic_solo_scout"`
				ChestsOpenedKitBasicSoloScout                 float64  `json:"chests_opened_kit_basic_solo_scout"`
				AssistsKitAdvancedSoloEnderman                float64  `json:"assists_kit_advanced_solo_enderman"`
				MeleeKillsKitAdvancedSoloCannoneer            float64  `json:"melee_kills_kit_advanced_solo_cannoneer"`
				LongestBowShotKitAdvancedSoloCannoneer        float64  `json:"longest_bow_shot_kit_advanced_solo_cannoneer"`
				ArrowsShotKitAdvancedSoloCannoneer            float64  `json:"arrows_shot_kit_advanced_solo_cannoneer"`
				ArrowsHitKitAdvancedSoloCannoneer             float64  `json:"arrows_hit_kit_advanced_solo_cannoneer"`
				FastestWinKitRankedRankedPyromancer           float64  `json:"fastest_win_kit_ranked_ranked_pyromancer"`
				FastestWinRanked                              float64  `json:"fastest_win_ranked"`
				ChestsOpenedRanked                            float64  `json:"chests_opened_ranked"`
				TimePlayedRanked                              float64  `json:"time_played_ranked"`
				MostKillsGameKitRankedRankedPyromancer        float64  `json:"most_kills_game_kit_ranked_ranked_pyromancer"`
				VoidKillsRanked                               float64  `json:"void_kills_ranked"`
				TimePlayedKitRankedRankedPyromancer           float64  `json:"time_played_kit_ranked_ranked_pyromancer"`
				MostKillsGameRanked                           float64  `json:"most_kills_game_ranked"`
				ChestsOpenedKitRankedRankedPyromancer         float64  `json:"chests_opened_kit_ranked_ranked_pyromancer"`
				VoidKillsKitRankedRankedPyromancer            float64  `json:"void_kills_kit_ranked_ranked_pyromancer"`
				LongestBowShotKitBasicSoloEnergix             float64  `json:"longest_bow_shot_kit_basic_solo_energix"`
				FastestWinKitBasicSoloEnergix                 float64  `json:"fastest_win_kit_basic_solo_energix"`
				KillsKitBasicSoloEnergix                      float64  `json:"kills_kit_basic_solo_energix"`
				VoidKillsKitBasicSoloEnergix                  float64  `json:"void_kills_kit_basic_solo_energix"`
				MostKillsGameKitBasicSoloEnergix              float64  `json:"most_kills_game_kit_basic_solo_energix"`
				ArrowsShotKitBasicSoloEnergix                 float64  `json:"arrows_shot_kit_basic_solo_energix"`
				ArrowsHitKitBasicSoloEnergix                  float64  `json:"arrows_hit_kit_basic_solo_energix"`
				ChestsOpenedKitBasicSoloEnergix               float64  `json:"chests_opened_kit_basic_solo_energix"`
				KillstreakKitBasicSoloEnergix                 float64  `json:"killstreak_kit_basic_solo_energix"`
				TimePlayedKitBasicSoloEnergix                 float64  `json:"time_played_kit_basic_solo_energix"`
				SurvivedPlayersKitBasicSoloEnergix            float64  `json:"survived_players_kit_basic_solo_energix"`
				MeleeKillsKitBasicSoloEnergix                 float64  `json:"melee_kills_kit_basic_solo_energix"`
				WinsKitBasicSoloEnergix                       float64  `json:"wins_kit_basic_solo_energix"`
				GamesKitBasicSoloEnergix                      float64  `json:"games_kit_basic_solo_energix"`
				WinstreakKitBasicSoloEnergix                  float64  `json:"winstreak_kit_basic_solo_energix"`
				LossesKitBasicSoloEnergix                     float64  `json:"losses_kit_basic_solo_energix"`
				DeathsKitBasicSoloEnergix                     float64  `json:"deaths_kit_basic_solo_energix"`
				FastestWinKitRankedRankedArmorer              float64  `json:"fastest_win_kit_ranked_ranked_armorer"`
				MeleeKillsKitRankedRankedArmorer              float64  `json:"melee_kills_kit_ranked_ranked_armorer"`
				TimePlayedKitRankedRankedArmorer              float64  `json:"time_played_kit_ranked_ranked_armorer"`
				MeleeKillsRanked                              float64  `json:"melee_kills_ranked"`
				VoidKillsKitRankedRankedArmorer               float64  `json:"void_kills_kit_ranked_ranked_armorer"`
				ChestsOpenedKitRankedRankedArmorer            float64  `json:"chests_opened_kit_ranked_ranked_armorer"`
				MostKillsGameKitRankedRankedArmorer           float64  `json:"most_kills_game_kit_ranked_ranked_armorer"`
				LongestBowShotRanked                          float64  `json:"longest_bow_shot_ranked"`
				LongestBowShotKitRankedRankedArmorer          float64  `json:"longest_bow_shot_kit_ranked_ranked_armorer"`
				ArrowsHitRanked                               float64  `json:"arrows_hit_ranked"`
				ArrowsHitKitRankedRankedArmorer               float64  `json:"arrows_hit_kit_ranked_ranked_armorer"`
				ArrowsShotRanked                              float64  `json:"arrows_shot_ranked"`
				ArrowsShotKitRankedRankedArmorer              float64  `json:"arrows_shot_kit_ranked_ranked_armorer"`
				AssistsKitBasicSoloEnergix                    float64  `json:"assists_kit_basic_solo_energix"`
				ArrowsShotKitAdvancedSoloEnderman             float64  `json:"arrows_shot_kit_advanced_solo_enderman"`
				LongestBowShotKitMegaMegaArmorsmith           float64  `json:"longest_bow_shot_kit_mega_mega_armorsmith"`
				LongestBowShotMega                            float64  `json:"longest_bow_shot_mega"`
				ArrowsShotKitMegaMegaArmorsmith               float64  `json:"arrows_shot_kit_mega_mega_armorsmith"`
				ArrowsHitKitMegaMegaArmorsmith                float64  `json:"arrows_hit_kit_mega_mega_armorsmith"`
				ArrowsShotMega                                float64  `json:"arrows_shot_mega"`
				ArrowsHitMega                                 float64  `json:"arrows_hit_mega"`
				LongestBowShotKitAdvancedSoloEnderman         float64  `json:"longest_bow_shot_kit_advanced_solo_enderman"`
				ArrowsHitKitAdvancedSoloEnderman              float64  `json:"arrows_hit_kit_advanced_solo_enderman"`
				MostKillsGameKitBasicSoloRookie               float64  `json:"most_kills_game_kit_basic_solo_rookie"`
				SurvivedPlayersKitBasicSoloRookie             float64  `json:"survived_players_kit_basic_solo_rookie"`
				GamesKitBasicSoloRookie                       float64  `json:"games_kit_basic_solo_rookie"`
				KillsKitBasicSoloRookie                       float64  `json:"kills_kit_basic_solo_rookie"`
				TimePlayedKitBasicSoloRookie                  float64  `json:"time_played_kit_basic_solo_rookie"`
				LossesKitBasicSoloRookie                      float64  `json:"losses_kit_basic_solo_rookie"`
				ChestsOpenedKitBasicSoloRookie                float64  `json:"chests_opened_kit_basic_solo_rookie"`
				DeathsKitBasicSoloRookie                      float64  `json:"deaths_kit_basic_solo_rookie"`
				MeleeKillsKitBasicSoloRookie                  float64  `json:"melee_kills_kit_basic_solo_rookie"`
				MobKillsKitBasicSoloEnergix                   float64  `json:"mob_kills_kit_basic_solo_energix"`
				GamesPlayedSkywars                            float64  `json:"games_played_skywars"`
				LastMode                                      string   `json:"lastMode"`
				MeleeKillsKitMegaMegaKnight                   float64  `json:"melee_kills_kit_mega_mega_knight"`
				ChestsOpenedKitMegaMegaKnight                 float64  `json:"chests_opened_kit_mega_mega_knight"`
				MostKillsGameKitMegaMegaKnight                float64  `json:"most_kills_game_kit_mega_mega_knight"`
				TimePlayedKitMegaMegaKnight                   float64  `json:"time_played_kit_mega_mega_knight"`
				VoidKillsKitMegaMegaKnight                    float64  `json:"void_kills_kit_mega_mega_knight"`
				LongestBowShotKitMegaMegaKnight               float64  `json:"longest_bow_shot_kit_mega_mega_knight"`
				ArrowsShotKitMegaMegaKnight                   float64  `json:"arrows_shot_kit_mega_mega_knight"`
				ArrowsHitKitMegaMegaKnight                    float64  `json:"arrows_hit_kit_mega_mega_knight"`
				LongestBowShotKitAttackingTeamEnderman        float64  `json:"longest_bow_shot_kit_attacking_team_enderman"`
				KillsKitAttackingTeamEnderman                 float64  `json:"kills_kit_attacking_team_enderman"`
				LossesKitAttackingTeamEnderman                float64  `json:"losses_kit_attacking_team_enderman"`
				MeleeKillsKitAttackingTeamEnderman            float64  `json:"melee_kills_kit_attacking_team_enderman"`
				SurvivedPlayersKitAttackingTeamEnderman       float64  `json:"survived_players_kit_attacking_team_enderman"`
				ArrowsShotKitAttackingTeamEnderman            float64  `json:"arrows_shot_kit_attacking_team_enderman"`
				DeathsKitAttackingTeamEnderman                float64  `json:"deaths_kit_attacking_team_enderman"`
				TimePlayedKitAttackingTeamEnderman            float64  `json:"time_played_kit_attacking_team_enderman"`
				ArrowsHitKitAttackingTeamEnderman             float64  `json:"arrows_hit_kit_attacking_team_enderman"`
				ChestsOpenedKitAttackingTeamEnderman          float64  `json:"chests_opened_kit_attacking_team_enderman"`
				MostKillsGameKitAttackingTeamEnderman         float64  `json:"most_kills_game_kit_attacking_team_enderman"`
				FastestWinKitMegaMegaArmorsmith               float64  `json:"fastest_win_kit_mega_mega_armorsmith"`
				FastestWinMega                                float64  `json:"fastest_win_mega"`
				KillstreakKitMegaMegaArmorsmith               float64  `json:"killstreak_kit_mega_mega_armorsmith"`
				MobKillsMega                                  float64  `json:"mob_kills_mega"`
				MobKillsKitMegaMegaArmorsmith                 float64  `json:"mob_kills_kit_mega_mega_armorsmith"`
				WinstreakKitMegaMegaArmorsmith                float64  `json:"winstreak_kit_mega_mega_armorsmith"`
				KitBasicSoloEnergixInventory                  struct {
					POTION9 string `json:"POTION:9"`
				} `json:"kit_basic_solo_energix_inventory"`
				KitAdvancedSoloEndermanInventory struct {
					ENDERPEARL0 string `json:"ENDER_PEARL:0"`
				} `json:"kit_advanced_solo_enderman_inventory"`
				KitBasicSoloArmorsmithInventory struct {
					ENCHANTEDBOOK0 string `json:"ENCHANTED_BOOK:0"`
					EXPBOTTLE0     string `json:"EXP_BOTTLE:0"`
					ANVIL0         string `json:"ANVIL:0"`
					DIAMONDHELMET0 string `json:"DIAMOND_HELMET:0"`
				} `json:"kit_basic_solo_armorsmith_inventory"`
				KitBasicSoloPrincessInventory struct {
					ARROW0      string `json:"ARROW:0"`
					GOLDHELMET0 string `json:"GOLD_HELMET:0"`
					BOW114      string `json:"BOW:114"`
				} `json:"kit_basic_solo_princess_inventory"`
				MobKillsKitBasicSoloGrenade            float64 `json:"mob_kills_kit_basic_solo_grenade"`
				LongestBowShotKitBasicSoloSnowman      float64 `json:"longest_bow_shot_kit_basic_solo_snowman"`
				ArrowsShotKitBasicSoloSnowman          float64 `json:"arrows_shot_kit_basic_solo_snowman"`
				ArrowsHitKitBasicSoloSnowman           float64 `json:"arrows_hit_kit_basic_solo_snowman"`
				AssistsKitBasicSoloSnowman             float64 `json:"assists_kit_basic_solo_snowman"`
				MostKillsGameKitAdvancedSoloKnight     float64 `json:"most_kills_game_kit_advanced_solo_knight"`
				SurvivedPlayersKitAdvancedSoloKnight   float64 `json:"survived_players_kit_advanced_solo_knight"`
				TimePlayedKitAdvancedSoloKnight        float64 `json:"time_played_kit_advanced_solo_knight"`
				KillsKitAdvancedSoloKnight             float64 `json:"kills_kit_advanced_solo_knight"`
				DeathsKitAdvancedSoloKnight            float64 `json:"deaths_kit_advanced_solo_knight"`
				MeleeKillsKitAdvancedSoloKnight        float64 `json:"melee_kills_kit_advanced_solo_knight"`
				LossesKitAdvancedSoloKnight            float64 `json:"losses_kit_advanced_solo_knight"`
				FastestWinKitAdvancedSoloKnight        float64 `json:"fastest_win_kit_advanced_solo_knight"`
				VoidKillsKitAdvancedSoloKnight         float64 `json:"void_kills_kit_advanced_solo_knight"`
				GamesKitAdvancedSoloKnight             float64 `json:"games_kit_advanced_solo_knight"`
				WinstreakKitAdvancedSoloKnight         float64 `json:"winstreak_kit_advanced_solo_knight"`
				KillstreakKitAdvancedSoloKnight        float64 `json:"killstreak_kit_advanced_solo_knight"`
				WinsKitAdvancedSoloKnight              float64 `json:"wins_kit_advanced_solo_knight"`
				ChestsOpenedKitAdvancedSoloKnight      float64 `json:"chests_opened_kit_advanced_solo_knight"`
				LongestBowShotKitBasicSoloSpeleologist float64 `json:"longest_bow_shot_kit_basic_solo_speleologist"`
				MeleeKillsKitBasicSoloSpeleologist     float64 `json:"melee_kills_kit_basic_solo_speleologist"`
				ArrowsHitKitBasicSoloSpeleologist      float64 `json:"arrows_hit_kit_basic_solo_speleologist"`
				DeathsKitBasicSoloSpeleologist         float64 `json:"deaths_kit_basic_solo_speleologist"`
				ArrowsShotKitBasicSoloSpeleologist     float64 `json:"arrows_shot_kit_basic_solo_speleologist"`
				MostKillsGameKitBasicSoloSpeleologist  float64 `json:"most_kills_game_kit_basic_solo_speleologist"`
				TimePlayedKitBasicSoloSpeleologist     float64 `json:"time_played_kit_basic_solo_speleologist"`
				ChestsOpenedKitBasicSoloSpeleologist   float64 `json:"chests_opened_kit_basic_solo_speleologist"`
				LossesKitBasicSoloSpeleologist         float64 `json:"losses_kit_basic_solo_speleologist"`
				VoidKillsKitBasicSoloSpeleologist      float64 `json:"void_kills_kit_basic_solo_speleologist"`
				FastestWinKitBasicSoloSpeleologist     float64 `json:"fastest_win_kit_basic_solo_speleologist"`
				ChestsOpenedKitAdvancedSoloArmorer     float64 `json:"chests_opened_kit_advanced_solo_armorer"`
				MostKillsGameKitAdvancedSoloArmorer    float64 `json:"most_kills_game_kit_advanced_solo_armorer"`
				MeleeKillsKitAdvancedSoloArmorer       float64 `json:"melee_kills_kit_advanced_solo_armorer"`
				TimePlayedKitAdvancedSoloArmorer       float64 `json:"time_played_kit_advanced_solo_armorer"`
				FastestWinKitAdvancedSoloArmorer       float64 `json:"fastest_win_kit_advanced_solo_armorer"`
				LongestBowShotKitAdvancedSoloArmorer   float64 `json:"longest_bow_shot_kit_advanced_solo_armorer"`
				ArrowsHitKitAdvancedSoloArmorer        float64 `json:"arrows_hit_kit_advanced_solo_armorer"`
				ArrowsShotKitAdvancedSoloArmorer       float64 `json:"arrows_shot_kit_advanced_solo_armorer"`
				VoidKillsKitAdvancedSoloArmorer        float64 `json:"void_kills_kit_advanced_solo_armorer"`
				FastestWinKitRankedRankedScout         float64 `json:"fastest_win_kit_ranked_ranked_scout"`
				GamesKitRankedRankedScout              float64 `json:"games_kit_ranked_ranked_scout"`
				SurvivedPlayersKitRankedRankedScout    float64 `json:"survived_players_kit_ranked_ranked_scout"`
				TimePlayedKitRankedRankedScout         float64 `json:"time_played_kit_ranked_ranked_scout"`
				MostKillsGameKitRankedRankedScout      float64 `json:"most_kills_game_kit_ranked_ranked_scout"`
				ChestsOpenedKitRankedRankedScout       float64 `json:"chests_opened_kit_ranked_ranked_scout"`
				MeleeKillsKitRankedRankedScout         float64 `json:"melee_kills_kit_ranked_ranked_scout"`
				ArrowsShotKitRankedRankedScout         float64 `json:"arrows_shot_kit_ranked_ranked_scout"`
				WinstreakKitRankedRankedScout          float64 `json:"winstreak_kit_ranked_ranked_scout"`
				KillstreakKitRankedRankedScout         float64 `json:"killstreak_kit_ranked_ranked_scout"`
				KillsKitRankedRankedScout              float64 `json:"kills_kit_ranked_ranked_scout"`
				WinsKitRankedRankedScout               float64 `json:"wins_kit_ranked_ranked_scout"`
				VoidKillsKitRankedRankedScout          float64 `json:"void_kills_kit_ranked_ranked_scout"`
				LongestBowShotKitRankedRankedScout     float64 `json:"longest_bow_shot_kit_ranked_ranked_scout"`
				ArrowsHitKitRankedRankedScout          float64 `json:"arrows_hit_kit_ranked_ranked_scout"`
				DeathsKitRankedRankedScout             float64 `json:"deaths_kit_ranked_ranked_scout"`
				LossesKitRankedRankedScout             float64 `json:"losses_kit_ranked_ranked_scout"`
				AssistsKitRankedRankedScout            float64 `json:"assists_kit_ranked_ranked_scout"`
				KitRankedRankedScoutInventory          struct {
					DIAMONDAXE0     string `json:"DIAMOND_AXE:0"`
					POTION2         string `json:"POTION:2"`
					IRONLEGGINGS0   string `json:"IRON_LEGGINGS:0"`
					IRONHELMET0     string `json:"IRON_HELMET:0"`
					IRONBOOTS0      string `json:"IRON_BOOTS:0"`
					IRONCHESTPLATE0 string `json:"IRON_CHESTPLATE:0"`
					DIAMONDPICKAXE0 string `json:"DIAMOND_PICKAXE:0"`
				} `json:"kit_ranked_ranked_scout_inventory"`
				FastestWinKitRankedRankedChampion            float64 `json:"fastest_win_kit_ranked_ranked_champion"`
				ChestsOpenedKitRankedRankedChampion          float64 `json:"chests_opened_kit_ranked_ranked_champion"`
				VoidKillsKitRankedRankedChampion             float64 `json:"void_kills_kit_ranked_ranked_champion"`
				MostKillsGameKitRankedRankedChampion         float64 `json:"most_kills_game_kit_ranked_ranked_champion"`
				TimePlayedKitRankedRankedChampion            float64 `json:"time_played_kit_ranked_ranked_champion"`
				LongestBowShotKitRankedRankedChampion        float64 `json:"longest_bow_shot_kit_ranked_ranked_champion"`
				ArrowsShotKitRankedRankedChampion            float64 `json:"arrows_shot_kit_ranked_ranked_champion"`
				ArrowsHitKitRankedRankedChampion             float64 `json:"arrows_hit_kit_ranked_ranked_champion"`
				MeleeKillsKitRankedRankedChampion            float64 `json:"melee_kills_kit_ranked_ranked_champion"`
				FastestWinKitBlacksmithRankedBlacksmith      float64 `json:"fastest_win_kit_blacksmith_ranked_blacksmith"`
				MostKillsGameKitBlacksmithRankedBlacksmith   float64 `json:"most_kills_game_kit_blacksmith_ranked_blacksmith"`
				GamesKitBlacksmithRankedBlacksmith           float64 `json:"games_kit_blacksmith_ranked_blacksmith"`
				MeleeKillsKitBlacksmithRankedBlacksmith      float64 `json:"melee_kills_kit_blacksmith_ranked_blacksmith"`
				KillstreakKitBlacksmithRankedBlacksmith      float64 `json:"killstreak_kit_blacksmith_ranked_blacksmith"`
				WinsKitBlacksmithRankedBlacksmith            float64 `json:"wins_kit_blacksmith_ranked_blacksmith"`
				SurvivedPlayersKitBlacksmithRankedBlacksmith float64 `json:"survived_players_kit_blacksmith_ranked_blacksmith"`
				TimePlayedKitBlacksmithRankedBlacksmith      float64 `json:"time_played_kit_blacksmith_ranked_blacksmith"`
				WinstreakKitBlacksmithRankedBlacksmith       float64 `json:"winstreak_kit_blacksmith_ranked_blacksmith"`
				KillsKitBlacksmithRankedBlacksmith           float64 `json:"kills_kit_blacksmith_ranked_blacksmith"`
				ChestsOpenedKitBlacksmithRankedBlacksmith    float64 `json:"chests_opened_kit_blacksmith_ranked_blacksmith"`
				LossesKitBlacksmithRankedBlacksmith          float64 `json:"losses_kit_blacksmith_ranked_blacksmith"`
				DeathsKitBlacksmithRankedBlacksmith          float64 `json:"deaths_kit_blacksmith_ranked_blacksmith"`
				LongestBowShotKitBlacksmithRankedBlacksmith  float64 `json:"longest_bow_shot_kit_blacksmith_ranked_blacksmith"`
				VoidKillsKitBlacksmithRankedBlacksmith       float64 `json:"void_kills_kit_blacksmith_ranked_blacksmith"`
				ArrowsShotKitBlacksmithRankedBlacksmith      float64 `json:"arrows_shot_kit_blacksmith_ranked_blacksmith"`
				ArrowsHitKitBlacksmithRankedBlacksmith       float64 `json:"arrows_hit_kit_blacksmith_ranked_blacksmith"`
				TimePlayedKitBasicSoloTroll                  float64 `json:"time_played_kit_basic_solo_troll"`
				ChestsOpenedKitBasicSoloTroll                float64 `json:"chests_opened_kit_basic_solo_troll"`
				TimePlayedKitBasicSoloPharaoh                float64 `json:"time_played_kit_basic_solo_pharaoh"`
				ChestsOpenedKitBasicSoloPharaoh              float64 `json:"chests_opened_kit_basic_solo_pharaoh"`
				MostKillsGameKitBasicSoloPharaoh             float64 `json:"most_kills_game_kit_basic_solo_pharaoh"`
				VoidKillsKitBasicSoloPharaoh                 float64 `json:"void_kills_kit_basic_solo_pharaoh"`
				MeleeKillsKitBasicSoloPharaoh                float64 `json:"melee_kills_kit_basic_solo_pharaoh"`
				FastestWinKitBasicSoloPharaoh                float64 `json:"fastest_win_kit_basic_solo_pharaoh"`
				KitBasicSoloScoutInventory                   struct {
					POTION2 string `json:"POTION:2"`
				} `json:"kit_basic_solo_scout_inventory"`
				KitAttackingTeamScoutInventory struct {
					POTION2     string `json:"POTION:2"`
					STONESWORD0 string `json:"STONE_SWORD:0"`
				} `json:"kit_attacking_team_scout_inventory"`
				ChestsOpenedKitAdvancedSoloFarmer   float64 `json:"chests_opened_kit_advanced_solo_farmer"`
				MostKillsGameKitAdvancedSoloFarmer  float64 `json:"most_kills_game_kit_advanced_solo_farmer"`
				TimePlayedKitAdvancedSoloFarmer     float64 `json:"time_played_kit_advanced_solo_farmer"`
				VoidKillsKitAdvancedSoloFarmer      float64 `json:"void_kills_kit_advanced_solo_farmer"`
				FastestWinKitAdvancedSoloFarmer     float64 `json:"fastest_win_kit_advanced_solo_farmer"`
				ArrowsShotKitAdvancedSoloFarmer     float64 `json:"arrows_shot_kit_advanced_solo_farmer"`
				MeleeKillsKitAdvancedSoloFarmer     float64 `json:"melee_kills_kit_advanced_solo_farmer"`
				LongestBowShotKitAdvancedSoloFarmer float64 `json:"longest_bow_shot_kit_advanced_solo_farmer"`
				ArrowsHitKitAdvancedSoloFarmer      float64 `json:"arrows_hit_kit_advanced_solo_farmer"`
				KitAdvancedSoloFarmerInventory      struct {
					IRONLEGGINGS0 string `json:"IRON_LEGGINGS:0"`
					EGG0          string `json:"EGG:0"`
					GOLDENAPPLE0  string `json:"GOLDEN_APPLE:0"`
				} `json:"kit_advanced_solo_farmer_inventory"`
				VoidKillsKitBasicSoloRookie                 float64  `json:"void_kills_kit_basic_solo_rookie"`
				FastestWinKitBasicSoloRookie                float64  `json:"fastest_win_kit_basic_solo_rookie"`
				WinsKitBasicSoloRookie                      float64  `json:"wins_kit_basic_solo_rookie"`
				WinstreakKitBasicSoloRookie                 float64  `json:"winstreak_kit_basic_solo_rookie"`
				KillstreakKitBasicSoloRookie                float64  `json:"killstreak_kit_basic_solo_rookie"`
				LongestBowShotKitBasicSoloRookie            float64  `json:"longest_bow_shot_kit_basic_solo_rookie"`
				ArrowsShotKitBasicSoloRookie                float64  `json:"arrows_shot_kit_basic_solo_rookie"`
				ArrowsHitKitBasicSoloRookie                 float64  `json:"arrows_hit_kit_basic_solo_rookie"`
				AssistsKitBasicSoloRookie                   float64  `json:"assists_kit_basic_solo_rookie"`
				FastestWinKitAttackingTeamSnowman           float64  `json:"fastest_win_kit_attacking_team_snowman"`
				VoidKillsKitAttackingTeamSnowman            float64  `json:"void_kills_kit_attacking_team_snowman"`
				MeleeKillsKitAttackingTeamSnowman           float64  `json:"melee_kills_kit_attacking_team_snowman"`
				MostKillsGameKitAttackingTeamSnowman        float64  `json:"most_kills_game_kit_attacking_team_snowman"`
				WinstreakKitAttackingTeamSnowman            float64  `json:"winstreak_kit_attacking_team_snowman"`
				KillstreakKitAttackingTeamSnowman           float64  `json:"killstreak_kit_attacking_team_snowman"`
				TimePlayedKitAttackingTeamSnowman           float64  `json:"time_played_kit_attacking_team_snowman"`
				ChestsOpenedKitAttackingTeamSnowman         float64  `json:"chests_opened_kit_attacking_team_snowman"`
				TeamLuckyCharm                              float64  `json:"team_lucky_charm"`
				TeamEnvironmentalExpert                     float64  `json:"team_environmental_expert"`
				TeamBridger                                 float64  `json:"team_bridger"`
				FastestWinKitSupportingTeamArmorsmith       float64  `json:"fastest_win_kit_supporting_team_armorsmith"`
				KillstreakKitSupportingTeamArmorsmith       float64  `json:"killstreak_kit_supporting_team_armorsmith"`
				WinstreakKitSupportingTeamArmorsmith        float64  `json:"winstreak_kit_supporting_team_armorsmith"`
				MostKillsGameKitSupportingTeamArmorsmith    float64  `json:"most_kills_game_kit_supporting_team_armorsmith"`
				ChestsOpenedKitSupportingTeamArmorsmith     float64  `json:"chests_opened_kit_supporting_team_armorsmith"`
				TimePlayedKitSupportingTeamArmorsmith       float64  `json:"time_played_kit_supporting_team_armorsmith"`
				MeleeKillsKitSupportingTeamArmorsmith       float64  `json:"melee_kills_kit_supporting_team_armorsmith"`
				LongestBowShotKitSupportingTeamArmorsmith   float64  `json:"longest_bow_shot_kit_supporting_team_armorsmith"`
				VoidKillsKitSupportingTeamArmorsmith        float64  `json:"void_kills_kit_supporting_team_armorsmith"`
				ArrowsHitKitSupportingTeamArmorsmith        float64  `json:"arrows_hit_kit_supporting_team_armorsmith"`
				ArrowsShotKitSupportingTeamArmorsmith       float64  `json:"arrows_shot_kit_supporting_team_armorsmith"`
				MobKillsKitDefendingTeamArmorer             float64  `json:"mob_kills_kit_defending_team_armorer"`
				MobKillsTeam                                float64  `json:"mob_kills_team"`
				FastestWinKitAttackingTeamKnight            float64  `json:"fastest_win_kit_attacking_team_knight"`
				TimePlayedKitAttackingTeamKnight            float64  `json:"time_played_kit_attacking_team_knight"`
				VoidKillsKitAttackingTeamKnight             float64  `json:"void_kills_kit_attacking_team_knight"`
				MeleeKillsKitAttackingTeamKnight            float64  `json:"melee_kills_kit_attacking_team_knight"`
				KillstreakKitAttackingTeamKnight            float64  `json:"killstreak_kit_attacking_team_knight"`
				MostKillsGameKitAttackingTeamKnight         float64  `json:"most_kills_game_kit_attacking_team_knight"`
				WinstreakKitAttackingTeamKnight             float64  `json:"winstreak_kit_attacking_team_knight"`
				ArrowsShotKitAttackingTeamKnight            float64  `json:"arrows_shot_kit_attacking_team_knight"`
				ChestsOpenedKitAttackingTeamKnight          float64  `json:"chests_opened_kit_attacking_team_knight"`
				FastestWinKitMiningTeamCannoneer            float64  `json:"fastest_win_kit_mining_team_cannoneer"`
				MeleeKillsKitMiningTeamCannoneer            float64  `json:"melee_kills_kit_mining_team_cannoneer"`
				MostKillsGameKitMiningTeamCannoneer         float64  `json:"most_kills_game_kit_mining_team_cannoneer"`
				ChestsOpenedKitMiningTeamCannoneer          float64  `json:"chests_opened_kit_mining_team_cannoneer"`
				VoidKillsKitMiningTeamCannoneer             float64  `json:"void_kills_kit_mining_team_cannoneer"`
				TimePlayedKitMiningTeamCannoneer            float64  `json:"time_played_kit_mining_team_cannoneer"`
				DeathsKitAttackingTeamEnergix               float64  `json:"deaths_kit_attacking_team_energix"`
				LossesKitAttackingTeamEnergix               float64  `json:"losses_kit_attacking_team_energix"`
				SurvivedPlayersKitAttackingTeamEnergix      float64  `json:"survived_players_kit_attacking_team_energix"`
				ChestsOpenedKitAttackingTeamEnergix         float64  `json:"chests_opened_kit_attacking_team_energix"`
				TimePlayedKitAttackingTeamEnergix           float64  `json:"time_played_kit_attacking_team_energix"`
				FastestWinKitAdvancedSoloHunter             float64  `json:"fastest_win_kit_advanced_solo_hunter"`
				LongestBowShotKitAdvancedSoloHunter         float64  `json:"longest_bow_shot_kit_advanced_solo_hunter"`
				TimePlayedKitAdvancedSoloHunter             float64  `json:"time_played_kit_advanced_solo_hunter"`
				ChestsOpenedKitAdvancedSoloHunter           float64  `json:"chests_opened_kit_advanced_solo_hunter"`
				MeleeKillsKitAdvancedSoloHunter             float64  `json:"melee_kills_kit_advanced_solo_hunter"`
				AssistsKitAdvancedSoloHunter                float64  `json:"assists_kit_advanced_solo_hunter"`
				VoidKillsKitAdvancedSoloHunter              float64  `json:"void_kills_kit_advanced_solo_hunter"`
				ArrowsHitKitAdvancedSoloHunter              float64  `json:"arrows_hit_kit_advanced_solo_hunter"`
				MostKillsGameKitAdvancedSoloHunter          float64  `json:"most_kills_game_kit_advanced_solo_hunter"`
				ArrowsShotKitAdvancedSoloHunter             float64  `json:"arrows_shot_kit_advanced_solo_hunter"`
				FastestWinKitBasicSoloEcologist             float64  `json:"fastest_win_kit_basic_solo_ecologist"`
				VoidKillsKitBasicSoloEcologist              float64  `json:"void_kills_kit_basic_solo_ecologist"`
				MostKillsGameKitBasicSoloEcologist          float64  `json:"most_kills_game_kit_basic_solo_ecologist"`
				TimePlayedKitBasicSoloEcologist             float64  `json:"time_played_kit_basic_solo_ecologist"`
				ChestsOpenedKitBasicSoloEcologist           float64  `json:"chests_opened_kit_basic_solo_ecologist"`
				WinstreakKitBasicSoloEcologist              float64  `json:"winstreak_kit_basic_solo_ecologist"`
				KillstreakKitBasicSoloEcologist             float64  `json:"killstreak_kit_basic_solo_ecologist"`
				LongestBowShotKitAttackingTeamGrenade       float64  `json:"longest_bow_shot_kit_attacking_team_grenade"`
				FastestWinKitAttackingTeamGrenade           float64  `json:"fastest_win_kit_attacking_team_grenade"`
				KillsKitAttackingTeamGrenade                float64  `json:"kills_kit_attacking_team_grenade"`
				TimePlayedKitAttackingTeamGrenade           float64  `json:"time_played_kit_attacking_team_grenade"`
				WinstreakKitAttackingTeamGrenade            float64  `json:"winstreak_kit_attacking_team_grenade"`
				ArrowsShotKitAttackingTeamGrenade           float64  `json:"arrows_shot_kit_attacking_team_grenade"`
				GamesKitAttackingTeamGrenade                float64  `json:"games_kit_attacking_team_grenade"`
				ArrowsHitKitAttackingTeamGrenade            float64  `json:"arrows_hit_kit_attacking_team_grenade"`
				ChestsOpenedKitAttackingTeamGrenade         float64  `json:"chests_opened_kit_attacking_team_grenade"`
				AssistsKitAttackingTeamGrenade              float64  `json:"assists_kit_attacking_team_grenade"`
				WinsKitAttackingTeamGrenade                 float64  `json:"wins_kit_attacking_team_grenade"`
				MostKillsGameKitAttackingTeamGrenade        float64  `json:"most_kills_game_kit_attacking_team_grenade"`
				SurvivedPlayersKitAttackingTeamGrenade      float64  `json:"survived_players_kit_attacking_team_grenade"`
				KillstreakKitAttackingTeamGrenade           float64  `json:"killstreak_kit_attacking_team_grenade"`
				MeleeKillsKitAttackingTeamGrenade           float64  `json:"melee_kills_kit_attacking_team_grenade"`
				GamesKitAttackingTeamEnergix                float64  `json:"games_kit_attacking_team_energix"`
				TimePlayedKitBasicSoloFisherman             float64  `json:"time_played_kit_basic_solo_fisherman"`
				ChestsOpenedKitBasicSoloFisherman           float64  `json:"chests_opened_kit_basic_solo_fisherman"`
				MostKillsGameKitBasicSoloFisherman          float64  `json:"most_kills_game_kit_basic_solo_fisherman"`
				VoidKillsKitBasicSoloFisherman              float64  `json:"void_kills_kit_basic_solo_fisherman"`
				LongestBowShotKitAdvancedSoloKnight         float64  `json:"longest_bow_shot_kit_advanced_solo_knight"`
				ArrowsHitKitAdvancedSoloKnight              float64  `json:"arrows_hit_kit_advanced_solo_knight"`
				ArrowsShotKitAdvancedSoloKnight             float64  `json:"arrows_shot_kit_advanced_solo_knight"`
				LongestBowShotKitSupportingTeamRookie       float64  `json:"longest_bow_shot_kit_supporting_team_rookie"`
				FastestWinKitSupportingTeamRookie           float64  `json:"fastest_win_kit_supporting_team_rookie"`
				WinsKitSupportingTeamRookie                 float64  `json:"wins_kit_supporting_team_rookie"`
				ArrowsHitKitSupportingTeamRookie            float64  `json:"arrows_hit_kit_supporting_team_rookie"`
				MostKillsGameKitSupportingTeamRookie        float64  `json:"most_kills_game_kit_supporting_team_rookie"`
				GamesKitSupportingTeamRookie                float64  `json:"games_kit_supporting_team_rookie"`
				AssistsKitSupportingTeamRookie              float64  `json:"assists_kit_supporting_team_rookie"`
				MeleeKillsKitSupportingTeamRookie           float64  `json:"melee_kills_kit_supporting_team_rookie"`
				SurvivedPlayersKitSupportingTeamRookie      float64  `json:"survived_players_kit_supporting_team_rookie"`
				VoidKillsKitSupportingTeamRookie            float64  `json:"void_kills_kit_supporting_team_rookie"`
				WinstreakKitSupportingTeamRookie            float64  `json:"winstreak_kit_supporting_team_rookie"`
				ChestsOpenedKitSupportingTeamRookie         float64  `json:"chests_opened_kit_supporting_team_rookie"`
				TimePlayedKitSupportingTeamRookie           float64  `json:"time_played_kit_supporting_team_rookie"`
				KillstreakKitSupportingTeamRookie           float64  `json:"killstreak_kit_supporting_team_rookie"`
				KillsKitSupportingTeamRookie                float64  `json:"kills_kit_supporting_team_rookie"`
				ArrowsShotKitSupportingTeamRookie           float64  `json:"arrows_shot_kit_supporting_team_rookie"`
				FastestWinKitBasicSoloFisherman             float64  `json:"fastest_win_kit_basic_solo_fisherman"`
				MeleeKillsKitBasicSoloFisherman             float64  `json:"melee_kills_kit_basic_solo_fisherman"`
				MobKillsKitBasicSoloFisherman               float64  `json:"mob_kills_kit_basic_solo_fisherman"`
				LongestBowShotKitAttackingTeamKnight        float64  `json:"longest_bow_shot_kit_attacking_team_knight"`
				ArrowsHitKitAttackingTeamKnight             float64  `json:"arrows_hit_kit_attacking_team_knight"`
				MeleeKillsKitAttackingTeamEnergix           float64  `json:"melee_kills_kit_attacking_team_energix"`
				MostKillsGameKitAttackingTeamEnergix        float64  `json:"most_kills_game_kit_attacking_team_energix"`
				KillsKitAttackingTeamEnergix                float64  `json:"kills_kit_attacking_team_energix"`
				FastestWinKitBasicSoloScout                 float64  `json:"fastest_win_kit_basic_solo_scout"`
				KillsKitBasicSoloScout                      float64  `json:"kills_kit_basic_solo_scout"`
				MostKillsGameKitBasicSoloScout              float64  `json:"most_kills_game_kit_basic_solo_scout"`
				WinsKitBasicSoloScout                       float64  `json:"wins_kit_basic_solo_scout"`
				WinstreakKitBasicSoloScout                  float64  `json:"winstreak_kit_basic_solo_scout"`
				KillstreakKitBasicSoloScout                 float64  `json:"killstreak_kit_basic_solo_scout"`
				GamesKitBasicSoloScout                      float64  `json:"games_kit_basic_solo_scout"`
				VoidKillsKitBasicSoloScout                  float64  `json:"void_kills_kit_basic_solo_scout"`
				FastestWinKitMegaMegaPaladin                float64  `json:"fastest_win_kit_mega_mega_paladin"`
				TimePlayedKitMegaMegaPaladin                float64  `json:"time_played_kit_mega_mega_paladin"`
				MeleeKillsKitBasicSoloScout                 float64  `json:"melee_kills_kit_basic_solo_scout"`
				LongestBowShotKitEnderchestSoloEnderchest   float64  `json:"longest_bow_shot_kit_enderchest_solo_enderchest"`
				FastestWinKitEnderchestSoloEnderchest       float64  `json:"fastest_win_kit_enderchest_solo_enderchest"`
				MostKillsGameKitEnderchestSoloEnderchest    float64  `json:"most_kills_game_kit_enderchest_solo_enderchest"`
				ArrowsHitKitEnderchestSoloEnderchest        float64  `json:"arrows_hit_kit_enderchest_solo_enderchest"`
				MeleeKillsKitEnderchestSoloEnderchest       float64  `json:"melee_kills_kit_enderchest_solo_enderchest"`
				VoidKillsKitEnderchestSoloEnderchest        float64  `json:"void_kills_kit_enderchest_solo_enderchest"`
				ArrowsShotKitEnderchestSoloEnderchest       float64  `json:"arrows_shot_kit_enderchest_solo_enderchest"`
				LongestBowShotKitBasicSoloScout             float64  `json:"longest_bow_shot_kit_basic_solo_scout"`
				ArrowsShotKitBasicSoloScout                 float64  `json:"arrows_shot_kit_basic_solo_scout"`
				ArrowsHitKitBasicSoloScout                  float64  `json:"arrows_hit_kit_basic_solo_scout"`
				LongestBowShotKitAttackingTeamScout         float64  `json:"longest_bow_shot_kit_attacking_team_scout"`
				FastestWinKitAttackingTeamScout             float64  `json:"fastest_win_kit_attacking_team_scout"`
				ChestsOpenedKitAttackingTeamScout           float64  `json:"chests_opened_kit_attacking_team_scout"`
				MobKillsKitAttackingTeamScout               float64  `json:"mob_kills_kit_attacking_team_scout"`
				TimePlayedKitAttackingTeamScout             float64  `json:"time_played_kit_attacking_team_scout"`
				WinstreakKitAttackingTeamScout              float64  `json:"winstreak_kit_attacking_team_scout"`
				KillstreakKitAttackingTeamScout             float64  `json:"killstreak_kit_attacking_team_scout"`
				MostKillsGameKitAttackingTeamScout          float64  `json:"most_kills_game_kit_attacking_team_scout"`
				ArrowsShotKitAttackingTeamScout             float64  `json:"arrows_shot_kit_attacking_team_scout"`
				MeleeKillsKitAttackingTeamScout             float64  `json:"melee_kills_kit_attacking_team_scout"`
				VoidKillsKitAttackingTeamScout              float64  `json:"void_kills_kit_attacking_team_scout"`
				ArrowsHitKitAttackingTeamScout              float64  `json:"arrows_hit_kit_attacking_team_scout"`
				VoidKillsKitAttackingTeamEnderman           float64  `json:"void_kills_kit_attacking_team_enderman"`
				FloorIsLavaExplained                        float64  `json:"floor_is_lava_explained"`
				RushExplained                               float64  `json:"rush_explained"`
				BlizzardExplained                           float64  `json:"blizzard_explained"`
				ColorManiaExplained                         float64  `json:"color_mania_explained"`
				TntMadnessExplained                         float64  `json:"tnt_madness_explained"`
				SoloBlackMagic                              float64  `json:"solo_black_magic"`
				TntMadnessExplainedLast                     float64  `json:"tnt_madness_explained_last"`
				RushExplainedLast                           float64  `json:"rush_explained_last"`
				FloorIsLavaExplainedLast                    float64  `json:"floor_is_lava_explained_last"`
				BlizzardExplainedLast                       float64  `json:"blizzard_explained_last"`
				KillByColorExplainedLast                    float64  `json:"kill_by_color_explained_last"`
				KillByColorExplained                        float64  `json:"kill_by_color_explained"`
				WinStreakLab                                float64  `json:"win_streak_lab"`
				ChestsOpenedLabSolo                         float64  `json:"chests_opened_lab_solo"`
				DeathsLab                                   float64  `json:"deaths_lab"`
				DeathsLabSolo                               float64  `json:"deaths_lab_solo"`
				TimePlayedLabSolo                           float64  `json:"time_played_lab_solo"`
				DeathsLabKitAdvancedSoloEnderman            float64  `json:"deaths_lab_kit_advanced_solo_enderman"`
				TimePlayedLab                               float64  `json:"time_played_lab"`
				LossesLabKitAdvancedSoloEnderman            float64  `json:"losses_lab_kit_advanced_solo_enderman"`
				CoinsGainedLab                              float64  `json:"coins_gained_lab"`
				LossesLabSolo                               float64  `json:"losses_lab_solo"`
				ChestsOpenedLab                             float64  `json:"chests_opened_lab"`
				ChestsOpenedLabKitAdvancedSoloEnderman      float64  `json:"chests_opened_lab_kit_advanced_solo_enderman"`
				LossesLab                                   float64  `json:"losses_lab"`
				QuitsLab                                    float64  `json:"quits_lab"`
				TimePlayedLabKitAdvancedSoloEnderman        float64  `json:"time_played_lab_kit_advanced_solo_enderman"`
				SurvivedPlayersLabSolo                      float64  `json:"survived_players_lab_solo"`
				KillsLabKitAdvancedSoloEnderman             float64  `json:"kills_lab_kit_advanced_solo_enderman"`
				VoidKillsLab                                float64  `json:"void_kills_lab"`
				GamesLab                                    float64  `json:"games_lab"`
				SurvivedPlayersLab                          float64  `json:"survived_players_lab"`
				VoidKillsLabKitAdvancedSoloEnderman         float64  `json:"void_kills_lab_kit_advanced_solo_enderman"`
				ArrowsShotLabSolo                           float64  `json:"arrows_shot_lab_solo"`
				KillsLab                                    float64  `json:"kills_lab"`
				GamesLabSolo                                float64  `json:"games_lab_solo"`
				BlocksPlacedLab                             float64  `json:"blocks_placed_lab"`
				KillsLabSolo                                float64  `json:"kills_lab_solo"`
				GamesLabKitAdvancedSoloEnderman             float64  `json:"games_lab_kit_advanced_solo_enderman"`
				ArrowsShotLabKitAdvancedSoloEnderman        float64  `json:"arrows_shot_lab_kit_advanced_solo_enderman"`
				ArrowsHitLabKitAdvancedSoloEnderman         float64  `json:"arrows_hit_lab_kit_advanced_solo_enderman"`
				ArrowsHitLab                                float64  `json:"arrows_hit_lab"`
				VoidKillsLabSolo                            float64  `json:"void_kills_lab_solo"`
				ArrowsShotLab                               float64  `json:"arrows_shot_lab"`
				ArrowsHitLabSolo                            float64  `json:"arrows_hit_lab_solo"`
				SurvivedPlayersLabKitAdvancedSoloEnderman   float64  `json:"survived_players_lab_kit_advanced_solo_enderman"`
				AssistsLab                                  float64  `json:"assists_lab"`
				AssistsLabSolo                              float64  `json:"assists_lab_solo"`
				AssistsLabKitAdvancedSoloEnderman           float64  `json:"assists_lab_kit_advanced_solo_enderman"`
				KillstreakLab                               float64  `json:"killstreak_lab"`
				WinsLab                                     float64  `json:"wins_lab"`
				WinsLabKitAdvancedSoloEnderman              float64  `json:"wins_lab_kit_advanced_solo_enderman"`
				KillstreakLabKitAdvancedSoloEnderman        float64  `json:"killstreak_lab_kit_advanced_solo_enderman"`
				WinstreakLab                                float64  `json:"winstreak_lab"`
				WinstreakLabKitAdvancedSoloEnderman         float64  `json:"winstreak_lab_kit_advanced_solo_enderman"`
				WinstreakLabSolo                            float64  `json:"winstreak_lab_solo"`
				KillstreakLabSolo                           float64  `json:"killstreak_lab_solo"`
				WinsLabSolo                                 float64  `json:"wins_lab_solo"`
				EnderpearlsThrownLab                        float64  `json:"enderpearls_thrown_lab"`
				BlocksBrokenLab                             float64  `json:"blocks_broken_lab"`
				MeleeKillsLabKitAdvancedSoloEnderman        float64  `json:"melee_kills_lab_kit_advanced_solo_enderman"`
				MeleeKillsLabSolo                           float64  `json:"melee_kills_lab_solo"`
				MeleeKillsLab                               float64  `json:"melee_kills_lab"`
				MobKillsLabSolo                             float64  `json:"mob_kills_lab_solo"`
				EggThrownLab                                float64  `json:"egg_thrown_lab"`
				MobKillsLab                                 float64  `json:"mob_kills_lab"`
				MobKillsLabKitAdvancedSoloEnderman          float64  `json:"mob_kills_lab_kit_advanced_solo_enderman"`
				VoidKillsLabKitAttackingTeamDefault         float64  `json:"void_kills_lab_kit_attacking_team_default"`
				SurvivedPlayersLabTeam                      float64  `json:"survived_players_lab_team"`
				TimePlayedLabKitAttackingTeamDefault        float64  `json:"time_played_lab_kit_attacking_team_default"`
				SurvivedPlayersLabKitAttackingTeamDefault   float64  `json:"survived_players_lab_kit_attacking_team_default"`
				MostKillsGameLabKitAttackingTeamDefault     float64  `json:"most_kills_game_lab_kit_attacking_team_default"`
				KillsLabKitAttackingTeamDefault             float64  `json:"kills_lab_kit_attacking_team_default"`
				AssistsLabKitAttackingTeamDefault           float64  `json:"assists_lab_kit_attacking_team_default"`
				DeathsLabKitAttackingTeamDefault            float64  `json:"deaths_lab_kit_attacking_team_default"`
				LossesLabTeam                               float64  `json:"losses_lab_team"`
				MostKillsGameLabTeam                        float64  `json:"most_kills_game_lab_team"`
				KillsLabTeam                                float64  `json:"kills_lab_team"`
				DeathsLabTeam                               float64  `json:"deaths_lab_team"`
				ChestsOpenedLabKitAttackingTeamDefault      float64  `json:"chests_opened_lab_kit_attacking_team_default"`
				ChestsOpenedLabTeam                         float64  `json:"chests_opened_lab_team"`
				TimePlayedLabTeam                           float64  `json:"time_played_lab_team"`
				AssistsLabTeam                              float64  `json:"assists_lab_team"`
				MostKillsGameLab                            float64  `json:"most_kills_game_lab"`
				LossesLabKitAttackingTeamDefault            float64  `json:"losses_lab_kit_attacking_team_default"`
				VoidKillsLabTeam                            float64  `json:"void_kills_lab_team"`
				FastestWinLabKitAttackingTeamDefault        float64  `json:"fastest_win_lab_kit_attacking_team_default"`
				LongestBowShotLab                           float64  `json:"longest_bow_shot_lab"`
				LongestBowShotLabKitAttackingTeamDefault    float64  `json:"longest_bow_shot_lab_kit_attacking_team_default"`
				FastestWinLabTeam                           float64  `json:"fastest_win_lab_team"`
				FastestWinLab                               float64  `json:"fastest_win_lab"`
				LongestBowShotLabTeam                       float64  `json:"longest_bow_shot_lab_team"`
				ArrowsHitLabTeam                            float64  `json:"arrows_hit_lab_team"`
				KillstreakLabKitAttackingTeamDefault        float64  `json:"killstreak_lab_kit_attacking_team_default"`
				ArrowsShotLabTeam                           float64  `json:"arrows_shot_lab_team"`
				WinstreakLabKitAttackingTeamDefault         float64  `json:"winstreak_lab_kit_attacking_team_default"`
				WinsLabKitAttackingTeamDefault              float64  `json:"wins_lab_kit_attacking_team_default"`
				MeleeKillsLabTeam                           float64  `json:"melee_kills_lab_team"`
				ArrowsShotLabKitAttackingTeamDefault        float64  `json:"arrows_shot_lab_kit_attacking_team_default"`
				WinstreakLabTeam                            float64  `json:"winstreak_lab_team"`
				GamesLabKitAttackingTeamDefault             float64  `json:"games_lab_kit_attacking_team_default"`
				GamesLabTeam                                float64  `json:"games_lab_team"`
				KillstreakLabTeam                           float64  `json:"killstreak_lab_team"`
				WinsLabTeam                                 float64  `json:"wins_lab_team"`
				ArrowsHitLabKitAttackingTeamDefault         float64  `json:"arrows_hit_lab_kit_attacking_team_default"`
				MeleeKillsLabKitAttackingTeamDefault        float64  `json:"melee_kills_lab_kit_attacking_team_default"`
				VoidKillsKitMegaMegaHellhound               float64  `json:"void_kills_kit_mega_mega_hellhound"`
				ItemsEnchantedLab                           float64  `json:"items_enchanted_lab"`
				SurvivedPlayersKitAdvancedSoloPigRider      float64  `json:"survived_players_kit_advanced_solo_pig-rider"`
				ChestsOpenedKitAdvancedSoloPigRider         float64  `json:"chests_opened_kit_advanced_solo_pig-rider"`
				LossesKitAdvancedSoloPigRider               float64  `json:"losses_kit_advanced_solo_pig-rider"`
				TimePlayedKitAdvancedSoloPigRider           float64  `json:"time_played_kit_advanced_solo_pig-rider"`
				DeathsKitAdvancedSoloPigRider               float64  `json:"deaths_kit_advanced_solo_pig-rider"`
				ActiveKitTEAMS                              string   `json:"activeKit_TEAMS"`
				SlimeExplained                              float64  `json:"slime_explained"`
				SlimeExplainedLast                          float64  `json:"slime_explained_last"`
				SoloNecromancer                             float64  `json:"solo_necromancer"`
				LongestBowShotKitAdvancedSoloPigRider       float64  `json:"longest_bow_shot_kit_advanced_solo_pig-rider"`
				ArrowsHitKitAdvancedSoloPigRider            float64  `json:"arrows_hit_kit_advanced_solo_pig-rider"`
				VoidKillsKitAdvancedSoloPigRider            float64  `json:"void_kills_kit_advanced_solo_pig-rider"`
				MostKillsGameKitAdvancedSoloPigRider        float64  `json:"most_kills_game_kit_advanced_solo_pig-rider"`
				ArrowsShotKitAdvancedSoloPigRider           float64  `json:"arrows_shot_kit_advanced_solo_pig-rider"`
				KillsKitAdvancedSoloPigRider                float64  `json:"kills_kit_advanced_solo_pig-rider"`
				MeleeKillsKitAdvancedSoloPigRider           float64  `json:"melee_kills_kit_advanced_solo_pig-rider"`
				GamesKitAdvancedSoloPigRider                float64  `json:"games_kit_advanced_solo_pig-rider"`
				FastestWinKitAdvancedSoloPigRider           float64  `json:"fastest_win_kit_advanced_solo_pig-rider"`
				WinsKitAdvancedSoloPigRider                 float64  `json:"wins_kit_advanced_solo_pig-rider"`
				WinstreakKitAdvancedSoloPigRider            float64  `json:"winstreak_kit_advanced_solo_pig-rider"`
				KillstreakKitAdvancedSoloPigRider           float64  `json:"killstreak_kit_advanced_solo_pig-rider"`
				SurvivedPlayersLabKitAdvancedSoloPigRider   float64  `json:"survived_players_lab_kit_advanced_solo_pig-rider"`
				TimePlayedLabKitAdvancedSoloPigRider        float64  `json:"time_played_lab_kit_advanced_solo_pig-rider"`
				ChestsOpenedLabKitAdvancedSoloPigRider      float64  `json:"chests_opened_lab_kit_advanced_solo_pig-rider"`
				LossesLabKitAdvancedSoloPigRider            float64  `json:"losses_lab_kit_advanced_solo_pig-rider"`
				DeathsLabKitAdvancedSoloPigRider            float64  `json:"deaths_lab_kit_advanced_solo_pig-rider"`
				KillsLabKitAdvancedSoloPigRider             float64  `json:"kills_lab_kit_advanced_solo_pig-rider"`
				MeleeKillsLabKitAdvancedSoloPigRider        float64  `json:"melee_kills_lab_kit_advanced_solo_pig-rider"`
				FastestWinLabSolo                           float64  `json:"fastest_win_lab_solo"`
				FastestWinLabKitAdvancedSoloPigRider        float64  `json:"fastest_win_lab_kit_advanced_solo_pig-rider"`
				WinsLabKitAdvancedSoloPigRider              float64  `json:"wins_lab_kit_advanced_solo_pig-rider"`
				LabWinSlimeLab                              float64  `json:"lab_win_slime_lab"`
				WinstreakLabKitAdvancedSoloPigRider         float64  `json:"winstreak_lab_kit_advanced_solo_pig-rider"`
				ArrowsHitLabKitAdvancedSoloPigRider         float64  `json:"arrows_hit_lab_kit_advanced_solo_pig-rider"`
				LabWinSlimeLabSolo                          float64  `json:"lab_win_slime_lab_solo"`
				GamesLabKitAdvancedSoloPigRider             float64  `json:"games_lab_kit_advanced_solo_pig-rider"`
				LabWinSlimeLabKitAdvancedSoloPigRider       float64  `json:"lab_win_slime_lab_kit_advanced_solo_pig-rider"`
				VoidKillsLabKitAdvancedSoloPigRider         float64  `json:"void_kills_lab_kit_advanced_solo_pig-rider"`
				ArrowsShotLabKitAdvancedSoloPigRider        float64  `json:"arrows_shot_lab_kit_advanced_solo_pig-rider"`
				KillstreakLabKitAdvancedSoloPigRider        float64  `json:"killstreak_lab_kit_advanced_solo_pig-rider"`
				LabWinTntMadnessLabSolo                     float64  `json:"lab_win_tnt_madness_lab_solo"`
				LabWinTntMadnessLabKitAdvancedSoloPigRider  float64  `json:"lab_win_tnt_madness_lab_kit_advanced_solo_pig-rider"`
				LabWinTntMadnessLab                         float64  `json:"lab_win_tnt_madness_lab"`
				LongestBowShotLabKitAdvancedSoloPigRider    float64  `json:"longest_bow_shot_lab_kit_advanced_solo_pig-rider"`
				LongestBowShotLabSolo                       float64  `json:"longest_bow_shot_lab_solo"`
				AssistsLabKitAdvancedSoloPigRider           float64  `json:"assists_lab_kit_advanced_solo_pig-rider"`
				LabWinTntMadnessLabKitAdvancedSoloEnderman  float64  `json:"lab_win_tnt_madness_lab_kit_advanced_solo_enderman"`
				ChestsOpenedKitAdvancedSoloSalmon           float64  `json:"chests_opened_kit_advanced_solo_salmon"`
				TimePlayedKitAdvancedSoloSalmon             float64  `json:"time_played_kit_advanced_solo_salmon"`
				LossesKitAdvancedSoloSalmon                 float64  `json:"losses_kit_advanced_solo_salmon"`
				DeathsKitAdvancedSoloSalmon                 float64  `json:"deaths_kit_advanced_solo_salmon"`
				FastestWinKitAdvancedSoloSalmon             float64  `json:"fastest_win_kit_advanced_solo_salmon"`
				KillstreakKitAdvancedSoloSalmon             float64  `json:"killstreak_kit_advanced_solo_salmon"`
				KillsKitAdvancedSoloSalmon                  float64  `json:"kills_kit_advanced_solo_salmon"`
				MeleeKillsKitAdvancedSoloSalmon             float64  `json:"melee_kills_kit_advanced_solo_salmon"`
				WinstreakKitAdvancedSoloSalmon              float64  `json:"winstreak_kit_advanced_solo_salmon"`
				SurvivedPlayersKitAdvancedSoloSalmon        float64  `json:"survived_players_kit_advanced_solo_salmon"`
				MostKillsGameKitAdvancedSoloSalmon          float64  `json:"most_kills_game_kit_advanced_solo_salmon"`
				WinsKitAdvancedSoloSalmon                   float64  `json:"wins_kit_advanced_solo_salmon"`
				GamesKitAdvancedSoloSalmon                  float64  `json:"games_kit_advanced_solo_salmon"`
				ArrowsShotKitAdvancedSoloSalmon             float64  `json:"arrows_shot_kit_advanced_solo_salmon"`
				VoidKillsKitAdvancedSoloSalmon              float64  `json:"void_kills_kit_advanced_solo_salmon"`
				AssistsKitAdvancedSoloSalmon                float64  `json:"assists_kit_advanced_solo_salmon"`
				LongestBowShotKitAdvancedSoloSalmon         float64  `json:"longest_bow_shot_kit_advanced_solo_salmon"`
				ArrowsHitKitAdvancedSoloSalmon              float64  `json:"arrows_hit_kit_advanced_solo_salmon"`
				MobKillsKitAdvancedSoloEnderman             float64  `json:"mob_kills_kit_advanced_solo_enderman"`
				MegaBlackMagic                              float64  `json:"mega_black_magic"`
				MegaNecromancer                             float64  `json:"mega_necromancer"`
				TeamNecromancer                             float64  `json:"team_necromancer"`
				TeamBlackMagic                              float64  `json:"team_black_magic"`
				MegaEnvironmentalExpert                     float64  `json:"mega_environmental_expert"`
				KitMegaMegaPyromaniac                       float64  `json:"kit_mega_mega_pyromaniac"`
				KitMegaMegaFisherman                        float64  `json:"kit_mega_mega_fisherman"`
				LongestBowKillKitAdvancedSoloEnderman       float64  `json:"longest_bow_kill_kit_advanced_solo_enderman"`
				LongestBowKill                              float64  `json:"longest_bow_kill"`
				LongestBowKillSolo                          float64  `json:"longest_bow_kill_solo"`
				InGamePresentsCap201719                     float64  `json:"inGamePresentsCap_2017_19"`
				InGamePresentsCap201722                     float64  `json:"inGamePresentsCap_2017_22"`
				LongestBowKillKitDefendingTeamArmorer       float64  `json:"longest_bow_kill_kit_defending_team_armorer"`
				LongestBowKillTeam                          float64  `json:"longest_bow_kill_team"`
				SkywarsChests                               float64  `json:"skywars_chests"`
				SkyWarsOpenedChests                         float64  `json:"SkyWars_openedChests"`
				SkywarsChestHistory                         []string `json:"skywars_chest_history"`
				SkyWarsOpenedCommons                        float64  `json:"SkyWars_openedCommons"`
				SkyWarsOpenedRares                          float64  `json:"SkyWars_openedRares"`
				CosmeticTokens                              float64  `json:"cosmetic_tokens"`
				SkyWarsOpenedEpics                          float64  `json:"SkyWars_openedEpics"`
				SkyWarsOpenedLegendaries                    float64  `json:"SkyWars_openedLegendaries"`
				ActiveBalloon                               string   `json:"active_balloon"`
				ActiveSprays                                string   `json:"active_sprays"`
				ActiveVictorydance                          string   `json:"active_victorydance"`
				LossesKitAdvancedSoloSlime                  float64  `json:"losses_kit_advanced_solo_slime"`
				GamesKitAdvancedSoloSlime                   float64  `json:"games_kit_advanced_solo_slime"`
				TimePlayedKitAdvancedSoloSlime              float64  `json:"time_played_kit_advanced_solo_slime"`
				DeathsKitAdvancedSoloSlime                  float64  `json:"deaths_kit_advanced_solo_slime"`
				ChestsOpenedKitAdvancedSoloSlime            float64  `json:"chests_opened_kit_advanced_solo_slime"`
				SurvivedPlayersKitAdvancedSoloSlime         float64  `json:"survived_players_kit_advanced_solo_slime"`
				ActiveProjectiletrail                       string   `json:"active_projectiletrail"`
				ActiveKilleffect                            string   `json:"active_killeffect"`
				ActiveDeathcry                              string   `json:"active_deathcry"`
				LuckyExplained                              float64  `json:"lucky_explained"`
				LuckyExplainedLast                          float64  `json:"lucky_explained_last"`
				TimePlayedLabKitAdvancedSoloSlime           float64  `json:"time_played_lab_kit_advanced_solo_slime"`
				SurvivedPlayersLabKitAdvancedSoloSlime      float64  `json:"survived_players_lab_kit_advanced_solo_slime"`
				MostKillsGameLabSolo                        float64  `json:"most_kills_game_lab_solo"`
				LossesLabKitAdvancedSoloSlime               float64  `json:"losses_lab_kit_advanced_solo_slime"`
				DeathsLabKitAdvancedSoloSlime               float64  `json:"deaths_lab_kit_advanced_solo_slime"`
				ArrowsShotLabKitAdvancedSoloSlime           float64  `json:"arrows_shot_lab_kit_advanced_solo_slime"`
				KillsLabKitAdvancedSoloSlime                float64  `json:"kills_lab_kit_advanced_solo_slime"`
				MostKillsGameLabKitAdvancedSoloSlime        float64  `json:"most_kills_game_lab_kit_advanced_solo_slime"`
				ChestsOpenedLabKitAdvancedSoloSlime         float64  `json:"chests_opened_lab_kit_advanced_solo_slime"`
				LongestBowKillLabSolo                       float64  `json:"longest_bow_kill_lab_solo"`
				FastestWinLabKitAdvancedSoloSlime           float64  `json:"fastest_win_lab_kit_advanced_solo_slime"`
				LongestBowKillLab                           float64  `json:"longest_bow_kill_lab"`
				LongestBowKillLabKitAdvancedSoloSlime       float64  `json:"longest_bow_kill_lab_kit_advanced_solo_slime"`
				LongestBowShotLabKitAdvancedSoloSlime       float64  `json:"longest_bow_shot_lab_kit_advanced_solo_slime"`
				WinsLabKitAdvancedSoloSlime                 float64  `json:"wins_lab_kit_advanced_solo_slime"`
				LabWinLuckyBlocksLabKitAdvancedSoloSlime    float64  `json:"lab_win_lucky_blocks_lab_kit_advanced_solo_slime"`
				VoidKillsLabKitAdvancedSoloSlime            float64  `json:"void_kills_lab_kit_advanced_solo_slime"`
				WinstreakLabKitAdvancedSoloSlime            float64  `json:"winstreak_lab_kit_advanced_solo_slime"`
				LabWinLuckyBlocksLabSolo                    float64  `json:"lab_win_lucky_blocks_lab_solo"`
				LabWinLuckyBlocksLab                        float64  `json:"lab_win_lucky_blocks_lab"`
				KillstreakLabKitAdvancedSoloSlime           float64  `json:"killstreak_lab_kit_advanced_solo_slime"`
				GamesLabKitAdvancedSoloSlime                float64  `json:"games_lab_kit_advanced_solo_slime"`
				MeleeKillsLabKitAdvancedSoloSlime           float64  `json:"melee_kills_lab_kit_advanced_solo_slime"`
				ArrowsHitLabKitAdvancedSoloSlime            float64  `json:"arrows_hit_lab_kit_advanced_solo_slime"`
				MobKillsLabKitAdvancedSoloSlime             float64  `json:"mob_kills_lab_kit_advanced_solo_slime"`
				AssistsLabKitAdvancedSoloSlime              float64  `json:"assists_lab_kit_advanced_solo_slime"`
				FreeLootChestNpc                            float64  `json:"freeLootChestNpc"`
				LongestBowKillKitBasicSoloScout             float64  `json:"longest_bow_kill_kit_basic_solo_scout"`
				LongestBowKillKitAttackingTeamScout         float64  `json:"longest_bow_kill_kit_attacking_team_scout"`
				TimePlayedLabKitBasicSoloScout              float64  `json:"time_played_lab_kit_basic_solo_scout"`
				SurvivedPlayersLabKitBasicSoloScout         float64  `json:"survived_players_lab_kit_basic_solo_scout"`
				ChestsOpenedLabKitBasicSoloScout            float64  `json:"chests_opened_lab_kit_basic_solo_scout"`
				DeathsLabKitBasicSoloScout                  float64  `json:"deaths_lab_kit_basic_solo_scout"`
				LossesLabKitBasicSoloScout                  float64  `json:"losses_lab_kit_basic_solo_scout"`
				LongestBowKillLabKitBasicSoloScout          float64  `json:"longest_bow_kill_lab_kit_basic_solo_scout"`
				WinstreakLabKitBasicSoloScout               float64  `json:"winstreak_lab_kit_basic_solo_scout"`
				ArrowsHitLabKitBasicSoloScout               float64  `json:"arrows_hit_lab_kit_basic_solo_scout"`
				KillsLabKitBasicSoloScout                   float64  `json:"kills_lab_kit_basic_solo_scout"`
				VoidKillsLabKitBasicSoloScout               float64  `json:"void_kills_lab_kit_basic_solo_scout"`
				GamesLabKitBasicSoloScout                   float64  `json:"games_lab_kit_basic_solo_scout"`
				LabWinLuckyBlocksLabKitBasicSoloScout       float64  `json:"lab_win_lucky_blocks_lab_kit_basic_solo_scout"`
				ArrowsShotLabKitBasicSoloScout              float64  `json:"arrows_shot_lab_kit_basic_solo_scout"`
				WinsLabKitBasicSoloScout                    float64  `json:"wins_lab_kit_basic_solo_scout"`
				MeleeKillsLabKitBasicSoloScout              float64  `json:"melee_kills_lab_kit_basic_solo_scout"`
				KillstreakLabKitBasicSoloScout              float64  `json:"killstreak_lab_kit_basic_solo_scout"`
				ActiveCage                                  string   `json:"active_cage"`
				VoidKillsLabKitAttackingTeamScout           float64  `json:"void_kills_lab_kit_attacking_team_scout"`
				DeathsLabKitAttackingTeamScout              float64  `json:"deaths_lab_kit_attacking_team_scout"`
				KillsLabKitAttackingTeamScout               float64  `json:"kills_lab_kit_attacking_team_scout"`
				SurvivedPlayersLabKitAttackingTeamScout     float64  `json:"survived_players_lab_kit_attacking_team_scout"`
				ChestsOpenedLabKitAttackingTeamScout        float64  `json:"chests_opened_lab_kit_attacking_team_scout"`
				MeleeKillsLabKitAttackingTeamScout          float64  `json:"melee_kills_lab_kit_attacking_team_scout"`
				LossesLabKitAttackingTeamScout              float64  `json:"losses_lab_kit_attacking_team_scout"`
				TimePlayedLabKitAttackingTeamScout          float64  `json:"time_played_lab_kit_attacking_team_scout"`
				ArrowsHitLabKitAttackingTeamScout           float64  `json:"arrows_hit_lab_kit_attacking_team_scout"`
				ArrowsShotLabKitAttackingTeamScout          float64  `json:"arrows_shot_lab_kit_attacking_team_scout"`
				LabWinTntMadnessLabTeam                     float64  `json:"lab_win_tnt_madness_lab_team"`
				GamesLabKitAttackingTeamScout               float64  `json:"games_lab_kit_attacking_team_scout"`
				WinstreakLabKitAttackingTeamScout           float64  `json:"winstreak_lab_kit_attacking_team_scout"`
				LabWinTntMadnessLabKitAttackingTeamScout    float64  `json:"lab_win_tnt_madness_lab_kit_attacking_team_scout"`
				WinsLabKitAttackingTeamScout                float64  `json:"wins_lab_kit_attacking_team_scout"`
				KillstreakLabKitAttackingTeamScout          float64  `json:"killstreak_lab_kit_attacking_team_scout"`
				MobKillsLabTeam                             float64  `json:"mob_kills_lab_team"`
				AssistsLabKitAttackingTeamScout             float64  `json:"assists_lab_kit_attacking_team_scout"`
				MobKillsLabKitAttackingTeamScout            float64  `json:"mob_kills_lab_kit_attacking_team_scout"`
				LongestBowKillKitMegaMegaArmorsmith         float64  `json:"longest_bow_kill_kit_mega_mega_armorsmith"`
				LongestBowKillMega                          float64  `json:"longest_bow_kill_mega"`
				LongestBowShotLabKitAttackingTeamScout      float64  `json:"longest_bow_shot_lab_kit_attacking_team_scout"`
				LabWinLuckyBlocksLabKitAdvancedSoloEnderman float64  `json:"lab_win_lucky_blocks_lab_kit_advanced_solo_enderman"`
				BowKills                                    float64  `json:"bow_kills"`
				BowKillsSolo                                float64  `json:"bow_kills_solo"`
				BowKillsKitAdvancedSoloEnderman             float64  `json:"bow_kills_kit_advanced_solo_enderman"`
				LabWinSlimeLabKitAdvancedSoloEnderman       float64  `json:"lab_win_slime_lab_kit_advanced_solo_enderman"`
				SkywarsEasterBoxes                          float64  `json:"skywars_easter_boxes"`
				ChestHistoryNew                             []string `json:"chest_history_new"`
				SoloRobbery                                 float64  `json:"solo_robbery"`
				ActiveKitSOLORandom                         bool     `json:"activeKit_SOLO_random"`
				KitAdvancedSoloJesterInventory              struct {
					POTION1    string `json:"POTION:1"`
					WOODSWORD0 string `json:"WOOD_SWORD:0"`
				} `json:"kit_advanced_solo_jester_inventory"`
				LongestBowKillKitAdvancedSoloJester                       float64 `json:"longest_bow_kill_kit_advanced_solo_jester"`
				FastestWinKitAdvancedSoloJester                           float64 `json:"fastest_win_kit_advanced_solo_jester"`
				MostKillsGameKitAdvancedSoloJester                        float64 `json:"most_kills_game_kit_advanced_solo_jester"`
				LongestBowShotKitAdvancedSoloJester                       float64 `json:"longest_bow_shot_kit_advanced_solo_jester"`
				WinstreakKitAdvancedSoloJester                            float64 `json:"winstreak_kit_advanced_solo_jester"`
				AssistsKitAdvancedSoloJester                              float64 `json:"assists_kit_advanced_solo_jester"`
				KillstreakKitAdvancedSoloJester                           float64 `json:"killstreak_kit_advanced_solo_jester"`
				KillsKitAdvancedSoloJester                                float64 `json:"kills_kit_advanced_solo_jester"`
				SurvivedPlayersKitAdvancedSoloJester                      float64 `json:"survived_players_kit_advanced_solo_jester"`
				MeleeKillsKitAdvancedSoloJester                           float64 `json:"melee_kills_kit_advanced_solo_jester"`
				TimePlayedKitAdvancedSoloJester                           float64 `json:"time_played_kit_advanced_solo_jester"`
				ArrowsShotKitAdvancedSoloJester                           float64 `json:"arrows_shot_kit_advanced_solo_jester"`
				GamesKitAdvancedSoloJester                                float64 `json:"games_kit_advanced_solo_jester"`
				WinsKitAdvancedSoloJester                                 float64 `json:"wins_kit_advanced_solo_jester"`
				VoidKillsKitAdvancedSoloJester                            float64 `json:"void_kills_kit_advanced_solo_jester"`
				ChestsOpenedKitAdvancedSoloJester                         float64 `json:"chests_opened_kit_advanced_solo_jester"`
				ArrowsHitKitAdvancedSoloJester                            float64 `json:"arrows_hit_kit_advanced_solo_jester"`
				LongestBowKillKitBasicSoloArmorsmith                      float64 `json:"longest_bow_kill_kit_basic_solo_armorsmith"`
				BowKillsKitBasicSoloArmorsmith                            float64 `json:"bow_kills_kit_basic_solo_armorsmith"`
				MobKillsKitBasicSoloArmorsmith                            float64 `json:"mob_kills_kit_basic_solo_armorsmith"`
				ActiveKitTEAMSRandom                                      bool    `json:"activeKit_TEAMS_random"`
				BowKillsTeam                                              float64 `json:"bow_kills_team"`
				BowKillsKitDefendingTeamArmorer                           float64 `json:"bow_kills_kit_defending_team_armorer"`
				HuntersVsBeastsExplained                                  float64 `json:"hunters_vs_beasts_explained"`
				HuntersVsBeastsExplainedLast                              float64 `json:"hunters_vs_beasts_explained_last"`
				BeastChance                                               float64 `json:"beast_chance"`
				TimePlayedLabKitAdvancedSoloCannoneer                     float64 `json:"time_played_lab_kit_advanced_solo_cannoneer"`
				KillsLabKitAdvancedSoloCannoneer                          float64 `json:"kills_lab_kit_advanced_solo_cannoneer"`
				LossesLabKitAdvancedSoloCannoneer                         float64 `json:"losses_lab_kit_advanced_solo_cannoneer"`
				GamesLabKitAdvancedSoloCannoneer                          float64 `json:"games_lab_kit_advanced_solo_cannoneer"`
				DeathsLabKitAdvancedSoloCannoneer                         float64 `json:"deaths_lab_kit_advanced_solo_cannoneer"`
				ChestsOpenedLabKitAdvancedSoloCannoneer                   float64 `json:"chests_opened_lab_kit_advanced_solo_cannoneer"`
				SurvivedPlayersLabKitAdvancedSoloCannoneer                float64 `json:"survived_players_lab_kit_advanced_solo_cannoneer"`
				FastestWinLabKitAdvancedSoloCannoneer                     float64 `json:"fastest_win_lab_kit_advanced_solo_cannoneer"`
				LabHvbHunterWinsLab                                       float64 `json:"lab_hvb_hunter_wins_lab"`
				WinsLabKitAdvancedSoloCannoneer                           float64 `json:"wins_lab_kit_advanced_solo_cannoneer"`
				LabWinHuntersVsBeastsLabSolo                              float64 `json:"lab_win_hunters_vs_beasts_lab_solo"`
				LabWinHuntersVsBeastsLab                                  float64 `json:"lab_win_hunters_vs_beasts_lab"`
				LabWinHuntersVsBeastsLabKitAdvancedSoloCannoneer          float64 `json:"lab_win_hunters_vs_beasts_lab_kit_advanced_solo_cannoneer"`
				WinstreakLabKitAdvancedSoloCannoneer                      float64 `json:"winstreak_lab_kit_advanced_solo_cannoneer"`
				LongestBowKillLabKitAdvancedSoloCannoneer                 float64 `json:"longest_bow_kill_lab_kit_advanced_solo_cannoneer"`
				MostKillsGameLabKitAdvancedSoloCannoneer                  float64 `json:"most_kills_game_lab_kit_advanced_solo_cannoneer"`
				KillstreakLabKitAdvancedSoloCannoneer                     float64 `json:"killstreak_lab_kit_advanced_solo_cannoneer"`
				MeleeKillsLabKitAdvancedSoloCannoneer                     float64 `json:"melee_kills_lab_kit_advanced_solo_cannoneer"`
				LabHvbBeastWinsLab                                        float64 `json:"lab_hvb_beast_wins_lab"`
				MobKillsLabKitAdvancedSoloCannoneer                       float64 `json:"mob_kills_lab_kit_advanced_solo_cannoneer"`
				VoidKillsLabKitAdvancedSoloCannoneer                      float64 `json:"void_kills_lab_kit_advanced_solo_cannoneer"`
				InGamePresentsCap20186                                    float64 `json:"inGamePresentsCap_2018_6"`
				SkywarsExperience                                         float64 `json:"skywars_experience"`
				TeamBulldozer                                             float64 `json:"team_bulldozer"`
				TeamRobbery                                               float64 `json:"team_robbery"`
				FastestWinKitAttackingTeamFisherman                       float64 `json:"fastest_win_kit_attacking_team_fisherman"`
				MostKillsGameKitAttackingTeamFisherman                    float64 `json:"most_kills_game_kit_attacking_team_fisherman"`
				KillstreakKitAttackingTeamFisherman                       float64 `json:"killstreak_kit_attacking_team_fisherman"`
				KillsKitAttackingTeamFisherman                            float64 `json:"kills_kit_attacking_team_fisherman"`
				GamesKitAttackingTeamFisherman                            float64 `json:"games_kit_attacking_team_fisherman"`
				WinstreakKitAttackingTeamFisherman                        float64 `json:"winstreak_kit_attacking_team_fisherman"`
				ChallengeWins                                             float64 `json:"challenge_wins"`
				ChallengeWins0Solo                                        float64 `json:"challenge_wins_0_solo"`
				ChallengeWins0KitAttackingTeamFisherman                   float64 `json:"challenge_wins_0_kit_attacking_team_fisherman"`
				TimePlayedKitAttackingTeamFisherman                       float64 `json:"time_played_kit_attacking_team_fisherman"`
				ChallengeWinsSolo                                         float64 `json:"challenge_wins_solo"`
				SurvivedPlayersKitAttackingTeamFisherman                  float64 `json:"survived_players_kit_attacking_team_fisherman"`
				VoidKillsKitAttackingTeamFisherman                        float64 `json:"void_kills_kit_attacking_team_fisherman"`
				ChallengeWins0                                            float64 `json:"challenge_wins_0"`
				WinsKitAttackingTeamFisherman                             float64 `json:"wins_kit_attacking_team_fisherman"`
				ChestsOpenedKitAttackingTeamFisherman                     float64 `json:"chests_opened_kit_attacking_team_fisherman"`
				ChallengeWinsKitAttackingTeamFisherman                    float64 `json:"challenge_wins_kit_attacking_team_fisherman"`
				InGamePresentsCap201818                                   float64 `json:"inGamePresentsCap_2018_18"`
				DeathsKitAttackingTeamFisherman                           float64 `json:"deaths_kit_attacking_team_fisherman"`
				LossesKitAttackingTeamFisherman                           float64 `json:"losses_kit_attacking_team_fisherman"`
				LongestBowShotKitAttackingTeamFisherman                   float64 `json:"longest_bow_shot_kit_attacking_team_fisherman"`
				LongestBowKillKitAttackingTeamFisherman                   float64 `json:"longest_bow_kill_kit_attacking_team_fisherman"`
				MeleeKillsKitAttackingTeamFisherman                       float64 `json:"melee_kills_kit_attacking_team_fisherman"`
				ArrowsHitKitAttackingTeamFisherman                        float64 `json:"arrows_hit_kit_attacking_team_fisherman"`
				MobKillsKitAttackingTeamFisherman                         float64 `json:"mob_kills_kit_attacking_team_fisherman"`
				ArrowsShotKitAttackingTeamFisherman                       float64 `json:"arrows_shot_kit_attacking_team_fisherman"`
				ChallengeAttemptsHalfHealth                               float64 `json:"challenge_attempts_half_health"`
				ChallengeAttemptsRookieSolo                               float64 `json:"challenge_attempts_rookie_solo"`
				ChallengeAttempts8Solo                                    float64 `json:"challenge_attempts_8_solo"`
				ChallengeAttemptsArcherSolo                               float64 `json:"challenge_attempts_archer_solo"`
				ChallengeAttemptsArcher                                   float64 `json:"challenge_attempts_archer"`
				ChallengeAttemptsNoChestKitAdvancedSoloEnderman           float64 `json:"challenge_attempts_no_chest_kit_advanced_solo_enderman"`
				ChallengeAttemptsHalfHealthKitAdvancedSoloEnderman        float64 `json:"challenge_attempts_half_health_kit_advanced_solo_enderman"`
				ChallengeAttemptsUltimateWarrior                          float64 `json:"challenge_attempts_ultimate_warrior"`
				ChallengeAttemptsNoChestSolo                              float64 `json:"challenge_attempts_no_chest_solo"`
				ChallengeAttemptsHalfHealthSolo                           float64 `json:"challenge_attempts_half_health_solo"`
				ChallengeAttemptsRookieKitAdvancedSoloEnderman            float64 `json:"challenge_attempts_rookie_kit_advanced_solo_enderman"`
				ChallengeAttemptsNoBlockKitAdvancedSoloEnderman           float64 `json:"challenge_attempts_no_block_kit_advanced_solo_enderman"`
				ChallengeAttemptsPaper                                    float64 `json:"challenge_attempts_paper"`
				ChallengeAttemptsUhcSolo                                  float64 `json:"challenge_attempts_uhc_solo"`
				ChallengeAttempts8KitAdvancedSoloEnderman                 float64 `json:"challenge_attempts_8_kit_advanced_solo_enderman"`
				ChallengeAttemptsNoBlockSolo                              float64 `json:"challenge_attempts_no_block_solo"`
				ChallengeAttemptsNoBlock                                  float64 `json:"challenge_attempts_no_block"`
				ChallengeAttemptsNoChest                                  float64 `json:"challenge_attempts_no_chest"`
				ChallengeAttemptsUhc                                      float64 `json:"challenge_attempts_uhc"`
				ChallengeAttemptsPaperKitAdvancedSoloEnderman             float64 `json:"challenge_attempts_paper_kit_advanced_solo_enderman"`
				ChallengeAttemptsPaperSolo                                float64 `json:"challenge_attempts_paper_solo"`
				ChallengeAttemptsUltimateWarriorKitAdvancedSoloEnderman   float64 `json:"challenge_attempts_ultimate_warrior_kit_advanced_solo_enderman"`
				ChallengeAttemptsSolo                                     float64 `json:"challenge_attempts_solo"`
				ChallengeAttempts8                                        float64 `json:"challenge_attempts_8"`
				ChallengeAttemptsArcherKitAdvancedSoloEnderman            float64 `json:"challenge_attempts_archer_kit_advanced_solo_enderman"`
				ChallengeAttemptsRookie                                   float64 `json:"challenge_attempts_rookie"`
				ChallengeAttemptsKitAdvancedSoloEnderman                  float64 `json:"challenge_attempts_kit_advanced_solo_enderman"`
				ChallengeAttemptsUhcKitAdvancedSoloEnderman               float64 `json:"challenge_attempts_uhc_kit_advanced_solo_enderman"`
				ChallengeAttemptsUltimateWarriorSolo                      float64 `json:"challenge_attempts_ultimate_warrior_solo"`
				ChallengeAttempts                                         float64 `json:"challenge_attempts"`
				ChallengeWinsNoChestSolo                                  float64 `json:"challenge_wins_no_chest_solo"`
				ChallengeWinsArcher                                       float64 `json:"challenge_wins_archer"`
				ChallengeWinsArcherSolo                                   float64 `json:"challenge_wins_archer_solo"`
				ChallengeAttempts8KitAttackingTeamFisherman               float64 `json:"challenge_attempts_8_kit_attacking_team_fisherman"`
				ChallengeWinsRookieKitAttackingTeamFisherman              float64 `json:"challenge_wins_rookie_kit_attacking_team_fisherman"`
				ChallengeWinsUltimateWarriorKitAttackingTeamFisherman     float64 `json:"challenge_wins_ultimate_warrior_kit_attacking_team_fisherman"`
				ChallengeWinsUltimateWarriorSolo                          float64 `json:"challenge_wins_ultimate_warrior_solo"`
				ChallengeWinsNoBlockSolo                                  float64 `json:"challenge_wins_no_block_solo"`
				ChallengeWinsUhcKitAttackingTeamFisherman                 float64 `json:"challenge_wins_uhc_kit_attacking_team_fisherman"`
				ChallengeWinsHalfHealth                                   float64 `json:"challenge_wins_half_health"`
				ChallengeAttemptsNoChestKitAttackingTeamFisherman         float64 `json:"challenge_attempts_no_chest_kit_attacking_team_fisherman"`
				ChallengeWinsRookieSolo                                   float64 `json:"challenge_wins_rookie_solo"`
				ChallengeWinsUhc                                          float64 `json:"challenge_wins_uhc"`
				ChallengeWinsNoChestKitAttackingTeamFisherman             float64 `json:"challenge_wins_no_chest_kit_attacking_team_fisherman"`
				ChallengeWinsNoBlock                                      float64 `json:"challenge_wins_no_block"`
				ChallengeWinsNoChest                                      float64 `json:"challenge_wins_no_chest"`
				ChallengeAttemptsPaperKitAttackingTeamFisherman           float64 `json:"challenge_attempts_paper_kit_attacking_team_fisherman"`
				ChallengeAttemptsNoBlockKitAttackingTeamFisherman         float64 `json:"challenge_attempts_no_block_kit_attacking_team_fisherman"`
				ChallengeAttemptsUltimateWarriorKitAttackingTeamFisherman float64 `json:"challenge_attempts_ultimate_warrior_kit_attacking_team_fisherman"`
				ChallengeWinsPaperSolo                                    float64 `json:"challenge_wins_paper_solo"`
				ChallengeWinsPaperKitAttackingTeamFisherman               float64 `json:"challenge_wins_paper_kit_attacking_team_fisherman"`
				ChallengeWinsHalfHealthSolo                               float64 `json:"challenge_wins_half_health_solo"`
				ChallengeWinsUhcSolo                                      float64 `json:"challenge_wins_uhc_solo"`
				ChallengeWinsNoBlockKitAttackingTeamFisherman             float64 `json:"challenge_wins_no_block_kit_attacking_team_fisherman"`
				ChallengeWinsUltimateWarrior                              float64 `json:"challenge_wins_ultimate_warrior"`
				ChallengeAttemptsRookieKitAttackingTeamFisherman          float64 `json:"challenge_attempts_rookie_kit_attacking_team_fisherman"`
				ChallengeWinsPaper                                        float64 `json:"challenge_wins_paper"`
				ChallengeWinsRookie                                       float64 `json:"challenge_wins_rookie"`
				ChallengeAttemptsArcherKitAttackingTeamFisherman          float64 `json:"challenge_attempts_archer_kit_attacking_team_fisherman"`
				ChallengeAttemptsKitAttackingTeamFisherman                float64 `json:"challenge_attempts_kit_attacking_team_fisherman"`
				ChallengeWinsArcherKitAttackingTeamFisherman              float64 `json:"challenge_wins_archer_kit_attacking_team_fisherman"`
				ChallengeWinsHalfHealthKitAttackingTeamFisherman          float64 `json:"challenge_wins_half_health_kit_attacking_team_fisherman"`
				ChallengeAttemptsUhcKitAttackingTeamFisherman             float64 `json:"challenge_attempts_uhc_kit_attacking_team_fisherman"`
				ChallengeWins8                                            float64 `json:"challenge_wins_8"`
				ChallengeWins8Solo                                        float64 `json:"challenge_wins_8_solo"`
				ChallengeAttemptsHalfHealthKitAttackingTeamFisherman      float64 `json:"challenge_attempts_half_health_kit_attacking_team_fisherman"`
				ChallengeWins8KitAttackingTeamFisherman                   float64 `json:"challenge_wins_8_kit_attacking_team_fisherman"`
				ChallengeAttemptsHalfHealthKitAdvancedSoloHunter          float64 `json:"challenge_attempts_half_health_kit_advanced_solo_hunter"`
				ChallengeAttempts4KitAdvancedSoloHunter                   float64 `json:"challenge_attempts_4_kit_advanced_solo_hunter"`
				ChallengeAttemptsArcherKitAdvancedSoloHunter              float64 `json:"challenge_attempts_archer_kit_advanced_solo_hunter"`
				ChallengeAttemptsUhcKitAdvancedSoloHunter                 float64 `json:"challenge_attempts_uhc_kit_advanced_solo_hunter"`
				ChallengeAttempts4Solo                                    float64 `json:"challenge_attempts_4_solo"`
				ChallengeAttemptsKitAdvancedSoloHunter                    float64 `json:"challenge_attempts_kit_advanced_solo_hunter"`
				ChallengeAttempts4                                        float64 `json:"challenge_attempts_4"`
				ChallengeAttemptsPaperKitAdvancedSoloHunter               float64 `json:"challenge_attempts_paper_kit_advanced_solo_hunter"`
				ChallengeAttempts5Solo                                    float64 `json:"challenge_attempts_5_solo"`
				ChallengeAttemptsNoChestKitAdvancedSoloHunter             float64 `json:"challenge_attempts_no_chest_kit_advanced_solo_hunter"`
				ChallengeAttempts5KitAdvancedSoloHunter                   float64 `json:"challenge_attempts_5_kit_advanced_solo_hunter"`
				ChallengeAttempts5                                        float64 `json:"challenge_attempts_5"`
				ChallengeAttemptsKitAttackingTeamEnderman                 float64 `json:"challenge_attempts_kit_attacking_team_enderman"`
				ChallengeAttempts1Solo                                    float64 `json:"challenge_attempts_1_solo"`
				ChallengeAttempts1                                        float64 `json:"challenge_attempts_1"`
				GamesKitAttackingTeamEnderman                             float64 `json:"games_kit_attacking_team_enderman"`
				ChallengeAttemptsPaperKitAttackingTeamEnderman            float64 `json:"challenge_attempts_paper_kit_attacking_team_enderman"`
				ChallengeAttempts1KitAttackingTeamEnderman                float64 `json:"challenge_attempts_1_kit_attacking_team_enderman"`
				KitAttackingTeamEndermanInventory                         struct {
					ENDERPEARL0 string `json:"ENDER_PEARL:0"`
				} `json:"kit_attacking_team_enderman_inventory"`
				AssistsKitAttackingTeamEnderman                  float64 `json:"assists_kit_attacking_team_enderman"`
				SelectedPrestigeIcon                             string  `json:"selected_prestige_icon"`
				AngelOfDeathLevel                                float64 `json:"angel_of_death_level"`
				LongestBowKillKitAttackingTeamEnderman           float64 `json:"longest_bow_kill_kit_attacking_team_enderman"`
				FastestWinKitAttackingTeamEnderman               float64 `json:"fastest_win_kit_attacking_team_enderman"`
				KillstreakKitAttackingTeamEnderman               float64 `json:"killstreak_kit_attacking_team_enderman"`
				BowKillsKitAttackingTeamEnderman                 float64 `json:"bow_kills_kit_attacking_team_enderman"`
				WinsKitAttackingTeamEnderman                     float64 `json:"wins_kit_attacking_team_enderman"`
				WinstreakKitAttackingTeamEnderman                float64 `json:"winstreak_kit_attacking_team_enderman"`
				LossesLabKitAttackingTeamEnderman                float64 `json:"losses_lab_kit_attacking_team_enderman"`
				DeathsLabKitAttackingTeamEnderman                float64 `json:"deaths_lab_kit_attacking_team_enderman"`
				TimePlayedLabKitAttackingTeamEnderman            float64 `json:"time_played_lab_kit_attacking_team_enderman"`
				ChestsOpenedLabKitAttackingTeamEnderman          float64 `json:"chests_opened_lab_kit_attacking_team_enderman"`
				ActiveKitMEGARandom                              bool    `json:"activeKit_MEGA_random"`
				LongestBowKillMegaDoubles                        float64 `json:"longest_bow_kill_mega_doubles"`
				ArrowsShotMegaDoubles                            float64 `json:"arrows_shot_mega_doubles"`
				TimePlayedMegaDoubles                            float64 `json:"time_played_mega_doubles"`
				LossesMegaDoublesNormal                          float64 `json:"losses_mega_doubles_normal"`
				GamesMegaDoubles                                 float64 `json:"games_mega_doubles"`
				KillsMegaDoublesNormal                           float64 `json:"kills_mega_doubles_normal"`
				DeathsMegaDoublesNormal                          float64 `json:"deaths_mega_doubles_normal"`
				SurvivedPlayersMegaDoubles                       float64 `json:"survived_players_mega_doubles"`
				LossesMegaDoubles                                float64 `json:"losses_mega_doubles"`
				MeleeKillsMegaDoubles                            float64 `json:"melee_kills_mega_doubles"`
				KillsMegaDoubles                                 float64 `json:"kills_mega_doubles"`
				VoidKillsMegaDoubles                             float64 `json:"void_kills_mega_doubles"`
				DeathsMegaDoubles                                float64 `json:"deaths_mega_doubles"`
				ChestsOpenedMegaDoubles                          float64 `json:"chests_opened_mega_doubles"`
				ArrowsHitMegaDoubles                             float64 `json:"arrows_hit_mega_doubles"`
				WinsMegaDoublesNormal                            float64 `json:"wins_mega_doubles_normal"`
				WinstreakMegaDoubles                             float64 `json:"winstreak_mega_doubles"`
				KillstreakMegaDoubles                            float64 `json:"killstreak_mega_doubles"`
				WinsMegaDoubles                                  float64 `json:"wins_mega_doubles"`
				LongestBowKillKitMegaMegaScout                   float64 `json:"longest_bow_kill_kit_mega_mega_scout"`
				LongestBowShotMegaDoubles                        float64 `json:"longest_bow_shot_mega_doubles"`
				LongestBowShotKitMegaMegaScout                   float64 `json:"longest_bow_shot_kit_mega_mega_scout"`
				MostKillsGameMegaDoubles                         float64 `json:"most_kills_game_mega_doubles"`
				MostKillsGameKitMegaMegaScout                    float64 `json:"most_kills_game_kit_mega_mega_scout"`
				ArrowsHitKitMegaMegaScout                        float64 `json:"arrows_hit_kit_mega_mega_scout"`
				TimePlayedKitMegaMegaScout                       float64 `json:"time_played_kit_mega_mega_scout"`
				MeleeKillsKitMegaMegaScout                       float64 `json:"melee_kills_kit_mega_mega_scout"`
				ArrowsShotKitMegaMegaScout                       float64 `json:"arrows_shot_kit_mega_mega_scout"`
				ChestsOpenedKitMegaMegaScout                     float64 `json:"chests_opened_kit_mega_mega_scout"`
				MeleeKillsLabKitAttackingTeamEnderman            float64 `json:"melee_kills_lab_kit_attacking_team_enderman"`
				VoidKillsLabKitAttackingTeamEnderman             float64 `json:"void_kills_lab_kit_attacking_team_enderman"`
				KillsLabKitAttackingTeamEnderman                 float64 `json:"kills_lab_kit_attacking_team_enderman"`
				SurvivedPlayersLabKitAttackingTeamEnderman       float64 `json:"survived_players_lab_kit_attacking_team_enderman"`
				FastestWinLabKitAttackingTeamEnderman            float64 `json:"fastest_win_lab_kit_attacking_team_enderman"`
				GamesLabKitAttackingTeamEnderman                 float64 `json:"games_lab_kit_attacking_team_enderman"`
				LabWinTntMadnessLabKitAttackingTeamEnderman      float64 `json:"lab_win_tnt_madness_lab_kit_attacking_team_enderman"`
				KillstreakLabKitAttackingTeamEnderman            float64 `json:"killstreak_lab_kit_attacking_team_enderman"`
				WinsLabKitAttackingTeamEnderman                  float64 `json:"wins_lab_kit_attacking_team_enderman"`
				WinstreakLabKitAttackingTeamEnderman             float64 `json:"winstreak_lab_kit_attacking_team_enderman"`
				ArrowsHitLabKitAttackingTeamEnderman             float64 `json:"arrows_hit_lab_kit_attacking_team_enderman"`
				ArrowsShotLabKitAttackingTeamEnderman            float64 `json:"arrows_shot_lab_kit_attacking_team_enderman"`
				LongestBowShotLabKitAttackingTeamEnderman        float64 `json:"longest_bow_shot_lab_kit_attacking_team_enderman"`
				AssistsLabKitAttackingTeamEnderman               float64 `json:"assists_lab_kit_attacking_team_enderman"`
				LabWinSlimeLabKitAttackingTeamEnderman           float64 `json:"lab_win_slime_lab_kit_attacking_team_enderman"`
				ActiveKitTEAMSTourneyRandom                      bool    `json:"activeKit_TEAMS_tourney_random"`
				ActiveKitTEAMSTourney                            string  `json:"activeKit_TEAMS_tourney"`
				LevelFormatted                                   string  `json:"levelFormatted"`
				LongestBowShotTourneyKitAttackingTeamEnderman    float64 `json:"longest_bow_shot_tourney_kit_attacking_team_enderman"`
				LongestBowShotTourney                            float64 `json:"longest_bow_shot_tourney"`
				WinStreakTourney                                 float64 `json:"win_streak_tourney"`
				LongestBowShotTourneyCrazytourney                float64 `json:"longest_bow_shot_tourney_crazytourney"`
				TourneySwCrazySolo0WinStreak                     float64 `json:"tourney_sw_crazy_solo_0_win_streak"`
				TimePlayedTourney                                float64 `json:"time_played_tourney"`
				EnderpearlsThrownTourney                         float64 `json:"enderpearls_thrown_tourney"`
				ArrowsShotTourneyKitAttackingTeamEnderman        float64 `json:"arrows_shot_tourney_kit_attacking_team_enderman"`
				TourneySwCrazySolo0MeleeKills                    float64 `json:"tourney_sw_crazy_solo_0_melee_kills"`
				TourneySwCrazySolo0EnderpearlsThrown             float64 `json:"tourney_sw_crazy_solo_0_enderpearls_thrown"`
				TourneySwCrazySolo0Quits                         float64 `json:"tourney_sw_crazy_solo_0_quits"`
				KillsTourneyCrazytourney                         float64 `json:"kills_tourney_crazytourney"`
				LossesTourneyCrazytourney                        float64 `json:"losses_tourney_crazytourney"`
				SurvivedPlayersTourneyCrazytourney               float64 `json:"survived_players_tourney_crazytourney"`
				TourneySwCrazySolo0ChestsOpened                  float64 `json:"tourney_sw_crazy_solo_0_chests_opened"`
				TourneySwCrazySolo0ArrowsShot                    float64 `json:"tourney_sw_crazy_solo_0_arrows_shot"`
				TourneySwCrazySolo0Kills                         float64 `json:"tourney_sw_crazy_solo_0_kills"`
				BlocksPlacedTourney                              float64 `json:"blocks_placed_tourney"`
				LossesTourney                                    float64 `json:"losses_tourney"`
				TourneySwCrazySolo0SurvivedPlayers               float64 `json:"tourney_sw_crazy_solo_0_survived_players"`
				CoinsGainedTourney                               float64 `json:"coins_gained_tourney"`
				ChestsOpenedTourneyKitAttackingTeamEnderman      float64 `json:"chests_opened_tourney_kit_attacking_team_enderman"`
				TourneySwCrazySolo0BlocksPlaced                  float64 `json:"tourney_sw_crazy_solo_0_blocks_placed"`
				KillsTourney                                     float64 `json:"kills_tourney"`
				TourneySwCrazySolo0Coins                         float64 `json:"tourney_sw_crazy_solo_0_coins"`
				TourneySwCrazySolo0ArrowsHit                     float64 `json:"tourney_sw_crazy_solo_0_arrows_hit"`
				SurvivedPlayersTourneyKitAttackingTeamEnderman   float64 `json:"survived_players_tourney_kit_attacking_team_enderman"`
				ArrowsHitTourney                                 float64 `json:"arrows_hit_tourney"`
				TimePlayedTourneyKitAttackingTeamEnderman        float64 `json:"time_played_tourney_kit_attacking_team_enderman"`
				ArrowsHitTourneyCrazytourney                     float64 `json:"arrows_hit_tourney_crazytourney"`
				QuitsTourney                                     float64 `json:"quits_tourney"`
				LossesTourneyKitAttackingTeamEnderman            float64 `json:"losses_tourney_kit_attacking_team_enderman"`
				ChestsOpenedTourneyCrazytourney                  float64 `json:"chests_opened_tourney_crazytourney"`
				TimePlayedTourneyCrazytourney                    float64 `json:"time_played_tourney_crazytourney"`
				ArrowsShotTourneyCrazytourney                    float64 `json:"arrows_shot_tourney_crazytourney"`
				ArrowsHitTourneyKitAttackingTeamEnderman         float64 `json:"arrows_hit_tourney_kit_attacking_team_enderman"`
				DeathsTourneyKitAttackingTeamEnderman            float64 `json:"deaths_tourney_kit_attacking_team_enderman"`
				DeathsTourney                                    float64 `json:"deaths_tourney"`
				DeathsTourneyCrazytourney                        float64 `json:"deaths_tourney_crazytourney"`
				MeleeKillsTourneyCrazytourney                    float64 `json:"melee_kills_tourney_crazytourney"`
				TourneySwCrazySolo0TimePlayed                    float64 `json:"tourney_sw_crazy_solo_0_time_played"`
				TourneySwCrazySolo0Losses                        float64 `json:"tourney_sw_crazy_solo_0_losses"`
				TourneySwCrazySolo0LongestBowShot                float64 `json:"tourney_sw_crazy_solo_0_longest_bow_shot"`
				SurvivedPlayersTourney                           float64 `json:"survived_players_tourney"`
				ChestsOpenedTourney                              float64 `json:"chests_opened_tourney"`
				TourneySwCrazySolo0Deaths                        float64 `json:"tourney_sw_crazy_solo_0_deaths"`
				MeleeKillsTourney                                float64 `json:"melee_kills_tourney"`
				KillsTourneyKitAttackingTeamEnderman             float64 `json:"kills_tourney_kit_attacking_team_enderman"`
				MeleeKillsTourneyKitAttackingTeamEnderman        float64 `json:"melee_kills_tourney_kit_attacking_team_enderman"`
				ArrowsShotTourney                                float64 `json:"arrows_shot_tourney"`
				TourneySwCrazySolo0CoinsGained                   float64 `json:"tourney_sw_crazy_solo_0_coins_gained"`
				LongestBowKillTourneyCrazytourney                float64 `json:"longest_bow_kill_tourney_crazytourney"`
				MostKillsGameTourneyKitDefendingTeamFarmer       float64 `json:"most_kills_game_tourney_kit_defending_team_farmer"`
				MostKillsGameTourney                             float64 `json:"most_kills_game_tourney"`
				LongestBowKillTourneyKitDefendingTeamFarmer      float64 `json:"longest_bow_kill_tourney_kit_defending_team_farmer"`
				MostKillsGameTourneyCrazytourney                 float64 `json:"most_kills_game_tourney_crazytourney"`
				LongestBowKillTourney                            float64 `json:"longest_bow_kill_tourney"`
				KillsTourneyKitDefendingTeamFarmer               float64 `json:"kills_tourney_kit_defending_team_farmer"`
				TourneySwCrazySolo0MostKillsGame                 float64 `json:"tourney_sw_crazy_solo_0_most_kills_game"`
				MeleeKillsTourneyKitDefendingTeamFarmer          float64 `json:"melee_kills_tourney_kit_defending_team_farmer"`
				SurvivedPlayersTourneyKitDefendingTeamFarmer     float64 `json:"survived_players_tourney_kit_defending_team_farmer"`
				LossesTourneyKitDefendingTeamFarmer              float64 `json:"losses_tourney_kit_defending_team_farmer"`
				DeathsTourneyKitDefendingTeamFarmer              float64 `json:"deaths_tourney_kit_defending_team_farmer"`
				TourneySwCrazySolo0LongestBowKill                float64 `json:"tourney_sw_crazy_solo_0_longest_bow_kill"`
				EggThrownTourney                                 float64 `json:"egg_thrown_tourney"`
				ChestsOpenedTourneyKitDefendingTeamFarmer        float64 `json:"chests_opened_tourney_kit_defending_team_farmer"`
				TourneySwCrazySolo0EggThrown                     float64 `json:"tourney_sw_crazy_solo_0_egg_thrown"`
				TimePlayedTourneyKitDefendingTeamFarmer          float64 `json:"time_played_tourney_kit_defending_team_farmer"`
				SurvivedPlayersTourneyKitSupportingTeamEcologist float64 `json:"survived_players_tourney_kit_supporting_team_ecologist"`
				DeathsTourneyKitSupportingTeamEcologist          float64 `json:"deaths_tourney_kit_supporting_team_ecologist"`
				LossesTourneyKitSupportingTeamEcologist          float64 `json:"losses_tourney_kit_supporting_team_ecologist"`
				TimePlayedTourneyKitSupportingTeamEcologist      float64 `json:"time_played_tourney_kit_supporting_team_ecologist"`
				LongestBowShotTourneyKitDefendingTeamFrog        float64 `json:"longest_bow_shot_tourney_kit_defending_team_frog"`
				LongestBowKillTourneyKitDefendingTeamFrog        float64 `json:"longest_bow_kill_tourney_kit_defending_team_frog"`
				FastestWinTourneyKitDefendingTeamFrog            float64 `json:"fastest_win_tourney_kit_defending_team_frog"`
				FastestWinTourneyCrazytourney                    float64 `json:"fastest_win_tourney_crazytourney"`
				MostKillsGameTourneyKitDefendingTeamFrog         float64 `json:"most_kills_game_tourney_kit_defending_team_frog"`
				FastestWinTourney                                float64 `json:"fastest_win_tourney"`
				SurvivedPlayersTourneyKitDefendingTeamFrog       float64 `json:"survived_players_tourney_kit_defending_team_frog"`
				TimePlayedTourneyKitDefendingTeamFrog            float64 `json:"time_played_tourney_kit_defending_team_frog"`
				MeleeKillsTourneyKitDefendingTeamFrog            float64 `json:"melee_kills_tourney_kit_defending_team_frog"`
				BlocksBrokenTourney                              float64 `json:"blocks_broken_tourney"`
				TourneySwCrazySolo0Winstreak                     float64 `json:"tourney_sw_crazy_solo_0_winstreak"`
				KillstreakTourney                                float64 `json:"killstreak_tourney"`
				WinstreakTourney                                 float64 `json:"winstreak_tourney"`
				GamesTourneyCrazytourney                         float64 `json:"games_tourney_crazytourney"`
				AssistsTourneyCrazytourney                       float64 `json:"assists_tourney_crazytourney"`
				TourneySwCrazySolo0Killstreak                    float64 `json:"tourney_sw_crazy_solo_0_killstreak"`
				KillstreakTourneyKitDefendingTeamFrog            float64 `json:"killstreak_tourney_kit_defending_team_frog"`
				GamesTourney                                     float64 `json:"games_tourney"`
				TourneySwCrazySolo0Games                         float64 `json:"tourney_sw_crazy_solo_0_games"`
				TourneySwCrazySolo0Assists                       float64 `json:"tourney_sw_crazy_solo_0_assists"`
				TourneySwCrazySolo0Wins                          float64 `json:"tourney_sw_crazy_solo_0_wins"`
				WinstreakTourneyKitDefendingTeamFrog             float64 `json:"winstreak_tourney_kit_defending_team_frog"`
				TourneySwCrazySolo0BlocksBroken                  float64 `json:"tourney_sw_crazy_solo_0_blocks_broken"`
				KillstreakTourneyCrazytourney                    float64 `json:"killstreak_tourney_crazytourney"`
				WinstreakTourneyCrazytourney                     float64 `json:"winstreak_tourney_crazytourney"`
				ArrowsShotTourneyKitDefendingTeamFrog            float64 `json:"arrows_shot_tourney_kit_defending_team_frog"`
				TourneySwCrazySolo0FastestWin                    float64 `json:"tourney_sw_crazy_solo_0_fastest_win"`
				ChestsOpenedTourneyKitDefendingTeamFrog          float64 `json:"chests_opened_tourney_kit_defending_team_frog"`
				WinsTourneyKitDefendingTeamFrog                  float64 `json:"wins_tourney_kit_defending_team_frog"`
				WinsTourneyCrazytourney                          float64 `json:"wins_tourney_crazytourney"`
				GamesTourneyKitDefendingTeamFrog                 float64 `json:"games_tourney_kit_defending_team_frog"`
				WinsTourney                                      float64 `json:"wins_tourney"`
				ArrowsHitTourneyKitDefendingTeamFrog             float64 `json:"arrows_hit_tourney_kit_defending_team_frog"`
				AssistsTourney                                   float64 `json:"assists_tourney"`
				AssistsTourneyKitDefendingTeamFrog               float64 `json:"assists_tourney_kit_defending_team_frog"`
				KillsTourneyKitDefendingTeamFrog                 float64 `json:"kills_tourney_kit_defending_team_frog"`
				VoidKillsTourneyKitDefendingTeamFrog             float64 `json:"void_kills_tourney_kit_defending_team_frog"`
				VoidKillsTourneyCrazytourney                     float64 `json:"void_kills_tourney_crazytourney"`
				VoidKillsTourney                                 float64 `json:"void_kills_tourney"`
				TourneySwCrazySolo0VoidKills                     float64 `json:"tourney_sw_crazy_solo_0_void_kills"`
				KitBasicSoloFrogInventory                        struct {
					POTION2            string `json:"POTION:2"`
					LEATHERCHESTPLATE0 string `json:"LEATHER_CHESTPLATE:0"`
					LEATHERLEGGINGS0   string `json:"LEATHER_LEGGINGS:0"`
					SKULLITEM3         string `json:"SKULL_ITEM:3"`
					LEATHERBOOTS0      string `json:"LEATHER_BOOTS:0"`
				} `json:"kit_basic_solo_frog_inventory"`
				DeathsTourneyKitDefendingTeamFrog   float64 `json:"deaths_tourney_kit_defending_team_frog"`
				LossesTourneyKitDefendingTeamFrog   float64 `json:"losses_tourney_kit_defending_team_frog"`
				MobKillsTourneyKitDefendingTeamFrog float64 `json:"mob_kills_tourney_kit_defending_team_frog"`
				MobKillsTourneyCrazytourney         float64 `json:"mob_kills_tourney_crazytourney"`
				TourneySwCrazySolo0MobKills         float64 `json:"tourney_sw_crazy_solo_0_mob_kills"`
				MobKillsTourney                     float64 `json:"mob_kills_tourney"`
				BowKillsTourneyKitDefendingTeamFrog float64 `json:"bow_kills_tourney_kit_defending_team_frog"`
				BowKillsTourney                     float64 `json:"bow_kills_tourney"`
				BowKillsTourneyCrazytourney         float64 `json:"bow_kills_tourney_crazytourney"`
				TourneySwCrazySolo0BowKills         float64 `json:"tourney_sw_crazy_solo_0_bow_kills"`
				ItemsEnchantedTourney               float64 `json:"items_enchanted_tourney"`
				TourneySwCrazySolo0ItemsEnchanted   float64 `json:"tourney_sw_crazy_solo_0_items_enchanted"`
				MostKillsGameKitDefendingTeamFrog   float64 `json:"most_kills_game_kit_defending_team_frog"`
				LongestBowKillKitDefendingTeamFrog  float64 `json:"longest_bow_kill_kit_defending_team_frog"`
				DeathsKitDefendingTeamFrog          float64 `json:"deaths_kit_defending_team_frog"`
				KillsKitDefendingTeamFrog           float64 `json:"kills_kit_defending_team_frog"`
				LossesKitDefendingTeamFrog          float64 `json:"losses_kit_defending_team_frog"`
				SurvivedPlayersKitDefendingTeamFrog float64 `json:"survived_players_kit_defending_team_frog"`
				TimePlayedKitDefendingTeamFrog      float64 `json:"time_played_kit_defending_team_frog"`
				MeleeKillsKitDefendingTeamFrog      float64 `json:"melee_kills_kit_defending_team_frog"`
				ChestsOpenedKitDefendingTeamFrog    float64 `json:"chests_opened_kit_defending_team_frog"`
				FastestWinKitDefendingTeamFrog      float64 `json:"fastest_win_kit_defending_team_frog"`
				VoidKillsKitDefendingTeamFrog       float64 `json:"void_kills_kit_defending_team_frog"`
				WinstreakKitDefendingTeamFrog       float64 `json:"winstreak_kit_defending_team_frog"`
				GamesKitDefendingTeamFrog           float64 `json:"games_kit_defending_team_frog"`
				WinsKitDefendingTeamFrog            float64 `json:"wins_kit_defending_team_frog"`
				KillstreakKitDefendingTeamFrog      float64 `json:"killstreak_kit_defending_team_frog"`
				TeamFrost                           float64 `json:"team_frost"`
				HeadsDecentKitDefendingTeamFrog     float64 `json:"heads_decent_kit_defending_team_frog"`
				HeadsSalty                          float64 `json:"heads_salty"`
				HeadsSaltyKitDefendingTeamFrog      float64 `json:"heads_salty_kit_defending_team_frog"`
				HeadsMeh                            float64 `json:"heads_meh"`
				Heads                               float64 `json:"heads"`
				HeadsYuckyKitDefendingTeamFrog      float64 `json:"heads_yucky_kit_defending_team_frog"`
				HeadsDivine                         float64 `json:"heads_divine"`
				HeadsYuckySolo                      float64 `json:"heads_yucky_solo"`
				HeadsDivineSolo                     float64 `json:"heads_divine_solo"`
				HeadsMehSolo                        float64 `json:"heads_meh_solo"`
				HeadsYucky                          float64 `json:"heads_yucky"`
				HeadsKitDefendingTeamFrog           float64 `json:"heads_kit_defending_team_frog"`
				HeadsDecent                         float64 `json:"heads_decent"`
				HeadsSaltySolo                      float64 `json:"heads_salty_solo"`
				HeadsDivineKitDefendingTeamFrog     float64 `json:"heads_divine_kit_defending_team_frog"`
				HeadsMehKitDefendingTeamFrog        float64 `json:"heads_meh_kit_defending_team_frog"`
				HeadsSolo                           float64 `json:"heads_solo"`
				HeadsDecentSolo                     float64 `json:"heads_decent_solo"`
				HeadCollection                      struct {
					Recent []struct {
						UUID      string  `json:"uuid"`
						Timestamp float64 `json:"timestamp"`
						Mode      string  `json:"mode"`
						Sacrifice string  `json:"sacrifice"`
					} `json:"recent"`
					Prestigious []struct {
						UUID      string  `json:"uuid"`
						Timestamp float64 `json:"timestamp"`
						Mode      string  `json:"mode"`
						Sacrifice string  `json:"sacrifice"`
					} `json:"prestigious"`
				} `json:"head_collection"`
				AssistsKitDefendingTeamFrog                  float64 `json:"assists_kit_defending_team_frog"`
				LongestBowShotKitDefendingTeamFrog           float64 `json:"longest_bow_shot_kit_defending_team_frog"`
				HeadsTastySolo                               float64 `json:"heads_tasty_solo"`
				HeadsTastyKitDefendingTeamFrog               float64 `json:"heads_tasty_kit_defending_team_frog"`
				ArrowsShotKitDefendingTeamFrog               float64 `json:"arrows_shot_kit_defending_team_frog"`
				HeadsTasty                                   float64 `json:"heads_tasty"`
				ArrowsHitKitDefendingTeamFrog                float64 `json:"arrows_hit_kit_defending_team_frog"`
				HeadsEwwSolo                                 float64 `json:"heads_eww_solo"`
				HeadsHeavenly                                float64 `json:"heads_heavenly"`
				HeadsMehKitAttackingTeamEnderman             float64 `json:"heads_meh_kit_attacking_team_enderman"`
				HeadsHeavenlyKitAttackingTeamEnderman        float64 `json:"heads_heavenly_kit_attacking_team_enderman"`
				HeadsEww                                     float64 `json:"heads_eww"`
				HeadsSaltyKitAttackingTeamEnderman           float64 `json:"heads_salty_kit_attacking_team_enderman"`
				HeadsEwwKitAttackingTeamEnderman             float64 `json:"heads_eww_kit_attacking_team_enderman"`
				HeadsKitAttackingTeamEnderman                float64 `json:"heads_kit_attacking_team_enderman"`
				HeadsHeavenlySolo                            float64 `json:"heads_heavenly_solo"`
				MostKillsGameKitDefendingTeamFarmer          float64 `json:"most_kills_game_kit_defending_team_farmer"`
				LongestBowKillKitDefendingTeamFarmer         float64 `json:"longest_bow_kill_kit_defending_team_farmer"`
				MeleeKillsKitDefendingTeamFarmer             float64 `json:"melee_kills_kit_defending_team_farmer"`
				TimePlayedKitDefendingTeamFarmer             float64 `json:"time_played_kit_defending_team_farmer"`
				SurvivedPlayersKitDefendingTeamFarmer        float64 `json:"survived_players_kit_defending_team_farmer"`
				LossesKitDefendingTeamFarmer                 float64 `json:"losses_kit_defending_team_farmer"`
				DeathsKitDefendingTeamFarmer                 float64 `json:"deaths_kit_defending_team_farmer"`
				GamesKitDefendingTeamFarmer                  float64 `json:"games_kit_defending_team_farmer"`
				KillsKitDefendingTeamFarmer                  float64 `json:"kills_kit_defending_team_farmer"`
				ChestsOpenedKitDefendingTeamFarmer           float64 `json:"chests_opened_kit_defending_team_farmer"`
				HeadsKitAdvancedSoloFarmer                   float64 `json:"heads_kit_advanced_solo_farmer"`
				HeadsHeavenlyKitAdvancedSoloFarmer           float64 `json:"heads_heavenly_kit_advanced_solo_farmer"`
				LabWinLuckyBlocksLabKitAttackingTeamEnderman float64 `json:"lab_win_lucky_blocks_lab_kit_attacking_team_enderman"`
				Privategames                                 struct {
					EnableTeleportMayhem  bool   `json:"enable_teleport_mayhem"`
					EnableMaxKitsAndPerks bool   `json:"enable_max_kits_and_perks"`
					ChestSwords           string `json:"chest_swords"`
					LowGravity            bool   `json:"low_gravity"`
					Dragons               string `json:"dragons"`
					EnableLegacyItems     bool   `json:"enable_legacy_items"`
					Speed                 string `json:"speed"`
					HealthBuff            string `json:"health_buff"`
					ChestArmour           string `json:"chest_armour"`
					ChestBows             string `json:"chest_bows"`
					OneHitOneKill         bool   `json:"one_hit_one_kill"`
					EnableNightTime       bool   `json:"enable_night_time"`
					NoKits                bool   `json:"no_kits"`
				} `json:"privategames"`
				HeadsSucculentKitAttackingTeamEnderman   float64 `json:"heads_succulent_kit_attacking_team_enderman"`
				HeadsSucculentSolo                       float64 `json:"heads_succulent_solo"`
				HeadsSucculent                           float64 `json:"heads_succulent"`
				HeadsYuckyKitAttackingTeamEnderman       float64 `json:"heads_yucky_kit_attacking_team_enderman"`
				MobKillsKitAttackingTeamEnderman         float64 `json:"mob_kills_kit_attacking_team_enderman"`
				BowKillsKitDefendingTeamFrog             float64 `json:"bow_kills_kit_defending_team_frog"`
				SkywarsHalloweenBoxes                    float64 `json:"skywars_halloween_boxes"`
				ChestsOpenedLabKitDefendingTeamFarmer    float64 `json:"chests_opened_lab_kit_defending_team_farmer"`
				DeathsLabKitDefendingTeamFarmer          float64 `json:"deaths_lab_kit_defending_team_farmer"`
				LossesLabKitDefendingTeamFarmer          float64 `json:"losses_lab_kit_defending_team_farmer"`
				SurvivedPlayersLabKitDefendingTeamFarmer float64 `json:"survived_players_lab_kit_defending_team_farmer"`
				TimePlayedLabKitDefendingTeamFarmer      float64 `json:"time_played_lab_kit_defending_team_farmer"`
				ChestsOpenedKitMegaMegaDefault           float64 `json:"chests_opened_kit_mega_mega_default"`
				TimePlayedKitMegaMegaDefault             float64 `json:"time_played_kit_mega_mega_default"`
				FreeEventKeySkywarsChristmasBoxes2019    bool    `json:"free_event_key_skywars_christmas_boxes_2019"`
				SkywarsChristmasBoxes                    float64 `json:"skywars_christmas_boxes"`
				BowKillsLab                              float64 `json:"bow_kills_lab"`
				BowKillsLabKitAttackingTeamEnderman      float64 `json:"bow_kills_lab_kit_attacking_team_enderman"`
				BowKillsLabSolo                          float64 `json:"bow_kills_lab_solo"`
				LabWinLuckyBlocksLabTeam                 float64 `json:"lab_win_lucky_blocks_lab_team"`
				Opals                                    float64 `json:"opals"`
				Shard                                    float64 `json:"shard"`
				ShardKitAttackingTeamEnderman            float64 `json:"shard_kit_attacking_team_enderman"`
				ShardSolo                                float64 `json:"shard_solo"`
				ShardTeam                                float64 `json:"shard_team"`
				LastTourneyAd                            float64 `json:"lastTourneyAd"`
			} `json:"SkyWars"`
			TrueCombat struct {
				CrazywallsKillsSoloChaos            float64  `json:"crazywalls_kills_solo_chaos"`
				ItemsEnchanted                      float64  `json:"items_enchanted"`
				ArrowsShot                          float64  `json:"arrows_shot"`
				Coins                               float64  `json:"coins"`
				Kills                               float64  `json:"kills"`
				Wins                                float64  `json:"wins"`
				SurvivedPlayers                     float64  `json:"survived_players"`
				CrazywallsWinsSoloChaos             float64  `json:"crazywalls_wins_solo_chaos"`
				WinStreak                           float64  `json:"win_streak"`
				Games                               float64  `json:"games"`
				CrazywallsLossesSoloChaos           float64  `json:"crazywalls_losses_solo_chaos"`
				CrazywallsGamesSoloChaos            float64  `json:"crazywalls_games_solo_chaos"`
				Deaths                              float64  `json:"deaths"`
				Losses                              float64  `json:"losses"`
				CrazywallsDeathsSoloChaos           float64  `json:"crazywalls_deaths_solo_chaos"`
				CrazywallsDeathsSolo                float64  `json:"crazywalls_deaths_solo"`
				CrazywallsGamesSolo                 float64  `json:"crazywalls_games_solo"`
				CrazywallsLossesSolo                float64  `json:"crazywalls_losses_solo"`
				CrazywallsKillsSolo                 float64  `json:"crazywalls_kills_solo"`
				CrazywallsWinsSolo                  float64  `json:"crazywalls_wins_solo"`
				GoldenSkulls                        float64  `json:"golden_skulls"`
				GiantZombie                         float64  `json:"giant_zombie"`
				SoloChaosBountyHunter               float64  `json:"solo_chaos_bounty_hunter"`
				LiveCombat                          bool     `json:"live_combat"`
				CrazywallsKillsWeeklyBSolo          float64  `json:"crazywalls_kills_weekly_b_solo"`
				CrazywallsKillsMonthlyBSolo         float64  `json:"crazywalls_kills_monthly_b_solo"`
				SkullsGathered                      float64  `json:"skulls_gathered"`
				TeamAdrenaline                      float64  `json:"team_adrenaline"`
				SoloChaosTombRobber                 float64  `json:"solo_chaos_tomb_robber"`
				GiantZombieLegendaries              float64  `json:"giant_zombie_legendaries"`
				SoloChaosBusinessman                float64  `json:"solo_chaos_businessman"`
				TeamBlazingArrow                    float64  `json:"team_blazing_arrow"`
				KitBasicChaosAlchemist              float64  `json:"kit_basic_chaos_alchemist"`
				ActiveKitSoloChaos                  string   `json:"activeKit_Solo_chaos"`
				CrazywallsKillsWeeklyBSoloChaos     float64  `json:"crazywalls_kills_weekly_b_solo_chaos"`
				CrazywallsKillsMonthlyBSoloChaos    float64  `json:"crazywalls_kills_monthly_b_solo_chaos"`
				Packages                            []string `json:"packages"`
				VotesGarden                         float64  `json:"votes_Garden"`
				GiantZombieRares                    float64  `json:"giant_zombie_rares"`
				SoloChaosRusher                     float64  `json:"solo_chaos_rusher"`
				GoldDust                            float64  `json:"gold_dust"`
				SoloChaosBerserk                    float64  `json:"solo_chaos_berserk"`
				SoloChaosSuperLuck                  float64  `json:"solo_chaos_super_luck"`
				TeamGroupHeal                       float64  `json:"team_group_heal"`
				SoloBlazingArrow                    float64  `json:"solo_blazing_arrow"`
				SoloLuckyDrops                      float64  `json:"solo_lucky_drops"`
				SoloSpeedMining                     float64  `json:"solo_speed_mining"`
				TeamRusher                          float64  `json:"team_rusher"`
				SoloRusher                          float64  `json:"solo_rusher"`
				SoloBountyHunter                    float64  `json:"solo_bounty_hunter"`
				TeamSwiftness                       float64  `json:"team_swiftness"`
				TeamBerserk                         float64  `json:"team_berserk"`
				SoloAdrenaline                      float64  `json:"solo_adrenaline"`
				KillsWeeklyB                        float64  `json:"kills_weekly_b"`
				KillsMonthlyB                       float64  `json:"kills_monthly_b"`
				TeamSmartMining                     float64  `json:"team_smart_mining"`
				SoloVampirism                       float64  `json:"solo_vampirism"`
				TeamBountyHunter                    float64  `json:"team_bounty_hunter"`
				TeamPrecision                       float64  `json:"team_precision"`
				KitBasicChaosWeaponsmith            float64  `json:"kit_basic_chaos_weaponsmith"`
				KillsWeeklyA                        float64  `json:"kills_weekly_a"`
				CrazywallsKillsMonthlyASoloChaos    float64  `json:"crazywalls_kills_monthly_a_solo_chaos"`
				CrazywallsKillsWeeklyASoloChaos     float64  `json:"crazywalls_kills_weekly_a_solo_chaos"`
				KillsMonthlyA                       float64  `json:"kills_monthly_a"`
				TeamLastStand                       float64  `json:"team_last_stand"`
				SoloPrecision                       float64  `json:"solo_precision"`
				SoloKnowledge                       float64  `json:"solo_knowledge"`
				KitBasicChaosArmorer                float64  `json:"kit_basic_chaos_armorer"`
				CraftingLuckyEnchantedBookSharpness float64  `json:"crafting_lucky_enchanted_book_sharpness"`
				SoloChaosBlazingArrow               float64  `json:"solo_chaos_blazing_arrow"`
				SoloSwiftness                       float64  `json:"solo_swiftness"`
				CraftingLuckyBrawler                float64  `json:"crafting_lucky_brawler"`
				SoloChaosKnowledge                  float64  `json:"solo_chaos_knowledge"`
				SoloChaosAdrenaline                 float64  `json:"solo_chaos_adrenaline"`
				ActiveKitTeamChaos                  string   `json:"activeKit_Team_chaos"`
				CrazywallsWinsTeamChaos             float64  `json:"crazywalls_wins_team_chaos"`
				CrazywallsKillsWeeklyBTeamChaos     float64  `json:"crazywalls_kills_weekly_b_team_chaos"`
				CrazywallsKillsMonthlyBTeamChaos    float64  `json:"crazywalls_kills_monthly_b_team_chaos"`
				CrazywallsKillsTeamChaos            float64  `json:"crazywalls_kills_team_chaos"`
				ActiveKitSolo                       string   `json:"activeKit_Solo"`
				CrazywallsKillsWeeklyASolo          float64  `json:"crazywalls_kills_weekly_a_solo"`
				TeamAnnoyOMite                      float64  `json:"team_annoy-o-mite"`
				SoloChaosVampirism                  float64  `json:"solo_chaos_vampirism"`
				SoloChaosSmartyPants                float64  `json:"solo_chaos_smarty_pants"`
				ArrowsHit                           float64  `json:"arrows_hit"`
				CrazywallsDeathsTeamChaos           float64  `json:"crazywalls_deaths_team_chaos"`
				CrazywallsGamesTeamChaos            float64  `json:"crazywalls_games_team_chaos"`
				CrazywallsKillsMonthlyATeamChaos    float64  `json:"crazywalls_kills_monthly_a_team_chaos"`
				CrazywallsKillsWeeklyATeamChaos     float64  `json:"crazywalls_kills_weekly_a_team_chaos"`
				CrazywallsLossesTeamChaos           float64  `json:"crazywalls_losses_team_chaos"`
			} `json:"TrueCombat"`
			SuperSmash struct {
				SmashLevel       float64 `json:"smashLevel"`
				LastLevelTHEBULK float64 `json:"lastLevel_THE_BULK"`
				ActiveClass      string  `json:"active_class"`
				Coins            float64 `json:"coins"`
				WinStreak        float64 `json:"win_streak"`
				DamageDealtTeams float64 `json:"damage_dealt_teams"`
				ClassStats       struct {
					Tinman struct {
						Kills          float64 `json:"kills"`
						HomingMissiles struct {
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							DamageDealt       float64 `json:"damage_dealt"`
							Kills             float64 `json:"kills"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smasher           float64 `json:"smasher"`
							SmasherNormal     float64 `json:"smasher_normal"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
						} `json:"homing_missiles"`
						ThrowCake struct {
							SmashedTeams float64 `json:"smashed_teams"`
							Smashed      float64 `json:"smashed"`
						} `json:"throw_cake"`
						LaserCannon struct {
							SmasherTeams      float64 `json:"smasher_teams"`
							Smasher           float64 `json:"smasher"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							DamageDealt       float64 `json:"damage_dealt"`
							Kills             float64 `json:"kills"`
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
						} `json:"laser_cannon"`
						Overload struct {
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							DamageDealt       float64 `json:"damage_dealt"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills2V2          float64 `json:"kills_2v2"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
						} `json:"overload"`
						GamesTeams  float64 `json:"games_teams"`
						WinsTeams   float64 `json:"wins_teams"`
						DeathsTeams float64 `json:"deaths_teams"`
						KillsTeams  float64 `json:"kills_teams"`
						RocketPunch struct {
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smasher           float64 `json:"smasher"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
						} `json:"rocket_punch"`
						SmashedTeams     float64 `json:"smashed_teams"`
						SmasherTeams     float64 `json:"smasher_teams"`
						Smashed          float64 `json:"smashed"`
						WinStreakTeams   float64 `json:"win_streak_teams"`
						Games            float64 `json:"games"`
						Smasher          float64 `json:"smasher"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						WinStreak        float64 `json:"win_streak"`
						Wins             float64 `json:"wins"`
						Deaths           float64 `json:"deaths"`
						DamageDealt      float64 `json:"damage_dealt"`
						GamesNormal      float64 `json:"games_normal"`
						Losses           float64 `json:"losses"`
						EggBazooka       struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"egg_bazooka"`
						KillsNormal   float64 `json:"kills_normal"`
						LossesNormal  float64 `json:"losses_normal"`
						SmasherNormal float64 `json:"smasher_normal"`
						SmashedNormal float64 `json:"smashed_normal"`
						SeismicSlam   struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"seismic_slam"`
						DeathsNormal      float64 `json:"deaths_normal"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						WinStreakNormal   float64 `json:"win_streak_normal"`
						WinsNormal        float64 `json:"wins_normal"`
						Melee             struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Smashed           float64 `json:"smashed"`
							SmashedTeams      float64 `json:"smashed_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Smasher           float64 `json:"smasher"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Kills             float64 `json:"kills"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							SmashedNormal     float64 `json:"smashed_normal"`
						} `json:"melee"`
						LossesTeams    float64 `json:"losses_teams"`
						ForceLightning struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"force_lightning"`
						Wins2V2      float64 `json:"wins_2v2"`
						Smashed2V2   float64 `json:"smashed_2v2"`
						Kills2V2     float64 `json:"kills_2v2"`
						WinStreak2V2 float64 `json:"win_streak_2v2"`
						CakeStorm    struct {
							Smashed2V2 float64 `json:"smashed_2v2"`
							Smashed    float64 `json:"smashed"`
						} `json:"cake_storm"`
						Deaths2V2      float64 `json:"deaths_2v2"`
						Smasher2V2     float64 `json:"smasher_2v2"`
						Games2V2       float64 `json:"games_2v2"`
						DamageDealt2V2 float64 `json:"damage_dealt_2v2"`
						Losses2V2      float64 `json:"losses_2v2"`
						Batarang       struct {
							Smashed2V2 float64 `json:"smashed_2v2"`
							Smashed    float64 `json:"smashed"`
						} `json:"batarang"`
						Telepunch struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"telepunch"`
					} `json:"TINMAN"`
					GeneralCluck struct {
						Bazooka struct {
							Kills             float64 `json:"kills"`
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smasher           float64 `json:"smasher"`
						} `json:"bazooka"`
						WinStreakNormal float64 `json:"win_streak_normal"`
						EggBazooka      struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							SmasherNormal     float64 `json:"smasher_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smasher           float64 `json:"smasher"`
							Kills             float64 `json:"kills"`
						} `json:"egg_bazooka"`
						Reinforcements struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
						} `json:"reinforcements"`
						Melee struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							KillsNormal       float64 `json:"kills_normal"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smasher           float64 `json:"smasher"`
							Kills             float64 `json:"kills"`
							Smashed           float64 `json:"smashed"`
							SmasherNormal     float64 `json:"smasher_normal"`
						} `json:"melee"`
						DeathsNormal      float64 `json:"deaths_normal"`
						GamesNormal       float64 `json:"games_normal"`
						Games             float64 `json:"games"`
						Deaths            float64 `json:"deaths"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						WinStreak         float64 `json:"win_streak"`
						DamageDealt       float64 `json:"damage_dealt"`
						Wins              float64 `json:"wins"`
						Smasher           float64 `json:"smasher"`
						SmasherNormal     float64 `json:"smasher_normal"`
						KillsNormal       float64 `json:"kills_normal"`
						Kills             float64 `json:"kills"`
						WinsNormal        float64 `json:"wins_normal"`
						Losses            float64 `json:"losses"`
						LossesNormal      float64 `json:"losses_normal"`
						Smashed           float64 `json:"smashed"`
						SmashedNormal     float64 `json:"smashed_normal"`
					} `json:"GENERAL_CLUCK"`
					Skullfire struct {
						FlamingDesertEagle struct {
							Kills             float64 `json:"kills"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smasher           float64 `json:"smasher"`
							DamageDealt       float64 `json:"damage_dealt"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							Kills2V2          float64 `json:"kills_2v2"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
						} `json:"flaming_desert_eagle"`
						Kills float64 `json:"kills"`
						Melee struct {
							SmashedNormal     float64 `json:"smashed_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Smashed           float64 `json:"smashed"`
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
						} `json:"melee"`
						Grenade struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsNormal       float64 `json:"kills_normal"`
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
						} `json:"grenade"`
						Smasher           float64 `json:"smasher"`
						DeathsNormal      float64 `json:"deaths_normal"`
						Deaths            float64 `json:"deaths"`
						Smashed           float64 `json:"smashed"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						DesertEagle       struct {
							Kills             float64 `json:"kills"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smasher           float64 `json:"smasher"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							Kills2V2          float64 `json:"kills_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
						} `json:"desert_eagle"`
						Games           float64 `json:"games"`
						WinStreakNormal float64 `json:"win_streak_normal"`
						DamageDealt     float64 `json:"damage_dealt"`
						WinsNormal      float64 `json:"wins_normal"`
						GamesNormal     float64 `json:"games_normal"`
						WinStreak       float64 `json:"win_streak"`
						SmasherNormal   float64 `json:"smasher_normal"`
						Wins            float64 `json:"wins"`
						SmashedNormal   float64 `json:"smashed_normal"`
						KillsNormal     float64 `json:"kills_normal"`
						LossesNormal    float64 `json:"losses_normal"`
						Losses          float64 `json:"losses"`
						Batarang        struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"batarang"`
						SmasherTeams float64 `json:"smasher_teams"`
						DeathsTeams  float64 `json:"deaths_teams"`
						LaserCannon  struct {
							Smashed      float64 `json:"smashed"`
							SmashedTeams float64 `json:"smashed_teams"`
							Smashed2V2   float64 `json:"smashed_2v2"`
						} `json:"laser_cannon"`
						KillsTeams       float64 `json:"kills_teams"`
						SmashedTeams     float64 `json:"smashed_teams"`
						WinStreakTeams   float64 `json:"win_streak_teams"`
						WinsTeams        float64 `json:"wins_teams"`
						GamesTeams       float64 `json:"games_teams"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						SwingPin         struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"swing_pin"`
						KiBlast struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"ki_blast"`
						Kills2V2       float64 `json:"kills_2v2"`
						Wins2V2        float64 `json:"wins_2v2"`
						Smasher2V2     float64 `json:"smasher_2v2"`
						Games2V2       float64 `json:"games_2v2"`
						WinStreak2V2   float64 `json:"win_streak_2v2"`
						DamageDealt2V2 float64 `json:"damage_dealt_2v2"`
						Deaths2V2      float64 `json:"deaths_2v2"`
						Smashed2V2     float64 `json:"smashed_2v2"`
						Losses2V2      float64 `json:"losses_2v2"`
					} `json:"SKULLFIRE"`
					Marauder struct {
						Melee struct {
							Smasher           float64 `json:"smasher"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							SmashedTeams      float64 `json:"smashed_teams"`
							Smashed           float64 `json:"smashed"`
							SmasherTeams      float64 `json:"smasher_teams"`
							KillsTeams        float64 `json:"kills_teams"`
						} `json:"melee"`
						ForceLightning struct {
							KillsNormal       float64 `json:"kills_normal"`
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
						} `json:"force_lightning"`
						WinsNormal  float64 `json:"wins_normal"`
						MonsterMash struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"monster_mash"`
						GamesNormal     float64 `json:"games_normal"`
						Smashed         float64 `json:"smashed"`
						SmasherNormal   float64 `json:"smasher_normal"`
						Smasher         float64 `json:"smasher"`
						DamageDealt     float64 `json:"damage_dealt"`
						WinStreakNormal float64 `json:"win_streak_normal"`
						ForcePull       struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							SmasherNormal     float64 `json:"smasher_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
						} `json:"force_pull"`
						KillsNormal       float64 `json:"kills_normal"`
						Games             float64 `json:"games"`
						DeathsNormal      float64 `json:"deaths_normal"`
						SmashedNormal     float64 `json:"smashed_normal"`
						Wins              float64 `json:"wins"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						Kills             float64 `json:"kills"`
						Deaths            float64 `json:"deaths"`
						WinStreak         float64 `json:"win_streak"`
						KillsTeams        float64 `json:"kills_teams"`
						WinsTeams         float64 `json:"wins_teams"`
						CakeStorm         struct {
							Smashed      float64 `json:"smashed"`
							SmashedTeams float64 `json:"smashed_teams"`
						} `json:"cake_storm"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						WinStreakTeams   float64 `json:"win_streak_teams"`
						SmashedTeams     float64 `json:"smashed_teams"`
						DeathsTeams      float64 `json:"deaths_teams"`
						SmasherTeams     float64 `json:"smasher_teams"`
						GamesTeams       float64 `json:"games_teams"`
						Botmubile        struct {
							SmashedTeams  float64 `json:"smashed_teams"`
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"botmubile"`
						Batarang struct {
							SmashedTeams  float64 `json:"smashed_teams"`
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"batarang"`
						Frostbolt struct {
							Smashed      float64 `json:"smashed"`
							SmashedTeams float64 `json:"smashed_teams"`
						} `json:"frostbolt"`
						Losses       float64 `json:"losses"`
						LossesTeams  float64 `json:"losses_teams"`
						LossesNormal float64 `json:"losses_normal"`
						LaserCannon  struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"laser_cannon"`
						HomingMissiles struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"homing_missiles"`
					} `json:"MARAUDER"`
					Frosty struct {
						Melee struct {
							DamageDealt       float64 `json:"damage_dealt"`
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Kills             float64 `json:"kills"`
							SmasherTeams      float64 `json:"smasher_teams"`
							Smasher           float64 `json:"smasher"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
						} `json:"melee"`
						DeathsTeams float64 `json:"deaths_teams"`
						Frostbolt   struct {
							DamageDealt       float64 `json:"damage_dealt"`
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							SmasherTeams      float64 `json:"smasher_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
						} `json:"frostbolt"`
						Losses           float64 `json:"losses"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						Smasher          float64 `json:"smasher"`
						SmashedTeams     float64 `json:"smashed_teams"`
						FreezingBreath   struct {
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Kills             float64 `json:"kills"`
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
						} `json:"freezing_breath"`
						Smashed float64 `json:"smashed"`
						Games   float64 `json:"games"`
						Bazooka struct {
							SmashedTeams float64 `json:"smashed_teams"`
							Smashed      float64 `json:"smashed"`
						} `json:"bazooka"`
						GamesTeams     float64 `json:"games_teams"`
						LossesTeams    float64 `json:"losses_teams"`
						KillsTeams     float64 `json:"kills_teams"`
						DamageDealt    float64 `json:"damage_dealt"`
						SmasherTeams   float64 `json:"smasher_teams"`
						Kills          float64 `json:"kills"`
						Deaths         float64 `json:"deaths"`
						SmashedNormal  float64 `json:"smashed_normal"`
						HomingMissiles struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"homing_missiles"`
						DeathsNormal      float64 `json:"deaths_normal"`
						WinStreak         float64 `json:"win_streak"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						GamesNormal       float64 `json:"games_normal"`
						Wins              float64 `json:"wins"`
						WinsNormal        float64 `json:"wins_normal"`
						SmasherNormal     float64 `json:"smasher_normal"`
						KillsNormal       float64 `json:"kills_normal"`
						WinStreakNormal   float64 `json:"win_streak_normal"`
					} `json:"FROSTY"`
					TheBulk struct {
						SeismicSlam struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smasher           float64 `json:"smasher"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
						} `json:"seismic_slam"`
						Boulder struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
						} `json:"boulder"`
						MonsterCharge struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							Smasher           float64 `json:"smasher"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
						} `json:"monster_charge"`
						Melee struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							Kills             float64 `json:"kills"`
							SmasherNormal     float64 `json:"smasher_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smasher           float64 `json:"smasher"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							SmashedTeams      float64 `json:"smashed_teams"`
							Smashed2V2        float64 `json:"smashed_2v2"`
						} `json:"melee"`
						Games             float64 `json:"games"`
						DeathsNormal      float64 `json:"deaths_normal"`
						LossesNormal      float64 `json:"losses_normal"`
						DamageDealt       float64 `json:"damage_dealt"`
						Losses            float64 `json:"losses"`
						Deaths            float64 `json:"deaths"`
						GamesNormal       float64 `json:"games_normal"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						WinStreak         float64 `json:"win_streak"`
						Wins              float64 `json:"wins"`
						SmasherNormal     float64 `json:"smasher_normal"`
						Batarang          struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"batarang"`
						WinsNormal  float64 `json:"wins_normal"`
						Smasher     float64 `json:"smasher"`
						Kills       float64 `json:"kills"`
						MonsterMash struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smasher           float64 `json:"smasher"`
							Kills             float64 `json:"kills"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							SmasherTeams      float64 `json:"smasher_teams"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed2V2        float64 `json:"smashed_2v2"`
						} `json:"monster_mash"`
						Smashed          float64 `json:"smashed"`
						SmashedNormal    float64 `json:"smashed_normal"`
						WinStreakNormal  float64 `json:"win_streak_normal"`
						KillsNormal      float64 `json:"kills_normal"`
						SmashedTeams     float64 `json:"smashed_teams"`
						WinsTeams        float64 `json:"wins_teams"`
						SmasherTeams     float64 `json:"smasher_teams"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						KillsTeams       float64 `json:"kills_teams"`
						ForceLightning   struct {
							SmashedTeams  float64 `json:"smashed_teams"`
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"force_lightning"`
						DeathsTeams    float64 `json:"deaths_teams"`
						GamesTeams     float64 `json:"games_teams"`
						WinStreakTeams float64 `json:"win_streak_teams"`
						Reinforcements struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"reinforcements"`
						KiBlast struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"ki_blast"`
						LaserCannon struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							SmashedTeams  float64 `json:"smashed_teams"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"laser_cannon"`
						Botmubile struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"botmubile"`
						ThrowCake struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"throw_cake"`
						ForcePull struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							Smashed2V2    float64 `json:"smashed_2v2"`
							SmashedTeams  float64 `json:"smashed_teams"`
						} `json:"force_pull"`
						SpooderBuddies struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"spooder_buddies"`
						Overload struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"overload"`
						DesertEagle struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"desert_eagle"`
						DamageDealt2V2 float64 `json:"damage_dealt_2v2"`
						Smasher2V2     float64 `json:"smasher_2v2"`
						Smashed2V2     float64 `json:"smashed_2v2"`
						Wins2V2        float64 `json:"wins_2v2"`
						WinStreak2V2   float64 `json:"win_streak_2v2"`
						Kills2V2       float64 `json:"kills_2v2"`
						Games2V2       float64 `json:"games_2v2"`
						Deaths2V2      float64 `json:"deaths_2v2"`
						SwingPin       struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"swing_pin"`
						Frostbolt struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"frostbolt"`
						HomingMissiles struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"homing_missiles"`
						SpiderKick struct {
							Smashed       float64 `json:"smashed"`
							Smashed2V2    float64 `json:"smashed_2v2"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"spider_kick"`
						LossesTeams float64 `json:"losses_teams"`
						Bazooka     struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"bazooka"`
						StaticLaser struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"static_laser"`
						FlamingDesertEagle struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"flaming_desert_eagle"`
						NotchedBow struct {
							Smashed      float64 `json:"smashed"`
							SmashedTeams float64 `json:"smashed_teams"`
						} `json:"notched_bow"`
						FriendWins       float64 `json:"friend_wins"`
						FriendWinsNormal float64 `json:"friend_wins_normal"`
						ChargedBeam      struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"charged_beam"`
						GrapplingHook struct {
							Smashed    float64 `json:"smashed"`
							Smashed2V2 float64 `json:"smashed_2v2"`
						} `json:"grappling_hook"`
						WebShot struct {
							Smashed2V2 float64 `json:"smashed_2v2"`
							Smashed    float64 `json:"smashed"`
						} `json:"web_shot"`
						Defecake struct {
							Smashed2V2 float64 `json:"smashed_2v2"`
							Smashed    float64 `json:"smashed"`
						} `json:"defecake"`
						EggBazooka struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"egg_bazooka"`
						FriendLosses        float64 `json:"friend_losses"`
						FriendLossesNormal  float64 `json:"friend_losses_normal"`
						OneVOneWinsNormal   float64 `json:"one_v_one_wins_normal"`
						OneVOneWins         float64 `json:"one_v_one_wins"`
						OneVOneLosses       float64 `json:"one_v_one_losses"`
						OneVOneLossesNormal float64 `json:"one_v_one_losses_normal"`
					} `json:"THE_BULK"`
					DuskCrawler struct {
						VoidSlash struct {
							DamageDealt       float64 `json:"damage_dealt"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills3V3          float64 `json:"kills_3v3"`
							Smasher3V3        float64 `json:"smasher_3v3"`
							Smasher           float64 `json:"smasher"`
							DamageDealt3V3    float64 `json:"damage_dealt_3v3"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Kills2V2          float64 `json:"kills_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
						} `json:"void_slash"`
						ForcePull struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"force_pull"`
						LossesNormal float64 `json:"losses_normal"`
						Melee        struct {
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher3V3        float64 `json:"smasher_3v3"`
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							Kills3V3          float64 `json:"kills_3v3"`
							DamageDealt3V3    float64 `json:"damage_dealt_3v3"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smashed2V2        float64 `json:"smashed_2v2"`
						} `json:"melee"`
						Losses            float64 `json:"losses"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						Deaths            float64 `json:"deaths"`
						DamageDealt       float64 `json:"damage_dealt"`
						Teleboom          struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealt3V3    float64 `json:"damage_dealt_3v3"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							Kills             float64 `json:"kills"`
							SmasherTeams      float64 `json:"smasher_teams"`
							Smasher           float64 `json:"smasher"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Kills2V2          float64 `json:"kills_2v2"`
							Smasher2V2        float64 `json:"smasher_2v2"`
						} `json:"teleboom"`
						Games         float64 `json:"games"`
						DeathsNormal  float64 `json:"deaths_normal"`
						KillsNormal   float64 `json:"kills_normal"`
						Kills         float64 `json:"kills"`
						SmashedNormal float64 `json:"smashed_normal"`
						Smashed       float64 `json:"smashed"`
						GamesNormal   float64 `json:"games_normal"`
						Losses3V3     float64 `json:"losses_3v3"`
						Telepunch     struct {
							Kills3V3          float64 `json:"kills_3v3"`
							Smasher           float64 `json:"smasher"`
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher3V3        float64 `json:"smasher_3v3"`
							DamageDealt3V3    float64 `json:"damage_dealt_3v3"`
							Kills             float64 `json:"kills"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Smasher2V2        float64 `json:"smasher_2v2"`
							DamageDealt2V2    float64 `json:"damage_dealt_2v2"`
							Kills2V2          float64 `json:"kills_2v2"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
						} `json:"telepunch"`
						SpooderBuddies struct {
							Smashed    float64 `json:"smashed"`
							Smashed3V3 float64 `json:"smashed_3v3"`
							Smashed2V2 float64 `json:"smashed_2v2"`
						} `json:"spooder_buddies"`
						Games3V3       float64 `json:"games_3v3"`
						Smashed3V3     float64 `json:"smashed_3v3"`
						Smasher        float64 `json:"smasher"`
						Deaths3V3      float64 `json:"deaths_3v3"`
						Smasher3V3     float64 `json:"smasher_3v3"`
						Kills3V3       float64 `json:"kills_3v3"`
						DamageDealt3V3 float64 `json:"damage_dealt_3v3"`
						SmasherTeams   float64 `json:"smasher_teams"`
						WinsTeams      float64 `json:"wins_teams"`
						LaserCannon    struct {
							SmashedTeams  float64 `json:"smashed_teams"`
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"laser_cannon"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						KillsTeams       float64 `json:"kills_teams"`
						DeathsTeams      float64 `json:"deaths_teams"`
						WinStreak        float64 `json:"win_streak"`
						SmashedTeams     float64 `json:"smashed_teams"`
						GamesTeams       float64 `json:"games_teams"`
						WinStreakTeams   float64 `json:"win_streak_teams"`
						Wins             float64 `json:"wins"`
						Deaths2V2        float64 `json:"deaths_2v2"`
						Smasher2V2       float64 `json:"smasher_2v2"`
						Kills2V2         float64 `json:"kills_2v2"`
						DamageDealt2V2   float64 `json:"damage_dealt_2v2"`
						Losses2V2        float64 `json:"losses_2v2"`
						Smashed2V2       float64 `json:"smashed_2v2"`
						Games2V2         float64 `json:"games_2v2"`
						Wins2V2          float64 `json:"wins_2v2"`
						WinStreak2V2     float64 `json:"win_streak_2v2"`
						Boulder          struct {
							SmashedTeams  float64 `json:"smashed_teams"`
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"boulder"`
						WinsNormal      float64 `json:"wins_normal"`
						WinStreakNormal float64 `json:"win_streak_normal"`
						SmasherNormal   float64 `json:"smasher_normal"`
						Botmubile       struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"botmubile"`
						Frostbolt struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"frostbolt"`
						CakeStorm struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"cake_storm"`
						Reinforcements struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"reinforcements"`
						Batarang struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"batarang"`
						ForceLightning struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"force_lightning"`
						SeismicSlam struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed2V2    float64 `json:"smashed_2v2"`
						} `json:"seismic_slam"`
						ThrowCake struct {
							Smashed2V2    float64 `json:"smashed_2v2"`
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"throw_cake"`
						HomingMissiles struct {
							Smashed2V2 float64 `json:"smashed_2v2"`
							Smashed    float64 `json:"smashed"`
						} `json:"homing_missiles"`
						KameBeam struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"kame_beam"`
						DesertEagle struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"desert_eagle"`
						Bazooka struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"bazooka"`
						KiBlast struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"ki_blast"`
						RocketPunch struct {
							Smashed float64 `json:"smashed"`
						} `json:"rocket_punch"`
						MonsterCharge struct {
							Smashed    float64 `json:"smashed"`
							Smashed2V2 float64 `json:"smashed_2v2"`
						} `json:"monster_charge"`
						SwingPin struct {
							Smashed    float64 `json:"smashed"`
							Smashed2V2 float64 `json:"smashed_2v2"`
						} `json:"swing_pin"`
						ChargedBeam struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"charged_beam"`
					} `json:"DUSK_CRAWLER"`
					Botmun struct {
						Batarang struct {
							Kills             float64 `json:"kills"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher           float64 `json:"smasher"`
							KillsNormal       float64 `json:"kills_normal"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
						} `json:"batarang"`
						Deaths          float64 `json:"deaths"`
						GamesNormal     float64 `json:"games_normal"`
						Games           float64 `json:"games"`
						WinStreak       float64 `json:"win_streak"`
						WinStreakNormal float64 `json:"win_streak_normal"`
						Kills           float64 `json:"kills"`
						Botmubile       struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills             float64 `json:"kills"`
							DamageDealt       float64 `json:"damage_dealt"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smasher           float64 `json:"smasher"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
							SmasherTeams      float64 `json:"smasher_teams"`
						} `json:"botmubile"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						Wins              float64 `json:"wins"`
						Smasher           float64 `json:"smasher"`
						Melee             struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smasher           float64 `json:"smasher"`
							SmasherTeams      float64 `json:"smasher_teams"`
							Kills             float64 `json:"kills"`
							KillsTeams        float64 `json:"kills_teams"`
							SmashedTeams      float64 `json:"smashed_teams"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
						} `json:"melee"`
						KillsNormal   float64 `json:"kills_normal"`
						SmasherNormal float64 `json:"smasher_normal"`
						DamageDealt   float64 `json:"damage_dealt"`
						DeathsNormal  float64 `json:"deaths_normal"`
						WinsNormal    float64 `json:"wins_normal"`
						WinsTeams     float64 `json:"wins_teams"`
						ForcePull     struct {
							SmashedTeams  float64 `json:"smashed_teams"`
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"force_pull"`
						KillsTeams       float64 `json:"kills_teams"`
						WinStreakTeams   float64 `json:"win_streak_teams"`
						DeathsTeams      float64 `json:"deaths_teams"`
						GamesTeams       float64 `json:"games_teams"`
						Smashed          float64 `json:"smashed"`
						SmasherTeams     float64 `json:"smasher_teams"`
						SmashedTeams     float64 `json:"smashed_teams"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						SmashedNormal    float64 `json:"smashed_normal"`
						LaserCannon      struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							SmashedTeams  float64 `json:"smashed_teams"`
						} `json:"laser_cannon"`
						EggBazooka struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"egg_bazooka"`
						Bazooka struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							SmashedTeams  float64 `json:"smashed_teams"`
						} `json:"bazooka"`
						Losses       float64 `json:"losses"`
						LossesNormal float64 `json:"losses_normal"`
						Frostbolt    struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"frostbolt"`
						ForceLightning struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"force_lightning"`
						DesertEagle struct {
							Smashed      float64 `json:"smashed"`
							SmashedTeams float64 `json:"smashed_teams"`
						} `json:"desert_eagle"`
						LossesTeams float64 `json:"losses_teams"`
						Boulder     struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"boulder"`
						Reinforcements struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"reinforcements"`
						FlamingDesertEagle struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"flaming_desert_eagle"`
						SpooderBuddies struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"spooder_buddies"`
						MonsterCharge struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"monster_charge"`
						KiBlast struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"ki_blast"`
						ThrowCake struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"throw_cake"`
						SpiderKick struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"spider_kick"`
						SwingPin struct {
							Smashed      float64 `json:"smashed"`
							SmashedTeams float64 `json:"smashed_teams"`
						} `json:"swing_pin"`
						SeismicSlam struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"seismic_slam"`
						HomingMissiles struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"homing_missiles"`
						FreezingBreath struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"freezing_breath"`
						OnionCannon struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"onion_cannon"`
					} `json:"BOTMUN"`
					Goku struct {
						Smasher float64 `json:"smasher"`
						KiBlast struct {
							Kills             float64 `json:"kills"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher           float64 `json:"smasher"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
						} `json:"ki_blast"`
						Wins  float64 `json:"wins"`
						Melee struct {
							DamageDealt       float64 `json:"damage_dealt"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
							Kills             float64 `json:"kills"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smasher           float64 `json:"smasher"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							KillsTeams        float64 `json:"kills_teams"`
						} `json:"melee"`
						Kills    float64 `json:"kills"`
						KameBeam struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
						} `json:"kame_beam"`
						Deaths            float64 `json:"deaths"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						WinStreak         float64 `json:"win_streak"`
						DamageDealt       float64 `json:"damage_dealt"`
						WinsNormal        float64 `json:"wins_normal"`
						Games             float64 `json:"games"`
						SmasherNormal     float64 `json:"smasher_normal"`
						WinStreakNormal   float64 `json:"win_streak_normal"`
						GamesNormal       float64 `json:"games_normal"`
						DeathsNormal      float64 `json:"deaths_normal"`
						KillsNormal       float64 `json:"kills_normal"`
						SmashedNormal     float64 `json:"smashed_normal"`
						Smashed           float64 `json:"smashed"`
						Reinforcements    struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"reinforcements"`
						Batarang struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"batarang"`
						Bazooka struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"bazooka"`
						GamesTeams       float64 `json:"games_teams"`
						SmasherTeams     float64 `json:"smasher_teams"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						WinStreakTeams   float64 `json:"win_streak_teams"`
						WinsTeams        float64 `json:"wins_teams"`
						KillsTeams       float64 `json:"kills_teams"`
						DeathsTeams      float64 `json:"deaths_teams"`
						CakeStorm        struct {
							SmashedTeams float64 `json:"smashed_teams"`
							Smashed      float64 `json:"smashed"`
						} `json:"cake_storm"`
						SmashedTeams   float64 `json:"smashed_teams"`
						SpooderBuddies struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"spooder_buddies"`
						Botmubile struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"botmubile"`
					} `json:"GOKU"`
					Spoderman struct {
						Melee struct {
							DamageDealt       float64 `json:"damage_dealt"`
							SmasherNormal     float64 `json:"smasher_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smasher           float64 `json:"smasher"`
							Kills             float64 `json:"kills"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
						} `json:"melee"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						SpooderBuddies    struct {
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher           float64 `json:"smasher"`
							KillsNormal       float64 `json:"kills_normal"`
							Kills             float64 `json:"kills"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
						} `json:"spooder_buddies"`
						WebShot struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Smasher           float64 `json:"smasher"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							SmasherTeams      float64 `json:"smasher_teams"`
						} `json:"web_shot"`
						DeathsNormal float64 `json:"deaths_normal"`
						DamageDealt  float64 `json:"damage_dealt"`
						GamesNormal  float64 `json:"games_normal"`
						Batarang     struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
							SmashedTeams  float64 `json:"smashed_teams"`
						} `json:"batarang"`
						SpiderKick struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealtTeams  float64 `json:"damage_dealt_teams"`
							KillsTeams        float64 `json:"kills_teams"`
							Kills             float64 `json:"kills"`
						} `json:"spider_kick"`
						Smashed         float64 `json:"smashed"`
						WinStreakNormal float64 `json:"win_streak_normal"`
						WinsNormal      float64 `json:"wins_normal"`
						Wins            float64 `json:"wins"`
						Smasher         float64 `json:"smasher"`
						Kills           float64 `json:"kills"`
						WinStreak       float64 `json:"win_streak"`
						Games           float64 `json:"games"`
						Deaths          float64 `json:"deaths"`
						KillsNormal     float64 `json:"kills_normal"`
						SmashedNormal   float64 `json:"smashed_normal"`
						SmasherNormal   float64 `json:"smasher_normal"`
						WallClimber     struct {
							DamageDealt      float64 `json:"damage_dealt"`
							DamageDealtTeams float64 `json:"damage_dealt_teams"`
						} `json:"wall_climber"`
						DamageDealtTeams float64 `json:"damage_dealt_teams"`
						SmashedTeams     float64 `json:"smashed_teams"`
						GamesTeams       float64 `json:"games_teams"`
						WinsTeams        float64 `json:"wins_teams"`
						KillsTeams       float64 `json:"kills_teams"`
						ForceLightning   struct {
							SmashedTeams float64 `json:"smashed_teams"`
							Smashed      float64 `json:"smashed"`
						} `json:"force_lightning"`
						DeathsTeams    float64 `json:"deaths_teams"`
						SmasherTeams   float64 `json:"smasher_teams"`
						WinStreakTeams float64 `json:"win_streak_teams"`
						Bazooka        struct {
							Smashed      float64 `json:"smashed"`
							SmashedTeams float64 `json:"smashed_teams"`
						} `json:"bazooka"`
						LaserCannon struct {
							SmashedTeams float64 `json:"smashed_teams"`
							Smashed      float64 `json:"smashed"`
						} `json:"laser_cannon"`
					} `json:"SPODERMAN"`
					CakeMonster struct {
						CakeStorm struct {
							SmasherNormal     float64 `json:"smasher_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills             float64 `json:"kills"`
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher           float64 `json:"smasher"`
						} `json:"cake_storm"`
						Defecake struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
						} `json:"defecake"`
						WinStreak float64 `json:"win_streak"`
						SwingPin  struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							KillsNormal       float64 `json:"kills_normal"`
						} `json:"swing_pin"`
						Smashed           float64 `json:"smashed"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						ThrowCake         struct {
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher           float64 `json:"smasher"`
							KillsNormal       float64 `json:"kills_normal"`
							Kills             float64 `json:"kills"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
						} `json:"throw_cake"`
						KillsNormal     float64 `json:"kills_normal"`
						WinStreakNormal float64 `json:"win_streak_normal"`
						Melee           struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Smasher           float64 `json:"smasher"`
							SmasherNormal     float64 `json:"smasher_normal"`
							Kills             float64 `json:"kills"`
							DamageDealt       float64 `json:"damage_dealt"`
							SmashedNormal     float64 `json:"smashed_normal"`
							Smashed           float64 `json:"smashed"`
						} `json:"melee"`
						Smasher    float64 `json:"smasher"`
						SpiderKick struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"spider_kick"`
						GamesNormal    float64 `json:"games_normal"`
						Deaths         float64 `json:"deaths"`
						WinsNormal     float64 `json:"wins_normal"`
						SmasherNormal  float64 `json:"smasher_normal"`
						SmashedNormal  float64 `json:"smashed_normal"`
						DamageDealt    float64 `json:"damage_dealt"`
						Kills          float64 `json:"kills"`
						Games          float64 `json:"games"`
						Wins           float64 `json:"wins"`
						DeathsNormal   float64 `json:"deaths_normal"`
						ForceLightning struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"force_lightning"`
						MonsterCharge struct {
							SmashedNormal float64 `json:"smashed_normal"`
							Smashed       float64 `json:"smashed"`
						} `json:"monster_charge"`
						LaserCannon struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"laser_cannon"`
					} `json:"CAKE_MONSTER"`
					ShoopDaWhoop struct {
						DamageDealt float64 `json:"damage_dealt"`
						Melee       struct {
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
						} `json:"melee"`
						StaticLaser struct {
							Smasher           float64 `json:"smasher"`
							Kills             float64 `json:"kills"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							KillsNormal       float64 `json:"kills_normal"`
						} `json:"static_laser"`
						FirMaLazer struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
						} `json:"fir_ma_lazer"`
						ChargedBeam struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Kills             float64 `json:"kills"`
						} `json:"charged_beam"`
						WinsNormal        float64 `json:"wins_normal"`
						GamesNormal       float64 `json:"games_normal"`
						Smasher           float64 `json:"smasher"`
						KillsNormal       float64 `json:"kills_normal"`
						WinStreakNormal   float64 `json:"win_streak_normal"`
						Games             float64 `json:"games"`
						DeathsNormal      float64 `json:"deaths_normal"`
						SmasherNormal     float64 `json:"smasher_normal"`
						Wins              float64 `json:"wins"`
						WinStreak         float64 `json:"win_streak"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						Deaths            float64 `json:"deaths"`
						Kills             float64 `json:"kills"`
					} `json:"SHOOP_DA_WHOOP"`
					Sanic struct {
						Dash struct {
							DamageDealt       float64 `json:"damage_dealt"`
							Smasher           float64 `json:"smasher"`
							Smashed           float64 `json:"smashed"`
							SmashedNormal     float64 `json:"smashed_normal"`
							KillsNormal       float64 `json:"kills_normal"`
							Kills             float64 `json:"kills"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
						} `json:"dash"`
						Kills float64 `json:"kills"`
						Boom  struct {
							Kills             float64 `json:"kills"`
							KillsNormal       float64 `json:"kills_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
						} `json:"boom"`
						Melee struct {
							Kills             float64 `json:"kills"`
							Smasher           float64 `json:"smasher"`
							SmasherNormal     float64 `json:"smasher_normal"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
							DamageDealt       float64 `json:"damage_dealt"`
							KillsNormal       float64 `json:"kills_normal"`
						} `json:"melee"`
						DesertEagle struct {
							Smashed       float64 `json:"smashed"`
							SmashedNormal float64 `json:"smashed_normal"`
						} `json:"desert_eagle"`
						SmashedNormal     float64 `json:"smashed_normal"`
						SmasherNormal     float64 `json:"smasher_normal"`
						Deaths            float64 `json:"deaths"`
						KillsNormal       float64 `json:"kills_normal"`
						DamageDealtNormal float64 `json:"damage_dealt_normal"`
						OnionCannon       struct {
							DamageDealt       float64 `json:"damage_dealt"`
							DamageDealtNormal float64 `json:"damage_dealt_normal"`
						} `json:"onion_cannon"`
						DeathsNormal float64 `json:"deaths_normal"`
						Smashed      float64 `json:"smashed"`
						Losses       float64 `json:"losses"`
						GamesNormal  float64 `json:"games_normal"`
						LossesNormal float64 `json:"losses_normal"`
						DamageDealt  float64 `json:"damage_dealt"`
						Games        float64 `json:"games"`
						Smasher      float64 `json:"smasher"`
					} `json:"SANIC"`
				} `json:"class_stats"`
				Smasher                     float64 `json:"smasher"`
				KillsMonthlyB               float64 `json:"kills_monthly_b"`
				GamesMonthlyB               float64 `json:"games_monthly_b"`
				Kills                       float64 `json:"kills"`
				Wins                        float64 `json:"wins"`
				GamesWeeklyB                float64 `json:"games_weekly_b"`
				KillsWeeklyB                float64 `json:"kills_weekly_b"`
				WinsTeams                   float64 `json:"wins_teams"`
				DamageDealt                 float64 `json:"damage_dealt"`
				SmashedTeams                float64 `json:"smashed_teams"`
				KillsTeams                  float64 `json:"kills_teams"`
				SmasherTeams                float64 `json:"smasher_teams"`
				Smashed                     float64 `json:"smashed"`
				DeathsTeams                 float64 `json:"deaths_teams"`
				Deaths                      float64 `json:"deaths"`
				WinsWeeklyB                 float64 `json:"wins_weekly_b"`
				GamesTeams                  float64 `json:"games_teams"`
				Games                       float64 `json:"games"`
				WinsMonthlyB                float64 `json:"wins_monthly_b"`
				XpTINMAN                    float64 `json:"xp_TINMAN"`
				LastLevelTINMAN             float64 `json:"lastLevel_TINMAN"`
				SmasherNormal               float64 `json:"smasher_normal"`
				LossesNormal                float64 `json:"losses_normal"`
				LossesWeeklyB               float64 `json:"losses_weekly_b"`
				LossesMonthlyB              float64 `json:"losses_monthly_b"`
				Losses                      float64 `json:"losses"`
				SmashedNormal               float64 `json:"smashed_normal"`
				DeathsNormal                float64 `json:"deaths_normal"`
				GamesNormal                 float64 `json:"games_normal"`
				KillsNormal                 float64 `json:"kills_normal"`
				DamageDealtNormal           float64 `json:"damage_dealt_normal"`
				SmashLevelTotal             float64 `json:"smash_level_total"`
				WinsNormal                  float64 `json:"wins_normal"`
				XpGENERALCLUCK              float64 `json:"xp_GENERAL_CLUCK"`
				LastLevelGENERALCLUCK       float64 `json:"lastLevel_GENERAL_CLUCK"`
				XpSKULLFIRE                 float64 `json:"xp_SKULLFIRE"`
				LastLevelSKULLFIRE          float64 `json:"lastLevel_SKULLFIRE"`
				ExpBoosterPurchases10Plays  float64 `json:"expBooster_purchases_10_plays"`
				XpMARAUDER                  float64 `json:"xp_MARAUDER"`
				LastLevelMARAUDER           float64 `json:"lastLevel_MARAUDER"`
				Quits                       float64 `json:"quits"`
				GamesWeeklyA                float64 `json:"games_weekly_a"`
				KillsWeeklyA                float64 `json:"kills_weekly_a"`
				WinsWeeklyA                 float64 `json:"wins_weekly_a"`
				VotesStrawberryTowers       float64 `json:"votes_Strawberry Towers"`
				LossesTeams                 float64 `json:"losses_teams"`
				LossesWeeklyA               float64 `json:"losses_weekly_a"`
				XpFROSTY                    float64 `json:"xp_FROSTY"`
				LastLevelFROSTY             float64 `json:"lastLevel_FROSTY"`
				ExpiredBooster              bool    `json:"expired_booster"`
				LastLevelDUSKCRAWLER        float64 `json:"lastLevel_DUSK_CRAWLER"`
				XpBOTMUN                    float64 `json:"xp_BOTMUN"`
				LastLevelBOTMUN             float64 `json:"lastLevel_BOTMUN"`
				VotesCubed                  float64 `json:"votes_Cubed"`
				ExpBoosterPurchases30Plays  float64 `json:"expBooster_purchases_30_plays"`
				Smasher3V3                  float64 `json:"smasher_3v3"`
				DamageDealt3V3              float64 `json:"damage_dealt_3v3"`
				Kills3V3                    float64 `json:"kills_3v3"`
				Deaths3V3                   float64 `json:"deaths_3v3"`
				Games3V3                    float64 `json:"games_3v3"`
				Smashed3V3                  float64 `json:"smashed_3v3"`
				Losses3V3                   float64 `json:"losses_3v3"`
				XpDUSKCRAWLER               float64 `json:"xp_DUSK_CRAWLER"`
				VotesGunmetal               float64 `json:"votes_Gunmetal"`
				XpGOKU                      float64 `json:"xp_GOKU"`
				LastLevelGOKU               float64 `json:"lastLevel_GOKU"`
				Deaths2V2                   float64 `json:"deaths_2v2"`
				Kills2V2                    float64 `json:"kills_2v2"`
				Games2V2                    float64 `json:"games_2v2"`
				Smashed2V2                  float64 `json:"smashed_2v2"`
				Smasher2V2                  float64 `json:"smasher_2v2"`
				Losses2V2                   float64 `json:"losses_2v2"`
				DamageDealt2V2              float64 `json:"damage_dealt_2v2"`
				Wins2V2                     float64 `json:"wins_2v2"`
				ExpBoosterPurchases100Plays float64 `json:"expBooster_purchases_100_plays"`
				XpSPODERMAN                 float64 `json:"xp_SPODERMAN"`
				LastLevelSPODERMAN          float64 `json:"lastLevel_SPODERMAN"`
				VotesGrove                  float64 `json:"votes_Grove"`
				VotesToybox                 float64 `json:"votes_Toybox"`
				VotesColorClash             float64 `json:"votes_Color Clash"`
				XpTHEBULK                   float64 `json:"xp_THE_BULK"`
				PgDUSKCRAWLER               float64 `json:"pg_DUSK_CRAWLER"`
				PgBOTMUN                    float64 `json:"pg_BOTMUN"`
				XpCAKEMONSTER               float64 `json:"xp_CAKE_MONSTER"`
				LastLevelCAKEMONSTER        float64 `json:"lastLevel_CAKE_MONSTER"`
				PgTHEBULK                   float64 `json:"pg_THE_BULK"`
				GamesMonthlyA               float64 `json:"games_monthly_a"`
				KillsMonthlyA               float64 `json:"kills_monthly_a"`
				WinsMonthlyA                float64 `json:"wins_monthly_a"`
				XpSHOOPDAWHOOP              float64 `json:"xp_SHOOP_DA_WHOOP"`
				LastLevelSHOOPDAWHOOP       float64 `json:"lastLevel_SHOOP_DA_WHOOP"`
				LossesMonthlyA              float64 `json:"losses_monthly_a"`
				LastLevelSANIC              float64 `json:"lastLevel_SANIC"`
				FRIENDSGamesDay             float64 `json:"FRIENDS_gamesDay"`
				FRIENDSFirstGame            float64 `json:"FRIENDS_firstGame"`
				FriendLosses                float64 `json:"friend_losses"`
				FriendWins                  float64 `json:"friend_wins"`
				FriendWinsNormal            float64 `json:"friend_wins_normal"`
				Assists                     float64 `json:"assists"`
				AssistsNormal               float64 `json:"assists_normal"`
				FriendLossesNormal          float64 `json:"friend_losses_normal"`
				HeroLevelBoosterActive      struct {
					Key        string  `json:"key"`
					Multiplier float64 `json:"multiplier"`
					Value      float64 `json:"value"`
					Plays      float64 `json:"plays"`
				} `json:"hero_level_booster_active"`
				ONEVJUANFirstGame   float64 `json:"ONE_V_JUAN_firstGame"`
				ONEVJUANGamesDay    float64 `json:"ONE_V_JUAN_gamesDay"`
				OneVOneLossesNormal float64 `json:"one_v_one_losses_normal"`
				OneVOneLosses       float64 `json:"one_v_one_losses"`
				OneVOneWins         float64 `json:"one_v_one_wins"`
				OneVOneWinsNormal   float64 `json:"one_v_one_wins_normal"`
			} `json:"SuperSmash"`
			SpeedUHC struct {
				ActiveKitINSANE                      string   `json:"activeKit_INSANE"`
				Coins                                float64  `json:"coins"`
				WinstreakKitBasicInsaneDefault       float64  `json:"winstreak_kit_basic_insane_default"`
				GamesSolo                            float64  `json:"games_solo"`
				KillsMonthlyA                        float64  `json:"kills_monthly_a"`
				Wins                                 float64  `json:"wins"`
				WinsMasteryWildSpecialist            float64  `json:"wins_mastery_wild_specialist"`
				WinstreakInsane                      float64  `json:"winstreak_insane"`
				KillstreakInsane                     float64  `json:"killstreak_insane"`
				SurvivedPlayersSolo                  float64  `json:"survived_players_solo"`
				WinsInsane                           float64  `json:"wins_insane"`
				KillsMasteryWildSpecialist           float64  `json:"kills_mastery_wild_specialist"`
				KillsWeeklyB                         float64  `json:"kills_weekly_b"`
				Kills                                float64  `json:"kills"`
				KillstreakSolo                       float64  `json:"killstreak_solo"`
				TearsGathered                        float64  `json:"tears_gathered"`
				Killstreak                           float64  `json:"killstreak"`
				KillstreakKitBasicInsaneDefault      float64  `json:"killstreak_kit_basic_insane_default"`
				ItemsEnchanted                       float64  `json:"items_enchanted"`
				WinstreakSolo                        float64  `json:"winstreak_solo"`
				WinsKitBasicInsaneDefault            float64  `json:"wins_kit_basic_insane_default"`
				Winstreak                            float64  `json:"winstreak"`
				BlocksBroken                         float64  `json:"blocks_broken"`
				Tears                                float64  `json:"tears"`
				KillsSolo                            float64  `json:"kills_solo"`
				KillsInsane                          float64  `json:"kills_insane"`
				WinsSolo                             float64  `json:"wins_solo"`
				Games                                float64  `json:"games"`
				SurvivedPlayers                      float64  `json:"survived_players"`
				KillsSoloInsane                      float64  `json:"kills_solo_insane"`
				SurvivedPlayersInsane                float64  `json:"survived_players_insane"`
				GamesInsane                          float64  `json:"games_insane"`
				GamesKitBasicInsaneDefault           float64  `json:"games_kit_basic_insane_default"`
				SurvivedPlayersKitBasicInsaneDefault float64  `json:"survived_players_kit_basic_insane_default"`
				WinsSoloInsane                       float64  `json:"wins_solo_insane"`
				BlocksPlaced                         float64  `json:"blocks_placed"`
				KillsKitBasicInsaneDefault           float64  `json:"kills_kit_basic_insane_default"`
				SurvivedPlayersNormal                float64  `json:"survived_players_normal"`
				WinsNormal                           float64  `json:"wins_normal"`
				KillsSoloNormal                      float64  `json:"kills_solo_normal"`
				GamesNormal                          float64  `json:"games_normal"`
				WinsKitBasicNormalDefault            float64  `json:"wins_kit_basic_normal_default"`
				WinsSoloNormal                       float64  `json:"wins_solo_normal"`
				KillstreakKitBasicNormalDefault      float64  `json:"killstreak_kit_basic_normal_default"`
				WinstreakKitBasicNormalDefault       float64  `json:"winstreak_kit_basic_normal_default"`
				WinstreakNormal                      float64  `json:"winstreak_normal"`
				SurvivedPlayersKitBasicNormalDefault float64  `json:"survived_players_kit_basic_normal_default"`
				GamesKitBasicNormalDefault           float64  `json:"games_kit_basic_normal_default"`
				KillstreakNormal                     float64  `json:"killstreak_normal"`
				KillsNormal                          float64  `json:"kills_normal"`
				KillsKitBasicNormalDefault           float64  `json:"kills_kit_basic_normal_default"`
				HighestWinstreak                     float64  `json:"highestWinstreak"`
				HighestKillstreak                    float64  `json:"highestKillstreak"`
				WinStreak                            float64  `json:"win_streak"`
				Losses                               float64  `json:"losses"`
				DeathsSolo                           float64  `json:"deaths_solo"`
				DeathsNormal                         float64  `json:"deaths_normal"`
				LossesSoloNormal                     float64  `json:"losses_solo_normal"`
				Quits                                float64  `json:"quits"`
				DeathsSoloNormal                     float64  `json:"deaths_solo_normal"`
				Deaths                               float64  `json:"deaths"`
				DeathsKitBasicNormalDefault          float64  `json:"deaths_kit_basic_normal_default"`
				DeathsMasteryWildSpecialist          float64  `json:"deaths_mastery_wild_specialist"`
				LossesMasteryWildSpecialist          float64  `json:"losses_mastery_wild_specialist"`
				LossesKitBasicNormalDefault          float64  `json:"losses_kit_basic_normal_default"`
				LossesNormal                         float64  `json:"losses_normal"`
				LossesSolo                           float64  `json:"losses_solo"`
				ArrowsShot                           float64  `json:"arrows_shot"`
				ArrowsHit                            float64  `json:"arrows_hit"`
				Packages                             []string `json:"packages"`
				TearWellUses                         float64  `json:"tearWellUses"`
				FoundCOMMON                          float64  `json:"found_COMMON"`
				DeathsSoloInsane                     float64  `json:"deaths_solo_insane"`
				LossesInsane                         float64  `json:"losses_insane"`
				LossesSoloInsane                     float64  `json:"losses_solo_insane"`
				DeathsInsane                         float64  `json:"deaths_insane"`
				DeathsKitBasicInsaneDefault          float64  `json:"deaths_kit_basic_insane_default"`
				LossesKitBasicInsaneDefault          float64  `json:"losses_kit_basic_insane_default"`
				FirstJoinLobby                       bool     `json:"firstJoinLobby"`
				FirstJoinLobbyInt                    float64  `json:"firstJoinLobbyInt"`
				NormalArrowRecovery                  float64  `json:"normal_arrow_recovery"`
				KillsKitBasicInsaneMiner             float64  `json:"kills_kit_basic_insane_miner"`
				KillsWeeklyA                         float64  `json:"kills_weekly_a"`
				GamesKitBasicInsaneMiner             float64  `json:"games_kit_basic_insane_miner"`
				KillstreakKitBasicInsaneMiner        float64  `json:"killstreak_kit_basic_insane_miner"`
				SurvivedPlayersKitBasicInsaneMiner   float64  `json:"survived_players_kit_basic_insane_miner"`
				WinstreakKitBasicInsaneMiner         float64  `json:"winstreak_kit_basic_insane_miner"`
				WinsKitBasicInsaneMiner              float64  `json:"wins_kit_basic_insane_miner"`
				AssistsKitBasicInsaneMiner           float64  `json:"assists_kit_basic_insane_miner"`
				AssistsInsane                        float64  `json:"assists_insane"`
				AssistsSolo                          float64  `json:"assists_solo"`
				Assists                              float64  `json:"assists"`
				Salt                                 float64  `json:"salt"`
				ExtraWheels                          float64  `json:"extra_wheels"`
				NormalMonsterTamer                   float64  `json:"normal_monster_tamer"`
				FoundRARE                            float64  `json:"found_RARE"`
				InsaneSwimmingChampion               float64  `json:"insane_swimming_champion"`
				InsaneKnowledge                      float64  `json:"insane_knowledge"`
				NormalSwimmingChampion               float64  `json:"normal_swimming_champion"`
				InsanePortalProtection               float64  `json:"insane_portal_protection"`
				MasteryFortune                       float64  `json:"mastery_fortune"`
				FoundLEGENDARY                       float64  `json:"found_LEGENDARY"`
				ActiveMasterPerk                     string   `json:"activeMasterPerk"`
				GamesKitBasicInsaneArcher            float64  `json:"games_kit_basic_insane_archer"`
				WinsMasteryFortune                   float64  `json:"wins_mastery_fortune"`
				KillstreakKitBasicInsaneArcher       float64  `json:"killstreak_kit_basic_insane_archer"`
				KillsMasteryFortune                  float64  `json:"kills_mastery_fortune"`
				WinsKitBasicInsaneArcher             float64  `json:"wins_kit_basic_insane_archer"`
				KillsKitBasicInsaneArcher            float64  `json:"kills_kit_basic_insane_archer"`
				SurvivedPlayersKitBasicInsaneArcher  float64  `json:"survived_players_kit_basic_insane_archer"`
				WinstreakKitBasicInsaneArcher        float64  `json:"winstreak_kit_basic_insane_archer"`
				ActiveKitNORMAL                      string   `json:"activeKit_NORMAL"`
				LossesMasteryFortune                 float64  `json:"losses_mastery_fortune"`
				DeathsKitBasicNormalScout            float64  `json:"deaths_kit_basic_normal_scout"`
				SurvivedPlayersKitBasicNormalScout   float64  `json:"survived_players_kit_basic_normal_scout"`
				GamesKitBasicNormalScout             float64  `json:"games_kit_basic_normal_scout"`
				DeathsMasteryFortune                 float64  `json:"deaths_mastery_fortune"`
				KillsKitBasicNormalScout             float64  `json:"kills_kit_basic_normal_scout"`
				LossesKitBasicNormalScout            float64  `json:"losses_kit_basic_normal_scout"`
				KillstreakKitBasicNormalScout        float64  `json:"killstreak_kit_basic_normal_scout"`
				WinsKitBasicNormalScout              float64  `json:"wins_kit_basic_normal_scout"`
				WinstreakKitBasicNormalScout         float64  `json:"winstreak_kit_basic_normal_scout"`
				LossesKitBasicInsaneHealer           float64  `json:"losses_kit_basic_insane_healer"`
				DeathsKitBasicInsaneHealer           float64  `json:"deaths_kit_basic_insane_healer"`
				SurvivedPlayersKitBasicInsaneHealer  float64  `json:"survived_players_kit_basic_insane_healer"`
				WinsKitBasicInsaneHealer             float64  `json:"wins_kit_basic_insane_healer"`
				KillsKitBasicInsaneHealer            float64  `json:"kills_kit_basic_insane_healer"`
				GamesKitBasicInsaneHealer            float64  `json:"games_kit_basic_insane_healer"`
				WinstreakKitBasicInsaneHealer        float64  `json:"winstreak_kit_basic_insane_healer"`
				KillstreakKitBasicInsaneHealer       float64  `json:"killstreak_kit_basic_insane_healer"`
				DeathsTeam                           float64  `json:"deaths_team"`
				DeathsTeamNormal                     float64  `json:"deaths_team_normal"`
				KillsTeamNormal                      float64  `json:"kills_team_normal"`
				SurvivedPlayersTeam                  float64  `json:"survived_players_team"`
				KillsTeam                            float64  `json:"kills_team"`
				LossesTeam                           float64  `json:"losses_team"`
				LossesTeamNormal                     float64  `json:"losses_team_normal"`
				KillsMonthlyB                        float64  `json:"kills_monthly_b"`
				GamesTeam                            float64  `json:"games_team"`
				DeathsTeamInsane                     float64  `json:"deaths_team_insane"`
				GamesKitBasicInsaneScout             float64  `json:"games_kit_basic_insane_scout"`
				WinstreakKitBasicInsaneScout         float64  `json:"winstreak_kit_basic_insane_scout"`
				DeathsKitBasicInsaneScout            float64  `json:"deaths_kit_basic_insane_scout"`
				WinsKitBasicInsaneScout              float64  `json:"wins_kit_basic_insane_scout"`
				WinstreakTeam                        float64  `json:"winstreak_team"`
				WinsTeamInsane                       float64  `json:"wins_team_insane"`
				WinsTeam                             float64  `json:"wins_team"`
				SurvivedPlayersKitBasicInsaneScout   float64  `json:"survived_players_kit_basic_insane_scout"`
				AssistsNormal                        float64  `json:"assists_normal"`
				AssistsKitBasicNormalScout           float64  `json:"assists_kit_basic_normal_scout"`
				InsaneTank                           float64  `json:"insane_tank"`
				SurvivedPlayersKitBasicNormalArcher  float64  `json:"survived_players_kit_basic_normal_archer"`
				WinstreakKitBasicNormalArcher        float64  `json:"winstreak_kit_basic_normal_archer"`
				WinsKitBasicNormalArcher             float64  `json:"wins_kit_basic_normal_archer"`
				KillstreakKitBasicNormalArcher       float64  `json:"killstreak_kit_basic_normal_archer"`
				GamesKitBasicNormalArcher            float64  `json:"games_kit_basic_normal_archer"`
				KillsKitBasicNormalArcher            float64  `json:"kills_kit_basic_normal_archer"`
				LossesKitBasicInsaneArcher           float64  `json:"losses_kit_basic_insane_archer"`
				DeathsKitBasicInsaneArcher           float64  `json:"deaths_kit_basic_insane_archer"`
				InsaneNotoriety                      float64  `json:"insane_notoriety"`
				NormalTelekinesis                    float64  `json:"normal_telekinesis"`
				NormalNourishment                    float64  `json:"normal_nourishment"`
				NormalNoMercy                        float64  `json:"normal_no_mercy"`
				NormalLowGravity                     float64  `json:"normal_low_gravity"`
				InsaneNoMercy                        float64  `json:"insane_no_mercy"`
				NormalTenacity                       float64  `json:"normal_tenacity"`
				InsaneAdrenaline                     float64  `json:"insane_adrenaline"`
				NormalVitamins                       float64  `json:"normal_vitamins"`
				MasteryHuntsman                      float64  `json:"mastery_huntsman"`
				NormalColdBlood                      float64  `json:"normal_cold_blood"`
				InsaneLowGravity                     float64  `json:"insane_low_gravity"`
				InsaneMasterBreeder                  float64  `json:"insane_master_breeder"`
				InsaneTelekinesis                    float64  `json:"insane_telekinesis"`
				NormalBowFlex                        float64  `json:"normal_bow_flex"`
				NormalMedicine                       float64  `json:"normal_medicine"`
				InsaneVitamins                       float64  `json:"insane_vitamins"`
				Score                                float64  `json:"score"`
				MovedOver                            bool     `json:"movedOver"`
			} `json:"SpeedUHC"`
			SkyClash struct {
				CardPacks                      float64  `json:"card_packs"`
				ActiveClass                    float64  `json:"active_class"`
				Killstreak                     float64  `json:"killstreak"`
				HighestKillstreak              float64  `json:"highestKillstreak"`
				Playstreak                     float64  `json:"playstreak"`
				LongestBowShotKitSwordsman     float64  `json:"longest_bow_shot_kit_swordsman"`
				LongestBowShot                 float64  `json:"longest_bow_shot"`
				WinStreak                      float64  `json:"win_streak"`
				BowShotsKitSwordsman           float64  `json:"bow_shots_kit_swordsman"`
				DeathsKitSwordsman             float64  `json:"deaths_kit_swordsman"`
				Losses                         float64  `json:"losses"`
				Kills                          float64  `json:"kills"`
				KillsKitSwordsman              float64  `json:"kills_kit_swordsman"`
				GamesPlayed                    float64  `json:"games_played"`
				BowHitsKitSwordsman            float64  `json:"bow_hits_kit_swordsman"`
				BowShots                       float64  `json:"bow_shots"`
				LossesSolo                     float64  `json:"losses_solo"`
				MostKillsGameKitSwordsman      float64  `json:"most_kills_game_kit_swordsman"`
				EnderchestsOpened              float64  `json:"enderchests_opened"`
				EnderchestsOpenedKitSwordsman  float64  `json:"enderchests_opened_kit_swordsman"`
				BowHits                        float64  `json:"bow_hits"`
				DeathsSolo                     float64  `json:"deaths_solo"`
				MostKillsGame                  float64  `json:"most_kills_game"`
				Quits                          float64  `json:"quits"`
				MeleeKills                     float64  `json:"melee_kills"`
				KillsSolo                      float64  `json:"kills_solo"`
				KillsMonthlyB                  float64  `json:"kills_monthly_b"`
				Deaths                         float64  `json:"deaths"`
				GamesPlayedKitSwordsman        float64  `json:"games_played_kit_swordsman"`
				Coins                          float64  `json:"coins"`
				MeleeKillsKitSwordsman         float64  `json:"melee_kills_kit_swordsman"`
				KillsWeeklyB                   float64  `json:"kills_weekly_b"`
				PlayStreak                     float64  `json:"play_streak"`
				FastestWinSoloKitSwordsman     float64  `json:"fastest_win_solo_kit_swordsman"`
				FastestWinSolo                 float64  `json:"fastest_win_solo"`
				VoidKills                      float64  `json:"void_kills"`
				Games                          float64  `json:"games"`
				Wins                           float64  `json:"wins"`
				VoidKillsKitSwordsman          float64  `json:"void_kills_kit_swordsman"`
				SoloWinsKitSwordsman           float64  `json:"solo_wins_kit_swordsman"`
				SoloWins                       float64  `json:"solo_wins"`
				WinsSolo                       float64  `json:"wins_solo"`
				SpawnAtWitch                   float64  `json:"spawn_at_witch"`
				AssistsKitSwordsman            float64  `json:"assists_kit_swordsman"`
				Assists                        float64  `json:"assists"`
				Packages                       []string `json:"packages"`
				Class2                         string   `json:"class_2"`
				AssistsKitAssassin             float64  `json:"assists_kit_assassin"`
				GamesPlayedKitAssassin         float64  `json:"games_played_kit_assassin"`
				DeathsKitAssassin              float64  `json:"deaths_kit_assassin"`
				MostKillsGameKitAssassin       float64  `json:"most_kills_game_kit_assassin"`
				MeleeKillsKitAssassin          float64  `json:"melee_kills_kit_assassin"`
				VoidKillsKitAssassin           float64  `json:"void_kills_kit_assassin"`
				KillsKitAssassin               float64  `json:"kills_kit_assassin"`
				PerkSharpenedSword             float64  `json:"perk_sharpened_sword"`
				PerkBlazingArrows              float64  `json:"perk_blazing_arrows"`
				PerkSharpenedSwordNew          bool     `json:"perk_sharpened_sword_new"`
				PerkCreeperNew                 bool     `json:"perk_creeper_new"`
				PerkBlazingArrowsNew           bool     `json:"perk_blazing_arrows_new"`
				PerkCreeper                    float64  `json:"perk_creeper"`
				PerkSupplyDrop                 float64  `json:"perk_supply_drop"`
				PerkVoidWarranty               float64  `json:"perk_void_warranty"`
				PerkVoidWarrantyNew            bool     `json:"perk_void_warranty_new"`
				PerkGuardian                   float64  `json:"perk_guardian"`
				PerkGuardianNew                bool     `json:"perk_guardian_new"`
				PerkSupplyDropNew              bool     `json:"perk_supply_drop_new"`
				PerkHonedBow                   float64  `json:"perk_honed_bow"`
				PerkSugarRushNew               bool     `json:"perk_sugar_rush_new"`
				PerkSugarRush                  float64  `json:"perk_sugar_rush"`
				PerkHonedBowNew                bool     `json:"perk_honed_bow_new"`
				PerkMarksmanDuplicates         float64  `json:"perk_marksman_duplicates"`
				LongestBowShotKitAssassin      float64  `json:"longest_bow_shot_kit_assassin"`
				LongestBowKillKitAssassin      float64  `json:"longest_bow_kill_kit_assassin"`
				LongestBowKill                 float64  `json:"longest_bow_kill"`
				BowKills                       float64  `json:"bow_kills"`
				BowKillsKitAssassin            float64  `json:"bow_kills_kit_assassin"`
				BowHitsKitAssassin             float64  `json:"bow_hits_kit_assassin"`
				EnderchestsOpenedKitAssassin   float64  `json:"enderchests_opened_kit_assassin"`
				BowShotsKitAssassin            float64  `json:"bow_shots_kit_assassin"`
				FastestWinSoloKitAssassin      float64  `json:"fastest_win_solo_kit_assassin"`
				SoloWinsKitAssassin            float64  `json:"solo_wins_kit_assassin"`
				MobsKilledKitAssassin          float64  `json:"mobs_killed_kit_assassin"`
				MobsKilled                     float64  `json:"mobs_killed"`
				PerkSugarRushDuplicates        float64  `json:"perk_sugar_rush_duplicates"`
				PerkNutritious                 float64  `json:"perk_nutritious"`
				PerkSkeletonJockey             float64  `json:"perk_skeleton_jockey"`
				PerkNutritiousNew              bool     `json:"perk_nutritious_new"`
				PerkSkeletonJockeyNew          bool     `json:"perk_skeleton_jockey_new"`
				PerkHitAndRunDuplicates        float64  `json:"perk_hit_and_run_duplicates"`
				PerkSkeletonJockeyDuplicates   float64  `json:"perk_skeleton_jockey_duplicates"`
				PerkHonedBowDuplicates         float64  `json:"perk_honed_bow_duplicates"`
				PerkHeartyStartDuplicates      float64  `json:"perk_hearty_start_duplicates"`
				PerkFruitFinderNew             bool     `json:"perk_fruit_finder_new"`
				PerkSnowGolem                  float64  `json:"perk_snow_golem"`
				PerkSnowGolemNew               bool     `json:"perk_snow_golem_new"`
				PerkFruitFinder                float64  `json:"perk_fruit_finder"`
				PerkBlastProtectionNew         bool     `json:"perk_blast_protection_new"`
				PerkEnergyDrink                float64  `json:"perk_energy_drink"`
				PerkEnergyDrinkNew             bool     `json:"perk_energy_drink_new"`
				PerkBlastProtection            float64  `json:"perk_blast_protection"`
				PerkWitch                      float64  `json:"perk_witch"`
				PerkWitchNew                   bool     `json:"perk_witch_new"`
				PerkBlastProtectionDuplicates  float64  `json:"perk_blast_protection_duplicates"`
				PerkInvisibility               float64  `json:"perk_invisibility"`
				PerkInvisibilityNew            bool     `json:"perk_invisibility_new"`
				PerkFruitFinderDuplicates      float64  `json:"perk_fruit_finder_duplicates"`
				FastestWinFourTeams            float64  `json:"fastest_win_four_teams"`
				FastestWinFourTeamsKitAssassin float64  `json:"fastest_win_four_teams_kit_assassin"`
				KillsFourTeams                 float64  `json:"kills_four_teams"`
				FourTeamsWinsKitAssassin       float64  `json:"four_teams_wins_kit_assassin"`
				FourTeamsWins                  float64  `json:"four_teams_wins"`
				WinsFourTeams                  float64  `json:"wins_four_teams"`
				DeathsFourTeams                float64  `json:"deaths_four_teams"`
				LossesFourTeams                float64  `json:"losses_four_teams"`
				PerkSupplyDropDuplicates       float64  `json:"perk_supply_drop_duplicates"`
				PerkPearlAbsorptionNew         bool     `json:"perk_pearl_absorption_new"`
				PerkSharpenedSwordDuplicates   float64  `json:"perk_sharpened_sword_duplicates"`
				PerkPearlAbsorption            float64  `json:"perk_pearl_absorption"`
				PerkPearlAbsorptionDuplicates  float64  `json:"perk_pearl_absorption_duplicates"`
				PerkEndlessQuiverNew           bool     `json:"perk_endless_quiver_new"`
				PerkEndlessQuiver              float64  `json:"perk_endless_quiver"`
				PerkEnergyDrinkDuplicates      float64  `json:"perk_energy_drink_duplicates"`
				PerkGuardianDuplicates         float64  `json:"perk_guardian_duplicates"`
				PerkEndlessQuiverDuplicates    float64  `json:"perk_endless_quiver_duplicates"`
				PerkDamagePotion               float64  `json:"perk_damage_potion"`
				PerkDamagePotionNew            bool     `json:"perk_damage_potion_new"`
				PerkPacify                     float64  `json:"perk_pacify"`
				PerkPacifyNew                  bool     `json:"perk_pacify_new"`
				PerkCreeperDuplicates          float64  `json:"perk_creeper_duplicates"`
				PerkResistantDuplicates        float64  `json:"perk_resistant_duplicates"`
				PerkPacifyDuplicates           float64  `json:"perk_pacify_duplicates"`
				PerkHitAndRun                  float64  `json:"perk_hit_and_run"`
				KillsDoubles                   float64  `json:"kills_doubles"`
				LossesDoubles                  float64  `json:"losses_doubles"`
				DeathsDoubles                  float64  `json:"deaths_doubles"`
				FastestWinDoubles              float64  `json:"fastest_win_doubles"`
				FastestWinDoublesKitAssassin   float64  `json:"fastest_win_doubles_kit_assassin"`
				DoublesWins                    float64  `json:"doubles_wins"`
				WinsDoubles                    float64  `json:"wins_doubles"`
				DoublesWinsKitAssassin         float64  `json:"doubles_wins_kit_assassin"`
				PerkBlazingArrowsDuplicates    float64  `json:"perk_blazing_arrows_duplicates"`
				PerkArrowDeflectionNew         bool     `json:"perk_arrow_deflection_new"`
				PerkArrowDeflection            float64  `json:"perk_arrow_deflection"`
				PerkDamagePotionDuplicates     float64  `json:"perk_damage_potion_duplicates"`
				PerkAlchemy                    float64  `json:"perk_alchemy"`
				PerkInvisibilityDuplicates     float64  `json:"perk_invisibility_duplicates"`
				PerkAlchemyNew                 bool     `json:"perk_alchemy_new"`
				KillsWeeklyA                   float64  `json:"kills_weekly_a"`
				PerkChickenBow                 float64  `json:"perk_chicken_bow"`
				PerkChickenBowNew              bool     `json:"perk_chicken_bow_new"`
				PerkIronGolem                  float64  `json:"perk_iron_golem"`
				PerkIronGolemNew               bool     `json:"perk_iron_golem_new"`
				KitAssassinMinor               float64  `json:"kit_assassin_minor"`
				PerkVoidWarrantyDuplicates     float64  `json:"perk_void_warranty_duplicates"`
				PerkChickenBowDuplicates       float64  `json:"perk_chicken_bow_duplicates"`
				PerkArrowDeflectionDuplicates  float64  `json:"perk_arrow_deflection_duplicates"`
				PerkRampageDuplicates          float64  `json:"perk_rampage_duplicates"`
				PerkAlchemyDuplicates          float64  `json:"perk_alchemy_duplicates"`
				PerkVoidMagnet                 float64  `json:"perk_void_magnet"`
				PerkVoidMagnetNew              bool     `json:"perk_void_magnet_new"`
				PerkVoidMagnetDuplicates       float64  `json:"perk_void_magnet_duplicates"`
				Class0                         string   `json:"class_0"`
				PerkSnowGolemDuplicates        float64  `json:"perk_snow_golem_duplicates"`
				PerkWitchDuplicates            float64  `json:"perk_witch_duplicates"`
				PerkIronGolemDuplicates        float64  `json:"perk_iron_golem_duplicates"`
				PerkTripleshot                 float64  `json:"perk_tripleshot"`
				PerkTripleshotNew              bool     `json:"perk_tripleshot_new"`
				PerkTripleshotDuplicates       float64  `json:"perk_tripleshot_duplicates"`
				PerkRegenerationDuplicates     float64  `json:"perk_regeneration_duplicates"`
				PerkExplosiveBowNew            bool     `json:"perk_explosive_bow_new"`
				PerkExplosiveBow               float64  `json:"perk_explosive_bow"`
				MobKills                       float64  `json:"mob_kills"`
				MobKillsKitAssassin            float64  `json:"mob_kills_kit_assassin"`
				PerkEnderman                   float64  `json:"perk_enderman"`
				PerkEndermanNew                bool     `json:"perk_enderman_new"`
				PerkNutritiousDuplicates       float64  `json:"perk_nutritious_duplicates"`
				PerkExplosiveBowDuplicates     float64  `json:"perk_explosive_bow_duplicates"`
				PerkHeadstart                  float64  `json:"perk_headstart"`
				PerkHeadstartNew               bool     `json:"perk_headstart_new"`
				LossesTeamWar                  float64  `json:"losses_team_war"`
				KillsTeamWar                   float64  `json:"kills_team_war"`
				KillsMonthlyA                  float64  `json:"kills_monthly_a"`
				PerkEndermanDuplicates         float64  `json:"perk_enderman_duplicates"`
				AssassinInventory              struct {
					LEATHERCHESTPLATE0 string `json:"LEATHER_CHESTPLATE:0"`
					ENDERPEARL0        string `json:"ENDER_PEARL:0"`
					SKULLITEM3         string `json:"SKULL_ITEM:3"`
					SKULLITEM2         string `json:"SKULL_ITEM:2"`
					SKULLITEM1         string `json:"SKULL_ITEM:1"`
					COMPASS0           string `json:"COMPASS:0"`
					STONESWORD0        string `json:"STONE_SWORD:0"`
					DIAMONDSWORD0      string `json:"DIAMOND_SWORD:0"`
					LEATHERHELMET0     string `json:"LEATHER_HELMET:0"`
					POTION14           string `json:"POTION:14"`
					POTION12           string `json:"POTION:12"`
					LEATHERLEGGINGS0   string `json:"LEATHER_LEGGINGS:0"`
					LEATHERBOOTS0      string `json:"LEATHER_BOOTS:0"`
				} `json:"assassin_inventory"`
				FastestWinMegaKitAssassin     float64 `json:"fastest_win_mega_kit_assassin"`
				FastestWinMega                float64 `json:"fastest_win_mega"`
				WinsMega                      float64 `json:"wins_mega"`
				KillsMega                     float64 `json:"kills_mega"`
				MegaWinsKitAssassin           float64 `json:"mega_wins_kit_assassin"`
				MegaWins                      float64 `json:"mega_wins"`
				KitAssassinMaster             float64 `json:"kit_assassin_master"`
				PerkLifesteal                 float64 `json:"perk_lifesteal"`
				PerkLifestealNew              bool    `json:"perk_lifesteal_new"`
				PerkMonsterHunterNew          bool    `json:"perk_monster_hunter_new"`
				PerkMonsterHunter             float64 `json:"perk_monster_hunter"`
				PerkHeadHunterNew             bool    `json:"perk_head_hunter_new"`
				PerkHeadHunter                float64 `json:"perk_head_hunter"`
				PerkElvenArcher               float64 `json:"perk_elven_archer"`
				PerkElvenArcherNew            bool    `json:"perk_elven_archer_new"`
				PerkVoidChestNew              bool    `json:"perk_void_chest_new"`
				PerkVoidChest                 float64 `json:"perk_void_chest"`
				PerkDoubleJumpNew             bool    `json:"perk_double_jump_new"`
				PerkDoubleJump                float64 `json:"perk_double_jump"`
				PerkBiggerBangs               float64 `json:"perk_bigger_bangs"`
				PerkBiggerBangsNew            bool    `json:"perk_bigger_bangs_new"`
				PerkResistant                 float64 `json:"perk_resistant"`
				PerkHeartyStart               float64 `json:"perk_hearty_start"`
				PerkRampage                   float64 `json:"perk_rampage"`
				PerkElvenArcherDuplicates     float64 `json:"perk_elven_archer_duplicates"`
				PerkMonsterHunterDuplicates   float64 `json:"perk_monster_hunter_duplicates"`
				PerkNuclearSolutionNew        bool    `json:"perk_nuclear_solution_new"`
				PerkNuclearSolution           float64 `json:"perk_nuclear_solution"`
				PerkMarksman                  float64 `json:"perk_marksman"`
				PerkRegeneration              float64 `json:"perk_regeneration"`
				PerkBountyHunterNew           bool    `json:"perk_bounty_hunter_new"`
				PerkBountyHunter              float64 `json:"perk_bounty_hunter"`
				PerkLifestealDuplicates       float64 `json:"perk_lifesteal_duplicates"`
				PerkMushroomAura              float64 `json:"perk_mushroom_aura"`
				PerkMushroomAuraNew           bool    `json:"perk_mushroom_aura_new"`
				PerkBatShieldNew              bool    `json:"perk_bat_shield_new"`
				PerkBatShield                 float64 `json:"perk_bat_shield"`
				PerkMushroomAuraDuplicates    float64 `json:"perk_mushroom_aura_duplicates"`
				PerkBatShieldDuplicates       float64 `json:"perk_bat_shield_duplicates"`
				PerkWingedBootsNew            bool    `json:"perk_winged_boots_new"`
				PerkNuclearSolutionDuplicates float64 `json:"perk_nuclear_solution_duplicates"`
				PerkWingedBoots               float64 `json:"perk_winged_boots"`
				PerkBountyHunterDuplicates    float64 `json:"perk_bounty_hunter_duplicates"`
				PerkVoidChestDuplicates       float64 `json:"perk_void_chest_duplicates"`
				PerkBiggerBangsDuplicates     float64 `json:"perk_bigger_bangs_duplicates"`
				PerkNoChestChallenge          float64 `json:"perk_no_chest_challenge"`
				PerkNoChestChallengeNew       bool    `json:"perk_no_chest_challenge_new"`
				PerkFlowerPowerNew            bool    `json:"perk_flower_power_new"`
				PerkFlowerPower               float64 `json:"perk_flower_power"`
				PerkArcherChallengeNew        bool    `json:"perk_archer_challenge_new"`
				PerkArcherChallenge           float64 `json:"perk_archer_challenge"`
				PerkDoubleJumpDuplicates      float64 `json:"perk_double_jump_duplicates"`
				PerkFlowerPowerDuplicates     float64 `json:"perk_flower_power_duplicates"`
				FastestWinTeamWarKitAssassin  float64 `json:"fastest_win_team_war_kit_assassin"`
				FastestWinTeamWar             float64 `json:"fastest_win_team_war"`
				DeathsPerkHeadstart           float64 `json:"deaths_perk_headstart"`
				DeathsPerkPearlAbsorption     float64 `json:"deaths_perk_pearl_absorption"`
				KillsPerkSharpenedSword       float64 `json:"kills_perk_sharpened_sword"`
				DeathsPerkSharpenedSword      float64 `json:"deaths_perk_sharpened_sword"`
				WinsPerkSharpenedSword        float64 `json:"wins_perk_sharpened_sword"`
				WinsTeamWar                   float64 `json:"wins_team_war"`
				WinsPerkPearlAbsorption       float64 `json:"wins_perk_pearl_absorption"`
				KillsPerkHeadstart            float64 `json:"kills_perk_headstart"`
				WinsPerkHeadstart             float64 `json:"wins_perk_headstart"`
				TeamWarWinsKitAssassin        float64 `json:"team_war_wins_kit_assassin"`
				TeamWarWins                   float64 `json:"team_war_wins"`
				KillsPerkPearlAbsorption      float64 `json:"kills_perk_pearl_absorption"`
				DeathsTeamWar                 float64 `json:"deaths_team_war"`
				LossesPerkHeadstart           float64 `json:"losses_perk_headstart"`
				LossesPerkPearlAbsorption     float64 `json:"losses_perk_pearl_absorption"`
				LossesPerkSharpenedSword      float64 `json:"losses_perk_sharpened_sword"`
				SwordsmanInventory            struct {
					SKULLITEM3     string `json:"SKULL_ITEM:3"`
					POTION9        string `json:"POTION:9"`
					SKULLITEM2     string `json:"SKULL_ITEM:2"`
					SKULLITEM1     string `json:"SKULL_ITEM:1"`
					IRONSWORD0     string `json:"IRON_SWORD:0"`
					COMPASS0       string `json:"COMPASS:0"`
					LEATHERHELMET0 string `json:"LEATHER_HELMET:0"`
				} `json:"swordsman_inventory"`
				ClassesUnlocked         float64 `json:"classes_unlocked"`
				Class3                  string  `json:"class_3"`
				LossesPerkRampage       float64 `json:"losses_perk_rampage"`
				DeathsPerkRampage       float64 `json:"deaths_perk_rampage"`
				LossesPerkHitAndRun     float64 `json:"losses_perk_hit_and_run"`
				DeathsPerkHitAndRun     float64 `json:"deaths_perk_hit_and_run"`
				DeathsPerkRegeneration  float64 `json:"deaths_perk_regeneration"`
				LossesPerkRegeneration  float64 `json:"losses_perk_regeneration"`
				MobsSpawnedKitSwordsman float64 `json:"mobs_spawned_kit_swordsman"`
				MobsSpawned             float64 `json:"mobs_spawned"`
				NecromancerInventory    struct {
					LEATHERCHESTPLATE0 string `json:"LEATHER_CHESTPLATE:0"`
					CHAINMAILLEGGINGS0 string `json:"CHAINMAIL_LEGGINGS:0"`
					POTION12           string `json:"POTION:12"`
					ARROW0             string `json:"ARROW:0"`
					BOW0               string `json:"BOW:0"`
					MONSTEREGG51       string `json:"MONSTER_EGG:51"`
					SKULLITEM3         string `json:"SKULL_ITEM:3"`
					SKULLITEM2         string `json:"SKULL_ITEM:2"`
					SKULLITEM1         string `json:"SKULL_ITEM:1"`
					COMPASS0           string `json:"COMPASS:0"`
				} `json:"necromancer_inventory"`
				LossesPerkWitch                 float64 `json:"losses_perk_witch"`
				LossesPerkSnowGolem             float64 `json:"losses_perk_snow_golem"`
				MobsKilledKitNecromancer        float64 `json:"mobs_killed_kit_necromancer"`
				DeathsKitNecromancer            float64 `json:"deaths_kit_necromancer"`
				DeathsPerkCreeper               float64 `json:"deaths_perk_creeper"`
				EnderchestsOpenedKitNecromancer float64 `json:"enderchests_opened_kit_necromancer"`
				DeathsPerkWitch                 float64 `json:"deaths_perk_witch"`
				DeathsPerkSnowGolem             float64 `json:"deaths_perk_snow_golem"`
				MobsSpawnedKitNecromancer       float64 `json:"mobs_spawned_kit_necromancer"`
				GamesPlayedKitNecromancer       float64 `json:"games_played_kit_necromancer"`
				LossesPerkCreeper               float64 `json:"losses_perk_creeper"`
				BowShotsKitNecromancer          float64 `json:"bow_shots_kit_necromancer"`
				MobsSpawnedKitAssassin          float64 `json:"mobs_spawned_kit_assassin"`
				ArcherInventory                 struct {
					ARROW0               string `json:"ARROW:0"`
					BOW0                 string `json:"BOW:0"`
					CHAINMAILBOOTS0      string `json:"CHAINMAIL_BOOTS:0"`
					CHAINMAILCHESTPLATE0 string `json:"CHAINMAIL_CHESTPLATE:0"`
					SKULLITEM3           string `json:"SKULL_ITEM:3"`
					SKULLITEM2           string `json:"SKULL_ITEM:2"`
					SKULLITEM1           string `json:"SKULL_ITEM:1"`
					COMPASS0             string `json:"COMPASS:0"`
				} `json:"archer_inventory"`
				LongestBowShotKitArcher    float64 `json:"longest_bow_shot_kit_archer"`
				KillsPerkHonedBow          float64 `json:"kills_perk_honed_bow"`
				BowHitsKitArcher           float64 `json:"bow_hits_kit_archer"`
				KillsKitArcher             float64 `json:"kills_kit_archer"`
				MobKillsKitArcher          float64 `json:"mob_kills_kit_archer"`
				KillsPerkEndlessQuiver     float64 `json:"kills_perk_endless_quiver"`
				BowShotsKitArcher          float64 `json:"bow_shots_kit_archer"`
				GamesPlayedKitArcher       float64 `json:"games_played_kit_archer"`
				DeathsPerkEndlessQuiver    float64 `json:"deaths_perk_endless_quiver"`
				KillsPerkExplosiveBow      float64 `json:"kills_perk_explosive_bow"`
				DeathsPerkExplosiveBow     float64 `json:"deaths_perk_explosive_bow"`
				MostKillsGameKitArcher     float64 `json:"most_kills_game_kit_archer"`
				MobsKilledKitArcher        float64 `json:"mobs_killed_kit_archer"`
				DeathsKitArcher            float64 `json:"deaths_kit_archer"`
				EnderchestsOpenedKitArcher float64 `json:"enderchests_opened_kit_archer"`
				DeathsPerkHonedBow         float64 `json:"deaths_perk_honed_bow"`
				FastestWinSoloKitArcher    float64 `json:"fastest_win_solo_kit_archer"`
				VoidKillsKitArcher         float64 `json:"void_kills_kit_archer"`
				WinsPerkEndlessQuiver      float64 `json:"wins_perk_endless_quiver"`
				SoloWinsKitArcher          float64 `json:"solo_wins_kit_archer"`
				WinsPerkExplosiveBow       float64 `json:"wins_perk_explosive_bow"`
				MeleeKillsKitArcher        float64 `json:"melee_kills_kit_archer"`
				WinsPerkHonedBow           float64 `json:"wins_perk_honed_bow"`
				CutePantsFound             float64 `json:"cute_pants_found"`
				MobsSpawnedKitArcher       float64 `json:"mobs_spawned_kit_archer"`
				CutePantsFoundKitArcher    float64 `json:"cute_pants_found_kit_archer"`
				LossesPerkHonedBow         float64 `json:"losses_perk_honed_bow"`
				LossesPerkExplosiveBow     float64 `json:"losses_perk_explosive_bow"`
				LossesPerkEndlessQuiver    float64 `json:"losses_perk_endless_quiver"`
				WinsPerkEnergyDrink        float64 `json:"wins_perk_energy_drink"`
				KillsPerkWingedBoots       float64 `json:"kills_perk_winged_boots"`
				KillsPerkEnergyDrink       float64 `json:"kills_perk_energy_drink"`
				WinsPerkWingedBoots        float64 `json:"wins_perk_winged_boots"`
			} `json:"SkyClash"`
			Legacy struct {
				Coins                   float64 `json:"coins"`
				TokensDaily             float64 `json:"tokens_daily"`
				TokensLastReceivedStamp float64 `json:"tokens_last_received_stamp"`
				NextTokensSeconds       float64 `json:"next_tokens_seconds"`
				PaintballTokens         float64 `json:"paintball_tokens"`
				TotalTokens             float64 `json:"total_tokens"`
				Tokens                  float64 `json:"tokens"`
				VampirezTokens          float64 `json:"vampirez_tokens"`
				QuakecraftTokens        float64 `json:"quakecraft_tokens"`
				GingerbreadTokens       float64 `json:"gingerbread_tokens"`
				PreferredChannel        string  `json:"preferredChannel"`
			} `json:"Legacy"`
			Bedwars struct {
				BedwarsBoxes                                      float64  `json:"bedwars_boxes"`
				Experience                                        float64  `json:"Experience"`
				FirstJoin7                                        bool     `json:"first_join_7"`
				GamesPlayedBedwars1                               float64  `json:"games_played_bedwars_1"`
				Winstreak                                         float64  `json:"winstreak"`
				FinalDeathsBedwars                                float64  `json:"final_deaths_bedwars"`
				GoldResourcesCollectedBedwars                     float64  `json:"gold_resources_collected_bedwars"`
				VoidDeathsBedwars                                 float64  `json:"void_deaths_bedwars"`
				FourThreeVoidKillsBedwars                         float64  `json:"four_three_void_kills_bedwars"`
				FourThreeEntityAttackKillsBedwars                 float64  `json:"four_three_entity_attack_kills_bedwars"`
				VoidKillsBedwars                                  float64  `json:"void_kills_bedwars"`
				DiamondResourcesCollectedBedwars                  float64  `json:"diamond_resources_collected_bedwars"`
				DeathsBedwars                                     float64  `json:"deaths_bedwars"`
				EntityAttackFinalDeathsBedwars                    float64  `json:"entity_attack_final_deaths_bedwars"`
				EmeraldResourcesCollectedBedwars                  float64  `json:"emerald_resources_collected_bedwars"`
				FourThreeIronResourcesCollectedBedwars            float64  `json:"four_three_iron_resources_collected_bedwars"`
				ResourcesCollectedBedwars                         float64  `json:"resources_collected_bedwars"`
				FourThreeEntityAttackFinalDeathsBedwars           float64  `json:"four_three_entity_attack_final_deaths_bedwars"`
				FallDeathsBedwars                                 float64  `json:"fall_deaths_bedwars"`
				FourThreeGamesPlayedBedwars                       float64  `json:"four_three_games_played_bedwars"`
				GamesPlayedBedwars                                float64  `json:"games_played_bedwars"`
				FourThreeFallDeathsBedwars                        float64  `json:"four_three_fall_deaths_bedwars"`
				FourThreeEmeraldResourcesCollectedBedwars         float64  `json:"four_three_emerald_resources_collected_bedwars"`
				FourThreeVoidDeathsBedwars                        float64  `json:"four_three_void_deaths_bedwars"`
				FourThreeEntityAttackDeathsBedwars                float64  `json:"four_three_entity_attack_deaths_bedwars"`
				FourThreeBedsLostBedwars                          float64  `json:"four_three_beds_lost_bedwars"`
				BedsLostBedwars                                   float64  `json:"beds_lost_bedwars"`
				KillsBedwars                                      float64  `json:"kills_bedwars"`
				FourThreeGoldResourcesCollectedBedwars            float64  `json:"four_three_gold_resources_collected_bedwars"`
				EntityAttackKillsBedwars                          float64  `json:"entity_attack_kills_bedwars"`
				FourThreeDeathsBedwars                            float64  `json:"four_three_deaths_bedwars"`
				FourThreeLossesBedwars                            float64  `json:"four_three_losses_bedwars"`
				EntityAttackDeathsBedwars                         float64  `json:"entity_attack_deaths_bedwars"`
				LossesBedwars                                     float64  `json:"losses_bedwars"`
				FourThreeDiamondResourcesCollectedBedwars         float64  `json:"four_three_diamond_resources_collected_bedwars"`
				IronResourcesCollectedBedwars                     float64  `json:"iron_resources_collected_bedwars"`
				FourThreeKillsBedwars                             float64  `json:"four_three_kills_bedwars"`
				FourThreeFinalDeathsBedwars                       float64  `json:"four_three_final_deaths_bedwars"`
				FourThreeResourcesCollectedBedwars                float64  `json:"four_three_resources_collected_bedwars"`
				FourThreeVoidFinalKillsBedwars                    float64  `json:"four_three_void_final_kills_bedwars"`
				FourThreeWinsBedwars                              float64  `json:"four_three_wins_bedwars"`
				Coins                                             float64  `json:"coins"`
				EntityAttackFinalKillsBedwars                     float64  `json:"entity_attack_final_kills_bedwars"`
				FourThreePermanentItemsPurchasedBedwars           float64  `json:"four_three_permanent _items_purchased_bedwars"`
				FourThreeBedsBrokenBedwars                        float64  `json:"four_three_beds_broken_bedwars"`
				FourThreeEntityAttackFinalKillsBedwars            float64  `json:"four_three_entity_attack_final_kills_bedwars"`
				ItemsPurchasedBedwars                             float64  `json:"items_purchased_bedwars"`
				WinsBedwars                                       float64  `json:"wins_bedwars"`
				FinalKillsBedwars                                 float64  `json:"final_kills_bedwars"`
				VoidFinalKillsBedwars                             float64  `json:"void_final_kills_bedwars"`
				FourThreeFinalKillsBedwars                        float64  `json:"four_three_final_kills_bedwars"`
				FourThreeItemsPurchasedBedwars                    float64  `json:"four_three_items_purchased_bedwars"`
				BedsBrokenBedwars                                 float64  `json:"beds_broken_bedwars"`
				Packages                                          []string `json:"packages"`
				BedwarsBoxCommons                                 float64  `json:"bedwars_box_commons"`
				ChestHistory                                      string   `json:"chest_history"`
				BedwarsBox                                        float64  `json:"bedwars_box"`
				BedwarsBoxRares                                   float64  `json:"bedwars_box_rares"`
				BedwarsBoxEpics                                   float64  `json:"bedwars_box_epics"`
				BedwarsBoxLegendaries                             float64  `json:"bedwars_box_legendaries"`
				FourThreeFallFinalDeathsBedwars                   float64  `json:"four_three_fall_final_deaths_bedwars"`
				FallFinalDeathsBedwars                            float64  `json:"fall_final_deaths_bedwars"`
				EightOneBedsLostBedwars                           float64  `json:"eight_one_beds_lost_bedwars"`
				EightOneGamesPlayedBedwars                        float64  `json:"eight_one_games_played_bedwars"`
				EightOneFinalKillsBedwars                         float64  `json:"eight_one_final_kills_bedwars"`
				EightOneBedsBrokenBedwars                         float64  `json:"eight_one_beds_broken_bedwars"`
				EightOneIronResourcesCollectedBedwars             float64  `json:"eight_one_iron_resources_collected_bedwars"`
				EightOneEmeraldResourcesCollectedBedwars          float64  `json:"eight_one_emerald_resources_collected_bedwars"`
				EightOneItemsPurchasedBedwars                     float64  `json:"eight_one_items_purchased_bedwars"`
				EightOneDiamondResourcesCollectedBedwars          float64  `json:"eight_one_diamond_resources_collected_bedwars"`
				EightOnePermanentItemsPurchasedBedwars            float64  `json:"eight_one_permanent _items_purchased_bedwars"`
				EightOneEntityAttackKillsBedwars                  float64  `json:"eight_one_entity_attack_kills_bedwars"`
				EightOneVoidKillsBedwars                          float64  `json:"eight_one_void_kills_bedwars"`
				EightOneVoidFinalKillsBedwars                     float64  `json:"eight_one_void_final_kills_bedwars"`
				EightOneWinsBedwars                               float64  `json:"eight_one_wins_bedwars"`
				EightOneKillsBedwars                              float64  `json:"eight_one_kills_bedwars"`
				EightOneResourcesCollectedBedwars                 float64  `json:"eight_one_resources_collected_bedwars"`
				EightOneGoldResourcesCollectedBedwars             float64  `json:"eight_one_gold_resources_collected_bedwars"`
				EightOneEntityAttackFinalKillsBedwars             float64  `json:"eight_one_entity_attack_final_kills_bedwars"`
				EightOneFinalDeathsBedwars                        float64  `json:"eight_one_final_deaths_bedwars"`
				EightOneLossesBedwars                             float64  `json:"eight_one_losses_bedwars"`
				VoidFinalDeathsBedwars                            float64  `json:"void_final_deaths_bedwars"`
				EightOneVoidFinalDeathsBedwars                    float64  `json:"eight_one_void_final_deaths_bedwars"`
				EightOneDeathsBedwars                             float64  `json:"eight_one_deaths_bedwars"`
				EightOneEntityAttackDeathsBedwars                 float64  `json:"eight_one_entity_attack_deaths_bedwars"`
				EightTwoBedsBrokenBedwars                         float64  `json:"eight_two_beds_broken_bedwars"`
				EightTwoEmeraldResourcesCollectedBedwars          float64  `json:"eight_two_emerald_resources_collected_bedwars"`
				EightTwoResourcesCollectedBedwars                 float64  `json:"eight_two_resources_collected_bedwars"`
				EightTwoDiamondResourcesCollectedBedwars          float64  `json:"eight_two_diamond_resources_collected_bedwars"`
				EightTwoWinsBedwars                               float64  `json:"eight_two_wins_bedwars"`
				EightTwoKillsBedwars                              float64  `json:"eight_two_kills_bedwars"`
				EightTwoIronResourcesCollectedBedwars             float64  `json:"eight_two_iron_resources_collected_bedwars"`
				EightTwoVoidKillsBedwars                          float64  `json:"eight_two_void_kills_bedwars"`
				EightTwoDeathsBedwars                             float64  `json:"eight_two_deaths_bedwars"`
				EightTwoGamesPlayedBedwars                        float64  `json:"eight_two_games_played_bedwars"`
				EightTwoItemsPurchasedBedwars                     float64  `json:"eight_two_items_purchased_bedwars"`
				EightTwoBedsLostBedwars                           float64  `json:"eight_two_beds_lost_bedwars"`
				EightTwoEntityAttackDeathsBedwars                 float64  `json:"eight_two_entity_attack_deaths_bedwars"`
				EightTwoGoldResourcesCollectedBedwars             float64  `json:"eight_two_gold_resources_collected_bedwars"`
				EightTwoVoidFinalKillsBedwars                     float64  `json:"eight_two_void_final_kills_bedwars"`
				EightTwoFinalKillsBedwars                         float64  `json:"eight_two_final_kills_bedwars"`
				EightTwoVoidDeathsBedwars                         float64  `json:"eight_two_void_deaths_bedwars"`
				EightTwoFallDeathsBedwars                         float64  `json:"eight_two_fall_deaths_bedwars"`
				EightTwoFinalDeathsBedwars                        float64  `json:"eight_two_final_deaths_bedwars"`
				EightTwoEntityAttackFinalDeathsBedwars            float64  `json:"eight_two_entity_attack_final_deaths_bedwars"`
				EightTwoLossesBedwars                             float64  `json:"eight_two_losses_bedwars"`
				EightTwoEntityAttackFinalKillsBedwars             float64  `json:"eight_two_entity_attack_final_kills_bedwars"`
				EightTwoVoidFinalDeathsBedwars                    float64  `json:"eight_two_void_final_deaths_bedwars"`
				EightTwoEntityAttackKillsBedwars                  float64  `json:"eight_two_entity_attack_kills_bedwars"`
				FallKillsBedwars                                  float64  `json:"fall_kills_bedwars"`
				EightTwoFallKillsBedwars                          float64  `json:"eight_two_fall_kills_bedwars"`
				ActiveKillEffect                                  string   `json:"activeKillEffect"`
				ProjectileKillsBedwars                            float64  `json:"projectile_kills_bedwars"`
				EightTwoProjectileKillsBedwars                    float64  `json:"eight_two_projectile_kills_bedwars"`
				EightTwoFallFinalDeathsBedwars                    float64  `json:"eight_two_fall_final_deaths_bedwars"`
				EightTwoProjectileFinalDeathsBedwars              float64  `json:"eight_two_projectile_final_deaths_bedwars"`
				ProjectileFinalDeathsBedwars                      float64  `json:"projectile_final_deaths_bedwars"`
				EightTwoFallFinalKillsBedwars                     float64  `json:"eight_two_fall_final_kills_bedwars"`
				FallFinalKillsBedwars                             float64  `json:"fall_final_kills_bedwars"`
				ProjectileFinalKillsBedwars                       float64  `json:"projectile_final_kills_bedwars"`
				EightTwoProjectileFinalKillsBedwars               float64  `json:"eight_two_projectile_final_kills_bedwars"`
				EightTwoEntityExplosionKillsBedwars               float64  `json:"eight_two_entity_explosion_kills_bedwars"`
				EntityExplosionKillsBedwars                       float64  `json:"entity_explosion_kills_bedwars"`
				ActiveVictoryDance                                string   `json:"activeVictoryDance"`
				ProjectileDeathsBedwars                           float64  `json:"projectile_deaths_bedwars"`
				EightTwoProjectileDeathsBedwars                   float64  `json:"eight_two_projectile_deaths_bedwars"`
				EightOneVoidDeathsBedwars                         float64  `json:"eight_one_void_deaths_bedwars"`
				FourThreeFallKillsBedwars                         float64  `json:"four_three_fall_kills_bedwars"`
				FourThreeVoidFinalDeathsBedwars                   float64  `json:"four_three_void_final_deaths_bedwars"`
				ActiveProjectileTrail                             string   `json:"activeProjectileTrail"`
				ActiveDeathCry                                    string   `json:"activeDeathCry"`
				GlyphStorageNew                                   string   `json:"glyph_storage_new"`
				FourThreeEntityExplosionKillsBedwars              float64  `json:"four_three_entity_explosion_kills_bedwars"`
				FourFourBedsBrokenBedwars                         float64  `json:"four_four_beds_broken_bedwars"`
				FourFourDeathsBedwars                             float64  `json:"four_four_deaths_bedwars"`
				FourFourGoldResourcesCollectedBedwars             float64  `json:"four_four_gold_resources_collected_bedwars"`
				FourFourEntityAttackDeathsBedwars                 float64  `json:"four_four_entity_attack_deaths_bedwars"`
				FourFourIronResourcesCollectedBedwars             float64  `json:"four_four_iron_resources_collected_bedwars"`
				FourFourGamesPlayedBedwars                        float64  `json:"four_four_games_played_bedwars"`
				FourFourWinsBedwars                               float64  `json:"four_four_wins_bedwars"`
				FourFourVoidFinalKillsBedwars                     float64  `json:"four_four_void_final_kills_bedwars"`
				FourFourEntityAttackFinalKillsBedwars             float64  `json:"four_four_entity_attack_final_kills_bedwars"`
				FourFourResourcesCollectedBedwars                 float64  `json:"four_four_resources_collected_bedwars"`
				FourFourItemsPurchasedBedwars                     float64  `json:"four_four_items_purchased_bedwars"`
				FourFourFinalKillsBedwars                         float64  `json:"four_four_final_kills_bedwars"`
				FourFourKillsBedwars                              float64  `json:"four_four_kills_bedwars"`
				FourFourDiamondResourcesCollectedBedwars          float64  `json:"four_four_diamond_resources_collected_bedwars"`
				FourFourEntityAttackKillsBedwars                  float64  `json:"four_four_entity_attack_kills_bedwars"`
				FourFourVoidKillsBedwars                          float64  `json:"four_four_void_kills_bedwars"`
				EightTwoEntityExplosionFinalKillsBedwars          float64  `json:"eight_two_entity_explosion_final_kills_bedwars"`
				EntityExplosionFinalKillsBedwars                  float64  `json:"entity_explosion_final_kills_bedwars"`
				ActiveNPCSkin                                     string   `json:"activeNPCSkin"`
				SprayGlyphField                                   string   `json:"spray_glyph_field"`
				ChestHistoryNew                                   []string `json:"chest_history_new"`
				EightOneEntityAttackFinalDeathsBedwars            float64  `json:"eight_one_entity_attack_final_deaths_bedwars"`
				EightTwoEntityExplosionFinalDeathsBedwars         float64  `json:"eight_two_entity_explosion_final_deaths_bedwars"`
				EntityExplosionFinalDeathsBedwars                 float64  `json:"entity_explosion_final_deaths_bedwars"`
				EntityExplosionDeathsBedwars                      float64  `json:"entity_explosion_deaths_bedwars"`
				EightTwoEntityExplosionDeathsBedwars              float64  `json:"eight_two_entity_explosion_deaths_bedwars"`
				ActiveSprays                                      string   `json:"activeSprays"`
				FourFourPermanentItemsPurchasedBedwars            float64  `json:"four_four_permanent _items_purchased_bedwars"`
				FourFourBedsLostBedwars                           float64  `json:"four_four_beds_lost_bedwars"`
				FourFourEmeraldResourcesCollectedBedwars          float64  `json:"four_four_emerald_resources_collected_bedwars"`
				FourFourVoidDeathsBedwars                         float64  `json:"four_four_void_deaths_bedwars"`
				FourFourFallKillsBedwars                          float64  `json:"four_four_fall_kills_bedwars"`
				FourFourFallFinalKillsBedwars                     float64  `json:"four_four_fall_final_kills_bedwars"`
				FourFourFallDeathsBedwars                         float64  `json:"four_four_fall_deaths_bedwars"`
				FourFourFinalDeathsBedwars                        float64  `json:"four_four_final_deaths_bedwars"`
				FourFourVoidFinalDeathsBedwars                    float64  `json:"four_four_void_final_deaths_bedwars"`
				FourFourLossesBedwars                             float64  `json:"four_four_losses_bedwars"`
				FourFourEntityAttackFinalDeathsBedwars            float64  `json:"four_four_entity_attack_final_deaths_bedwars"`
				FourFourEntityExplosionFinalKillsBedwars          float64  `json:"four_four_entity_explosion_final_kills_bedwars"`
				FourThreeEntityExplosionFinalKillsBedwars         float64  `json:"four_three_entity_explosion_final_kills_bedwars"`
				FourThreeFallFinalKillsBedwars                    float64  `json:"four_three_fall_final_kills_bedwars"`
				FourFourEntityExplosionDeathsBedwars              float64  `json:"four_four_entity_explosion_deaths_bedwars"`
				FourFourProjectileKillsBedwars                    float64  `json:"four_four_projectile_kills_bedwars"`
				FourFourEntityExplosionKillsBedwars               float64  `json:"four_four_entity_explosion_kills_bedwars"`
				FourThreeEntityExplosionDeathsBedwars             float64  `json:"four_three_entity_explosion_deaths_bedwars"`
				FourThreeProjectileDeathsBedwars                  float64  `json:"four_three_projectile_deaths_bedwars"`
				FourFourProjectileDeathsBedwars                   float64  `json:"four_four_projectile_deaths_bedwars"`
				Favourites1                                       string   `json:"favourites_1"`
				BedwarsHalloweenBoxes                             float64  `json:"bedwars_halloween_boxes"`
				BedwarsOpenedChests                               float64  `json:"Bedwars_openedChests"`
				BedwarsOpenedCommons                              float64  `json:"Bedwars_openedCommons"`
				BedwarsOpenedRares                                float64  `json:"Bedwars_openedRares"`
				SpookyOpenAch                                     float64  `json:"spooky_open_ach"`
				BedwarsOpenedLegendaries                          float64  `json:"Bedwars_openedLegendaries"`
				BedwarsOpenedEpics                                float64  `json:"Bedwars_openedEpics"`
				ActiveKillMessages                                string   `json:"activeKillMessages"`
				FreeEventKeyBedwarsHalloweenBoxes2017             bool     `json:"free_event_key_bedwars_halloween_boxes_2017"`
				FourFourProjectileFinalKillsBedwars               float64  `json:"four_four_projectile_final_kills_bedwars"`
				FourFourFallFinalDeathsBedwars                    float64  `json:"four_four_fall_final_deaths_bedwars"`
				FireTickDeathsBedwars                             float64  `json:"fire_tick_deaths_bedwars"`
				EightTwoFireTickDeathsBedwars                     float64  `json:"eight_two_fire_tick_deaths_bedwars"`
				EightOneFallKillsBedwars                          float64  `json:"eight_one_fall_kills_bedwars"`
				EightTwoWrappedPresentResourcesCollectedBedwars   float64  `json:"eight_two_wrapped_present_resources_collected_bedwars"`
				WrappedPresentResourcesCollectedBedwars           float64  `json:"wrapped_present_resources_collected_bedwars"`
				BedwarsChristmasBoxes                             float64  `json:"bedwars_christmas_boxes"`
				FreeEventKeyBedwarsChristmasBoxes2017             bool     `json:"free_event_key_bedwars_christmas_boxes_2017"`
				ActiveGlyph                                       string   `json:"activeGlyph"`
				EightOneWrappedPresentResourcesCollectedBedwars   float64  `json:"eight_one_wrapped_present_resources_collected_bedwars"`
				FourFourWrappedPresentResourcesCollectedBedwars   float64  `json:"four_four_wrapped_present_resources_collected_bedwars"`
				SeenBetaMenu                                      float64  `json:"seen_beta_menu"`
				Favourites2                                       string   `json:"favourites_2"`
				BedwarsLunarBoxes                                 float64  `json:"bedwars_lunar_boxes"`
				FreeEventKeyBedwarsLunarBoxes2018                 bool     `json:"free_event_key_bedwars_lunar_boxes_2018"`
				BedwarsEasterBoxes                                float64  `json:"bedwars_easter_boxes"`
				ActiveBedDestroy                                  string   `json:"activeBedDestroy"`
				FireTickFinalKillsBedwars                         float64  `json:"fire_tick_final_kills_bedwars"`
				EightTwoFireTickFinalKillsBedwars                 float64  `json:"eight_two_fire_tick_final_kills_bedwars"`
				EightTwoUltimateWinstreak                         float64  `json:"eight_two_ultimate_winstreak"`
				EightTwoUltimateVoidKillsBedwars                  float64  `json:"eight_two_ultimate_void_kills_bedwars"`
				EightTwoUltimateResourcesCollectedBedwars         float64  `json:"eight_two_ultimate_resources_collected_bedwars"`
				EightTwoUltimateBedsLostBedwars                   float64  `json:"eight_two_ultimate_beds_lost_bedwars"`
				EightTwoUltimateGamesPlayedBedwars                float64  `json:"eight_two_ultimate_games_played_bedwars"`
				EightTwoUltimateIronResourcesCollectedBedwars     float64  `json:"eight_two_ultimate_iron_resources_collected_bedwars"`
				EightTwoUltimateFinalDeathsBedwars                float64  `json:"eight_two_ultimate_final_deaths_bedwars"`
				EightTwoUltimateLossesBedwars                     float64  `json:"eight_two_ultimate_losses_bedwars"`
				EightTwoUltimateGoldResourcesCollectedBedwars     float64  `json:"eight_two_ultimate_gold_resources_collected_bedwars"`
				EightTwoUltimateKillsBedwars                      float64  `json:"eight_two_ultimate_kills_bedwars"`
				EightTwoUltimateFinalKillsBedwars                 float64  `json:"eight_two_ultimate_final_kills_bedwars"`
				EightTwoUltimateBedsBrokenBedwars                 float64  `json:"eight_two_ultimate_beds_broken_bedwars"`
				EightTwoUltimateItemsPurchasedBedwars             float64  `json:"eight_two_ultimate_items_purchased_bedwars"`
				EightTwoUltimateVoidFinalKillsBedwars             float64  `json:"eight_two_ultimate_void_final_kills_bedwars"`
				EightTwoUltimateEntityAttackFinalDeathsBedwars    float64  `json:"eight_two_ultimate_entity_attack_final_deaths_bedwars"`
				EightTwoUltimateDeathsBedwars                     float64  `json:"eight_two_ultimate_deaths_bedwars"`
				EightTwoUltimateVoidDeathsBedwars                 float64  `json:"eight_two_ultimate_void_deaths_bedwars"`
				EightTwoUltimateEntityAttackKillsBedwars          float64  `json:"eight_two_ultimate_entity_attack_kills_bedwars"`
				EightTwoUltimateWinsBedwars                       float64  `json:"eight_two_ultimate_wins_bedwars"`
				EightTwoUltimateEntityAttackDeathsBedwars         float64  `json:"eight_two_ultimate_entity_attack_deaths_bedwars"`
				EightTwoUltimateEntityAttackFinalKillsBedwars     float64  `json:"eight_two_ultimate_entity_attack_final_kills_bedwars"`
				EightTwoUltimateEmeraldResourcesCollectedBedwars  float64  `json:"eight_two_ultimate_emerald_resources_collected_bedwars"`
				EightTwoUltimateDiamondResourcesCollectedBedwars  float64  `json:"eight_two_ultimate_diamond_resources_collected_bedwars"`
				FourFourUltimateWinstreak                         float64  `json:"four_four_ultimate_winstreak"`
				FourFourUltimateDeathsBedwars                     float64  `json:"four_four_ultimate_deaths_bedwars"`
				FourFourUltimateVoidKillsBedwars                  float64  `json:"four_four_ultimate_void_kills_bedwars"`
				FourFourUltimateFinalKillsBedwars                 float64  `json:"four_four_ultimate_final_kills_bedwars"`
				FourFourUltimatePermanentItemsPurchasedBedwars    float64  `json:"four_four_ultimate_permanent _items_purchased_bedwars"`
				FourFourUltimateEntityAttackKillsBedwars          float64  `json:"four_four_ultimate_entity_attack_kills_bedwars"`
				FourFourUltimateResourcesCollectedBedwars         float64  `json:"four_four_ultimate_resources_collected_bedwars"`
				FourFourUltimateBedsLostBedwars                   float64  `json:"four_four_ultimate_beds_lost_bedwars"`
				FourFourUltimateWinsBedwars                       float64  `json:"four_four_ultimate_wins_bedwars"`
				FourFourUltimateItemsPurchasedBedwars             float64  `json:"four_four_ultimate_items_purchased_bedwars"`
				FourFourUltimateDiamondResourcesCollectedBedwars  float64  `json:"four_four_ultimate_diamond_resources_collected_bedwars"`
				FourFourUltimateIronResourcesCollectedBedwars     float64  `json:"four_four_ultimate_iron_resources_collected_bedwars"`
				FourFourUltimateGoldResourcesCollectedBedwars     float64  `json:"four_four_ultimate_gold_resources_collected_bedwars"`
				FourFourUltimateBedsBrokenBedwars                 float64  `json:"four_four_ultimate_beds_broken_bedwars"`
				FourFourUltimateGamesPlayedBedwars                float64  `json:"four_four_ultimate_games_played_bedwars"`
				FourFourUltimateKillsBedwars                      float64  `json:"four_four_ultimate_kills_bedwars"`
				FourFourUltimateEntityAttackFinalDeathsBedwars    float64  `json:"four_four_ultimate_entity_attack_final_deaths_bedwars"`
				FourFourUltimateVoidDeathsBedwars                 float64  `json:"four_four_ultimate_void_deaths_bedwars"`
				FourFourUltimateVoidFinalKillsBedwars             float64  `json:"four_four_ultimate_void_final_kills_bedwars"`
				FourFourUltimateFinalDeathsBedwars                float64  `json:"four_four_ultimate_final_deaths_bedwars"`
				FourFourUltimateEntityAttackFinalKillsBedwars     float64  `json:"four_four_ultimate_entity_attack_final_kills_bedwars"`
				FourFourUltimateEntityAttackDeathsBedwars         float64  `json:"four_four_ultimate_entity_attack_deaths_bedwars"`
				FourFourUltimateLossesBedwars                     float64  `json:"four_four_ultimate_losses_bedwars"`
				FourFourUltimateVoidFinalDeathsBedwars            float64  `json:"four_four_ultimate_void_final_deaths_bedwars"`
				FourFourUltimateEntityExplosionFinalDeathsBedwars float64  `json:"four_four_ultimate_entity_explosion_final_deaths_bedwars"`
				UnderstandsStreaks                                bool     `json:"understands_streaks"`
				UnderstandsResourceBank                           bool     `json:"understands_resource_bank"`
				CastleWinstreak                                   float64  `json:"castle_winstreak"`
				CastleDeathsBedwars                               float64  `json:"castle_deaths_bedwars"`
				CastleWinsBedwars                                 float64  `json:"castle_wins_bedwars"`
				CastleEntityAttackFinalKillsBedwars               float64  `json:"castle_entity_attack_final_kills_bedwars"`
				CastleEntityExplosionFinalKillsBedwars            float64  `json:"castle_entity_explosion_final_kills_bedwars"`
				CastleItemsPurchasedBedwars                       float64  `json:"castle_items_purchased_bedwars"`
				CastleDiamondResourcesCollectedBedwars            float64  `json:"castle_diamond_resources_collected_bedwars"`
				CastleGoldResourcesCollectedBedwars               float64  `json:"castle_gold_resources_collected_bedwars"`
				CastleResourcesCollectedBedwars                   float64  `json:"castle_resources_collected_bedwars"`
				CastleEntityAttackDeathsBedwars                   float64  `json:"castle_entity_attack_deaths_bedwars"`
				CastleEmeraldResourcesCollectedBedwars            float64  `json:"castle_emerald_resources_collected_bedwars"`
				CastleFinalKillsBedwars                           float64  `json:"castle_final_kills_bedwars"`
				CastleVoidDeathsBedwars                           float64  `json:"castle_void_deaths_bedwars"`
				CastleIronResourcesCollectedBedwars               float64  `json:"castle_iron_resources_collected_bedwars"`
				CastleVoidKillsBedwars                            float64  `json:"castle_void_kills_bedwars"`
				CastleKillsBedwars                                float64  `json:"castle_kills_bedwars"`
				CastleBedsLostBedwars                             float64  `json:"castle_beds_lost_bedwars"`
				CastleEntityExplosionDeathsBedwars                float64  `json:"castle_entity_explosion_deaths_bedwars"`
				CastleGamesPlayedBedwars                          float64  `json:"castle_games_played_bedwars"`
				CastleEntityAttackKillsBedwars                    float64  `json:"castle_entity_attack_kills_bedwars"`
				CastleVoidFinalKillsBedwars                       float64  `json:"castle_void_final_kills_bedwars"`
				CastleFinalDeathsBedwars                          float64  `json:"castle_final_deaths_bedwars"`
				CastleVoidFinalDeathsBedwars                      float64  `json:"castle_void_final_deaths_bedwars"`
				CastleFallKillsBedwars                            float64  `json:"castle_fall_kills_bedwars"`
				CastleEntityExplosionKillsBedwars                 float64  `json:"castle_entity_explosion_kills_bedwars"`
				CastleFallDeathsBedwars                           float64  `json:"castle_fall_deaths_bedwars"`
				CastleProjectileKillsBedwars                      float64  `json:"castle_projectile_kills_bedwars"`
				CastleLossesBedwars                               float64  `json:"castle_losses_bedwars"`
				CastleFallFinalKillsBedwars                       float64  `json:"castle_fall_final_kills_bedwars"`
				CastleBedsBrokenBedwars                           float64  `json:"castle_beds_broken_bedwars"`
				CastleFallFinalDeathsBedwars                      float64  `json:"castle_fall_final_deaths_bedwars"`
				FourFourWinstreak                                 float64  `json:"four_four_winstreak"`
				EightTwoWinstreak                                 float64  `json:"eight_two_winstreak"`
				SelectedUltimate                                  string   `json:"selected_ultimate"`
				EightTwoUltimatePermanentItemsPurchasedBedwars    float64  `json:"eight_two_ultimate_permanent _items_purchased_bedwars"`
				CastleEntityAttackFinalDeathsBedwars              float64  `json:"castle_entity_attack_final_deaths_bedwars"`
				FreeEventKeyBedwarsHalloweenBoxes2018             bool     `json:"free_event_key_bedwars_halloween_boxes_2018"`
				FreeEventKeyBedwarsChristmasBoxes2018             bool     `json:"free_event_key_bedwars_christmas_boxes_2018"`
				LastHytaleAd                                      float64  `json:"lastHytaleAd"`
				EightTwoRushWinstreak                             float64  `json:"eight_two_rush_winstreak"`
				EightTwoRushEntityExplosionKillsBedwars           float64  `json:"eight_two_rush_entity_explosion_kills_bedwars"`
				EightTwoRushItemsPurchasedBedwars                 float64  `json:"eight_two_rush_items_purchased_bedwars"`
				EightTwoRushBedsBrokenBedwars                     float64  `json:"eight_two_rush_beds_broken_bedwars"`
				EightTwoRushResourcesCollectedBedwars             float64  `json:"eight_two_rush_resources_collected_bedwars"`
				EightTwoRushWinsBedwars                           float64  `json:"eight_two_rush_wins_bedwars"`
				EightTwoRushEntityAttackFinalKillsBedwars         float64  `json:"eight_two_rush_entity_attack_final_kills_bedwars"`
				EightTwoRushEmeraldResourcesCollectedBedwars      float64  `json:"eight_two_rush_emerald_resources_collected_bedwars"`
				EightTwoRushKillsBedwars                          float64  `json:"eight_two_rush_kills_bedwars"`
				EightTwoRushDeathsBedwars                         float64  `json:"eight_two_rush_deaths_bedwars"`
				EightTwoRushVoidKillsBedwars                      float64  `json:"eight_two_rush_void_kills_bedwars"`
				EightTwoRushVoidDeathsBedwars                     float64  `json:"eight_two_rush_void_deaths_bedwars"`
				EightTwoRushFinalKillsBedwars                     float64  `json:"eight_two_rush_final_kills_bedwars"`
				EightTwoRushGoldResourcesCollectedBedwars         float64  `json:"eight_two_rush_gold_resources_collected_bedwars"`
				EightTwoRushGamesPlayedBedwars                    float64  `json:"eight_two_rush_games_played_bedwars"`
				EightTwoRushEntityAttackKillsBedwars              float64  `json:"eight_two_rush_entity_attack_kills_bedwars"`
				EightTwoRushIronResourcesCollectedBedwars         float64  `json:"eight_two_rush_iron_resources_collected_bedwars"`
				EightTwoRushLossesBedwars                         float64  `json:"eight_two_rush_losses_bedwars"`
				EightTwoRushVoidFinalDeathsBedwars                float64  `json:"eight_two_rush_void_final_deaths_bedwars"`
				EightTwoRushDiamondResourcesCollectedBedwars      float64  `json:"eight_two_rush_diamond_resources_collected_bedwars"`
				EightTwoRushVoidFinalKillsBedwars                 float64  `json:"eight_two_rush_void_final_kills_bedwars"`
				EightTwoRushBedsLostBedwars                       float64  `json:"eight_two_rush_beds_lost_bedwars"`
				EightTwoRushPermanentItemsPurchasedBedwars        float64  `json:"eight_two_rush_permanent _items_purchased_bedwars"`
				EightTwoRushFinalDeathsBedwars                    float64  `json:"eight_two_rush_final_deaths_bedwars"`
				EightTwoRushEntityAttackDeathsBedwars             float64  `json:"eight_two_rush_entity_attack_deaths_bedwars"`
				Privategames                                      struct {
					OneHitOneKill          bool   `json:"one_hit_one_kill"`
					HealthBuff             string `json:"health_buff"`
					Speed                  string `json:"speed"`
					RespawnTime            string `json:"respawn_time"`
					DisableBlockProtection bool   `json:"disable_block_protection"`
					MaxTeamUpgrades        bool   `json:"max_team_upgrades"`
					BedInstabreak          bool   `json:"bed_instabreak"`
					EventTime              string `json:"event_time"`
					LowGravity             bool   `json:"low_gravity"`
				} `json:"privategames"`
				FourThreeWinstreak                             float64 `json:"four_three_winstreak"`
				EightTwoLuckyBedsBrokenBedwars                 float64 `json:"eight_two_lucky_beds_broken_bedwars"`
				EightTwoLuckyPermanentItemsPurchasedBedwars    float64 `json:"eight_two_lucky_permanent _items_purchased_bedwars"`
				EightTwoLuckyIronResourcesCollectedBedwars     float64 `json:"eight_two_lucky_iron_resources_collected_bedwars"`
				EightTwoLuckyWinsBedwars                       float64 `json:"eight_two_lucky_wins_bedwars"`
				EightTwoLuckyEntityAttackKillsBedwars          float64 `json:"eight_two_lucky_entity_attack_kills_bedwars"`
				EightTwoLuckyKillsBedwars                      float64 `json:"eight_two_lucky_kills_bedwars"`
				EightTwoLuckyGoldResourcesCollectedBedwars     float64 `json:"eight_two_lucky_gold_resources_collected_bedwars"`
				EightTwoLuckyResourcesCollectedBedwars         float64 `json:"eight_two_lucky_resources_collected_bedwars"`
				EightTwoLuckyFinalKillsBedwars                 float64 `json:"eight_two_lucky_final_kills_bedwars"`
				EightTwoLuckyEmeraldResourcesCollectedBedwars  float64 `json:"eight_two_lucky_emerald_resources_collected_bedwars"`
				EightTwoLuckyItemsPurchasedBedwars             float64 `json:"eight_two_lucky_items_purchased_bedwars"`
				EightTwoLuckyVoidFinalKillsBedwars             float64 `json:"eight_two_lucky_void_final_kills_bedwars"`
				EightTwoLuckyBedsLostBedwars                   float64 `json:"eight_two_lucky_beds_lost_bedwars"`
				EightTwoLuckyGamesPlayedBedwars                float64 `json:"eight_two_lucky_games_played_bedwars"`
				EightTwoLuckyWinstreak                         float64 `json:"eight_two_lucky_winstreak"`
				EightTwoLuckyVoidKillsBedwars                  float64 `json:"eight_two_lucky_void_kills_bedwars"`
				EightTwoLuckyEntityAttackFinalDeathsBedwars    float64 `json:"eight_two_lucky_entity_attack_final_deaths_bedwars"`
				EightTwoLuckyLossesBedwars                     float64 `json:"eight_two_lucky_losses_bedwars"`
				EightTwoLuckyFireTickKillsBedwars              float64 `json:"eight_two_lucky_fire_tick_kills_bedwars"`
				EightTwoLuckyFinalDeathsBedwars                float64 `json:"eight_two_lucky_final_deaths_bedwars"`
				EightTwoLuckyDeathsBedwars                     float64 `json:"eight_two_lucky_deaths_bedwars"`
				EightTwoLuckyDiamondResourcesCollectedBedwars  float64 `json:"eight_two_lucky_diamond_resources_collected_bedwars"`
				EightTwoLuckyVoidDeathsBedwars                 float64 `json:"eight_two_lucky_void_deaths_bedwars"`
				FourFourLuckyWinstreak                         float64 `json:"four_four_lucky_winstreak"`
				FourFourLuckyDiamondResourcesCollectedBedwars  float64 `json:"four_four_lucky_diamond_resources_collected_bedwars"`
				FourFourLuckyEntityAttackFinalKillsBedwars     float64 `json:"four_four_lucky_entity_attack_final_kills_bedwars"`
				FourFourLuckyVoidDeathsBedwars                 float64 `json:"four_four_lucky_void_deaths_bedwars"`
				FourFourLuckyPermanentItemsPurchasedBedwars    float64 `json:"four_four_lucky_permanent _items_purchased_bedwars"`
				FourFourLuckyEmeraldResourcesCollectedBedwars  float64 `json:"four_four_lucky_emerald_resources_collected_bedwars"`
				FourFourLuckyItemsPurchasedBedwars             float64 `json:"four_four_lucky_items_purchased_bedwars"`
				FourFourLuckyGamesPlayedBedwars                float64 `json:"four_four_lucky_games_played_bedwars"`
				FourFourLuckyEntityAttackKillsBedwars          float64 `json:"four_four_lucky_entity_attack_kills_bedwars"`
				FourFourLuckyKillsBedwars                      float64 `json:"four_four_lucky_kills_bedwars"`
				FourFourLuckyBedsBrokenBedwars                 float64 `json:"four_four_lucky_beds_broken_bedwars"`
				FourFourLuckyResourcesCollectedBedwars         float64 `json:"four_four_lucky_resources_collected_bedwars"`
				FourFourLuckyFinalKillsBedwars                 float64 `json:"four_four_lucky_final_kills_bedwars"`
				FourFourLuckyDeathsBedwars                     float64 `json:"four_four_lucky_deaths_bedwars"`
				FourFourLuckyIronResourcesCollectedBedwars     float64 `json:"four_four_lucky_iron_resources_collected_bedwars"`
				FourFourLuckyVoidKillsBedwars                  float64 `json:"four_four_lucky_void_kills_bedwars"`
				FourFourLuckyGoldResourcesCollectedBedwars     float64 `json:"four_four_lucky_gold_resources_collected_bedwars"`
				FourFourLuckyWinsBedwars                       float64 `json:"four_four_lucky_wins_bedwars"`
				EightTwoLuckyEntityAttackFinalKillsBedwars     float64 `json:"eight_two_lucky_entity_attack_final_kills_bedwars"`
				EightTwoLuckyProjectileKillsBedwars            float64 `json:"eight_two_lucky_projectile_kills_bedwars"`
				FourFourLuckyEntityAttackDeathsBedwars         float64 `json:"four_four_lucky_entity_attack_deaths_bedwars"`
				FourFourLuckyVoidFinalKillsBedwars             float64 `json:"four_four_lucky_void_final_kills_bedwars"`
				FourFourLuckyFireTickFinalKillsBedwars         float64 `json:"four_four_lucky_fire_tick_final_kills_bedwars"`
				FourFourLuckyProjectileFinalKillsBedwars       float64 `json:"four_four_lucky_projectile_final_kills_bedwars"`
				FourFourLuckyFinalDeathsBedwars                float64 `json:"four_four_lucky_final_deaths_bedwars"`
				FourFourLuckyBedsLostBedwars                   float64 `json:"four_four_lucky_beds_lost_bedwars"`
				FourFourLuckyVoidFinalDeathsBedwars            float64 `json:"four_four_lucky_void_final_deaths_bedwars"`
				FourFourLuckyLossesBedwars                     float64 `json:"four_four_lucky_losses_bedwars"`
				FourFourLuckyFireTickDeathsBedwars             float64 `json:"four_four_lucky_fire_tick_deaths_bedwars"`
				EightTwoLuckyVoidFinalDeathsBedwars            float64 `json:"eight_two_lucky_void_final_deaths_bedwars"`
				EightOneWinstreak                              float64 `json:"eight_one_winstreak"`
				FourFourLuckyEntityAttackFinalDeathsBedwars    float64 `json:"four_four_lucky_entity_attack_final_deaths_bedwars"`
				FourFourLuckyFallDeathsBedwars                 float64 `json:"four_four_lucky_fall_deaths_bedwars"`
				FourFourLuckyEntityExplosionFinalDeathsBedwars float64 `json:"four_four_lucky_entity_explosion_final_deaths_bedwars"`
				FourFourLuckyFallFinalKillsBedwars             float64 `json:"four_four_lucky_fall_final_kills_bedwars"`
				FourFourLuckyFireTickKillsBedwars              float64 `json:"four_four_lucky_fire_tick_kills_bedwars"`
				FourFourLuckyEntityExplosionDeathsBedwars      float64 `json:"four_four_lucky_entity_explosion_deaths_bedwars"`
				FourFourLuckyProjectileDeathsBedwars           float64 `json:"four_four_lucky_projectile_deaths_bedwars"`
				FourFourLuckyProjectileKillsBedwars            float64 `json:"four_four_lucky_projectile_kills_bedwars"`
				FourFourLuckyFireFinalKillsBedwars             float64 `json:"four_four_lucky_fire_final_kills_bedwars"`
				FourFourLuckyFireTickFinalDeathsBedwars        float64 `json:"four_four_lucky_fire_tick_final_deaths_bedwars"`
				FourFourLuckyEntityExplosionFinalKillsBedwars  float64 `json:"four_four_lucky_entity_explosion_final_kills_bedwars"`
				FourFourLuckyFallKillsBedwars                  float64 `json:"four_four_lucky_fall_kills_bedwars"`
				EightTwoLuckyFallDeathsBedwars                 float64 `json:"eight_two_lucky_fall_deaths_bedwars"`
				EightTwoLuckyEntityExplosionKillsBedwars       float64 `json:"eight_two_lucky_entity_explosion_kills_bedwars"`
				ActiveIslandTopper                             string  `json:"activeIslandTopper"`
				FourFourArmedWinstreak                         float64 `json:"four_four_armed_winstreak"`
				FourFourArmedBedsBrokenBedwars                 float64 `json:"four_four_armed_beds_broken_bedwars"`
				FourFourArmedBedsLostBedwars                   float64 `json:"four_four_armed_beds_lost_bedwars"`
				FourFourArmedDeathsBedwars                     float64 `json:"four_four_armed_deaths_bedwars"`
				FourFourArmedDiamondResourcesCollectedBedwars  float64 `json:"four_four_armed_diamond_resources_collected_bedwars"`
				FourFourArmedEntityAttackFinalDeathsBedwars    float64 `json:"four_four_armed_entity_attack_final_deaths_bedwars"`
				FourFourArmedEntityAttackKillsBedwars          float64 `json:"four_four_armed_entity_attack_kills_bedwars"`
				FourFourArmedFinalDeathsBedwars                float64 `json:"four_four_armed_final_deaths_bedwars"`
				FourFourArmedFinalKillsBedwars                 float64 `json:"four_four_armed_final_kills_bedwars"`
				FourFourArmedGamesPlayedBedwars                float64 `json:"four_four_armed_games_played_bedwars"`
				FourFourArmedGoldResourcesCollectedBedwars     float64 `json:"four_four_armed_gold_resources_collected_bedwars"`
				FourFourArmedIronResourcesCollectedBedwars     float64 `json:"four_four_armed_iron_resources_collected_bedwars"`
				FourFourArmedItemsPurchasedBedwars             float64 `json:"four_four_armed_items_purchased_bedwars"`
				FourFourArmedKillsBedwars                      float64 `json:"four_four_armed_kills_bedwars"`
				FourFourArmedLossesBedwars                     float64 `json:"four_four_armed_losses_bedwars"`
				FourFourArmedProjectileDeathsBedwars           float64 `json:"four_four_armed_projectile_deaths_bedwars"`
				FourFourArmedProjectileKillsBedwars            float64 `json:"four_four_armed_projectile_kills_bedwars"`
				FourFourArmedResourcesCollectedBedwars         float64 `json:"four_four_armed_resources_collected_bedwars"`
				FourFourArmedVoidDeathsBedwars                 float64 `json:"four_four_armed_void_deaths_bedwars"`
				FourFourArmedVoidFinalKillsBedwars             float64 `json:"four_four_armed_void_final_kills_bedwars"`
				FourFourArmedVoidKillsBedwars                  float64 `json:"four_four_armed_void_kills_bedwars"`
				FourFourArmedEmeraldResourcesCollectedBedwars  float64 `json:"four_four_armed_emerald_resources_collected_bedwars"`
				FourFourArmedEntityExplosionDeathsBedwars      float64 `json:"four_four_armed_entity_explosion_deaths_bedwars"`
				FourFourArmedPermanentItemsPurchasedBedwars    float64 `json:"four_four_armed_permanent _items_purchased_bedwars"`
				FourFourArmedEntityAttackDeathsBedwars         float64 `json:"four_four_armed_entity_attack_deaths_bedwars"`
				FourFourArmedEntityExplosionKillsBedwars       float64 `json:"four_four_armed_entity_explosion_kills_bedwars"`
				FourFourArmedFallKillsBedwars                  float64 `json:"four_four_armed_fall_kills_bedwars"`
				FourFourArmedProjectileFinalKillsBedwars       float64 `json:"four_four_armed_projectile_final_kills_bedwars"`
				FourFourArmedWinsBedwars                       float64 `json:"four_four_armed_wins_bedwars"`
				FourFourArmedProjectileFinalDeathsBedwars      float64 `json:"four_four_armed_projectile_final_deaths_bedwars"`
				FourFourArmedFallFinalDeathsBedwars            float64 `json:"four_four_armed_fall_final_deaths_bedwars"`
				FourFourArmedEntityExplosionFinalKillsBedwars  float64 `json:"four_four_armed_entity_explosion_final_kills_bedwars"`
				TwoFourWinstreak                               float64 `json:"two_four_winstreak"`
				TwoFourEntityAttackFinalKillsBedwars           float64 `json:"two_four_entity_attack_final_kills_bedwars"`
				TwoFourEntityAttackKillsBedwars                float64 `json:"two_four_entity_attack_kills_bedwars"`
				TwoFourFinalKillsBedwars                       float64 `json:"two_four_final_kills_bedwars"`
				TwoFourGamesPlayedBedwars                      float64 `json:"two_four_games_played_bedwars"`
				TwoFourGoldResourcesCollectedBedwars           float64 `json:"two_four_gold_resources_collected_bedwars"`
				TwoFourIronResourcesCollectedBedwars           float64 `json:"two_four_iron_resources_collected_bedwars"`
				TwoFourItemsPurchasedBedwars                   float64 `json:"two_four_items_purchased_bedwars"`
				TwoFourKillsBedwars                            float64 `json:"two_four_kills_bedwars"`
				TwoFourResourcesCollectedBedwars               float64 `json:"two_four_resources_collected_bedwars"`
				TwoFourVoidKillsBedwars                        float64 `json:"two_four_void_kills_bedwars"`
				TwoFourWinsBedwars                             float64 `json:"two_four_wins_bedwars"`
				TwoFourBedsLostBedwars                         float64 `json:"two_four_beds_lost_bedwars"`
				TwoFourDeathsBedwars                           float64 `json:"two_four_deaths_bedwars"`
				TwoFourEntityAttackDeathsBedwars               float64 `json:"two_four_entity_attack_deaths_bedwars"`
				TwoFourEntityAttackFinalDeathsBedwars          float64 `json:"two_four_entity_attack_final_deaths_bedwars"`
				TwoFourFinalDeathsBedwars                      float64 `json:"two_four_final_deaths_bedwars"`
				TwoFourLossesBedwars                           float64 `json:"two_four_losses_bedwars"`
				TwoFourVoidDeathsBedwars                       float64 `json:"two_four_void_deaths_bedwars"`
				TwoFourVoidFinalKillsBedwars                   float64 `json:"two_four_void_final_kills_bedwars"`
				TwoFourBedsBrokenBedwars                       float64 `json:"two_four_beds_broken_bedwars"`
				CastleMagicKillsBedwars                        float64 `json:"castle_magic_kills_bedwars"`
				EightTwoPermanentItemsPurchasedBedwars         float64 `json:"eight_two_permanent_items_purchased_bedwars"`
				PermanentItemsPurchasedBedwars                 float64 `json:"permanent_items_purchased_bedwars"`
				CastlePermanentItemsPurchasedBedwars           float64 `json:"castle_permanent_items_purchased_bedwars"`
				Practice                                       struct {
					Selected string `json:"selected"`
					Mlg      struct {
						SuccessfulAttempts float64 `json:"successful_attempts"`
					} `json:"mlg"`
					FireballJumping struct {
						BlocksPlaced       float64 `json:"blocks_placed"`
						SuccessfulAttempts float64 `json:"successful_attempts"`
						FailedAttempts     float64 `json:"failed_attempts"`
					} `json:"fireball_jumping"`
					Records struct {
						BridgingDistance30ElevationNONEAngleSTRAIGHT float64 `json:"bridging_distance_30:elevation_NONE:angle_STRAIGHT:"`
					} `json:"records"`
					Bridging struct {
						FailedAttempts     float64 `json:"failed_attempts"`
						SuccessfulAttempts float64 `json:"successful_attempts"`
						BlocksPlaced       float64 `json:"blocks_placed"`
					} `json:"bridging"`
				} `json:"practice"`
			} `json:"Bedwars"`
			MurderMystery struct {
				Coins                                                    float64  `json:"coins"`
				GamesMURDERASSASSINS                                     float64  `json:"games_MURDER_ASSASSINS"`
				GamesCruiseShipMURDERASSASSINS                           float64  `json:"games_cruise_ship_MURDER_ASSASSINS"`
				KillsCruiseShipMURDERASSASSINS                           float64  `json:"kills_cruise_ship_MURDER_ASSASSINS"`
				CoinsPickedupCruiseShip                                  float64  `json:"coins_pickedup_cruise_ship"`
				KnifeKills                                               float64  `json:"knife_kills"`
				GamesCruiseShip                                          float64  `json:"games_cruise_ship"`
				KillsCruiseShip                                          float64  `json:"kills_cruise_ship"`
				DeathsCruiseShipMURDERASSASSINS                          float64  `json:"deaths_cruise_ship_MURDER_ASSASSINS"`
				DeathsCruiseShip                                         float64  `json:"deaths_cruise_ship"`
				KnifeKillsCruiseShip                                     float64  `json:"knife_kills_cruise_ship"`
				Deaths                                                   float64  `json:"deaths"`
				DeathsMURDERASSASSINS                                    float64  `json:"deaths_MURDER_ASSASSINS"`
				CoinsPickedupCruiseShipMURDERASSASSINS                   float64  `json:"coins_pickedup_cruise_ship_MURDER_ASSASSINS"`
				KnifeKillsCruiseShipMURDERASSASSINS                      float64  `json:"knife_kills_cruise_ship_MURDER_ASSASSINS"`
				Games                                                    float64  `json:"games"`
				Kills                                                    float64  `json:"kills"`
				KnifeKillsMURDERASSASSINS                                float64  `json:"knife_kills_MURDER_ASSASSINS"`
				CoinsPickedupMURDERASSASSINS                             float64  `json:"coins_pickedup_MURDER_ASSASSINS"`
				CoinsPickedup                                            float64  `json:"coins_pickedup"`
				KillsMURDERASSASSINS                                     float64  `json:"kills_MURDER_ASSASSINS"`
				MurdermysteryBooks                                       []string `json:"murdermystery_books"`
				Wins                                                     float64  `json:"wins"`
				WinsMURDERHARDCORE                                       float64  `json:"wins_MURDER_HARDCORE"`
				GamesHeadquartersMURDERHARDCORE                          float64  `json:"games_headquarters_MURDER_HARDCORE"`
				GamesMURDERHARDCORE                                      float64  `json:"games_MURDER_HARDCORE"`
				CoinsPickedupHeadquarters                                float64  `json:"coins_pickedup_headquarters"`
				WinsHeadquarters                                         float64  `json:"wins_headquarters"`
				CoinsPickedupHeadquartersMURDERHARDCORE                  float64  `json:"coins_pickedup_headquarters_MURDER_HARDCORE"`
				GamesHeadquarters                                        float64  `json:"games_headquarters"`
				CoinsPickedupMURDERHARDCORE                              float64  `json:"coins_pickedup_MURDER_HARDCORE"`
				WinsHeadquartersMURDERHARDCORE                           float64  `json:"wins_headquarters_MURDER_HARDCORE"`
				GrantedChests                                            float64  `json:"granted_chests"`
				MmChests                                                 float64  `json:"mm_chests"`
				DetectiveChance                                          float64  `json:"detective_chance"`
				WinsGoldRushMURDERCLASSIC                                float64  `json:"wins_gold_rush_MURDER_CLASSIC"`
				GamesGoldRushMURDERCLASSIC                               float64  `json:"games_gold_rush_MURDER_CLASSIC"`
				WasHero                                                  float64  `json:"was_hero"`
				KillsGoldRushMURDERCLASSIC                               float64  `json:"kills_gold_rush_MURDER_CLASSIC"`
				BowKillsMURDERCLASSIC                                    float64  `json:"bow_kills_MURDER_CLASSIC"`
				WasHeroMURDERCLASSIC                                     float64  `json:"was_hero_MURDER_CLASSIC"`
				BowKillsGoldRushMURDERCLASSIC                            float64  `json:"bow_kills_gold_rush_MURDER_CLASSIC"`
				WasHeroGoldRushMURDERCLASSIC                             float64  `json:"was_hero_gold_rush_MURDER_CLASSIC"`
				GamesGoldRush                                            float64  `json:"games_gold_rush"`
				KillsMURDERCLASSIC                                       float64  `json:"kills_MURDER_CLASSIC"`
				CoinsPickedupGoldRushMURDERCLASSIC                       float64  `json:"coins_pickedup_gold_rush_MURDER_CLASSIC"`
				WinsMURDERCLASSIC                                        float64  `json:"wins_MURDER_CLASSIC"`
				CoinsPickedupMURDERCLASSIC                               float64  `json:"coins_pickedup_MURDER_CLASSIC"`
				DetectiveWinsGoldRushMURDERCLASSIC                       float64  `json:"detective_wins_gold_rush_MURDER_CLASSIC"`
				BowKills                                                 float64  `json:"bow_kills"`
				KillsGoldRush                                            float64  `json:"kills_gold_rush"`
				DetectiveWinsGoldRush                                    float64  `json:"detective_wins_gold_rush"`
				DetectiveWins                                            float64  `json:"detective_wins"`
				WasHeroGoldRush                                          float64  `json:"was_hero_gold_rush"`
				WinsGoldRush                                             float64  `json:"wins_gold_rush"`
				DetectiveWinsMURDERCLASSIC                               float64  `json:"detective_wins_MURDER_CLASSIC"`
				GamesMURDERCLASSIC                                       float64  `json:"games_MURDER_CLASSIC"`
				CoinsPickedupGoldRush                                    float64  `json:"coins_pickedup_gold_rush"`
				BowKillsGoldRush                                         float64  `json:"bow_kills_gold_rush"`
				GamesAncientTomb                                         float64  `json:"games_ancient_tomb"`
				DetectiveWinsAncientTombMURDERCLASSIC                    float64  `json:"detective_wins_ancient_tomb_MURDER_CLASSIC"`
				GamesAncientTombMURDERCLASSIC                            float64  `json:"games_ancient_tomb_MURDER_CLASSIC"`
				WinsAncientTombMURDERCLASSIC                             float64  `json:"wins_ancient_tomb_MURDER_CLASSIC"`
				CoinsPickedupAncientTombMURDERCLASSIC                    float64  `json:"coins_pickedup_ancient_tomb_MURDER_CLASSIC"`
				DetectiveWinsAncientTomb                                 float64  `json:"detective_wins_ancient_tomb"`
				CoinsPickedupAncientTomb                                 float64  `json:"coins_pickedup_ancient_tomb"`
				WinsAncientTomb                                          float64  `json:"wins_ancient_tomb"`
				WinsHollywood                                            float64  `json:"wins_hollywood"`
				WinsHollywoodMURDERCLASSIC                               float64  `json:"wins_hollywood_MURDER_CLASSIC"`
				DetectiveWinsHollywood                                   float64  `json:"detective_wins_hollywood"`
				CoinsPickedupHollywoodMURDERCLASSIC                      float64  `json:"coins_pickedup_hollywood_MURDER_CLASSIC"`
				CoinsPickedupHollywood                                   float64  `json:"coins_pickedup_hollywood"`
				GamesHollywood                                           float64  `json:"games_hollywood"`
				GamesHollywoodMURDERCLASSIC                              float64  `json:"games_hollywood_MURDER_CLASSIC"`
				DetectiveWinsHollywoodMURDERCLASSIC                      float64  `json:"detective_wins_hollywood_MURDER_CLASSIC"`
				CoinsPickedupHollywoodMURDERHARDCORE                     float64  `json:"coins_pickedup_hollywood_MURDER_HARDCORE"`
				WinsHollywoodMURDERHARDCORE                              float64  `json:"wins_hollywood_MURDER_HARDCORE"`
				DetectiveWinsMURDERHARDCORE                              float64  `json:"detective_wins_MURDER_HARDCORE"`
				GamesHollywoodMURDERHARDCORE                             float64  `json:"games_hollywood_MURDER_HARDCORE"`
				DetectiveWinsHollywoodMURDERHARDCORE                     float64  `json:"detective_wins_hollywood_MURDER_HARDCORE"`
				WinsLibraryMURDERCLASSIC                                 float64  `json:"wins_library_MURDER_CLASSIC"`
				WinsLibrary                                              float64  `json:"wins_library"`
				GamesLibraryMURDERCLASSIC                                float64  `json:"games_library_MURDER_CLASSIC"`
				CoinsPickedupLibrary                                     float64  `json:"coins_pickedup_library"`
				CoinsPickedupLibraryMURDERCLASSIC                        float64  `json:"coins_pickedup_library_MURDER_CLASSIC"`
				GamesLibrary                                             float64  `json:"games_library"`
				WasHeroCruiseShipMURDERCLASSIC                           float64  `json:"was_hero_cruise_ship_MURDER_CLASSIC"`
				BowKillsCruiseShip                                       float64  `json:"bow_kills_cruise_ship"`
				DetectiveWinsCruiseShip                                  float64  `json:"detective_wins_cruise_ship"`
				WinsCruiseShipMURDERCLASSIC                              float64  `json:"wins_cruise_ship_MURDER_CLASSIC"`
				CoinsPickedupCruiseShipMURDERCLASSIC                     float64  `json:"coins_pickedup_cruise_ship_MURDER_CLASSIC"`
				GamesCruiseShipMURDERCLASSIC                             float64  `json:"games_cruise_ship_MURDER_CLASSIC"`
				WasHeroCruiseShip                                        float64  `json:"was_hero_cruise_ship"`
				BowKillsCruiseShipMURDERCLASSIC                          float64  `json:"bow_kills_cruise_ship_MURDER_CLASSIC"`
				DetectiveWinsCruiseShipMURDERCLASSIC                     float64  `json:"detective_wins_cruise_ship_MURDER_CLASSIC"`
				WinsCruiseShip                                           float64  `json:"wins_cruise_ship"`
				KillsCruiseShipMURDERCLASSIC                             float64  `json:"kills_cruise_ship_MURDER_CLASSIC"`
				MurdererChance                                           float64  `json:"murderer_chance"`
				BowKillsLibraryMURDERCLASSIC                             float64  `json:"bow_kills_library_MURDER_CLASSIC"`
				KillsLibraryMURDERCLASSIC                                float64  `json:"kills_library_MURDER_CLASSIC"`
				KnifeKillsMURDERCLASSIC                                  float64  `json:"knife_kills_MURDER_CLASSIC"`
				KnifeKillsLibraryMURDERCLASSIC                           float64  `json:"knife_kills_library_MURDER_CLASSIC"`
				BowKillsLibrary                                          float64  `json:"bow_kills_library"`
				MurdererWinsLibraryMURDERCLASSIC                         float64  `json:"murderer_wins_library_MURDER_CLASSIC"`
				KillsAsMurdererMURDERCLASSIC                             float64  `json:"kills_as_murderer_MURDER_CLASSIC"`
				KnifeKillsLibrary                                        float64  `json:"knife_kills_library"`
				KillsAsMurdererLibrary                                   float64  `json:"kills_as_murderer_library"`
				KillsAsMurdererLibraryMURDERCLASSIC                      float64  `json:"kills_as_murderer_library_MURDER_CLASSIC"`
				KillsAsMurderer                                          float64  `json:"kills_as_murderer"`
				MurdererWinsMURDERCLASSIC                                float64  `json:"murderer_wins_MURDER_CLASSIC"`
				KillsLibrary                                             float64  `json:"kills_library"`
				MurdererWinsLibrary                                      float64  `json:"murderer_wins_library"`
				MurdererWins                                             float64  `json:"murderer_wins"`
				MurdererWinsCruiseShip                                   float64  `json:"murderer_wins_cruise_ship"`
				MurdererWinsCruiseShipMURDERCLASSIC                      float64  `json:"murderer_wins_cruise_ship_MURDER_CLASSIC"`
				ThrownKnifeKillsCruiseShip                               float64  `json:"thrown_knife_kills_cruise_ship"`
				ThrownKnifeKills                                         float64  `json:"thrown_knife_kills"`
				KillsAsMurdererCruiseShip                                float64  `json:"kills_as_murderer_cruise_ship"`
				ThrownKnifeKillsCruiseShipMURDERCLASSIC                  float64  `json:"thrown_knife_kills_cruise_ship_MURDER_CLASSIC"`
				KillsAsMurdererCruiseShipMURDERCLASSIC                   float64  `json:"kills_as_murderer_cruise_ship_MURDER_CLASSIC"`
				ThrownKnifeKillsMURDERCLASSIC                            float64  `json:"thrown_knife_kills_MURDER_CLASSIC"`
				KnifeKillsCruiseShipMURDERCLASSIC                        float64  `json:"knife_kills_cruise_ship_MURDER_CLASSIC"`
				BowKillsHollywoodMURDERCLASSIC                           float64  `json:"bow_kills_hollywood_MURDER_CLASSIC"`
				KillsHollywood                                           float64  `json:"kills_hollywood"`
				KillsHollywoodMURDERCLASSIC                              float64  `json:"kills_hollywood_MURDER_CLASSIC"`
				WasHeroHollywood                                         float64  `json:"was_hero_hollywood"`
				WasHeroHollywoodMURDERCLASSIC                            float64  `json:"was_hero_hollywood_MURDER_CLASSIC"`
				BowKillsHollywood                                        float64  `json:"bow_kills_hollywood"`
				DeathsGoldRush                                           float64  `json:"deaths_gold_rush"`
				DeathsGoldRushMURDERCLASSIC                              float64  `json:"deaths_gold_rush_MURDER_CLASSIC"`
				DeathsMURDERCLASSIC                                      float64  `json:"deaths_MURDER_CLASSIC"`
				WinsTransport                                            float64  `json:"wins_transport"`
				CoinsPickedupTransport                                   float64  `json:"coins_pickedup_transport"`
				WinsTransportMURDERCLASSIC                               float64  `json:"wins_transport_MURDER_CLASSIC"`
				GamesTransportMURDERCLASSIC                              float64  `json:"games_transport_MURDER_CLASSIC"`
				GamesTransport                                           float64  `json:"games_transport"`
				CoinsPickedupTransportMURDERCLASSIC                      float64  `json:"coins_pickedup_transport_MURDER_CLASSIC"`
				CoinsPickedupHypixelWorld                                float64  `json:"coins_pickedup_hypixel_world"`
				WasHeroHypixelWorldMURDERCLASSIC                         float64  `json:"was_hero_hypixel_world_MURDER_CLASSIC"`
				WinsHypixelWorld                                         float64  `json:"wins_hypixel_world"`
				KillsHypixelWorldMURDERCLASSIC                           float64  `json:"kills_hypixel_world_MURDER_CLASSIC"`
				DetectiveWinsHypixelWorld                                float64  `json:"detective_wins_hypixel_world"`
				KillsHypixelWorld                                        float64  `json:"kills_hypixel_world"`
				GamesHypixelWorldMURDERCLASSIC                           float64  `json:"games_hypixel_world_MURDER_CLASSIC"`
				BowKillsHypixelWorldMURDERCLASSIC                        float64  `json:"bow_kills_hypixel_world_MURDER_CLASSIC"`
				WinsHypixelWorldMURDERCLASSIC                            float64  `json:"wins_hypixel_world_MURDER_CLASSIC"`
				GamesHypixelWorld                                        float64  `json:"games_hypixel_world"`
				BowKillsHypixelWorld                                     float64  `json:"bow_kills_hypixel_world"`
				CoinsPickedupHypixelWorldMURDERCLASSIC                   float64  `json:"coins_pickedup_hypixel_world_MURDER_CLASSIC"`
				WasHeroHypixelWorld                                      float64  `json:"was_hero_hypixel_world"`
				DetectiveWinsHypixelWorldMURDERCLASSIC                   float64  `json:"detective_wins_hypixel_world_MURDER_CLASSIC"`
				DetectiveWinsLibrary                                     float64  `json:"detective_wins_library"`
				WasHeroLibraryMURDERCLASSIC                              float64  `json:"was_hero_library_MURDER_CLASSIC"`
				WasHeroLibrary                                           float64  `json:"was_hero_library"`
				DetectiveWinsLibraryMURDERCLASSIC                        float64  `json:"detective_wins_library_MURDER_CLASSIC"`
				KnifeKillsHollywood                                      float64  `json:"knife_kills_hollywood"`
				MurdererWinsHollywood                                    float64  `json:"murderer_wins_hollywood"`
				ThrownKnifeKillsHollywood                                float64  `json:"thrown_knife_kills_hollywood"`
				KnifeKillsHollywoodMURDERCLASSIC                         float64  `json:"knife_kills_hollywood_MURDER_CLASSIC"`
				MurdererWinsHollywoodMURDERCLASSIC                       float64  `json:"murderer_wins_hollywood_MURDER_CLASSIC"`
				ThrownKnifeKillsHollywoodMURDERCLASSIC                   float64  `json:"thrown_knife_kills_hollywood_MURDER_CLASSIC"`
				KillsAsMurdererHollywoodMURDERCLASSIC                    float64  `json:"kills_as_murderer_hollywood_MURDER_CLASSIC"`
				KillsAsMurdererHollywood                                 float64  `json:"kills_as_murderer_hollywood"`
				Packages                                                 []string `json:"packages"`
				MmChestHistory                                           []string `json:"mm_chest_history"`
				MurderMysteryOpenedCommons                               float64  `json:"MurderMystery_openedCommons"`
				MurderMysteryOpenedChests                                float64  `json:"MurderMystery_openedChests"`
				MurderMysteryOpenedRares                                 float64  `json:"MurderMystery_openedRares"`
				MurderMysteryOpenedEpics                                 float64  `json:"MurderMystery_openedEpics"`
				MurderMysteryOpenedLegendaries                           float64  `json:"MurderMystery_openedLegendaries"`
				ThrownKnifeKillsArchives                                 float64  `json:"thrown_knife_kills_archives"`
				KillsAsMurdererArchives                                  float64  `json:"kills_as_murderer_archives"`
				DeathsArchivesMURDERCLASSIC                              float64  `json:"deaths_archives_MURDER_CLASSIC"`
				CoinsPickedupArchives                                    float64  `json:"coins_pickedup_archives"`
				KnifeKillsArchives                                       float64  `json:"knife_kills_archives"`
				KillsArchives                                            float64  `json:"kills_archives"`
				GamesArchivesMURDERCLASSIC                               float64  `json:"games_archives_MURDER_CLASSIC"`
				KillsArchivesMURDERCLASSIC                               float64  `json:"kills_archives_MURDER_CLASSIC"`
				KnifeKillsArchivesMURDERCLASSIC                          float64  `json:"knife_kills_archives_MURDER_CLASSIC"`
				ThrownKnifeKillsArchivesMURDERCLASSIC                    float64  `json:"thrown_knife_kills_archives_MURDER_CLASSIC"`
				DeathsArchives                                           float64  `json:"deaths_archives"`
				CoinsPickedupArchivesMURDERCLASSIC                       float64  `json:"coins_pickedup_archives_MURDER_CLASSIC"`
				GamesArchives                                            float64  `json:"games_archives"`
				KillsAsMurdererArchivesMURDERCLASSIC                     float64  `json:"kills_as_murderer_archives_MURDER_CLASSIC"`
				KillsAsMurdererAncientTombMURDERCLASSIC                  float64  `json:"kills_as_murderer_ancient_tomb_MURDER_CLASSIC"`
				DeathsAncientTomb                                        float64  `json:"deaths_ancient_tomb"`
				KillsAncientTombMURDERCLASSIC                            float64  `json:"kills_ancient_tomb_MURDER_CLASSIC"`
				KillsAncientTomb                                         float64  `json:"kills_ancient_tomb"`
				KillsAsMurdererAncientTomb                               float64  `json:"kills_as_murderer_ancient_tomb"`
				KnifeKillsAncientTombMURDERCLASSIC                       float64  `json:"knife_kills_ancient_tomb_MURDER_CLASSIC"`
				KnifeKillsAncientTomb                                    float64  `json:"knife_kills_ancient_tomb"`
				DeathsAncientTombMURDERCLASSIC                           float64  `json:"deaths_ancient_tomb_MURDER_CLASSIC"`
				ThrownKnifeKillsAncientTombMURDERCLASSIC                 float64  `json:"thrown_knife_kills_ancient_tomb_MURDER_CLASSIC"`
				ThrownKnifeKillsAncientTomb                              float64  `json:"thrown_knife_kills_ancient_tomb"`
				MurdererWinsGoldRushMURDERCLASSIC                        float64  `json:"murderer_wins_gold_rush_MURDER_CLASSIC"`
				MurdererWinsGoldRush                                     float64  `json:"murderer_wins_gold_rush"`
				KillsAsMurdererGoldRushMURDERCLASSIC                     float64  `json:"kills_as_murderer_gold_rush_MURDER_CLASSIC"`
				ThrownKnifeKillsGoldRush                                 float64  `json:"thrown_knife_kills_gold_rush"`
				KillsAsMurdererGoldRush                                  float64  `json:"kills_as_murderer_gold_rush"`
				ThrownKnifeKillsGoldRushMURDERCLASSIC                    float64  `json:"thrown_knife_kills_gold_rush_MURDER_CLASSIC"`
				KnifeKillsGoldRush                                       float64  `json:"knife_kills_gold_rush"`
				KnifeKillsGoldRushMURDERCLASSIC                          float64  `json:"knife_kills_gold_rush_MURDER_CLASSIC"`
				MurdererWinsArchives                                     float64  `json:"murderer_wins_archives"`
				MurdererWinsArchivesMURDERCLASSIC                        float64  `json:"murderer_wins_archives_MURDER_CLASSIC"`
				WinsArchives                                             float64  `json:"wins_archives"`
				WinsArchivesMURDERCLASSIC                                float64  `json:"wins_archives_MURDER_CLASSIC"`
				DetectiveWinsArchivesMURDERCLASSIC                       float64  `json:"detective_wins_archives_MURDER_CLASSIC"`
				DetectiveWinsArchives                                    float64  `json:"detective_wins_archives"`
				ActiveVictoryDance                                       string   `json:"active_victory_dance"`
				ActiveProjectileTrail                                    string   `json:"active_projectile_trail"`
				ActiveLastWords                                          string   `json:"active_last_words"`
				ActiveKnifeSkin                                          string   `json:"active_knife_skin"`
				BowKillsArchives                                         float64  `json:"bow_kills_archives"`
				WasHeroArchives                                          float64  `json:"was_hero_archives"`
				WasHeroArchivesMURDERCLASSIC                             float64  `json:"was_hero_archives_MURDER_CLASSIC"`
				BowKillsArchivesMURDERCLASSIC                            float64  `json:"bow_kills_archives_MURDER_CLASSIC"`
				GamesTowerfallMURDERCLASSIC                              float64  `json:"games_towerfall_MURDER_CLASSIC"`
				CoinsPickedupTowerfall                                   float64  `json:"coins_pickedup_towerfall"`
				GamesTowerfall                                           float64  `json:"games_towerfall"`
				WinsTowerfallMURDERCLASSIC                               float64  `json:"wins_towerfall_MURDER_CLASSIC"`
				WinsTowerfall                                            float64  `json:"wins_towerfall"`
				CoinsPickedupTowerfallMURDERCLASSIC                      float64  `json:"coins_pickedup_towerfall_MURDER_CLASSIC"`
				MurdererWinsAncientTombMURDERCLASSIC                     float64  `json:"murderer_wins_ancient_tomb_MURDER_CLASSIC"`
				MurdererWinsAncientTomb                                  float64  `json:"murderer_wins_ancient_tomb"`
				WasHeroTransport                                         float64  `json:"was_hero_transport"`
				KillsTransportMURDERCLASSIC                              float64  `json:"kills_transport_MURDER_CLASSIC"`
				KillsTransport                                           float64  `json:"kills_transport"`
				DetectiveWinsTransport                                   float64  `json:"detective_wins_transport"`
				DetectiveWinsTransportMURDERCLASSIC                      float64  `json:"detective_wins_transport_MURDER_CLASSIC"`
				BowKillsTransportMURDERCLASSIC                           float64  `json:"bow_kills_transport_MURDER_CLASSIC"`
				BowKillsTransport                                        float64  `json:"bow_kills_transport"`
				WasHeroTransportMURDERCLASSIC                            float64  `json:"was_hero_transport_MURDER_CLASSIC"`
				GamesAncientTombMURDERASSASSINS                          float64  `json:"games_ancient_tomb_MURDER_ASSASSINS"`
				ThrownKnifeKillsMURDERASSASSINS                          float64  `json:"thrown_knife_kills_MURDER_ASSASSINS"`
				KillsAncientTombMURDERASSASSINS                          float64  `json:"kills_ancient_tomb_MURDER_ASSASSINS"`
				CoinsPickedupAncientTombMURDERASSASSINS                  float64  `json:"coins_pickedup_ancient_tomb_MURDER_ASSASSINS"`
				WinsAncientTombMURDERASSASSINS                           float64  `json:"wins_ancient_tomb_MURDER_ASSASSINS"`
				KnifeKillsAncientTombMURDERASSASSINS                     float64  `json:"knife_kills_ancient_tomb_MURDER_ASSASSINS"`
				WinsMURDERASSASSINS                                      float64  `json:"wins_MURDER_ASSASSINS"`
				ThrownKnifeKillsAncientTombMURDERASSASSINS               float64  `json:"thrown_knife_kills_ancient_tomb_MURDER_ASSASSINS"`
				BowKillsAncientTomb                                      float64  `json:"bow_kills_ancient_tomb"`
				WasHeroAncientTomb                                       float64  `json:"was_hero_ancient_tomb"`
				BowKillsAncientTombMURDERCLASSIC                         float64  `json:"bow_kills_ancient_tomb_MURDER_CLASSIC"`
				WasHeroAncientTombMURDERCLASSIC                          float64  `json:"was_hero_ancient_tomb_MURDER_CLASSIC"`
				Showqueuebook                                            bool     `json:"showqueuebook"`
				DeathsHollywoodMURDERCLASSIC                             float64  `json:"deaths_hollywood_MURDER_CLASSIC"`
				DeathsHollywood                                          float64  `json:"deaths_hollywood"`
				KillsAsMurdererTowerfallMURDERCLASSIC                    float64  `json:"kills_as_murderer_towerfall_MURDER_CLASSIC"`
				KnifeKillsTowerfall                                      float64  `json:"knife_kills_towerfall"`
				KillsTowerfall                                           float64  `json:"kills_towerfall"`
				MurdererWinsTowerfallMURDERCLASSIC                       float64  `json:"murderer_wins_towerfall_MURDER_CLASSIC"`
				KnifeKillsTowerfallMURDERCLASSIC                         float64  `json:"knife_kills_towerfall_MURDER_CLASSIC"`
				ThrownKnifeKillsTowerfall                                float64  `json:"thrown_knife_kills_towerfall"`
				ThrownKnifeKillsTowerfallMURDERCLASSIC                   float64  `json:"thrown_knife_kills_towerfall_MURDER_CLASSIC"`
				KillsTowerfallMURDERCLASSIC                              float64  `json:"kills_towerfall_MURDER_CLASSIC"`
				KillsAsMurdererTowerfall                                 float64  `json:"kills_as_murderer_towerfall"`
				MurdererWinsTowerfall                                    float64  `json:"murderer_wins_towerfall"`
				KnifeKillsHypixelWorld                                   float64  `json:"knife_kills_hypixel_world"`
				ThrownKnifeKillsHypixelWorld                             float64  `json:"thrown_knife_kills_hypixel_world"`
				KillsAsMurdererHypixelWorld                              float64  `json:"kills_as_murderer_hypixel_world"`
				MurdererWinsHypixelWorldMURDERCLASSIC                    float64  `json:"murderer_wins_hypixel_world_MURDER_CLASSIC"`
				ThrownKnifeKillsHypixelWorldMURDERCLASSIC                float64  `json:"thrown_knife_kills_hypixel_world_MURDER_CLASSIC"`
				KillsAsMurdererHypixelWorldMURDERCLASSIC                 float64  `json:"kills_as_murderer_hypixel_world_MURDER_CLASSIC"`
				MurdererWinsHypixelWorld                                 float64  `json:"murderer_wins_hypixel_world"`
				KnifeKillsHypixelWorldMURDERCLASSIC                      float64  `json:"knife_kills_hypixel_world_MURDER_CLASSIC"`
				DeathsHypixelWorldMURDERCLASSIC                          float64  `json:"deaths_hypixel_world_MURDER_CLASSIC"`
				DeathsHypixelWorld                                       float64  `json:"deaths_hypixel_world"`
				GamesHeadquartersMURDERCLASSIC                           float64  `json:"games_headquarters_MURDER_CLASSIC"`
				WinsHeadquartersMURDERCLASSIC                            float64  `json:"wins_headquarters_MURDER_CLASSIC"`
				DetectiveWinsHeadquartersMURDERCLASSIC                   float64  `json:"detective_wins_headquarters_MURDER_CLASSIC"`
				DetectiveWinsHeadquarters                                float64  `json:"detective_wins_headquarters"`
				CoinsPickedupHeadquartersMURDERCLASSIC                   float64  `json:"coins_pickedup_headquarters_MURDER_CLASSIC"`
				KnifeKillsGoldRushMURDERASSASSINS                        float64  `json:"knife_kills_gold_rush_MURDER_ASSASSINS"`
				GamesGoldRushMURDERASSASSINS                             float64  `json:"games_gold_rush_MURDER_ASSASSINS"`
				CoinsPickedupGoldRushMURDERASSASSINS                     float64  `json:"coins_pickedup_gold_rush_MURDER_ASSASSINS"`
				KillsGoldRushMURDERASSASSINS                             float64  `json:"kills_gold_rush_MURDER_ASSASSINS"`
				WinsGoldRushMURDERASSASSINS                              float64  `json:"wins_gold_rush_MURDER_ASSASSINS"`
				WinsHeadquartersMURDERASSASSINS                          float64  `json:"wins_headquarters_MURDER_ASSASSINS"`
				KillsHeadquarters                                        float64  `json:"kills_headquarters"`
				KillsHeadquartersMURDERASSASSINS                         float64  `json:"kills_headquarters_MURDER_ASSASSINS"`
				CoinsPickedupHeadquartersMURDERASSASSINS                 float64  `json:"coins_pickedup_headquarters_MURDER_ASSASSINS"`
				KnifeKillsHeadquarters                                   float64  `json:"knife_kills_headquarters"`
				GamesHeadquartersMURDERASSASSINS                         float64  `json:"games_headquarters_MURDER_ASSASSINS"`
				ThrownKnifeKillsHeadquarters                             float64  `json:"thrown_knife_kills_headquarters"`
				ThrownKnifeKillsHeadquartersMURDERASSASSINS              float64  `json:"thrown_knife_kills_headquarters_MURDER_ASSASSINS"`
				KnifeKillsHeadquartersMURDERASSASSINS                    float64  `json:"knife_kills_headquarters_MURDER_ASSASSINS"`
				ThrownKnifeKillsCruiseShipMURDERASSASSINS                float64  `json:"thrown_knife_kills_cruise_ship_MURDER_ASSASSINS"`
				WinsCruiseShipMURDERASSASSINS                            float64  `json:"wins_cruise_ship_MURDER_ASSASSINS"`
				ThrownKnifeKillsGoldRushMURDERASSASSINS                  float64  `json:"thrown_knife_kills_gold_rush_MURDER_ASSASSINS"`
				GamesTowerfallMURDERASSASSINS                            float64  `json:"games_towerfall_MURDER_ASSASSINS"`
				CoinsPickedupTowerfallMURDERASSASSINS                    float64  `json:"coins_pickedup_towerfall_MURDER_ASSASSINS"`
				KnifeKillsTowerfallMURDERASSASSINS                       float64  `json:"knife_kills_towerfall_MURDER_ASSASSINS"`
				WinsTowerfallMURDERASSASSINS                             float64  `json:"wins_towerfall_MURDER_ASSASSINS"`
				KillsTowerfallMURDERASSASSINS                            float64  `json:"kills_towerfall_MURDER_ASSASSINS"`
				CoinsPickedupArchivesMURDERASSASSINS                     float64  `json:"coins_pickedup_archives_MURDER_ASSASSINS"`
				DeathsArchivesMURDERASSASSINS                            float64  `json:"deaths_archives_MURDER_ASSASSINS"`
				ThrownKnifeKillsArchivesMURDERASSASSINS                  float64  `json:"thrown_knife_kills_archives_MURDER_ASSASSINS"`
				KnifeKillsArchivesMURDERASSASSINS                        float64  `json:"knife_kills_archives_MURDER_ASSASSINS"`
				KillsArchivesMURDERASSASSINS                             float64  `json:"kills_archives_MURDER_ASSASSINS"`
				GamesArchivesMURDERASSASSINS                             float64  `json:"games_archives_MURDER_ASSASSINS"`
				DeathsAncientTombMURDERASSASSINS                         float64  `json:"deaths_ancient_tomb_MURDER_ASSASSINS"`
				KillsTransportMURDERASSASSINS                            float64  `json:"kills_transport_MURDER_ASSASSINS"`
				CoinsPickedupTransportMURDERASSASSINS                    float64  `json:"coins_pickedup_transport_MURDER_ASSASSINS"`
				KnifeKillsTransport                                      float64  `json:"knife_kills_transport"`
				GamesTransportMURDERASSASSINS                            float64  `json:"games_transport_MURDER_ASSASSINS"`
				ThrownKnifeKillsTransport                                float64  `json:"thrown_knife_kills_transport"`
				WinsTransportMURDERASSASSINS                             float64  `json:"wins_transport_MURDER_ASSASSINS"`
				ThrownKnifeKillsTransportMURDERASSASSINS                 float64  `json:"thrown_knife_kills_transport_MURDER_ASSASSINS"`
				KnifeKillsTransportMURDERASSASSINS                       float64  `json:"knife_kills_transport_MURDER_ASSASSINS"`
				WinsHypixelWorldMURDERASSASSINS                          float64  `json:"wins_hypixel_world_MURDER_ASSASSINS"`
				CoinsPickedupHypixelWorldMURDERASSASSINS                 float64  `json:"coins_pickedup_hypixel_world_MURDER_ASSASSINS"`
				GamesHypixelWorldMURDERASSASSINS                         float64  `json:"games_hypixel_world_MURDER_ASSASSINS"`
				ThrownKnifeKillsHypixelWorldMURDERASSASSINS              float64  `json:"thrown_knife_kills_hypixel_world_MURDER_ASSASSINS"`
				KillsHypixelWorldMURDERASSASSINS                         float64  `json:"kills_hypixel_world_MURDER_ASSASSINS"`
				KnifeKillsHypixelWorldMURDERASSASSINS                    float64  `json:"knife_kills_hypixel_world_MURDER_ASSASSINS"`
				DeathsHypixelWorldMURDERASSASSINS                        float64  `json:"deaths_hypixel_world_MURDER_ASSASSINS"`
				WinsArchivesMURDERASSASSINS                              float64  `json:"wins_archives_MURDER_ASSASSINS"`
				BowKillsHypixelWorldMURDERASSASSINS                      float64  `json:"bow_kills_hypixel_world_MURDER_ASSASSINS"`
				BowKillsMURDERASSASSINS                                  float64  `json:"bow_kills_MURDER_ASSASSINS"`
				WinsHollywoodMURDERASSASSINS                             float64  `json:"wins_hollywood_MURDER_ASSASSINS"`
				GamesHollywoodMURDERASSASSINS                            float64  `json:"games_hollywood_MURDER_ASSASSINS"`
				CoinsPickedupHollywoodMURDERASSASSINS                    float64  `json:"coins_pickedup_hollywood_MURDER_ASSASSINS"`
				ThrownKnifeKillsHollywoodMURDERASSASSINS                 float64  `json:"thrown_knife_kills_hollywood_MURDER_ASSASSINS"`
				KillsHollywoodMURDERASSASSINS                            float64  `json:"kills_hollywood_MURDER_ASSASSINS"`
				KnifeKillsHollywoodMURDERASSASSINS                       float64  `json:"knife_kills_hollywood_MURDER_ASSASSINS"`
				DeathsGoldRushMURDERASSASSINS                            float64  `json:"deaths_gold_rush_MURDER_ASSASSINS"`
				KnifeKillsLibraryMURDERASSASSINS                         float64  `json:"knife_kills_library_MURDER_ASSASSINS"`
				KillsLibraryMURDERASSASSINS                              float64  `json:"kills_library_MURDER_ASSASSINS"`
				ThrownKnifeKillsLibraryMURDERASSASSINS                   float64  `json:"thrown_knife_kills_library_MURDER_ASSASSINS"`
				ThrownKnifeKillsLibrary                                  float64  `json:"thrown_knife_kills_library"`
				CoinsPickedupLibraryMURDERASSASSINS                      float64  `json:"coins_pickedup_library_MURDER_ASSASSINS"`
				GamesLibraryMURDERASSASSINS                              float64  `json:"games_library_MURDER_ASSASSINS"`
				WinsLibraryMURDERASSASSINS                               float64  `json:"wins_library_MURDER_ASSASSINS"`
				DeathsHollywoodMURDERASSASSINS                           float64  `json:"deaths_hollywood_MURDER_ASSASSINS"`
				DeathsHeadquarters                                       float64  `json:"deaths_headquarters"`
				DeathsHeadquartersMURDERASSASSINS                        float64  `json:"deaths_headquarters_MURDER_ASSASSINS"`
				BowKillsCruiseShipMURDERASSASSINS                        float64  `json:"bow_kills_cruise_ship_MURDER_ASSASSINS"`
				ThrownKnifeKillsTowerfallMURDERASSASSINS                 float64  `json:"thrown_knife_kills_towerfall_MURDER_ASSASSINS"`
				DeathsTowerfallMURDERASSASSINS                           float64  `json:"deaths_towerfall_MURDER_ASSASSINS"`
				DeathsTowerfall                                          float64  `json:"deaths_towerfall"`
				DeathsLibraryMURDERASSASSINS                             float64  `json:"deaths_library_MURDER_ASSASSINS"`
				DeathsLibrary                                            float64  `json:"deaths_library"`
				BowKillsTowerfall                                        float64  `json:"bow_kills_towerfall"`
				BowKillsTowerfallMURDERASSASSINS                         float64  `json:"bow_kills_towerfall_MURDER_ASSASSINS"`
				DeathsCruiseShipMURDERCLASSIC                            float64  `json:"deaths_cruise_ship_MURDER_CLASSIC"`
				TotalTimeSurvivedSecondsGoldRush                         float64  `json:"total_time_survived_seconds_gold_rush"`
				TotalTimeSurvivedSeconds                                 float64  `json:"total_time_survived_seconds"`
				LongestTimeAsSurvivorSeconds                             float64  `json:"longest_time_as_survivor_seconds"`
				LongestTimeAsSurvivorSecondsGoldRush                     float64  `json:"longest_time_as_survivor_seconds_gold_rush"`
				TotalTimeSurvivedSecondsGoldRushMURDERINFECTION          float64  `json:"total_time_survived_seconds_gold_rush_MURDER_INFECTION"`
				LongestTimeAsSurvivorSecondsMURDERINFECTION              float64  `json:"longest_time_as_survivor_seconds_MURDER_INFECTION"`
				TotalTimeSurvivedSecondsMURDERINFECTION                  float64  `json:"total_time_survived_seconds_MURDER_INFECTION"`
				LongestTimeAsSurvivorSecondsGoldRushMURDERINFECTION      float64  `json:"longest_time_as_survivor_seconds_gold_rush_MURDER_INFECTION"`
				LastOneAliveGoldRush                                     float64  `json:"last_one_alive_gold_rush"`
				GamesMURDERINFECTION                                     float64  `json:"games_MURDER_INFECTION"`
				KillsAsSurvivorMURDERINFECTION                           float64  `json:"kills_as_survivor_MURDER_INFECTION"`
				KillsAsSurvivor                                          float64  `json:"kills_as_survivor"`
				GamesGoldRushMURDERINFECTION                             float64  `json:"games_gold_rush_MURDER_INFECTION"`
				KillsAsSurvivorGoldRushMURDERINFECTION                   float64  `json:"kills_as_survivor_gold_rush_MURDER_INFECTION"`
				LastOneAliveGoldRushMURDERINFECTION                      float64  `json:"last_one_alive_gold_rush_MURDER_INFECTION"`
				LastOneAliveMURDERINFECTION                              float64  `json:"last_one_alive_MURDER_INFECTION"`
				LastOneAlive                                             float64  `json:"last_one_alive"`
				KillsAsSurvivorGoldRush                                  float64  `json:"kills_as_survivor_gold_rush"`
				MmChristmasChests                                        float64  `json:"mm_christmas_chests"`
				QuickestShowdownWinTimeSecondsMURDERSHOWDOWN             float64  `json:"quickest_showdown_win_time_seconds_MURDER_SHOWDOWN"`
				QuickestShowdownWinTimeSecondsMountain                   float64  `json:"quickest_showdown_win_time_seconds_mountain"`
				QuickestShowdownWinTimeSeconds                           float64  `json:"quickest_showdown_win_time_seconds"`
				QuickestShowdownWinTimeSecondsMountainMURDERSHOWDOWN     float64  `json:"quickest_showdown_win_time_seconds_mountain_MURDER_SHOWDOWN"`
				GamesMountainMURDERSHOWDOWN                              float64  `json:"games_mountain_MURDER_SHOWDOWN"`
				WinsMountain                                             float64  `json:"wins_mountain"`
				WinsMountainMURDERSHOWDOWN                               float64  `json:"wins_mountain_MURDER_SHOWDOWN"`
				ShowdownPotg                                             float64  `json:"showdown_potg"`
				DeathsMountain                                           float64  `json:"deaths_mountain"`
				ShowdownPotgMURDERSHOWDOWN                               float64  `json:"showdown_potg_MURDER_SHOWDOWN"`
				CoinsPickedupMountain                                    float64  `json:"coins_pickedup_mountain"`
				CoinsPickedupMountainMURDERSHOWDOWN                      float64  `json:"coins_pickedup_mountain_MURDER_SHOWDOWN"`
				BowKillsMountainMURDERSHOWDOWN                           float64  `json:"bow_kills_mountain_MURDER_SHOWDOWN"`
				KillsMountainMURDERSHOWDOWN                              float64  `json:"kills_mountain_MURDER_SHOWDOWN"`
				CoinsPickedupMURDERSHOWDOWN                              float64  `json:"coins_pickedup_MURDER_SHOWDOWN"`
				BowKillsMURDERSHOWDOWN                                   float64  `json:"bow_kills_MURDER_SHOWDOWN"`
				ShowdownPotgMountainMURDERSHOWDOWN                       float64  `json:"showdown_potg_mountain_MURDER_SHOWDOWN"`
				KillsMountain                                            float64  `json:"kills_mountain"`
				KillsMURDERSHOWDOWN                                      float64  `json:"kills_MURDER_SHOWDOWN"`
				DeathsMountainMURDERSHOWDOWN                             float64  `json:"deaths_mountain_MURDER_SHOWDOWN"`
				GamesMountain                                            float64  `json:"games_mountain"`
				GamesMURDERSHOWDOWN                                      float64  `json:"games_MURDER_SHOWDOWN"`
				WinsMURDERSHOWDOWN                                       float64  `json:"wins_MURDER_SHOWDOWN"`
				BowKillsMountain                                         float64  `json:"bow_kills_mountain"`
				DeathsMURDERSHOWDOWN                                     float64  `json:"deaths_MURDER_SHOWDOWN"`
				ShowdownPotgMountain                                     float64  `json:"showdown_potg_mountain"`
				QuickestShowdownWinTimeSecondsArchives                   float64  `json:"quickest_showdown_win_time_seconds_archives"`
				QuickestShowdownWinTimeSecondsArchivesMURDERSHOWDOWN     float64  `json:"quickest_showdown_win_time_seconds_archives_MURDER_SHOWDOWN"`
				CoinsPickedupArchivesMURDERSHOWDOWN                      float64  `json:"coins_pickedup_archives_MURDER_SHOWDOWN"`
				BowKillsArchivesMURDERSHOWDOWN                           float64  `json:"bow_kills_archives_MURDER_SHOWDOWN"`
				DeathsArchivesMURDERSHOWDOWN                             float64  `json:"deaths_archives_MURDER_SHOWDOWN"`
				KnifeKillsMURDERSHOWDOWN                                 float64  `json:"knife_kills_MURDER_SHOWDOWN"`
				KillsArchivesMURDERSHOWDOWN                              float64  `json:"kills_archives_MURDER_SHOWDOWN"`
				ShowdownPotgArchivesMURDERSHOWDOWN                       float64  `json:"showdown_potg_archives_MURDER_SHOWDOWN"`
				GamesArchivesMURDERSHOWDOWN                              float64  `json:"games_archives_MURDER_SHOWDOWN"`
				WinsArchivesMURDERSHOWDOWN                               float64  `json:"wins_archives_MURDER_SHOWDOWN"`
				KnifeKillsArchivesMURDERSHOWDOWN                         float64  `json:"knife_kills_archives_MURDER_SHOWDOWN"`
				ShowdownPotgArchives                                     float64  `json:"showdown_potg_archives"`
				ShowdownPotgCruiseShip                                   float64  `json:"showdown_potg_cruise_ship"`
				WinsCruiseShipMURDERSHOWDOWN                             float64  `json:"wins_cruise_ship_MURDER_SHOWDOWN"`
				KillsCruiseShipMURDERSHOWDOWN                            float64  `json:"kills_cruise_ship_MURDER_SHOWDOWN"`
				CoinsPickedupCruiseShipMURDERSHOWDOWN                    float64  `json:"coins_pickedup_cruise_ship_MURDER_SHOWDOWN"`
				KnifeKillsCruiseShipMURDERSHOWDOWN                       float64  `json:"knife_kills_cruise_ship_MURDER_SHOWDOWN"`
				ShowdownPotgCruiseShipMURDERSHOWDOWN                     float64  `json:"showdown_potg_cruise_ship_MURDER_SHOWDOWN"`
				BowKillsCruiseShipMURDERSHOWDOWN                         float64  `json:"bow_kills_cruise_ship_MURDER_SHOWDOWN"`
				GamesCruiseShipMURDERSHOWDOWN                            float64  `json:"games_cruise_ship_MURDER_SHOWDOWN"`
				DeathsCruiseShipMURDERSHOWDOWN                           float64  `json:"deaths_cruise_ship_MURDER_SHOWDOWN"`
				KnifeKillsMountainMURDERSHOWDOWN                         float64  `json:"knife_kills_mountain_MURDER_SHOWDOWN"`
				KnifeKillsMountain                                       float64  `json:"knife_kills_mountain"`
				BowKillsTowerfallMURDERSHOWDOWN                          float64  `json:"bow_kills_towerfall_MURDER_SHOWDOWN"`
				KnifeKillsTowerfallMURDERSHOWDOWN                        float64  `json:"knife_kills_towerfall_MURDER_SHOWDOWN"`
				CoinsPickedupTowerfallMURDERSHOWDOWN                     float64  `json:"coins_pickedup_towerfall_MURDER_SHOWDOWN"`
				ShowdownPotgTowerfall                                    float64  `json:"showdown_potg_towerfall"`
				GamesTowerfallMURDERSHOWDOWN                             float64  `json:"games_towerfall_MURDER_SHOWDOWN"`
				DeathsTowerfallMURDERSHOWDOWN                            float64  `json:"deaths_towerfall_MURDER_SHOWDOWN"`
				KillsTowerfallMURDERSHOWDOWN                             float64  `json:"kills_towerfall_MURDER_SHOWDOWN"`
				WinsTowerfallMURDERSHOWDOWN                              float64  `json:"wins_towerfall_MURDER_SHOWDOWN"`
				ShowdownPotgTowerfallMURDERSHOWDOWN                      float64  `json:"showdown_potg_towerfall_MURDER_SHOWDOWN"`
				GamesHypixelWorldMURDERSHOWDOWN                          float64  `json:"games_hypixel_world_MURDER_SHOWDOWN"`
				KillsHypixelWorldMURDERSHOWDOWN                          float64  `json:"kills_hypixel_world_MURDER_SHOWDOWN"`
				CoinsPickedupHypixelWorldMURDERSHOWDOWN                  float64  `json:"coins_pickedup_hypixel_world_MURDER_SHOWDOWN"`
				ShowdownPotgHypixelWorldMURDERSHOWDOWN                   float64  `json:"showdown_potg_hypixel_world_MURDER_SHOWDOWN"`
				BowKillsHypixelWorldMURDERSHOWDOWN                       float64  `json:"bow_kills_hypixel_world_MURDER_SHOWDOWN"`
				KnifeKillsHypixelWorldMURDERSHOWDOWN                     float64  `json:"knife_kills_hypixel_world_MURDER_SHOWDOWN"`
				DeathsHypixelWorldMURDERSHOWDOWN                         float64  `json:"deaths_hypixel_world_MURDER_SHOWDOWN"`
				ShowdownPotgHypixelWorld                                 float64  `json:"showdown_potg_hypixel_world"`
				GamesLibraryMURDERSHOWDOWN                               float64  `json:"games_library_MURDER_SHOWDOWN"`
				DeathsLibraryMURDERSHOWDOWN                              float64  `json:"deaths_library_MURDER_SHOWDOWN"`
				KnifeKillsLibraryMURDERSHOWDOWN                          float64  `json:"knife_kills_library_MURDER_SHOWDOWN"`
				KillsLibraryMURDERSHOWDOWN                               float64  `json:"kills_library_MURDER_SHOWDOWN"`
				CoinsPickedupLibraryMURDERSHOWDOWN                       float64  `json:"coins_pickedup_library_MURDER_SHOWDOWN"`
				BowKillsLibraryMURDERSHOWDOWN                            float64  `json:"bow_kills_library_MURDER_SHOWDOWN"`
				WinsLibraryMURDERSHOWDOWN                                float64  `json:"wins_library_MURDER_SHOWDOWN"`
				WinsHypixelWorldMURDERSHOWDOWN                           float64  `json:"wins_hypixel_world_MURDER_SHOWDOWN"`
				CoinsPickedupAncientTombMURDERSHOWDOWN                   float64  `json:"coins_pickedup_ancient_tomb_MURDER_SHOWDOWN"`
				KnifeKillsAncientTombMURDERSHOWDOWN                      float64  `json:"knife_kills_ancient_tomb_MURDER_SHOWDOWN"`
				GamesAncientTombMURDERSHOWDOWN                           float64  `json:"games_ancient_tomb_MURDER_SHOWDOWN"`
				BowKillsAncientTombMURDERSHOWDOWN                        float64  `json:"bow_kills_ancient_tomb_MURDER_SHOWDOWN"`
				KillsAncientTombMURDERSHOWDOWN                           float64  `json:"kills_ancient_tomb_MURDER_SHOWDOWN"`
				DeathsAncientTombMURDERSHOWDOWN                          float64  `json:"deaths_ancient_tomb_MURDER_SHOWDOWN"`
				ShowdownPotgHollywood                                    float64  `json:"showdown_potg_hollywood"`
				GamesHollywoodMURDERSHOWDOWN                             float64  `json:"games_hollywood_MURDER_SHOWDOWN"`
				WinsHollywoodMURDERSHOWDOWN                              float64  `json:"wins_hollywood_MURDER_SHOWDOWN"`
				BowKillsHollywoodMURDERSHOWDOWN                          float64  `json:"bow_kills_hollywood_MURDER_SHOWDOWN"`
				ShowdownPotgHollywoodMURDERSHOWDOWN                      float64  `json:"showdown_potg_hollywood_MURDER_SHOWDOWN"`
				KnifeKillsHollywoodMURDERSHOWDOWN                        float64  `json:"knife_kills_hollywood_MURDER_SHOWDOWN"`
				KillsHollywoodMURDERSHOWDOWN                             float64  `json:"kills_hollywood_MURDER_SHOWDOWN"`
				CoinsPickedupHollywoodMURDERSHOWDOWN                     float64  `json:"coins_pickedup_hollywood_MURDER_SHOWDOWN"`
				QuickestShowdownWinTimeSecondsHypixelWorld               float64  `json:"quickest_showdown_win_time_seconds_hypixel_world"`
				QuickestShowdownWinTimeSecondsHypixelWorldMURDERSHOWDOWN float64  `json:"quickest_showdown_win_time_seconds_hypixel_world_MURDER_SHOWDOWN"`
				CoinsPickedupHeadquartersMURDERSHOWDOWN                  float64  `json:"coins_pickedup_headquarters_MURDER_SHOWDOWN"`
				KnifeKillsHeadquartersMURDERSHOWDOWN                     float64  `json:"knife_kills_headquarters_MURDER_SHOWDOWN"`
				BowKillsHeadquartersMURDERSHOWDOWN                       float64  `json:"bow_kills_headquarters_MURDER_SHOWDOWN"`
				GamesHeadquartersMURDERSHOWDOWN                          float64  `json:"games_headquarters_MURDER_SHOWDOWN"`
				KillsHeadquartersMURDERSHOWDOWN                          float64  `json:"kills_headquarters_MURDER_SHOWDOWN"`
				WinsHeadquartersMURDERSHOWDOWN                           float64  `json:"wins_headquarters_MURDER_SHOWDOWN"`
				DeathsHeadquartersMURDERSHOWDOWN                         float64  `json:"deaths_headquarters_MURDER_SHOWDOWN"`
				BowKillsHeadquarters                                     float64  `json:"bow_kills_headquarters"`
				DeathsHollywoodMURDERSHOWDOWN                            float64  `json:"deaths_hollywood_MURDER_SHOWDOWN"`
				QuickestDetectiveWinTimeSecondsMountainMURDERCLASSIC     float64  `json:"quickest_detective_win_time_seconds_mountain_MURDER_CLASSIC"`
				QuickestDetectiveWinTimeSecondsMURDERCLASSIC             float64  `json:"quickest_detective_win_time_seconds_MURDER_CLASSIC"`
				QuickestDetectiveWinTimeSecondsMountain                  float64  `json:"quickest_detective_win_time_seconds_mountain"`
				QuickestDetectiveWinTimeSeconds                          float64  `json:"quickest_detective_win_time_seconds"`
				CoinsPickedupMountainMURDERCLASSIC                       float64  `json:"coins_pickedup_mountain_MURDER_CLASSIC"`
				DetectiveWinsMountain                                    float64  `json:"detective_wins_mountain"`
				GamesMountainMURDERCLASSIC                               float64  `json:"games_mountain_MURDER_CLASSIC"`
				WinsMountainMURDERCLASSIC                                float64  `json:"wins_mountain_MURDER_CLASSIC"`
				DetectiveWinsMountainMURDERCLASSIC                       float64  `json:"detective_wins_mountain_MURDER_CLASSIC"`
				DetectiveWinsTowerfall                                   float64  `json:"detective_wins_towerfall"`
				DetectiveWinsTowerfallMURDERCLASSIC                      float64  `json:"detective_wins_towerfall_MURDER_CLASSIC"`
				QuickestMurdererWinTimeSecondsLibraryMURDERCLASSIC       float64  `json:"quickest_murderer_win_time_seconds_library_MURDER_CLASSIC"`
				QuickestMurdererWinTimeSeconds                           float64  `json:"quickest_murderer_win_time_seconds"`
				QuickestMurdererWinTimeSecondsLibrary                    float64  `json:"quickest_murderer_win_time_seconds_library"`
				QuickestMurdererWinTimeSecondsMURDERCLASSIC              float64  `json:"quickest_murderer_win_time_seconds_MURDER_CLASSIC"`
				ThrownKnifeKillsLibraryMURDERCLASSIC                     float64  `json:"thrown_knife_kills_library_MURDER_CLASSIC"`
				DeathsLibraryMURDERCLASSIC                               float64  `json:"deaths_library_MURDER_CLASSIC"`
				ChestHistoryNew                                          []string `json:"chest_history_new"`
				CoinsPickedupSanPeratico                                 float64  `json:"coins_pickedup_san_peratico"`
				WinsSanPeratico                                          float64  `json:"wins_san_peratico"`
				KnifeKillsSanPeraticoMURDERSHOWDOWN                      float64  `json:"knife_kills_san_peratico_MURDER_SHOWDOWN"`
				KillsSanPeratico                                         float64  `json:"kills_san_peratico"`
				BowKillsSanPeraticoMURDERSHOWDOWN                        float64  `json:"bow_kills_san_peratico_MURDER_SHOWDOWN"`
				KillsSanPeraticoMURDERSHOWDOWN                           float64  `json:"kills_san_peratico_MURDER_SHOWDOWN"`
				GamesSanPeratico                                         float64  `json:"games_san_peratico"`
				BowKillsSanPeratico                                      float64  `json:"bow_kills_san_peratico"`
				DeathsSanPeraticoMURDERSHOWDOWN                          float64  `json:"deaths_san_peratico_MURDER_SHOWDOWN"`
				KnifeKillsSanPeratico                                    float64  `json:"knife_kills_san_peratico"`
				WinsSanPeraticoMURDERSHOWDOWN                            float64  `json:"wins_san_peratico_MURDER_SHOWDOWN"`
				GamesSanPeraticoMURDERSHOWDOWN                           float64  `json:"games_san_peratico_MURDER_SHOWDOWN"`
				CoinsPickedupSanPeraticoMURDERSHOWDOWN                   float64  `json:"coins_pickedup_san_peratico_MURDER_SHOWDOWN"`
				DeathsSanPeratico                                        float64  `json:"deaths_san_peratico"`
				ShowdownPotgSanPeratico                                  float64  `json:"showdown_potg_san_peratico"`
				ShowdownPotgSanPeraticoMURDERSHOWDOWN                    float64  `json:"showdown_potg_san_peratico_MURDER_SHOWDOWN"`
				ShowdownPotgAncientTombMURDERSHOWDOWN                    float64  `json:"showdown_potg_ancient_tomb_MURDER_SHOWDOWN"`
				WinsAncientTombMURDERSHOWDOWN                            float64  `json:"wins_ancient_tomb_MURDER_SHOWDOWN"`
				ShowdownPotgAncientTomb                                  float64  `json:"showdown_potg_ancient_tomb"`
				ActiveGesture                                            string   `json:"active_gesture"`
				GamesGoldRushMURDERSHOWDOWN                              float64  `json:"games_gold_rush_MURDER_SHOWDOWN"`
				WinsGoldRushMURDERSHOWDOWN                               float64  `json:"wins_gold_rush_MURDER_SHOWDOWN"`
				KillsGoldRushMURDERSHOWDOWN                              float64  `json:"kills_gold_rush_MURDER_SHOWDOWN"`
				CoinsPickedupGoldRushMURDERSHOWDOWN                      float64  `json:"coins_pickedup_gold_rush_MURDER_SHOWDOWN"`
				DeathsGoldRushMURDERSHOWDOWN                             float64  `json:"deaths_gold_rush_MURDER_SHOWDOWN"`
				KnifeKillsGoldRushMURDERSHOWDOWN                         float64  `json:"knife_kills_gold_rush_MURDER_SHOWDOWN"`
				BowKillsGoldRushMURDERSHOWDOWN                           float64  `json:"bow_kills_gold_rush_MURDER_SHOWDOWN"`
				QuickestShowdownWinTimeSecondsSanPeratico                float64  `json:"quickest_showdown_win_time_seconds_san_peratico"`
				QuickestShowdownWinTimeSecondsSanPeraticoMURDERSHOWDOWN  float64  `json:"quickest_showdown_win_time_seconds_san_peratico_MURDER_SHOWDOWN"`
				WinsSanPeraticoMURDERCLASSIC                             float64  `json:"wins_san_peratico_MURDER_CLASSIC"`
				GamesSanPeraticoMURDERCLASSIC                            float64  `json:"games_san_peratico_MURDER_CLASSIC"`
				CoinsPickedupSanPeraticoMURDERCLASSIC                    float64  `json:"coins_pickedup_san_peratico_MURDER_CLASSIC"`
				BowKillsTowerfallMURDERCLASSIC                           float64  `json:"bow_kills_towerfall_MURDER_CLASSIC"`
				WasHeroTowerfall                                         float64  `json:"was_hero_towerfall"`
				WasHeroTowerfallMURDERCLASSIC                            float64  `json:"was_hero_towerfall_MURDER_CLASSIC"`
				DeathsHeadquartersMURDERCLASSIC                          float64  `json:"deaths_headquarters_MURDER_CLASSIC"`
				KillsAsMurdererHeadquarters                              float64  `json:"kills_as_murderer_headquarters"`
				KillsAsMurdererHeadquartersMURDERCLASSIC                 float64  `json:"kills_as_murderer_headquarters_MURDER_CLASSIC"`
				KillsHeadquartersMURDERCLASSIC                           float64  `json:"kills_headquarters_MURDER_CLASSIC"`
				KnifeKillsHeadquartersMURDERCLASSIC                      float64  `json:"knife_kills_headquarters_MURDER_CLASSIC"`
				ThrownKnifeKillsHeadquartersMURDERCLASSIC                float64  `json:"thrown_knife_kills_headquarters_MURDER_CLASSIC"`
				BowKillsArchivesTopFloor                                 float64  `json:"bow_kills_archives_top_floor"`
				BowKillsArchivesTopFloorMURDERCLASSIC                    float64  `json:"bow_kills_archives_top_floor_MURDER_CLASSIC"`
				CoinsPickedupArchivesTopFloor                            float64  `json:"coins_pickedup_archives_top_floor"`
				CoinsPickedupArchivesTopFloorMURDERCLASSIC               float64  `json:"coins_pickedup_archives_top_floor_MURDER_CLASSIC"`
				GamesArchivesTopFloor                                    float64  `json:"games_archives_top_floor"`
				GamesArchivesTopFloorMURDERCLASSIC                       float64  `json:"games_archives_top_floor_MURDER_CLASSIC"`
				KillsArchivesTopFloor                                    float64  `json:"kills_archives_top_floor"`
				KillsArchivesTopFloorMURDERCLASSIC                       float64  `json:"kills_archives_top_floor_MURDER_CLASSIC"`
				WasHeroArchivesTopFloor                                  float64  `json:"was_hero_archives_top_floor"`
				WasHeroArchivesTopFloorMURDERCLASSIC                     float64  `json:"was_hero_archives_top_floor_MURDER_CLASSIC"`
				WinsArchivesTopFloor                                     float64  `json:"wins_archives_top_floor"`
				WinsArchivesTopFloorMURDERCLASSIC                        float64  `json:"wins_archives_top_floor_MURDER_CLASSIC"`
				CoinsPickedupWidowSDen                                   float64  `json:"coins_pickedup_widow\u0027s_den"`
				CoinsPickedupWidowSDenMURDERCLASSIC                      float64  `json:"coins_pickedup_widow\u0027s_den_MURDER_CLASSIC"`
				DeathsWidowSDen                                          float64  `json:"deaths_widow\u0027s_den"`
				DeathsWidowSDenMURDERCLASSIC                             float64  `json:"deaths_widow\u0027s_den_MURDER_CLASSIC"`
				GamesWidowSDen                                           float64  `json:"games_widow\u0027s_den"`
				GamesWidowSDenMURDERCLASSIC                              float64  `json:"games_widow\u0027s_den_MURDER_CLASSIC"`
				WinsWidowSDen                                            float64  `json:"wins_widow\u0027s_den"`
				WinsWidowSDenMURDERCLASSIC                               float64  `json:"wins_widow\u0027s_den_MURDER_CLASSIC"`
				CoinsPickedupSanPeraticoV2                               float64  `json:"coins_pickedup_san_peratico_v2"`
				CoinsPickedupSanPeraticoV2MURDERCLASSIC                  float64  `json:"coins_pickedup_san_peratico_v2_MURDER_CLASSIC"`
				DeathsSanPeraticoV2                                      float64  `json:"deaths_san_peratico_v2"`
				DeathsSanPeraticoV2MURDERCLASSIC                         float64  `json:"deaths_san_peratico_v2_MURDER_CLASSIC"`
				GamesSanPeraticoV2                                       float64  `json:"games_san_peratico_v2"`
				GamesSanPeraticoV2MURDERCLASSIC                          float64  `json:"games_san_peratico_v2_MURDER_CLASSIC"`
				WinsSanPeraticoV2                                        float64  `json:"wins_san_peratico_v2"`
				WinsSanPeraticoV2MURDERCLASSIC                           float64  `json:"wins_san_peratico_v2_MURDER_CLASSIC"`
				MmHalloweenChests                                        float64  `json:"mm_halloween_chests"`
				CoinsPickedupMURDERDOUBLEUP                              float64  `json:"coins_pickedup_MURDER_DOUBLE_UP"`
				CoinsPickedupMountainMURDERDOUBLEUP                      float64  `json:"coins_pickedup_mountain_MURDER_DOUBLE_UP"`
				DeathsMURDERDOUBLEUP                                     float64  `json:"deaths_MURDER_DOUBLE_UP"`
				DeathsMountainMURDERDOUBLEUP                             float64  `json:"deaths_mountain_MURDER_DOUBLE_UP"`
				GamesMURDERDOUBLEUP                                      float64  `json:"games_MURDER_DOUBLE_UP"`
				GamesMountainMURDERDOUBLEUP                              float64  `json:"games_mountain_MURDER_DOUBLE_UP"`
				WinsMURDERDOUBLEUP                                       float64  `json:"wins_MURDER_DOUBLE_UP"`
				WinsMountainMURDERDOUBLEUP                               float64  `json:"wins_mountain_MURDER_DOUBLE_UP"`
				CoinsPickedupArchivesTopFloorMURDERDOUBLEUP              float64  `json:"coins_pickedup_archives_top_floor_MURDER_DOUBLE_UP"`
				GamesArchivesTopFloorMURDERDOUBLEUP                      float64  `json:"games_archives_top_floor_MURDER_DOUBLE_UP"`
				KnifeKillsMURDERDOUBLEUP                                 float64  `json:"knife_kills_MURDER_DOUBLE_UP"`
				KnifeKillsArchivesTopFloor                               float64  `json:"knife_kills_archives_top_floor"`
				KnifeKillsArchivesTopFloorMURDERDOUBLEUP                 float64  `json:"knife_kills_archives_top_floor_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsMURDERDOUBLEUP                           float64  `json:"thrown_knife_kills_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsArchivesTopFloor                         float64  `json:"thrown_knife_kills_archives_top_floor"`
				ThrownKnifeKillsArchivesTopFloorMURDERDOUBLEUP           float64  `json:"thrown_knife_kills_archives_top_floor_MURDER_DOUBLE_UP"`
				CoinsPickedupSkywayPier                                  float64  `json:"coins_pickedup_skyway_pier"`
				CoinsPickedupSkywayPierMURDERCLASSIC                     float64  `json:"coins_pickedup_skyway_pier_MURDER_CLASSIC"`
				GamesSkywayPier                                          float64  `json:"games_skyway_pier"`
				GamesSkywayPierMURDERCLASSIC                             float64  `json:"games_skyway_pier_MURDER_CLASSIC"`
				CoinsPickedupSnowfall                                    float64  `json:"coins_pickedup_snowfall"`
				CoinsPickedupSnowfallMURDERCLASSIC                       float64  `json:"coins_pickedup_snowfall_MURDER_CLASSIC"`
				GamesSnowfall                                            float64  `json:"games_snowfall"`
				GamesSnowfallMURDERCLASSIC                               float64  `json:"games_snowfall_MURDER_CLASSIC"`
				CoinsPickedupAncientTombMURDERDOUBLEUP                   float64  `json:"coins_pickedup_ancient_tomb_MURDER_DOUBLE_UP"`
				GamesAncientTombMURDERDOUBLEUP                           float64  `json:"games_ancient_tomb_MURDER_DOUBLE_UP"`
				KnifeKillsAncientTombMURDERDOUBLEUP                      float64  `json:"knife_kills_ancient_tomb_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsAncientTombMURDERDOUBLEUP                float64  `json:"thrown_knife_kills_ancient_tomb_MURDER_DOUBLE_UP"`
				CoinsPickedupSanPeraticoV2MURDERDOUBLEUP                 float64  `json:"coins_pickedup_san_peratico_v2_MURDER_DOUBLE_UP"`
				GamesSanPeraticoV2MURDERDOUBLEUP                         float64  `json:"games_san_peratico_v2_MURDER_DOUBLE_UP"`
				WasHeroMURDERDOUBLEUP                                    float64  `json:"was_hero_MURDER_DOUBLE_UP"`
				WasHeroSanPeraticoV2                                     float64  `json:"was_hero_san_peratico_v2"`
				WasHeroSanPeraticoV2MURDERDOUBLEUP                       float64  `json:"was_hero_san_peratico_v2_MURDER_DOUBLE_UP"`
				KnifeKillsSanPeraticoV2                                  float64  `json:"knife_kills_san_peratico_v2"`
				KnifeKillsSanPeraticoV2MURDERDOUBLEUP                    float64  `json:"knife_kills_san_peratico_v2_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsSanPeraticoV2                            float64  `json:"thrown_knife_kills_san_peratico_v2"`
				ThrownKnifeKillsSanPeraticoV2MURDERDOUBLEUP              float64  `json:"thrown_knife_kills_san_peratico_v2_MURDER_DOUBLE_UP"`
				BowKillsMURDERDOUBLEUP                                   float64  `json:"bow_kills_MURDER_DOUBLE_UP"`
				BowKillsHeadquartersMURDERDOUBLEUP                       float64  `json:"bow_kills_headquarters_MURDER_DOUBLE_UP"`
				CoinsPickedupHeadquartersMURDERDOUBLEUP                  float64  `json:"coins_pickedup_headquarters_MURDER_DOUBLE_UP"`
				GamesHeadquartersMURDERDOUBLEUP                          float64  `json:"games_headquarters_MURDER_DOUBLE_UP"`
				KnifeKillsHeadquartersMURDERDOUBLEUP                     float64  `json:"knife_kills_headquarters_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsHeadquartersMURDERDOUBLEUP               float64  `json:"thrown_knife_kills_headquarters_MURDER_DOUBLE_UP"`
				BowKillsSanPeraticoV2                                    float64  `json:"bow_kills_san_peratico_v2"`
				BowKillsSanPeraticoV2MURDERDOUBLEUP                      float64  `json:"bow_kills_san_peratico_v2_MURDER_DOUBLE_UP"`
				WinsSkywayPier                                           float64  `json:"wins_skyway_pier"`
				WinsSkywayPierMURDERCLASSIC                              float64  `json:"wins_skyway_pier_MURDER_CLASSIC"`
				CoinsPickedupCruiseShipMURDERDOUBLEUP                    float64  `json:"coins_pickedup_cruise_ship_MURDER_DOUBLE_UP"`
				GamesCruiseShipMURDERDOUBLEUP                            float64  `json:"games_cruise_ship_MURDER_DOUBLE_UP"`
				KillsMURDERDOUBLEUP                                      float64  `json:"kills_MURDER_DOUBLE_UP"`
				KillsAsMurdererMURDERDOUBLEUP                            float64  `json:"kills_as_murderer_MURDER_DOUBLE_UP"`
				KillsAsMurdererCruiseShipMURDERDOUBLEUP                  float64  `json:"kills_as_murderer_cruise_ship_MURDER_DOUBLE_UP"`
				KillsCruiseShipMURDERDOUBLEUP                            float64  `json:"kills_cruise_ship_MURDER_DOUBLE_UP"`
				KnifeKillsCruiseShipMURDERDOUBLEUP                       float64  `json:"knife_kills_cruise_ship_MURDER_DOUBLE_UP"`
				MurdererWinsMURDERDOUBLEUP                               float64  `json:"murderer_wins_MURDER_DOUBLE_UP"`
				MurdererWinsCruiseShipMURDERDOUBLEUP                     float64  `json:"murderer_wins_cruise_ship_MURDER_DOUBLE_UP"`
				QuickestMurdererWinTimeSecondsMURDERDOUBLEUP             float64  `json:"quickest_murderer_win_time_seconds_MURDER_DOUBLE_UP"`
				QuickestMurdererWinTimeSecondsCruiseShip                 float64  `json:"quickest_murderer_win_time_seconds_cruise_ship"`
				QuickestMurdererWinTimeSecondsCruiseShipMURDERDOUBLEUP   float64  `json:"quickest_murderer_win_time_seconds_cruise_ship_MURDER_DOUBLE_UP"`
				ThrownKnifeKillsCruiseShipMURDERDOUBLEUP                 float64  `json:"thrown_knife_kills_cruise_ship_MURDER_DOUBLE_UP"`
				WinsCruiseShipMURDERDOUBLEUP                             float64  `json:"wins_cruise_ship_MURDER_DOUBLE_UP"`
				BowKillsMountainMURDERDOUBLEUP                           float64  `json:"bow_kills_mountain_MURDER_DOUBLE_UP"`
				KillsMountainMURDERDOUBLEUP                              float64  `json:"kills_mountain_MURDER_DOUBLE_UP"`
				CoinsPickedupGoldRushMURDERDOUBLEUP                      float64  `json:"coins_pickedup_gold_rush_MURDER_DOUBLE_UP"`
				GamesGoldRushMURDERDOUBLEUP                              float64  `json:"games_gold_rush_MURDER_DOUBLE_UP"`
				WasHeroGoldRushMURDERDOUBLEUP                            float64  `json:"was_hero_gold_rush_MURDER_DOUBLE_UP"`
				WinsGoldRushMURDERDOUBLEUP                               float64  `json:"wins_gold_rush_MURDER_DOUBLE_UP"`
			} `json:"MurderMystery"`
			BuildBattle struct {
				WinsSoloNormal     float64  `json:"wins_solo_normal"`
				WinsTeamsNormal    float64  `json:"wins_teams_normal"`
				Wins               float64  `json:"wins"`
				BuildbattleLoadout []string `json:"buildbattle_loadout"`
				Coins              float64  `json:"coins"`
				WinsGuessTheBuild  float64  `json:"wins_guess_the_build"`
				GamesPlayed        float64  `json:"games_played"`
				Score              float64  `json:"score"`
				MonthlyCoinsA      float64  `json:"monthly_coins_a"`
				WeeklyCoinsB       float64  `json:"weekly_coins_b"`
				CorrectGuesses     float64  `json:"correct_guesses"`
				SoloMostPoints     float64  `json:"solo_most_points"`
				WeeklyCoinsA       float64  `json:"weekly_coins_a"`
				TotalVotes         float64  `json:"total_votes"`
				Packages           []string `json:"packages"`
				SuperVotes         float64  `json:"super_votes"`
				NewSelectedHat     string   `json:"new_selected_hat"`
				MonthlyCoinsB      float64  `json:"monthly_coins_b"`
			} `json:"BuildBattle"`
			Duels struct {
				GamesPlayedDuels               float64 `json:"games_played_duels"`
				MeleeSwings                    float64 `json:"melee_swings"`
				MeleeHits                      float64 `json:"melee_hits"`
				Wins                           float64 `json:"wins"`
				RoundsPlayed                   float64 `json:"rounds_played"`
				UhcTournamentMeleeHits         float64 `json:"uhc_tournament_melee_hits"`
				UhcTournamentHealthRegenerated float64 `json:"uhc_tournament_health_regenerated"`
				UhcTournamentDamageDealt       float64 `json:"uhc_tournament_damage_dealt"`
				BowShots                       float64 `json:"bow_shots"`
				UhcTournamentKills             float64 `json:"uhc_tournament_kills"`
				Kills                          float64 `json:"kills"`
				BowHits                        float64 `json:"bow_hits"`
				UhcTournamentMeleeSwings       float64 `json:"uhc_tournament_melee_swings"`
				UhcTournamentRoundsPlayed      float64 `json:"uhc_tournament_rounds_played"`
				HealthRegenerated              float64 `json:"health_regenerated"`
				UhcTournamentBowHits           float64 `json:"uhc_tournament_bow_hits"`
				DamageDealt                    float64 `json:"damage_dealt"`
				UhcTournamentBowShots          float64 `json:"uhc_tournament_bow_shots"`
				UhcTournamentWins              float64 `json:"uhc_tournament_wins"`
				Coins                          float64 `json:"coins"`
				Losses                         float64 `json:"losses"`
				UhcTournamentLosses            float64 `json:"uhc_tournament_losses"`
				UhcTournamentDeaths            float64 `json:"uhc_tournament_deaths"`
				Deaths                         float64 `json:"deaths"`
				Selected1                      string  `json:"selected_1"`
				Selected2                      string  `json:"selected_2"`
				KitMenuOption                  string  `json:"kit_menu_option"`
				BlitzDuelsKit                  string  `json:"blitz_duels_kit"`
				ShowLbOption                   string  `json:"show_lb_option"`
				RematchOption1                 string  `json:"rematch_option_1"`
				BlitzDuelMeleeHits             float64 `json:"blitz_duel_melee_hits"`
				BlitzDuelMeleeSwings           float64 `json:"blitz_duel_melee_swings"`
				BlitzDuelRoundsPlayed          float64 `json:"blitz_duel_rounds_played"`
				BlitzDuelHealthRegenerated     float64 `json:"blitz_duel_health_regenerated"`
				BlitzDuelDamageDealt           float64 `json:"blitz_duel_damage_dealt"`
				BlitzDuelBowShots              float64 `json:"blitz_duel_bow_shots"`
				BlitzDuelBowHits               float64 `json:"blitz_duel_bow_hits"`
				BowspleefDuelRoundsPlayed      float64 `json:"bowspleef_duel_rounds_played"`
				BowspleefDuelBowShots          float64 `json:"bowspleef_duel_bow_shots"`
				DuelsRecentlyPlayed            string  `json:"duels_recently_played"`
				SwDuelsKit                     string  `json:"sw_duels_kit"`
				SwTournamentRoundsPlayed       float64 `json:"sw_tournament_rounds_played"`
				SwTournamentMeleeHits          float64 `json:"sw_tournament_melee_hits"`
				SwTournamentDamageDealt        float64 `json:"sw_tournament_damage_dealt"`
				SwTournamentMeleeSwings        float64 `json:"sw_tournament_melee_swings"`
				SwTournamentHealthRegenerated  float64 `json:"sw_tournament_health_regenerated"`
				SwTournamentKills              float64 `json:"sw_tournament_kills"`
				SwTournamentDeaths             float64 `json:"sw_tournament_deaths"`
				CurrentBlitzWinstreak          float64 `json:"current_blitz_winstreak"`
				CurrentWinstreak               float64 `json:"current_winstreak"`
				BestOverallWinstreak           float64 `json:"best_overall_winstreak"`
				BestBlitzWinstreak             float64 `json:"best_blitz_winstreak"`
				BlitzDuelWins                  float64 `json:"blitz_duel_wins"`
				BlitzDuelKills                 float64 `json:"blitz_duel_kills"`
				DuelsWinstreakBestBlitzDuel    float64 `json:"duels_winstreak_best_blitz_duel"`
				BlitzDuelLosses                float64 `json:"blitz_duel_losses"`
				BlitzDuelDeaths                float64 `json:"blitz_duel_deaths"`
				BlitzDuelLayoutHunter5         struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num6 string `json:"6"`
				} `json:"blitz_duel_layout_hunter_5"`
				ChatEnabled                       string   `json:"chat_enabled"`
				BridgeDoublesWins                 float64  `json:"bridge_doubles_wins"`
				BridgeDuelWins                    float64  `json:"bridge_duel_wins"`
				BridgeDuelDeaths                  float64  `json:"bridge_duel_deaths"`
				BridgeDoublesDeaths               float64  `json:"bridge_doubles_deaths"`
				Goals                             float64  `json:"goals"`
				BridgeDuelGoals                   float64  `json:"bridge_duel_goals"`
				BridgeDoublesGoals                float64  `json:"bridge_doubles_goals"`
				BridgeDoublesRoundsPlayed         float64  `json:"bridge_doubles_rounds_played"`
				BridgeDoublesKills                float64  `json:"bridge_doubles_kills"`
				BridgeDuelKills                   float64  `json:"bridge_duel_kills"`
				BridgeDuelRoundsPlayed            float64  `json:"bridge_duel_rounds_played"`
				Packages                          []string `json:"packages"`
				AllModesRookieTitlePrestige       float64  `json:"all_modes_rookie_title_prestige"`
				NoDebuffRookieTitlePrestige       float64  `json:"no_debuff_rookie_title_prestige"`
				UhcRookieTitlePrestige            float64  `json:"uhc_rookie_title_prestige"`
				SkywarsRookieTitlePrestige        float64  `json:"skywars_rookie_title_prestige"`
				SumoRookieTitlePrestige           float64  `json:"sumo_rookie_title_prestige"`
				TournamentRookieTitlePrestige     float64  `json:"tournament_rookie_title_prestige"`
				MegaWallsRookieTitlePrestige      float64  `json:"mega_walls_rookie_title_prestige"`
				TntGamesRookieTitlePrestige       float64  `json:"tnt_games_rookie_title_prestige"`
				BlitzRookieTitlePrestige          float64  `json:"blitz_rookie_title_prestige"`
				BridgeRookieTitlePrestige         float64  `json:"bridge_rookie_title_prestige"`
				ComboRookieTitlePrestige          float64  `json:"combo_rookie_title_prestige"`
				BowRookieTitlePrestige            float64  `json:"bow_rookie_title_prestige"`
				OpRookieTitlePrestige             float64  `json:"op_rookie_title_prestige"`
				ClassicRookieTitlePrestige        float64  `json:"classic_rookie_title_prestige"`
				Selected2New                      string   `json:"selected_2_new"`
				Selected1New                      string   `json:"selected_1_new"`
				BridgeMapWins                     []string `json:"bridgeMapWins"`
				MapsWonOn                         []string `json:"maps_won_on"`
				CurrentBridgeWinstreak            float64  `json:"current_bridge_winstreak"`
				BestWinstreakModeBridgeDoubles    float64  `json:"best_winstreak_mode_bridge_doubles"`
				BestBridgeWinstreak               float64  `json:"best_bridge_winstreak"`
				CurrentWinstreakModeBridgeDoubles float64  `json:"current_winstreak_mode_bridge_doubles"`
				DuelsChests                       float64  `json:"duels_chests"`
				BridgeDoublesBridgeDeaths         float64  `json:"bridge_doubles_bridge_deaths"`
				BlocksPlaced                      float64  `json:"blocks_placed"`
				BridgeDoublesBridgeKills          float64  `json:"bridge_doubles_bridge_kills"`
				BridgeDoublesDamageDealt          float64  `json:"bridge_doubles_damage_dealt"`
				BridgeDoublesHealthRegenerated    float64  `json:"bridge_doubles_health_regenerated"`
				BridgeDoublesBlocksPlaced         float64  `json:"bridge_doubles_blocks_placed"`
				BridgeDoublesMeleeSwings          float64  `json:"bridge_doubles_melee_swings"`
				BridgeDoublesMeleeHits            float64  `json:"bridge_doubles_melee_hits"`
				BridgeKills                       float64  `json:"bridge_kills"`
				BridgeDeaths                      float64  `json:"bridge_deaths"`
				DuelsWinstreakBestBridgeDoubles   float64  `json:"duels_winstreak_best_bridge_doubles"`
				BridgeDoublesBowShots             float64  `json:"bridge_doubles_bow_shots"`
				BridgeDoublesBowHits              float64  `json:"bridge_doubles_bow_hits"`
				BridgeFourRoundsPlayed            float64  `json:"bridge_four_rounds_played"`
				BridgeFourDamageDealt             float64  `json:"bridge_four_damage_dealt"`
				BridgeFourBlocksPlaced            float64  `json:"bridge_four_blocks_placed"`
				BridgeFourBowShots                float64  `json:"bridge_four_bow_shots"`
				BridgeFourMeleeSwings             float64  `json:"bridge_four_melee_swings"`
				BridgeFourHealthRegenerated       float64  `json:"bridge_four_health_regenerated"`
				BridgeFourMeleeHits               float64  `json:"bridge_four_melee_hits"`
				SwDuelsKitNew3                    string   `json:"sw_duels_kit_new3"`
				CurrentWinstreakModeBridgeDuel    float64  `json:"current_winstreak_mode_bridge_duel"`
				BestWinstreakModeBridgeDuel       float64  `json:"best_winstreak_mode_bridge_duel"`
				BridgeDuelHealthRegenerated       float64  `json:"bridge_duel_health_regenerated"`
				BridgeDuelDamageDealt             float64  `json:"bridge_duel_damage_dealt"`
				BridgeDuelMeleeSwings             float64  `json:"bridge_duel_melee_swings"`
				BridgeDuelBowHits                 float64  `json:"bridge_duel_bow_hits"`
				BridgeDuelBridgeDeaths            float64  `json:"bridge_duel_bridge_deaths"`
				BridgeDuelBridgeKills             float64  `json:"bridge_duel_bridge_kills"`
				BridgeDuelMeleeHits               float64  `json:"bridge_duel_melee_hits"`
				BridgeDuelBlocksPlaced            float64  `json:"bridge_duel_blocks_placed"`
				BridgeDuelBowShots                float64  `json:"bridge_duel_bow_shots"`
				DuelsWinstreakBestBridgeDuel      float64  `json:"duels_winstreak_best_bridge_duel"`
				LayoutBridgeDuelLayout            struct {
					Num0 string `json:"0"`
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
					Num4 string `json:"4"`
					Num5 string `json:"5"`
					Num7 string `json:"7"`
					Num8 string `json:"8"`
				} `json:"layout_bridge_duel_layout"`
				DuelsOpenedCommons            float64  `json:"Duels_openedCommons"`
				DuelsOpenedChests             float64  `json:"Duels_openedChests"`
				DuelsChestHistory             []string `json:"duels_chest_history"`
				DuelsOpenedRares              float64  `json:"Duels_openedRares"`
				BestWinstreakModeBlitzDuel    float64  `json:"best_winstreak_mode_blitz_duel"`
				CurrentWinstreakModeBlitzDuel float64  `json:"current_winstreak_mode_blitz_duel"`
				ArcherKitWins                 float64  `json:"archer_kit_wins"`
				BlitzDuelKitWins              float64  `json:"blitz_duel_kit_wins"`
				BlitzDuelArcherKitWins        float64  `json:"blitz_duel_archer_kit_wins"`
				KitWins                       float64  `json:"kit_wins"`
				PigmanKitWins                 float64  `json:"pigman_kit_wins"`
				BlitzDuelBlocksPlaced         float64  `json:"blitz_duel_blocks_placed"`
				BlitzDuelPigmanKitWins        float64  `json:"blitz_duel_pigman_kit_wins"`
				GolemKitWins                  float64  `json:"golem_kit_wins"`
				BlitzDuelGolemKitWins         float64  `json:"blitz_duel_golem_kit_wins"`
				BlitzDuelNecromancerKitWins   float64  `json:"blitz_duel_necromancer_kit_wins"`
				NecromancerKitWins            float64  `json:"necromancer_kit_wins"`
				ProgressMode                  string   `json:"progress_mode"`
				CurrentWinstreakModeComboDuel float64  `json:"current_winstreak_mode_combo_duel"`
				CurrentComboWinstreak         float64  `json:"current_combo_winstreak"`
				ComboDuelMeleeHits            float64  `json:"combo_duel_melee_hits"`
				ComboDuelLosses               float64  `json:"combo_duel_losses"`
				ComboDuelRoundsPlayed         float64  `json:"combo_duel_rounds_played"`
				ComboDuelHealthRegenerated    float64  `json:"combo_duel_health_regenerated"`
				GoldenApplesEaten             float64  `json:"golden_apples_eaten"`
				ComboDuelDeaths               float64  `json:"combo_duel_deaths"`
				ComboDuelMeleeSwings          float64  `json:"combo_duel_melee_swings"`
				ComboDuelGoldenApplesEaten    float64  `json:"combo_duel_golden_apples_eaten"`
				LeaderboardPageWinStreak      float64  `json:"leaderboardPage_win_streak"`
				SumoDuelMeleeSwings           float64  `json:"sumo_duel_melee_swings"`
				SumoDuelRoundsPlayed          float64  `json:"sumo_duel_rounds_played"`
				SumoDuelMeleeHits             float64  `json:"sumo_duel_melee_hits"`
				UhcDuelHealthRegenerated      float64  `json:"uhc_duel_health_regenerated"`
				UhcDuelBowShots               float64  `json:"uhc_duel_bow_shots"`
				UhcDuelMeleeSwings            float64  `json:"uhc_duel_melee_swings"`
				UhcDuelBowHits                float64  `json:"uhc_duel_bow_hits"`
				UhcDuelRoundsPlayed           float64  `json:"uhc_duel_rounds_played"`
				UhcDuelDamageDealt            float64  `json:"uhc_duel_damage_dealt"`
				UhcDuelBlocksPlaced           float64  `json:"uhc_duel_blocks_placed"`
				UhcDuelGoldenApplesEaten      float64  `json:"uhc_duel_golden_apples_eaten"`
				UhcDuelMeleeHits              float64  `json:"uhc_duel_melee_hits"`
				CurrentSumoWinstreak          float64  `json:"current_sumo_winstreak"`
				CurrentWinstreakModeSumoDuel  float64  `json:"current_winstreak_mode_sumo_duel"`
				SumoDuelLosses                float64  `json:"sumo_duel_losses"`
				DuelsRecentlyPlayed2          string   `json:"duels_recently_played2"`
				CurrentUhcWinstreak           float64  `json:"current_uhc_winstreak"`
				CurrentWinstreakModeUhcDuel   float64  `json:"current_winstreak_mode_uhc_duel"`
				UhcDuelDeaths                 float64  `json:"uhc_duel_deaths"`
				UhcDuelLosses                 float64  `json:"uhc_duel_losses"`
				SwDuelBlocksPlaced            float64  `json:"sw_duel_blocks_placed"`
				SwDuelDamageDealt             float64  `json:"sw_duel_damage_dealt"`
				SwDuelHealthRegenerated       float64  `json:"sw_duel_health_regenerated"`
				SwDuelMeleeHits               float64  `json:"sw_duel_melee_hits"`
				SwDuelMeleeSwings             float64  `json:"sw_duel_melee_swings"`
				SwDuelRoundsPlayed            float64  `json:"sw_duel_rounds_played"`
				BridgeFourBowHits             float64  `json:"bridge_four_bow_hits"`
				BestSkywarsWinstreak          float64  `json:"best_skywars_winstreak"`
				CurrentWinstreakModeSwDoubles float64  `json:"current_winstreak_mode_sw_doubles"`
				CurrentSkywarsWinstreak       float64  `json:"current_skywars_winstreak"`
				BestWinstreakModeSwDoubles    float64  `json:"best_winstreak_mode_sw_doubles"`
				ScoutKitWins                  float64  `json:"scout_kit_wins"`
				SwDoublesDamageDealt          float64  `json:"sw_doubles_damage_dealt"`
				SwDoublesHealthRegenerated    float64  `json:"sw_doubles_health_regenerated"`
				SwDoublesKills                float64  `json:"sw_doubles_kills"`
				SwDoublesKitWins              float64  `json:"sw_doubles_kit_wins"`
				SwDoublesMeleeHits            float64  `json:"sw_doubles_melee_hits"`
				SwDoublesMeleeSwings          float64  `json:"sw_doubles_melee_swings"`
				SwDoublesRoundsPlayed         float64  `json:"sw_doubles_rounds_played"`
				SwDoublesScoutKitWins         float64  `json:"sw_doubles_scout_kit_wins"`
				SwDoublesWins                 float64  `json:"sw_doubles_wins"`
				LayoutUhcDuelLayout           struct {
					Num0  string `json:"0"`
					Num1  string `json:"1"`
					Num2  string `json:"2"`
					Num3  string `json:"3"`
					Num4  string `json:"4"`
					Num5  string `json:"5"`
					Num6  string `json:"6"`
					Num7  string `json:"7"`
					Num8  string `json:"8"`
					Num9  string `json:"9"`
					Num33 string `json:"33"`
					Num34 string `json:"34"`
					Num35 string `json:"35"`
				} `json:"layout_uhc_duel_layout"`
				SwDoublesBlocksPlaced           float64 `json:"sw_doubles_blocks_placed"`
				DuelsWinstreakBestSwDoubles     float64 `json:"duels_winstreak_best_sw_doubles"`
				SwDoublesBowHits                float64 `json:"sw_doubles_bow_hits"`
				SwDoublesBowShots               float64 `json:"sw_doubles_bow_shots"`
				SwDoublesDeaths                 float64 `json:"sw_doubles_deaths"`
				SwDoublesLosses                 float64 `json:"sw_doubles_losses"`
				OpDuelDamageDealt               float64 `json:"op_duel_damage_dealt"`
				OpDuelHealthRegenerated         float64 `json:"op_duel_health_regenerated"`
				OpDuelMeleeHits                 float64 `json:"op_duel_melee_hits"`
				OpDuelMeleeSwings               float64 `json:"op_duel_melee_swings"`
				OpDuelRoundsPlayed              float64 `json:"op_duel_rounds_played"`
				ChallengesEnabled               bool    `json:"challenges_enabled"`
				BowDuelBowHits                  float64 `json:"bow_duel_bow_hits"`
				BowDuelBowShots                 float64 `json:"bow_duel_bow_shots"`
				BowDuelDamageDealt              float64 `json:"bow_duel_damage_dealt"`
				BowDuelHealthRegenerated        float64 `json:"bow_duel_health_regenerated"`
				BowDuelRoundsPlayed             float64 `json:"bow_duel_rounds_played"`
				BestClassicWinstreak            float64 `json:"best_classic_winstreak"`
				BestWinstreakModeClassicDuel    float64 `json:"best_winstreak_mode_classic_duel"`
				CurrentWinstreakModeClassicDuel float64 `json:"current_winstreak_mode_classic_duel"`
				CurrentClassicWinstreak         float64 `json:"current_classic_winstreak"`
				ClassicDuelBowHits              float64 `json:"classic_duel_bow_hits"`
				ClassicDuelBowShots             float64 `json:"classic_duel_bow_shots"`
				ClassicDuelDamageDealt          float64 `json:"classic_duel_damage_dealt"`
				ClassicDuelHealthRegenerated    float64 `json:"classic_duel_health_regenerated"`
				ClassicDuelKills                float64 `json:"classic_duel_kills"`
				ClassicDuelMeleeHits            float64 `json:"classic_duel_melee_hits"`
				ClassicDuelMeleeSwings          float64 `json:"classic_duel_melee_swings"`
				ClassicDuelRoundsPlayed         float64 `json:"classic_duel_rounds_played"`
				ClassicDuelWins                 float64 `json:"classic_duel_wins"`
				DuelsWinstreakBestClassicDuel   float64 `json:"duels_winstreak_best_classic_duel"`
				ClassicDuelDeaths               float64 `json:"classic_duel_deaths"`
				ClassicDuelLosses               float64 `json:"classic_duel_losses"`
				BestUhcWinstreak                float64 `json:"best_uhc_winstreak"`
				BestWinstreakModeUhcDuel        float64 `json:"best_winstreak_mode_uhc_duel"`
				UhcDuelKills                    float64 `json:"uhc_duel_kills"`
				UhcDuelWins                     float64 `json:"uhc_duel_wins"`
				DuelsWinstreakBestUhcDuel       float64 `json:"duels_winstreak_best_uhc_duel"`
				BridgeDuelLosses                float64 `json:"bridge_duel_losses"`
			} `json:"Duels"`
			Pit struct {
				Profile struct {
					OutgoingOffers  []interface{} `json:"outgoing_offers"`
					ContractChoices interface{}   `json:"contract_choices"`
					LastSave        float64       `json:"last_save"`
					KingQuest       struct {
						Kills  float64 `json:"kills"`
						Renown float64 `json:"renown"`
					} `json:"king_quest"`
					HatColor        float64       `json:"hat_color"`
					TradeTimestamps []interface{} `json:"trade_timestamps"`
					Unlocks1        []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks_1"`
					DeathRecaps struct {
						Type float64   `json:"type"`
						Data []float64 `json:"data"`
					} `json:"death_recaps"`
					InvEnderchest struct {
						Type float64   `json:"type"`
						Data []float64 `json:"data"`
					} `json:"inv_enderchest"`
					Unlocks2 []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks_2"`
					Unlocks3 []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks_3"`
					Unlocks4 []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks_4"`
					Cash     float64 `json:"cash"`
					Unlocks5 []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks_5"`
					Unlocks6 []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks_6"`
					Unlocks7 []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks_7"`
					LeaderboardStats struct {
						PitRaffleJackpot2019Summer float64 `json:"Pit_raffle_jackpot_2019_summer"`
						PitKotlTime2019Spring      float64 `json:"Pit_kotl_time_2019_spring"`
						PitRaffleTickets2019Summer float64 `json:"Pit_raffle_tickets_2019_summer"`
						PitKotlGold2019Spring      float64 `json:"Pit_kotl_gold_2019_spring"`
						PitRaffleTickets2019Spring float64 `json:"Pit_raffle_tickets_2019_spring"`
					} `json:"leaderboard_stats"`
					SelectedPerk3       string  `json:"selected_perk_3"`
					SelectedPerk2       string  `json:"selected_perk_2"`
					SelectedPerk1       string  `json:"selected_perk_1"`
					SelectedPerk0       string  `json:"selected_perk_0"`
					LastContract        float64 `json:"last_contract"`
					CashDuringPrestige3 float64 `json:"cash_during_prestige_3"`
					GoldTransactions    []struct {
						Amount    float64 `json:"amount"`
						Timestamp float64 `json:"timestamp"`
					} `json:"gold_transactions"`
					CashDuringPrestige2 float64 `json:"cash_during_prestige_2"`
					CashDuringPrestige5 float64 `json:"cash_during_prestige_5"`
					CashDuringPrestige4 float64 `json:"cash_during_prestige_4"`
					CashDuringPrestige7 float64 `json:"cash_during_prestige_7"`
					CashDuringPrestige6 float64 `json:"cash_during_prestige_6"`
					InvContents         struct {
						Type float64   `json:"type"`
						Data []float64 `json:"data"`
					} `json:"inv_contents"`
					Unlocks []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"unlocks"`
					CashDuringPrestige1 float64 `json:"cash_during_prestige_1"`
					CashDuringPrestige0 float64 `json:"cash_during_prestige_0"`
					Renown              float64 `json:"renown"`
					MovedAchievements1  bool    `json:"moved_achievements_1"`
					MovedAchievements2  bool    `json:"moved_achievements_2"`
					Prestiges           []struct {
						Index        float64 `json:"index"`
						XpOnPrestige float64 `json:"xp_on_prestige"`
						Timestamp    float64 `json:"timestamp"`
					} `json:"prestiges"`
					SpireStashInv struct {
						Type float64   `json:"type"`
						Data []float64 `json:"data"`
					} `json:"spire_stash_inv"`
					ZeroPointThreeGoldTransfer bool `json:"zero_point_three_gold_transfer"`
					RenownUnlocks              []struct {
						Tier        float64 `json:"tier"`
						AcquireDate float64 `json:"acquireDate"`
						Key         string  `json:"key"`
					} `json:"renown_unlocks"`
					SpireStashArmor struct {
						Type float64   `json:"type"`
						Data []float64 `json:"data"`
					} `json:"spire_stash_armor"`
					LastMidfightDisconnect float64 `json:"last_midfight_disconnect"`
					InvArmor               struct {
						Type float64   `json:"type"`
						Data []float64 `json:"data"`
					} `json:"inv_armor"`
					LoginMessages   []interface{} `json:"login_messages"`
					HotbarFavorites []float64     `json:"hotbar_favorites"`
					Xp              float64       `json:"xp"`
					EndedContracts  []struct {
						Difficulty   string  `json:"difficulty"`
						GoldReward   float64 `json:"gold_reward"`
						Requirements struct {
							Streak float64 `json:"streak"`
						} `json:"requirements"`
						Progress struct {
							Streak float64 `json:"streak"`
						} `json:"progress"`
						CompletionDate float64 `json:"completion_date"`
						RemainingTicks float64 `json:"remaining_ticks"`
						Key            string  `json:"key"`
					} `json:"ended_contracts"`
					Bounties []interface{} `json:"bounties"`
				} `json:"profile"`
				PitStatsPtl struct {
					Deaths                float64 `json:"deaths"`
					EnderchestOpened      float64 `json:"enderchest_opened"`
					MeleeDamageDealt      float64 `json:"melee_damage_dealt"`
					CashEarned            float64 `json:"cash_earned"`
					Joins                 float64 `json:"joins"`
					PlaytimeMinutes       float64 `json:"playtime_minutes"`
					BowDamageReceived     float64 `json:"bow_damage_received"`
					Kills                 float64 `json:"kills"`
					DamageReceived        float64 `json:"damage_received"`
					JumpedIntoPit         float64 `json:"jumped_into_pit"`
					MeleeDamageReceived   float64 `json:"melee_damage_received"`
					LeftClicks            float64 `json:"left_clicks"`
					DamageDealt           float64 `json:"damage_dealt"`
					Assists               float64 `json:"assists"`
					BlocksPlaced          float64 `json:"blocks_placed"`
					ContractsCompleted    float64 `json:"contracts_completed"`
					LaunchedByLaunchers   float64 `json:"launched_by_launchers"`
					ArrowsFired           float64 `json:"arrows_fired"`
					BowDamageDealt        float64 `json:"bow_damage_dealt"`
					ContractsStarted      float64 `json:"contracts_started"`
					ChatMessages          float64 `json:"chat_messages"`
					ArrowHits             float64 `json:"arrow_hits"`
					DiamondItemsPurchased float64 `json:"diamond_items_purchased"`
					MaxStreak             float64 `json:"max_streak"`
				} `json:"pit_stats_ptl"`
			} `json:"Pit"`
			Housing struct {
				LayoutItems struct {
				} `json:"layout_items"`
			} `json:"Housing"`
			SkyBlock struct {
				Profiles struct {
					B876Ec32E396476Ba1158438D83C67D4 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"b876ec32e396476ba1158438d83c67d4"`
					Eight92Dde8Fda9B494Aa4139040758E3Ae7 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"892dde8fda9b494aa4139040758e3ae7"`
					Six4722047F9B34E69B67B76A62351Eb05 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"64722047f9b34e69b67b76a62351eb05"`
					Three564B910B00C40698B6Cf652A344A41E struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"3564b910b00c40698b6cf652a344a41e"`
					NineD3C26E0078D47A9Be7F0292B710Bed3 struct {
						ProfileID string `json:"profile_id"`
						CuteName  string `json:"cute_name"`
					} `json:"9d3c26e0078d47a9be7f0292b710bed3"`
				} `json:"profiles"`
			} `json:"SkyBlock"`
		} `json:"stats"`
		TestPass         bool    `json:"testPass"`
		ThanksReceived   float64 `json:"thanksReceived"`
		ThanksSent       float64 `json:"thanksSent"`
		TimePlaying      float64 `json:"timePlaying"`
		TournamentTokens float64 `json:"tournamentTokens"`
		UUID             string  `json:"uuid"`
		VanityMeta       struct {
			Packages []string `json:"packages"`
		} `json:"vanityMeta"`
		Wardrobe string `json:"wardrobe"`
		Eugene   struct {
			DailyTwoKExp float64 `json:"dailyTwoKExp"`
		} `json:"eugene"`
		Voting struct {
			SecondaryMinestatus float64 `json:"secondary_minestatus"`
			TotalMinestatus     float64 `json:"total_minestatus"`
			Total               float64 `json:"total"`
			LastMinestatus      float64 `json:"last_minestatus"`
			LastVote            float64 `json:"last_vote"`
			TotalMcsorg         float64 `json:"total_mcsorg"`
			SecondaryMcsorg     float64 `json:"secondary_mcsorg"`
			LastMcsorg          float64 `json:"last_mcsorg"`
			SecondaryMcsl       float64 `json:"secondary_mcsl"`
			TotalMcsl           float64 `json:"total_mcsl"`
			LastMcsl            float64 `json:"last_mcsl"`
			TotalPmc            float64 `json:"total_pmc"`
			SecondaryPmc        float64 `json:"secondary_pmc"`
			LastPmc             float64 `json:"last_pmc"`
			SecondaryTopg       float64 `json:"secondary_topg"`
			TotalTopg           float64 `json:"total_topg"`
			LastTopg            float64 `json:"last_topg"`
			VotesToday          float64 `json:"votesToday"`
			TotalMcmp           float64 `json:"total_mcmp"`
			SecondaryMcmp       float64 `json:"secondary_mcmp"`
			LastMcmp            float64 `json:"last_mcmp"`
			SecondaryMcipl      float64 `json:"secondary_mcipl"`
			TotalMcipl          float64 `json:"total_mcipl"`
			LastMcipl           float64 `json:"last_mcipl"`
			SecondaryPact       float64 `json:"secondary_pact"`
			TotalPact           float64 `json:"total_pact"`
			LastPact            float64 `json:"last_pact"`
			TotalMcf            float64 `json:"total_mcf"`
			SecondaryMcf        float64 `json:"secondary_mcf"`
			LastMcf             float64 `json:"last_mcf"`
		} `json:"voting"`
		McVersionRp             string `json:"mcVersionRp"`
		MostRecentlyThankedUUID string `json:"mostRecentlyThankedUuid"`
		PetConsumables          struct {
			BakedPotato  float64 `json:"BAKED_POTATO"`
			Cookie       float64 `json:"COOKIE"`
			Feather      float64 `json:"FEATHER"`
			HayBlock     float64 `json:"HAY_BLOCK"`
			SlimeBall    float64 `json:"SLIME_BALL"`
			CookedBeef   float64 `json:"COOKED_BEEF"`
			RedRose      float64 `json:"RED_ROSE"`
			WaterBucket  float64 `json:"WATER_BUCKET"`
			Melon        float64 `json:"MELON"`
			Stick        float64 `json:"STICK"`
			WoodSword    float64 `json:"WOOD_SWORD"`
			MilkBucket   float64 `json:"MILK_BUCKET"`
			GoldRecord   float64 `json:"GOLD_RECORD"`
			Leash        float64 `json:"LEASH"`
			LavaBucket   float64 `json:"LAVA_BUCKET"`
			Bone         float64 `json:"BONE"`
			MagmaCream   float64 `json:"MAGMA_CREAM"`
			Wheat        float64 `json:"WHEAT"`
			MushroomSoup float64 `json:"MUSHROOM_SOUP"`
			Bread        float64 `json:"BREAD"`
			PumpkinPie   float64 `json:"PUMPKIN_PIE"`
			Apple        float64 `json:"APPLE"`
			CarrotItem   float64 `json:"CARROT_ITEM"`
			RawFish      float64 `json:"RAW_FISH"`
			Pork         float64 `json:"PORK"`
			Cake         float64 `json:"CAKE"`
			RottenFlesh  float64 `json:"ROTTEN_FLESH"`
		} `json:"petConsumables"`
		HousingMeta struct {
			TutorialStep       string   `json:"tutorialStep"`
			Packages           []string `json:"packages"`
			AllowedBlocks      []string `json:"allowedBlocks"`
			VisibilityDisabled bool     `json:"visibilityDisabled"`
			PlayerSettings     struct {
				Border     string `json:"BORDER"`
				YtRepulsor string `json:"YT_REPULSOR"`
				Tips       string `json:"TIPS"`
				Visibility string `json:"VISIBILITY"`
			} `json:"playerSettings"`
			FirstHouseJoinMs   float64  `json:"firstHouseJoinMs"`
			GivenCookies105000 []string `json:"given_cookies_105000"`
			GivenCookies105006 []string `json:"given_cookies_105006"`
			GivenCookies105016 []string `json:"given_cookies_105016"`
			GivenCookies105025 []string `json:"given_cookies_105025"`
			SelectedChannelsV3 []string `json:"selectedChannels_v3"`
			GivenCookies105036 []string `json:"given_cookies_105036"`
			GivenCookies105048 []string `json:"given_cookies_105048"`
			Playlist           string   `json:"playlist"`
		} `json:"housingMeta"`
		PetStats struct {
			Pig struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"PIG"`
			IronGolem struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"IRON_GOLEM"`
			Snowman struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"SNOWMAN"`
			HorseBrown struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_BROWN"`
			HorseBrownBaby struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_BROWN_BABY"`
			SlimeBig struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Experience float64 `json:"experience"`
				Exercise   struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
			} `json:"SLIME_BIG"`
			Wolf struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"WOLF"`
			Zombie struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"ZOMBIE"`
			Silverfish struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SILVERFISH"`
			SheepWhite struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_WHITE"`
			SheepGray struct {
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_GRAY"`
			SheepBrown struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_BROWN"`
			SheepSilver struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_SILVER"`
			SheepWhiteBaby struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_WHITE_BABY"`
			SheepGrayBaby struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_GRAY_BABY"`
			SheepBrownBaby struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_BROWN_BABY"`
			SheepSilverBaby struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_SILVER_BABY"`
			SheepOrange struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_ORANGE"`
			SheepMagenta struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_MAGENTA"`
			SheepLightBlue struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_LIGHT_BLUE"`
			SheepYellow struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_YELLOW"`
			SheepLime struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_LIME"`
			SheepCyan struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_CYAN"`
			SheepPurple struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_PURPLE"`
			SheepBlue struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_BLUE"`
			SheepGreen struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_GREEN"`
			SheepRed struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_RED"`
			HorseGrayBaby struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_GRAY_BABY"`
			HorseSkeleton struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_SKELETON"`
			SlimeTiny struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"SLIME_TINY"`
			HorseDarkBrown struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_DARK_BROWN"`
			HorseGrey struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_GREY"`
			HorseChestnut struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_CHESTNUT"`
			HorseCreamy struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_CREAMY"`
			HorseCreamyBaby struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_CREAMY_BABY"`
			HorseChestnutBaby struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_CHESTNUT_BABY"`
			WildOcelot struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"WILD_OCELOT"`
			WildOcelotBaby struct {
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"WILD_OCELOT_BABY"`
			HorseUndead struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_UNDEAD"`
			MooshroomBaby struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"MOOSHROOM_BABY"`
			SlimeSmall struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SLIME_SMALL"`
			SheepPinkBaby struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_PINK_BABY"`
			SheepBlackBaby struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_BLACK_BABY"`
			HorseBlack struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_BLACK"`
			HorseWhite struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"HORSE_WHITE"`
			SheepPink struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_PINK"`
			SheepOrangeBaby struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_ORANGE_BABY"`
			SheepMagentaBaby struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_MAGENTA_BABY"`
			SheepBlack struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_BLACK"`
			SheepLightBlueBaby struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SHEEP_LIGHT_BLUE_BABY"`
			PigZombie struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"PIG_ZOMBIE"`
			CreeperPowered struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"CREEPER_POWERED"`
			Witch struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"WITCH"`
			MagmaCubeBig struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"MAGMA_CUBE_BIG"`
			Herobrine struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"HEROBRINE"`
			Creeper struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"CREEPER"`
			Spider struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SPIDER"`
			Enderman struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"ENDERMAN"`
			PigZombieBaby struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"PIG_ZOMBIE_BABY"`
			PigBaby struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"PIG_BABY"`
			Gorrila struct {
				Name string `json:"name"`
			} `json:"GORRILA"`
			Gifterino struct {
			} `json:"GIFTERINO"`
			Bat struct {
				Name   string `json:"name"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"BAT"`
			Elephant struct {
				Name string `json:"name"`
			} `json:"ELEPHANT"`
			Penguin struct {
				Name string `json:"name"`
			} `json:"PENGUIN"`
			CatBlack struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Experience float64 `json:"experience"`
			} `json:"CAT_BLACK"`
			CatRed struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"CAT_RED"`
			CatSiamese struct {
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Experience float64 `json:"experience"`
			} `json:"CAT_SIAMESE"`
			CatBlackBaby struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Experience float64 `json:"experience"`
			} `json:"CAT_BLACK_BABY"`
			Chicken struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"CHICKEN"`
			ChickenBaby struct {
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"CHICKEN_BABY"`
			CatRedBaby struct {
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"CAT_RED_BABY"`
			CatSiameseBaby struct {
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"CAT_SIAMESE_BABY"`
			WolfBaby struct {
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"WOLF_BABY"`
			ZombieBaby struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"ZOMBIE_BABY"`
			CaveSpider struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"CAVE_SPIDER"`
			BlackRabbit struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"BLACK_RABBIT"`
			BrownRabbit struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"BROWN_RABBIT"`
			GoldRabbit struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"GOLD_RABBIT"`
			SaltPepperRabbit struct {
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"SALT_PEPPER_RABBIT"`
			WhiteRabbit struct {
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"WHITE_RABBIT"`
			VillagerBlacksmith struct {
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"VILLAGER_BLACKSMITH"`
			VillagerBlacksmithBaby struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"VILLAGER_BLACKSMITH_BABY"`
			VillagerButcher struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"VILLAGER_BUTCHER"`
			RedHelper struct {
				Hunger struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Experience float64 `json:"experience"`
			} `json:"RED_HELPER"`
			GreenHelper struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
			} `json:"GREEN_HELPER"`
			Blaze struct {
				Thirst struct {
					Value     float64 `json:"value"`
					Timestamp float64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp float64 `json:"timestamp"`
					Value     float64 `json:"value"`
				} `json:"HUNGER"`
				Experience float64 `json:"experience"`
				Name       string  `json:"name"`
			} `json:"BLAZE"`
			LittleYou struct {
				Name string `json:"name"`
			} `json:"LITTLE_YOU"`
		} `json:"petStats"`
		PetJourneyTimestamp float64 `json:"petJourneyTimestamp"`
		Transformation      string  `json:"transformation"`
		SocialMedia         struct {
			Youtube string `json:"YOUTUBE"`
			Twitter string `json:"TWITTER"`
			Links   struct {
				Youtube string `json:"YOUTUBE"`
				Twitter string `json:"TWITTER"`
			} `json:"links"`
		} `json:"socialMedia"`
		DisableTipMessages bool `json:"disableTipMessages"`
		RewardConsumed     bool `json:"rewardConsumed"`
		GiftingMeta        struct {
			RealBundlesReceivedInc float64 `json:"realBundlesReceivedInc"`
			BundlesReceived        float64 `json:"bundlesReceived"`
			RealBundlesReceived    float64 `json:"realBundlesReceived"`
			GiftsGiven             float64 `json:"giftsGiven"`
			BundlesGiven           float64 `json:"bundlesGiven"`
			RealBundlesGiven       float64 `json:"realBundlesGiven"`
		} `json:"giftingMeta"`
		LevelingReward0         bool    `json:"levelingReward_0"`
		LevelingReward1         bool    `json:"levelingReward_1"`
		LevelingReward2         bool    `json:"levelingReward_2"`
		LevelingReward3         bool    `json:"levelingReward_3"`
		LevelingReward4         bool    `json:"levelingReward_4"`
		LevelingReward5         bool    `json:"levelingReward_5"`
		LevelingReward6         bool    `json:"levelingReward_6"`
		LevelingReward7         bool    `json:"levelingReward_7"`
		LevelingReward8         bool    `json:"levelingReward_8"`
		LevelingReward9         bool    `json:"levelingReward_9"`
		LevelingReward10        bool    `json:"levelingReward_10"`
		LevelingReward11        bool    `json:"levelingReward_11"`
		LevelingReward12        bool    `json:"levelingReward_12"`
		LevelingReward13        bool    `json:"levelingReward_13"`
		LevelingReward14        bool    `json:"levelingReward_14"`
		LevelingReward15        bool    `json:"levelingReward_15"`
		LevelingReward16        bool    `json:"levelingReward_16"`
		LevelingReward17        bool    `json:"levelingReward_17"`
		LevelingReward18        bool    `json:"levelingReward_18"`
		LevelingReward19        bool    `json:"levelingReward_19"`
		LevelingReward20        bool    `json:"levelingReward_20"`
		LevelingReward21        bool    `json:"levelingReward_21"`
		LevelingReward22        bool    `json:"levelingReward_22"`
		LevelingReward23        bool    `json:"levelingReward_23"`
		LevelingReward24        bool    `json:"levelingReward_24"`
		LevelingReward25        bool    `json:"levelingReward_25"`
		LevelingReward26        bool    `json:"levelingReward_26"`
		LevelingReward27        bool    `json:"levelingReward_27"`
		LevelingReward28        bool    `json:"levelingReward_28"`
		LevelingReward29        bool    `json:"levelingReward_29"`
		LevelingReward30        bool    `json:"levelingReward_30"`
		LevelingReward31        bool    `json:"levelingReward_31"`
		LevelingReward32        bool    `json:"levelingReward_32"`
		LevelingReward33        bool    `json:"levelingReward_33"`
		LevelingReward34        bool    `json:"levelingReward_34"`
		LevelingReward35        bool    `json:"levelingReward_35"`
		LevelingReward36        bool    `json:"levelingReward_36"`
		LevelingReward37        bool    `json:"levelingReward_37"`
		LevelingReward38        bool    `json:"levelingReward_38"`
		LevelingReward39        bool    `json:"levelingReward_39"`
		LevelingReward40        bool    `json:"levelingReward_40"`
		LevelingReward41        bool    `json:"levelingReward_41"`
		LevelingReward42        bool    `json:"levelingReward_42"`
		LevelingReward43        bool    `json:"levelingReward_43"`
		LevelingReward44        bool    `json:"levelingReward_44"`
		LevelingReward45        bool    `json:"levelingReward_45"`
		LevelingReward46        bool    `json:"levelingReward_46"`
		LevelingReward47        bool    `json:"levelingReward_47"`
		LevelingReward48        bool    `json:"levelingReward_48"`
		LevelingReward49        bool    `json:"levelingReward_49"`
		LevelingReward50        bool    `json:"levelingReward_50"`
		LevelingReward51        bool    `json:"levelingReward_51"`
		LevelingReward52        bool    `json:"levelingReward_52"`
		LevelingReward53        bool    `json:"levelingReward_53"`
		LevelingReward54        bool    `json:"levelingReward_54"`
		LevelingReward55        bool    `json:"levelingReward_55"`
		LevelingReward56        bool    `json:"levelingReward_56"`
		LevelingReward57        bool    `json:"levelingReward_57"`
		LevelingReward58        bool    `json:"levelingReward_58"`
		LevelingReward59        bool    `json:"levelingReward_59"`
		LevelingReward60        bool    `json:"levelingReward_60"`
		LevelingReward61        bool    `json:"levelingReward_61"`
		LevelingReward62        bool    `json:"levelingReward_62"`
		LevelingReward63        bool    `json:"levelingReward_63"`
		LevelingReward64        bool    `json:"levelingReward_64"`
		LevelingReward65        bool    `json:"levelingReward_65"`
		LevelingReward66        bool    `json:"levelingReward_66"`
		LevelingReward67        bool    `json:"levelingReward_67"`
		LevelingReward68        bool    `json:"levelingReward_68"`
		LevelingReward69        bool    `json:"levelingReward_69"`
		LevelingReward70        bool    `json:"levelingReward_70"`
		LevelingReward71        bool    `json:"levelingReward_71"`
		LevelingReward72        bool    `json:"levelingReward_72"`
		LevelingReward73        bool    `json:"levelingReward_73"`
		LevelingReward74        bool    `json:"levelingReward_74"`
		LevelingReward75        bool    `json:"levelingReward_75"`
		LevelingReward76        bool    `json:"levelingReward_76"`
		LevelingReward77        bool    `json:"levelingReward_77"`
		LevelingReward78        bool    `json:"levelingReward_78"`
		LevelingReward79        bool    `json:"levelingReward_79"`
		LevelingReward80        bool    `json:"levelingReward_80"`
		LevelingReward81        bool    `json:"levelingReward_81"`
		LevelingReward82        bool    `json:"levelingReward_82"`
		LevelingReward83        bool    `json:"levelingReward_83"`
		LevelingReward84        bool    `json:"levelingReward_84"`
		LevelingReward85        bool    `json:"levelingReward_85"`
		LevelingReward86        bool    `json:"levelingReward_86"`
		LevelingReward87        bool    `json:"levelingReward_87"`
		LevelingReward88        bool    `json:"levelingReward_88"`
		LevelingReward89        bool    `json:"levelingReward_89"`
		LevelingReward90        bool    `json:"levelingReward_90"`
		LevelingReward91        bool    `json:"levelingReward_91"`
		LevelingReward92        bool    `json:"levelingReward_92"`
		LevelingReward93        bool    `json:"levelingReward_93"`
		LevelingReward94        bool    `json:"levelingReward_94"`
		LevelingReward95        bool    `json:"levelingReward_95"`
		LevelingReward96        bool    `json:"levelingReward_96"`
		LevelingReward97        bool    `json:"levelingReward_97"`
		LevelingReward98        bool    `json:"levelingReward_98"`
		LevelingReward99        bool    `json:"levelingReward_99"`
		LevelingReward109       bool    `json:"levelingReward_109"`
		LevelingReward119       bool    `json:"levelingReward_119"`
		LevelingReward129       bool    `json:"levelingReward_129"`
		LevelingReward139       bool    `json:"levelingReward_139"`
		LevelingReward149       bool    `json:"levelingReward_149"`
		VanityConvertedBoxToday float64 `json:"vanityConvertedBoxToday"`
		VanityFirstConvertedBox float64 `json:"vanityFirstConvertedBox"`
		SpecialtyCooldowns      struct {
			Mvp0     bool `json:"MVP0"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			Vip0     bool `json:"VIP0"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			Vip1     bool `json:"VIP1"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			Mvp1     bool `json:"MVP1"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			Vip2     bool `json:"VIP2"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			Mvp2     bool `json:"MVP2"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
			Vip3     bool `json:"VIP3"`
			VipPlus3 bool `json:"VIP_PLUS3"`
			Mvp3     bool `json:"MVP3"`
			MvpPlus3 bool `json:"MVP_PLUS3"`
			Mvp4     bool `json:"MVP4"`
			Vip4     bool `json:"VIP4"`
			VipPlus4 bool `json:"VIP_PLUS4"`
			MvpPlus4 bool `json:"MVP_PLUS4"`
			Vip5     bool `json:"VIP5"`
			VipPlus5 bool `json:"VIP_PLUS5"`
			Mvp5     bool `json:"MVP5"`
			MvpPlus5 bool `json:"MVP_PLUS5"`
			Vip6     bool `json:"VIP6"`
			VipPlus6 bool `json:"VIP_PLUS6"`
			Mvp6     bool `json:"MVP6"`
			MvpPlus6 bool `json:"MVP_PLUS6"`
		} `json:"specialtyCooldowns"`
		QuestSettings struct {
			AutoActivate bool `json:"autoActivate"`
		} `json:"questSettings"`
		LastAdsenseGenerateTime float64 `json:"lastAdsenseGenerateTime"`
		LastClaimedReward       float64 `json:"lastClaimedReward"`
		TotalRewards            float64 `json:"totalRewards"`
		TotalDailyRewards       float64 `json:"totalDailyRewards"`
		RewardStreak            float64 `json:"rewardStreak"`
		RewardScore             float64 `json:"rewardScore"`
		RewardHighScore         float64 `json:"rewardHighScore"`
		RankPlusColor           string  `json:"rankPlusColor"`
		AdsenseTokens           float64 `json:"adsense_tokens"`
		LevelingReward159       bool    `json:"levelingReward_159"`
		FlashingSalePopup       float64 `json:"flashingSalePopup"`
		FlashingSalePoppedUp    float64 `json:"flashingSalePoppedUp"`
		FlashingSaleOpens       float64 `json:"flashingSaleOpens"`
		FlashingSaleClicks      float64 `json:"flashingSaleClicks"`
		Halloween2016Cooldowns  struct {
			Vip0     bool `json:"VIP0"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			Mvp0     bool `json:"MVP0"`
			Vip1     bool `json:"VIP1"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			Mvp1     bool `json:"MVP1"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			Vip2     bool `json:"VIP2"`
			Mvp2     bool `json:"MVP2"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
		} `json:"halloween2016Cooldowns"`
		LevelingReward169    bool `json:"levelingReward_169"`
		SantaQuestStarted    bool `json:"SANTA_QUEST_STARTED"`
		Holiday2016Cooldowns struct {
			Vip0     bool `json:"VIP0"`
			Vip1     bool `json:"VIP1"`
			Vip2     bool `json:"VIP2"`
			Vip3     bool `json:"VIP3"`
			Vip4     bool `json:"VIP4"`
			Mvp0     bool `json:"MVP0"`
			Mvp1     bool `json:"MVP1"`
			Mvp2     bool `json:"MVP2"`
			Mvp3     bool `json:"MVP3"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			VipPlus3 bool `json:"VIP_PLUS3"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
			MvpPlus3 bool `json:"MVP_PLUS3"`
			VipPlus4 bool `json:"VIP_PLUS4"`
			MvpPlus4 bool `json:"MVP_PLUS4"`
			Mvp4     bool `json:"MVP4"`
			Vip5     bool `json:"VIP5"`
			VipPlus5 bool `json:"VIP_PLUS5"`
			Mvp5     bool `json:"MVP5"`
			MvpPlus5 bool `json:"MVP_PLUS5"`
		} `json:"holiday2016Cooldowns"`
		LevelingReward179 bool `json:"levelingReward_179"`
		CompassStats      struct {
			Compass struct {
				Battleground float64 `json:"battleground"`
				Arcade       float64 `json:"arcade"`
				Prototype    float64 `json:"prototype"`
				Skywars      float64 `json:"skywars"`
				Tntgames     float64 `json:"tntgames"`
			} `json:"compass"`
		} `json:"compassStats"`
		FortuneBuff           float64  `json:"fortuneBuff"`
		NetworkUpdateBook     string   `json:"network_update_book"`
		LastLogout            float64  `json:"lastLogout"`
		AchievementTracking   []string `json:"achievementTracking"`
		AchievementRewardsNew struct {
			ForPoints200   float64 `json:"for_points_200"`
			ForPoints300   float64 `json:"for_points_300"`
			ForPoints400   float64 `json:"for_points_400"`
			ForPoints500   float64 `json:"for_points_500"`
			ForPoints600   float64 `json:"for_points_600"`
			ForPoints700   float64 `json:"for_points_700"`
			ForPoints800   float64 `json:"for_points_800"`
			ForPoints900   float64 `json:"for_points_900"`
			ForPoints1000  float64 `json:"for_points_1000"`
			ForPoints1100  float64 `json:"for_points_1100"`
			ForPoints1200  float64 `json:"for_points_1200"`
			ForPoints1300  float64 `json:"for_points_1300"`
			ForPoints1400  float64 `json:"for_points_1400"`
			ForPoints1500  float64 `json:"for_points_1500"`
			ForPoints1600  float64 `json:"for_points_1600"`
			ForPoints1700  float64 `json:"for_points_1700"`
			ForPoints1800  float64 `json:"for_points_1800"`
			ForPoints1900  float64 `json:"for_points_1900"`
			ForPoints2000  float64 `json:"for_points_2000"`
			ForPoints2100  float64 `json:"for_points_2100"`
			ForPoints2200  float64 `json:"for_points_2200"`
			ForPoints2300  float64 `json:"for_points_2300"`
			ForPoints2400  float64 `json:"for_points_2400"`
			ForPoints2500  float64 `json:"for_points_2500"`
			ForPoints2600  float64 `json:"for_points_2600"`
			ForPoints2700  float64 `json:"for_points_2700"`
			ForPoints2800  float64 `json:"for_points_2800"`
			ForPoints2900  float64 `json:"for_points_2900"`
			ForPoints3000  float64 `json:"for_points_3000"`
			ForPoints3100  float64 `json:"for_points_3100"`
			ForPoints3200  float64 `json:"for_points_3200"`
			ForPoints3300  float64 `json:"for_points_3300"`
			ForPoints3400  float64 `json:"for_points_3400"`
			ForPoints3500  float64 `json:"for_points_3500"`
			ForPoints3600  float64 `json:"for_points_3600"`
			ForPoints3700  float64 `json:"for_points_3700"`
			ForPoints3800  float64 `json:"for_points_3800"`
			ForPoints3900  float64 `json:"for_points_3900"`
			ForPoints4000  float64 `json:"for_points_4000"`
			ForPoints4100  float64 `json:"for_points_4100"`
			ForPoints4200  float64 `json:"for_points_4200"`
			ForPoints4300  float64 `json:"for_points_4300"`
			ForPoints4400  float64 `json:"for_points_4400"`
			ForPoints4500  float64 `json:"for_points_4500"`
			ForPoints4600  float64 `json:"for_points_4600"`
			ForPoints4700  float64 `json:"for_points_4700"`
			ForPoints4800  float64 `json:"for_points_4800"`
			ForPoints4900  float64 `json:"for_points_4900"`
			ForPoints5000  float64 `json:"for_points_5000"`
			ForPoints5100  float64 `json:"for_points_5100"`
			ForPoints5200  float64 `json:"for_points_5200"`
			ForPoints5300  float64 `json:"for_points_5300"`
			ForPoints5400  float64 `json:"for_points_5400"`
			ForPoints5500  float64 `json:"for_points_5500"`
			ForPoints5600  float64 `json:"for_points_5600"`
			ForPoints5700  float64 `json:"for_points_5700"`
			ForPoints5800  float64 `json:"for_points_5800"`
			ForPoints5900  float64 `json:"for_points_5900"`
			ForPoints6000  float64 `json:"for_points_6000"`
			ForPoints6100  float64 `json:"for_points_6100"`
			ForPoints6200  float64 `json:"for_points_6200"`
			ForPoints6300  float64 `json:"for_points_6300"`
			ForPoints6400  float64 `json:"for_points_6400"`
			ForPoints6500  float64 `json:"for_points_6500"`
			ForPoints6600  float64 `json:"for_points_6600"`
			ForPoints6700  float64 `json:"for_points_6700"`
			ForPoints6800  float64 `json:"for_points_6800"`
			ForPoints6900  float64 `json:"for_points_6900"`
			ForPoints7000  float64 `json:"for_points_7000"`
			ForPoints7100  float64 `json:"for_points_7100"`
			ForPoints7200  float64 `json:"for_points_7200"`
			ForPoints7300  float64 `json:"for_points_7300"`
			ForPoints7400  float64 `json:"for_points_7400"`
			ForPoints7500  float64 `json:"for_points_7500"`
			ForPoints7600  float64 `json:"for_points_7600"`
			ForPoints7700  float64 `json:"for_points_7700"`
			ForPoints7800  float64 `json:"for_points_7800"`
			ForPoints7900  float64 `json:"for_points_7900"`
			ForPoints8000  float64 `json:"for_points_8000"`
			ForPoints8100  float64 `json:"for_points_8100"`
			ForPoints8200  float64 `json:"for_points_8200"`
			ForPoints8300  float64 `json:"for_points_8300"`
			ForPoints8400  float64 `json:"for_points_8400"`
			ForPoints8500  float64 `json:"for_points_8500"`
			ForPoints8600  float64 `json:"for_points_8600"`
			ForPoints8700  float64 `json:"for_points_8700"`
			ForPoints8800  float64 `json:"for_points_8800"`
			ForPoints9000  float64 `json:"for_points_9000"`
			ForPoints8900  float64 `json:"for_points_8900"`
			ForPoints9100  float64 `json:"for_points_9100"`
			ForPoints9200  float64 `json:"for_points_9200"`
			ForPoints9300  float64 `json:"for_points_9300"`
			ForPoints9400  float64 `json:"for_points_9400"`
			ForPoints9500  float64 `json:"for_points_9500"`
			ForPoints9600  float64 `json:"for_points_9600"`
			ForPoints9700  float64 `json:"for_points_9700"`
			ForPoints9800  float64 `json:"for_points_9800"`
			ForPoints9900  float64 `json:"for_points_9900"`
			ForPoints10000 float64 `json:"for_points_10000"`
			ForPoints10100 float64 `json:"for_points_10100"`
			ForPoints10200 float64 `json:"for_points_10200"`
			ForPoints10300 float64 `json:"for_points_10300"`
			ForPoints10400 float64 `json:"for_points_10400"`
			ForPoints10500 float64 `json:"for_points_10500"`
			ForPoints10600 float64 `json:"for_points_10600"`
			ForPoints10700 float64 `json:"for_points_10700"`
			ForPoints10800 float64 `json:"for_points_10800"`
			ForPoints10900 float64 `json:"for_points_10900"`
			ForPoints11000 float64 `json:"for_points_11000"`
			ForPoints11100 float64 `json:"for_points_11100"`
			ForPoints11200 float64 `json:"for_points_11200"`
			ForPoints11300 float64 `json:"for_points_11300"`
			ForPoints11400 float64 `json:"for_points_11400"`
			ForPoints11500 float64 `json:"for_points_11500"`
			ForPoints11600 float64 `json:"for_points_11600"`
			ForPoints11700 float64 `json:"for_points_11700"`
			ForPoints11800 float64 `json:"for_points_11800"`
			ForPoints11900 float64 `json:"for_points_11900"`
			ForPoints12000 float64 `json:"for_points_12000"`
			ForPoints12100 float64 `json:"for_points_12100"`
			ForPoints12200 float64 `json:"for_points_12200"`
			ForPoints12300 float64 `json:"for_points_12300"`
		} `json:"achievementRewardsNew"`
		AchievementTotem struct {
			CanCustomize     bool     `json:"canCustomize"`
			AllowedMaxHeight float64  `json:"allowed_max_height"`
			UnlockedParts    []string `json:"unlockedParts"`
			SelectedParts    struct {
				Slot0 string `json:"slot_0"`
				Slot1 string `json:"slot_1"`
				Slot2 string `json:"slot_2"`
				Slot3 string `json:"slot_3"`
			} `json:"selectedParts"`
			UnlockedColors []string `json:"unlockedColors"`
			SelectedColors struct {
				Slotcolor0 string `json:"slotcolor_0"`
			} `json:"selectedColors"`
		} `json:"achievementTotem"`
		LevelingReward189            bool    `json:"levelingReward_189"`
		LevelingReward199            bool    `json:"levelingReward_199"`
		LevelingReward209            bool    `json:"levelingReward_209"`
		LevelingReward219            bool    `json:"levelingReward_219"`
		LevelingReward224            bool    `json:"levelingReward_224"`
		CompletedChristmasQuests2017 float64 `json:"completed_christmas_quests_2017"`
		Challenges                   struct {
			AllTime struct {
				BEDWARSOffensive                 float64 `json:"BEDWARS__offensive"`
				BEDWARSSupport                   float64 `json:"BEDWARS__support"`
				SKYWARSRushChallenge             float64 `json:"SKYWARS__rush_challenge"`
				SKYWARSEndermanChallenge         float64 `json:"SKYWARS__enderman_challenge"`
				SKYWARSFeedingTheVoidChallenge   float64 `json:"SKYWARS__feeding_the_void_challenge"`
				SUPERSMASHLeaderboardChallenge   float64 `json:"SUPER_SMASH__leaderboard_challenge"`
				SUPERSMASHSmashChallenge         float64 `json:"SUPER_SMASH__smash_challenge"`
				SUPERSMASHFlawlessChallenge      float64 `json:"SUPER_SMASH__flawless_challenge"`
				ARCADEPartyGamesChallenge        float64 `json:"ARCADE__party_games_challenge"`
				BUILDBATTLEGuesserChallenge      float64 `json:"BUILD_BATTLE__guesser_challenge"`
				BUILDBATTLETop3Challenge         float64 `json:"BUILD_BATTLE__top_3_challenge"`
				QUAKECRAFTDonTBlinkChallenge     float64 `json:"QUAKECRAFT__don\u0027t_blink_challenge"`
				QUAKECRAFTComboChallenge         float64 `json:"QUAKECRAFT__combo_challenge"`
				QUAKECRAFTKillingStreakChallenge float64 `json:"QUAKECRAFT__killing_streak_challenge"`
				PAINTBALLKillStreakChallenge     float64 `json:"PAINTBALL__kill_streak_challenge"`
				PAINTBALLKillingSpreeChallenge   float64 `json:"PAINTBALL__killing_spree_challenge"`
				PAINTBALLFinishChallenge         float64 `json:"PAINTBALL__finish_challenge"`
				GINGERBREADFirstPlaceChallenge   float64 `json:"GINGERBREAD__first_place_challenge"`
				GINGERBREADCoinChallenge         float64 `json:"GINGERBREAD__coin_challenge"`
				MURDERMYSTERYMurderSpree         float64 `json:"MURDER_MYSTERY__murder_spree"`
				MURDERMYSTERYHero                float64 `json:"MURDER_MYSTERY__hero"`
				MCGOKillingSpreeChallenge        float64 `json:"MCGO__killing_spree_challenge"`
				PAINTBALLNukeChallenge           float64 `json:"PAINTBALL__nuke_challenge"`
				TNTGAMESTntWizardChallenge       float64 `json:"TNTGAMES__tnt_wizard_challenge"`
				GINGERBREADBananaChallenge       float64 `json:"GINGERBREAD__banana_challenge"`
				DUELSFeedTheVoidChallenge        float64 `json:"DUELS__feed_the_void_challenge"`
				ARCADEHoleInTheWallChallenge     float64 `json:"ARCADE__hole_in_the_wall_challenge"`
				ARCADEFootballChallenge          float64 `json:"ARCADE__football_challenge"`
				GINGERBREADLeaderboardChallenge  float64 `json:"GINGERBREAD__leaderboard_challenge"`
				ARCADEHypixelSaysChallenge       float64 `json:"ARCADE__hypixel_says_challenge"`
				DUELSTeamsChallenge              float64 `json:"DUELS__teams_challenge"`
				ARCADEZombiesChallenge           float64 `json:"ARCADE__zombies_challenge"`
				ARCADEPixelPaintersChallenge     float64 `json:"ARCADE__pixel_painters_challenge"`
				SKYCLASHEnderchestChallenge      float64 `json:"SKYCLASH__enderchest_challenge"`
				SKYCLASHFighterChallenge         float64 `json:"SKYCLASH__fighter_challenge"`
				SKYCLASHMonsterKillerChallenge   float64 `json:"SKYCLASH__monster_killer_challenge"`
				BATTLEGROUNDBruteChallenge       float64 `json:"BATTLEGROUND__brute_challenge"`
				BATTLEGROUNDCarryChallenge       float64 `json:"BATTLEGROUND__carry_challenge"`
				BATTLEGROUNDSupportChallenge     float64 `json:"BATTLEGROUND__support_challenge"`
				BATTLEGROUNDCaptureChallenge     float64 `json:"BATTLEGROUND__capture_challenge"`
				TNTGAMESTntRunChallenge          float64 `json:"TNTGAMES__tnt_run_challenge"`
				TNTGAMESBowSpleefChallenge       float64 `json:"TNTGAMES__bow_spleef_challenge"`
				DUELSTargetPracticeChallenge     float64 `json:"DUELS__target_practice_challenge"`
				TNTGAMESPvpRunChallenge          float64 `json:"TNTGAMES__pvp_run_challenge"`
				TNTGAMESTntWizardsChallenge      float64 `json:"TNTGAMES__tnt_wizards_challenge"`
				TNTGAMESTntTagChallenge          float64 `json:"TNTGAMES__tnt_tag_challenge"`
				VAMPIREZLastStandChallenge       float64 `json:"VAMPIREZ__last_stand_challenge"`
				VAMPIREZFangChallenge            float64 `json:"VAMPIREZ__fang_challenge"`
				ARCADEThrowOutChallenge          float64 `json:"ARCADE__throw_out_challenge"`
				VAMPIREZPurifyingChallenge       float64 `json:"VAMPIREZ__purifying_challenge"`
			} `json:"all_time"`
		} `json:"challenges"`
		AchievementSync struct {
			QuakeTiered float64 `json:"quake_tiered"`
		} `json:"achievementSync"`
		LevelingReward249 bool    `json:"levelingReward_249"`
		AchievementPoints float64 `json:"achievementPoints"`
		Tourney           struct {
			FirstJoinLobby float64 `json:"first_join_lobby"`
			SwCrazySolo0   struct {
				Playtime             float64 `json:"playtime"`
				TributesEarned       float64 `json:"tributes_earned"`
				FirstWin             float64 `json:"first_win"`
				ClaimedRankingReward float64 `json:"claimed_ranking_reward"`
			} `json:"sw_crazy_solo_0"`
			TotalTributes float64 `json:"total_tributes"`
			QuakeSolo20   struct {
				GamesPlayed          float64 `json:"games_played"`
				Playtime             float64 `json:"playtime"`
				TributesEarned       float64 `json:"tributes_earned"`
				FirstGame            float64 `json:"first_game"`
				ClaimedRankingReward float64 `json:"claimed_ranking_reward"`
			} `json:"quake_solo2_0"`
			GingerbreadSolo0 struct {
				SeenRPbook           bool    `json:"seenRPbook"`
				GamesPlayed          float64 `json:"games_played"`
				Playtime             float64 `json:"playtime"`
				TributesEarned       float64 `json:"tributes_earned"`
				FirstWin             float64 `json:"first_win"`
				ClaimedRankingReward float64 `json:"claimed_ranking_reward"`
			} `json:"gingerbread_solo_0"`
		} `json:"tourney"`
		CompletedChristmasQuests2018 float64 `json:"completed_christmas_quests_2018"`
		VanityFavorites              string  `json:"vanityFavorites"`
		CurrentClickEffect           string  `json:"currentClickEffect"`
		ParkourCheckpointBests       struct {
			Warlords struct {
				Num0 float64 `json:"0"`
				Num1 float64 `json:"1"`
				Num2 float64 `json:"2"`
			} `json:"Warlords"`
			BlitzLobby struct {
				Num0 float64 `json:"0"`
			} `json:"BlitzLobby"`
			SpeedUHC struct {
				Num0 float64 `json:"0"`
				Num1 float64 `json:"1"`
				Num2 float64 `json:"2"`
			} `json:"SpeedUHC"`
			Tnt struct {
				Num0 float64 `json:"0"`
				Num1 float64 `json:"1"`
				Num2 float64 `json:"2"`
				Num3 float64 `json:"3"`
			} `json:"TNT"`
			SkywarsAug2017 struct {
				Num0 float64 `json:"0"`
				Num1 float64 `json:"1"`
				Num2 float64 `json:"2"`
				Num3 float64 `json:"3"`
			} `json:"SkywarsAug2017"`
			Bedwars struct {
				Num0 float64 `json:"0"`
				Num1 float64 `json:"1"`
				Num2 float64 `json:"2"`
				Num3 float64 `json:"3"`
			} `json:"Bedwars"`
		} `json:"ParkourCheckpointBests"`
		BattlePassGlowStatus bool   `json:"battlePassGlowStatus"`
		Prefix               string `json:"prefix"`
		Dmcrates102019       struct {
			Regular bool `json:"REGULAR"`
			Vip     bool `json:"VIP"`
			VipPlus bool `json:"VIP_PLUS"`
			Mvp     bool `json:"MVP"`
			MvpPlus bool `json:"MVP_PLUS"`
		} `json:"dmcrates-10-2019"`
		Monthlycrates struct {
			One2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				MvpPlus bool `json:"MVP_PLUS"`
				Mvp     bool `json:"MVP"`
			} `json:"1-2017"`
			One2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"1-2018"`
			One2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"1-2019"`
			One02016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2016"`
			One02017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2017"`
			One02018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				Mvp     bool `json:"MVP"`
				VipPlus bool `json:"VIP_PLUS"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2018"`
			One02019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"10-2019"`
			One12016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"11-2016"`
			One12017 struct {
				VipPlus bool `json:"VIP_PLUS"`
				Vip     bool `json:"VIP"`
				Regular bool `json:"REGULAR"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"11-2017"`
			One22016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2016"`
			One22017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2017"`
			One22018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2018"`
			Two2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2017"`
			Two2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2018"`
			Two2019 struct {
				Regular bool `json:"REGULAR"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
				Vip     bool `json:"VIP"`
			} `json:"2-2019"`
			Three2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2017"`
			Three2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2018"`
			Three2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2019"`
			Four2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2017"`
			Four2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2018"`
			Four2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2019"`
			Five2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2017"`
			Five2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2018"`
			Five2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2019"`
			Six2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2017"`
			Six2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2019"`
			Seven2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2016"`
			Seven2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2017"`
			Seven2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2019"`
			Eight2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"8-2016"`
			Eight2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"8-2017"`
			Eight2018 struct {
				Regular bool `json:"REGULAR"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
				Vip     bool `json:"VIP"`
			} `json:"8-2018"`
			Eight2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"8-2019"`
			Nine2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2016"`
			Nine2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2017"`
			Nine2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2019"`
			One22019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2019"`
			One2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"1-2020"`
			Two2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2020"`
			Three2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2020"`
			Five2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2020"`
			Six2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2020"`
			Seven2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2020"`
			One12020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"11-2020"`
			Three2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2021"`
			Four2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"4-2021"`
			Five2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2021"`
			Six2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"6-2021"`
			Nine2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"9-2021"`
		} `json:"monthlycrates"`
		Halloween2019Cooldowns struct {
			VipPlus3   bool `json:"VIP_PLUS3"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			Normal0    bool `json:"NORMAL0"`
			Normal1    bool `json:"NORMAL1"`
			Normal2    bool `json:"NORMAL2"`
			Normal3    bool `json:"NORMAL3"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			MvpPlus1   bool `json:"MVP_PLUS1"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			MvpPlus3   bool `json:"MVP_PLUS3"`
			Superstar0 bool `json:"SUPERSTAR0"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar3 bool `json:"SUPERSTAR3"`
			Vip3       bool `json:"VIP3"`
			Vip2       bool `json:"VIP2"`
			Vip1       bool `json:"VIP1"`
			Mvp3       bool `json:"MVP3"`
			Mvp2       bool `json:"MVP2"`
			Mvp1       bool `json:"MVP1"`
			Mvp0       bool `json:"MVP0"`
			Vip0       bool `json:"VIP0"`
		} `json:"halloween2019Cooldowns"`
		AdventRewards2017 struct {
			Day9  float64 `json:"day9"`
			Day10 float64 `json:"day10"`
			Day11 float64 `json:"day11"`
			Day12 float64 `json:"day12"`
			Day14 float64 `json:"day14"`
			Day15 float64 `json:"day15"`
			Day16 float64 `json:"day16"`
			Day17 float64 `json:"day17"`
			Day18 float64 `json:"day18"`
			Day19 float64 `json:"day19"`
			Day20 float64 `json:"day20"`
			Day22 float64 `json:"day22"`
			Day23 float64 `json:"day23"`
			Day24 float64 `json:"day24"`
			Day25 float64 `json:"day25"`
			Day1  float64 `json:"day1"`
			Day2  float64 `json:"day2"`
			Day3  float64 `json:"day3"`
			Day4  float64 `json:"day4"`
			Day5  float64 `json:"day5"`
			Day6  float64 `json:"day6"`
			Day7  float64 `json:"day7"`
			Day8  float64 `json:"day8"`
			Day13 float64 `json:"day13"`
			Day21 float64 `json:"day21"`
		} `json:"adventRewards2017"`
		AdventRewardsV22018 struct {
			Day1  float64 `json:"day1"`
			Day2  float64 `json:"day2"`
			Day6  float64 `json:"day6"`
			Day8  float64 `json:"day8"`
			Day9  float64 `json:"day9"`
			Day11 float64 `json:"day11"`
			Day16 float64 `json:"day16"`
			Day18 float64 `json:"day18"`
			Day19 float64 `json:"day19"`
			Day20 float64 `json:"day20"`
			Day22 float64 `json:"day22"`
			Day23 float64 `json:"day23"`
			Day24 float64 `json:"day24"`
			Day25 float64 `json:"day25"`
		} `json:"adventRewards_v2_2018"`
		AdventRewards2019 struct {
			Day1  float64 `json:"day1"`
			Day2  float64 `json:"day2"`
			Day4  float64 `json:"day4"`
			Day9  float64 `json:"day9"`
			Day20 float64 `json:"day20"`
			Day22 float64 `json:"day22"`
			Day24 float64 `json:"day24"`
			Day25 float64 `json:"day25"`
		} `json:"adventRewards2019"`
		GiftsGrinch                  float64 `json:"gifts_grinch"`
		CompletedChristmasQuests2019 float64 `json:"completed_christmas_quests_2019"`
		Christmas2019Cooldowns       struct {
			Superstar0 bool `json:"SUPERSTAR0"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			Normal0    bool `json:"NORMAL0"`
			Normal1    bool `json:"NORMAL1"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			Mvp0       bool `json:"MVP0"`
			Mvp1       bool `json:"MVP1"`
			Vip0       bool `json:"VIP0"`
			Vip1       bool `json:"VIP1"`
			MvpPlus1   bool `json:"MVP_PLUS1"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Normal2    bool `json:"NORMAL2"`
			MvpPlus3   bool `json:"MVP_PLUS3"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			Normal3    bool `json:"NORMAL3"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus3   bool `json:"VIP_PLUS3"`
			Mvp2       bool `json:"MVP2"`
			Mvp3       bool `json:"MVP3"`
			Vip2       bool `json:"VIP2"`
			Vip3       bool `json:"VIP3"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar3 bool `json:"SUPERSTAR3"`
		} `json:"christmas2019Cooldowns"`
		Xmas2019ARCADE2            bool      `json:"xmas2019_ARCADE_2"`
		Xmas2019ARCADE3            bool      `json:"xmas2019_ARCADE_3"`
		Xmas2019ARCADE1            bool      `json:"xmas2019_ARCADE_1"`
		Xmas2019PTL3               bool      `json:"xmas2019_PTL_3"`
		ClaimedSoloBank            float64   `json:"claimed_solo_bank"`
		CurrentGadget              string    `json:"currentGadget"`
		AnniversaryNPCVisited2020  []float64 `json:"anniversaryNPCVisited2020"`
		AnniversaryNPCProgress2020 float64   `json:"anniversaryNPCProgress2020"`
		ClaimedPotatoTalisman      float64   `json:"claimed_potato_talisman"`
		ClaimPotatoWarCrown        float64   `json:"claim_potato_war_crown"`
		ClaimedPotatoBasket        float64   `json:"claimed_potato_basket"`
		Summer2020Cooldowns        struct {
			MvpPlus1   bool `json:"MVP_PLUS1"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Superstar0 bool `json:"SUPERSTAR0"`
			Vip2       bool `json:"VIP2"`
			Vip1       bool `json:"VIP1"`
			Vip0       bool `json:"VIP0"`
			Mvp2       bool `json:"MVP2"`
			Mvp1       bool `json:"MVP1"`
			Mvp0       bool `json:"MVP0"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			Normal0    bool `json:"NORMAL0"`
			Normal1    bool `json:"NORMAL1"`
			Normal2    bool `json:"NORMAL2"`
		} `json:"summer2020Cooldowns"`
		Rank                 string  `json:"rank"`
		SkyblockFreeCookie   float64 `json:"skyblock_free_cookie"`
		Easter2021Cooldowns2 struct {
			MvpPlus3   bool `json:"MVP_PLUS3"`
			MvpPlus2   bool `json:"MVP_PLUS2"`
			MvpPlus1   bool `json:"MVP_PLUS1"`
			MvpPlus0   bool `json:"MVP_PLUS0"`
			Normal3    bool `json:"NORMAL3"`
			Normal2    bool `json:"NORMAL2"`
			Normal1    bool `json:"NORMAL1"`
			Normal0    bool `json:"NORMAL0"`
			VipPlus0   bool `json:"VIP_PLUS0"`
			VipPlus1   bool `json:"VIP_PLUS1"`
			VipPlus2   bool `json:"VIP_PLUS2"`
			VipPlus3   bool `json:"VIP_PLUS3"`
			Mvp0       bool `json:"MVP0"`
			Mvp1       bool `json:"MVP1"`
			Mvp2       bool `json:"MVP2"`
			Mvp3       bool `json:"MVP3"`
			Vip0       bool `json:"VIP0"`
			Vip1       bool `json:"VIP1"`
			Vip2       bool `json:"VIP2"`
			Superstar3 bool `json:"SUPERSTAR3"`
			Superstar2 bool `json:"SUPERSTAR2"`
			Superstar1 bool `json:"SUPERSTAR1"`
			Superstar0 bool `json:"SUPERSTAR0"`
			Vip3       bool `json:"VIP3"`
		} `json:"easter2021Cooldowns2"`
		Outfit struct {
		} `json:"outfit"`
		CurrentPet         string  `json:"currentPet"`
		ScorpiusBribe144   float64 `json:"scorpius_bribe_144"`
		ClaimedYear143Cake float64 `json:"claimed_year143_cake"`
		MostRecentGameType string  `json:"mostRecentGameType"`
	} `json:"player"`
}
