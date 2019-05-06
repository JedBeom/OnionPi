/*
function vote(id, up) {
    let url = "";
    if (up) {
       url = "/vote/"+id+"/+";
    } else {
       url = "/vote/"+id+"/+";
    }

    $.ajax({
        type: "GET",
        url: url,
        dataType: "json",

        success: function(response) {

        }

    }

}
 */

function vote_up() {
    alert(this.Attr("name"))
}