package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swol-api/config"
	"swol-api/controller"
	"swol-api/router"
)

var (
	appRouter      = router.NewChaiRouter()
	authController = controller.NewAuthController()
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
	appRouter.Get("/accounts", authController.All)
	appRouter.Serve(fmt.Sprintf(":%s", config.Conf.Port))
}
