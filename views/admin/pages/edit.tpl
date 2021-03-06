<div class="col-lg-12 ">
<form method="POST" >
  <div class="form-group">
    <label for="Title">Page Title</label>
    <input type="text" class="form-control required" value="{{ .Page.Title }}" id="Title" name="Title" placeholder="Enter Title">
  </div>
  <div class="form-group">
    <label for="Tagline">Tagline</label>
    <input type="text" class="form-control" value="{{ .Page.Tagline}}" name="Tagline" id="Tagline" placeholder="Enter Tagline">
  </div>
  <div class="form-group">
      <label for="Content">Page Content</label>
      <textarea class="form-control required tinymce" id="ContentArea" name="Content" rows="15">{{ .Page.Content }}	</textarea>
  </div>
  <div class="form-group">
    <button type="button" class="btn btn-primary btn-sm" data-toggle="modal" data-target="#mediaModal">
      Add Media
    </button>
    
  <a class="btn btn-primary btn-sm" data-toggle="collapse" href="#collapseExample" aria-expanded="false" aria-controls="collapseExample">
    SEO Info
  </a>

    <select form-control name="ContentType">
      <option value="markdown" selected="selected">Markdown</option>
      <option value="HTML">HTML</option>
      <option value="text">Plain Text </option>            
    </select>
  </div>


  <div class="well collapse" id="collapseExample">
  <div class="form-group">
    <label for="Keywords">SEO Keywords</label>
    <input type="text" class="form-control" name="Keywords" id="Tagline" placeholder="Enter SEO Keywords" value="{{ .Page.Keywords}}">
  </div>
  <div class="form-group">
      <label for="Description">SEO Description</label>
      <textarea class="form-control required" name="Description">{{.Page.Description}}	</textarea>
  </div>
</div>

  <div class="form-group">
      <label for="Published">Published</label>
      <input type="checkbox" value="1" name="Published" {{ if .Page.Published }} checked="checked" {{ end }}/>
</div>

  <button type="submit" class="btn btn-default">Submit</button>

</form>



</div>

<!-- Media Form -->
<div class="modal fade" id="mediaModal" tabindex="-1" role="dialog" aria-labelledby="mediaModalLabel" aria-hidden="true">
  <div class="modal-dialog modal-lg">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title">Upload Files</h4>
      </div>
      <div class="modal-body">
<div class="container-fluid">


          <!-- Standar Form -->
          <h4>Select files from your computer</h4>
          <form action="" method="post" enctype="multipart/form-data" id="js-upload-form">
            <div class="form-inline">
              <div class="form-group">
                <input type="file" name="files[]" id="js-upload-files" multiple>
              </div>
              <button type="submit" class="btn btn-sm btn-primary" id="js-upload-submit">Upload files</button>
            </div>
          </form>


          <!-- Upload Finished -->
          <div class="js-upload-finished" >
            <h3>Available media</h3>
            <div class="list-group" id="uploadedFiles">
         
	      
      </div>
    </div> <!-- /container -->
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
</div>




<script src="/static/js/admin.js"></script>    

