package table

type CostTable struct {
	UpgradesCost struct {
		Damage struct {
			Value        []int  `json:"value"`
			ResourceType string `json:"resourceType"`
		} `json:"damage"`
		Attack struct {
			Value        []int  `json:"value"`
			ResourceType string `json:"resourceType"`
		} `json:"attack"`
		Health struct {
			Value        []int  `json:"value"`
			ResourceType string `json:"resourceType"`
		} `json:"health"`
		Defence struct {
			Value        []int  `json:"value"`
			ResourceType string `json:"resourceType"`
		} `json:"defence"`
	} `json:"upgradesCost"`
	UpgradeMultiplier struct {
		Damage  []float64 `json:"damage"`
		Attack  []float64 `json:"attack"`
		Health  []float64 `json:"health"`
		Defence []float64 `json:"defence"`
	} `json:"upgradeMultiplier"`
	UpgradeBaseStats struct {
		Damage struct {
			Min []int `json:"min"`
			Max []int `json:"max"`
		} `json:"damage"`
		Attack  []float64 `json:"attack"`
		Health  []int     `json:"health"`
		Defence []float64 `json:"defence"`
	} `json:"upgradeBaseStats"`
	Bonuses struct {
		StrongerMinions struct {
			MaxLevel     int       `json:"maxLevel"`
			Cost         []int     `json:"cost"`
			ResourceType string    `json:"resourceType"`
			Value        []float64 `json:"value"`
		} `json:"strongerMinions"`
		TowerMultiAttacks struct {
			MaxLevel     int    `json:"maxLevel"`
			Cost         []int  `json:"cost"`
			ResourceType string `json:"resourceType"`
			Value        []int  `json:"value"`
		} `json:"towerMultiAttacks"`
		TowerAura struct {
			MaxLevel     int           `json:"maxLevel"`
			Cost         []int         `json:"cost"`
			ResourceType string        `json:"resourceType"`
			Value        []interface{} `json:"value"`
		} `json:"towerAura"`
		TowerVolley struct {
			MaxLevel     int           `json:"maxLevel"`
			Cost         []int         `json:"cost"`
			ResourceType string        `json:"resourceType"`
			Value        []interface{} `json:"value"`
		} `json:"towerVolley"`
		XpSeeking struct {
			MaxLevel     int    `json:"maxLevel"`
			Cost         []int  `json:"cost"`
			ResourceType string `json:"resourceType"`
			Value        []int  `json:"value"`
		} `json:"xpSeeking"`
		TomeSeeking struct {
			MaxLevel     int       `json:"maxLevel"`
			Cost         []int     `json:"cost"`
			ResourceType string    `json:"resourceType"`
			Value        []float64 `json:"value"`
		} `json:"tomeSeeking"`
		EmeraldsSeeking struct {
			MaxLevel     int       `json:"maxLevel"`
			Cost         []int     `json:"cost"`
			ResourceType string    `json:"resourceType"`
			Value        []float64 `json:"value"`
		} `json:"emeraldsSeeking"`
		LargerResourceStorage struct {
			MaxLevel     int    `json:"maxLevel"`
			Cost         []int  `json:"cost"`
			ResourceType string `json:"resourceType"`
			Value        []int  `json:"value"`
		} `json:"largerResourceStorage"`
		LargerEmeraldsStorage struct {
			MaxLevel     int    `json:"maxLevel"`
			Cost         []int  `json:"cost"`
			ResourceType string `json:"resourceType"`
			Value        []int  `json:"value"`
		} `json:"largerEmeraldsStorage"`
		EfficientResource struct {
			MaxLevel     int       `json:"maxLevel"`
			Cost         []int     `json:"cost"`
			ResourceType string    `json:"resourceType"`
			Value        []float64 `json:"value"`
		} `json:"efficientResource"`
		EfficientEmeralds struct {
			MaxLevel     int       `json:"maxLevel"`
			Cost         []int     `json:"cost"`
			ResourceType string    `json:"resourceType"`
			Value        []float64 `json:"value"`
		} `json:"efficientEmeralds"`
		ResourceRate struct {
			MaxLevel     int    `json:"maxLevel"`
			Cost         []int  `json:"cost"`
			ResourceType string `json:"resourceType"`
			Value        []int  `json:"value"`
		} `json:"resourceRate"`
		EmeraldsRate struct {
			MaxLevel     int    `json:"maxLevel"`
			Cost         []int  `json:"cost"`
			ResourceType string `json:"resourceType"`
			Value        []int  `json:"value"`
		} `json:"emeraldsRate"`
	} `json:"bonuses"`
}

type TerritoryProperty struct {
	Upgrades     TerritoryPropertyUpgradeData
	Bonuses      TerritoryPropertyBonusesData
	Tax          Tax
	Border       string
	TradingStyle string
	HQ           bool
}

type TerritoryPropertyUpgradeData struct {
	Damage  int
	Attack  int
	Health  int
	Defence int
}

type TerritoryPropertyBonusesData struct {
	StrongerMinions       int
	TowerMultiAttack      int
	TowerAura             int
	TowerVolley           int
	LargerResourceStorage int
	LargerEmeraldStorage  int
	EfficientResource     int
	EfficientEmerald      int
	ResourceRate          int
	EmeraldRate           int
}

