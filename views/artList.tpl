{{range $key, $ja := .jas}}					
					<div>
						<div style="border-bottom: 3px solid;">
							<div class="titleDiv">
								{{$ja.Title}}
							</div>
							<div class="buttonDiv">
								{{if $ja.IsPublished}}
								<input type="button" value="取消发布" onclick="publish('$ja.Id','publish')"/>
								{{else}}
								<input type="button" value="  发布  " onclick="publish('$ja.Id','unpublish')"/>
								{{end}}
								<input type="button" value="编辑" onclick="edit('$ja.Id')"/>
								<input type="button" value="删除" onclick="del('$ja.Id')"/>
							</div>
						</div>
						<div style="padding-left: 25px;">
							{{$ja.Text}}
						</div>
					</div>
{{end}} 