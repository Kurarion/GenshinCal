# GenshinData
```go
//人物定义
AvatarExcelConfigData = GenshinData\ExcelBinOutput\AvatarExcelConfigData.json
//人物基础升级提升值(累加)
AvatarCurveExcelConfigData = GenshinData\ExcelBinOutput\AvatarCurveExcelConfigData.json
//人物突破提升值(对应)
AvatarPromoteExcelConfigData = GenshinData\ExcelBinOutput\AvatarPromoteExcelConfigData.json
//人物技能集定义
AvatarSkillDepotExcelConfigData = GenshinData\ExcelBinOutput\AvatarSkillDepotExcelConfigData.json
//技能详细定义
AvatarSkillExcelConfigData = GenshinData\ExcelBinOutput\AvatarSkillExcelConfigData.json

//人物定义
for _,v := range AvatarExcelConfigData {
    //血量
    v.HpBase
    //基础攻击
    v.AttackBase
    //基础防御
    v.DefenseBase
    //基础暴击率
    v.Critical
    //基础暴击伤害(不含基础1.0倍率)
    v.CriticalHurt
    //基础充能效率
    v.ChargeEfficiency
    //角色名(TextMap)
    v.NameTextMapHash
    //图标
    v.IconName
    //武器类型
    v.WeaponType
    //[]成长参数
    v.PropGrowCurves
    for _,vv := range v.CurveInfos {
        //属性类型
        vv.Type
        //成长类型
        vv.Value
    }
    //人物成长ID
    v.AvatarPromoteId
    //技能ID(AvatarSkillDepotExcelConfigData.id)
    v.SkillDepotId
}

//人物基础升级提升值(累加)[1-100]
for _,v := range AvatarCurveExcelConfigData {
    //级别
    v.Level
    //[]成长类型
    v.CurveInfos
    for _,vv := range v.CurveInfos {
        //成长类型
        vv.Type
        //值
        vv.Value
    }
}

//人物突破提升值(对应)
for _,v := range AvatarPromoteExcelConfigData {
    //人物成长ID
    v.AvatarPromoteId
    //突破等级(星数)[0-6]
    v.PromoteLevel
    //解锁等级
    v.UnlockMaxLevel
    //增加各属性值(包含已经突破的相对0级的合计值)
    v.AddProps
    for _,vv := range v.AddProps {
        //属性类型
        vv.PropType
        //值
        vv.Value
    }
    //突破需要等级
    v.RequiredPlayerLevel
}

//人物技能集定义
for _,v := range AvatarSkillDepotExcelConfigData {
    //ID
    v.Id
    //Q技能
    v.EnergySkill
    //[]小技能ID(A,E)
    v.Skills
    //[]其他技能ID(武器CD，重击，队伍天赋共CD，命座效果，空中攻击)
    v.SubSkills
    //[]命座ID
    v.Talents
    //[]固有天赋
    v.InherentProudSkillOpens
}

//技能详细定义
for _,v := range AvatarSkillExcelConfigData {
    //ID
    v.Id
    //技能名(TextMap)
    v.NameTextMapHash
    //技能描述(TextMap)
    v.DescTextMapHash
    //技能CD
    v.CdTime
    //使用上限次数
    v.MaxChargeNum
    //能量消耗
    v.CostElemVal
    //能量消耗元素类型
    v.CostElemType
}
```