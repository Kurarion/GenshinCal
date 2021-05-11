package genshindata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"regexp"
)

//资源URL
var downloadList map[string]string

//文件
var fileList map[string]fileInfo

//角色
var avatar map[uint64]*Avatar

//角色名ID映射表
var avatarNameMap map[string]uint64

//武器
var weapon map[uint64]*Weapon

//武器名ID映射表
var weaponNameMap map[string]uint64

//圣遗物词条刻度
var reliquaryAffixMap map[string][]float64

//圣遗物主词条值
var reliquaryMainMap map[string]float64

//怪物
var monster map[uint64]*Monster

//怪物名ID映射表
var monsterNameMap map[string]uint64

//人物技能集
var avatarSkills map[uint64]*AvatarSkills

//序列化列表
var saveObjList map[string]interface{}

//正则
var regx []*regexp.Regexp
var regxReplaceList []string

//保存路径
const savePath = "./data"
const slash = "/"
const avatarFileName = "avatar_map.json"
const weaponFileName = "weapon_map.json"
const reliquaryAffixFileName = "reliquary_affix_map.json"
const reliquaryMainFileName = "reliquary_main_map.json"
const avatarSkillsFileName = "avatar_skills_map.json"
const monsterFileName = "monster_map.json"

//文件完整路径
const avatarFileFullPath = savePath + slash + avatarFileName
const weaponFileFullPath = savePath + slash + weaponFileName
const reliquaryAffixFileFullPath = savePath + slash + reliquaryAffixFileName
const reliquaryMainFileFullPath = savePath + slash + reliquaryMainFileName
const avatarSkillsFileFullPath = savePath + slash + avatarSkillsFileName
const monsterFileFullPath = savePath + slash + monsterFileName

//圣遗物词条depotID
const artiDepotId = 501

//圣遗物等级
const artiLeveL = 21

//武器最低星级
const minWeaponRankLevel = 3

//正则
const regexColorToFront = `<color`
const regexColorToFrontReplaced = `<font color`
const regexColorToFrontSalsh = `</color`
const regexColorToFrontSalshReplaced = `</font`

//文件定义
type fileType int

const (
	dir fileType = iota
	js
)

type fileInfo struct {
	path  string
	class fileType
}

//默认Buff大小
const defaultBuffSize = 15000

//名
const (
	avatarExcelConfig           = "avatarExcelConfigData"
	avatarCurveExcelConfig      = "avatarCurveExcelConfigData"
	avatarPromoteExcelConfig    = "avatarPromoteExcelConfigData"
	weaponExcelConfig           = "weaponExcelConfigData"
	weaponCurveExcelConfig      = "weaponCurveExcelConfigData"
	weaponPromoteExcelConfig    = "weaponPromoteExcelConfigData"
	equipAffixExcelConfig       = "EquipAffixExcelConfigData"
	reliquaryAffixExcelConfig   = "ReliquaryAffixExcelConfigData"
	reliquaryLevelExcelConfig   = "ReliquaryLevelExcelConfigData"
	monsterExcelConfig          = "MonsterExcelConfigData"
	monsterCurveExcelConfig     = "MonsterCurveExcelConfigData"
	AvatarSkillDepotExcelConfig = "AvatarSkillDepotExcelConfigData"
	AvatarSkillExcelConfig      = "AvatarSkillExcelConfigData"
	ProudSkillExcelConfig       = "ProudSkillExcelConfigData"
	AvatarTalentExcelConfig     = "AvatarTalentExcelConfigData"
	textMapFile                 = "textMapData"
)

//级别
const (
	levelMin        = 1
	levelMax        = 90
	promotedMark    = "+"
	MonsterLevelMin = 1
	MonsterLevelMax = 100
)

