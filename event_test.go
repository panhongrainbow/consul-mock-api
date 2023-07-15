package mockconsul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestMyAPI_Event2(t *testing.T) {
	m := NewConsul(t)

	// http.Get will add both of the headers but we don't want to care about them.
	m.SetFilteredHeaders([]string{
		"Content-Length",
		"Content-Type",
		"Accept-Encoding",
		"User-Agent",
	})

	// set up an expectation for a GET /v1/acl/tokens
	m.EventFire("my-event", []byte("Lorem ipsum dolor sit amet, consectetur adipisicing elit..."), 200, `{"ID":"f54e9a29-acc3-d573-aca4-6f5c751b3b06","Name":"my-event","Payload":"TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2ljaW5nIGVsaXQuLi4=","NodeFilter":"","ServiceFilter":"","TagFilter":"","Version":1,"LTime":0}`).Once()

	cfg := api.DefaultConfig()
	cfg.Address = m.URL()

	client, err := api.NewClient(cfg)
	if err != nil {
		t.Fatalf("error when creating a new HTTP client: %v", err)
	}

	// 建立Event
	event := &api.UserEvent{
		Name:    "my-event",
		Payload: []byte("Lorem ipsum dolor sit amet, consectetur adipisicing elit..."),
	}

	// 發送Event
	test1, test2, err := client.Event().Fire(event, nil)
	fmt.Println(test1, test2)
	/*if err != nil {
		log.Fatal("Failed to fire event:", err)
	}*/

	fmt.Println("Event fired successfully.")
}
