package metrics

import "github.com/prometheus/client_golang/prometheus"

var DbCall = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "db_calls_total", Help: "Number of database call"},
	[]string{"type_name", "operation_name", "status"})
