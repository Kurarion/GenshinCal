package calc

//伤害类型
type DAMAGETYPE int

//伤害定义
const (
	NORMALATK DAMAGETYPE = iota + 1
	THUMP
	E
	Q
	PHYSICAL
	FIRE
	WATER
	ICE
	ELEC
	WIND
	ROCK
	GRASS
	ALL
)

//最小有效伤害
const MinActiveDamageType = NORMALATK

//最大有效伤害
const MaxActiveDamageType = ALL

//伤害中文名
var damgeTypeNameCN [MaxActiveDamageType + 1]string

func init() {
	//伤害名初始化
	damgeTypeNameCN = [...]string{"", "普攻", "重击", "元素战技", "元素爆发", "物理", "火元素", "水元素", "冰元素", "雷元素", "风元素", "岩元素", "草元素", "全伤害"}
	//圣遗物位置可选词条
	selectableEntry = [...][]string{
		make([]string, 0, 0),
		[]string{"Boold"},
		[]string{"Atk"},
		[]string{"ChargeEffBuff", "BooldBuffRate", "AtkBuffRate", "DefBuffRate", "EleValueBuff"},
		[]string{"DamageBoosts", "BooldBuffRate", "AtkBuffRate", "DefBuffRate", "EleValueBuff"},
		[]string{"CriticalBuff", "CriticalHurtBuff", "BooldBuffRate", "AtkBuffRate", "DefBuffRate", "EleValueBuff"},
	}
}

//GetName 名字取得
func (c DAMAGETYPE) GetName() string {
	return damgeTypeNameCN[int(c)]
}

//基础属性
type StateBase struct {
	Atk          float32        `json:"Atk"`
	Def          float32        `json:"Def"`
	Boold        float32        `json:"Boold"`
	Critical     float32        `json:"Critical"`
	CriticalHurt float32        `json:"CriticalHurt"`
	EleValue     float32        `json:"EleValue"`
	ChargeEff    float32        `json:"ChargeEff"`
	DamageBoosts DamageBoostMap `json:"DamageBoosts"`
}

func newStateBase() StateBase {
	return StateBase{
		DamageBoosts: newDamageBoostMap(),
	}
}

//基础增幅属性
type StateBuff struct {
	AtkBuffRate      float32        `json:"AtkBuffRate"`
	DefBuffRate      float32        `json:"DefBuffRate"`
	BooldBuffRate    float32        `json:"BooldBuffRate"`
	CriticalBuff     float32        `json:"CriticalBuff"`
	CriticalHurtBuff float32        `json:"CriticalHurtBuff"`
	EleValueBuff     float32        `json:"EleValueBuff"`
	ChargeEffBuff    float32        `json:"ChargeEffBuff"`
	DamageBoostBuffs DamageBoostMap `json:"DamageBoostBuffs"`
}

func newStateBuff() StateBuff {
	return StateBuff{
		DamageBoostBuffs: newDamageBoostMap(),
	}
}

// //增害类型
// type DamageBoost struct {
// 	DamageBoostRate float32 `json:"DamageBoostRate"`
// 	DamageType      string  `json:"DamageType"`
// 	DamageBoostType DAMAGETYPE
// }

// //DamageBoost加法
// func (c *DamageBoost) add(toAdd DamageBoost) *DamageBoost {
// 	if c.DamageBoostType == toAdd.DamageBoostType {
// 		c.DamageBoostRate += toAdd.DamageBoostRate
// 	}
// 	return c
// }

// //DamageBoost减法
// func (c *DamageBoost) sub(toAdd DamageBoost) *DamageBoost {
// 	if c.DamageBoostType == toAdd.DamageBoostType {
// 		c.DamageBoostRate -= toAdd.DamageBoostRate
// 	}
// 	return c
// }

//增害类型列表
type DamageBoostMap map[DAMAGETYPE]float32

func newDamageBoostMap() (x DamageBoostMap) {
	for i := int(MinActiveDamageType); i <= int(MaxActiveDamageType); i++ {
		x[DAMAGETYPE(i)] = 0
	}
	return
}

//Add 加算
func (c DamageBoostMap) Add(toAdd ...DamageBoostMap) DamageBoostMap {
	for i := range c {
		for ii := range toAdd {
			c[i] += toAdd[ii][i]
		}
	}
	return c
}

//Sub 减算
func (c DamageBoostMap) Sub(toAdd ...DamageBoostMap) DamageBoostMap {
	for i := range c {
		for ii := range toAdd {
			c[i] += toAdd[ii][i]
		}
	}
	return c
}

//ForEach 循环执行
func (c DamageBoostMap) ForEach(f func(float32) float32) {
	for i := range c {
		c[i] = f(c[i])
	}
}

//人物信息
type Identity struct {
	Id    uint64 `json:"Id"`
	Name  string `json:"Name"`
	Level int    `json:"Level"`
}

//圣遗物类型
type ARTIFACTTYPE int

//圣遗物类型定义
const (
	Flower ARTIFACTTYPE = iota + 1
	Feather
	Hourglass
	Cup
	HeadWear
)

//最小圣遗物类型
const MinActiveAritfactType = Flower

//最大圣遗物类型
const MaxActiveAritfactType = HeadWear

//圣遗物位置可选词条
var selectableEntry [HeadWear + 1][]string

//GetMainPropertyList 圣遗物位置可选词条名列表
func (c ARTIFACTTYPE) GetMainPropertyList() []string {
	return selectableEntry[int(c)]
}

//圣遗物信息
type ArtifactIdentity struct {
	Location ARTIFACTTYPE `json:"Location"`
}
