package main_test

import (
	"errors"
	"testing"

	"github.com/k3forx/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	mock_main "github.com/k3forx/gomock/mock"
)

func TestSample(t *testing.T) {
	t.Parallel()
	err := errors.New("test error")
	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) main.ApiClient
		in    string
		out   string
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) main.ApiClient {
				mockApiClient := mock_main.NewApiClient(ctrl)
				mockApiClient.EXPECT().Request("bar").Return("out", nil)
				return mockApiClient
			},
			in:  "bar",
			out: "out",
		},
		"fail": {
			setup: func(ctrl *gomock.Controller) main.ApiClient {
				mockApiClient := mock_main.NewApiClient(ctrl)
				mockApiClient.EXPECT().Request("bar").Return("", err)
				return mockApiClient
			},
			in:  "bar",
			out: "",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockApiClient := c.setup(ctrl)
			d := &main.DataRegister{}
			d.Client = mockApiClient
			res, _ := d.Register(c.in)
			assert.Equal(t, res, c.out)
		})
	}
}
