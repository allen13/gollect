package inputs

import (
  "log"
  "time"
  "os/exec"
  "errors"
  "github.com/allen13/gollect/data"
  "github.com/allen13/gollect/parsers"
  "fmt"
  "bytes"
  "github.com/gonuts/go-shellquote"
)

var (
  TimeoutErr = errors.New("Command timed out.")
)

type Exec struct {
  command string
  timeout time.Duration
  parser parsers.Parser
}

func NewExec(command string, timeout string, parser string)(error, *Exec){
  timeoutDuration, err := time.ParseDuration(timeout)
  if  err != nil{
    return err, nil
  }

  c := &parsers.Config{}
  c.DataFormat = parser
  c.MetricName = "exec"
  newParser, err := parsers.NewParser(c)
  if err != nil{
    return err, nil
  }

  return nil, &Exec{
    command,
    timeoutDuration,
    newParser,
  }
}

func (e* Exec)Gather(metricsC chan data.Metric) {
  out, err := e.run()
  if err != nil {
    return
  }

  metrics, err := e.parser.Parse(out)
  if err != nil {
    log.Println(err)
  } else {
    for _, metric := range metrics {
      metricsC <- metric
    }
  }
}

//Assembles and runs exec.Cmd object
func (e *Exec) run() ([]byte, error) {
  split_cmd, err := shellquote.Split(e.command)
  if err != nil || len(split_cmd) == 0 {
    return nil, fmt.Errorf("exec: unable to parse command, %s", err)
  }

  cmd := exec.Command(split_cmd[0], split_cmd[1:]...)

  var out bytes.Buffer
  cmd.Stdout = &out

  if err := runTimeout(cmd, e.timeout); err != nil {
    return nil, fmt.Errorf("exec: %s for command '%s'", err, e.command)
  }
  return out.Bytes(), nil
}

// RunTimeout runs the given command with the given timeout.
// If the command times out, it attempts to kill the process.
func runTimeout(c *exec.Cmd, timeout time.Duration) error {
  if err := c.Start(); err != nil {
    return err
  }
  return waitTimeout(c, timeout)
}

// WaitTimeout waits for the given command to finish with a timeout.
// It assumes the command has already been started.
// If the command times out, it attempts to kill the process.
func waitTimeout(c *exec.Cmd, timeout time.Duration) error {
  timer := time.NewTimer(timeout)
  done := make(chan error)
  go func() {
    done <- c.Wait()
  }()
  select {
  case err := <-done:
    timer.Stop()
    return err
  case <-timer.C:
    if err := c.Process.Kill(); err != nil {
      fmt.Errorf("FATAL error killing process: %s", err)
      return err
    }
  // wait for the command to return after killing it
    <-done
    return TimeoutErr
  }
}
