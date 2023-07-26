package handlers

import (
	"errors"
	"fmt"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/pkg/response"
	"github.com/ka1i/cli/pkg/logger"

	"go.uber.org/zap"
)

func SetHandlers(engine *gin.Engine) {
	if err := engine.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	// Trust Upstream Header "X-Real-Ip"
	engine.TrustedPlatform = "X-Real-Ip"

	// registry gin middleware
	engine.Use(ginLogger(), ginRecovery())
}

func ginLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		ctx.Next()

		// Stop timer
		cost := time.Since(start)

		// log process
		responStatus := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.String("ip", ctx.ClientIP()),
			zap.Int("status", responStatus),
			zap.String("method", ctx.Request.Method),
			zap.String("query", ctx.Request.URL.Path),
			zap.String("proto", ctx.Request.Proto),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("time", cost.String()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		}

		if responStatus >= 400 && responStatus <= 499 {
			logger.GinWarn(fmt.Sprintf("%s Access [%d]", ctx.Request.Proto, responStatus), logFields...)
		} else if responStatus >= 500 && responStatus <= 599 {
			logger.GinError(fmt.Sprintf("%s Access [%d]", ctx.Request.Proto, responStatus), logFields...)
		} else {
			logger.GinInfo(fmt.Sprintf("%s Access [%d]", ctx.Request.Proto, responStatus), logFields...)
		}
	}
}

func ginRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)

				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					logger.GinError("HTTP broken pipe", zap.Any("error", err), zap.String("request", string(httpRequest)))
					response.Abnormal(ctx, nil, "Internal Server Error")
					ctx.Error(err.(error)) //nolint: errcheck
					ctx.Abort()
					return
				} else {
					logger.GinError("Panic recovery", zap.Any("error", err), zap.String("request", string(httpRequest)), zap.Stack("stacktrace"))
					response.Abnormal(ctx, nil, "Internal Server Error")
					ctx.Abort()
				}
			}
		}()
		ctx.Next()
	}
}
