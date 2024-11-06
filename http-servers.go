package nest

import (
	"context"
	"log"
	"net/http"
	"time"

	logService "github.com/vanyastar/nest/log-service"
)

type httpServers struct {
	servers []*http.Server
	mux     *http.ServeMux
}

// Start all added servers
func (s *httpServers) ListenAndServe() {
	for _, server := range s.servers {
		server.Handler = s.mux
		go func() {
			if server.TLSConfig != nil && len(server.TLSConfig.Certificates) > 0 {
				logService.Log("HTTP/2 Server", "This application is running on: "+server.Addr)
				server.ListenAndServeTLS("", "")
			} else {
				logService.Log("HTTP/1 Server", "This application is running on: "+server.Addr)
				server.ListenAndServe()
			}
		}()
	}
}

func (s *httpServers) Shutdown() {
	for _, server := range s.servers {
		go func() {
			shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
			defer shutdownRelease()
			err := server.Shutdown(shutdownCtx)
			if err != nil {
				log.Fatal(err)
				return
			}
			logService.Log("HTTP Server", "Shutdown gracefuly: "+server.Addr)
		}()
	}
}

func newHttpServers(handler *http.ServeMux, sArr ...*http.Server) *httpServers {
	return &httpServers{
		mux:     handler,
		servers: sArr,
	}
}
