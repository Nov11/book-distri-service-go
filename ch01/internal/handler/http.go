package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Nov11/proglog/ch01/internal/log"
	"html"
	"net/http"
)

type httpServer struct {
	Log *log.Log
}

func NewHTTPServer() *httpServer {
	return &httpServer{
		Log: log.NewLog(),
	}
}

type ProduceRequest struct {
	Record log.Record `json:"record"`
}
type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}
type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}
type ConsumeResponse struct {
	Record log.Record `json:"record"`
}

func errorAndReturn(w http.ResponseWriter, err error, status int) bool {
	if err != nil {
		http.Error(w, err.Error(), status)
		return true
	}
	return false
}

func (h *httpServer) handGET(w http.ResponseWriter, r *http.Request) {
	req := &ConsumeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if errorAndReturn(w, err, http.StatusBadRequest) {
		return
	}

	read, err := h.Log.Read(req.Offset)
	if errorAndReturn(w, err, http.StatusInternalServerError) {
		return
	}

	resp := &ConsumeResponse{Record: read}
	err = json.NewEncoder(w).Encode(resp)
	if errorAndReturn(w, err, http.StatusInternalServerError) {
		return
	}
}

func (h *httpServer) handlePOST(w http.ResponseWriter, r *http.Request) {
	req := &ProduceRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if errorAndReturn(w, err, http.StatusBadRequest) {
		return
	}
	u, err := h.Log.Append(req.Record)
	if errorAndReturn(w, err, http.StatusInternalServerError) {
		return
	}

	resp := &ProduceResponse{}
	resp.Offset = u
	err = json.NewEncoder(w).Encode(resp)
	if errorAndReturn(w, err, http.StatusInternalServerError) {
		return
	}
}

func (h *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hello path:%s", html.EscapeString(r.URL.Path))

	switch r.Method {
	case "GET":
		h.handGET(w, r)
	case "POST":
		h.handlePOST(w, r)
	default:
		_, _ = fmt.Fprintf(w, "only GET / POST methods are supported.")
	}
}
