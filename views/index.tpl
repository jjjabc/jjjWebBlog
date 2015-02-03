<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
"http://www.w3.org/TR/html4/loose.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<meta name="viewport" content="width=device-width,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0">
		<link href="/static/css/jjj.css" rel="stylesheet" type="text/css">
		<script src="/static/js/jquery-1.11.2.js"></script>
		<script src="/static/js/index.js"></script>
		<script>
			$(document).ready(function() {
				imgresize();
				$(window).resize(function() {
					imgresize();

				});
			});

			$(document).ready(function() {
				imgresize();
				$(window).resize(function() {
					imgresize();

				});
			});
			function imgresize() {
				if ($(window).width() > 1600) {
					winLag();
				} else if ($(window).width() > 1200) {
					winMax();
				} else if ($(window).width() > 1000) {
					winMid();
				} else if ($(window).width() > 640) {
					winSmall();
				} else {
					winMin();
				}
				resizefun();
			}

			function expart(artid) {
				{{range $key, $val := .jas}}
				if (artid == "artId{{$val.Id}}") {
					$("#artId{{$val.Id}}").toggle();
					if ($("#artId{{$val.Id}}").is(":visible")) {
						$("#art{{$val.Id}}").css("background-color", "#ffffff")
						$("#artpuls{{$val.Id}}").html("-");
					} else {
						$("#art{{$val.Id}}").css("background-color", "#E2E2E2");
						$("#artpuls{{$val.Id}}").html("+");
					}
				} else {
					$("#artId{{$val.Id}}").hide();
					$("#art{{$val.Id}}").css("background-color", "#E2E2E2");
					$("#artpuls{{$val.Id}}").html("+");
				}
				{{end}}
			}
		</script>
		<title>Wicwin.com</title>
	</head>
	<body>
		<div id="headdiv">
			<div id="logodiv">
				<img src="/static/img/logo_b_50.png">
			</div>
			<div id="menudiv">
				<a class="menuA" href="/">Home</a>
				<a class="menuA" href="/static/contact_us.html">Conatact us</a>
				<a class="menuA" href="#" onclick="$('html,body').animate({scrollTop:$('#cont').offset().top},1000);">Case</a>
			</div>
		</div>
		<div id="mobile">
			<div id="Slogomenudiv"  style="padding-top: 10px;padding-bottom: 10px;">
				<div id="Slogodiv" style="display: inline-block">
					<img src="/static/img/logo_b_50.png">
				</div>
				<div id="Smenudiv" style="display: inline-block;float:right;">
					<img src="/static/img/icon-menu-black.png" onclick="$('#Smenulist').toggle();">
				</div>
			</div>
			<div id="Smenulist" style="display: none;">
				<div class="Smenu">
					Home
				</div>
				<div class="Smenu">
					Conatact us
				</div>
				<div class="Smenu" onclick="$('html,body').animate({scrollTop:$('#cont').offset().top},1000);">
					Case
				</div>
			</div>

		</div>
		<div id="imgdiv">
			<div id="contdiv">
				<div id="briefdiv">
					<div id="titlediv">
						智信创赢
						<br>
						<br>

					</div>
					<div id="textdiv">
						成都智信创赢科技有限公司是专门从事管理信息系统研发、咨询及服务的高新技术企业，由富有创新力且在企业信息化经验丰富的智慧型高效能团队组成。团队成员均从事企业信息化多年，在企业信息化咨询与设计方面卓有成效。公司秉承“智慧、诚信、创新、共赢”核心文化与价值，立志于打造数字化中国、信息化全球的信息咨询、设计、服务提供商。
					</div>
				</div>

			</div>
			<div id="Sbriefdiv" >
				<div id="Stitlediv">
					智信创赢
				</div>

			</div>
		</div>
		<div id="Stxtdiv">
			<div id="Stextdiv">
				成都智信创赢科技有限公司是专门从事管理信息系统研发、咨询及服务的高新技术企业，由富有创新力且在企业信息化经验丰富的智慧型高效能团队组成。团队成员均从事企业信息化多年，在企业信息化咨询与设计方面卓有成效。公司秉承“智慧、诚信、创新、共赢”核心文化与价值，立志于打造数字化中国、信息化全球的信息咨询、设计、服务提供商。
			</div>
		</div>
		<div id="cont" style="background-color: #E2E2E2">
			<div id="contIn">
				<div id="listCont">
					<div id="listTitle" class="listTitle">
						产品及案例
					</div>
					<div id="artlist">

						{{range $key, $val := .jas}}
						<div id="art{{$val.Id}}" class="art">
							<div>
								<div style="padding-bottom: 5px;cursor:pointer;" onclick="expart('artId{{$val.Id}}');$('html,body').animate({scrollTop:$('#art{{$val.Id}}').offset().top},500);">
									<div id="artpuls{{$val.Id}}" style="display: inline-block;width:12px;text-align:center;">
										+
									</div>
									<div style="display: inline-block;font-size: 18px;">
										{{$val.Title}}
									</div>
								</div>

							</div>
							<div id="artId{{$val.Id}}" style="display:none;">
								<div class="artline">
									<div class="t1"></div>
									<div class="t2"></div>
								</div>
								<div style="padding-left:20px;padding-right:20px;font-size:14px;">
								{{$val.Text}}
								</div>
							</div>
						</div>
						{{end}} 


					</div>
				</div>
			</div>
		</div>
		<div id="footdiv">
			<div id="footcont"><img src="/static/img/logoew26.png" style="height: 26px;display: inline-block;float:right;padding-right:80px"></div>
		</div>
	</body>
</html>
