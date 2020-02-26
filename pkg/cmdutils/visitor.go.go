package cmdutils

import (
	"os"

	"github.com/stretchr/testify/mock"

	"github.com/golang/protobuf/jsonpb"
	"github.com/joshcarp/sysl-printing/pkg/sysl"
)

func (m *mockEndpointLabeler) LabelEndpoint(p *EndpointLabelerParam) string {
	args := m.Called(p)

	return args.String(0)
}

func readModule(p string) (*sysl.Module, error) {
	m := &sysl.Module{}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	if err := jsonpb.Unmarshal(f, m); err != nil {
		return nil, err
	}
	return m, nil
}

type Labeler struct{}

func (l *Labeler) LabelApp(appName, controls string, attrs map[string]*sysl.Attribute) string {
	return appName
}

func (l *Labeler) LabelEndpoint(p *EndpointLabelerParam) string {
	return p.EndpointName
}

type mockEndpointLabeler struct {
	mock.Mock
}
