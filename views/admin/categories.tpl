<div class="col-lg-10 col-lg-offset-1"> 
     <ul class="list-group">
      {{ range $key, $category := .Categories }}

      	  <li class="list-group-item">
	<a href="/admin/categories/edit/{{ $category.Id }}">

	    {{ $category.Title }}


	</a>
			  <a href="/admin/category/delete/{{ $category.Id }}" class="btn btn-primary">Delete</a>
      </li> 
      {{ end }}
      </ul>
</div>
