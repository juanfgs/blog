<div class="col-lg-12">
<form method="POST" >
  <div class="form-group">
    <label for="Title">Category Title</label>
    <input type="text" class="form-control required" id="Title" name="Title" placeholder="Enter Title" value="{{ .Category.Title }}">
  </div>
  <div class="form-group">
      <label for="Description">Description</label>
      <textarea class="form-control required tinymce" name="Description">{{.Category.Description}}	</textarea>
  </div>

  <button type="submit" class="btn btn-default">Submit</button>

</form>



</div>
