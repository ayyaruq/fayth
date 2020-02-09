package models

type Gauge struct {
	ClassJobID uint8
	Duration2  time.Duration // uint16, time in ms for Gauge2
	Duration   time.Dutation // uint16, time in ms for Gauge
	Gauge      int8          // Primary gauge, negative can indicate Umbral Ice for example
	Gauge1     int8          // Umbral Hearts lives here for example
	Gauge2     uint8         // eg Enochian uses bitfields in this gauge
	unk1       uint32
	unk2       uint32
}
