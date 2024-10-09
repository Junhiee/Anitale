// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: anime.proto

package animeservice

import (
	"context"

	"Anitale/apps/anime/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Anime           = pb.Anime
	AnimeAddReq     = pb.AnimeAddReq
	AnimeAddResp    = pb.AnimeAddResp
	AnimeDeleteReq  = pb.AnimeDeleteReq
	AnimeDeleteResp = pb.AnimeDeleteResp
	AnimeListReq    = pb.AnimeListReq
	AnimeListResp   = pb.AnimeListResp
	AnimeUpdateReq  = pb.AnimeUpdateReq
	AnimeUpdateResp = pb.AnimeUpdateResp

	AnimeService interface {
		AnimeAdd(ctx context.Context, in *AnimeAddReq, opts ...grpc.CallOption) (*AnimeAddResp, error)
		AnimeDelete(ctx context.Context, in *AnimeDeleteReq, opts ...grpc.CallOption) (*AnimeDeleteResp, error)
		AnimeUpdate(ctx context.Context, in *AnimeUpdateReq, opts ...grpc.CallOption) (*AnimeUpdateResp, error)
		AnimeList(ctx context.Context, in *AnimeListReq, opts ...grpc.CallOption) (*AnimeListResp, error)
	}

	defaultAnimeService struct {
		cli zrpc.Client
	}
)

func NewAnimeService(cli zrpc.Client) AnimeService {
	return &defaultAnimeService{
		cli: cli,
	}
}

func (m *defaultAnimeService) AnimeAdd(ctx context.Context, in *AnimeAddReq, opts ...grpc.CallOption) (*AnimeAddResp, error) {
	client := pb.NewAnimeServiceClient(m.cli.Conn())
	return client.AnimeAdd(ctx, in, opts...)
}

func (m *defaultAnimeService) AnimeDelete(ctx context.Context, in *AnimeDeleteReq, opts ...grpc.CallOption) (*AnimeDeleteResp, error) {
	client := pb.NewAnimeServiceClient(m.cli.Conn())
	return client.AnimeDelete(ctx, in, opts...)
}

func (m *defaultAnimeService) AnimeUpdate(ctx context.Context, in *AnimeUpdateReq, opts ...grpc.CallOption) (*AnimeUpdateResp, error) {
	client := pb.NewAnimeServiceClient(m.cli.Conn())
	return client.AnimeUpdate(ctx, in, opts...)
}

func (m *defaultAnimeService) AnimeList(ctx context.Context, in *AnimeListReq, opts ...grpc.CallOption) (*AnimeListResp, error) {
	client := pb.NewAnimeServiceClient(m.cli.Conn())
	return client.AnimeList(ctx, in, opts...)
}
