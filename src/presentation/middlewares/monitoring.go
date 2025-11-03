package middlewares

import (
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// TracingMiddleware adds OpenTelemetry tracing to HTTP requests
func TracingMiddleware(monitoringService interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use reflection to call StartSpan if available
		if service := reflect.ValueOf(monitoringService); service.IsValid() {
			if method := service.MethodByName("StartSpan"); method.IsValid() {
				// Call StartSpan method
				results := method.Call([]reflect.Value{
					reflect.ValueOf(c.Request.Context()),
					reflect.ValueOf("HTTP " + c.Request.Method + " " + c.FullPath()),
				})

				if len(results) >= 2 {
					// Update context if returned
					if ctx, ok := results[0].Interface().(context.Context); ok {
						c.Request = c.Request.WithContext(ctx)
					}

					// End span when request completes
					if span := results[1]; span.IsValid() {
						defer func() {
							if endMethod := span.MethodByName("End"); endMethod.IsValid() {
								endMethod.Call([]reflect.Value{})
							}
						}()

						// Set attributes
						if setAttrMethod := span.MethodByName("SetAttributes"); setAttrMethod.IsValid() {
							// Create attributes using reflection
							attrs := []reflect.Value{}

							// Add HTTP method attribute
							if newStringAttr := service.MethodByName("NewStringAttribute"); newStringAttr.IsValid() {
								attr := newStringAttr.Call([]reflect.Value{
									reflect.ValueOf("http.method"),
									reflect.ValueOf(c.Request.Method),
								})
								if len(attr) > 0 {
									attrs = append(attrs, attr[0])
								}
							}

							// Add URL attribute
							if newStringAttr := service.MethodByName("NewStringAttribute"); newStringAttr.IsValid() {
								attr := newStringAttr.Call([]reflect.Value{
									reflect.ValueOf("http.url"),
									reflect.ValueOf(c.Request.URL.String()),
								})
								if len(attr) > 0 {
									attrs = append(attrs, attr[0])
								}
							}

							if len(attrs) > 0 {
								setAttrMethod.Call(attrs)
							}
						}
					}
				}
			}
		}

		c.Next()
	}
}

// MetricsMiddleware records HTTP request metrics
func MetricsMiddleware(monitoringService interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Increment in-flight requests
		if service := reflect.ValueOf(monitoringService); service.IsValid() {
			if method := service.MethodByName("IncrementInFlightRequests"); method.IsValid() {
				method.Call([]reflect.Value{})
			}
		}

		c.Next()

		// Decrement in-flight requests and record metrics
		if service := reflect.ValueOf(monitoringService); service.IsValid() {
			if method := service.MethodByName("DecrementInFlightRequests"); method.IsValid() {
				method.Call([]reflect.Value{})
			}

			// Record HTTP request
			if method := service.MethodByName("RecordHTTPRequest"); method.IsValid() {
				duration := time.Since(start)
				method.Call([]reflect.Value{
					reflect.ValueOf(c.Request.Method),
					reflect.ValueOf(c.FullPath()),
					reflect.ValueOf(strconv.Itoa(c.Writer.Status())),
					reflect.ValueOf(duration),
				})
			}
		}
	}
}

// MonitoringMiddleware combines tracing and metrics
func MonitoringMiddleware(monitoringService interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Increment in-flight requests
		if service := reflect.ValueOf(monitoringService); service.IsValid() {
			if method := service.MethodByName("IncrementInFlightRequests"); method.IsValid() {
				method.Call([]reflect.Value{})
			}
		}

		// Start tracing span
		var spanValue reflect.Value
		if service := reflect.ValueOf(monitoringService); service.IsValid() {
			if method := service.MethodByName("StartSpan"); method.IsValid() {
				results := method.Call([]reflect.Value{
					reflect.ValueOf(c.Request.Context()),
					reflect.ValueOf("HTTP " + c.Request.Method + " " + c.FullPath()),
				})

				if len(results) >= 2 {
					if ctx, ok := results[0].Interface().(context.Context); ok {
						c.Request = c.Request.WithContext(ctx)
					}
					spanValue = results[1]
				}
			}
		}

		c.Next()

		// End span and record metrics
		if spanValue.IsValid() {
			// Set span attributes
			if setAttrMethod := spanValue.MethodByName("SetAttributes"); setAttrMethod.IsValid() {
				if service := reflect.ValueOf(monitoringService); service.IsValid() {
					// Create attributes using reflection
					attrs := []reflect.Value{}

					// Add HTTP method attribute
					if newStringAttr := service.MethodByName("NewStringAttribute"); newStringAttr.IsValid() {
						attr := newStringAttr.Call([]reflect.Value{
							reflect.ValueOf("http.method"),
							reflect.ValueOf(c.Request.Method),
						})
						if len(attr) > 0 {
							attrs = append(attrs, attr[0])
						}
					}

					// Add HTTP status code attribute
					if newIntAttr := service.MethodByName("NewIntAttribute"); newIntAttr.IsValid() {
						attr := newIntAttr.Call([]reflect.Value{
							reflect.ValueOf("http.status_code"),
							reflect.ValueOf(int64(c.Writer.Status())),
						})
						if len(attr) > 0 {
							attrs = append(attrs, attr[0])
						}
					}

					// Add URL attribute
					if newStringAttr := service.MethodByName("NewStringAttribute"); newStringAttr.IsValid() {
						attr := newStringAttr.Call([]reflect.Value{
							reflect.ValueOf("http.url"),
							reflect.ValueOf(c.Request.URL.String()),
						})
						if len(attr) > 0 {
							attrs = append(attrs, attr[0])
						}
					}

					if len(attrs) > 0 {
						setAttrMethod.Call(attrs)
					}
				}
			}

			// End span
			if endMethod := spanValue.MethodByName("End"); endMethod.IsValid() {
				endMethod.Call([]reflect.Value{})
			}
		}

		// Decrement in-flight requests and record metrics
		if service := reflect.ValueOf(monitoringService); service.IsValid() {
			if method := service.MethodByName("DecrementInFlightRequests"); method.IsValid() {
				method.Call([]reflect.Value{})
			}

			// Record HTTP request
			if method := service.MethodByName("RecordHTTPRequest"); method.IsValid() {
				duration := time.Since(start)
				method.Call([]reflect.Value{
					reflect.ValueOf(c.Request.Method),
					reflect.ValueOf(c.FullPath()),
					reflect.ValueOf(strconv.Itoa(c.Writer.Status())),
					reflect.ValueOf(duration),
				})
			}
		}
	}
}
