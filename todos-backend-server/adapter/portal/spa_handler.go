package portal

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Handler struct {
	RootPath   string
	IndexPath  string
	StaticPath string
}

func NewSpaHandler() *Handler {
	return &Handler{
		RootPath:   "www",
		IndexPath:  "index.html",
		StaticPath: "/static",
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedPath, err := absPath(r.URL)
	if err != nil {
		log.Println("Invalid URL:", r.URL, "- Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	localPath := filepath.Join(h.RootPath, requestedPath)
	_, err = os.Stat(localPath)
	if os.IsNotExist(err) {
		// Never cache index file.
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, filepath.Join(h.RootPath, h.IndexPath))
	} else if err != nil {
		log.Println("File not readable:", localPath, "- Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if h.StaticPath != "" {
			if strings.HasPrefix(requestedPath, h.StaticPath) {
				// Static files can be cached for 1 year.
				w.Header().Set("Cache-Control", "max-age=31536000")
			} else {
				// Other files must revalidate on server.
				w.Header().Set("Cache-Control", "no-cache")
			}
		}
		http.FileServer(http.Dir(h.RootPath)).ServeHTTP(w, r)
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
