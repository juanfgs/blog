<div class="col-lg-10 col-lg-offset-1">
<form method="POST" >
  <div class="form-group">
    <label for="Title">Post Title</label>
    <input type="text" class="form-control required" value="{{ .Post.Title }}" id="Title" name="Title" placeholder="Enter Title">
  </div>
  <div class="form-group">
    <label for="Tagline">Tagline</label>
    <input type="text" class="form-control" value="{{ .Post.Tagline}}" name="Tagline" id="Tagline" placeholder="Enter Tagline">
  </div>
  <div class="form-group">
      <label for="Content">Post Content</label>
      <textarea class="form-control required tinymce" name="Content">{{ .Post.Content }}	</textarea>
  </div>
  <div class="form-group">
    <label for="Keywords">SEO Keywords</label>
    <input type="text" class="form-control" name="Keywords" id="Tagline" placeholder="Enter SEO Keywords" value="{{ .Post.Keywords}}">
  </div>
  <div class="form-group">
      <label for="Description">SEO Description</label>
      <textarea class="form-control required" name="Description">{{.Post.Description}}	</textarea>
  </div>
  <div class="form-group">
      <label for="Category">Category</label>
      <select class="form-control" name="CategoryId">

      {{ range $key, $category := .Categories}}

	       	 <option value="{{ $category.Id}}">{{ $category.Title}} </option>

      {{ end }}

      </select>

 </div>
  <div class="form-group">
      <label for="Published">Published</label>
      <input type="checkbox" value="1" name="Published" {{ if .Post.Published }} checked="checked" {{ end }}/>
</div>

  <button type="submit" class="btn btn-default">Submit</button>

</form>



</div>
