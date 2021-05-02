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
const PROP_BOOST_ALL = "All";
const PROP_BOOST_NORMALATK = "NormalAtk";
const PROP_BOOST_HEAVYATK = "HeavyAtk";
const PROP_BOOST_SKILL = "Skill";
const PROP_BOOST_ULT = "Ult";


//更新函数
function update(){
    setToAll();
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
    
}