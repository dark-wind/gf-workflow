$("#start-button").on("click", function(event){
    let Status = $("#Status").val()
    console.log(Status)


    $.post( "/start",{
        "ProcessID": "1",
        "UserID": "16",
        "Conditions": Status
    }, function( data ) {
        console.log(data)
    });
});

$("#zsb-submit").on("click", function(event){
    console.log("zsb-submit");

    console.log(temp);

    // $.post( "/complete",{
    //     "TaskID": "1",
    //     "UserID": "9"
    // }, function( data ) {
    //     console.log(data)
    // });
});

$("#fdy-submit").on("click", function(event){
    console.log("zsb-submit");

    $.post( "/complete",{
        "TaskID": "1",
        "UserID": "10"
    }, function( data ) {
        console.log(data)
    });
});

$("#cwc-submit").on("click", function(event){
    console.log("zsb-submit");

    $.post( "/complete",{
        "TaskID": "1",
        "UserID": "11"
    }, function( data ) {
        console.log(data)
    });
});

$("#ds-submit").on("click", function(event){
    console.log("zsb-submit");

    $.post( "/complete",{
        "TaskID": "1",
        "UserID": "12"
    }, function( data ) {
        console.log(data)
    });
});

$("#dzb-submit").on("click", function(event){
    console.log("zsb-submit");

    $.post( "/complete",{
        "TaskID": "1",
        "UserID": "13"
    }, function( data ) {
        console.log(data)
    });
});

$("#tzb-submit").on("click", function(event){
    console.log("zsb-submit");

    $.post( "/complete",{
        "TaskID": "5",
        "UserID": "14"
    }, function( data ) {
        console.log(data)
    });
});

$("#sg-submit").on("click", function(event){
    console.log("zsb-submit");

    $.post( "/complete",{
        "TaskID": "5",
        "UserID": "15"
    }, function( data ) {
        console.log(data)
    });
});