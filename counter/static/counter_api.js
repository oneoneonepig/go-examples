
function flushInputKeyName() {
    $('#inputKeyName').val("");
}

function flushInputValue() {
    $('#inputValue').val("");
}

function flushAllInput() {
    $('#inputKeyName').val("");
    $('#inputValue').val("");
}

function flushResponseDump() {
    $('#responseDump').html("");
}

function get() {
    var keyName = $('#inputKeyName').val()
    var settings = {
        "url": `http://localhost:8080/key/${keyName}`,
        "method": "GET",
    }

    $.ajax(settings).done(function (response) {
        console.log(response);
        $('#counterMessage').text(response.message);
        $('#inputValue').val(response.value);
        flushResponseDump();
    });
}
function set() {
    var keyName = $('#inputKeyName').val()
    var value = $('#inputValue').val()
    var settings = {
        "url": `http://localhost:8080/key/${keyName}/${value}`,
        "method": "POST",
    }

    $.ajax(settings).done(function (response) {
        console.log(response);
        $('#counterMessage').text(response.message);
        flushResponseDump();
    });
}

function incr() {
    var keyName = $('#inputKeyName').val()
    var currentValue;
    $.ajax({
        "url": `http://localhost:8080/key/${keyName}`,
        "method": "GET",
    }).done(function (response) {
        console.log(response);
        currentValue = response.value;
        currentValue += 1
        $.ajax({
            "url": `http://localhost:8080/key/${keyName}/${currentValue}`,
            "method": "POST",
        }).done(function (response) {
            console.log(response);
            $("#inputValue").val(currentValue);
            $("#counterMessage").html("key increased");
            flushResponseDump();
        });
    });
}

function del(){
    var keyName = $('#inputKeyName').val()
    var settings = {
        "url": `http://localhost:8080/key/${keyName}`,
        "method": "DELETE",
    }

    $.ajax(settings).done(function (response) {
        console.log(response);
        $('#counterMessage').text(response.message);
        flushAllInput();
        flushResponseDump();
    });
}
function getAll(){
    var settings = {
        "url": "http://localhost:8080/keys",
        "method": "GET",
    }

    $.ajax(settings).done(function (response) {
        console.log(response);
        $("#responseDump").html("");
        $("#responseDump").append(response.message);
        $('#counterMessage').text("keys dumped");
        flushAllInput();
    });
}

function delAll(){
    var settings = {
        "url": "http://localhost:8080/keys",
        "method": "DELETE",
    }

    $.ajax(settings).done(function (response) {
        console.log(response);
        $('#counterMessage').text(response.message);
        flushResponseDump();
    });
}

function cleanUp(){
    flushAllInput();
    flushResponseDump();

}