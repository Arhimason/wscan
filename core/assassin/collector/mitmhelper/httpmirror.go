/**
* @Author: shaochuyu
* @Date: 5/7/2022 11:30
 */
package mitmhelper

import (
	"context"
	"github.com/panjf2000/ants/v2"
	"io"
	"sync"
	"wscan/core/assassin/http"
	"wscan/core/assassin/resource"
	"wscan/core/utils/checker"
)

type HTTPMirrorModifier struct {
	ctx          context.Context
	pool         *ants.Pool
	allowChecker *checker.RequestChecker
	httpOpts     *http.ClientOptions
	output       chan resource.Resource
}

func (*HTTPMirrorModifier) ModifyRequest(*http.Request) error {
	return nil
}

func (*HTTPMirrorModifier) ModifyResponse(*http.Response) error {
	return nil
}

func NewHTTPMirrorModifier() {

}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool
	once       sync.Once
	data       []uint8
	name       string
}
type multiReadCloser struct {
	io.Reader
	closer []io.Closer
}
