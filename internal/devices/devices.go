package devices

import "github.com/tuya/tuya-cloud-sdk-go/api/device"

type DevicesSection struct {
	Devices []Devices `mapstructure:"devices"`
}

type Devices struct {
	Name     string `mapstructure:"name"`
	DeviceID string  `mapstructure:"deviceId"`
}

// SetDeviceState - set the state of a device.
func SetDeviceState(deviceID string, state bool) error {
	_, err := device.PostDeviceCommand(deviceID, []device.Command{
		{
			Code:  "switch_1",
			Value: state,
		},
	})

	return err
}
