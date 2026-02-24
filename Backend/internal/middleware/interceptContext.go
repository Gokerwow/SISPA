package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Extract the current context and wrap it with a timeout.
		timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		// Step 2: Defer the cancellation.
		defer cancel()

		// Step 3: Replace the HTTP Request inside Gin with a new one that holds your timeout context.
		c.Request = c.Request.WithContext(timeoutCtx)

		// Step 4: Tell Gin to proceed to the actual Handler.
		c.Next()
	}
}
