package genshindata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

//资源URL
var downloadList map[string]string

//文件
var fileList map[string]fileInfo

//角色
var avatar map[uint64]*Avatar

//角色名ID映射表
var avatarNameMap map[string]uint64

//保存路径
const savePath = "./data"
const slash = "/"
const fileName = "avatar_map.json"

//文件完整路径
const fileFullPath = savePath + slash + fileName

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
	avatarExcelConfig        = "avatarExcelConfigData"
	avatarCurveExcelConfig   = "avatarCurveExcelConfigData"
	avatarPromoteExcelConfig = "avatarPromoteExcelConfigData"
	textMapFile              = "textMapData"
)

//级别
const (
	levelMin     = 1
	levelMax     = 90
	promotedMark = "+"
)

func init() {
	//下载URL初始化
	downloadList = map[string]string{
		avatarExcelConfig:        repositoryURL + avatarExcelConfigData,
		avatarCurveExcelConfig:   repositoryURL + avatarCurveExcelConfigData,
		avatarPromoteExcelConfig: repositoryURL + avatarPromoteExcelConfigData,
		textMapFile:              repositoryURL + textMapData,
	}
	//文件列表初始化
	fileList = map[string]fileInfo{
		"map_dir":    {path: savePath, class: dir},
		"avatar_map": {path: fileFullPath, class: js},
	}
	//角色对应初始化
	avatar = make(map[uint64]*Avatar)
	avatarNameMap = make(map[string]uint64)

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
		fmt.Println("[获取最新数据]")
		err = getDataFromRepository()
	} else {
		//读取
		fmt.Println("[读取已经有数据]")
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
	avatarBaseDataList := make(avatarBaseListData, 0)
	growCurvesDataList := make(growCurvesListData, 0)
	promoteDataList := make(promoteListData, 0)
	textMap := make(map[uint64]string)
	for i, v := range content {
		switch i {
		case avatarExcelConfig:
			json.Unmarshal(v.Bytes(), &avatarBaseDataList)
		case avatarCurveExcelConfig:
			json.Unmarshal(v.Bytes(), &growCurvesDataList)
		case avatarPromoteExcelConfig:
			json.Unmarshal(v.Bytes(), &promoteDataList)
		case textMapFile:
			json.Unmarshal(v.Bytes(), &textMap)
		}
	}
	//数据处理
	growCurvesDataMap := make(map[int]*growCurvesData)
	promoteDataMap := make(map[uint64][]*promoteData)
	for i := range growCurvesDataList {
		growCurvesDataMap[growCurvesDataList[i].Level] = &growCurvesDataList[i]
	}
	CurvesIndexMap := make(map[string]int)
	for i, v := range growCurvesDataMap[1].CurveInfos {
		CurvesIndexMap[v.Type] = i
	}
	for i := range promoteDataList {
		temp, ok := promoteDataMap[promoteDataList[i].AvatarPromoteId]
		if !ok {
			temp = make([]*promoteData, 0)
		}
		promoteDataMap[promoteDataList[i].AvatarPromoteId] = append(temp, &promoteDataList[i])
	}
	//计算
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
		}
		//级别曲线参数
		var hpTypeIndex int
		var attackTypeIndex int
		var defenseTypeIndex int
		for _, vv := range currentAvatarData.PropGrowCurves {
			switch vv.Type {
			case HP:
				hpTypeIndex = CurvesIndexMap[vv.Value]
			case ATTACK:
				attackTypeIndex = CurvesIndexMap[vv.Value]
			case DEFENSE:
				defenseTypeIndex = CurvesIndexMap[vv.Value]
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
			currentProperty.Hp *= growCurvesDataMap[ii].CurveInfos[hpTypeIndex].Value
			currentProperty.Attack *= growCurvesDataMap[ii].CurveInfos[attackTypeIndex].Value
			currentProperty.Defense *= growCurvesDataMap[ii].CurveInfos[defenseTypeIndex].Value

			avatar[currentAvatarData.Id].LevelMap[strconv.Itoa(ii)] = currentProperty
		}
		//突破参数
		list := promoteDataMap[currentAvatarData.AvatarPromoteId]
		var currentProperty *Property
		var addPropNames []string = make([]string, 0)
		for _, vv := range list[0].AddProps {
			tempName := ""
			switch vv.PropType {
			case HP:
				tempName = "Hp"
			case ATTACK:
				tempName = "Attack"
			case DEFENSE:
				tempName = "Defense"
			case HP_PERCENT:
				tempName = "Hp_percent"
			case ATTACK_PERCENT:
				tempName = "Attack_percent"
			case DEFENSE_PERCENT:
				tempName = "Defense_percent"
			case CRITICAL:
				tempName = "Critical"
			case CRITICAL_HURT:
				tempName = "CriticalHurt"
			case ICE:
				tempName = "Ice"
			case WIND:
				tempName = "Wind"
			case PHYSICAL:
				tempName = "Physical"
			case ELEC:
				tempName = "Elec"
			case ROCK:
				tempName = "Rock"
			case FIRE:
				tempName = "Fire"
			case WATER:
				tempName = "Water"
			case CHANGE:
				tempName = "ChargeEfficiency"
			case ELEMENT_MASTER:
				tempName = "ElementMaster"
			case HEAL:
				tempName = "HealActiveUp"

			}
			addPropNames = append(addPropNames, tempName)
		}

		for ii := len(list) - 1; ii > 0; ii-- {
			currentPromote := list[ii]
			requiredLevel := list[ii-1].UnlockMaxLevel
			unlockMaxLevel := currentPromote.UnlockMaxLevel

			for iii := unlockMaxLevel; iii >= requiredLevel; iii-- {
				currentProperty = avatar[currentAvatarData.Id].LevelMap[strconv.Itoa(iii)]
				//突破界限
				if iii == requiredLevel {
					newCurrentProperty := &Property{}
					avatar[currentAvatarData.Id].LevelMap[strconv.Itoa(iii)+promotedMark] = newCurrentProperty
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

//结果保存至本地
func saveResult() error {
	content, err := json.Marshal(avatar)
	if err != nil {
		return err
	}
	return writeToFile(fileFullPath, bytes.NewBuffer(content))
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
	content := bytes.NewBuffer(make([]byte, 0, defaultBuffSize))
	err := readFromFile(fileFullPath, content)
	if err != nil {
		return err
	}
	return json.Unmarshal(content.Bytes(), &avatar)
}
