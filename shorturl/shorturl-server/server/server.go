package server

import (
	"context"
	"shorturl/pkg/config"
	"shorturl/pkg/log"
	proto "shorturl/proc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type shortUrlService struct {
	proto.UnimplementedShortUrlServer
	config *config.Config
	log    log.ILogger
}

func NewService(config *config.Config, logger log.ILogger) proto.ShortUrlServer {
	return &shortUrlService{config: config, log: logger}
}

func (s *shortUrlService) GetShortUrl(ctx context.Context, in *proto.Url) (*proto.Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortUrl not implemented")
}
func (s *shortUrlService) GetOriginalUrl(ctx context.Context, in *proto.ShortKey) (*proto.Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalUrl not implemented")
}