func init() {
	//下载URL初始化
	downloadList = map[string]string{
		avatarExcelConfig:           repositoryURL + avatarExcelConfigData,
		avatarCurveExcelConfig:      repositoryURL + avatarCurveExcelConfigData,
		avatarPromoteExcelConfig:    repositoryURL + avatarPromoteExcelConfigData,
		weaponExcelConfig:           repositoryURL + weaponExcelConfigData,
		weaponCurveExcelConfig:      repositoryURL + weaponCurveExcelConfigData,
		weaponPromoteExcelConfig:    repositoryURL + weaponPromoteExcelConfigData,
		equipAffixExcelConfig:       repositoryURL + EquipAffixExcelConfigData,
		reliquaryAffixExcelConfig:   repositoryURL + ReliquaryAffixExcelConfigData,
		reliquaryLevelExcelConfig:   repositoryURL + ReliquaryLevelExcelConfigData,
		monsterExcelConfig:          repositoryURL + MonsterExcelConfigData,
		monsterCurveExcelConfig:     repositoryURL + MonsterCurveExcelConfigData,
		AvatarSkillDepotExcelConfig: repositoryURL + AvatarSkillDepotExcelConfigData,
		AvatarSkillExcelConfig:      repositoryURL + AvatarSkillExcelConfigData,
		ProudSkillExcelConfig:       repositoryURL + ProudSkillExcelConfigData,
		AvatarTalentExcelConfig:     repositoryURL + AvatarTalentExcelConfigData,
		textMapFile:                 repositoryURL + textMapData,
	}
	//文件列表初始化
	fileList = map[string]fileInfo{
		"map_dir":             {path: savePath, class: dir},
		"avatar_map":          {path: avatarFileFullPath, class: js},
		"weapon_map":          {path: weaponFileFullPath, class: js},
		"reliquary_affix_map": {path: reliquaryAffixFileFullPath, class: js},
		"reliquary_main_map":  {path: reliquaryMainFileFullPath, class: js},
		"monster_map":         {path: monsterFileFullPath, class: js},
		"avatar_skills_map":   {path: avatarSkillsFileFullPath, class: js},
	}
	saveObjList = map[string]interface{}{
		avatarFileFullPath:         &avatar,
		weaponFileFullPath:         &weapon,
		reliquaryAffixFileFullPath: &reliquaryAffixMap,
		reliquaryMainFileFullPath:  &reliquaryMainMap,
		monsterFileFullPath:        &monster,
		avatarSkillsFileFullPath:   &avatarSkills,
	}
	//角色对应初始化
	avatar = make(map[uint64]*Avatar)
	avatarNameMap = make(map[string]uint64)

	//武器对应初始化
	weapon = make(map[uint64]*Weapon)
	weaponNameMap = make(map[string]uint64)

	//圣遗物对应初始化
	//圣遗物词条刻度
	reliquaryAffixMap = make(map[string][]float64)
	//圣遗物主词条值
	reliquaryMainMap = make(map[string]float64)

	//怪物对应初始化
	monster = make(map[uint64]*Monster)
	monsterNameMap = make(map[string]uint64)

	//人物技能对应初始化
	avatarSkills = make(map[uint64]*AvatarSkills)

	//正则
	regx = make([]*regexp.Regexp, 2, 2)
	regx[0] = regexp.MustCompile(regexColorToFront)
	regx[1] = regexp.MustCompile(regexColorToFrontSalsh)
	regxReplaceList = make([]string, 2, 2)
	regxReplaceList[0] = regexColorToFrontReplaced
	regxReplaceList[1] = regexColorToFrontSalshReplaced

	//初始化
	initialize(false)
}

func getJSON(url string) (buf *bytes.Buffer, err error) {
	rp, err := http.Get(url)
	if err != nil {
		return
	}
	defer rp.Body.Close()
	return readBody(rp)
}

func readBody(rp *http.Response) (buf *bytes.Buffer, err error) {
	defer func() {
		r := recover()
		if r == nil {
			return
		} else if er, ok := r.(error); ok {
			err = er
			return
		}
		panic(r)
	}()
	buf = bytes.NewBuffer(make([]byte, 0, defaultBuffSize))
	buf.ReadFrom(rp.Body)
	return buf, nil
}

