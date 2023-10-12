package main

import (
	"github.com/gin-gonic/gin"

	"github.com/penglongli/gin-metrics/ginmetrics"
)

func main() {
	r := gin.Default()


	gaugeMetric := &ginmetrics.Metric{
		Type:        ginmetrics.Gauge,
		Name:        "example_gauge_metric",
		Description: "an example of gauge type metric",
		Labels:      []string{"label1"},
	}
	
	// Add metric to global monitor object
	_ = ginmetrics.GetMonitor().AddMetric(gaugeMetric)


	// +optional set metric path, default /debug/metrics
	ginmetrics.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	// m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	// m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	ginmetrics.Use(r)

	r.GET("/product/:id", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"productId": ctx.Param("id"),
		})
	})

	r.GET("/product/add", func(ctx *gin.Context) {
		_ = ginmetrics.GetMonitor().GetMetric("example_gauge_metric").Add([]string{"label_value1"}, 0.2)

		ctx.JSON(200, map[string]string{
			"productId": "add",
		})
	})

	r.GET("/product/incr", func(ctx *gin.Context) {
		_ = ginmetrics.GetMonitor().GetMetric("example_gauge_metric").Inc([]string{"label_value1"})

		ctx.JSON(200, map[string]string{
			"productId": "incr",
		})
	})

	r.GET("/product/decr", func(ctx *gin.Context) {
		_ = ginmetrics.GetMonitor().GetMetric("example_gauge_metric").Inc([]string{"label_value1"})

		ctx.JSON(200, map[string]string{
			"productId": "incr",
		})
	})

	r.GET("/product/set", func(ctx *gin.Context) {
		_ = ginmetrics.GetMonitor().GetMetric("example_gauge_metric").SetGaugeValue([]string{"label_value1"}, 0)

		ctx.JSON(200, map[string]string{
			"productId": "set",
		})
	})

	_ = r.Run(":8088")
}
