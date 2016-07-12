package inputs

type Metric struct {
  Data string
}

type Input interface {
  Gather(metricsC chan Metric)
  Description() string
}
