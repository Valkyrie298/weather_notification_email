package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Valkyrie298/weather_notification_email/models"
	"github.com/joho/godotenv"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func main() {
	godotenv.Load(".env")
	apiKey := os.Getenv("OPEN_WEATHER_API")
	//url to get weather from Hanoi
	UrlString := "http://dataservice.accuweather.com/forecasts/v1/daily/1day/353412"

	UrlString = fmt.Sprintf(UrlString + "?apikey=" + apiKey)

	fmt.Println(UrlString)

	resp, err := http.Get(UrlString)


	if err != nil {
		fmt.Println(err)
	}

	var weather models.Weather	
	json.NewDecoder(resp.Body).Decode(&weather)

	// fmt.Println(weather.Headline.Severity)

dailyWeather := weather.ConvertWeatherToGeneralizedStructWeather()

var messages []string

for i :=0 ; i<len(dailyWeather); i++ {
	message := dailyWeather[i].MessageFormat()
	messages = append(messages, message)
}

telegramApi := os.Getenv("WEATHER_BOT_TELEGRAM_API")

telegramService, err := telegram.New(telegramApi)
if err!= nil {
	fmt.Println(err)
}


telegramService.AddReceivers(1658858395)

notify.UseServices(telegramService)


for i:=0; i<len(messages); i++ {
	err = notify.Send(
		context.Background(),
		"Weather forecast in the following day",
		messages[i],
	)
	
	if err!= nil {
		fmt.Println(err)
	}
}

// Send a test message.

}