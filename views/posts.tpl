
<header class="intro-header" style="background-image: url('/static/img/home-bg.jpg')">
  <div class="container">
    <div class="row">
      <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
        <div class="site-heading">
          <h1>{{ .HeroTitle }}</h1>
          <hr class="small">
          <span class="subheading">{{.HeroTagline }}</span>
        </div>
      </div>
    </div>
  </div>
</header>

<!-- Main Content -->
<div class="container">
  <div class="row">
    <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
      {{ range $key, $post := .posts }}
      <div class="post-preview">
	{{ if $post.Slug }}
	<a href="/post/{{ $post.Slug }}">	
	  {{ else }}
	  <a href="/post/{{ $post.Id }}">	  
	{{ end }}
	  <h2 class="post-title">
	    {{ $post.Title }}
	  </h2>	
	  </a>
	  <p>

	    {{ renderPost $post.Content $post.ContentType | str2html   }}
	    <div class="btn-group-sm" role="group">
	      <a  class="btn  btn-default" href="/post/{{ $post.Id }}">Leer Más</a>
	      <a class="btn  btn-primary" href="/post/{{ $post.Id }}"> {{ len $post.Comments }}</a>	    
	    </div>
	</p>

	<p class="post-meta">Posted by <a href="#">Juan</a> on {{ .CreatedAt}}</p>
      </div>
      <hr>
      {{ end }}
      <!-- Pager -->

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
  </div>
</div>

<hr>