func initialize(forceUpdate bool) (err error) {
	//检查本地是否已有数据
	hasLocalData := true
	//检查是否存在
	for _, v := range fileList {
		fi, er := os.Stat(v.path)
		if er != nil {
			if os.IsNotExist(er) {
				hasLocalData = false
				if v.class == dir {
					os.MkdirAll(v.path, 0666)
				}
				break
			}
		}
		//文件大小检查
		if fi.Size() <= 0 && v.class == js {
			hasLocalData = false
			break
		}
	}
	if forceUpdate || !hasLocalData {
		//更新
		fmt.Println("Getting the latest data...")
		err = getDataFromRepository()
	} else {
		//读取
		fmt.Println("Reading existing data...")
		err = readMapFormLocal()
	}

	if err != nil {
		return
	}

	//角色名ID映射
	for i, v := range avatar {
		temp, ok := avatarNameMap[v.Name]
		if !ok {
			avatarNameMap[v.Name] = i
		} else {
			if len(v.Desc) > len(avatar[temp].Desc) {
				avatarNameMap[v.Name] = i
			}
		}
	}

	//武器名ID映射
	for i, v := range weapon {
		temp, ok := weaponNameMap[v.Name]
		if !ok {
			weaponNameMap[v.Name] = i
		} else {
			if len(v.Desc) > len(weapon[temp].Desc) {
				weaponNameMap[v.Name] = i
			}
		}
	}

	//怪物名ID映射
	for i, v := range monster {
		_, ok := monsterNameMap[v.Name]
		if !ok {
			monsterNameMap[v.Name] = i
		}
	}

	return
}

//写入文件
func writeToFile(path string, content *bytes.Buffer) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = content.WriteTo(f)
	return err
}

//读取文件
func readFromFile(fileName string, content *bytes.Buffer) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = content.ReadFrom(f)
	return err
}

