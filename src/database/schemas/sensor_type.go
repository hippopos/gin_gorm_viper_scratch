package schemas

type SensorType struct {
	Model     `gorm:"embedded"` //引用基本字段
	Unit      string            `json:"unit" gorm:"size:100;"`
	Frequency string            `json:"frequency" gorm:"size:100;"`
	// From      string            `json:"from" gorm:"size:100;"`
}

func (p SensorType) TableName() string {
	return "sensor_type"
}

var InitSensorTypeData = []SensorType{
	{Model{Code: "temper", Name: "温度传感器"}, "摄氏度", ""},
	{Model{Code: "vib", Name: "振动传感器，速度/加速度/位移"}, "mm/s, mm/s2, mm", ">3kHz"},
	{Model{Code: "speed", Name: "转速传感器"}, "mm/s, mm/s2, mm", ">3kHz"},
	{Model{Code: "noise", Name: "噪音传感器"}, "dB", ""},
	{Model{Code: "health", Name: "健康度"}, "100%", ""},
}
