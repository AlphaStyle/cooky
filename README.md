# Cooky

## This package is for learning purposes.

### Example

```go
// New will create a new Cookie
Cookie := cooky.New("name", "value")

// Get will return the "named" cookie
cooky.Get(r, "Name/key")

// GetAll will return all cookies from the request
cooky.GetAll(r)

// Add a Cookie to the request
cooky.Set(w, Cookie)

// Delete a Cookie from the request
cooky.Delete(w, Cookie)
```