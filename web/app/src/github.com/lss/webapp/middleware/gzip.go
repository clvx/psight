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
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()
	var rw http.ResponseWriter
	//verifying the response implements the Pusher interface
	if pusher, ok := w.(http.Pusher); ok {
		//overwriting gzip ResponseWriter with the http.Pusher
		rw = gzipPusherResponseWriter{
			gzipResponseWriter: gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gzipWriter,
			},
			Pusher: pusher,
		}
	} else {
		//implementing the http.ResponseWritter to pass our gzipWriter to the next handler
		rw = gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzipWriter,
		}
	}
	gm.Next.ServeHTTP(rw, r)
}

//Implementing the http.ResponseWriter and writing interface for gzip
type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

//Implementing the gzipResponseWriter with http.Pusher
type gzipPusherResponseWriter struct {
	gzipResponseWriter
	http.Pusher
}

//Overwriting the writing interface so it does not collide with gzipWriter Writer.
func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}
