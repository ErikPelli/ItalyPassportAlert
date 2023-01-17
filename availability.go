package ItalyPassportAlert

import (
	"github.com/gen2brain/beeep"
	"log"
	"time"
)

func RequestLoop(req Request, enableSound, enableNotify bool, wait int) {
	precedentAvailable := false
	for {
		available := req.CheckAvailability()
		if available && !precedentAvailable {
			// Alert
			if enableSound {
				_ = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			}
			if enableNotify {
				_ = beeep.Notify("Passport", "Passport is available", "assets/information.png")
			}
		}
		log.Println("HTTP request done, result is", available)
		precedentAvailable = available
		time.Sleep(time.Second * time.Duration(wait))
	}
}
