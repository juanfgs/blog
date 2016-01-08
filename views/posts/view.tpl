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
                            <textarea class="form-control" name="comment"rows="3"></textarea>
		      </div>
		      <input type="submit" class="btn btn-sm btn-primary pull-right" value="Save" />		      
		      <div class="g-recaptcha" data-sitekey="6LcBH_oSAAAAAPCKHLRUVglylkLo2xQ_bOycnYow"></div>
		      <input type="hidden" value="{{ .Post.Id }}" name="post_id" />
                      
                    </form>
                </div>

                <hr>

                <!-- Posted Comments -->

                <!-- Comment -->
		{{ range $key, $comment := .Comments}}
                <div class="media">
		  
                    <div class="media-body">
                        <h4 class="media-heading">{{ $comment.Commenter}}
                            <small>{{ $comment.CreatedAt}}</small>
                        </h4>
			{{ renderSafeMarkDown $comment.Comment | str2html }}
			<a href="#"  class=" btn btn-sm btn-default reply" data-id="{{ $comment.Id }}">Reply</a>
                    </div>
                </div>
		{{ end }}
	</div>

      </div>
    </div>
{{ end }}

<script src="/static/js/comment.js"></script>
