// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: server/rank/rank.proto

package rank

import (
	_ "e.coding.net/mmstudio/blade/server/proto/global"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RankService service

func NewRankServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "RankService.QueryRankByObjId",
			Path:    []string{"/query"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for RankService service

type RankService interface {
	QueryRankByObjId(ctx context.Context, in *QueryRankByObjIdRq, opts ...client.CallOption) (*QueryRankByObjIdRs, error)
	QueryRankByRange(ctx context.Context, in *QueryRankByRangeRq, opts ...client.CallOption) (*QueryRankByRangeRs, error)
	SetRankScore(ctx context.Context, in *SetRankScoreRq, opts ...client.CallOption) (*SetRankScoreRs, error)
	KickRankData(ctx context.Context, in *KickRankDataRq, opts ...client.CallOption) (*KickRankDataRs, error)
}

type rankService struct {
	c    client.Client
	name string
}

func NewRankService(name string, c client.Client) RankService {
	return &rankService{
		c:    c,
		name: name,
	}
}

func (c *rankService) QueryRankByObjId(ctx context.Context, in *QueryRankByObjIdRq, opts ...client.CallOption) (*QueryRankByObjIdRs, error) {
	req := c.c.NewRequest(c.name, "RankService.QueryRankByObjId", in)
	out := new(QueryRankByObjIdRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankService) QueryRankByRange(ctx context.Context, in *QueryRankByRangeRq, opts ...client.CallOption) (*QueryRankByRangeRs, error) {
	req := c.c.NewRequest(c.name, "RankService.QueryRankByRange", in)
	out := new(QueryRankByRangeRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankService) SetRankScore(ctx context.Context, in *SetRankScoreRq, opts ...client.CallOption) (*SetRankScoreRs, error) {
	req := c.c.NewRequest(c.name, "RankService.SetRankScore", in)
	out := new(SetRankScoreRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankService) KickRankData(ctx context.Context, in *KickRankDataRq, opts ...client.CallOption) (*KickRankDataRs, error) {
	req := c.c.NewRequest(c.name, "RankService.KickRankData", in)
	out := new(KickRankDataRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RankService service

type RankServiceHandler interface {
	QueryRankByObjId(context.Context, *QueryRankByObjIdRq, *QueryRankByObjIdRs) error
	QueryRankByRange(context.Context, *QueryRankByRangeRq, *QueryRankByRangeRs) error
	SetRankScore(context.Context, *SetRankScoreRq, *SetRankScoreRs) error
	KickRankData(context.Context, *KickRankDataRq, *KickRankDataRs) error
}

func RegisterRankServiceHandler(s server.Server, hdlr RankServiceHandler, opts ...server.HandlerOption) error {
	type rankService interface {
		QueryRankByObjId(ctx context.Context, in *QueryRankByObjIdRq, out *QueryRankByObjIdRs) error
		QueryRankByRange(ctx context.Context, in *QueryRankByRangeRq, out *QueryRankByRangeRs) error
		SetRankScore(ctx context.Context, in *SetRankScoreRq, out *SetRankScoreRs) error
		KickRankData(ctx context.Context, in *KickRankDataRq, out *KickRankDataRs) error
	}
	type RankService struct {
		rankService
	}
	h := &rankServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RankService.QueryRankByObjId",
		Path:    []string{"/query"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&RankService{h}, opts...))
}

type rankServiceHandler struct {
	RankServiceHandler
}

func (h *rankServiceHandler) QueryRankByObjId(ctx context.Context, in *QueryRankByObjIdRq, out *QueryRankByObjIdRs) error {
	return h.RankServiceHandler.QueryRankByObjId(ctx, in, out)
}

func (h *rankServiceHandler) QueryRankByRange(ctx context.Context, in *QueryRankByRangeRq, out *QueryRankByRangeRs) error {
	return h.RankServiceHandler.QueryRankByRange(ctx, in, out)
}

func (h *rankServiceHandler) SetRankScore(ctx context.Context, in *SetRankScoreRq, out *SetRankScoreRs) error {
	return h.RankServiceHandler.SetRankScore(ctx, in, out)
}

func (h *rankServiceHandler) KickRankData(ctx context.Context, in *KickRankDataRq, out *KickRankDataRs) error {
	return h.RankServiceHandler.KickRankData(ctx, in, out)
}
