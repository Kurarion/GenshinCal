package genshindata

import (
	"strconv"
)

//取得角色
func GetAvatar(id interface{}) *Avatar {

	var targeID uint64
	switch v := id.(type) {
	case string:
		targeID, _ = strconv.ParseUint(v, 10, 64)
	case uint64:
		targeID = v
	case int:
		targeID = uint64(v)
	}

	return avatar[targeID]
}

//取得角色
func GetAvatarByName(name string) *Avatar {
	targeID, ok := avatarNameMap[name]
	if !ok {
		return nil
	}
	return GetAvatar(targeID)
}

//取得角色列表
func GetAvatarMap() map[uint64]*Avatar {
	return avatar
}

//取得角色全技能
func GetAvatarSkills(id interface{}) *AvatarSkills {

	var targeID uint64
	switch v := id.(type) {
	case string:
		targeID, _ = strconv.ParseUint(v, 10, 64)
	case uint64:
		targeID = v
	case int:
		targeID = uint64(v)
	}

	return avatarSkills[targeID]
}

//取得角色全技能列表
func GetAvatarSkillsMap() map[uint64]*AvatarSkills {
	return avatarSkills
}

//取得武器
func GetWeapon(id interface{}) *Weapon {

	var targeID uint64
	switch v := id.(type) {
	case string:
		targeID, _ = strconv.ParseUint(v, 10, 64)
	case uint64:
		targeID = v
	case int:
		targeID = uint64(v)
	}

	return weapon[targeID]
}

//取得武器
func GetWeaponByName(name string) *Weapon {
	targeID, ok := weaponNameMap[name]
	if !ok {
		return nil
	}
	return GetWeapon(targeID)
}

//取得武器列表
func GetWeaponMap() map[uint64]*Weapon {
	return weapon
}

//取得满级圣遗物主词条
func GetReliquaryMainMap() map[string]float64 {
	return reliquaryMainMap
}

//取得满级圣遗物小词条刻度
func GetReliquaryAffixMap() map[string][]float64 {
	return reliquaryAffixMap
}

//取得怪物
func GetMonster(id interface{}) *Monster {

	var targeID uint64
	switch v := id.(type) {
	case string:
		targeID, _ = strconv.ParseUint(v, 10, 64)
	case uint64:
		targeID = v
	case int:
		targeID = uint64(v)
	}

	return monster[targeID]
}

//取得怪物
func GetMonsterByName(name string) *Monster {
	targeID, ok := monsterNameMap[name]
	if !ok {
		return nil
	}
	return GetMonster(targeID)
}

//取得怪物列表
func GetMonsterMap() map[uint64]*Monster {
	return monster
}

//GetNameFromTypeCode genshindataType名转换属性名
func GetNameFromTypeCode(code string) string {
	name := ""
	switch code {
	case HP, HP_ADD:
		name = "Hp"
	case ATTACK, ATTACK_ADD:
		name = "Attack"
	case DEFENSE, DEFENSE_ADD:
		name = "Defense"
	case HP_PERCENT:
		name = "Hp_percent"
	case ATTACK_PERCENT:
		name = "Attack_percent"
	case DEFENSE_PERCENT:
		name = "Defense_percent"
	case CRITICAL:
		name = "Critical"
	case CRITICAL_HURT:
		name = "CriticalHurt"
	case ICE:
		name = "Ice"
	case WIND:
		name = "Wind"
	case PHYSICAL:
		name = "Physical"
	case ELEC:
		name = "Elec"
	case ROCK:
		name = "Rock"
	case FIRE:
		name = "Fire"
	case WATER:
		name = "Water"
	case GRASS:
		name = "Grass"
	case CHANGE:
		name = "ChargeEfficiency"
	case ELEMENT_MASTER:
		name = "ElementMaster"
	case HEAL:
		name = "HealActiveUp"
	}
	return name
}
