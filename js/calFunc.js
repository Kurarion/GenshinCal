const PROP_ATTACK = "Attack";
const PROP_ATTACK_PERCENT = "Attack_percent";
const PROP_CHARGE_EFFICIENCY = "ChargeEfficiency";
const PROP_CRITICAL = "Critical";
const PROP_CRITICAL_HURT = "CriticalHurt";
const PROP_DEFENSE = "Defense";
const PROP_DEFENSE_PERCENT = "Defense_percent";
const PROP_ELEMENT_MASTERY = "ElementMaster";
const PROP_HP = "Hp";
const PROP_HP_PERCENT = "Hp_percent";

const PROP_HEALACTIVEUP = "HealActiveUp";
const PROP_HEALPASSIVEUP = "HealPassiveUp";

const PROP_BOOST_ICE = "Ice";
const PROP_BOOST_WIND = "Wind";
const PROP_BOOST_PHYSICAL = "Physical";
const PROP_BOOST_ELEC = "Elec";
const PROP_BOOST_ROCK = "Rock";
const PROP_BOOST_FIRE = "Fire";
const PROP_BOOST_WATER = "Water";
const PROP_BOOST_GRASS = "Grass";
const PROP_BOOST_ALL = "All";
const PROP_BOOST_NORMALATK = "NormalAtk";
const PROP_BOOST_HEAVYATK = "HeavyAtk";
const PROP_BOOST_SKILL = "Skill";
const PROP_BOOST_ULT = "Ult";

const Reaction_Rate_1_5 = 1.5;
const Reaction_Rate_2_0 = 2.0;

const NAME_ATTAK_RATIO = "AttakRatio";
const NAME_DEFENSE_RATIO = "DefenseRatio";
const NAME_HP_RATIO = "HpRatio";
const NAME_ABS_VALUE = "AbsValue";
const NAME_BOOST_HURT = "BoostHurt";
const NAME_SUB_HURT = "SubHurt";
const NAME_ALL_SUB_HURT = "AllSubHurt";
const NAME_ELEMENT_SUB_HURT = "ElementSubHurt";
const NAME_DEFENSE_SUB_HURT = "DefenseSubHurt";
const NAME_VALUE_WITHOUTCRI = "ValueWithoutCri";
const NAME_VALUE_WITHCRI = "ValueWithCri";
const NAME_VALUE_EXPECT = "ValueExpect";
const NAME_1_5_RATIO = "15";
const NAME_2_0_RATIO = "20";

const NAME_TYPE = "Type";
const NAME_ELEMENT_TYPE = "ElementType";

var currentDamageType;
var currentDamageElementType;
var currentDamageReactionRate1_5;
var currentDamageReactionRate2_0;

//更新函数
function update(){
    setToAll();
    calDamage();
}

function getDamageInputObj(type){
    return $("#" + TYPE_DAMAGE + type);
}

function getAllInputObj(type){
    return $("#" + TYPE_ALL + type);
}

function setAllInputObjGeneral(type){
    var toSetVal = getAllPropValue(type);
    getAllInputObj(type).val(toSetVal);
    return toSetVal;
}

function getPropValueBaseByID(type, name){
    var obj = $("#" + type + name);
    var value = obj.val();
    if(value){
        return parseFloat(value);
    }
    return 0
}

function getPropValueBaseByName(type, name){
    var obj = $("input[name='" + type + name + "']");
    var value = 0;
    for(var i = 0; i < obj.length; ++i){
        var temp = obj[i].value;
        if(temp){
            value += parseFloat(temp);
        }
    }
    return value;
}
//================================

//基础攻击力
function getBaseAtk(){
    var val = 0;
    val += getPropValueBaseByID(TYPE_CHARACTER, PROP_ATTACK);
    val += getPropValueBaseByID(TYPE_WEAPON, PROP_ATTACK);
    return val;
}

//基础防御力
function getBaseDefense(){
    var val = 0;
    val += getPropValueBaseByID(TYPE_CHARACTER, PROP_DEFENSE);
    return val;
}

