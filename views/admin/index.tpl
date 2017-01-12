<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">


    <title>{{ .Title }}</title>


    <link href="/static/css/bootstrap.admin.min.css" rel="stylesheet">
    <link href="/static/css/admin.css" rel="stylesheet">

    <!-- Custom CSS -->


    <!-- Custom Fonts -->
    <link href="http://maxcdn.bootstrapcdn.com/font-awesome/4.1.0/css/font-awesome.min.css" rel="stylesheet" type="text/css">
    <link href='http://fonts.googleapis.com/css?family=Lora:400,700,400italic,700italic' rel='stylesheet' type='text/css'>
    <link href='http://fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,800italic,400,300,600,700,800' rel='stylesheet' type='text/css'>


    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

</head>

<body>

    <div id="wrapper">
      {{  if .User }}
        <!-- Navigation -->
        <nav class="navbar navbar-inverse navbar-static-top" role="navigation" style="margin-bottom: 0">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand" href="/admin/">Blog</a>
            </div>
            <!-- /.navbar-header -->

            <ul class="nav navbar-top-links navbar-right">
                <!-- /.dropdown -->
                <li class="dropdown">
                    <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                        <i class="fa fa-bell fa-fw"></i>
			<i class="fa fa-caret-down">
			    
			</i>
                    </a>
                    <ul class="dropdown-menu dropdown-alerts">
			{{ range $key, $notification := .Notifications}}
                        <li>
                                <div>
                                    <i class="fa fa-comment fa-fw"></i> <a href="{{ $notification.Message }}">New Comment </a>
                                    <span class="pull-right text-muted small">{{ $notification.CreatedAt}}</span>
                                </div>
                        </li>
			<li class="divider"></li>
			{{ end }}
                        <li>
                            <a class="text-center" href="">
                                <strong>See All Alerts ({{ .NotificationCount }})</strong>
                                <i class="fa fa-angle-right"></i>
                            </a>
                        </li>
                    </ul>
                    <!-- /.dropdown-alerts -->
                </li>
                <!-- /.dropdown -->
                <li class="dropdown">
                    <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                        <i class="fa fa-user fa-fw"></i>  <i class="fa fa-caret-down"></i>
                    </a>
                    <ul class="dropdown-menu dropdown-user">
                        <li><a href="#"><i class="fa fa-user fa-fw"></i> User Profile</a>
                        </li>
                        <li><a href="/admin/settings"><i class="fa fa-gear fa-fw"></i> Settings</a>
                        </li>
                        <li class="divider"></li>
                        <li><a href="/logout"><i class="fa fa-sign-out fa-fw"></i> Logout</a>
                        </li>
                    </ul>
                    <!-- /.dropdown-user -->
                </li>
                <!-- /.dropdown -->
            </ul>
            <!-- /.navbar-top-links -->

            <div class="navbar-default sidebar" role="navigation">
                <div class="sidebar-nav navbar-collapse">
                    <ul class="nav" id="side-menu">
                        <li>
                            <a href="/admin/"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
                        </li>
                        <li>
                            <a href="#"><i class="fa fa-archive fa-fw"></i>CMS<span class="fa arrow"></span></a>
                            <ul class="nav nav-second-level">
                              <li>
                                <a href="#"><i class="fa fa-bars fa-fw"></i>Blog <span class="fa arrow"></span></a>				

				  <ul class="nav nav-third-level">
                                        <li>
                                            <a href="/admin/posts"><i class="fa fa-file fa-fw"></i>Posts</a>
                                        </li>
                                        <li>
                                            <a href="/admin/categories"><i class="fa fa-folder fa-fw"></i>Categories</a>
                                        </li>
					<li>
					  <a href="/admin/comments"><i class="fa fa-comments fa-fw"></i>Comments</a>
					</li>
                                        <li>
                                            <a href="/admin/posts/export"><i class="fa fa-file fa-fw"></i>Export Posts</a>
                                        </li>					
				  </ul>
                                </li>
                                <li>
                                    <a href="/admin/pages"><i class="fa fa-globe fa-fw"></i>Pages</a>
                                </li>
                            </ul>
                            <!-- /.nav-second-level -->
                        </li>
                        <li>
                          <a href="/admin/settings"><i class="fa fa-cog fa-fw"></i>Settings</a>
                        </li>			
                    </ul>
                </div>
                <!-- /.sidebar-collapse -->
            </div>
            <!-- /.navbar-static-side -->
        </nav>
	{{ end }}
        <div id="page-wrapper" class="row">
   <!-- Page Header -->
    {{ if .flash.error  }}
    <div class="col-lg-12" id="ErrorArea">
      <div class="well col-md-6 col-md-offset-3">
	{{.flash.error}}
      </div>
    </div>
    {{ end }}

    {{ if .flash.notice  }}
    <div class="col-lg-12" id="ErrorArea">
      <div class="well col-md-6 col-md-offset-3">
	{{.flash.notice}}
      </div>
    </div>
    {{ end }}

    {{ if .flash.warning  }}
    <div class="col-lg-12" id="ErrorArea">
      <div class="well col-md-6 col-md-offset-3">
	{{.flash.warning}}
      </div>
    </div>
    {{ end }}


    {{ .LayoutContent }}
	  
        </div>
        <!-- /#page-wrapper -->
      </div>
    <!-- /#wrapper -->
    <!-- jQuery -->
    <script src="/static/js/jquery.js"></script>

    <!-- Bootstrap Core JavaScript -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/admin.js"></script>
    <script type="text/javascript">
      $(document).ready(function(){

      $('.dropdown-toggle').dropdown();
      $('.collapse').collapse()
      });

    </script>
  </body>
  

</html>
