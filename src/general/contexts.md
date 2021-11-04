# Contexts

## context.WithValue
Simple example of passing around a value through a context
```go
func main() {
	ctx := context.Background()
	ctx = addValue(ctx)
	readValue(ctx)
}

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "key", "test-value")
}

func readValue(ctx context.Context) {
	val := ctx.Value("key")
	fmt.Println(val)
}
```

### Example wrapping contexts
```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = addValue(ctx)
	fmt.Println(ctx.Value("first"))
}

func addValue(ctx context.Context) context.Context {
	a := context.WithValue(ctx, "first", "first-v")
	b := context.WithValue(a, "second", "second-v")
	c := context.WithValue(b, "third", "third-v")
	return c
}

```
Each Context is nested, but you can still grab the value at the bottom level

### Example passing UUID using middleware

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(guidMiddleware)
	router.HandleFunc("/ishealthy", handleIsHealthy).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}

func handleIsHealthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	uuid := r.Context().Value("uuid")
	log.Printf("[%v] Returning 200 - Healthy", uuid)
	w.Write([]byte("Healthy"))
}

func guidMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		r = r.WithContext(context.WithValue(r.Context(), "uuid", uuid))
		next.ServeHTTP(w, r)
	})
}
```
