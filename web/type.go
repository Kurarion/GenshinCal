package web

import (
	data "genshincal/genshindata"
)

type InitData struct {
	Avatar  AvatarMap
	Weapon  WeaponMap
	Monster MonsterMap
}

type AvatarMap map[uint64]*data.Avatar
type WeaponMap map[uint64]*data.Weapon
type MonsterMap map[uint64]*data.Monster

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

//GetLevelList
func (c MonsterMap) GetLevelList() map[string]*data.MonsterProperty {
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

//GetReliquaryList
func (c InitData) GetReliquaryList() [][]map[string][][]string {
	getNameFunc := data.GetNameFromTypeCode

	affix := func() [][]string {
		a := []([]string){
			[]string{getNameFunc(data.ATTACK_ADD), data.ATTACK_ADD, data.ATTACK_ADD_NAME},
			[]string{getNameFunc(data.HP_PERCENT), data.HP_PERCENT, data.HP_PERCENT_NAME},
			[]string{getNameFunc(data.ATTACK_ADD), data.ATTACK_ADD, data.ATTACK_ADD_NAME},
			[]string{getNameFunc(data.ATTACK_PERCENT), data.ATTACK_PERCENT, data.ATTACK_PERCENT_NAME},
			[]string{getNameFunc(data.DEFENSE_ADD), data.DEFENSE_ADD, data.DEFENSE_ADD_NAME},
			[]string{getNameFunc(data.DEFENSE_PERCENT), data.DEFENSE_PERCENT, data.DEFENSE_PERCENT_NAME},
			[]string{getNameFunc(data.CHANGE), data.CHANGE, data.CHANGE_NAME},
			[]string{getNameFunc(data.ELEMENT_MASTER), data.ELEMENT_MASTER, data.ELEMENT_MASTER_NAME},
			[]string{getNameFunc(data.CRITICAL), data.CRITICAL, data.CRITICAL_NAME},
			[]string{getNameFunc(data.CRITICAL_HURT), data.CRITICAL_HURT, data.CRITICAL_HURT_NAME},
		}

		return a
	}()

	res := []([](map[string][][]string)){
		[](map[string][][]string){
			map[string][][]string{
				"Main": ([]([]string){[]string{getNameFunc(data.HP_ADD), data.HP_ADD, data.HP_ADD_NAME}}),
			},
			map[string][][]string{
				"Affix0": affix,
				"Affix1": affix,
				"Affix2": affix,
				"Affix3": affix,
			},
		},
		[](map[string][][]string){
			map[string][][]string{
				"Main": ([]([]string){[]string{getNameFunc(data.ATTACK_ADD), data.ATTACK_ADD, data.ATTACK_ADD_NAME}}),
			},
			map[string][][]string{
				"Affix0": affix,
				"Affix1": affix,
				"Affix2": affix,
				"Affix3": affix,
			},
		},
		[](map[string][][]string){
			map[string][][]string{
				"Main": ([]([]string){
					[]string{getNameFunc(data.CHANGE), data.CHANGE, data.CHANGE_NAME},
					[]string{getNameFunc(data.HP_PERCENT), data.HP_PERCENT, data.HP_PERCENT_NAME},
					[]string{getNameFunc(data.ATTACK_PERCENT), data.ATTACK_PERCENT, data.ATTACK_PERCENT_NAME},
					[]string{getNameFunc(data.DEFENSE_PERCENT), data.DEFENSE_PERCENT, data.DEFENSE_PERCENT_NAME},
					[]string{getNameFunc(data.ELEMENT_MASTER), data.ELEMENT_MASTER, data.ELEMENT_MASTER_NAME},
				}),
			},
			map[string][][]string{
				"Affix0": affix,
				"Affix1": affix,
				"Affix2": affix,
				"Affix3": affix,
			},
		},
		[](map[string][][]string){
			map[string][][]string{
				"Main": ([]([]string){
					[]string{getNameFunc(data.ICE), data.ICE, data.ICE_NAME},
					[]string{getNameFunc(data.WIND), data.WIND, data.WIND_NAME},
					[]string{getNameFunc(data.PHYSICAL), data.PHYSICAL, data.PHYSICAL_NAME},
					[]string{getNameFunc(data.ELEC), data.ELEC, data.ELEC_NAME},
					[]string{getNameFunc(data.ROCK), data.ROCK, data.ROCK_NAME},
					[]string{getNameFunc(data.FIRE), data.FIRE, data.FIRE_NAME},
					[]string{getNameFunc(data.WATER), data.WATER, data.WATER_NAME},
					[]string{getNameFunc(data.HP_PERCENT), data.HP_PERCENT, data.HP_PERCENT_NAME},
					[]string{getNameFunc(data.ATTACK_PERCENT), data.ATTACK_PERCENT, data.ATTACK_PERCENT_NAME},
					[]string{getNameFunc(data.DEFENSE_PERCENT), data.DEFENSE_PERCENT, data.DEFENSE_PERCENT_NAME},
					[]string{getNameFunc(data.ELEMENT_MASTER), data.ELEMENT_MASTER, data.ELEMENT_MASTER_NAME},
				}),
			},
			map[string][][]string{
				"Affix0": affix,
				"Affix1": affix,
				"Affix2": affix,
				"Affix3": affix,
			},
		},
		[](map[string][][]string){
			map[string][][]string{
				"Main": ([]([]string){
					[]string{getNameFunc(data.CRITICAL), data.CRITICAL, data.CRITICAL_NAME},
					[]string{getNameFunc(data.CRITICAL_HURT), data.CRITICAL_HURT, data.CRITICAL_HURT_NAME},
					[]string{getNameFunc(data.HP_PERCENT), data.HP_PERCENT, data.HP_PERCENT_NAME},
					[]string{getNameFunc(data.ATTACK_PERCENT), data.ATTACK_PERCENT, data.ATTACK_PERCENT_NAME},
					[]string{getNameFunc(data.DEFENSE_PERCENT), data.DEFENSE_PERCENT, data.DEFENSE_PERCENT_NAME},
					[]string{getNameFunc(data.ELEMENT_MASTER), data.ELEMENT_MASTER, data.ELEMENT_MASTER_NAME},
				}),
			},
			map[string][][]string{
				"Affix0": affix,
				"Affix1": affix,
				"Affix2": affix,
				"Affix3": affix,
			},
		},
	}

	return res
}
