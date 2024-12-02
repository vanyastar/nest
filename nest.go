package nest

import (
	"net/http"

	"github.com/vanyastar/nest/nestlog"
)

// Create NestGO App from nest factory
func CreateApp(fn AppContextFunc, sArr ...*http.Server) *httpServers {
	nestlog.Log("NestGoFactory", "Starting application...")
	router := newRouter()
	servers := newHttpServers(router.mux, sArr...)
	appContext := newAppContext(router, servers)
	fn(appContext)
	return servers
}
