package route

import (
	"getLocation/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Router) {
	r.Add(http.MethodGet, "/getLocation", controllers.GetLocation)
}
