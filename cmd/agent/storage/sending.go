package storage

type SendingOne interface {
	SendOne(metric OneMetric)
}

type SendingAll interface {
	SendAll(metrics []OneMetric)
}

func SendAll(s SendingAll, metrics []OneMetric) {
	s.SendAll(metrics)
}

func SendOne(s SendingOne, metrics OneMetric) {
	s.SendOne(metrics)
}
