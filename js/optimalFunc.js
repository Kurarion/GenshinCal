const OptimalStep = 10;
const MinValid = 0;
const MaxValid = (4 + 5) * 5 * OptimalStep;

const OPTIMAL_SWITCH_NAME = TYPE_OPTIMAL + "Switch";

var isOptimalMode = false;
var hadOptimalCal = false;

var optimalReliquaryAffixStepObj = {};
var optimalReliquaryAffixNameList = [];
var optimalReliquaryResult = [];

function changeCalMode(){
    isOptimalMode = $("#" + OPTIMAL_SWITCH_NAME)[0].checked;
    if(isOptimalMode){
        currentAllType = TYPE_OPTIMAL;
        currentReliquaryType = TYPE_OPTIMAL_RELIQUARY;
        optimalUpdate();
    }else{
        currentAllType = TYPE_ALL;
        currentReliquaryType = TYPE_RELIQUARY;
        update();
    }
}

function setOptimalMainValue(pos,index){
    var valueObj = getOptimalMainValue(pos,index);
    var value = 0;
    var obj = getOptimalValueInputObj(pos,index);
    if(valueObj.error){
        obj.val(value);
        // update();
        return
    }
    value = valueObj.data;
    obj.val(value);
    obj.attr("name",TYPE_OPTIMAL_RELIQUARY + propertyNameMap[getOptimalSelectObj(pos,index).val()]);
    // update();
}

function getOptimalValueInputObj(pos,index){
    return $("#" + TYPE_OPTIMAL_RELIQUARY + pos + index + "Value");
}

function getOptimalSelectObj(pos,index){
    return $("#" + TYPE_OPTIMAL_RELIQUARY + pos + index);
}

function getOptimalMainValue(pos,index){
    var resObj = {};
    var res = reliquaryMainObj[getOptimalSelectObj(pos,index).val()];
    if(!res){
        res = null;
        resObj.error = true;
    }
    resObj.data = res
    return resObj;
}

//最优更新函数
function optimalUpdate(){
    setToAll();
    calDamage();
}

function setOptimalValidNum(num){
    var optimalValidObj = $("#" + TYPE_OPTIMAL + "ValidNum");
    var optimalValidShowObj = $("#" + TYPE_OPTIMAL + "ValidNum" + "Show");
    if(typeof num == "undefined"){
        num = optimalValidObj.val();
    }
    optimalValidObj.val(num);
    optimalValidShowObj.val(num);
    setToOptimal(num);
    optimalUpdate();
}

function setOptimalValidNumByWheel(){
    var e = window.event;
    var delta = Math.max(-1, Math.min(1, (e.wheelDelta || -e.detail)));
    var num = parseInt($("#" + TYPE_OPTIMAL + "ValidNum").val());
    if(delta > 0){
        ++num;
    }else if (delta < 0){
        --num;
    }
    if(num > MaxValid){
        num = MaxValid;
    }else if(num < MinValid){
        num = MinValid;
    }
    setOptimalValidNum(num);
    return false;
}

function setOptimalCriMax(num){
    var optimalCriMaxObj = $("#" + TYPE_OPTIMAL + "CriMax");
    var optimalCriMaxShowObj = $("#" + TYPE_OPTIMAL + "CriMax" + "Show");
    if(typeof num == "undefined"){
        num = optimalCriMaxObj.val();
    }
    optimalCriMaxObj.val(num);
    optimalCriMaxShowObj.val(num);
    if(!checkOptimal()){
        //nothing
    }
}

function setOptimalCriMaxByWheel(){
    var e = window.event;
    var delta = Math.max(-1, Math.min(1, (e.wheelDelta || -e.detail)));
    var num = parseFloat($("#" + TYPE_OPTIMAL + "CriMax").val());
    if(delta > 0){
        num += 0.01;
    }else if (delta < 0){
        num -= 0.01;
    }
    if(num > 1.0){
        num = 1.0;
    }else if(num < 0.05){
        num = 0.05;
    }
    setOptimalCriMax(num.toFixed(2));
    return false;
}

