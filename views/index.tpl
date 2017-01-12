<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    {{ if .Post}}
    <meta name="description" content=" {{.Post.Description }}">
    <meta name="keywords" content="{{ .Post.Keywords }}">
    {{ else   }}
    <meta name="description" content="My personal blog, it's main topics are programming, my personal life, and some food for thought.">
    <meta name="keywords" content="Golang, Go, Blog, Laravel, PHP, MVC, Javascript, JSON, NoSQL, MongoDB, Web Developer">
    {{ end}}
    <meta name="author" content="Juan F. Giménez Silva">
      <link href="/feed.xml" rel="alternate" type="application/rss+xml" title="{{ .Title }}" />
    <title>{{ .Title }}</title>

    <!-- Bootstrap Core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link href="/static/css/prism.css" rel="stylesheet">	      
      <link href="/static/css/style.css" rel="stylesheet">


    <!-- Custom Fonts -->
    <link href="http://maxcdn.bootstrapcdn.com/font-awesome/4.1.0/css/font-awesome.min.css" rel="stylesheet" type="text/css">
    <link href="https://fonts.googleapis.com/css?family=Crimson+Text|Libre+Baskerville|Inconsolata|Raleway" rel="stylesheet"> 
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <script src='https://www.google.com/recaptcha/api.js?onload=loadRecaptcha&render=explicit'></script>
</head>

<body>

    <!-- Navigation -->
    <nav class="navbar navbar-default ">
        <div class="container-fluid">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header page-scroll">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                    <span class="sr-only">Activar Navegación</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>

                <a href="/" class="navbar-brand">Inicio</a>

            </div>

            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav navbar-right">


                    <li>
                        <a href="#search">Búsqueda</a>
                    </li>

		    {{ if .User }}
                    <li>
                        <a href="/admin/">Bienvenido {{ .User.Username }}</a>
                    </li>
		    {{ else }}
		    <li>
                        <a href="/login/">Iniciar Sesión</a>
                    </li>
		    {{ end }}
                </ul>
            </div>
            <!-- /.navbar-collapse -->
        </div>
        <!-- /.container -->
    </nav>

    <!-- Page Header -->
 <div class="container-fluid">
     <div class="row">

         <div class=" col-lg-2 col-md-1">
             <div class="wrapper ">

                 <ul class="nav nav-stacked" id="sidebar">
                     <li class="well"><a href="#author">Juan F. Giménez Silva</a>
                         <img src="/static/img/author.png" class="img-circle center" id="author"/>
                         <p class="author-bio">Hola, bienvenidos a mi blog. Aquí encontrarán articulos sobre informática, programación y política.</p>
                     </li> 

              <li><a href="#sidebar-search">Búsqueda</a>

                  <form action="/search/" method="GET" class="form-inline search" id="sidebar-search">
                      <div class="form-group">
                          <div class="input-group">
                              <input type="text" value="" name="imnotahuman" style="display:none;" />	
                              <input type="search" value=""  class="form-control" name="keyword" placeholder="palabra(s) clave(s)" />
                              <div class="input-group-addon"><input type="submit" id="search_button" value="Buscar"></div>
                          </div>
                      </div>
                  </form>
        
              </li>
                   <li> <a href="#pages">Páginas</a>
		      <ul id="pages">
			{{ range $key, $page :=  .Pages}}
			<li><a href="/page/{{ $page.Slug }}">{{ $page.Title }}</a></li>
			{{ end }}
		      </ul>

                    </li>

                    <li><a href="#categories">Categorías</a>
		      <ul id="categories">
			{{ range $key, $category :=  .Categories}}
			<li><a href="/categories/{{ $category.Id }}">{{ $category.Title }}</a></li>
			{{ end }}
		      </ul>

                    </li>
                    <li><a href="#links">Enlaces</a>
		      <ul id="links">
			{{ range $key, $link :=  .Links}}
			<li><a href="{{ $link.Url }}">{{ $link.Title }}</a></li>
			{{ end }}
		      </ul>
                    </li>
          </ul>
      </div>
      </div>
      {{ .LayoutContent }}
  </div>
 </div>
    <!-- Set your background image for this header on the line below. -->
    <div id="search">
      <button type="button" class="close">×</button>
      <form action="/search/" method="GET">
        <input type="search" value="" name="keyword" placeholder="ingrese palabra(s) clave(s)" />
        <input type="text" value="" name="imnotahuman" style="display:none;" />	
        <input type="submit" class="btn btn-primary" value="Buscar">
      </form>
    </div>

    <!-- Footer -->
    <footer>
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
                    <ul class="list-inline text-center">
                        <li>
                            <a href="http://twitter.com/juanfgs">
                                <span class="fa-stack fa-lg">
                                    <i class="fa fa-circle fa-stack-2x"></i>
                                    <i class="fa fa-twitter fa-stack-1x fa-inverse"></i>
                                </span>
                            </a>
                        </li>
                        <li>
                            <a href="http://facebook.com/juanfgs">
                                <span class="fa-stack fa-lg">
                                    <i class="fa fa-circle fa-stack-2x"></i>
                                    <i class="fa fa-facebook fa-stack-1x fa-inverse"></i>
                                </span>
                            </a>
                        </li>
                        <li>
                            <a href="http://github.com/juanfgs">
                                <span class="fa-stack fa-lg">
                                    <i class="fa fa-circle fa-stack-2x"></i>
                                    <i class="fa fa-github fa-stack-1x fa-inverse"></i>
                                </span>
                            </a>
                        </li>
                    </ul>
                    <p class="copyright text-muted">Copyright &copy; Juan Francisco Gim&eacute;nez Silva 2016</p>
                </div>
            </div>
        </div>
    </footer>

    <!-- jQuery -->
    <script src="/static/js/jquery.js"></script>

    <!-- Bootstrap Core JavaScript -->
    <script src="/static/js/bootstrap.min.js"></script>
    

    <script src="/static/js/prism.js"></script>
    
    <script src="//cdnjs.cloudflare.com/ajax/libs/KaTeX/0.3.0/katex.min.js"></script>
  <script>
        $(document).ready(function () {
            $('script[type="text/katex"]').each(function () {
                var math = $(this).text();
                var span = document.createElement('span');
                katex.render(math, span);
                $(this).replaceWith(span);
            })
        });
    </script>


    

    <!-- Custom Theme JavaScript -->
    <script src="/static/js/clean-blog.js"></script>
    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

      ga('create', 'UA-51811716-1', 'auto');
      ga('send', 'pageview');

    </script>
</body>

</html>
