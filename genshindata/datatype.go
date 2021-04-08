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
}
type avatarBaseListData []avatarBaseData

//成长参数
type propGrowCurves struct {
	Type  string `json:"Type"`
	Value string `json:"GrowCurve"`
}

//人物基础升级提升值
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

//人物突破提升值
type promoteData struct {
	AvatarPromoteId     uint64    `json:"AvatarPromoteId"`
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
