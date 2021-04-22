
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
                // $("#" + type + "Result").html(parseJSON(type, result))
                setJSON(type, result);
            }
            break;
        case TYPE_WEAPON:
            func = function(result){
                // $("#" + type + "Result").html(parseJSON(type, result))
                setJSON(type, result);
            }
            break;
        case TYPE_WEAPON_SKILL_AFFIX:
            func = function(result){
                $("#" + type + "Result").attr("title", parseJSON(type, result))
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