<div class="col-lg-10 col-lg-offset-1"> 
     <ul class="list-group">
      {{ range $key, $post := .posts }}

      	  <li class="list-group-item">
	<a href="/post/{{ $post.Id }}">
	  <h2 class="post-title">
	    {{ $post.Title }}
	  </h2>
	  <h3 class="post-subtitle">
	    {{ str2html $post.Content }}
	  </h3>
	</a>
	<p class="post-meta">Posted by <a href="#">Juan</a> on {{ .CreatedAt}}</p>

      <hr>
      </li> 
      {{ end }}
      </ul>
</div>