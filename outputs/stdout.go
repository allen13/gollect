package outputs

import (
  "log"
  "github.com/allen13/gollect/data"
  "strconv"
)

type Stdout struct{
}

func (stdout* Stdout)Write(metric data.Metric){
  fields := metric.Fields()
  log.Println("stdout:" + "It took " + metric.Name() + " " + strconv.FormatFloat(fields["responseTimeSeconds"].(float64), 'f', 0, 64) + " seconds to respond")
}
