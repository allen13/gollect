package outputs

import (
  "github.com/allen13/gollect/data"
  "github.com/allen13/gollect/outputs/file"
  "github.com/allen13/gollect/serializers/graphite"
)

var Outputs = []data.Output{}

func Add(output data.Output) {
	Outputs = append(Outputs, output)
}

func InitOutputs(){
  fileOutput := file.File{}
  fileOutput.Files = []string{"stdout"}
  graphiteSerializer := graphite.GraphiteSerializer{}
  fileOutput.SetSerializer(&graphiteSerializer)
  fileOutput.Connect()
  Add(&fileOutput)
}