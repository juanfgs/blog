    <!-- Page Header -->
    <!-- Set your background image for this header on the line below. -->
    <header class="intro-header" style="background-image: url('/static/img/post-bg.jpg')">
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
                    <div class="post-heading">
                        <h1>{{ .Post.Title }}</h1>
                        <h2 class="subheading">{{ .Post.Tagline}}</h2>
                        <span class="meta">Posted by <a href="http://juanfgs.eosweb.info">{{ .Post.Author.Username}}</a> {{ .Post.CreatedAt}}</span>
                    </div>
                </div>
            </div>
        </div> 
    </header>

    <!-- Post Content -->
    <article>
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1" id="post">
               	{{ renderPost .Post.Content .Post.ContentType | str2html }}
                </div>
            </div>
        </div>
    </article>

 <hr>

{{ if .EnableComments }}
    <div class="container">
     
      <div class="row">
	<div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
                <!-- Comments Form -->
                <div class="well">
                    <h4>Leave a Comment:</h4>
                    <form action="/comment/new/" method="POST" role="form">
                      <div class="form-group">
			<input class="form-control" type="text" value="" placeholder="Your name..." name="commenter" />

		      </div>
		      <div  class="form-group">
                            <textarea class="form-control" name="comment" rows="3"></textarea>
		      </div>
		      <input type="submit" class="btn btn-sm btn-primary pull-right" value="Save" />		      
		      <div id="recaptcha1"></div>
		      <input type="hidden" value="{{ .Post.Id }}" name="post_id" />
                      
                    </form>
                </div>

                <hr>

                <!-- Posted Comments -->

                <!-- Comment -->
		{{ range $key, $comment := .Comments}}
                <div class="media" {{ if $comment.Parent }} data-parent-id="{{$comment.Parent.Id}}" {{ end }} data-comment-id="{{$comment.Id}}">
		  
                    <div class="media-body">
                        <h4 class="media-heading">{{ $comment.Commenter}}
                            <small>{{ $comment.CreatedAt}}</small>
                        </h4>
			{{ renderSafeMarkDown $comment.Comment | str2html }}
			{{ if $comment.Parent }} 			
			{{ else }}
			<a href="#"  class="btn-sm btn-primary  reply" data-toggle="modal" data-target="#commentBox" data-id="{{ $comment.Id }}">Reply</a>

			{{ end }}
                    </div>
                </div>
		
		{{ end }}
	</div>

      </div>
    </div>
{{ end }}

    
<div class="modal fade" tabindex="-1" role="dialog" id="commentBox">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title">Reply to comment</h4>
      </div>
      <div class="modal-body">
	<form action="/comment/new/" method="POST" role="form">
          <div class="form-group">
	    <input class="form-control" type="text" value="" placeholder="Your name..." name="commenter" />
	    
	  </div>
	  <div  class="form-group">
            <textarea class="form-control" name="comment" rows="b3"></textarea>
	  </div>
	  
	  <div id="recaptcha2"></div>
	  
	  <input type="hidden" value="{{ .Post.Id }}" name="post_id" />
	  <input type="hidden" value="" name="parent_id"  id="commentId"/>	  
          
        </form>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-sm btn-default" data-dismiss="modal">Close</button>
        <button type="button" class="btn btn-sm btn-primary" id="sendComment">Comment</button>
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

<script type="text/javascript">
  var recaptcha1;
  var recaptcha2;  
  var loadRecaptcha = function() {
  recaptcha1 = grecaptcha.render('recaptcha1', {
  'sitekey': '6LcBH_oSAAAAAPCKHLRUVglylkLo2xQ_bOycnYow',
  'theme': 'light'
  });
  recaptcha2 = grecaptcha.render('recaptcha2', {
  'sitekey': '6LcBH_oSAAAAAPCKHLRUVglylkLo2xQ_bOycnYow',
  'theme': 'light'
  });  
  };
</script>
