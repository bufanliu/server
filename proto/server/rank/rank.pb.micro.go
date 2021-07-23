// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: server/rank/rank.proto

package rank

import (
	_ "e.coding.net/mmstudio/blade/server/proto/global"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	return []*api.Endpoint{}
}

// Client API for RankService service

type RankService interface {
	//   rpc CreateMail(CreateMailRq) returns (CreateMailRs) {}
	//   rpc QueryPlayerMails(QueryPlayerMailsRq) returns (QueryPlayerMailsRs) {}
	//   rpc ReadMail(ReadMailRq) returns (ReadMailRs) {}
	//   rpc GainAttachments(GainAttachmentsRq) returns (GainAttachmentsRs) {}
	//   rpc DelMail(DelMailRq) returns (DelMailRs) {}
	QueryRankByKey(ctx context.Context, in *QueryRankByKeyRq, opts ...client.CallOption) (*QueryRankByKeyRs, error)
	QueryRankByIndex(ctx context.Context, in *QueryRankByIndexRq, opts ...client.CallOption) (*QueryRankByIndexRs, error)
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

func (c *rankService) QueryRankByKey(ctx context.Context, in *QueryRankByKeyRq, opts ...client.CallOption) (*QueryRankByKeyRs, error) {
	req := c.c.NewRequest(c.name, "RankService.QueryRankByKey", in)
	out := new(QueryRankByKeyRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankService) QueryRankByIndex(ctx context.Context, in *QueryRankByIndexRq, opts ...client.CallOption) (*QueryRankByIndexRs, error) {
	req := c.c.NewRequest(c.name, "RankService.QueryRankByIndex", in)
	out := new(QueryRankByIndexRs)
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
	//   rpc CreateMail(CreateMailRq) returns (CreateMailRs) {}
	//   rpc QueryPlayerMails(QueryPlayerMailsRq) returns (QueryPlayerMailsRs) {}
	//   rpc ReadMail(ReadMailRq) returns (ReadMailRs) {}
	//   rpc GainAttachments(GainAttachmentsRq) returns (GainAttachmentsRs) {}
	//   rpc DelMail(DelMailRq) returns (DelMailRs) {}
	QueryRankByKey(context.Context, *QueryRankByKeyRq, *QueryRankByKeyRs) error
	QueryRankByIndex(context.Context, *QueryRankByIndexRq, *QueryRankByIndexRs) error
	SetRankScore(context.Context, *SetRankScoreRq, *SetRankScoreRs) error
	KickRankData(context.Context, *KickRankDataRq, *KickRankDataRs) error
}

func RegisterRankServiceHandler(s server.Server, hdlr RankServiceHandler, opts ...server.HandlerOption) error {
	type rankService interface {
		QueryRankByKey(ctx context.Context, in *QueryRankByKeyRq, out *QueryRankByKeyRs) error
		QueryRankByIndex(ctx context.Context, in *QueryRankByIndexRq, out *QueryRankByIndexRs) error
		SetRankScore(ctx context.Context, in *SetRankScoreRq, out *SetRankScoreRs) error
		KickRankData(ctx context.Context, in *KickRankDataRq, out *KickRankDataRs) error
	}
	type RankService struct {
		rankService
	}
	h := &rankServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RankService{h}, opts...))
}

type rankServiceHandler struct {
	RankServiceHandler
}

func (h *rankServiceHandler) QueryRankByKey(ctx context.Context, in *QueryRankByKeyRq, out *QueryRankByKeyRs) error {
	return h.RankServiceHandler.QueryRankByKey(ctx, in, out)
}

func (h *rankServiceHandler) QueryRankByIndex(ctx context.Context, in *QueryRankByIndexRq, out *QueryRankByIndexRs) error {
	return h.RankServiceHandler.QueryRankByIndex(ctx, in, out)
}

func (h *rankServiceHandler) SetRankScore(ctx context.Context, in *SetRankScoreRq, out *SetRankScoreRs) error {
	return h.RankServiceHandler.SetRankScore(ctx, in, out)
}

func (h *rankServiceHandler) KickRankData(ctx context.Context, in *KickRankDataRq, out *KickRankDataRs) error {
	return h.RankServiceHandler.KickRankData(ctx, in, out)
}
