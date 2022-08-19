package grpc

import (
	"context"
	"net"

	pbhighscore "github.com/jacksonbeecher/m-apis/m-highscore/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 9999999.0

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest, opts ...grpc.CallOption) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil //Return no error.
}

func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest, opts ...grpc.CallOption) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil //Return no error.
}

func (g *Grpc) ListenAndServer() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "Failed to open tcp port.")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("Starting grpc server for m-highscore microservice.")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "Failed to start gRPC server for highscore microservice.")
	}
	return nil
}
