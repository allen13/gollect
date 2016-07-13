package inputs

import (
  "log"
  "time"
  "os/exec"
  "errors"
  "github.com/allen13/gollect/data"
)

var (
  TimeoutErr = errors.New("Command timed out.")
)

type Exec struct{
  Name string
  Duration string
}

func (exec* Exec)Gather(metricsC chan data.Metric){
  log.Println("gathering " + exec.Name)
  duration,_ := time.ParseDuration(exec.Duration)
  time.Sleep(duration)
  result := make(map[string]interface{})
  result["responseTimeSeconds"] = duration.Seconds()
  metric,_ := data.NewMetric("sl-api",nil, result, time.Now())
  metricsC <- metric
}

func (exec* Exec)Description()(string){
  return exec.Name
}

// RunTimeout runs the given command with the given timeout.
// If the command times out, it attempts to kill the process.
func RunTimeout(c *exec.Cmd, timeout time.Duration) error {
  if err := c.Start(); err != nil {
    return err
  }
  return WaitTimeout(c, timeout)
}

// WaitTimeout waits for the given command to finish with a timeout.
// It assumes the command has already been started.
// If the command times out, it attempts to kill the process.
func WaitTimeout(c *exec.Cmd, timeout time.Duration) error {
  timer := time.NewTimer(timeout)
  done := make(chan error)
  go func() { done <- c.Wait() }()
  select {
  case err := <-done:
    timer.Stop()
    return err
  case <-timer.C:
    if err := c.Process.Kill(); err != nil {
      log.Printf("FATAL error killing process: %s", err)
      return err
    }
  // wait for the command to return after killing it
    <-done
    return TimeoutErr
  }
}
