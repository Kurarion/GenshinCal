package genshindata

import "strconv"

type Avatar struct {
	Id              uint64               `json:"Id"`
	Name            string               `json:"Name"`
	NameTextMapHash uint64               `json:"NameTextMapHash"`
	Desc            string               `json:"Desc"`
	DescTextMapHash uint64               `json:"DescTextMapHash"`
	IconName        string               `json:"IconName"`
	WeaponType      string               `json:"WeaponType"`
	LevelMap        map[string]*Property `json:"LevelList"`
	SkillDepotId    uint64               `json:"SkillDepotId"`
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
	Grass     float64 `json:"Grass"`
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

//人物技能集
type AvatarSkills struct {
	Id          uint64            `json:"Id"`
	ESkill      AvatarSkillInfo   `json:"ESkill"`
	QSkill      AvatarSkillInfo   `json:"QSkill"`
	ProudSkills []AvatarSkillInfo `json:"ProudSkills"`
	Talents     []AvatarSkillInfo `json:"Talents"`
}

func (c *AvatarSkills) String() string {
	var res = "元素战技: " + c.ESkill.Name +
		"\n" + c.ESkill.Desc +
		"\n元素爆发: " + c.QSkill.Name +
		"\n" + c.QSkill.Desc
	for i := range c.ProudSkills {
		res += "\n天赋" + strconv.Itoa(i+1) + ": " + c.ProudSkills[i].Name +
			"\n" + c.ProudSkills[i].Desc
	}
	for i := range c.Talents {
		res += "\n命座" + strconv.Itoa(i+1) + ": " + c.Talents[i].Name +
			"\n" + c.Talents[i].Desc
	}
	return res
}

//人物技能
type AvatarSkillInfo struct {
	Name string `json:"Name"`
	Desc string `json:"Desc"`
}
