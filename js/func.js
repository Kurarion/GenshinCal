
function getCharacter(value,toSetName){
    htmlobj=$.post(
        "./character",
        value,
        function(result){$("#" + toSetName).html(result)}
    );
}