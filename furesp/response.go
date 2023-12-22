package furesp

import (
	"fmt"
	"github.com/ErtugrulAcar/futil/fuconf"
	"github.com/valyala/fasthttp"
	"time"
)

type success struct {
	Status  string      `json:"status"`
	Time    int64       `json:"time"`
	NextUrl string      `json:"next_url,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(data interface{}) *success {
	return &success{
		Status: "OK",
		Time:   time.Now().UTC().Unix(),
		Data:   data,
	}
}

func OKWithNextURL(data interface{}, URI *fasthttp.URI, totalCount int) *success {
	succ := success{
		Status: "OK",
		Time:   time.Now().UTC().Unix(),
		Data:   data,
	}

	page, err := URI.QueryArgs().GetUint(fuconf.Conf().GetPageQueryKey())
	if err != nil {
		return &succ
	}

	size, err := URI.QueryArgs().GetUint(fuconf.Conf().GetSizeQueryKey())
	if err != nil {
		return &succ
	}

	if (page+1)*size < totalCount {
		URI.QueryArgs().Set(fuconf.Conf().GetPageQueryKey(), fmt.Sprintf("%d", page+1))
		succ.NextUrl = URI.String()
	}
	return &succ
}

func Created(data interface{}) *success {
	return &success{
		Status: "CREATED",
		Time:   time.Now().UTC().Unix(),
		Data:   data,
	}
}
