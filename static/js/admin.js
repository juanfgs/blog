document.addEventListener("DOMContentLoaded",function(evt){
    var formElement = document.getElementById("js-upload-form");
    var uploadedFiles = document.getElementById("js-upload-files");    

    formElement.addEventListener("submit", function(e){
	if (e.preventDefault) e.preventDefault();
	var fd = new FormData();

	for (var i = 0; i < uploadedFiles.files.length; i++){
	    fd.append("media",uploadedFiles.files[i]);
	    var request = new XMLHttpRequest();
	    request.open("POST", "/admin/media/create");
	    request.send(fd);

	    request.onreadystatechange = function(){
		if(request.readyState == 4 && request.status == 200){
		    var file = JSON.parse(request.responseText);
		    var element = document.createElement('div');
		    var insert = document.createElement('a');
		    insert.appendChild(document.createTextNode("Insert"));
		    var del = document.createElement('a');
		    del.appendChild(document.createTextNode("Delete"));
		    insert.addEventListener('click', function(){
			switch(file.Type){
			case 'image/png':
			case 'image/jpeg':
			case 'image/gif':
			    insertImage(file);
			    break;
			case 'application/x-download':
			case 'application/x-rar-compressed':
			case 'application/x-zip-compressed':
			    insertDownloadLink(file);
			    break;
			case 'video/webm':
			    insertVideo(file);
			    break;			    
			}
		    });
		    del.addEventListener('click',function(){
			deleteFile(file);
		    });
		    
		    element.setAttribute('class', 'list-group-item list-group-item-success');
		    element.setAttribute('id', 'file-'+file.Id);		    
		    insert.setAttribute('class', 'pull-right btn btn-sm btn-primary');
		    del.setAttribute('class', 'pull-right btn btn-sm btn-default');		    		    
		    element.appendChild(document.createTextNode(file.Filename));
		    element.appendChild(insert);
		    element.appendChild(del);		    

		    var uploadedFilesList = document.getElementById('uploadedFiles');
		    uploadedFilesList.insertBefore(element,uploadedFilesList.firstChild);
		} 

		
	    }
	}

	return false;
    });


});

function deleteFile(file){
    var request = new XMLHttpRequest();
    request.open("GET", "/admin/media/delete/" + file.Id);
    request.send();
    request.onreadystatechange = function(){
	if(request.readyState == 4 && request.status == 200){
	    var element = document.getElementById("file-"+file.Id);
	    element.parentNode.removeChild(element);
	    
	}
    }
}


function insertImage(file){
    var contentArea = document.getElementById("ContentArea");
    var text = '![fit-width](/static/uploads/'+file.Filename+' )';
    insertAtCursor(contentArea,text);
}

function insertDownloadLink(file){
    var contentArea = document.getElementById("ContentArea");
    var text = '[ '+file.Filename+'](/static/uploads/'+file.Filename+' )';
    insertAtCursor(contentArea,text);
}
function insertVideo(file){
    var contentArea = document.getElementById("ContentArea");
    var text = '\<video src="/static/uploads/'+file.Filename+'" controls>Your browser does not support the <code>video</code> element.</video>';
    insertAtCursor(contentArea,text);    
}




function insertAtCursor(myField, myValue) {
    //IE support
    if (document.selection) {
        myField.focus();
        sel = document.selection.createRange();
        sel.text = myValue;
	
    }
    //MOZILLA and others
    else if (myField.selectionStart || myField.selectionStart == '0') {
        var startPos = myField.selectionStart;
        var endPos = myField.selectionEnd;
        myField.value = myField.value.substring(0, startPos)
            + myValue
            + myField.value.substring(endPos, myField.value.length);
	myField.selectionStart = startPos + myValue.length;
	myField.selectionEnd = startPos + myValue.length;
    } else {
        myField.value += myValue;
    }
}




