package genshindata

type Avatar struct {
	Id              uint64               `json:"Id"`
	Name            string               `json:"Name"`
	NameTextMapHash uint64               `json:"NameTextMapHash"`
	Desc            string               `json:"Desc"`
	DescTextMapHash uint64               `json:"DescTextMapHash"`
	IconName        string               `json:"IconName"`
	WeaponType      string               `json:"WeaponType"`
	LevelMap        map[string]*Property `json:"LevelList"`
}

type Property struct {
	Level            int     `json:"Level"`
	PromoteLevel     float64 `json:"PromoteLevel"`
	Hp               float64 `json:"Hp"`
	Attack           float64 `json:"Attack"`
	Defense          float64 `json:"Defense"`
	Hp_percent       float64 `json:"Hp_percent"`
	Attack_percent   float64 `json:"Attack_percent"`
	Defense_percent  float64 `json:"Defense_percent"`
	Critical         float64 `json:"Critical"`
	CriticalHurt     float64 `json:"CriticalHurt"`
	ChargeEfficiency float64 `json:"ChargeEfficiency"`
	HealActiveUp     float64 `json:"HealActiveUp"`
	HealPassiveUp    float64 `json:"HealPassiveUp"`
	ElementMaster    float64 `json:"ElementMaster"`
	ElementsAddHurt
}

type ElementsAddHurt struct {
	Ice       float64 `json:"Ice"`
	Wind      float64 `json:"Wind"`
	Physical  float64 `json:"Physical"`
	Elec      float64 `json:"Elec"`
	Rock      float64 `json:"Rock"`
	Fire      float64 `json:"Fire"`
	Water     float64 `json:"Water"`
	All       float64 `json:"All"`
	NormalAtk float64 `json:"NormalAtk"`
	HeavyAtk  float64 `json:"HeavyAtk"`
	Skill     float64 `json:"Skill"`
	Ult       float64 `json:"Ult"`
}

type Weapon struct {
	Id              uint64               `json:"Id"`
	Name            string               `json:"Name"`
	NameTextMapHash uint64               `json:"NameTextMapHash"`
	Desc            string               `json:"Desc"`
	DescTextMapHash uint64               `json:"DescTextMapHash"`
	IconName        string               `json:"IconName"`
	WeaponType      string               `json:"WeaponType"`
	SkillAffixMap   map[int]skillAffix   `json:"SkillAffixMap"`
	LevelMap        map[string]*Property `json:"LevelList"`
}

//武器特效
type skillAffix struct {
	Name string `json:"Name"`
	Desc string `json:"Desc"`
	skillAffixData
}

//怪物
type Monster struct {
	Id              uint64                      `json:"Id"`
	Name            string                      `json:"Name"`
	NameTextMapHash uint64                      `json:"NameTextMapHash"`
	MonsterName     string                      `json:"MonsterName"`
	Type            string                      `json:"Type"`
	LevelMap        map[string]*MonsterProperty `json:"LevelList"`
}

type MonsterProperty struct {
	Level   int     `json:"Level"`
	Hp      float64 `json:"Hp"`
	Attack  float64 `json:"Attack"`
	Defense float64 `json:"Defense"`
	subHurtData
}
