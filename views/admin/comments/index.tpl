<div class="col-lg-12"> 
  <div class="comment-list">

    {{ range $key, $comment := .comments }}
    <div class="media">
      <div class="media-body">
	<h3><a href="/admin/comments/edit/{{ $comment.Id }}">
	  {{ $comment.Commenter }}
	</a></h3>
	<p> {{ renderSafeMarkDown $comment.Comment | str2html }}</p>
      </div>
      <div class="btn-group pull-right">
	<a href="/admin/comments/delete/{{ $comment.Id }}" class="btn btn-sm btn-primary">Delete</a>
      </div>

    {{ end }}
      
  </div>
</div>
