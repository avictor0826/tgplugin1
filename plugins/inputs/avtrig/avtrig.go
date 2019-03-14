package avtrig

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type AVtrig struct {
	x	 	float64
	Amplitude	float64
}

var TrigConfig = `
  ## Set the amplitude
  amplitude = 10.0
`

func (s *AVtrig) SampleConfig() string {
	return TrigConfig 
}

func (s *AVtrig) Description() string {
	return "#AV Trig description"
}

func (s *AVtrig) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("avtrig",func() telegraf.Input { return &AVtrig{x: 0.0} })
}
