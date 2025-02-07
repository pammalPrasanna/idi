package lib

import "net/http"

type metricsResponseWriter struct {
	StatusCode    int
	BytesCount    int
	headerWritten bool
	wrapped       http.ResponseWriter
}

func newMetricsResponseWriter(w http.ResponseWriter) *metricsResponseWriter {
	return &metricsResponseWriter{
		StatusCode: http.StatusOK,
		wrapped:    w,
	}
}

func (mw *metricsResponseWriter) Header() http.Header {
	return mw.wrapped.Header()
}

func (mw *metricsResponseWriter) WriteHeader(statusCode int) {
	mw.wrapped.WriteHeader(statusCode)

	if !mw.headerWritten {
		mw.StatusCode = statusCode
		mw.headerWritten = true
	}
}

func (mw *metricsResponseWriter) Write(b []byte) (int, error) {
	mw.headerWritten = true

	n, err := mw.wrapped.Write(b)
	mw.BytesCount += n
	return n, err
}

func (mw *metricsResponseWriter) Unwrap() http.ResponseWriter {
	return mw.wrapped
}
