package helpers


func RenderPost(content string, contentType string, excerpt bool) string {
	if excerpt {
		content = Excerpt(content, contentType)
	}
	
	if contentType == "markdown" {
		return RenderMarkDown(content)
	} else {
		return content
	}
}

func Excerpt(content string, contentType string) string {
	if contentType == "markdown" {
		return content[0:300] + "..."
	} else {
		return content
	}
	
}
