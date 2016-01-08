<div class="col-lg-12"> 
  <div class="comment-list">

    {{ range $key, $comment := .comments }}
    <div class="media">
      <div class="btn-group pull-right">
	<a href="/admin/comments/edit/{{ $comment.Id }}" class="btn btn-sm btn-default">Edit</a>
	<a href="/admin/comments/delete/{{ $comment.Id }}" class="btn btn-sm btn-primary">Delete</a>	
      </div>            
      <div class="media-body">
	
	<p>Author:{{ $comment.Commenter }}</p>

	
	<p> {{ renderSafeMarkDown $comment.Comment | str2html }}</p>
	<span class="small">on <a href="/admin/posts/edit/{{$comment.Post.Id}}">"{{ $comment.Post.Title }}"</a></span>

	
      </div>
      <hr />
    </div>
    {{ end }}
      
  </div>
</div>
