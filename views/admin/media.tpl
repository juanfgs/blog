<div class="col-lg-10 col-lg-offset-1"> 
     <ul class="list-group">
      {{ range $key, $media := .media }}

      	  <li class="list-group-item">
	  <h2 class="media-title">
	    {{ $media.Filename }}
	  </h2>

	  <a href="/admin/media/delete/{{ $media.Id }}" class="btn btn-primary">Delete media</a>

      <hr>
      </li> 
      {{ end }}
      </ul>

{{if .paginator.HasPages}}
<ul class="pagination pagination">
    {{if .paginator.HasPrev}}
        <li><a href="{{.paginator.PageLinkFirst}}">First</a></li>
        <li><a href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
    {{else}}
        <li class="disabled"><a>First</a></li>
        <li class="disabled"><a>&laquo;</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <li{{if $.paginator.IsActive .}} class="active"{{end}}>
            <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </li>
    {{end}}
    {{if .paginator.HasNext}}
        <li><a href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
        <li><a href="{{.paginator.PageLinkLast}}">Last</a></li>
    {{else}}
        <li class="disabled"><a>&raquo;</a></li>
        <li class="disabled"><a>Last</a></li>
    {{end}}
</ul>
{{end}}

</div>
