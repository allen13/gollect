package outputs

import "github.com/allen13/gollect/inputs"

type Output interface {
  Write(data inputs.Metric)
}