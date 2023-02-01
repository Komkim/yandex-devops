package agent

import (
	"context"
	"fmt"
	"time"
	"yandex-devops/config"
	myclient "yandex-devops/provider"
)

func SendMetric(ctx context.Context, cfg *config.Agent, client *myclient.MyClient, chMetrics <-chan *[]myclient.Metrics, counter *Counter) error {
	ticker := time.NewTicker(cfg.Report)
	var metrics *[]myclient.Metrics

	for {
		select {

		case <-ticker.C:
			err := client.SendAllMetric(metrics)
			if err != nil {
				return err
			}
			counter.Reset()

		case metrics = <-chMetrics:

		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		}
	}
}
