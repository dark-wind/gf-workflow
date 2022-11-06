$("#start-button").on("click", function(event){
    console.log("ababbaa");
    let Status = $("#Status").val()
    console.log(Status)


    $.post( "/start",{
        "ProcessID": "1",
        "UserID": "9",
        "Conditions": Status
    }, function( data ) {
        console.log(data)
    });
});