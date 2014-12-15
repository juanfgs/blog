<div class="col-lg-10 col-lg-offset-1"> 
     <ul class="list-group">
      {{ range $key, $post := .posts }}

      	  <li class="list-group-item">
	<a href="/admin/post/edit/{{ $post.Id }}">
	  <h2 class="post-title">
	    {{ $post.Title }}
	  </h2>

	</a>
		<p>	
		    {{ str2html $post.Tagline }}
	    	  </p>
		  <a href="/admin/post/delete/{{ $post.Id }}" class="btn btn-primary">Delete post</a>
	<p class="post-meta">Posted by <a href="#">Juan</a> on {{ .CreatedAt}}</p>

      <hr>
      </li> 
      {{ end }}
      </ul>
</div>
