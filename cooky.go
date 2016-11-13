// Package cooky is the package for liby that handles the cookies
package cooky

import (
	"net/http"
	"time"
)

// Get will get a single cookie from request
func Get(r *http.Request, name string) (*http.Cookie, error) {
	c, err := r.Cookie(name)
	return c, err
}

// GetAll will get all the cookies from  request
func GetAll(r *http.Request) []*http.Cookie {
	return r.Cookies()
}

// Set will add a cookie to the request
func Set(w http.ResponseWriter, c *http.Cookie) {
	http.SetCookie(w, c)
}

// Delete will delete the cookie
func Delete(w http.ResponseWriter, c *http.Cookie) {
	// TODO not working cause happen after header load
	c.MaxAge = -1
	c.Expires = time.Now()
	http.SetCookie(w, c)
}

// New will create new cookie
func New(name string, value string) *http.Cookie {
	// expiration = 1 day
	expiration := time.Now().Add(24 * time.Hour)

	// New cookie
	c := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expiration,
		Path:    "/",
	}

	// return the cookie
	return c
}
