
const TYPE_CHARACTER = "character";
const TYPE_WEAPON = "weapon";
const TYPE_WEAPON_SKILL_AFFIX = "weaponSkillAffix";

function getData(types){
    if(!(types instanceof Array)){
        return;
    }
    for(var i = 0; i < types.length; ++i){
        var loopType = types[i];
        var loopData = createData(loopType);
        if(loopData.length != 0){
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
}

function createData(type){
    var data = "";
    switch (type){
        case TYPE_CHARACTER:
            id = $("#"+type+"Select").val();
            level = encodeURIComponent($("#"+type+"LevelSelect").val());
            if(id.length != 0 && level.length != 0){
                data = "id=" + id + "&level=" + level;
            }
            break;
        case TYPE_WEAPON:
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
    }
    return data
}

function createCallback(type) {
    var func = function(){};
    switch (type){
        case TYPE_CHARACTER:
            func = function(result){
                $("#" + type + "Result").html(parseJSON(type, result))
            }
            break;
        case TYPE_WEAPON:
            func = function(result){
                $("#" + type + "Result").html(parseJSON(type, result))
            }
            break;
        case TYPE_WEAPON_SKILL_AFFIX:
            func = function(result){
                $("#" + type + "Result").html(parseJSON(type, result))
            }
            break;
    }
    return func
}

function parseJSON(type, input){
    res = "";
    var obj = JSON.parse(input)
    switch (type){
        case TYPE_CHARACTER:
            res += "生命值: " + obj.Hp
            res += "<br>攻击值: " + obj.Attack
            res += "<br>防御值: " + obj.Defense
            res += "<br>生命百分比: " + obj.Hp_percent
            res += "<br>攻击百分比: " + obj.Attack_percent
            res += "<br>防御百分比: " + obj.Defense_percent
            res += "<br>暴击率: " + obj.Critical
            res += "<br>暴击伤害: " + obj.CriticalHurt
            res += "<br>充能效率: " + obj.ChargeEfficiency
            res += "<br>治疗增益: " + obj.HealActiveUp
            res += "<br>受治疗增益: " + obj.HealPassiveUp
            res += "<br>元素精通: " + obj.ElementMaster
            res += "<br>冰伤: " + obj.Ice
            res += "<br>风伤: " + obj.Wind
            res += "<br>物伤: " + obj.Physical
            res += "<br>雷伤: " + obj.Elec
            res += "<br>岩伤: " + obj.Rock
            res += "<br>火伤: " + obj.Fire
            res += "<br>水伤: " + obj.Water
            break;
        case TYPE_WEAPON:
            res += "攻击值: " + obj.Attack
            res += "<br>生命百分比: " + obj.Hp_percent
            res += "<br>攻击百分比: " + obj.Attack_percent
            res += "<br>防御百分比: " + obj.Defense_percent
            res += "<br>暴击率: " + obj.Critical
            res += "<br>暴击伤害: " + obj.CriticalHurt
            res += "<br>充能效率: " + obj.ChargeEfficiency
            res += "<br>治疗增益: " + obj.HealActiveUp
            res += "<br>受治疗增益: " + obj.HealPassiveUp
            res += "<br>元素精通: " + obj.ElementMaster
            res += "<br>冰伤: " + obj.Ice
            res += "<br>风伤: " + obj.Wind
            res += "<br>物伤: " + obj.Physical
            res += "<br>雷伤: " + obj.Elec
            res += "<br>岩伤: " + obj.Rock
            res += "<br>火伤: " + obj.Fire
            res += "<br>水伤: " + obj.Water
            break;
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