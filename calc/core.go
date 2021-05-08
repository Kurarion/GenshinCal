package calc

//废弃

//				|-------------------------------------------------------
//				| 					| [人物攻击力 + 武器攻击力](基础攻击力)
//				| 					| *
//				| 					1 <攻击力增加百分比> + 1
//				| 					| +
//				| 					| <攻击力增加绝对值>(防御转换，生命转换，圣遗物小数值)
//				|					| ----------------------------------
//				| 					| [人物防御力]
//				| 					| *
//				|		基础数值区	 2 <防御力增加百分比> + 1
//				| 					| +
//				| 					| <防御力增加绝对值>(圣遗物小数值)
//				|					| ----------------------------------
//				| 					| [人物生命]
//				| 					| *
//				| 					3 <生命增加百分比> + 1
//				| 					| +
//				| 					| <生命增加绝对值>(圣遗物小数值)
//				|-------------------------------------------------------
//	基础伤害值 =		[1x3]	*	[3x1]
//				|-------------------------------------------------------
//				|					1 <攻击力倍率>
//				|					| ----------------------------------
//				|		技能倍率	 2 <防御力倍率>(阿贝多E绽放伤害)
//				|					| ----------------------------------
//				|					3 <生命倍率>(钟离:天赋, 胡桃:血梅香)
//				|-------------------------------------------------------

//		*

//							特定技能类(E, Q)
//	增伤百分比 + 1			<元素增伤>(物理, 各元素)
//							攻击类型(普攻, 重击, 下落)
//							全增伤

//		*

//							[人物暴击率]
//	暴击率					[武器暴击率]
//							<圣遗物暴击率>
//
//		*
//
//							[人物暴击伤害]
//	暴击伤害 + 1			[武器暴击伤害]
//							<圣遗物暴击伤害>

//		*

//	增幅反应加成系数 + 1	<元素精通>

//		*

//							角色等级
//	等级，减防相关增减伤	施加减防
//							怪物等级

//		*

//							怪物自带抗性
//	各元素抗性
//							施加减抗

// //基础伤害值
// func (c InputData) getBaseDamageValue(art Artifacts) float64 {
// 	return 1.0
// }

// //增加伤害值
// func (c InputData) getBoostDamageValue(art Artifacts) float64 {
// 	BaseDamageValueMat := getBaseDamageValue()
// 	BoostMat := .0
// 	return BaseDamageValueMat * BoostMat
// }

// //基础伤害值 + 增加伤害值
// func (c InputData) getBaseDamageValueWithBoost() float64 {
// 	return getBaseDamageValue() + getBoostDamageValue()
// }

// //暴击率
// func (c InputData) getCriticalRate() float64 {
// 	return 0.05
// }

// //暴击伤害加成系数
// func (c InputData) getCriticalBoost() float64 {
// 	return 1.0 + 0.5
// }

// //增幅反应加成系数
// func (c InputData) getElementBoost() float64 {
// 	return 1.0 + 0.0
// }

// //等级，减防相关增减伤系数
// func (c InputData) getLevelAndDefEffect() float64 {
// 	return 0.5
// }

// //各元素抗性系数
// func (c InputData) getResistanceEffect() float64 {
// 	return 0.9
// }

// //最终结果
// func (c InputData) GetDamageValue() float64 {
// 	valuebase := getBaseDamageValueWithBoost() * getLevelAndDefEffect() * getResistanceEffect()
// 	valueWithCrit := valuebase * getCriticalBoost()
// 	valueWithCritExpect := valueWithCrit * getCriticalRate()
// 	valueWithEle := valuebase * getElementBoost()
// 	valueWithCritWithEleExpect := valueWithCrit * getElementBoost()
// 	valueAllExpect := valueWithCritExpect * getElementBoost()
// 	return 0.0
// }
