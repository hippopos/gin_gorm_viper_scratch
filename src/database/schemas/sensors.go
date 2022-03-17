package schemas

type Sensors struct {
	Model `gorm:"embedded"` //引用基本字段
	// Wheels Wheels            `json:"-" gorm:"many2many:wheel_bearing_fk;foreignKey:Code;joinForeignKey:BearingCode;References:Code;joinReferences:WheelCode;" preload:"Wheels"`
	BearingCode string `json:"bearing_code" gorm:"size:100;"`

	SensorTypeCode string     `json:"sensor_type_code" gorm:"size:100;"`
	SensorType     SensorType `json:"sensor_type" gorm:"foreignKey:SensorTypeCode;references:Code" preload:"SensorType"`
	From           string     `json:"from" gorm:"size:100"` //传感器采集  特征值
	BeeClient      string     `json:"bee_client" gorm:"size:100"`
	BeePort        string     `json:"bee_port" gorm:"size:100"`
	BeeAddress     string     `json:"bee_address" gorm:"size:100"`
}

func (p Sensors) TableName() string {
	return "sensors"
}
