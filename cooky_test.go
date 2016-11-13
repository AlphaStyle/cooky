package cooky

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCookie(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
		c1 := New("demo", "1234")
		c2 := New("test", "abcd")

		// Set(w, c1)
		// Set(w, c2)
		r.AddCookie(c1)
		r.AddCookie(c2)

		cookie, err := Get(r, "demo")
		if err != nil {
			t.Error(err)
		}
		expect := "demo=1234"
		if cookie.String() != expect {
			t.Errorf("got %s, expected %s\n", cookie, expect)
		}

		cookies := GetAll(r)
		expect2 := "test=abcd"
		for _, v := range cookies {
			if v.String() != expect || v.String() != expect2 {
				t.Errorf("got %s, expected %s or %s\n", v, expect, expect2)
			}
		}
	}))
	defer ts.Close()
}