//基础生命值
function getBaseHp(){
    var val = 0;
    val += getPropValueBaseByID(TYPE_CHARACTER, PROP_HP);
    return val;
}

//================================

//
function getAllPropValue(typeName){
    var val = 0;
    val += getPropValueBaseByID(TYPE_CHARACTER, typeName);
    val += getPropValueBaseByID(TYPE_WEAPON, typeName);
    val += getPropValueBaseByName(TYPE_RELIQUARY, typeName);
    val += getPropValueBaseByID(TYPE_RELIQUARY_SET, typeName);
    val += getPropValueBaseByID(TYPE_OTHER, typeName);
    return val;
}

//攻击力
function getAtk(buffPercent){
    if(!buffPercent){
        //攻击力百分比
        buffPercent = 0;
        buffPercent += getAllPropValue(PROP_ATTACK_PERCENT);
    }
    //攻击值
    var absVal = 0;
    //基础攻击
    absVal += getPropValueBaseByName(TYPE_RELIQUARY, PROP_ATTACK);
    absVal += getPropValueBaseByID(TYPE_RELIQUARY_SET, PROP_ATTACK);
    absVal += getPropValueBaseByID(TYPE_OTHER, PROP_ATTACK);
    //总攻击
    var totalVal = getBaseAtk() * ( 1 + buffPercent ) + absVal;
    return totalVal;
}

//防御力
function getDefense(buffPercent){
    if(!buffPercent){
        //防御力百分比
        buffPercent = 0;
        buffPercent += getAllPropValue(PROP_DEFENSE_PERCENT);
    }
    //防御值
    var absVal = 0;
    absVal += getPropValueBaseByName(TYPE_RELIQUARY, PROP_DEFENSE);
    absVal += getPropValueBaseByID(TYPE_RELIQUARY_SET, PROP_DEFENSE);
    absVal += getPropValueBaseByID(TYPE_OTHER, PROP_DEFENSE);
    //总防御
    var totalVal = getBaseDefense() * ( 1 + buffPercent ) + absVal;
    return totalVal;
}

//生命值
function getHp(buffPercent){
    if(!buffPercent){
        //生命值百分比
        buffPercent = 0;
        buffPercent += getAllPropValue(PROP_HP_PERCENT);
    }
    //生命值
    var absVal = 0;
    absVal += getPropValueBaseByName(TYPE_RELIQUARY, PROP_HP);
    absVal += getPropValueBaseByID(TYPE_RELIQUARY_SET, PROP_HP);
    absVal += getPropValueBaseByID(TYPE_OTHER, PROP_HP);
    //总生命
    var totalVal = getBaseHp() * ( 1 + buffPercent ) + absVal;
    return totalVal;
}

//================================

//附加属性
function setExtra(index){
    var targetPropName = getExtraSelectObj(index).val();
    var calPropName = getExtraCalSelectObj(index).val();
    var calValue = parseFloat(getExtraCalValueObj(index).val());
    var maxValue = parseFloat(getExtraMaxValueObj(index).val());

    if(targetPropName.length != 0
        && calPropName.length != 0
        && calValue){
            var toSetVal = parseFloat(getAllInputObj(calPropName).val()) * calValue;
            if(toSetVal > maxValue && maxValue != 0){
                toSetVal = maxValue;
            }
            var targetObj = getAllInputObj(targetPropName);
            targetObj.val(parseFloat(targetObj.val()) + toSetVal);
        }
}

