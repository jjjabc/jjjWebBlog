function showList(id){
	$.get("/artList", function(result) {
		$(id).html(result);
	});
}

function refList() {
	$.get("/artList", function(result) {
		$("#allList").html(result);
	});
}

function checkReturn(result) {
	if (result != "OK") {
		$("#allList").html(result);
	} else {
		refList();
	}
}

function publish(artId, publish) {
	$.post("/publishArt", {
		artId : artId,
		Action : publish
	}, checkReturn);
}

function edit(artId) {
	$("#editTextarea").val($("#text" + artId).text());
	$("#editTitle").val($("#title" + artId).text());
	$("#editartId").val(artId);
	$("#bgDiv").show();
	$("#updataDiv").show();
}

function updataart() {
	ja = {
		title : $("#editTitle").val(),
		text : $("#editTextarea").val(),
		id : $("#editartId").val(),
		imgurl : $("#imgurl1").val()
	};
	$.post("/updataArt", ja, updataOver);
}

function updataOver(result) {
	$("#updataDiv").hide();
	$("#bgDiv").hide();
	checkReturn(result);
}

function del(artId) {
	$.post("/delArt", {
		artId : artId
	}, checkReturn);

}

function addart() {
	ja = {
		title : $("#arttitle").val(),
		text : $("#arttext").val(),
		ispublish : $("#artaction").val(),
		imgurl : $("#imgurl").val()
	};
	$.post("/addArt", ja, checkReturn);
}

function UpladFile() {
	var fileObj = document.getElementById("file").files[0];
	// js 获取文件对象
	var FileController = "/upload";
	// 接收上传文件的后台地址
	// FormData 对象
	var form = new FormData();
	//form.append("author", "hooyes");
	// 可以增加表单数据
	form.append("file", fileObj);
	// 文件对象
	// XMLHttpRequest 对象
	var xhr = new XMLHttpRequest();
	xhr.open("post", FileController, true);
	xhr.onload = function() {
		var imgDiv = document.getElementById("imgpreview");
		imgDiv.innerHTML = "<img src='" + xhr.responseText + "' class='img-responsive'>";
		$("#imgurl").val(xhr.responseText);
	};
	xhr.upload.addEventListener("progress", progressFunction, false);
	xhr.send(form);
}

function progressFunction(evt) {

	var progressBar = document.getElementById("progressBar");

	var percentageDiv = document.getElementById("percentage");

	if (evt.lengthComputable) {

		progressBar.max = evt.total;

		progressBar.value = evt.loaded;

		percentageDiv.innerHTML = Math.round(evt.loaded / evt.total * 100) + "%";

	}

}

function UpladFile1() {
	var fileObj = document.getElementById("file1").files[0];
	// js 获取文件对象
	var FileController = "/upload";
	// 接收上传文件的后台地址
	// FormData 对象
	var form = new FormData();
	//form.append("author", "hooyes");
	// 可以增加表单数据
	form.append("file", fileObj);
	// 文件对象
	// XMLHttpRequest 对象
	var xhr = new XMLHttpRequest();
	xhr.open("post", FileController, true);
	xhr.onload = function() {
		var imgDiv = document.getElementById("imgpreview1");
		imgDiv.innerHTML = "<img src='" + xhr.responseText + "' class='img-responsive'>";
		$("#imgurl1").val(xhr.responseText);
	};
	xhr.upload.addEventListener("progress", progressFunction1, false);
	xhr.send(form);
}

function progressFunction1(evt) {

	var progressBar = document.getElementById("progressBar1");

	var percentageDiv = document.getElementById("percentage1");

	if (evt.lengthComputable) {

		progressBar.max = evt.total;

		progressBar.value = evt.loaded;

		percentageDiv.innerHTML = Math.round(evt.loaded / evt.total * 100) + "%";

	}

}