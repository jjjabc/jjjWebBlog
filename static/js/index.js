function winMin() {
	$("#imgdiv").css("background-image", "url(img/bg_min.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.4055" + "px");
	$("#Sbriefdiv").show();
	$("#Stxtdiv").show();
	$("#contdiv").hide();
	$("#Stitlediv").css("font-size", $("#imgdiv").height() * "0.3" + "px")
	$("#Stitlediv").css("padding-left", (($("#imgdiv").width() * 0.05) - 5) + "px")
	$("#Stitlediv").css("padding-top", ($("#imgdiv").height() *0.2) + "px")
	$("#Stxtdiv").css("padding-left", (($("#imgdiv").width() * 0.05) - 5) + "px")

}

function winSmall() {
	$("#imgdiv").css("background-image", "url(img/bg_s.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() + "px");
	$("#Sbriefdiv").hide();
	$("#Stxtdiv").hide();
	$("#contdiv").show();
	$("#textdiv").css("font-size", $("#imgdiv").width() * "0.03" + "px");
	$("#titlediv").css("font-size", $("#imgdiv").width() * "0.06" + "px");
	$("#contdiv").css("margin", "0");
	$("#contdiv").css("width", "auto");
}

function winMid() {
	$("#imgdiv").css("background-image", "url(img/bg_m.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.6043" + "px");
	$("#Sbriefdiv").hide();
	$("#Stxtdiv").hide();
	$("#contdiv").show();
	$("#textdiv").css("font-size", $("#imgdiv").width() * "0.02" + "px");
	$("#titlediv").css("font-size", $("#imgdiv").width() * "0.04" + "px");
	$("#contdiv").css("margin", "0");
	$("#contdiv").css("width", "auto");
}

function winMax() {
	$("#imgdiv").css("background-image", "url(img/bg.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.4055" + "px");
	$("#Sbriefdiv").hide();
	$("#Stxtdiv").hide();
	$("#contdiv").show();
	$("#textdiv").css("font-size", $("#imgdiv").width() * "0.02" + "px");
	$("#titlediv").css("font-size", $("#imgdiv").width() * "0.04" + "px");
	$("#contdiv").css("margin", "0");
	$("#contdiv").css("width", "auto");

}

function winLag() {
	$("#imgdiv").css("background-image", "url(img/bg.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.4055" + "px");
	$("#Sbriefdiv").hide();
	$("#Stxtdiv").hide();
	$("#contdiv").show();
	$("#textdiv").css("font-size", "32px");
	$("#contdiv").css("margin", "0 auto");
	$("#contdiv").css("width", "1280");
	$("#titlediv").css("font-size", "64px");

}

function resizefun() {
	$("#contdiv").css("padding-top", ($("#imgdiv").height() - $("#briefdiv").height()) / 2 + "px");
}
