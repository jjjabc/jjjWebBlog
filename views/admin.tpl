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
		</style>
		<script src="static/js/jquery-1.11.2.js"></script>

		<script>
			$(document).ready(function() {
				refList();
			});

			function refList() {
				$.get("/artList", function(result) {
					$(allList).html(result);
				});
			}

			function checkReturn(result) {
				if (result != "OK") {
					$(allList).html(result);
				} else {
					refList();
				}
			}

			function publish(artId, publish) {
				$.post("/publishArt", {
					artId : artId,
					Action : publish
				}, checkReturn());
			}

			function editart(artId) {
				$("editTextarea").html($("artId" + artId).html());
				$("editTitle").html($("artId" + artId).html());
			}

			function del(artId) {
				$.post("/delArt", artId, checkReturn());

			}

			function addart() {
				ja = {
					title : $("title").text(),
					text : $("text").text(),
					ispublish : $("action").checked
				};
				$.post("/addArt", ja, checkReturn());
			}

		</script>
		<title>Wicwin.com\admin</title>
	</head>
	<body>
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
					<input id="title" type="text" />
					<br>
					<div>
						内容：
					</div>
					<textarea id="text" rows="5"></textarea>
				</div>
				<div>
					<input type="checkbox" id="action" />
					提交并发布
				</div>
				<div>
					<input type="button" value="提交" onclick="addart();"/>
				</div>
			</div>
		</div>
	</body>

</html>