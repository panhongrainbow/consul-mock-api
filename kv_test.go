package mockconsul

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/require"
	"testing"
)

// Test_KVPut is used for unit testing to correct the KVPut function.
func Test_KVPut(t *testing.T) {
	// Create a new Consul mock instance
	consulMock := NewConsul(t)

	// Set the filtered headers
	consulMock.SetFilteredHeaders([]string{
		"Accept-Encoding",
		"Content-Length",
		"Content-Type",
		"User-Agent",
	})

	// Define a struct for the value
	type value struct {
		ID int `json:"id"`
	}

	// Create a new value instance
	v := value{
		ID: 1,
	}

	// Marshal the value to JSON
	jsonValue, err := json.Marshal(v)
	require.NoError(t, err)

	// Set up the mock for KVPut
	consulMock.KVPut("KVPut_test", nil, jsonValue, 200).Once()

	// Create a new Consul client instance
	cfg := api.DefaultConfig()
	cfg.Address = consulMock.URL()
	client, err := api.NewClient(cfg)
	require.NoError(t, err)

	// Call KV().Put() on the client and record the RTT in writeMeta
	// When the function makes a request, it returns the RTT, which is recorded in writeMeta.
	var writeMeta *api.WriteMeta
	writeMeta, err = client.KV().Put(&api.KVPair{Key: "KVPut_test", Value: jsonValue}, nil)
	require.NoError(t, err)

	// Assert that the RTT is greater than zero
	require.Less(t, int64(0), writeMeta.RequestTime.Nanoseconds())
}
