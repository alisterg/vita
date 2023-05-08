package api

import (
	"fmt"
	"net/http"
)

func LambdaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
