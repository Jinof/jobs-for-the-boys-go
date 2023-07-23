package server

import (
	"context"
	"testing"

	"github.com/Jinof/jobs-for-the-boys/mock/mock_calculater"
	"github.com/Jinof/jobs-for-the-boys/proto/calculater"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	resp := calculater.GetModelByPathResponse{
		Matrix: []*calculater.CombinedMatrix{{
			Name: "test",
			Data: []byte("test"),
		}},
	}
	m := mock_calculater.NewMockCalculaterClient(ctrl)
	m.EXPECT().GetModelByPath(ctx, gomock.Any()).Return(&resp, nil).Times(1)

	s := server{calculaterClient: m}
	err := s.GetModelByPath(ctx)
	assert.NoError(t, err)
}
