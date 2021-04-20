
function getCharacter(value,toSetName){
    htmlobj=$.post(
        "./api/character",
        value,
        function(result){$("#" + toSetName).html(result)}
    );
}