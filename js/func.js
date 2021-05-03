
const TYPE_CHARACTER = "character";
const TYPE_WEAPON = "weapon";
const TYPE_MONSTER = "monster";
const TYPE_WEAPON_SKILL_AFFIX = "weaponSkillAffix";
const TYPE_RELIQUARY = "reliquary";
const TYPE_RELIQUARY_MAIN = "reliquaryMain";
const TYPE_RELIQUARY_AFFIX = "reliquaryAffix";
const TYPE_RELIQUARY_SET = "reliquarySet";
const TYPE_OTHER = "other";
const TYPE_EXTRA = "extra";
const TYPE_ALL = "all";
const TYPE_DAMAGE = "damage";
const TYPE_MONSTER_DEBUFF = "monsterDebuff";
var reliquaryMainObj = {}
var reliquaryAffixObj = {}
const reliquaryMainName = "Main"
const reliquaryAffixName = "Affix"
const reliquaryPonitName = "Ponit"
const maxrReliquaryPonit = 5
var reliquaryAffixAllValue = {}

var propertyNameMap = {}

const MAX_EXTRA_NUM = 4

var lockVars = {}

function init(param){
    getData(param);
    propertyNameMap = {
        "FIGHT_PROP_ATTACK": PROP_ATTACK,
        "FIGHT_PROP_ATTACK_PERCENT": PROP_ATTACK_PERCENT,
        "FIGHT_PROP_CHARGE_EFFICIENCY": PROP_CHARGE_EFFICIENCY,
        "FIGHT_PROP_CRITICAL": PROP_CRITICAL,
        "FIGHT_PROP_CRITICAL_HURT": PROP_CRITICAL_HURT,
        "FIGHT_PROP_DEFENSE": PROP_DEFENSE,
        "FIGHT_PROP_DEFENSE_PERCENT": PROP_DEFENSE_PERCENT,
        "FIGHT_PROP_ELEMENT_MASTERY": PROP_ELEMENT_MASTERY,
        "FIGHT_PROP_HP": PROP_HP,
        "FIGHT_PROP_HP_PERCENT": PROP_HP_PERCENT,
        "FIGHT_PROP_ICE_ADD_HURT": PROP_BOOST_ICE,
        "FIGHT_PROP_WIND_ADD_HURT": PROP_BOOST_WIND,
        "FIGHT_PROP_PHYSICAL_ADD_HURT": PROP_BOOST_PHYSICAL,
        "FIGHT_PROP_ELEC_ADD_HURT": PROP_BOOST_ELEC,
        "FIGHT_PROP_ROCK_ADD_HURT": PROP_BOOST_ROCK,
        "FIGHT_PROP_FIRE_ADD_HURT": PROP_BOOST_FIRE,
        "FIGHT_PROP_WATER_ADD_HURT": PROP_BOOST_WATER,
        "FIGHT_PROP_GRASS_ADD_HURT": PROP_BOOST_GRASS,
    }
    $("input").change(update);
    $("#" + TYPE_DAMAGE + NAME_TYPE).change(update);
    $("#" + TYPE_DAMAGE + NAME_ELEMENT_TYPE).change(update);
    $("span#extraProperty select").change(update);
    var objs = $("input[id^=" + TYPE_ALL + "]");
    objs.attr("readonly","readonly");
    objs.attr("style","color:gray");
    lockVars = {
        "0": false,
        "1": false,
        "2": false,
        "3": false,
        "4": false,
    }
}

function initAllAffixValue(){
    for(i in reliquaryAffixObj){
        var baseValues = reliquaryAffixObj[i];
        reliquaryAffixAllValue[i] = baseValues.slice();;
        var toSetValues = reliquaryAffixAllValue[i];
        var unAddIndex = 0;
        for(var ii = 0; ii < maxrReliquaryPonit; ++ii){
            var currentLoopTime = toSetValues.length;
            for(; unAddIndex < currentLoopTime; ++unAddIndex){
                for(var iii = 0; iii < baseValues.length; ++iii){
                    toSetValues.push(toSetValues[unAddIndex] + baseValues[iii]);
                }
            }
        }
        toSetValues = toSetValues.sort(function(a,b)
        {
        return a - b
        })
        reliquaryAffixAllValue[i] = toSetValues.filter(function(item, index, toSetValues) {
        return toSetValues.indexOf(item, 0) === index;
        });
    }
}

