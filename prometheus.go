package main

import (
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "log"
  "net/http"
)

func prometheusHandler(prometheus bool) {
  if prometheus {
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8090", nil))
  }
}
