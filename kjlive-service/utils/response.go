package utils

import (
    "github.com/gin-gonic/gin"
    "bytes"
    "net"
    "bufio"
    "net/http"
)
const (
    noWritten	= -1
    defaultStatus	= 200
)

// 重写Response，为了取得Response的data
type ResponseBuffer struct {
    Response gin.ResponseWriter // the actual ResponseWriter to flush to
    status   int                 // the HTTP response code from WriteHeader
    Body     *bytes.Buffer       // the response content body
    Flushed  bool
}

func NewResponseBuffer(w gin.ResponseWriter) *ResponseBuffer {
    return &ResponseBuffer{
        Response: w, status: defaultStatus, Body: &bytes.Buffer{},
    }
}

func (w *ResponseBuffer) Header() http.Header {
    return w.Response.Header() // use the actual response header
}

func (w *ResponseBuffer) Write(buf []byte) (int, error) {
    w.Body.Write(buf)
    return len(buf), nil
}

func (w *ResponseBuffer) WriteString(s string) (n int, err error) {
    //w.WriteHeaderNow()
    //n, err = io.WriteString(w.ResponseWriter, s)
    //w.size += n
    n, err = w.Write([]byte(s))
    return
}

func (w *ResponseBuffer) Written() bool {
    return w.Body.Len() != noWritten
}

func (w *ResponseBuffer) WriteHeader(status int) {
    w.status = status
}

func (w *ResponseBuffer) WriteHeaderNow() {
    //if !w.Written() {
    //	w.size = 0
    //	w.ResponseWriter.WriteHeader(w.status)
    //}
}

func (w *ResponseBuffer) Status() int {
    return w.status
}

func (w *ResponseBuffer) Size() int {
    return w.Body.Len()
}

func (w *ResponseBuffer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
    //if w.size < 0 {
    //	w.size = 0
    //}
    return w.Response.(http.Hijacker).Hijack()
}

func (w *ResponseBuffer) CloseNotify() <-chan bool {
    return w.Response.(http.CloseNotifier).CloseNotify()
}

// Fake Flush
// TBD
func (w *ResponseBuffer) Flush() {
    w.realFlush()
}

func (w *ResponseBuffer) realFlush() {
    if w.Flushed {
        return
    }
    w.Response.WriteHeader(w.status)
    if w.Body.Len() > 0 {
        _, err := w.Response.Write(w.Body.Bytes())
        if err != nil {
            panic(err)
        }
        w.Body.Reset()
    }
    w.Flushed = true
}