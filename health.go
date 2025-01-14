// Code generated by "mock-api-gen -type Consul -pkg mockconsul -endpoints health.json -output ./health.go"; DO NOT EDIT.

package mockconsul

import (
	"fmt"
	mockapi "github.com/panhongrainbow/mock-http-api"

	"github.com/hashicorp/consul/api"
)

func (m *Consul) HealthChecks(service string, queryParams map[string]string, status int, reply api.HealthChecks) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/health/checks/%s", service)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) HealthConnect(service string, queryParams map[string]string, status int, reply []*api.ServiceEntry) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/health/connect/%s", service)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) HealthIngress(service string, queryParams map[string]string, status int, reply []*api.ServiceEntry) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/health/ingress/%s", service)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) HealthNode(node string, queryParams map[string]string, status int, reply api.HealthChecks) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/health/node/%s", node)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) HealthService(service string, queryParams map[string]string, status int, reply []*api.ServiceEntry) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/health/service/%s", service)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}