type Tax struct {
	Ally   int
	Others int
}

type TerritoryResource struct {
	Emerald int
	Ore     int
	Wood    int
	Fish    int
	Crop    int
}

type Territory struct {
	Name                   string
	BaseResourceProduction TerritoryResource
	Property               TerritoryProperty
	Storage                TerritoryResourceStorage
	TransversingResource   TerritoryResource
	ResourceProduction     TerritoryResource
	TerritoryUsage         TerritoryResource
	Type                   string
	TradingRoutes          []string
	RouteToHQ              []string
	Claim                  bool
	ID                     int
}

type TerritoryResourceStorage struct {
	Capacity struct {
		Emerald int
		Ore     int
		Wood    int
		Fish    int
		Crop    int
	}
	Current struct {
		Emerald int
		Ore     int
		Wood    int
		Fish    int
		Crop    int
	}
}

type RawTerritoryData struct {
	Resources struct {
		Emeralds int `json:"emeralds"`
		Ore      int `json:"ore"`
		Crops    int `json:"crops"`
		Fish     int `json:"fish"`
		Wood     int `json:"wood"`
	} `json:"resources"`
	TradingRoutes []string `json:"Trading Routes"`
	Type          string
}

type TerritoryUpdateData struct {
	Property TerritoryProperty
}

func (t *Territory) Set(d TerritoryUpdateData) *Territory {
	// validate the territory before setting
	// Damage, Attack, Health and Defence must be between 0 and 11
	// StrongerMinions 0 - 3, Tower MultiAttack 0 - 1, Tower Aura and Volley 0 - 3
	// Larger Emerald and Resource storage 0 - 6, Efficient Resource 0 - 6, Efficient Emerald 0 - 3 and Resource and Emerald Rate 0 - 3
	if d.Property.Upgrades.Damage < 0 || d.Property.Upgrades.Damage > 11 {
		return t
	} else if d.Property.Upgrades.Attack < 0 || d.Property.Upgrades.Attack > 11 {
		return t
	} else if d.Property.Upgrades.Health < 0 || d.Property.Upgrades.Health > 11 {
		return t
	} else if d.Property.Upgrades.Defence < 0 || d.Property.Upgrades.Defence > 11 {
		return t
	} else if d.Property.Bonuses.StrongerMinions < 0 || d.Property.Bonuses.StrongerMinions > 3 {
		return t
	} else if d.Property.Bonuses.TowerMultiAttack < 0 || d.Property.Bonuses.TowerMultiAttack > 1 {
		return t
	} else if d.Property.Bonuses.TowerAura < 0 || d.Property.Bonuses.TowerAura > 3 {
		return t
	} else if d.Property.Bonuses.TowerVolley < 0 || d.Property.Bonuses.TowerVolley > 3 {
		return t
	} else if d.Property.Bonuses.LargerResourceStorage < 0 || d.Property.Bonuses.LargerResourceStorage > 6 {
		return t
	} else if d.Property.Bonuses.LargerEmeraldStorage < 0 || d.Property.Bonuses.LargerEmeraldStorage > 6 {
		return t
	} else if d.Property.Bonuses.EfficientResource < 0 || d.Property.Bonuses.EfficientResource > 6 {
		return t
	} else if d.Property.Bonuses.EfficientEmerald < 0 || d.Property.Bonuses.EfficientEmerald > 3 {
		return t
	} else if d.Property.Bonuses.ResourceRate < 0 || d.Property.Bonuses.ResourceRate > 3 {
		return t
	} else if d.Property.Bonuses.EmeraldRate < 0 || d.Property.Bonuses.EmeraldRate > 3 {
		return t
	}

	t.Property = d.Property
	return t
}

func (t *Territory) Update() *Territory {
	// update the territory stats and cost

	return t
}

func (t *Territory) CloseBorder() *Territory {
	t.Property.Border = "Closed"
	return t
}

func (t *Territory) OpenBorder() *Territory {
	t.Property.Border = "Open"
	return t
}

func (t *Territory) Fastest() *Territory {
	t.Property.TradingStyle = "Fastest"
	return t
}

func (t *Territory) Cheapest() *Territory {
	t.Property.TradingStyle = "Cheapest"
	return t
}

func (t *Territory) ToggleBorder() *Territory {
	if t.Property.Border == "Closed" {
		t.Property.Border = "Open"
	} else {
		t.Property.Border = "Closed"
	}
	return t
}

func (t *Territory) UnsetHQ() *Territory {
	t.Property.HQ = false
	return t
}

func (t *Territory) SetAllyTax(n int) *Territory {
	// tax has to be within 5 and 60
	if n < 5 || n > 60 {
		return t
	}
	t.Property.Tax.Ally = n
	return t
}

func (t *Territory) SetOthersTax(n int) *Territory {
	// tax has to be within 5 and 60
	if n < 5 || n > 60 {
		return t
	}
	t.Property.Tax.Others = n
	return t
}

func (t *Territory) AddTradingRoute(r string) *Territory {
	t.TradingRoutes = append(t.TradingRoutes, r)
	return t
}

func (t *Territory) SetHQ() *Territory {
	t.Property.HQ = true
	return t
}
