package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler
}

//http.ResponseWriter implements the writing interface - Write([]byte) (int, error) between them
func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Checking if this is the last middleware on the chain
	//If yes, call the http.DefaultServerMux
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	encodings := r.Header.Get("Accept-Encoding")
	//if request does not support encoding, passing to the next handler
	if !strings.Contains(encodings, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}
	//Compressing the reponse
	w.Header().Add("Content-Encoding", "gzip")
	//Writes to the returned writer are compressed and written to w.
	gzipwriter := gzip.NewWriter(w)
	defer gzipwriter.Close()
	//implementing the http.ResponseWritter to pass our gzipwriter to the next handler
	grw := gzipResponseWriter{
		ResponseWriter: w,
		Writer:         gzipwriter,
	}
	gm.Next.ServeHTTP(grw, r)

}

//Implementing the http.ResponseWriter and writing interface
type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

//Overwriting the writing interface so it does not collide with gzipwriter Writer.
func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}
