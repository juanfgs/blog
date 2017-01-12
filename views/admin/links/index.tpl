<div class="col-lg-12 ">
    <div class="row actions" role="">
    <a href="/admin/links/new" class="btn btn-sm btn-primary pull-right">New</a>
  </div>
  <table class="table table-striped">
    <thead>
      <tr>
	<th colspan="2">Name</th>
	<th>Actions</th>	
      </tr>
    </thead>
    <tbody>
      {{ range $key, $link := .Links }}
        
      <tr>
	<td colspan="2">
	<a href="/admin/links/edit/{{ $link.Id }}">
	    {{ $link.Title }}
	</a>
	</td>
	<td>
	  <a href="/admin/links/delete/{{ $link.Id }}" class="btn btn-sm btn-default">Delete</a>
	</td>
      </tr>
      {{ end }}
  </table>
</div>
