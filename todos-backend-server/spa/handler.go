// Package spa defines a handler to serve a single page application.
package spa

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

// A Handler serves a single page application from static files.
type Handler struct {
	// Root path of static content, using default "www" if not set.
	StaticPath string
	// Path to index of SPA, using default "index.html" if not set.
	IndexPath string
}

// ServeHTTP delivers existing files or redirect to index.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := absPath(r.URL)
	if err != nil {
		log.Println("Invalid URL:", r.URL, "- Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	staticPath := "www"
	if h.StaticPath != "" {
		staticPath = h.StaticPath
	}
	indexPath := "index.html"
	if h.IndexPath != "" {
		indexPath = h.IndexPath
	}

	path = filepath.Join(staticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
	} else if err != nil {
		log.Println("File not readable:", path, "- Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.FileServer(http.Dir(staticPath)).ServeHTTP(w, r)
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
