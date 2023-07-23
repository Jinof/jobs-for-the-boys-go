package server

import (
	"context"

	"github.com/Jinof/jobs-for-the-boys/proto/calculater"
)

type Server interface {
	GetModelByPath(ctx context.Context) error
}

var _ Server = (*server)(nil)

type server struct {
	calculaterClient calculater.CalculaterClient

	matrx []*calculater.CombinedMatrix
}

func (s *server) GetModelByPath(ctx context.Context) error {
	resp, err := s.calculaterClient.GetModelByPath(ctx, &calculater.GetModelByPathRequest{})
	if err != nil {
		return err
	}

	s.matrx = resp.Matrix
	return nil
}
