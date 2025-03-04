// Code generated by "mock-api-gen -type Consul -pkg mockconsul -endpoints kv.json -output ./kv.go"; DO NOT EDIT.

package mockconsul

import (
	"fmt"
	mockapi "github.com/panhongrainbow/mock-http-api"

	"github.com/hashicorp/consul/api"
)

func (m *Consul) KVAcquire(key string, queryParams map[string]string, body *api.KVPair, status int, reply bool) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/kv/%s", key)).WithBody(body).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) KVDelete(key string, queryParams map[string]string, body *api.KVPair, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("DELETE", fmt.Sprintf("/v1/kv/%s", key)).WithBody(body).WithQueryParams(queryParams)

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) KVDeleteCAS(key string, queryParams map[string]string, body *api.KVPair, status int, reply bool) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("DELETE", fmt.Sprintf("/v1/kv/%s", key)).WithBody(body).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) KVGet(key string, queryParams map[string]string, status int, reply *api.KVPairs) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/kv/%s", key)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) KVKeys(prefix string, queryParams map[string]string, status int, reply []string) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/kv/%s", prefix)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) KVList(prefix string, queryParams map[string]string, status int, reply *api.KVPairs) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/kv/%s", prefix)).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) KVPut(key string, queryParams map[string]string, body []byte, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/kv/%s", key)).WithBody(body).WithQueryParams(queryParams)

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) KVPutCAS(key string, queryParams map[string]string, body *api.KVPair, status int, reply bool) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/kv/%s", key)).WithBody(body).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) KVRelease(key string, queryParams map[string]string, body *api.KVPair, status int, reply bool) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/kv/%s", key)).WithBody(body).WithQueryParams(queryParams)

	return m.WithJSONReply(req, status, reply)
}