function initOptimalStep(){
    for(i in reliquaryAffixObj){
        var temp = reliquaryAffixObj[i];
        optimalReliquaryAffixStepObj[i] = (temp[temp.length - 1] / 10).toFixed(5);
    }
    optimalReliquaryAffixNameList = [];
    optimalReliquaryAffixNameList.push("FIGHT_PROP_ATTACK_PERCENT");
    optimalReliquaryAffixNameList.push("FIGHT_PROP_DEFENSE_PERCENT");
    optimalReliquaryAffixNameList.push("FIGHT_PROP_HP_PERCENT");
    optimalReliquaryAffixNameList.push("FIGHT_PROP_ELEMENT_MASTERY");
    optimalReliquaryAffixNameList.push("FIGHT_PROP_CRITICAL");
    optimalReliquaryAffixNameList.push("FIGHT_PROP_CRITICAL_HURT");
}

function getOptimalDamageType(){
    return $("select#" + TYPE_OPTIMAL + "TargetDamageType").val();
}

function getOptimalMaxCritical(){
    return parseFloat($("input#" + TYPE_OPTIMAL + "CriMax").val());
}

function optimalCal(){
    if(!isOptimalMode){
        alert("请勾选【最优计算模式】后进行计算");
        return;
    }
    optimalReliquaryResult = [];
    hadOptimalCal = true;
    var targetDamageType = getOptimalDamageType();
    var maxCritical = getOptimalMaxCritical();
    var totalCritObj =$("input#"+TYPE_OPTIMAL+PROP_CRITICAL);
    for(var i = 0; i <= MaxValid; ++i){
        var thisLoopMax = parseFloat($("input#"+targetDamageType).val());
        var UpdatedName = "";
        var UpdatedValue = 0.0;
        var saveObj = {};
        saveObj.data = {};
        setToOptimal(i - 1);

        for(var j = 0; j < optimalReliquaryAffixNameList.length; ++j){
            var currentPropName = optimalReliquaryAffixNameList[j];
            var valueName = propertyNameMap[currentPropName];
            var value = parseFloat(optimalReliquaryAffixStepObj[currentPropName]);
            var tempInput = $("input[name=" + TYPE_OPTIMAL_RELIQUARY + valueName + "][optimal='true']");
            var tempOldValue = parseFloat(tempInput.val());
            //check max crit
            if(PROP_CRITICAL == valueName && parseFloat(totalCritObj.val()) >= maxCritical){
                continue;
            }
            var newValue = tempOldValue + value;
            tempInput.val(newValue);

            //save
            saveObj.data[valueName] = tempOldValue;
            optimalUpdate();
            var currentRes = parseFloat($("input#"+targetDamageType).val());
            if(currentRes > thisLoopMax){
                UpdatedName = valueName;
                UpdatedValue = newValue;
                thisLoopMax = currentRes;
            }
            tempInput.val(tempOldValue);
        }
        //save
        saveObj.best = UpdatedName;
        saveObj.data[saveObj.best] = UpdatedValue;

        optimalReliquaryResult.push(saveObj);
    }

    setOptimalValidNum(MaxValid);
}

function setToOptimal(index){
    if(index < 0 || index > optimalReliquaryResult.length){
        $("input[optimal='true']").val(0);
        return;
    }
    var temp = optimalReliquaryResult[index];
    var tempData = temp.data;
    for(i in tempData){
        $("input[name=" + TYPE_OPTIMAL_RELIQUARY + i + "][optimal='true']").val(tempData[i]);
    }
    //color
    $("input[optimal='true']").removeAttr("style");
    //此次最优
    $("input[name=" + TYPE_OPTIMAL_RELIQUARY + temp.best + "][optimal='true']").attr("style","color:rgb(255, 174, 0)");
    // //下次最优
    // if(index + 1 < optimalReliquaryResult.length){
    //     $("input[name=" + TYPE_OPTIMAL_RELIQUARY + optimalReliquaryResult[index + 1].best + "][optimal='true']").attr("style","color:rgb(255, 174, 0)");
    // }
}