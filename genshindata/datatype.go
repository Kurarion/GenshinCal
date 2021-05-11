package genshindata

//人物定义
type avatarBaseData struct {
	Id               uint64           `json:"Id"`
	NameTextMapHash  uint64           `json:"NameTextMapHash"`
	DescTextMapHash  uint64           `json:"DescTextMapHash"`
	IconName         string           `json:"IconName"`
	WeaponType       string           `json:"WeaponType"`
	HpBase           float64          `json:"HpBase"`
	AttackBase       float64          `json:"AttackBase"`
	DefenseBase      float64          `json:"DefenseBase"`
	Critical         float64          `json:"Critical"`
	CriticalHurt     float64          `json:"CriticalHurt"`
	ChargeEfficiency float64          `json:"ChargeEfficiency"`
	PropGrowCurves   []propGrowCurves `json:"PropGrowCurves"`
	AvatarPromoteId  uint64           `json:"AvatarPromoteId"`
	SkillDepotId     uint64           `json:"SkillDepotId"`
}
type avatarBaseListData []avatarBaseData

//武器定义
type weaponBaseData struct {
	Id              uint64                 `json:"Id"`
	NameTextMapHash uint64                 `json:"NameTextMapHash"`
	DescTextMapHash uint64                 `json:"DescTextMapHash"`
	RankLevel       int                    `json:"RankLevel"`
	IconName        string                 `json:"Icon"`
	WeaponType      string                 `json:"WeaponType"`
	PropGrowCurves  []weaponPropGrowCurves `json:"WeaponProp"`
	WeaponPromoteId uint64                 `json:"WeaponPromoteId"`
	SkillAffix      []uint64               `json:"SkillAffix"`
}
type weaponBaseListData []weaponBaseData

//成长参数
type propGrowCurves struct {
	Type  string `json:"Type"`
	Value string `json:"GrowCurve"`
}

//武器成长参数
type weaponPropGrowCurves struct {
	PropType  string  `json:"PropType"`
	InitValue float64 `json:"InitValue"`
	Type      string  `json:"Type"`
}

//人物/武器 基础升级提升
type growCurvesData struct {
	Level      int     `json:"Level"`
	CurveInfos []curve `json:"CurveInfos"`
}
type growCurvesListData []growCurvesData

//成长类型
type curve struct {
	Type  string  `json:"Type"`
	Value float64 `json:"Value"`
}

//人物/武器 突破提升值
type promoteData struct {
	AvatarPromoteId     uint64    `json:"AvatarPromoteId"`
	WeaponPromoteId     uint64    `json:"WeaponPromoteId"`
	PromoteLevel        int       `json:"PromoteLevel"`
	AddProps            []propAdd `json:"AddProps"`
	RequiredPlayerLevel int       `json:"RequiredPlayerLevel"`
	UnlockMaxLevel      int       `json:"UnlockMaxLevel"`
}
type promoteListData []promoteData

func (this promoteListData) Len() int      { return len(this) }
func (this promoteListData) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this promoteListData) Less(i, j int) bool {
	return this[i].AvatarPromoteId < this[j].AvatarPromoteId
}

//属性增量
type propAdd struct {
	PropType string  `json:"PropType"`
	Value    float64 `json:"Value"`
}

//武器特效
type skillAffixData struct {
	Id              uint64    `json:"Id"`
	AffixId         uint64    `json:"AffixId"`
	Level           int       `json:"Level"`
	NameTextMapHash uint64    `json:"NameTextMapHash"`
	DescTextMapHash uint64    `json:"DescTextMapHash"`
	ParamList       []float64 `json:"ParamList"`
}
type skillAffixListData []skillAffixData

//圣遗物小词条刻度
type reliquaryAffix struct {
	Id        uint64  `json:"Id"`
	DepotId   int     `json:"DepotId"`
	PropType  string  `json:"PropType"`
	PropValue float64 `json:"PropValue"`
	Weight    int     `json:"Weight"`
}
type reliquaryAffixListData []reliquaryAffix

//圣遗物主词条
type reliquaryMain struct {
	Rank     int       `json:"Rank"`
	Level    int       `json:"Level"`
	AddProps []propAdd `json:"AddProps"`
}
type reliquaryMainListData []reliquaryMain

//怪物定义
type monsterBaseData struct {
	Id              uint64           `json:"Id"`
	MonsterName     string           `json:"MonsterName"`
	Type            string           `json:"Type"`
	NameTextMapHash uint64           `json:"NameTextMapHash"`
	HpBase          float64          `json:"HpBase"`
	AttackBase      float64          `json:"AttackBase"`
	DefenseBase     float64          `json:"DefenseBase"`
	PropGrowCurves  []propGrowCurves `json:"PropGrowCurves"`
	subHurtData
}
type monsterBaseListData []monsterBaseData

//抗性定义
type subHurtData struct {
	FireSubHurt     float64 `json:"FireSubHurt"`
	GrassSubHurt    float64 `json:"GrassSubHurt"`
	WaterSubHurt    float64 `json:"WaterSubHurt"`
	ElecSubHurt     float64 `json:"ElecSubHurt"`
	WindSubHurt     float64 `json:"WindSubHurt"`
	IceSubHurt      float64 `json:"IceSubHurt"`
	RockSubHurt     float64 `json:"RockSubHurt"`
	PhysicalSubHurt float64 `json:"PhysicalSubHurt"`
}

//人物全技能列表
type avatarSkillsData struct {
	Id                      uint64                     `json:"Id"`
	EnergySkill             uint64                     `json:"EnergySkill"`
	Skills                  []uint64                   `json:"Skills"`
	InherentProudSkillOpens []avatarInherentProudSkill `json:"InherentProudSkillOpens"`
	Talents                 []uint64                   `json:"Talents"`
}
type avatarSkillsListData []avatarSkillsData

//人物天赋子结构
type avatarInherentProudSkill struct {
	ProudSkillGroupId      uint64 `json:"ProudSkillGroupId"`
	NeedAvatarPromoteLevel int    `json:"NeedAvatarPromoteLevel"`
}

//人物EQ技能
type avatarSkillData struct {
	Id              uint64  `json:"Id"`
	NameTextMapHash uint64  `json:"NameTextMapHash"`
	DescTextMapHash uint64  `json:"DescTextMapHash"`
	CdTime          float64 `json:"CdTime"`
}
type avatarSkillListData []avatarSkillData

//人物天赋
type avatarProudSkillData struct {
	ProudSkillGroupId uint64 `json:"ProudSkillGroupId"`
	NameTextMapHash   uint64 `json:"NameTextMapHash"`
	DescTextMapHash   uint64 `json:"DescTextMapHash"`
}
type avatarProudSkillListData []avatarProudSkillData

//人物命座
type avatarTalentData struct {
	TalentId        uint64 `json:"TalentId"`
	NameTextMapHash uint64 `json:"NameTextMapHash"`
	DescTextMapHash uint64 `json:"DescTextMapHash"`
}
type avatarTalentListData []avatarTalentData
