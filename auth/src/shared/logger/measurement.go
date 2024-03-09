package logger

type Measurement struct {
	object      map[string]interface{}
	timeElapsed string
}

func NewMeasurement(object map[string]interface{}, timeElapsed string) *Measurement {
	return &Measurement{object, timeElapsed}
}
