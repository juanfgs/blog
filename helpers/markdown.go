package helpers
import(
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

func renderMarkDown(input string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(input))
	return string(unsafe)
}
func renderSafeMarkDown(input string) string {

	unsafe := blackfriday.MarkdownCommon([]byte(input))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(html)
}
func renderPost(content string, contentType string) string {
	if contentType == "markdown" {
		return renderMarkDown(content)
	} else {
		return content
	}
}
