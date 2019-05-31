package controller

import (
	"github.com/reef-pi/reef-pi/controller/types"
	"github.com/reef-pi/reef-pi/controller/utils"
	"log"
	"net/http"
	"time"
)

func initializeTelemetry(store types.Store, notify bool) types.Telemetry {
	var t utils.TelemetryConfig
	if err := store.Get(Bucket, "telemetry", &t); err != nil {
		log.Println("ERROR: Failed to load telemtry config from saved settings. Initializing")
		t = utils.DefaultTelemetryConfig
		store.Update(Bucket, "telemetry", t)
	}
	// for upgrades, this value will be 0. Remove in 3.0
	if t.HistoricalLimit < 1 {
		t.HistoricalLimit = types.HistoricalLimit
	}
	if t.CurrentLimit < 1 {
		t.CurrentLimit = types.CurrentLimit
	}
	fn := func(t, m string) error { return logError(store, t, m) }
	return utils.NewTelemetry(t, fn)
}

func (r *ReefPi) getTelemetry(w http.ResponseWriter, req *http.Request) {
	fn := func(_ string) (interface{}, error) {
		var t utils.TelemetryConfig
		if err := r.store.Get(Bucket, "telemetry", &t); err != nil {
			return nil, err
		}
		t.AdafruitIO.Token = ""
		t.Mailer.Password = ""
		return &t, nil
	}
	utils.JSONGetResponse(fn, w, req)
}

func (r *ReefPi) updateTelemetry(w http.ResponseWriter, req *http.Request) {
	var t utils.TelemetryConfig
	fn := func(_ string) error {
		if err := r.store.Update(Bucket, "telemetry", t); err != nil {
			return err
		}
		fn := func(t, m string) error { return logError(r.store, t, m) }
		r.telemetry = utils.NewTelemetry(t, fn)
		return nil
	}
	utils.JSONUpdateResponse(&t, fn, w, req)
}

func (r *ReefPi) sendTestMessage(w http.ResponseWriter, req *http.Request) {
	fn := func(_ string) error {
		_, err := r.telemetry.Alert("[reef-pi] Test email", "This is a test email, generated by reef-pi at:"+time.Now().Format(time.RFC822))
		return err
	}
	utils.JSONDeleteResponse(fn, w, req)
}