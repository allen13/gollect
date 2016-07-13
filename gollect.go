package main

import (
  "os"
  "os/signal"
  "log"
  "github.com/allen13/gollect/agent"
  "sync"
  "github.com/allen13/gollect/data"
)


func main(){

  shutdown := createShutdownChannel()
  metricsC := make(chan data.Metric)

  var wg sync.WaitGroup
  wg.Add(2)

  go agent.FlushAgent(metricsC, shutdown, &wg)
  go agent.GatherAgent(metricsC, shutdown, &wg)

  wg.Wait()
  log.Println("shutting down gollect")
}

func createShutdownChannel()(chan struct{}){
  shutdown := make(chan struct{})
  signals := make(chan os.Signal)
  signal.Notify(signals, os.Interrupt)

  go func() {
    sig := <-signals
    if sig == os.Interrupt {
      close(shutdown)
    }
  }()

  return shutdown
}

