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
			#bgDiv{
				display:none;
	background-color: #000;
	width: 100%;
	height: 100%;
	left:0;
	top:0;/*FF IE7*/
	filter:alpha(opacity=40);/*IE*/
	opacity:0.4;/*FF*/
	z-index:1;
	position:fixed!important;/*FF IE7*/
	position:absolute;/*IE6*/
	_top:       expression(eval(document.compatMode &&
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
			function updataart(){
				ja = {
					title : $("#editTitle").val(),
					text : $("#editTextarea").val(),
					id : $("#editartId").val()
				};
				$.post("/updataArt",ja,updataOver)
			}
			function updataOver(result){
				$("#updataDiv").hide();
				$("#bgDiv").hide();
				checkReturn(result);
			}
			function del(artId) {
				$.post("/delArt", {artId:artId}, checkReturn);

			}

			function addart() {
				ja = {
					title : $("#arttitle").val(),
					text : $("#arttext").val(),
					ispublish : $("#artaction").val()
				};
				$.post("/addArt", ja, checkReturn);
			}

		</script>
		<title>Wicwin.com\admin</title>
	</head>
	<body>
	<div id="bgDiv"></div>
	<div id="updataDiv" style="display:none;z-index:2;position:absolute;top:50%;left:50%;margin:-100px 0 0 -250px;">
	<div>编辑</div>
	<div style="background-color:#ffffff;">
		<div><input id="editTitle" type="text" style="width:100%;"/></div>
		<div><textarea id="editTextarea" rows="5" style="width:500px;"></textarea></div>
		<div  class="buttonDiv"><input id="editartId" style="display:none;">
		<input type="button" value="提交" onclick="updataart();"/>
		<input type="button" value="取消" onclick="$('#updataDiv').hide();$('#bgDiv').hide();"/>
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
			</div>
		</div>
	</body>
</html>