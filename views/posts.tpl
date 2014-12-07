
<header class="intro-header" style="background-image: url('/static/img/home-bg.jpg')">
  <div class="container">
    <div class="row">
      <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
        <div class="site-heading">
          <h1>{{ .HeroTitle }}</h1>
          <hr class="small">
          <span class="subheading">{{.HeroTagline }}</span>
        </div>
      </div>
    </div>
  </div>
</header>

<!-- Main Content -->
<div class="container">
  <div class="row">
    <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
      {{ range $key, $post := .posts }}
      <div class="post-preview">
	<a href="/post/{{ $post.Id }}">
	  <h2 class="post-title">
	    {{ $post.Title }}
	  </h2>
	  <h3 class="post-subtitle">
	    {{ str2html $post.Content }}
	  </h3>
	</a>
	<p class="post-meta">Posted by <a href="#">Juan</a> on {{ .CreatedAt}}</p>
      </div>
      <hr>
      {{ end }}
      <!-- Pager -->
      <ul class="pager">
	<li class="next">
	  <a href="#">Older Posts &rarr;</a>
	</li>
      </ul>	     
    </div>
  </div>
</div>

<hr>




