<div class="col-lg-12 ">
  <div class="row actions" role="">
    <a href="/admin/posts/new" class="btn btn-sm btn-primary pull-right">New</a>
  </div>
     <div class="post-lists">
      {{ range $key, $post := .posts }}
      <div class="media">
	<div class="media-body">
	  <a href="/admin/posts/edit/{{ $post.Id }}">
	    <h2 class="post-title">
	      {{ $post.Title }}
	    </h2>

	  </a>
	  <p>	
	    {{ str2html $post.Tagline }}
	  </p>
	</div>	  
	  <div class="pull-right btn-group" role="group">
	    <a href="/admin/posts/delete/{{ $post.Id }}" class="btn btn-sm btn-default">Delete</a>
	    <a href="/admin/posts/edit/{{ $post.Id }}" class="btn btn-sm btn-default">Edit</a>	    
	  </div>
	  <p class="post-meta"><small>Posted by <a href="#">Juan</a> on {{ .CreatedAt}}</small></p>

      </div>
      <hr>

	{{ end }}
      </div>

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
