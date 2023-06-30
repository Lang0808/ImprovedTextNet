package GrpcBlogService

import (
	"context"

	LibUtils "github.com/Lang0808/libutils"
)

type RequestHandler struct {
	inner RequestHandlerInterface
}

type RequestHandlerInterface interface {
	GetCommandId() int32
}

func (r RequestHandler) process(ctx context.Context, req *interface{}) (*interface{}, error) {
	logEntry := LibUtils.LogEntry{}
	logEntry.CommandId = r.inner.GetCommandId()
	logEntry.SrcId = -1
	logEntry.Extra = make(map[string]string)
}
