<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">


    <title>{{ .Title }}</title>


  <!-- Bootstrap core CSS-->
  <link href="/static/admin/vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">
  <!-- Custom fonts for this template-->
  <link href="/static/admin/vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">
  <!-- Page level plugin CSS-->
  <link href="/static/admin/vendor/datatables/dataTables.bootstrap4.css" rel="stylesheet">
  <!-- Custom styles for this template-->
  <link href="/static/admin/css/sb-admin.css" rel="stylesheet">
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
    <body class="fixed-nav sticky-footer bg-dark" id="page-top">
            
          {{ if .User }}
         
          <app-root></app-root>
          {{ else }}
          
        <div class="content-wrapper">
          <div class="card card-login mx-auto mt-5">
              <div class="card-header">Login</div>
              <div class="card-body">
                  <form action="/login/process" method="POST">
                      <div class="form-group">
                          <label for="exampleInputEmail1">Email address</label>
                          <input class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email" name="Username">
                          </div>
                          <div class="form-group">
                              <label for="exampleInputPassword1">Password</label>
                              <input class="form-control" id="exampleInputPassword1" type="password" placeholder="Password" name="Password">
                          </div>
                          <div class="form-group">
                              <div class="form-check">
                                  <label class="form-check-label">
                                      <input class="form-check-input" type="checkbox"> Remember Password</label>
                                      </div>
                          </div>
                          <input class="btn btn-primary btn-block" type="submit" value="Login" />
                          </form>
              </div>
              </div>
     </div>

          {{ end }}
    </div>
    <!-- Contents -->
    <!-- Bootstrap core JavaScript-->
    <script src="/static/admin/vendor/jquery/jquery.min.js"></script>
    <script src="/static/admin/vendor/bootstrap/js/bootstrap.bundle.min.js"></script>
    <!-- Core plugin JavaScript-->
    <script src="/static/admin/vendor/jquery-easing/jquery.easing.min.js"></script>
    <!-- Page level plugin JavaScript-->
    <script src="/static/admin/vendor/datatables/jquery.dataTables.js"></script>
    <script src="/static/admin/vendor/datatables/dataTables.bootstrap4.js"></script>
    <!-- Custom scripts for all pages-->
    <script src="/static/admin/js/sb-admin.min.js"></script>
    <!-- Custom scripts for this page-->
    <script src="/static/admin/js/sb-admin-datatables.min.js"></script>
    </div>

    
    <script type="text/javascript" src="/static/ts/admin/dist/inline.bundle.js"></script>
    <script type="text/javascript" src="/static/ts/admin/dist/polyfills.bundle.js"></script>
    <script type="text/javascript" src="/static/ts/admin/dist/styles.bundle.js"></script>
    <script type="text/javascript" src="/static/ts/admin/dist/vendor.bundle.js"></script>
    <script type="text/javascript" src="/static/ts/admin/dist/main.bundle.js"></script>
</body>


</html>