//更新到合计属性
function setToAll(){
    var atkPercent = setAllInputObjGeneral(PROP_ATTACK_PERCENT);
    // setAllInputObjGeneral(PROP_ATTACK);
    getAllInputObj(PROP_ATTACK).val(getAtk(atkPercent));
    setAllInputObjGeneral(PROP_CHARGE_EFFICIENCY);
    setAllInputObjGeneral(PROP_CRITICAL);
    setAllInputObjGeneral(PROP_CRITICAL_HURT);
    var defensePercent = setAllInputObjGeneral(PROP_DEFENSE_PERCENT);
    // setAllInputObjGeneral(PROP_DEFENSE);
    getAllInputObj(PROP_DEFENSE).val(getDefense(defensePercent));
    setAllInputObjGeneral(PROP_ELEMENT_MASTERY);
    var hpPercent = setAllInputObjGeneral(PROP_HP_PERCENT);
    // setAllInputObjGeneral(PROP_HP);
    getAllInputObj(PROP_HP).val(getHp(hpPercent));
    setAllInputObjGeneral(PROP_HEALACTIVEUP);
    setAllInputObjGeneral(PROP_HEALPASSIVEUP);
    setAllInputObjGeneral(PROP_BOOST_ICE);
    setAllInputObjGeneral(PROP_BOOST_WIND);
    setAllInputObjGeneral(PROP_BOOST_PHYSICAL);
    setAllInputObjGeneral(PROP_BOOST_ELEC);
    setAllInputObjGeneral(PROP_BOOST_ROCK);
    setAllInputObjGeneral(PROP_BOOST_FIRE);
    setAllInputObjGeneral(PROP_BOOST_WATER);
    setAllInputObjGeneral(PROP_BOOST_GRASS);
    setAllInputObjGeneral(PROP_BOOST_ALL);
    setAllInputObjGeneral(PROP_BOOST_NORMALATK);
    setAllInputObjGeneral(PROP_BOOST_HEAVYATK);
    setAllInputObjGeneral(PROP_BOOST_SKILL);
    setAllInputObjGeneral(PROP_BOOST_ULT);
    
    //附加值
    for(var i = 0; i < MAX_EXTRA_NUM; ++i){
        setExtra(i);
    }
}

//================================

function calDamage(){
    currentDamageType = getDamageInputObj(NAME_TYPE).val();
    currentDamageElementType = getDamageInputObj(NAME_ELEMENT_TYPE).val();

    if(currentDamageType.length !=0 && currentDamageElementType !=0){
        //1.伤害值
        updateDamageOriginValue();
        //2.暴击区间
        updateDamageCritiacl();
        //3.增伤区间
        updateDamageBoostDamage();
        //4.减抗区间
        updateDamageDebuffEle();
        //5.防御区间
        updateDamageDefense();
        //6.反应区间
        updateDamageReaction();
        //计算
        calculate();
    }
}

//伤害值
function updateDamageOriginValue(){
    var attackRatio = getPropValueBaseByID(TYPE_DAMAGE, NAME_ATTAK_RATIO);
    var defenseRatio = getPropValueBaseByID(TYPE_DAMAGE, NAME_DEFENSE_RATIO);
    var hpRatio = getPropValueBaseByID(TYPE_DAMAGE, NAME_HP_RATIO);
    var absObj = getDamageInputObj(NAME_ABS_VALUE);
    var value = getPropValueBaseByID(TYPE_ALL,PROP_ATTACK)*attackRatio + getPropValueBaseByID(TYPE_ALL,PROP_DEFENSE)*defenseRatio + getPropValueBaseByID(TYPE_ALL,PROP_HP)*hpRatio;
    absObj.val(value);
}

//暴击区间
function updateDamageCritiacl(){
    getDamageInputObj(PROP_CRITICAL).val(getAllInputObj(PROP_CRITICAL).val());
    getDamageInputObj(PROP_CRITICAL_HURT).val(getAllInputObj(PROP_CRITICAL_HURT).val());
}

//增伤区间
function updateDamageBoostDamage(){
    var value = 0;
    value += getPropValueBaseByID(TYPE_ALL, currentDamageType);
    value += getPropValueBaseByID(TYPE_ALL, currentDamageElementType);
    value += getPropValueBaseByID(TYPE_ALL, PROP_BOOST_ALL);
    getDamageInputObj(NAME_BOOST_HURT).val(value);
}

