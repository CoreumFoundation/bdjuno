package logging

import (
	"github.com/prometheus/client_golang/prometheus"
)

// ActionResponseTime represents the Telemetry counter used to classify each executed action by response time
var ActionResponseTime = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "bdjuno_action_response_time",
		Help:    "Time it has taken to execute an action",
		Buckets: []float64{0.5, 1, 2, 3, 4, 5},
	}, []string{"path"})

// ActionCounter represents the Telemetry counter used to track the total number of actions executed
var ActionCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bdjuno_actions_total_count",
		Help: "Total number of actions executed.",
	}, []string{"path", "http_status_code"})

// ActionErrorCounter represents the Telemetry counter used to track the number of action's errors emitted
var ActionErrorCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bdjuno_actions_error_count",
		Help: "Total number of errors emitted.",
	}, []string{"path", "http_status_code"},
)

// BlockTimeGauge represents the Telemetry gauge used to track chain block time
var BlockTimeGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "bdjuno_block_time",
		Help: "The current bdjuno block time.",
	}, []string{
		"period",
	},
)

// MissedProposerCounter represents the Telemetry counter used to track missed proposals
var MissedProposerCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bdjuno_missed_proposer",
		Help: "How many times each validator missed to propose the block.",
	}, []string{
		"validator",
	},
)

func init() {
	err := prometheus.Register(ActionResponseTime)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(ActionCounter)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(ActionErrorCounter)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(BlockTimeGauge)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(MissedProposerCounter)
	if err != nil {
		panic(err)
	}
}
