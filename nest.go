package nest

import (
	"net/http"

	logService "github.com/vanyastar/nest/log-service"
)

// Create NestGO App from nest factory
func CreateApp(fn AppContextFunc, sArr ...*http.Server) *httpServers {
	logService.Log("NestGoFactory", "Starting application...")
	router := newRouter()
	servers := newHttpServers(router.mux, sArr...)
	appContext := newAppContext(router, servers)
	fn(appContext)
	return servers
}
