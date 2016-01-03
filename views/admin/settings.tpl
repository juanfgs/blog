<div class="col-lg-12 ">
  <!-- Nav tabs -->
  <ul class="nav nav-pills" role="tablist">
    <li role="presentation" class="active"><a href="#home" aria-controls="home" role="tab" data-toggle="tab">Site Information</a></li>
    <li role="presentation"><a href="#url" aria-controls="profile" role="tab" data-toggle="tab">URL Settings</a></li>
    <li role="presentation"><a href="#messages" aria-controls="messages" role="tab" data-toggle="tab">Comments</a></li>
    <li role="presentation"><a href="#settings" aria-controls="settings" role="tab" data-toggle="tab">Settings</a></li>
  </ul>

  <!-- Tab panes -->
  <div class="tab-content">
    <div role="tabpanel" class="tab-pane active" id="home">
      <div class="col-md-4">
      <form method="POST">
	<h2>Website Information</h2>
	<div class="form-group">
	  <label for="Title">Website Title (html)</label>
	  <input type="text" class="form-control required" id="Title" name="Title" placeholder="Enter Title" />
	</div>
	<div class="form-group">
	  <label for="Title">Main Title</label>
	  <input type="text" class="form-control required" id="Title" name="MainTitle" placeholder="Enter Title" />
	</div>	
	
	<div class="form-group">
	  <label for="Title">Secondary Title</label>
	  <input type="text" class="form-control required" id="Title" name="SecondaryTitle" placeholder="Enter Secondary Title" />
	</div>
	
	  <button type="submit" class="btn btn-default">Submit</button>

      </form>
      </div>
    </div>
    
    <div role="tabpanel" class="tab-pane" id="url">
      <h2>URL</h2>
      <p>Regenerate slugs for all pages</p>
      <a href="/admin/posts/regenerate_slugs/" class="btn btn-sm btn-primary">Regenerate URLs</a>
    </div>
    <div role="tabpanel" class="tab-pane" id="messages">
      <h2>Comment settings</h2>
      <input class="pull-left" type="checkbox" name="enable-comments" value="1" /> Enable Comments
    </div>
    <div role="tabpanel" class="tab-pane" id="settings">...</div>
  </div>
</div>
