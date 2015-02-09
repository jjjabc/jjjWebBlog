function artwinMin() {

	artnotSmall();
	artimgSmall();
}

function artwinSmall() {
	artimgSmall();
	artnotSmall();
}

function artwinMid() {
	artimgBig();

	artnotSmall();
}

function artwinMax() {
	artimgBig();

	artnotSmall();
}

function artwinLag() {
	artimgBig();
	artnotSmall();
}

function artnotSmall() {
	$("#arttextimg").css("margin-right", $("#headdiv").css("padding-left"));
	$("#arttextimg").css("margin-left", $("#headdiv").css("padding-left"));
}

function artimgBig() {
	var artimg = $("#artimg");
	alert(artimg.width());
	var bl = artimg.heigth() / artimg.width();
	$("#text-under-image-block").css("width", "67%");
	artimg.width = $("#text-under-image-block").width;
	artimg.hight = artimg.width * bl;
}

function artimgSmall() {
	var artimg = $("#artimg");
	var bl = artimg.height / artimg.width;
	artimg.css("width", "100%");
	artimg.hight = artimg.width * bl;
}

