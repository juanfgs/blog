    <!-- Page Header -->
    <!-- Set your background image for this header on the line below. -->
    <header class="intro-header" style="background-image: url('/static/img/post-bg.jpg')">
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
                    <div class="post-heading">
                        <h1>{{ .Page.Title }}</h1>
                        <h2 class="subheading">{{ .Page.Tagline}}</h2>
                    </div>
                </div>
            </div>
        </div> 
    </header>

    <!-- Page Content -->
    <article>
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1" id="post">
               	{{ renderPost .Page.Content .Page.ContentType | str2html }}
                </div>
            </div>
        </div>
    </article>

 <hr>


