package main

import (
	"eco-engine/table"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

var (
	t                 map[string]*table.Territory // loaded
	ut                map[string]*table.Territory // unloaded
	st                int                         // second tick
	loadedTerritories = make(map[string]*table.Territory)
)

// setInterval function to use later for resgen
func setInterval(f func(map[string]*table.Territory), milliseconds int) chan bool {
	ticker := time.NewTicker(time.Duration(milliseconds) * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				// run func to work on the data
				f(t)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
	return done
}

func clearInterval(done chan bool) {
	done <- true
}

func setTimeout(f func(args ...interface{}), ms int, args ...interface{}) chan bool {
	ticker := time.NewTicker(time.Duration(ms) * time.Millisecond)
	done := make(chan bool)
	go func() {
		select {
		case <-ticker.C:
			f(args...)
		case <-done:
			ticker.Stop()
			return
		}
	}()
	return done
}

func init() {
	// load all upgrades data
	var bytes, err = os.ReadFile("./data/upgrades.json")
	if err != nil {
		panic(err)
	}

	var upgrades table.CostTable

	err = json.Unmarshal(bytes, &upgrades)
	if err != nil {
		panic(err)
	}

	var uninitTerritories map[string]table.RawTerritoryData
	bytes, err = os.ReadFile("./data/baseProperty.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &uninitTerritories)
	if err != nil {
		panic(err)
	}

	// initialise territory
	var territories = make(map[string]*table.Territory, len(uninitTerritories))
	var counter = 0
	for name, data := range uninitTerritories {
		territories[name] = &table.Territory{
			Name: name,
			BaseResourceProduction: table.TerritoryResource{
				Emerald: data.Resources.Emeralds,
				Ore:     data.Resources.Ore,
				Wood:    data.Resources.Wood,
				Fish:    data.Resources.Fish,
				Crop:    data.Resources.Crops,
			},
			Property: table.TerritoryProperty{
				Upgrades: table.TerritoryPropertyUpgradeData{
					Damage:  0,
					Attack:  0,
					Health:  0,
					Defence: 0,
				},
				Bonuses: table.TerritoryPropertyBonusesData{
					StrongerMinions:       0,
					TowerMultiAttack:      0,
					TowerAura:             0,
					TowerVolley:           0,
					LargerResourceStorage: 0,
					LargerEmeraldStorage:  0,
					EfficientResource:     0,
					EfficientEmerald:      0,
					ResourceRate:          0,
					EmeraldRate:           0,
				},
				Tax: table.Tax{
					Ally:   5,
					Others: 5,
				},
				Border:       "Open",
				TradingStyle: "Cheapest",
				HQ:           false,
			},
			Storage: table.TerritoryResourceStorage{
				Capacity: table.TerritoryResource{
					Emerald: 3000,
					Ore:     300,
					Wood:    300,
					Fish:    300,
					Crop:    300,
				},
				Current: table.TerritoryResource{
					Emerald: 0,
					Ore:     0,
					Wood:    0,
					Fish:    0,
					Crop:    0,
				},
			},
			ResourceProduction: table.TerritoryResource{
				Emerald: data.Resources.Emeralds,
				Ore:     data.Resources.Ore,
				Wood:    data.Resources.Wood,
				Fish:    data.Resources.Fish,
				Crop:    data.Resources.Crops,
			},
			TerritoryUsage: table.TerritoryResource{
				Emerald: 0,
				Ore:     0,
				Wood:    0,
				Fish:    0,
				Crop:    0,
			},
			TradingRoutes: data.TradingRoutes,
			ID:            counter,
		}
		counter++
	}
	t = territories
}

func main() {

	var port = os.Args[1]

	if port == "" {
		log.Panicln("$PORT must be set")
	}

	// start http server
	http.HandleFunc("/eco", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

		} else if r.Method == "POST" {
			switch r.URL.Path {
			case "/eco/init":
				// init territories specified in the request body
				var bytes, err = io.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				var territories struct {
					Territories []string `json:"territories"`
					HQ          string   `json:"hq"`
				}
				json.Unmarshal(bytes, &territories)

				var hq = territories.HQ

				// if no hq provided
				if hq == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				var zeroData = &table.TerritoryUpdateData{
					Property: table.TerritoryProperty{
						Upgrades: table.TerritoryPropertyUpgradeData{
							Damage:  0,
							Attack:  0,
							Health:  0,
							Defence: 0,
						},
						Bonuses: table.TerritoryPropertyBonusesData{
							StrongerMinions:       0,
							TowerMultiAttack:      0,
							TowerAura:             0,
							TowerVolley:           0,
							LargerResourceStorage: 0,
							LargerEmeraldStorage:  0,
							EfficientResource:     0,
							EfficientEmerald:      0,
							ResourceRate:          0,
							EmeraldRate:           0,
						},
						Tax: table.Tax{
							Ally:   5,
							Others: 5,
						},
						Border:       "Open",
						TradingStyle: "Cheapest",
						HQ:           false,
					},
				}

				for _, name := range territories.Territories {
					// set all the territory properties to 0 or default
					t[name].Set(*zeroData).SetAllyTax(5).SetOthersTax(60).OpenBorder().Cheapest()
					loadedTerritories[name] = t[name]
					if name == hq {
						t[name].SetHQ()
					}
				}

			}
		}
	})

	http.ListenAndServe(port, nil)

	// run generateResource every 1s and resTick every 60s
	var done = make(chan struct{})
	var secTicker = time.NewTicker(time.Second)
	var minTicker = time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-secTicker.C:
				generateResorce(t)
			case <-minTicker.C:
				resourceTick(t)
			case <-done:
				return
			}
		}
	}()

	fmt.Printf("test")
}

