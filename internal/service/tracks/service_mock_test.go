// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock_test.go -package=tracks
//

// Package tracks is a generated GoMock package.
package tracks

import (
	context "context"
	reflect "reflect"

	spotify "github.com/mdafaardiansyah/musicacu-backend/internal/repository/spotify"
	gomock "go.uber.org/mock/gomock"
)

// MockspotifyOutbound is a mock of spotifyOutbound interface.
type MockspotifyOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockspotifyOutboundMockRecorder
	isgomock struct{}
}

// MockspotifyOutboundMockRecorder is the mock recorder for MockspotifyOutbound.
type MockspotifyOutboundMockRecorder struct {
	mock *MockspotifyOutbound
}

// NewMockspotifyOutbound creates a new mock instance.
func NewMockspotifyOutbound(ctrl *gomock.Controller) *MockspotifyOutbound {
	mock := &MockspotifyOutbound{ctrl: ctrl}
	mock.recorder = &MockspotifyOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockspotifyOutbound) EXPECT() *MockspotifyOutboundMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockspotifyOutbound) Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, query, limit, offset)
	ret0, _ := ret[0].(*spotify.SpotifySearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockspotifyOutboundMockRecorder) Search(ctx, query, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockspotifyOutbound)(nil).Search), ctx, query, limit, offset)
}