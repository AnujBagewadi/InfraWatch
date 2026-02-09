package service

import (
	"context"
	"net/http"
	"time"

	"github.com/AnujBagewadi/InfraWatch/internal/models"
	repository "github.com/AnujBagewadi/InfraWatch/internal/repo"
)

func StartMonitorEngine() {
	ticker := time.NewTicker(20 * time.Second)

	go func() {
		for {
			<-ticker.C
			MonitorAllAPIs()
		}
	}()
}

func MonitorAllAPIs() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	apis, err := repository.GetAllAPIs(ctx)
	if err != nil {
		return
	}

	for _, api := range apis {
		go MonitorSingleAPI(api)
	}
}

func MonitorSingleAPI(api models.API) {
	start := time.Now()
	resp, err := http.Get(api.URL)
	elapsed := time.Since(start).Milliseconds()

	log := models.APILog{
		APIID:          api.ID,
		Timestamp:      time.Now(),
		ResponseTimeMS: elapsed,
	}

	if err != nil {
		log.StatusCode = 0
		log.Success = false

		// Insert log first
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		repository.InsertLog(ctx, log)
		cancel()

		// Check last 3 failures
		if ShouldTriggerFailureAlert(api.ID) {
			CreateAlert(api.ID, "API failed 3 consecutive times")
		}
		return
	} else {
		log.StatusCode = resp.StatusCode
		log.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
		resp.Body.Close()

		// Alert if latency high
		if elapsed > int64(api.ThresholdMS) {
			CreateAlert(api.ID, "High latency detected")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repository.InsertLog(ctx, log)
}
