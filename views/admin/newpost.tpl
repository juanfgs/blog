<div class="col-lg-10 col-lg-offset-1">
<form method="POST" >
  <div class="form-group">
    <label for="Title">Post Title</label>
    <input type="text" class="form-control required" id="Title" name="Title" placeholder="Enter Title">
  </div>
  <div class="form-group">
    <label for="Tagline">Tagline</label>
    <input type="text" class="form-control" name="Tagline" id="Tagline" placeholder="Enter Tagline">
  </div>
  <div class="form-group">
      <label for="Content">Post Content</label>
      <textarea class="form-control required tinymce" name="Content">	</textarea>
  </div>

  <div class="form-group">
    <label for="Keywords">SEO Keywords</label>
    <input type="text" class="form-control" name="Keywords" id="Tagline" placeholder="Enter ">
  </div>
  <div class="form-group">
      <label for="Description">SEO Description</label>
      <textarea class="form-control required" name="Description">	</textarea>
  </div>

  <div class="form-group">
      <label for="Category">Category</label>
      <select class="form-control" name="category_id">
      {{ range $key, $category := .Categories}}
      <option value="{{ $category.Id}}">{{ $category.Title}} </option>
      {{ end }}
      </select>
 </div>
  <div class="form-group">
      <label for="Published">Published</label>
      <input type="checkbox" value="1" name="Published"/>
</div>


  <button type="submit" class="btn btn-default">Submit</button>

</form>



</div>