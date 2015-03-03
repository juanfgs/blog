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
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
               	{{ str2html .Post.Content}}
                </div>
            </div>
        </div>
    </article>

 <hr>


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
		      <div class="g-recaptcha" data-sitekey="6LcBH_oSAAAAAPCKHLRUVglylkLo2xQ_bOycnYow"></div>
		      <input type="hidden" value="{{ .Post.Id }}" name="post_id" />
                        <input type="submit" class="btn btn-primary" value="Save" />
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
			{{ html2str $comment.Comment }}
                    </div>
                </div>
		{{ end }}
	</div>

      </div>
      </div>

