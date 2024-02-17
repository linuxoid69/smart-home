package main

import (
	"log/slog"
	"os"
	"time"

	cfg "github.com/linuxoid69/smart-home/internal/config"
	"github.com/linuxoid69/smart-home/internal/devices"
	"github.com/spf13/viper"
	"github.com/tuya/tuya-cloud-sdk-go/api/common"
	"github.com/tuya/tuya-cloud-sdk-go/config"
)

func main() {
	cfg.GetConfig()

	config.SetEnv(common.URLEU, viper.GetString("accessId"), viper.GetString("accessKey"))

	switchDevice := false

	for true {

		//switchDevice := scapeExporter()

		for _, dev := range viper.Get("devices").([]interface{}) {
			devMap, ok := dev.(map[string]interface{})
			if !ok {
				continue
			}

			if devMap["name"] == os.Getenv("DEVICE_NAME") {
				devId, ok := devMap["deviceid"].(string)
				if !ok {
					continue
				}

				if err := devices.SetDeviceState(devId, switchDevice); err != nil {
					slog.Error("Work with device:", "error", err, "deviceName", os.Getenv("DEVICE_NAME"))
				}

				slog.Info("Device set status:", "deviceName", os.Getenv("DEVICE_NAME"), "status", switchDevice)
			}
		}

		time.Sleep(time.Second * time.Duration(viper.GetInt("scrape.checkInterval")))
	}
}
