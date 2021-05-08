package calc

//废弃

// type InputData struct {
// 	Kyara     *Character `json:"Kyara"`
// 	Buki      *Weapon    `json:"Buki"`
// 	Seiibutsu *Artifacts `json:"Seiibutsu"`
// 	Monsuta   *Monster   `json:"Monsuta"`
// 	Hoka      *Other     `json:"Hoka"`
// }

// func NewInputData() *InputData {
// 	return &InputData{
// 		Kyara:     NewCharacter(),
// 		Buki:      NewWeapon(),
// 		Seiibutsu: NewArtifacts(),
// 		Monsuta:   NewMonster(),
// 		Hoka:      NewOther(),
// 	}
// }

// //人物属性
// type Character struct {
// 	Identity
// 	StateBase
// 	StateBuff
// }

// func NewCharacter() *Character {
// 	return &Character{
// 		StateBase: newStateBase(),
// 		StateBuff: newStateBuff(),
// 	}
// }

// //武器属性
// type Weapon struct {
// 	Identity
// 	StateBase
// 	StateBuff
// }

// func NewWeapon() *Weapon {
// 	return &Weapon{
// 		StateBase: newStateBase(),
// 		StateBuff: newStateBuff(),
// 	}
// }

// //圣遗物(全)
// type Artifacts struct {
// 	ArtifactList
// 	ArtifactsEffect
// }

// func NewArtifacts() *Artifacts {
// 	return &Artifacts{
// 		ArtifactsEffect: *NewArtifactsEffect(),
// 	}
// }

// //圣遗物套装效果
// type ArtifactsEffect struct {
// 	StateBase
// 	StateBuff
// }

// func NewArtifactsEffect() *ArtifactsEffect {
// 	return &ArtifactsEffect{
// 		StateBase: newStateBase(),
// 		StateBuff: newStateBuff(),
// 	}
// }

// //圣遗物列表
// type ArtifactList []Artifact

// func NewArtifactList() (x ArtifactList) {
// 	x = make(ArtifactList, 6, 6)
// 	for i := MinActiveAritfactType; i <= MaxActiveAritfactType; i++ {
// 		x[int(i)] = *NewArtifact(i)
// 	}
// 	return
// }

// //圣遗物
// type Artifact struct {
// 	ArtifactIdentity
// 	StateBase
// 	StateBuff
// }

// func NewArtifact(i ARTIFACTTYPE) *Artifact {
// 	return &Artifact{
// 		ArtifactIdentity: ArtifactIdentity{Location: i},
// 		StateBase:        newStateBase(),
// 		StateBuff:        newStateBuff(),
// 	}
// }

// //怪物属性
// type Monster struct {
// 	Identity
// 	//怪物自身抗性
// 	EleResisRates DamageBoostMap `json:"EleResisRates"`
// 	//怪物被减抗
// 	EleDecreaseRates   DamageBoostMap `json:"EleDecreaseRates"`
// 	FinalEleResisRates DamageBoostMap `json:"FinalEleResisRates"`
// 	DefDeBuffRate      float32        `json:"DefDeBuffRate"`
// }

// func NewMonster() *Monster {
// 	return &Monster{
// 		EleResisRates:      newDamageBoostMap(),
// 		EleDecreaseRates:   newDamageBoostMap(),
// 		FinalEleResisRates: newDamageBoostMap(),
// 	}
// }

// //其他加成
// type Other struct {
// 	StateBase
// 	StateBuff
// }

// func NewOther() *Other {
// 	return &Other{
// 		StateBase: newStateBase(),
// 		StateBuff: newStateBuff(),
// 	}
// }
