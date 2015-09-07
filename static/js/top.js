$(document).ready(function() {
	resizefun();
	$(window).resize(function() {
		resizefun();
	});
});
$(window).onload=function(){
      resizefun();
     };
function resizefun() {
	$("#topcontent").css("top", ($("#topcontainer").height() - $("#topcontent").height()) / 2 + "px");
	$("#topcontenttitle").css("top", ($("#topcontainer").height() - $("#topcontenttitle").height()) / 2 + "px");
	$("#navmenu").css("padding-top",(70 - $("#navmenu").height()) / 2 + "px");
}
