package main

import (
	"encoding/json"
	"net/http"
	"swol-api/router"
)

var (
	appRouter router.Router = router.NewChaiRouter()
)

func main() {
	appRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Success bool `json:"success"`
		}{
			Success: true,
		})
	})
	appRouter.Serve(":8000")
}
