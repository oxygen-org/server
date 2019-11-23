package admin

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/qor/admin"

	db "github.com/oxygen-org/server/db"
)


// InitAdmin Config Admin
func InitAdmin(e *echo.Echo) {
	mux := http.NewServeMux()
	adminConfig := &admin.AdminConfig{
		SiteName: "github.com/oxygen-org Admin",
		DB: db.DB,
	}
	a := admin.New(adminConfig)
	configVersion(a)
	configUser(a)
	configRole(a)
	configResource(a)
	configImages(a)
	configRPCService(a)
	configJob(a)
	configDispatcher(a)

	a.MountTo("/admin", mux)
	e.Any("/admin/*", echo.WrapHandler(mux))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,

		Skipper: func(e echo.Context) bool {
			uri := e.Request().RequestURI
			if strings.HasPrefix(uri, "/admin/") || "/admin" == uri {
				return true
			} else {
				return false
			}
		},
	}))
}

func configVersion(a *admin.Admin){
	a.AddResource(&db.Version{})
}

func configResource(a *admin.Admin){
	a.AddResource(&db.Resource{})
}

func configUser(a *admin.Admin){
	a.AddResource(&db.User{})
}
func configRole(a *admin.Admin){
	a.AddResource(&db.Role{})
}
func configJob(a *admin.Admin){
	a.AddResource(&db.Job{})
}
func configDispatcher(a *admin.Admin){
	a.AddResource(&db.Dispatcher{})
}

func configImages(a *admin.Admin){
	a.AddResource(&db.Images{})
}

func configRPCService(a *admin.Admin){
	a.AddResource(&db.RPCService{})
}
