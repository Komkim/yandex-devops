package router

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func (h *Router) TrustedSubnetMiddleware(c *gin.Context) {
	ipstr := c.RemoteIP()
	if len(ipstr) > 0 {
		ip := net.ParseIP(ipstr)
		_, cidr, err := net.ParseCIDR(h.cfg.TrustedSubnet)
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "bad trusted subnet")
			return
		}
		if !cidr.Contains(ip) {
			respondWithError(c, 403, "Forbidden")
			return
		}
	}

	c.Next()
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
