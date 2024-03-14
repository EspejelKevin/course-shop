package logger

type Measurement struct {
	Object      map[string]interface{}
	TimeElapsed string
}

func NewMeasurement(object map[string]interface{}, timeElapsed string) *Measurement {
	return &Measurement{object, timeElapsed}
}
