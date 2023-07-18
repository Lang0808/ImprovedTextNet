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
	

}
