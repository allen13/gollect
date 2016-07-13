package outputs

import (
  "github.com/allen13/gollect/data"
)

type Output interface {
  Write(data data.Metric)
}