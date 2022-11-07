$("#start-button").on("click", function(event){
    let Status = $("#Status").val()
    console.log(Status)

    $.post( "/start",{
        "ProcessID": "1",
        "UserID": "1",
        "Conditions": Status
    }, function( data ) {
        console.log(data)
        location.reload()
    });
});

$("#zsb-submit").on("click", function(event){
    console.log("zsb-submit");

    let taskId = $("#taskId").text()

    $.post( "/complete",{
        "TaskID": taskId,
        "UserID": "2"
    }, function( data ) {
        console.log(data)
        location.reload()
    });
});

$("#fdy-submit").on("click", function(event){
    console.log("zsb-submit");
    let taskId = $("#taskId").text()


    $.post( "/complete",{
        "TaskID": taskId,
        "UserID": "3"
    }, function( data ) {
        console.log(data)
        location.reload()
    });
});

$("#cwc-submit").on("click", function(event){
    console.log("zsb-submit");
    let taskId = $("#taskId").text()


    $.post( "/complete",{
        "TaskID": taskId,
        "UserID": "4"
    }, function( data ) {
        console.log(data)
        location.reload()
    });

});

$("#ds-submit").on("click", function(event){
    console.log("zsb-submit");
    let taskId = $("#taskId").text()


    $.post( "/complete",{
        "TaskID": taskId,
        "UserID": "5"
    }, function( data ) {
        console.log(data)
        location.reload()
    });
});

$("#dzb-submit").on("click", function(event){
    console.log("zsb-submit");
    let taskId = $("#taskId").text()


    $.post( "/complete",{
        "TaskID": taskId,
        "UserID": "6"
    }, function( data ) {
        console.log(data)
        location.reload()
    });
});

$("#tzb-submit").on("click", function(event){
    console.log("zsb-submit");
    let taskId = $("#taskId").text()


    $.post( "/complete",{
        "TaskID": taskId,
        "UserID": "7"
    }, function( data ) {
        console.log(data)
        location.reload()
    });
});

$("#sg-submit").on("click", function(event){
    console.log("zsb-submit");
    let taskId = $("#taskId").text()


    $.post( "/complete",{
        "TaskID": taskId,
        "UserID": "8"
    }, function( data ) {
        console.log(data)
        location.reload()
    });
});