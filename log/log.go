package log

import (
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// MLog is a base parent logger for the telemetry package
var MLog = logf.Log.WithName("telemetry")
