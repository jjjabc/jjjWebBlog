function winMin() {
	$("#imgdiv").css("background-image", "url(static/img/bg_min.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.6043" + "px");
	$("#Sbriefdiv").show();
	$("#Stxtdiv").show();
	$("#contdiv").hide();
	$("#mobile").show();
	$("#headdiv").hide();
	$("#Stitlediv").css("font-size", $("#imgdiv").height() * "0.3" + "px");
	$("#Stitlediv").css("padding-left", (($("#imgdiv").width() * 0.05) - 5) + "px");
	$("#Stitlediv").css("padding-top", ($("#imgdiv").height() * 0.2) + "px");
	$("#Stxtdiv").css("padding-left", (($("#imgdiv").width() * 0.05) - 5) + "px");
	$("#Slogomenudiv").css("margin-left", $("#Stxtdiv").css("padding-left"));
	$("#Slogomenudiv").css("margin-right", $("#Stxtdiv").css("padding-left"));
	$("#contIn").css("margin-right", $("#Stxtdiv").css("padding-left"));
	$("#contIn").css("margin-left", $("#Stxtdiv").css("padding-left"));
	$("#listTitle").css("font-size", $("#textdiv").css("font-size"));
}

function winSmall() {
	$("#imgdiv").css("background-image", "url(static/img/bg_s.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() + "px");

	$("#textdiv").css("font-size", $("#imgdiv").width() * "0.03" + "px");
	$("#titlediv").css("font-size", $("#imgdiv").width() * "0.06" + "px");
	$("#contdiv").css("margin", "0");
	$("#contdiv").css("width", "auto");
	notSmall();
}

function winMid() {
	$("#imgdiv").css("background-image", "url(static/img/bg_m.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.5409" + "px");

	$("#textdiv").css("font-size", $("#imgdiv").width() * "0.02" + "px");
	$("#titlediv").css("font-size", $("#imgdiv").width() * "0.04" + "px");
	$("#contdiv").css("margin", "0");
	$("#contdiv").css("width", "auto");
	notSmall();
}

function winMax() {
	$("#imgdiv").css("background-image", "url(static/img/bg.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.4055" + "px");
	$("#textdiv").css("font-size", $("#imgdiv").width() * "0.02" + "px");
	$("#titlediv").css("font-size", $("#imgdiv").width() * "0.04" + "px");
	$("#contdiv").css("margin", "0");
	$("#contdiv").css("width", "auto");
	notSmall();
}

function winLag() {
	$("#imgdiv").css("background-image", "url(static/img/bg.jpg)");
	$("#imgdiv").css("height", $("#imgdiv").width() * "0.4055" + "px");
	$("#textdiv").css("font-size", "32px");
	$("#contdiv").css("margin", "0 auto");
	$("#contdiv").css("width", "1280");
	$("#titlediv").css("font-size", "64px");
	notSmall();
}

function resizefun() {
	$("#contdiv").css("padding-top", ($("#imgdiv").height() - $("#briefdiv").height()) / 2 + "px");
}

function notSmall() {
	$("#Sbriefdiv").hide();
	$("#Stxtdiv").hide();
	$("#mobile").hide();
	$("#headdiv").show();
	$("#contdiv").show();
	$("#headdiv").css("margin-left", $("#briefdiv").css("padding-left"));
	$("#headdiv").css("margin-right", $("#briefdiv").css("padding-left"));
	$("#contIn").css("margin-right", $("#briefdiv").css("padding-left"));
	$("#contIn").css("margin-left", $("#briefdiv").css("padding-left"));
	$("#listTitle").css("font-size", $("#textdiv").css("font-size"));
}

