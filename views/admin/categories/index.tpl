<div class="col-lg-12 ">
    <div class="row actions" role="">
    <a href="/admin/categories/new" class="btn btn-sm btn-primary pull-right">New</a>
  </div>
  <table class="table table-striped">
    <thead>
      <tr>
	<th colspan="2">Name</th>
	<th>Actions</th>	
      </tr>
    </thead>
    <tbody>
      {{ range $key, $category := .Categories }}
      <tr>
	<td colspan="2">
	<a href="/admin/categories/edit/{{ $category.Id }}">
	    {{ $category.Title }}
	</a>
	</td>
	<td>
	  <a href="/admin/category/delete/{{ $category.Id }}" class="btn btn-sm btn-default">Delete</a>
	</td>
      </tr>
      {{ end }}
  </table>
</div>
