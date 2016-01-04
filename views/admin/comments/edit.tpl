<div class="col-lg-12">
  <h2>Editing Comment</h2>
<form method="POST">
  <div class="form-group">
    <label for="Commenter">Commenter</label>
    <input type="text" class="form-control required" id="Commenter" name="Commenter" placeholder="" value="{{ .Commenter}}">
  </div>
    <div class="form-group">
      <label for="Comment">Comment </label>
      <textarea class="form-control required tinymce" id="ContentArea" name="Comment" rows="15">{{ .Comment }}	</textarea>
    </div>

  <button type="submit" class="btn btn-default">Submit</button>

</form>


</div>


<script src="/static/js/admin.js"></script>    
