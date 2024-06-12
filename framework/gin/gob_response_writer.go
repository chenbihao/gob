package gin

import "net/http"

// NewResponseWriter ...
func NewResponseWriter(writer http.ResponseWriter) *responseWriter {
	var w responseWriter
	w.ResponseWriter = writer
	w.size = noWritten
	w.status = defaultStatus
	return &w
}
