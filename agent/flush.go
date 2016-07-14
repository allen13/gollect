package agent

import (
  "fmt"
  "log"
  "sync"
  "github.com/allen13/gollect/outputs"
  "github.com/allen13/gollect/data"
)

func FlushAgent(metricsC chan data.Metric, shutdown chan struct{}, wg* sync.WaitGroup){
  defer wg.Done()

  fmt.Println("starting flush agent...")

  for{
    select {
    case <-shutdown:
      log.Println("shutting down flush agent")
      return
    case metric := <-metricsC:
      flushToOutputs(metric, shutdown)
    }
  }
}

func flushToOutputs(metric data.Metric, shutdown chan struct{}){
  for _,output := range outputs.Outputs {
    go output.Write([]data.Metric{metric})
  }
}
