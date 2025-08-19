package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Initial Test
// func TestRacer(t *testing.T) {
// 	slowURL := "http://www.facebook.com"
// 	fastURL := "http://www.quii.dev"

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// ----------------------------------------------------------------------------------- //
// Test 2: with mock server
// func TestRacer(t *testing.T) {
// 	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(20 * time.Millisecond)
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	slowServer := makeDelayedServer(20 * time.Millisecond)
// 	defer slowServer.Close()
// 	fastServer := makeDelayedServer(1 * time.Millisecond)
// 	defer fastServer.Close()

// 	slowURL := slowServer.URL
// 	fastURL := fastServer.URL

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// ----------------------------------------------------------------------------------- //
// Test 3
// func TestRacer(t *testing.T) {

// 	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
// 		slowServer := makeDelayedServer(20 * time.Millisecond)
// 		defer slowServer.Close()
// 		fastServer := makeDelayedServer(0 * time.Millisecond)
// 		defer fastServer.Close()

// 		want := fastServer.URL
// 		got, _ := Racer(slowServer.URL, fastServer.URL)

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})

// 	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
// 		server1 := makeDelayedServer(2 * time.Second)
// 		defer server1.Close()
// 		server2 := makeDelayedServer(3 * time.Second)
// 		defer server2.Close()

// 		_, err := Racer(server1.URL, server2.URL)

// 		if err == nil {
// 			t.Error("expected an error but didn't get one")
// 		}
// 	})
// }

// ----------------------------------------------------------------------------------- //
// Test 4
func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(20 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
