package helpers
import(
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)


func RenderMarkDown(input string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(input))
	return string(unsafe)
}
func RenderSafeMarkDown(input string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(input))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(html)
}

