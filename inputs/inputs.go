package inputs

import "github.com/allen13/gollect/data"

type Input interface {
  Gather(metricsC chan data.Metric)
  Description() string
}
