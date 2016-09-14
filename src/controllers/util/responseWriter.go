package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type CloseableResponseWriter interface {
	http.ResponseWriter
	Close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (this gzipResponseWriter) Close() {
	this.Writer.Close()
}

func (this gzipResponseWriter) Write(data []byte) (int, error) {
	return this.Writer.Write(data)
}

func (this gzipResponseWriter) Header() http.Header {
	return this.ResponseWriter.Header()
}

type closeableResponseWriter struct {
	http.ResponseWriter
}

func (this closeableResponseWriter) Close() {}

func GetReponseWriter(w http.ResponseWriter, req *http.Request) CloseableResponseWriter {
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gRw := gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzip.NewWriter(w),
		}
		return gRw
	} else {
		return closeableResponseWriter{ResponseWriter: w}
	}
}
