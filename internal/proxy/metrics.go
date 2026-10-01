package proxy

import (
	"sync/atomic"
)

type MetricsTracker struct {
	totalRequests uint64
	totalBlocked  uint64
}

func NewMetricsTracker() *MetricsTracker {
	return &MetricsTracker{}
}

func (m *MetricsTracker) IncRequests() {
	atomic.AddUint64(&m.totalRequests, 1)
}

func (m *MetricsTracker) IncBlocked() {
	atomic.AddUint64(&m.totalBlocked, 1)
}

func (m *MetricsTracker) Snapshot() (uint64, uint64) {
	return atomic.LoadUint64(&m.totalRequests), atomic.LoadUint64(&m.totalBlocked)
}
