package service

import (
	"log"
	"net/http"
	"runtime"

	"github.com/les-cours/gateway/env"
	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
)

var (
	registry       = prometheus.NewRegistry()
	requestCounter = prometheus.NewGauge(prometheus.GaugeOpts{Name: "request_Counter", Help: "count request that hit"})
	memoryUsage    = prometheus.NewGauge(prometheus.GaugeOpts{Name: "gateway_memory_usage", Help: "gateway memory usage"})
	goRoutineNum   = prometheus.NewGauge(prometheus.GaugeOpts{Name: "go_routines_num", Help: "the number of go routine "})
	cpuPercentage  = prometheus.NewGauge(prometheus.GaugeOpts{Name: "cpu_percentage", Help: "cpu percentage"})
)

func monitoringMiddleware(originalHandler http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memoryUsage.Set(float64(m.Alloc))
		goRoutineNum.Set(float64(runtime.NumGoroutine()))
		percent, _ := cpu.Percent(0, false)
		cpuPercentage.Set(percent[0])
		originalHandler.ServeHTTP(w, r)
	})
}
func Start() {
	registry.MustRegister(requestCounter, memoryUsage, cpuPercentage, goRoutineNum)
	log.Println(env.Settings.LearningApiURL)
	schemas, err := graphql.IntrospectRemoteSchemas(checkApis(
		env.Settings.UserApiURL,
		env.Settings.LearningApiURL,
		//env.Settings.Apis ...
	)...)
	if err != nil {
		log.Fatalf("Error in IntrospectRemoteSchemas: %v", err)
	}

	gw, err := gateway.New(schemas, gateway.WithMiddlewares(forwardUser))
	if err != nil {
		log.Fatalf("Error when creating gateway: %v", err)
	}
	http.HandleFunc("/api", cors(decodeToken(gw.GraphQLHandler)))
	promHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	http.HandleFunc("/metrics", cors(decodeToken(monitoringMiddleware(promHandler))))
	http.HandleFunc("/", cors(decodeToken(gw.PlaygroundHandler)))
	log.Printf("Starting https server on port " + env.Settings.HttpPort)
	log.Printf("Starting graphql gateway...")
	err = http.ListenAndServe(":"+env.Settings.HttpPort, nil)
	if err != nil {
		log.Fatalf("Error https server on port %v: %v", env.Settings.HttpPort, err)
	}
}
