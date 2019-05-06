function vote(id, is_vote_up) {
    let url = "";
    let add = 0;
    if (is_vote_up) {
       url = "/vote/"+id+"/+";
        add = 1
    } else {
        url = "/vote/" + id + "/-";
        add = -1
    }

    $.ajax({
        type: "GET",
        url: url,
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

