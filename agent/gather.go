package agent

import (
  "fmt"
  "log"
  "time"
  "github.com/allen13/gollect/inputs"
  "sync"
  "github.com/allen13/gollect/data"
)

func GatherAgent(metricsC chan data.Metric, shutdown chan struct{}, wg* sync.WaitGroup){
  defer wg.Done()

  fmt.Println("starting gather agent...")
  ticker := time.NewTicker(time.Second * 1)
  defer ticker.Stop()

  for{
    select {
    case <-shutdown:
      log.Println("shutting down gather agent")
      return
    case <-ticker.C:
      gatherAllInputs(metricsC, shutdown)
    }
  }
}

func gatherAllInputs(metricsC chan data.Metric, shutdown chan struct{}){
  for _,input := range []inputs.Input{&inputs.Exec{"one","1s"},&inputs.Exec{"two","2s"}}{
    go input.Gather(metricsC)
  }
}