function getData(types){
    if(!(types instanceof Array)){
        return;
    }
    for(var i = 0; i < types.length; ++i){
        var loopType = types[i];
        var loopData = createData(loopType);
        if(!loopData){
            return;
        }
        !function(ajaxType, ajaxData){
            $.ajax({
                url: "./api/" + ajaxType,
                type: "post",
                data: ajaxData,
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                success: function(res){
                    createCallback(ajaxType)(res);
                },
                error: function(e) {
                }
            })
        }(loopType, loopData);
    }
}

function createData(type){
    var data = null;
    switch (type){
        case TYPE_CHARACTER:
        case TYPE_WEAPON:
        case TYPE_MONSTER:
            id = $("#"+type+"Select").val();
            level = encodeURIComponent($("#"+type+"LevelSelect").val());
            if(id.length != 0 && level.length != 0){
                data = "id=" + id + "&level=" + level;
            }
            break;
        case TYPE_WEAPON_SKILL_AFFIX:
            id = $("#"+TYPE_WEAPON+"Select").val();
            level = encodeURIComponent($("#"+type+"Select").val());
            if(id.length != 0 && level.length != 0){
                data = "id=" + id + "&level=" + level;
            }
            break;
        case TYPE_RELIQUARY_MAIN:
        case TYPE_RELIQUARY_AFFIX:
            data = "1";
            break;
    }
    return data;
}

function createCallback(type) {
    var func = function(){};
    switch (type){
        case TYPE_CHARACTER:
        case TYPE_WEAPON:
        case TYPE_MONSTER:
            func = function(result){
                if(result != null && result.length != 0){
                    setJSON(type, result);
                    update();
                }
            }
            break;
        case TYPE_WEAPON_SKILL_AFFIX:
            func = function(result){
                $("#" + type + "Result").attr("title", parseJSON(type, result));
            }
            break;
        case TYPE_RELIQUARY_MAIN:
            func = function(result){
                reliquaryMainObj = JSON.parse(result);
            }
            break;
        case TYPE_RELIQUARY_AFFIX:
            func = function(result){
                reliquaryAffixObj = JSON.parse(result);
                initAllAffixValue();
            }
            break;
    }
    return func
}

function setJSON(type,input){
    var obj = JSON.parse(input)
    for(var i in obj){
        $("#" + type + i).val(obj[i])
    }
}

function parseJSON(type, input){
    res = "";
    var obj = JSON.parse(input)
    switch (type){
        case TYPE_WEAPON_SKILL_AFFIX:
            res += "特效名: " + obj.Name
            res += "<br>特效: <br>" + obj.Desc
            break;
    }
    return res;
}

function filterSelect(type){
    var value = $("#" + type + "TypeSelect").val();
    var targetSelect = $("#" + type + "Select");
    if(value == null || value.length == 0){
        targetSelect.children().each(function(i,v){
            $(v).show();
        });
        return;
    }

    targetSelect.children().each(function(i,v){
        if($(v).attr("type") == value){
            $(v).show();
        }else{
            $(v).hide();
        }
    });
}

function hideOtherSelect(pos,index){
    thisSelect = getReliquarySelectObj(pos,index);
    thisSelectID = thisSelect.attr('id');
    thisDiv = thisSelect.parent();
    thisDivID = thisDiv.attr('id');
    value = thisSelect.val();
    allSelectObjs = $("#" + thisDivID + " select");
    var values = [];
    allSelectObjs.each(function(i,v){
        val = $(v).val();
        if(val.length != 0){
            values.push(val);
        }
    });
    allSelectObjs.not("#" + thisSelectID).children().each(function(i,v){
        if(-1 == values.indexOf($(v).attr("value"))){
            $(v).show();
        }else{
            $(v).hide();
        }
    });
}

function getReliquaryValueInputObj(pos,index){
    return $("#" + TYPE_RELIQUARY + pos + index + "Value");
}

function getReliquarySelectObj(pos,index){
    return $("#" + TYPE_RELIQUARY + pos + index);
}

function getReliquaryMainValue(pos,index){
    var resObj = {};
    var res = reliquaryMainObj[getReliquarySelectObj(pos,index).val()];
    if(!res){
        res = null;
        resObj.error = true;
    }
    resObj.data = res
    return resObj;
}


function getReliquaryAffixValues(pos,index){
    var resObj = {};
    var res = reliquaryAffixAllValue[getReliquarySelectObj(pos,index).val()];
    if(!res){
        res = null;
        resObj.error = true;
    }
    resObj.data = res
    return resObj;
}

