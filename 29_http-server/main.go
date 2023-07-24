package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	// 익명 함수 활용
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// http.Handler 인터페이스 활용
	http.Handle("/test", new(testHandler))

	// Request URL Path에 표시된 정적 파일을 그대로 전달
	http.Handle("/", new(staticHandler))

	// 해당 경로의 모든 정적 리소스를 1:1로 매핑하여 그대로 전달
	http.Handle("/fileserver/", http.StripPrefix("/fileserver/", http.FileServer(http.Dir("fileserver"))))

	http.ListenAndServe(":5000", nil)

}

type testHandler struct {
	http.Handler
}

// http.Handler 인터페이스는 ServeHTTP 메서드를 가짐
// testHandler 구조체를 정의하고, testHandler의 메서드를 구현
func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your request path is " + req.URL.Path
	w.Write([]byte(str))
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// http://localhost:5000/index.html
	localPath := "wwwroot" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}
