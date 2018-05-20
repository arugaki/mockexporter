package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	"collector"
	"strconv"
	"log"
)

var (
	metricsPath = envOrDefault("PROMETHEUS_PATH", "/metrics")
	prometheusPort = envOrDefault("PROMETHEUS_PORT","9120")
)

func main() {

	n := envOrDefault("n", "48")
	m := envOrDefault("m", "100")

	nByte, err := strconv.Atoi(n)
	if err != nil {
		log.Printf("The Env n is illegel.")
		return
	}
	nMetrics, err := strconv.Atoi(m)
	if err != nil {
		log.Printf("The Env m is illegel.")
		return
	}

	c := collector.NewExporter(nByte, nMetrics)
	prometheus.MustRegister(c)
	http.Handle(metricsPath, prometheus.Handler())
	http.ListenAndServe(":"+ prometheusPort, nil)
}

func envOrDefault(env, def string) string {
	e := os.Getenv(env)
	if e != "" {
		return e
	}
	return def
}