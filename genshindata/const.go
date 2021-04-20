package genshindata

//突破属性类型
//基础
const (
	HP      = "FIGHT_PROP_BASE_HP"
	ATTACK  = "FIGHT_PROP_BASE_ATTACK"
	DEFENSE = "FIGHT_PROP_BASE_DEFENSE"
)

//额外
const (
	HP_PERCENT      = "FIGHT_PROP_HP_PERCENT"
	ATTACK_PERCENT  = "FIGHT_PROP_ATTACK_PERCENT"
	DEFENSE_PERCENT = "FIGHT_PROP_DEFENSE_PERCENT"
	CRITICAL        = "FIGHT_PROP_CRITICAL"
	CRITICAL_HURT   = "FIGHT_PROP_CRITICAL_HURT"
	ICE             = "FIGHT_PROP_ICE_ADD_HURT"
	WIND            = "FIGHT_PROP_WIND_ADD_HURT"
	PHYSICAL        = "FIGHT_PROP_PHYSICAL_ADD_HURT"
	ELEC            = "FIGHT_PROP_ELEC_ADD_HURT"
	ROCK            = "FIGHT_PROP_ROCK_ADD_HURT"
	FIRE            = "FIGHT_PROP_FIRE_ADD_HURT"
	WATER           = "FIGHT_PROP_WATER_ADD_HURT"
	CHANGE          = "FIGHT_PROP_CHARGE_EFFICIENCY"
	ELEMENT_MASTER  = "FIGHT_PROP_ELEMENT_MASTERY"
	HEAL            = "FIGHT_PROP_HEAL_ADD"
)

//人物成长曲线类型
const (
	C_HP5      = "GROW_CURVE_HP_S5"
	C_HP4      = "GROW_CURVE_HP_S4"
	C_ATTACK5  = "GROW_CURVE_ATTACK_S5"
	C_ATTACK4  = "GROW_CURVE_ATTACK_S4"
	C_DEFENSE5 = "GROW_CURVE_HP_S5"
	C_DEFENSE4 = "GROW_CURVE_HP_S4"
)

//武器成长曲线类型
const (
	C_W_ATTACK101   = "GROW_CURVE_ATTACK_101"
	C_W_ATTACK102   = "GROW_CURVE_ATTACK_102"
	C_W_ATTACK103   = "GROW_CURVE_ATTACK_103"
	C_W_ATTACK104   = "GROW_CURVE_ATTACK_104"
	C_W_ATTACK105   = "GROW_CURVE_ATTACK_105"
	C_W_CRITICAL101 = "GROW_CURVE_CRITICAL_101"
	C_W_ATTACK201   = "GROW_CURVE_ATTACK_201"
	C_W_ATTACK202   = "GROW_CURVE_ATTACK_202"
	C_W_ATTACK203   = "GROW_CURVE_ATTACK_203"
	C_W_ATTACK204   = "GROW_CURVE_ATTACK_204"
	C_W_ATTACK205   = "GROW_CURVE_ATTACK_205"
	C_W_CRITICAL201 = "GROW_CURVE_CRITICAL_201"
	C_W_ATTACK301   = "GROW_CURVE_ATTACK_301"
	C_W_ATTACK302   = "GROW_CURVE_ATTACK_302"
	C_W_ATTACK303   = "GROW_CURVE_ATTACK_303"
	C_W_ATTACK304   = "GROW_CURVE_ATTACK_304"
	C_W_ATTACK305   = "GROW_CURVE_ATTACK_305"
	C_W_CRITICAL301 = "GROW_CURVE_CRITICAL_301"
)

//文件信息
const (
	repositoryURL = `https://raw.githubusercontent.com/Dimbreath/GenshinData/master`
	//人物定义
	avatarExcelConfigData = "/ExcelBinOutput/AvatarExcelConfigData.json"
	//人物基础升级提升值(累加)
	avatarCurveExcelConfigData = "/ExcelBinOutput/AvatarCurveExcelConfigData.json"
	//人物突破提升值(对应)
	avatarPromoteExcelConfigData = "/ExcelBinOutput/AvatarPromoteExcelConfigData.json"
	//文字代码对应表
	textMapData = "/TextMap/TextCHS.json"
	//武器定义
	weaponExcelConfigData = "/ExcelBinOutput/WeaponExcelConfigData.json"
	//武器基础升级提升值(对应)
	weaponCurveExcelConfigData = "/ExcelBinOutput/WeaponCurveExcelConfigData.json"
	//武器突破提升值(对应)
	weaponPromoteExcelConfigData = "/ExcelBinOutput/WeaponPromoteExcelConfigData.json"
	//武器特效
	EquipAffixExcelConfigData = "/ExcelBinOutput/EquipAffixExcelConfigData.json"
	//圣遗物词条提升值
	ReliquaryAffixExcelConfigData = "/ExcelBinOutput/ReliquaryAffixExcelConfigData.json"
	//圣遗物主词条值
	ReliquaryLevelExcelConfigData = "/ExcelBinOutput/ReliquaryLevelExcelConfigData.json"
)
