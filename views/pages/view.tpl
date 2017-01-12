
    <!-- Page Content -->
    <article>
        <div class="container">
            <div class="row">
                <div class="col-lg-8  col-md-10 " id="post">
               	{{ renderPost .Page.Content .Page.ContentType | str2html }}
                </div>
            </div>
        </div>
    </article>



