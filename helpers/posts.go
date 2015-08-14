package helpers

func RenderPost(content string, contentType string) string {
	if contentType == "markdown" {
		return RenderMarkDown(content)
	} else {
		return content
	}
}
