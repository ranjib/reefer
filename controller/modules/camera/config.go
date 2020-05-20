package camera

import (
	"fmt"
	"time"

	"github.com/reef-pi/reef-pi/controller/storage"
)

const DefaulCaptureFlags = ""

type MotionConfig struct {
	Enable bool   `json:"enable"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// swagger:model cameraConfig
type Config struct {
	Enable         bool          `json:"enable"`
	ImageDirectory string        `json:"image_directory"`
	CaptureFlags   string        `json:"capture_flags"`
	TickInterval   time.Duration `json:"tick_interval"`
	Upload         bool          `json:"upload"`
	Motion         MotionConfig  `json:"motion"`
}

var Default = Config{
	ImageDirectory: "/var/lib/reef-pi/images",
	TickInterval:   120,
	Motion: MotionConfig{
		URL:    "http://localhost:8081",
		Width:  800,
		Height: 600,
	},
}

func loadConfig(store storage.Store) (Config, error) {
	var conf Config
	return conf, store.Get(Bucket, "config", &conf)
}

func saveConfig(store storage.Store, conf Config) error {
	if conf.TickInterval <= 0 {
		return fmt.Errorf("Tick Interval for camera controller must be greater than zero")
	}
	if conf.ImageDirectory == "" {
		return fmt.Errorf("Image directory cant not be empty")
	}
	return store.Update(Bucket, "config", conf)
}
