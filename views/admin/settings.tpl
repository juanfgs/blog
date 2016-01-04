<div class="col-lg-12 ">
  <!-- Nav tabs -->
  <ul class="nav nav-pills" role="tablist">
    <li role="presentation" class="active"><a href="#home" aria-controls="home" role="tab" data-toggle="tab">Site Information</a></li>
    <li role="presentation"><a href="#url" aria-controls="profile" role="tab" data-toggle="tab">URL Settings</a></li>
    <li role="presentation"><a href="#messages" aria-controls="messages" role="tab" data-toggle="tab">Comments</a></li>
    <li role="presentation"><a href="#settings" aria-controls="settings" role="tab" data-toggle="tab">Writing</a></li>
  </ul>

  <!-- Tab panes -->
  <div class="tab-content">
    <div role="tabpanel" class="tab-pane active" id="home">
      <div class="col-md-4">
      <form method="POST" action="/admin/settings/save">
	<h2>Website Information</h2>
	<div class="form-group">
	  <label for="Title">Website Title (html)</label>
	  <input type="text" class="form-control required" id="Title" name="Title" placeholder="Enter Title" value="{{ index .Settings "Title" }}" />
	</div>
	<div class="form-group">
	  <label for="Title">Main Title</label>
	  <input type="text" class="form-control required" id="Title" name="MainTitle" placeholder="Enter Title" value="{{ index .Settings "MainTitle" }}" />
	</div>	
	
	<div class="form-group">
	  <label for="Title">Secondary Title</label>
	  <input type="text" class="form-control required" id="Title" name="SecondaryTitle" placeholder="Enter Secondary Title" value="{{ index .Settings "SecondaryTitle" }}" />
	</div>
	
	  <button type="submit" class="btn btn-default">Submit</button>

      </form>
      </div>
    </div>
    
    <div role="tabpanel" class="tab-pane" id="url">
      <form method="POST" action="/admin/settings/save">
		<div class="form-group">		
      <h2>URL</h2>

	  <input class="pull-left" type="checkbox" name="EnablePrettyUrls" value="1"
		 {{ if index .Settings "EnableComments" }}
		 checked="checked"
		 {{ end }}
/> Enable Pretty URLs
</div>
	<button type="submit" class="btn btn-default">Submit</button>
      </form>
    </div>
    <div role="tabpanel" class="tab-pane" id="messages">
      <form method="POST" action="/admin/settings/save">

	<h2>Comment settings</h2>
	<div class="form-group">		
	  <input class="pull-left" type="checkbox" name="EnableComments" value="1"
		 {{ if index .Settings "EnableComments" }}
		 checked="checked"
		 {{ end }}
/> Enable Comments
	</div>
	<button type="submit" class="btn btn-default">Submit</button>
      </form>
    </div>
    <div role="tabpanel" class="tab-pane" id="settings">
      <form method="POST" action="/admin/settings/save">

	<h2>Writing Settings</h2>
	<div class="form-group">
	  <label for="Title">Default Format</label>
	  <p> Default format for posts, Current format: {{ index .Settings "PostsDefaultFormat" }} </p>
	  <select name="PostsDefaultFormat" class="form-control" >
	    <option value="html">Html</option>
	    <option value="markdown">Markdown</option>
	    <option value="plain text">Plain text</option>
	  </select>
	  
	</div>
	<div class="form-group">
	  <label for="Title">WYSIWYG </label>
	  <p> Enable WYSIWYG for HTML posts </p>
	  <input class="pull-left" type="checkbox" name="EnableWYSIWYG" value="1"
		 {{ if index .Settings "EnableWYSIWYG" }}
		 checked="checked"
		 {{ end }}
/> Enable WYSIWYG
	  
	</div>
	<button type="submit" class="btn btn-default">Submit</button>
      </form>
    </div>
  </div>
</div>
