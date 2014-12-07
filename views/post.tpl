    <!-- Page Header -->
    <!-- Set your background image for this header on the line below. -->
    <header class="intro-header" style="background-image: url('/static/img/post-bg.jpg')">
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
                    <div class="post-heading">
                        <h1>{{ .Post.Title }}</h1>
                        <h2 class="subheading">{{ .Post.Tagline}}</h2>
                        <span class="meta">Posted by <a href="http://juanfgs.eosweb.info">Juan</a> {{ .Post.CreatedAt}}</span>
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


