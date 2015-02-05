<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />

		<style type="text/css">
			.buttonDiv {
				display: inline-block;
				float: right;
			}
			.titleDiv {
				display: inline-block;
			}
			#bgDiv {
				display: none;
				background-color: #000;
				width: 100%;
				height: 100%;
				left: 0;
				top: 0;/*FF IE7*/
				filter: alpha(opacity=40);/*IE*/
				opacity: 0.4;/*FF*/
				z-index: 1;
				position: fixed !important;/*FF IE7*/
				position: absolute;/*IE6*/
				_top: expression(eval(document.compatMode &&
				document.compatMode=='CSS1Compat') ?
				documentElement.scrollTop + (document.documentElement.clientHeight-this.offsetHeight)/2 :/*IE6*/
				document.body.scrollTop + (document.body.clientHeight - this.clientHeight)/2);/*IE5 IE5.5*/
			}
		</style>
		<script src="static/js/jquery-1.11.2.js"></script>

		<script>
			$(document).ready(function() {
				refList();
			});

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
					imgDiv.innerHTML = "<img src='" + xhr.responseText + "' style='max-width:300px;_width:expression(this.width > 300 ? '300px' : this.width);'>";
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
					imgDiv.innerHTML = "<img src='" + xhr.responseText + "' style='max-width:300px;_width:expression(this.width > 300 ? '300px' : this.width);'>";
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
		</script>
		<title>Wicwin.com\admin</title>
	</head>
	<body>
		<div id="bgDiv"></div>
		<div id="updataDiv" style="display:none;z-index:2;position:absolute;top:50%;left:50%;margin:-100px 0 0 -250px;">
			<div>
				编辑
			</div>
			<div style="background-color:#ffffff;">
				<div>
					<input id="editTitle" type="text" style="width:100%;"/>
				</div>
				<div>
					<textarea id="editTextarea" rows="5" style="width:500px;"></textarea>
				</div>
				<div  class="buttonDiv">
					<input id="editartId" style="display:none;">
					<input type="button" value="提交" onclick="updataart();"/>
					<input type="button" value="取消" onclick="$('#updataDiv').hide();$('#bgDiv').hide();"/>
				</div>

				<div id="uploadDiv1">
					<div>
						<input id="file1" type="file" />
					</div>
					<div>
						<input type="button" onclick="UpladFile1()" value="上传" />
						<input id="imgurl1" style="display:none;" value=""/>
					</div>
					<progress id="progressBar1" value="0" max="100"></progress>
					<div id="percentage1"></div>

					<div id="imgpreview1"></div>
				</div>

			</div>
		</div>
		<div>
			<div>
				<div>
					文章列表
				</div>

				<div id="allList">

				</div>

			</div>
			<div style="border: 1px solid;">
				<div>
					新增加文章
				</div>
				<div>
					标题：
					<input id="arttitle" type="text" />
					<br>
					<div>
						内容：
					</div>
					<textarea id="arttext" rows="5"></textarea>
				</div>
				<div>
					<input type="checkbox" id="artaction" />
					提交并发布
				</div>
				<div>
					<input type="button" value="提交" onclick="addart();"/>
				</div>

				<div id="uploadDiv">
					<div>
						<input id="file" type="file" />
					</div>
					<div>
						<input type="button" onclick="UpladFile()" value="上传" />
						<input id="imgurl" style="display:none;" value=""/>
					</div>
					<progress id="progressBar" value="0" max="100"></progress>
					<div id="percentage"></div>

					<div id="imgpreview"></div>
				</div>

			</div>
		</div>
	</body>
</html>