//更新
func update() error {
	content := make(map[string]*bytes.Buffer)
	for i, v := range downloadList {
		temp, err := getJSON(v)
		if err != nil {
			//error
			return err
		}
		content[i] = temp
	}
	//解析
	//人物
	avatarBaseDataList := make(avatarBaseListData, 0)
	avatarGrowCurvesDataList := make(growCurvesListData, 0)
	avatarPromoteDataList := make(promoteListData, 0)
	//武器
	weaponBaseDataList := make(weaponBaseListData, 0)
	weaponGrowCurvesDataList := make(growCurvesListData, 0)
	weaponPromoteDataList := make(promoteListData, 0)
	weaponSkillAffixDataList := make(skillAffixListData, 0)
	//圣遗物
	reliquaryAffixDataList := make(reliquaryAffixListData, 0)
	reliquaryMainDataList := make(reliquaryMainListData, 0)
	//怪物
	monsterBaseDataList := make(monsterBaseListData, 0)
	monsterGrowCurvesDataList := make(growCurvesListData, 0)
	//人物技能
	avatarSkillsDataList := make(avatarSkillsListData, 0)
	avatarSkillDataList := make(avatarSkillListData, 0)
	avatarProudSkillDataList := make(avatarProudSkillListData, 0)
	avatarTalentDataList := make(avatarTalentListData, 0)
	//名
	textMap := make(map[uint64]string)
	for i, v := range content {
		switch i {
		case avatarExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarBaseDataList)
		case avatarCurveExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarGrowCurvesDataList)
		case avatarPromoteExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarPromoteDataList)
		case weaponExcelConfig:
			json.Unmarshal(v.Bytes(), &weaponBaseDataList)
		case weaponCurveExcelConfig:
			json.Unmarshal(v.Bytes(), &weaponGrowCurvesDataList)
		case weaponPromoteExcelConfig:
			json.Unmarshal(v.Bytes(), &weaponPromoteDataList)
		case equipAffixExcelConfig:
			json.Unmarshal(v.Bytes(), &weaponSkillAffixDataList)
		case reliquaryAffixExcelConfig:
			json.Unmarshal(v.Bytes(), &reliquaryAffixDataList)
		case reliquaryLevelExcelConfig:
			json.Unmarshal(v.Bytes(), &reliquaryMainDataList)
		case monsterExcelConfig:
			json.Unmarshal(v.Bytes(), &monsterBaseDataList)
		case monsterCurveExcelConfig:
			json.Unmarshal(v.Bytes(), &monsterGrowCurvesDataList)
		case AvatarSkillDepotExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarSkillsDataList)
		case AvatarSkillExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarSkillDataList)
		case ProudSkillExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarProudSkillDataList)
		case AvatarTalentExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarTalentDataList)
		case textMapFile:
			json.Unmarshal(v.Bytes(), &textMap)
		}
	}
	//数据处理
	textMap[0] = ""
	//人物
	avatarGrowCurvesDataMap := make(map[int]*growCurvesData)
	avatarPromoteDataMap := make(map[uint64][]*promoteData)
	for i := range avatarGrowCurvesDataList {
		avatarGrowCurvesDataMap[avatarGrowCurvesDataList[i].Level] = &avatarGrowCurvesDataList[i]
	}
	avatarCurvesIndexMap := make(map[string]int)
	for i, v := range avatarGrowCurvesDataMap[1].CurveInfos {
		avatarCurvesIndexMap[v.Type] = i
	}
	for i := range avatarPromoteDataList {
		temp, ok := avatarPromoteDataMap[avatarPromoteDataList[i].AvatarPromoteId]
		if !ok {
			temp = make([]*promoteData, 0)
		}
		avatarPromoteDataMap[avatarPromoteDataList[i].AvatarPromoteId] = append(temp, &avatarPromoteDataList[i])
	}
	//武器
	weaponGrowCurvesDataMap := make(map[int]*growCurvesData)
	weaponPromoteDataMap := make(map[uint64][]*promoteData)
	for i := range weaponGrowCurvesDataList {
		weaponGrowCurvesDataMap[weaponGrowCurvesDataList[i].Level] = &weaponGrowCurvesDataList[i]
	}
	weaponCurvesIndexMap := make(map[string]int)
	for i, v := range weaponGrowCurvesDataMap[1].CurveInfos {
		weaponCurvesIndexMap[v.Type] = i
	}
	for i := range weaponPromoteDataList {
		temp, ok := weaponPromoteDataMap[weaponPromoteDataList[i].WeaponPromoteId]
		if !ok {
			temp = make([]*promoteData, 0)
		}
		weaponPromoteDataMap[weaponPromoteDataList[i].WeaponPromoteId] = append(temp, &weaponPromoteDataList[i])
	}
	//特效
	weaponSkillAffixDataMap := make(map[uint64]map[int]skillAffix)
	for i := range weaponSkillAffixDataList {
		_, ok := weaponSkillAffixDataMap[weaponSkillAffixDataList[i].Id]
		if !ok {
			weaponSkillAffixDataMap[weaponSkillAffixDataList[i].Id] = make(map[int]skillAffix)
		}
		weaponSkillAffixDataMap[weaponSkillAffixDataList[i].Id][weaponSkillAffixDataList[i].Level+1] =
			skillAffix{
				skillAffixData: weaponSkillAffixDataList[i],
				Name:           textMap[weaponSkillAffixDataList[i].NameTextMapHash],
				Desc:           htmlColorTag(textMap[weaponSkillAffixDataList[i].DescTextMapHash]),
			}
	}
	//圣遗物
	for i := range reliquaryAffixDataList {
		if reliquaryAffixDataList[i].DepotId != artiDepotId {
			continue
		}
		temp, have := reliquaryAffixMap[reliquaryAffixDataList[i].PropType]
		if !have {
			temp = make([]float64, 0)
		}
		reliquaryAffixMap[reliquaryAffixDataList[i].PropType] = append(temp, reliquaryAffixDataList[i].PropValue)
	}
	for i := len(reliquaryMainDataList) - 1; i >= 0; i-- {
		if reliquaryMainDataList[i].Level == artiLeveL {
			for ii := range reliquaryMainDataList[i].AddProps {
				reliquaryMainMap[reliquaryMainDataList[i].AddProps[ii].PropType] = reliquaryMainDataList[i].AddProps[ii].Value
			}
			break
		}
	}
	//怪物
	monsterGrowCurvesDataMap := make(map[int]*growCurvesData)
	for i := range monsterGrowCurvesDataList {
		monsterGrowCurvesDataMap[monsterGrowCurvesDataList[i].Level] = &monsterGrowCurvesDataList[i]
	}
	monsterCurvesIndexMap := make(map[string]int)
	for i, v := range monsterGrowCurvesDataMap[1].CurveInfos {
		monsterCurvesIndexMap[v.Type] = i
	}
	//人物技能
	avatarSkillsDataMap := make(map[uint64]*avatarSkillsData)
	avatarSkillDataMap := make(map[uint64]*avatarSkillData)
	avatarProudSkillDataMap := make(map[uint64]*avatarProudSkillData)
	avatarTalentDataMap := make(map[uint64]*avatarTalentData)
	for i := range avatarSkillsDataList {
		avatarSkillsDataMap[avatarSkillsDataList[i].Id] = &avatarSkillsDataList[i]
	}
	for i := range avatarSkillDataList {
		avatarSkillDataMap[avatarSkillDataList[i].Id] = &avatarSkillDataList[i]
	}
	for i := range avatarProudSkillDataList {
		avatarProudSkillDataMap[avatarProudSkillDataList[i].ProudSkillGroupId] = &avatarProudSkillDataList[i]
	}
	for i := range avatarTalentDataList {
		avatarTalentDataMap[avatarTalentDataList[i].TalentId] = &avatarTalentDataList[i]
	}
	for i := range avatarSkillsDataMap {
		temp := avatarSkillsDataMap[i]
		if temp.Skills[1] == 0 {
			continue
		}
		avatarSkills[temp.Id] = &AvatarSkills{
			Id: temp.Id,
			ESkill: AvatarSkillInfo{
				Name: textMap[avatarSkillDataMap[temp.Skills[1]].NameTextMapHash],
				Desc: htmlColorTag(textMap[avatarSkillDataMap[temp.Skills[1]].DescTextMapHash]),
			},
			QSkill: AvatarSkillInfo{
				Name: textMap[avatarSkillDataMap[temp.EnergySkill].NameTextMapHash],
				Desc: htmlColorTag(textMap[avatarSkillDataMap[temp.EnergySkill].DescTextMapHash]),
			},
		}
		for ii := range temp.InherentProudSkillOpens {
			temp2 := temp.InherentProudSkillOpens[ii]
			if temp2.ProudSkillGroupId == 0 {
				continue
			}
			avatarSkills[temp.Id].ProudSkills = append(avatarSkills[temp.Id].ProudSkills, AvatarSkillInfo{
				Name: textMap[avatarProudSkillDataMap[temp2.ProudSkillGroupId].NameTextMapHash],
				Desc: htmlColorTag(textMap[avatarProudSkillDataMap[temp2.ProudSkillGroupId].DescTextMapHash]),
			})
		}
		for ii := range temp.Talents {
			temp2 := temp.Talents[ii]
			if temp2 == 0 {
				continue
			}
			avatarSkills[temp.Id].Talents = append(avatarSkills[temp.Id].Talents, AvatarSkillInfo{
				Name: textMap[avatarTalentDataMap[temp2].NameTextMapHash],
				Desc: htmlColorTag(textMap[avatarTalentDataMap[temp2].DescTextMapHash]),
			})
		}
	}
	//计算
	//人物
	for i := range avatarBaseDataList {
		currentAvatarData := &avatarBaseDataList[i]
		avatar[currentAvatarData.Id] = &Avatar{
			Id:              currentAvatarData.Id,
			Name:            textMap[currentAvatarData.NameTextMapHash],
			NameTextMapHash: currentAvatarData.NameTextMapHash,
			Desc:            textMap[currentAvatarData.DescTextMapHash],
			DescTextMapHash: currentAvatarData.DescTextMapHash,
			IconName:        currentAvatarData.IconName,
			WeaponType:      currentAvatarData.WeaponType,
			LevelMap:        make(map[string]*Property),
			SkillDepotId:    currentAvatarData.SkillDepotId,
		}
		//级别曲线参数
		var hpTypeIndex int
		var attackTypeIndex int
		var defenseTypeIndex int
		for _, vv := range currentAvatarData.PropGrowCurves {
			switch vv.Type {
			case HP:
				hpTypeIndex = avatarCurvesIndexMap[vv.Value]
			case ATTACK:
				attackTypeIndex = avatarCurvesIndexMap[vv.Value]
			case DEFENSE:
				defenseTypeIndex = avatarCurvesIndexMap[vv.Value]
			}
		}
		for ii := levelMin; ii <= levelMax; ii++ {
			currentProperty := &Property{
				Hp:               currentAvatarData.HpBase,
				Attack:           currentAvatarData.AttackBase,
				Defense:          currentAvatarData.DefenseBase,
				Critical:         currentAvatarData.Critical,
				CriticalHurt:     currentAvatarData.CriticalHurt,
				ChargeEfficiency: currentAvatarData.ChargeEfficiency,
			}
			currentProperty.Level = ii
			//此等级数值
			currentProperty.Hp *= avatarGrowCurvesDataMap[ii].CurveInfos[hpTypeIndex].Value
			currentProperty.Attack *= avatarGrowCurvesDataMap[ii].CurveInfos[attackTypeIndex].Value
			currentProperty.Defense *= avatarGrowCurvesDataMap[ii].CurveInfos[defenseTypeIndex].Value

			avatar[currentAvatarData.Id].LevelMap[fmt.Sprintf("%02d", ii)] = currentProperty
		}
		//突破参数
		list := avatarPromoteDataMap[currentAvatarData.AvatarPromoteId]
		var currentProperty *Property
		var addPropNames []string = make([]string, 0)
		for _, vv := range list[0].AddProps {
			tempName := GetNameFromTypeCode(vv.PropType)
			addPropNames = append(addPropNames, tempName)
		}

		for ii := len(list) - 1; ii > 0; ii-- {
			currentPromote := list[ii]
			requiredLevel := list[ii-1].UnlockMaxLevel
			unlockMaxLevel := currentPromote.UnlockMaxLevel

			for iii := unlockMaxLevel; iii >= requiredLevel; iii-- {
				currentProperty = avatar[currentAvatarData.Id].LevelMap[fmt.Sprintf("%02d", iii)]
				//突破界限
				if iii == requiredLevel {
					newCurrentProperty := &Property{}
					avatar[currentAvatarData.Id].LevelMap[fmt.Sprintf("%02d", iii)+promotedMark] = newCurrentProperty
					copyStruct(newCurrentProperty, currentProperty)
					currentProperty = newCurrentProperty
				}
				currentProperty.PromoteLevel += float64(currentPromote.PromoteLevel)
				temp := reflect.ValueOf(currentProperty).Elem()
				for iiii := range addPropNames {
					temp.FieldByName(addPropNames[iiii]).SetFloat(temp.FieldByName(addPropNames[iiii]).Float() + currentPromote.AddProps[iiii].Value)
				}
			}
		}
	}
	//武器
	for i := range weaponBaseDataList {
		currentWeaponData := &weaponBaseDataList[i]
		if currentWeaponData.RankLevel < minWeaponRankLevel {
			continue
		}
		weapon[currentWeaponData.Id] = &Weapon{
			Id:              currentWeaponData.Id,
			Name:            textMap[currentWeaponData.NameTextMapHash],
			NameTextMapHash: currentWeaponData.NameTextMapHash,
			Desc:            textMap[currentWeaponData.DescTextMapHash],
			DescTextMapHash: currentWeaponData.DescTextMapHash,
			IconName:        currentWeaponData.IconName,
			WeaponType:      currentWeaponData.WeaponType,
			SkillAffixMap:   weaponSkillAffixDataMap[currentWeaponData.SkillAffix[0]],
			LevelMap:        make(map[string]*Property),
		}
		//级别曲线参数
		var weaponBaseAtkIndex int
		var weaponSubAffixIndex int
		var weaponSubAffixName string
		for _, vv := range currentWeaponData.PropGrowCurves {
			switch vv.PropType {
			case ATTACK:
				weaponBaseAtkIndex = weaponCurvesIndexMap[vv.Type]
			default:
				weaponSubAffixName = GetNameFromTypeCode(vv.PropType)
				weaponSubAffixIndex = weaponCurvesIndexMap[vv.Type]
			}
		}

		for ii := levelMin; ii <= levelMax; ii++ {
			currentProperty := &Property{
				Attack: currentWeaponData.PropGrowCurves[0].InitValue,
			}
			currentProperty.Level = ii
			//此等级数值
			currentProperty.Attack *= weaponGrowCurvesDataMap[ii].CurveInfos[weaponBaseAtkIndex].Value
			if len(weaponSubAffixName) != 0 {
				temp := reflect.ValueOf(currentProperty).Elem()
				temp.FieldByName(weaponSubAffixName).SetFloat(currentWeaponData.PropGrowCurves[1].InitValue * weaponGrowCurvesDataMap[ii].CurveInfos[weaponSubAffixIndex].Value)
			}

			weapon[currentWeaponData.Id].LevelMap[fmt.Sprintf("%02d", ii)] = currentProperty
		}
		//突破参数
		list := weaponPromoteDataMap[currentWeaponData.WeaponPromoteId]
		var currentProperty *Property
		var addPropNames []string = make([]string, 0)
		for _, vv := range list[0].AddProps {
			tempName := GetNameFromTypeCode(vv.PropType)
			addPropNames = append(addPropNames, tempName)
		}

		for ii := len(list) - 1; ii > 0; ii-- {
			currentPromote := list[ii]
			requiredLevel := list[ii-1].UnlockMaxLevel
			unlockMaxLevel := currentPromote.UnlockMaxLevel

			for iii := unlockMaxLevel; iii >= requiredLevel; iii-- {
				currentProperty = weapon[currentWeaponData.Id].LevelMap[fmt.Sprintf("%02d", iii)]
				//突破界限
				if iii == requiredLevel {
					newCurrentProperty := &Property{}
					weapon[currentWeaponData.Id].LevelMap[fmt.Sprintf("%02d", iii)+promotedMark] = newCurrentProperty
					copyStruct(newCurrentProperty, currentProperty)
					currentProperty = newCurrentProperty
				}
				currentProperty.PromoteLevel += float64(currentPromote.PromoteLevel)
				temp := reflect.ValueOf(currentProperty).Elem()
				for iiii := range addPropNames {
					temp.FieldByName(addPropNames[iiii]).SetFloat(temp.FieldByName(addPropNames[iiii]).Float() + currentPromote.AddProps[iiii].Value)
				}
			}
		}
	}
	//怪物
	for i := range monsterBaseDataList {
		currentMonsterData := &monsterBaseDataList[i]
		monster[currentMonsterData.Id] = &Monster{
			Id:              currentMonsterData.Id,
			Name:            textMap[currentMonsterData.NameTextMapHash],
			NameTextMapHash: currentMonsterData.NameTextMapHash,
			MonsterName:     currentMonsterData.MonsterName,
			Type:            currentMonsterData.Type,
			LevelMap:        make(map[string]*MonsterProperty),
		}
		//级别曲线参数
		var hpTypeIndex int
		var attackTypeIndex int
		var defenseTypeIndex int
		for _, vv := range currentMonsterData.PropGrowCurves {
			switch vv.Type {
			case HP:
				hpTypeIndex = monsterCurvesIndexMap[vv.Value]
			case ATTACK:
				attackTypeIndex = monsterCurvesIndexMap[vv.Value]
			case DEFENSE:
				defenseTypeIndex = monsterCurvesIndexMap[vv.Value]
			}
		}
		for ii := MonsterLevelMin; ii <= MonsterLevelMax; ii++ {
			currentProperty := &MonsterProperty{
				Hp:          currentMonsterData.HpBase,
				Attack:      currentMonsterData.AttackBase,
				Defense:     currentMonsterData.DefenseBase,
				subHurtData: currentMonsterData.subHurtData,
			}
			currentProperty.Level = ii
			//此等级数值
			currentProperty.Hp *= monsterGrowCurvesDataMap[ii].CurveInfos[hpTypeIndex].Value
			currentProperty.Attack *= monsterGrowCurvesDataMap[ii].CurveInfos[attackTypeIndex].Value
			currentProperty.Defense *= monsterGrowCurvesDataMap[ii].CurveInfos[defenseTypeIndex].Value

			monster[currentMonsterData.Id].LevelMap[fmt.Sprintf("%03d", ii)] = currentProperty
		}
	}
	return nil
}

