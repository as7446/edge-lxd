package route

import (
	"github.com/as7446/lxd/internal/api"
	"net/http"
)

func InitRoute() {
	instanceAPI := api.NewInstanceAPI()

	http.HandleFunc("/instance/create", instanceAPI.CreateInstance)
	http.HandleFunc("/instance/delete", instanceAPI.DeleteInstance)

	http.ListenAndServe(":8080", nil)
}
