package mockconsul

import mockapi "github.com/panhongrainbow/mock-http-api"

type Consul struct {
	*mockapi.MockAPI
}

func NewConsul(t mockapi.TestingT) *Consul {
	return &Consul{
		MockAPI: mockapi.NewMockAPI(t),
	}
}
