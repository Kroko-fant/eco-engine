use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
pub struct CostTable {
    upgrades_cost: UpgradesCost,
    upgrade_multiplier: UpgradeMultiplier,
    upgrade_base_stats: UpgradeBaseStats,
    bonuses: Bonuses,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct UpgradesCost {
    damage: UpgradeCost,
    attack: UpgradeCost,
    health: UpgradeCost,
    defence: UpgradeCost,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct UpgradeCost {
    value: Vec<i32>,
    resource_type: String,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct UpgradeMultiplier {
    damage: Vec<f64>,
    attack: Vec<f64>,
    health: Vec<f64>,
    defence: Vec<f64>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct UpgradeBaseStats {
    damage: BaseStats,
    attack: Vec<f64>,
    health: Vec<i32>,
    defence: Vec<f64>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct BaseStats {
    min: Vec<i32>,
    max: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct Bonuses {
    stronger_minions: StrongerMinions,
    tower_multi_attacks: TowerMultiAttacks,
    tower_aura: TowerAura,
    tower_volley: TowerVolley,
    xp_seeking: XpSeeking,
    tome_seeking: TomeSeeking,
    emeralds_seeking: EmeraldsSeeking,
    larger_resource_storage: LargerResourceStorage,
    larger_emeralds_storage: LargerEmeraldsStorage,
    efficient_resource: EfficientResource,
    efficient_emeralds: EfficientEmeralds,
    resource_rate: ResourceRate,
    emeralds_rate: EmeraldsRate,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct StrongerMinions {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<f64>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TowerMultiAttacks {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TowerAura {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TowerVolley {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct XpSeeking {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TomeSeeking {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<f64>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct EmeraldsSeeking {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<f64>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct LargerResourceStorage {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct LargerEmeraldsStorage {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct LargerEmeraldsStorage {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct EfficientResource {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<f64>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct EfficientEmeralds {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<f64>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct ResourceRate {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct EmeraldsRate {
    max_level: i32,
    cost: Vec<i32>,
    resource_type: String,
    value: Vec<i32>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TerritoryProperty {
  upgrade: TerritoryPropertyUpgradeData,
  bonuses: TerritoryPropertyBonusesData,
  tax: i32,
  border: String,
  trading_style: String,
  hq: String,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TerritoryPropertyUpgradeData {
  damage: i32,
  attack: i32,
  health: i32,
  defence: i32,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TerritoryPropertyBonusesData {
  stronger_minions: i32,
  tower_multi_attack: i32,
  tower_aura: i32,
  tower_volley: i32,
  larger_resource_storage: i32,
  larger_emeralds_storage: i32,
  efficient_resource: i32,
  efficient_emeralds: i32,
  resource_rate: i32,
  emeralds_rate: i32,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct Tax {
  ally: i32,
  other: i32
}

/*
type TerritoryResource struct {
	Emerald int
	Ore     int
	Wood    int
	Fish    int
	Crop    int
}

*/
#[derive(Debug, Deserialize, Serialize)]
pub struct TerritoryResource {
  emeralds: i32,
  ore: i32,
  wood: i32,
  fish: i32,
  crops: i32,
}

/*
type Territory struct {
	Name                   string
	BaseResourceProduction TerritoryResource
	Property               TerritoryProperty
	Storage                TerritoryResourceStorage
	ResourceProduction     TerritoryResource
	TerritoryUsage         TerritoryResource
	Type                   string
	TradingRoutes          []string
	RouteToHQ              []string
}
*/

#[derive(Debug, Deserialize, Serialize)]
pub struct Territory {
  name: String,
  base_resource_production: TerritoryResource,
  property: TerritoryProperty,
  storage: TerritoryResourceStorage,
  resource_production: TerritoryResource,
  territory_usage: TerritoryResource,
  territory_type: String,
  trading_routes: Vec<String>,
  route_to_hq: Vec<String>,
}

/*
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
*/

#[derive(Debug, Deserialize, Serialize)]
pub struct TerritoryResourceStorage {
  capacity: TerritoryResource,
  current: TerritoryResource,
}

/*
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
*/

#[derive(Debug, Deserialize, Serialize)]
pub struct RawTerritoryData {
  resources: TerritoryResource,
  trading_routes: Vec<String>,
  territory_type: String,
}
/*
type TerritoryUpdateData struct {
	Territory string
	Property  TerritoryProperty
}
*/

#[derive(Debug, Deserialize, Serialize)]
pub struct TerritoryUpdateData {
  territory: String,
  property: TerritoryProperty,
}

fn main() {
    // Deserialize the JSON into a CostTable struct
    
    let cost_table: CostTable = serde_json::from_str(json_data).unwrap();

    // Print the deserialized cost_table struct
    println!("{:#?}", cost_table);
}