//减抗区间
function updateDamageDebuffEle(){
    var value = 0;
    value += getPropValueBaseByID(TYPE_MONSTER, currentDamageElementType + NAME_SUB_HURT);
    value -= getPropValueBaseByID(TYPE_MONSTER_DEBUFF, currentDamageElementType + NAME_SUB_HURT);
    value -= getPropValueBaseByID(TYPE_MONSTER_DEBUFF, NAME_ALL_SUB_HURT);
    if(value < 0){
        value /= 2;
    }
    getDamageInputObj(NAME_ELEMENT_SUB_HURT).val(value);
}

//防御区间
function updateDamageDefense(){
    var value = 0;
    // [ref] https://bbs.nga.cn/read.php?tid=23708327
    // Def_c/(Def_c+5l+500)
    var monsterDef = 0;
    monsterDef += getPropValueBaseByID(TYPE_MONSTER, PROP_DEFENSE);
    monsterDef *= 1 - getPropValueBaseByID(TYPE_MONSTER_DEBUFF, PROP_DEFENSE);
    var characterLevel = 0;
    characterLevel = parseInt($("#characterLevelSelect").val());
    value = monsterDef/(monsterDef + characterLevel*5 + 500);
    getDamageInputObj(NAME_DEFENSE_SUB_HURT).val(value);
}

//反应区间
function updateDamageReaction(){
    // [ref] https://bbs.nga.cn/read.php?tid=26143970#postauthor5
    var ele = getPropValueBaseByID(TYPE_ALL, PROP_ELEMENT_MASTERY);
    var temp = 2.78/(1 + 1400/ele);
    currentDamageReactionRate1_5 = Reaction_Rate_1_5 * (1 + temp);
    currentDamageReactionRate2_0 = Reaction_Rate_2_0 * (1 + temp);
}

//最终计算
function calculate(){
    var val1 = getPropValueBaseByID(TYPE_DAMAGE, NAME_ABS_VALUE);
    var val2 = getPropValueBaseByID(TYPE_DAMAGE, PROP_CRITICAL);
    if(val2 > 1){
        val2 = 1;
    }
    var val3 = 1 + getPropValueBaseByID(TYPE_DAMAGE, PROP_CRITICAL_HURT);
    var val4 = 1 + getPropValueBaseByID(TYPE_DAMAGE, NAME_BOOST_HURT);
    var val5 = 1 - getPropValueBaseByID(TYPE_DAMAGE, NAME_ELEMENT_SUB_HURT);
    var val6 = 1 - getPropValueBaseByID(TYPE_DAMAGE, NAME_DEFENSE_SUB_HURT);
    var val7 = currentDamageReactionRate1_5;
    var val8 = currentDamageReactionRate2_0;

    var temp1 = val1*val4*val5*val6;
    var temp2 = temp1*val3;
    var temp3 = temp2*val2 + temp1*(1-val2);
    var temp4 = temp1*val7;
    var temp5 = temp2*val7;
    var temp6 = temp3*val7;
    var temp7 = temp1*val8;
    var temp8 = temp2*val8;
    var temp9 = temp3*val8;

    getDamageInputObj(NAME_VALUE_WITHOUTCRI).val(temp1);
    getDamageInputObj(NAME_VALUE_WITHCRI).val(temp2);
    getDamageInputObj(NAME_VALUE_EXPECT).val(temp3);
    getDamageInputObj(NAME_VALUE_WITHOUTCRI + NAME_1_5_RATIO).val(temp4);
    getDamageInputObj(NAME_VALUE_WITHCRI + NAME_1_5_RATIO).val(temp5);
    getDamageInputObj(NAME_VALUE_EXPECT + NAME_1_5_RATIO).val(temp6);
    getDamageInputObj(NAME_VALUE_WITHOUTCRI + NAME_2_0_RATIO).val(temp7);
    getDamageInputObj(NAME_VALUE_WITHCRI + NAME_2_0_RATIO).val(temp8);
    getDamageInputObj(NAME_VALUE_EXPECT + NAME_2_0_RATIO).val(temp9);
}