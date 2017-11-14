
    <!-- Page Content -->
    <article>
        <div class="container">
            <div class="row">
                <div class="col-lg-8  col-md-10 " id="post">
                                        <h1>{{ .Page.Title }}</h1>
                    <hr/>
                    <h2>{{ .Page.Tagline }}</h2>

               	{{ renderPost .Page.Content .Page.ContentType false | str2html }}
                </div>
            </div>
        </div>
    </article>



