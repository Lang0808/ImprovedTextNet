package handler

import (
	utils "WebServer/utils"
	"fmt"
	"net/http"
	"time"

	LibUtils "github.com/Lang0808/libutils"
)

type RequestHandlerInterface interface {
	GetSrcId(r *http.Request) (int32, error)
	GetCommandId() int32
	Handle(srcId int32, w http.ResponseWriter, r *http.Request, extra *map[string]string)
	PostProcess(w *http.ResponseWriter, r *http.Request)
}

type RequestHandler struct {
	inner RequestHandlerInterface
}

func (h *RequestHandler) process(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now().UnixMilli()

	logEntry := LibUtils.LogEntry{}
	srcId, err := h.inner.GetSrcId(r)
	if err != nil {
		fmt.Fprintf(w, GetApiMessage(UNAUTHORIZE_MSG))
		return
	}
	logEntry.SrcId = srcId
	logEntry.CommandId = h.inner.GetCommandId()
	logEntry.Extra = make(map[string]string)

	h.inner.PostProcess(&w, r)

	h.inner.Handle(srcId, w, r, &logEntry.Extra)

	endTime := time.Now().UnixMilli()
	fmt.Printf("endTime = %v, startTime = %v\n", endTime, startTime)
	logEntry.ExecTime = int32(endTime - startTime)
	LibUtils.Info_entry("handler", "process", logEntry)
}

func DefaultGetCommandId() int32 {
	return -1
}

func DefaultHandle(w http.ResponseWriter, r *http.Request,
	Extra map[string]string) {

}

func DefaultPostProcess(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Origin", LibUtils.Get("client.domain"))
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "content-type")
}

func DefaultGetSrcId(r *http.Request) (int32, error) {
	AuthCookie, err := r.Cookie("jwt-token")
	if err != nil {
		return -1, err
	}
	srcId, err := utils.GetUserIdInJWTToken(AuthCookie.Value)
	if err != nil {
		return -1, err
	}
	return srcId, nil
}