func generateResorce(territories map[string]*table.Territory) {
	// rate means how many seconds it takes to generate n resource
	// n resource is calculated like this
	// nres = base res prod * efficient resource
	// so if rate is level 0 then it takes 1 second to generate 1/4 of the resource
	// and if resource stored in the storage excees the capacity then the excess resource will be lost
	// stoarge capacity is calculated like this
	// cap = base cap * larger storage

	// emerald generation
	for name, territory := range territories {
		// calculate the resource production
		var emeraldRate = float64(territory.Property.Bonuses.EmeraldRate)
		var emeraldProduction = float64(territory.ResourceProduction.Emerald) * (1 + emeraldRate/100)
		var emeraldStorage = float64(territory.Storage.Capacity.Emerald) * (1 + float64(territory.Property.Bonuses.LargerEmeraldStorage)/100)

		// if the storage is full then do nothing
		if float64(territory.Storage.Current.Emerald) >= emeraldStorage {
			continue
		}

		// if the storage is not full then generate the resource

		if float64(territory.Storage.Current.Emerald)+emeraldProduction <= emeraldStorage {
			currEms := float64((territories)[name].Storage.Current.Emerald)
			currEms += emeraldProduction
		} else {
			currEms := float64((territories)[name].Storage.Current.Emerald)
			currEms += emeraldStorage
		}
		fmt.Println(territories)
	}
	if st < 60 {
		st++
	} else {
		for range territories {
			resourceTick(territories)
		}
		st = 0
	}
}

func resourceTick(territories map[string]*table.Territory) {
	// move
	/*
		for _, territory := range territories {
			// dump the resource into next terr's transversing resource
		}
	*/
}

func getPathToHQ(territories map[string]*table.Territory, name string) {
	// get path to hq using djikstra, depending on the trading style
	// fastest  int
	// will find the shortest path while ccheapest will find the shortest path with the least GLOBAL tax
	// if the territory is the hq then return empty array

	// connected nodes (territory) can be found at territories[name].TradingRoutes

	var dist int64 = 1

	if territories[name].Property.HQ {
		return
	}

	var path []string
	var graph = dijkstra.NewGraph()
	var HQID *int
	var currID = territories[name].ID

	// find the id of hq
	for _, territory := range territories {
		if territory.Property.HQ {
			*HQID = territory.ID
			break
		}
	}

	graph.AddVertex(currID)

	for name := range territories {
		// Add logic to compute the shortest path to HQ using Dijkstra's algorithm
		// add current node
		graph.AddVertex(territories[name].ID)

		// add nearby nodes
		for _, connection := range territories[name].TradingRoutes {

			if territories[name].Property.TradingStyle == "Cheapest" && !territories[connection].Claim {
				dist = int64(territories[name].Property.Tax.Others)
			} else if territories[name].Property.TradingStyle == "Cheapest" && territories[connection].Claim {
				dist = 1
			} else {
				dist = 1
			}

			var conn = territories[connection]
			graph.AddVertex(conn.ID)
			graph.AddArc(territories[name].ID, conn.ID, dist)
		}
	}

	var RouteToHQRaw, _ = graph.ShortestSafe(currID, *HQID)
	for _, id := range RouteToHQRaw.Path {
		for name := range territories {
			if territories[name].ID == id {
				path = append(path, name)
			}
		}
	}

	territories[name].RouteToHQ = path
}
