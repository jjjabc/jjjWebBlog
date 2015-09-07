{{range $key, $ja := .jas}}
<div>
	<div style="border-bottom: 3px solid;">
		<div id="title{{$ja.Id}}" class="titleDiv">
			{{$ja.Title}}
		</div>
		<div class="buttonDiv">
			{{if $ja.IsPublished}}
			<input type="button" value="取消发布" onclick="publish('{{$ja.Id}}','unpublish')"/>
			{{else}}
			<input type="button" value="  发布  " onclick="publish('{{$ja.Id}}','publish')"/>
			{{end}}
			<input type="button" value="编辑" onclick="edit('{{$ja.Id}}')"/>
			<input type="button" value="删除" onclick="del('{{$ja.Id}}')"/>
			<span>属于：{{$ja.Category}}</span>
		</div>
	</div>
	<table>
		<tr>
			<td style="width:300px;">
			<div id="imgurl{{$ja.Id}}">

				<img src="{{$ja.Imgurl}}" style="max-width:300px;_width:expression(this.width > 300 ? '300px' : this.width);">
			</div></td>
			<td>
			<div id="text{{$ja.Id}}" style="padding-left: 25px;display:inline-block;">
				{{$ja.Text}}
			</div></td>
		</tr>
	</table>
</div>
{{end}} 