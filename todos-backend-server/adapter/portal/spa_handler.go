package portal

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

type Handler struct {
	StaticPath string
	IndexPath  string
}

func NewSpaHandler() *Handler {
	return &Handler{StaticPath: "www", IndexPath: "index.html"}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := absPath(r.URL)
	if err != nil {
		log.Println("Invalid URL:", r.URL, "- Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.StaticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
	} else if err != nil {
		log.Println("File not readable:", path, "- Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
	}
}

func absPath(url *url.URL) (path string, err error) {
	if path, err = filepath.Abs(url.Path); err != nil {
		return
	}

	// Remove drive like `C:` at start of absolute path on Windows
	re := regexp.MustCompile(`^[a-zA-Z]:(.*)`)
	path = re.ReplaceAllString(path, "${1}")

	return
}
