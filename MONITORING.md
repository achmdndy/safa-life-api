# Monitoring Setup Guide 📊

This guide explains how to set up and use monitoring for the Safa Life API using **Jaeger** for distributed tracing and **Prometheus** for metrics collection.

## 🎯 Features

- **Prometheus Metrics**: HTTP request metrics, application info, and Go runtime metrics
- **Jaeger Tracing**: Distributed tracing for request flow analysis
- **Health Monitoring**: Application health check endpoints
- **Docker Integration**: Easy setup with Docker Compose
- **Grafana Dashboard**: Optional visualization (included in docker-compose)

## 🚀 Quick Start

### 1. Start Monitoring Stack

```bash
# Start Jaeger and Prometheus
docker-compose -f docker-compose.monitoring.yml up -d

# Or start complete stack (app + monitoring + database)
docker-compose up -d
```

### 2. Start Application

```bash
# Build and run
go build -o bin/app ./src/cmd/server/main.go
./bin/app start

# Or run directly
go run src/cmd/server/main.go start
```

### 3. Verify Setup

```bash
# Check application health
curl http://localhost:8080/health

# Check metrics endpoint
curl http://localhost:8080/metrics

# Test API endpoint
curl http://localhost:8080/api/v1/health
```

## 🔗 Access Monitoring Tools

| Service | URL | Credentials |
|---------|-----|-------------|
| **Application** | http://localhost:8080 | - |
| **Jaeger UI** | http://localhost:16686 | - |
| **Prometheus** | http://localhost:9090 | - |
| **Grafana** | http://localhost:3000 | admin/admin |

## ⚙️ Configuration

### Jaeger Configuration

The application automatically connects to Jaeger using OpenTelemetry:

```go
// Jaeger endpoint
jaegerEndpoint := "http://localhost:14268/api/traces"

// Service configuration
serviceName := "safa-life-api"
serviceVersion := "1.0.0"
environment := "development"
```

### Prometheus Configuration

Prometheus scrapes metrics from the application:

```yaml
# prometheus.yml
scrape_configs:
  - job_name: 'safa-life-api'
    static_configs:
      - targets: ['host.docker.internal:8080']  # For Docker on Mac/Windows
      # - targets: ['172.17.0.1:8080']         # For Docker on Linux
    metrics_path: '/metrics'
    scrape_interval: 15s
```

## 📊 Available Metrics

### HTTP Request Metrics

```
# Total HTTP requests
http_requests_total{method="GET", endpoint="/health", status="200"}

# HTTP requests currently in flight
http_requests_in_flight

# HTTP request duration
http_request_duration_seconds{method="GET", endpoint="/health"}
```

### Application Metrics

```
# Application information
application_info{version="1.0.0", service="safa-life-api"}

# Go runtime metrics
go_goroutines
go_gc_duration_seconds
go_memstats_alloc_bytes
```

### Custom Business Metrics

You can add custom metrics in your handlers:

```go
// Example: Counter for specific business events
businessEventCounter := prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "business_events_total",
        Help: "Total number of business events",
    },
    []string{"event_type"},
)

// Usage in handler
businessEventCounter.WithLabelValues("user_registration").Inc()
```

## 🔍 Using Jaeger Tracing

### Viewing Traces

1. Open Jaeger UI: http://localhost:16686
2. Select service: `safa-life-api`
3. Click "Find Traces"
4. Explore request flows and performance

### Trace Information

Each trace includes:
- **Request ID**: Unique identifier for each request
- **Span Duration**: Time taken for each operation
- **HTTP Details**: Method, URL, status code, headers
- **Error Information**: Stack traces and error messages

### Custom Spans

Add custom spans for detailed tracing:

```go
func (h *HealthHandler) GetHealth(c *gin.Context) {
    // Get tracer from context
    tracer := otel.Tracer("safa-life-api")
    
    // Start custom span
    ctx, span := tracer.Start(c.Request.Context(), "health-check-operation")
    defer span.End()
    
    // Your business logic here
    result := h.healthUseCase.CheckHealth(ctx)
    
    // Add attributes to span
    span.SetAttributes(
        attribute.String("health.status", result.Status),
        attribute.Int("health.checks", len(result.Checks)),
    )
    
    c.JSON(200, result)
}
```

## 🐳 Docker Compose Setup

### Monitoring Only

```yaml
# docker-compose.monitoring.yml
version: '3.8'
services:
  jaeger:
    image: jaegertracing/all-in-one:latest
    ports:
      - "16686:16686"
      - "14268:14268"
    environment:
      - COLLECTOR_OTLP_ENABLED=true

  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
```

### Complete Stack

```yaml
# docker-compose.yml
version: '3.8'
services:
  postgres:
    image: postgres:14
    environment:
      POSTGRES_DB: safa_life
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"

  redis:
    image: redis:6-alpine
    ports:
      - "6379:6379"

  # ... monitoring services
```

## 🧪 Testing Monitoring

### Generate Test Traffic

```bash
# Generate multiple requests
for i in {1..10}; do
  curl http://localhost:8080/health
  curl http://localhost:8080/api/v1/health
  sleep 1
done

# Check metrics
curl http://localhost:8080/metrics | grep http_requests_total
```

### Verify Traces

1. Generate requests with errors:
   ```bash
   curl http://localhost:8080/nonexistent-endpoint
   ```

2. Check Jaeger UI for error traces
3. Verify error spans and stack traces

## 🚨 Alerting (Optional)

### Prometheus Alerting Rules

```yaml
# alerts.yml
groups:
  - name: safa-life-api
    rules:
      - alert: HighErrorRate
        expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.1
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High error rate detected"

      - alert: HighResponseTime
        expr: histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) > 1
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High response time detected"
```

## 🎯 Benefits

### Performance Monitoring
- **Real-time Metrics**: Monitor request rates, response times, and error rates
- **Resource Usage**: Track memory, CPU, and goroutine usage
- **Bottleneck Identification**: Find slow operations and optimize performance

### Debugging & Troubleshooting
- **Distributed Tracing**: Follow requests across services
- **Error Tracking**: Detailed error information with stack traces
- **Request Flow**: Understand how requests flow through the application

### Operational Insights
- **Health Monitoring**: Continuous health checks and alerts
- **Capacity Planning**: Historical data for scaling decisions
- **SLA Monitoring**: Track service level objectives

## 🔧 Troubleshooting

### Common Issues

1. **Jaeger not receiving traces**
   - Check if application can reach Jaeger endpoint
   - Verify OpenTelemetry configuration
   - Check firewall settings

2. **Prometheus not scraping metrics**
   - Verify target configuration in prometheus.yml
   - Check if metrics endpoint is accessible
   - Ensure correct network configuration

3. **High memory usage**
   - Adjust trace sampling rate
   - Configure metric retention policies
   - Monitor resource limits

### Debug Commands

```bash
# Check if services are running
docker-compose ps

# View service logs
docker-compose logs jaeger
docker-compose logs prometheus

# Test connectivity
curl http://localhost:8080/metrics
curl http://localhost:16686/api/services
```

## 📚 Additional Resources

- [Prometheus Documentation](https://prometheus.io/docs/)
- [Jaeger Documentation](https://www.jaegertracing.io/docs/)
- [OpenTelemetry Go Documentation](https://opentelemetry.io/docs/instrumentation/go/)
- [Grafana Documentation](https://grafana.com/docs/)

---

**🎯 With proper monitoring, you can ensure your Safa Life API runs smoothly and efficiently in production!**