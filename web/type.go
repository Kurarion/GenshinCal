package web

import (
	data "genshincal/genshindata"
)

type InitData struct {
	Avatar AvatarMap
	Weapon WeaponMap
}

type AvatarMap map[uint64]*data.Avatar
type WeaponMap map[uint64]*data.Weapon

//GetLevelList
func (c AvatarMap) GetLevelList() map[string]*data.Property {
	for i := range c {
		return c[i].LevelMap
	}
	return nil
}

//GetLevelList
func (c WeaponMap) GetLevelList() map[string]*data.Property {
	for i := range c {
		return c[i].LevelMap
	}
	return nil
}

func getWeaponTypeMap() map[string]string {
	list := map[string]string{
		"单手剑":  data.WEAPON_SWORD_ONE_HAND,
		"双手剑":  data.WEAPON_CLAYMORE,
		"长柄武器": data.WEAPON_POLE,
		"法器":   data.WEAPON_CATALYST,
		"弓":    data.WEAPON_BOW,
	}
	return list
}

//GetWeaponTypeList
func (c AvatarMap) GetWeaponTypeList() map[string]string {
	return getWeaponTypeMap()
}

//GetWeaponTypeList
func (c WeaponMap) GetWeaponTypeList() map[string]string {
	return getWeaponTypeMap()
}

func getSkillAffixLevelMap() map[string]string {
	list := map[string]string{
		"精1": "1",
		"精2": "2",
		"精3": "3",
		"精4": "4",
		"精5": "5",
	}
	return list
}

//GetWeaponTypeList
func (c WeaponMap) GetSkillAffixLevelList() map[string]string {
	return getSkillAffixLevelMap()
}
