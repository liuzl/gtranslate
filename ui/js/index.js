$(document).ready(function() {
    var input = GetQueryString("input");
    if (input != null) {
        $('#input').val(input);
        T(url);
    }
    $('#submit').click(function() {
        var input = $('#input').val().trim();
        if (input != "") {
            T(input);
        }
    });
    $("#input").on('keypress', function(e) {
        if (e.ctrlKey && e.which == 13) {
            var input = $('#input').val().trim();
            if (input != "") {
                T(input);
            }
        }
    });
});

function T(input) {
    $('#editor_holder').html("<h4>loading...</h4>");
    $("#visual").html("<h4>loading...</h4>");
    $.ajax({
        url: "/api/?input="+encodeURIComponent(input), cache: false,
        success: function(result) {
            $('#editor_holder').jsonview(result);
            visual(result)
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert(XMLHttpRequest.responseText);
        }
    });
}

function one(k, v) {
    return "<fieldset><legend class=\"label label-info left\">"+
        k+"</legend>"+v+"</fieldset>";
}

function visual(docs) {
    var html = "";
    for (var i = 0; i < docs.length; ++i) {
        html += one(docs[i].lang, docs[i].text);
    }
    $("#visual").html(html);
}
