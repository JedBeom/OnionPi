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
                $("#post-" + id).removeClass("table-success table-danger")
            } else if (data.user_vote === 1) {
                $("#post-" + id).removeClass("table-danger").addClass("table-success");
            } else if (data.user_vote === 2) {
                $("#post-" + id).removeClass("table-success").addClass("table-danger");
            }

        }

    })
}

$("#submit_button").click(async function () {
    let content = $("#submit").val();
    $("#alert").css("visibility", "hidden");
    $("#spinner").css("visibility", "visible");

    await sleep(500);

    content = content.trim();
    if (content === "") {
        $("#alert").css("visibility", "visible").text("내용이 없습니다.");
        $(".spinner-border").css("visibility", "hidden");
        return
    }

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



