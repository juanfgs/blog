<div class="col-lg-10 col-lg-offset-1"> 
     <ul class="list-group">
      {{ range $key, $comment := .comments }}

      	  <li class="list-group-item">
	<a href="/admin/comments/edit/{{ $comment.Id }}">

	    {{ $comment.Commenter }}


	</a>
	<p> {{ renderSafeMarkDown $comment.Comment | str2html }}</p>
			  <a href="/admin/comments/delete/{{ $comment.Id }}" class="btn btn-primary">Delete</a>
      </li> 
      {{ end }}
      </ul>
</div>
