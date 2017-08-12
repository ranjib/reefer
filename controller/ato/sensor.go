package ato

import (
	"github.com/reef-pi/reef-pi/controller/utils"
	"log"
)

func (c *Controller) Read() (int, error) {
	if c.config.DevMode {
		log.Println("ATO is running under dev mode. Sending fixed sensor reading of 1")
		return 1, nil
	}
	return utils.ReadGPIO(c.config.Sensor)
}

func (c *Controller) Control(reading int) error {
	if c.config.DevMode {
		log.Println("ATO is running under dev mode. Skipping controls")
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if reading == 1 {
		if c.pumpOn {
			if err := utils.SwitchOff(c.config.Pump); err != nil {
				return err
			}
			c.pumpOn = false
			log.Println("Switched off ATO pump")
		}
	} else {
		if !c.pumpOn {
			if err := utils.SwitchOn(c.config.Pump); err != nil {
				return err
			}
			c.pumpOn = true
			log.Println("Switched on ATO pump")
		}
	}
	return nil
}
