package storage

type Sending interface {
	SendOne(metric OneMetric)
	SendAll(metrics []OneMetric)
}
