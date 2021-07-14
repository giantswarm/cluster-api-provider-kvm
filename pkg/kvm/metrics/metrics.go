package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/giantswarm/cluster-api-provider-kvm/controllers"
)

const (
	PrometheusNamespace = "operatorkit"
	PrometheusSubsystem = "controller"
)

var (
	lastReconciledGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: PrometheusNamespace,
			Subsystem: PrometheusSubsystem,
			Name:      "last_reconciled",
			Help:      "Last reconciled Timestamp of watched runtime objects.",
		},
		[]string{"controller"},
	)
)

func init() {
	prometheus.MustRegister(lastReconciledGauge)
}

// CaptureLastReconciled will monitor and capture metrics.
func CaptureLastReconciled(controller string) {
	lastReconciledGauge.WithLabelValues(
		controllers.KVMClusterController,
	).SetToCurrentTime()
}
