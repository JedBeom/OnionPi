function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function getCookie(name) {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length === 2) return parts.pop().split(";").shift();
}

function vote(id, is_vote_up) {

    $.ajax({
        type: "POST",
        url: "/vote",
        data: {'id': id, 'is_vote_up': is_vote_up},
        dataType: "json",

        success: function (data) {


            let value = $("#vote_value_" + id);
            if (data.total_vote > 0) {
                data.total_vote = "+" + String(data.total_vote)
            }
            value.text(data.total_vote);

            if (data.user_vote === 0) {
                $("#vote_up_" + id).removeClass("btn-success").addClass("btn-secondary");
                $("#vote_down_" + id).removeClass("btn-success").addClass("btn-secondary");
            } else if (data.user_vote === 1) {
                $("#vote_up_" + id).removeClass("btn-secondary").addClass("btn-success");
                $("#vote_down_" + id).removeClass("btn-success").addClass("btn-secondary");
            } else if (data.user_vote === 2) {
                $("#vote_up_" + id).removeClass("btn-success").addClass("btn-secondary");
                $("#vote_down_" + id).removeClass("btn-secondary").addClass("btn-success");
            }

        }

    })
}

$("#submit_button").click(async function () {
    let content = $("#submit").val();
    $("#alert").css("visibility", "hidden");
    $("#spinner").css("visibility", "visible");

    await sleep(700);

    $.ajax({
        type: "POST",
        url: "/submit",
        data: {"content": content, "cookie": getCookie("_yayoiori")},
        dataType: "text",
        success: function (data) {
            alert("글을 올렸습니다!");
            $("#submit").val("");
            $(".spinner-border").css("visibility", "hidden");
            location.reload();
        },

        error: function (req, status, err) {
            if (req.responseText === "") {
                req.responseText = "오류가 발생했습니다!"
            }
            $("#alert").css("visibility", "visible").text(req.responseText);
        },
    });

    $(".spinner-border").css("visibility", "hidden");

});