func copyStruct(dst, src interface{}) {
	tempA := reflect.ValueOf(dst).Elem()
	tempB := reflect.ValueOf(src).Elem()
	for i := 0; i < tempA.NumField(); i++ {
		name := tempB.Type().Field(i).Name
		value := tempB.FieldByName(name)
		tempA.FieldByName(name).Set(value)
	}
}

func htmlColorTag(origin string) string {
	var res = origin
	for i := range regx {
		res = regx[i].ReplaceAllString(res, regxReplaceList[i])
	}
	return res
}

//结果保存至本地
func saveResult() error {
	for i := range saveObjList {
		content, err := json.Marshal(saveObjList[i])
		if err != nil {
			return err
		}
		err = writeToFile(i, bytes.NewBuffer(content))
		if err != nil {
			return err
		}
	}
	return nil
}

//获取最新数据
func getDataFromRepository() error {
	err := update()
	if err != nil {
		return err
	}
	return saveResult()
}

//从本地读取已往结果
func readMapFormLocal() error {
	for i := range saveObjList {
		content := bytes.NewBuffer(make([]byte, 0, defaultBuffSize))
		err := readFromFile(i, content)
		if err != nil {
			return err
		}
		err = json.Unmarshal(content.Bytes(), saveObjList[i])
		if err != nil {
			return err
		}
	}
	return nil
}
