package agent

import (
	"time"

	"github.com/torotonnato/gobarebones/model"
)

func PushMetric(m *model.Metric, value float64) error {
	if !state.isRunning {
		return Error{AgentNotRunning}
	}
	item := MetricItem{
		ID: m.ID,
		Point: model.Point{
			Value:     value,
			Timestamp: time.Now().Unix(),
		},
	}
	state.dataChan.WriteChan <- item
	return nil
}
