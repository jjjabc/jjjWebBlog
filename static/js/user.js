function deluser(userId) {
	 $.ajax({
      url: "/userdel",
      dataType: 'html',
      type: 'POST',
      data: {userId : userId},
      success: checkuserReturn,
      error: function(xhr, status, err) {
        console.error("/user", status, err.toString());
      }.bind(this)
    });
	


}
function checkuserReturn(result) {
	if (result != "OK") {
		$("#contdiv").html(result);
	} else {
		userMangeView("#contdiv");
	}
}