function setReliquaryMainValue(pos,index){
    if(lockVars[pos]){
        return;
    }
    var valueObj = getReliquaryMainValue(pos,index);
    var value = 0;
    if(valueObj.error){
        return
    }
    value = valueObj.data;
    var obj = getReliquaryValueInputObj(pos,index);
    obj.val(value);
    obj.attr("name",TYPE_RELIQUARY + propertyNameMap[getReliquarySelectObj(pos,index).val()]);
    update();
}

function resetReliquaryAffixValue(pos,index){
    if(lockVars[pos]){
        return;
    }
    var obj = getReliquaryValueInputObj(pos,index);
    obj.val(0);
    obj.attr("name",TYPE_RELIQUARY + propertyNameMap[getReliquarySelectObj(pos,index).val()]);
    update();
}

function addReliquaryAffixValue(pos,index,step){
    if(lockVars[pos]){
        return;
    }
    var valuesObj = getReliquaryAffixValues(pos,index);
    var values = [];
    if(valuesObj.error){
        return
    }
    values = valuesObj.data;
    var iObj = getReliquaryCloseValue(pos,index);
    var i = 0;
    if(iObj.error){
        return
    }
    i = iObj.data;
    var targetNum = 0;
    targetNum = i + step;
    if(targetNum >= values.length){
        targetNum = values.length - 1;
    }
    if(i < values.length - 1){
        if(iObj.isLower){
            targetNum = i;
        }
        getReliquaryValueInputObj(pos,index).val(values[targetNum]);
    }
    update();
}

function subReliquaryAffixValue(pos,index,step){
    if(lockVars[pos]){
        return;
    }
    var valuesObj = getReliquaryAffixValues(pos,index);
    var values = [];
    if(valuesObj.error){
        return
    }
    values = valuesObj.data;
    var iObj = getReliquaryCloseValue(pos,index);
    var i = 0;
    if(iObj.error){
        return
    }
    i = iObj.data;
    var targetNum = 0;
    targetNum = i - step;
    if(targetNum < 0){
        targetNum = 0;
    }
    if(i > 0){
        getReliquaryValueInputObj(pos,index).val(values[targetNum]);
    }
    update();
}

function setReliquaryAffixValueByWheel(pos,index){
    if(lockVars[pos]){
        return;
    }
    var e = window.event;
    var delta = Math.max(-1, Math.min(1, (e.wheelDelta || -e.detail)));
    if(delta > 0){
        addReliquaryAffixValue(pos,index,1);
    }else if (delta < 0){
        subReliquaryAffixValue(pos,index,1);
    }
    return false;
}

function getReliquaryCloseValue(pos,index){
    var resObj = {};
    var data = 0;
    var valuesObj = getReliquaryAffixValues(pos,index);
    var values = [];
    if(valuesObj.error){
        resObj.error = true;
        resObj.data = data;
        return resObj;
    }
    values = valuesObj.data;
    if(values.length == 0){
        resObj.error = true;
        resObj.data = data;
        return resObj;
    }
    var currentValue = getReliquaryValueInputObj(pos,index).val();

    var min = 0;
    var max = values.length - 1;
    var middle = 0;
    while(max - min > 1){
        middle = Math.floor((min + max)/2);
        if(values[middle] > currentValue){
            max = middle;
        }else{
            min = middle;
        }
    }

    if(values[max] - currentValue < currentValue - values[min] ){
        data = max;
    }else{
        data = min;
    }

    resObj.data = data;
    resObj.isLower = currentValue < values[data];
    return resObj;
}

function setReliquaryCloseValue(pos,index){
    var valuesObj = getReliquaryAffixValues(pos,index);
    var values = [];
    if(valuesObj.error){
        return
    }
    values = valuesObj.data;
    var iObj = getReliquaryCloseValue(pos,index);
    var i = iObj.data;
    if(iObj.error){
        return
    }
    i = iObj.data;
    getReliquaryValueInputObj(pos,index).val(values[i]);
    update();
}

function getExtraSelectObj(index){
    return $("#" + TYPE_EXTRA + "Name" + index);
}

function getExtraCalSelectObj(index){
    return $("#" + TYPE_EXTRA + "CalName" + index);
}

function getExtraCalValueObj(index){
    return $("#" + TYPE_EXTRA + "CalValue" + index);
}

function getExtraMaxValueObj(index){
    return $("#" + TYPE_EXTRA + "MaxValue" + index);
}

function lockReliquary(pos){
    var ele = $("#reliquaryLock" + pos);
    var selects = $("div#reliquaryDiv" + pos +" select");
    if(ele[0].checked){
        lockVars[pos] = true;
        selects.attr("disabled","disabled");
    }else{
        lockVars[pos] = false;
        selects.removeAttr("disabled");
    }
}