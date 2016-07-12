package outputs

import (
  "log"
  "github.com/allen13/gollect/inputs"
)

type Stdout struct{
}

func (stdout* Stdout)Write(metric inputs.Metric){
  log.Println("stdout:" + metric.Data)
}
