package models

import "fmt"

type Weather struct {
	Headline       *Headline        `json:"Headline"`
	DailyForecasts []DailyForecasts `json:"DailyForecasts"`
}

type Headline struct {
	EffectiveDate string `json:"EffectiveDate"`
	// EffectiveEpochDate int `json:"EffectiveEpochDate"`
	Severity int    `json:"Severity"`
	Text     string `json:"Text"`
	// Category string `json:"Category"`
	// EndDate string `json:"EndDate"`
	// EndEpochDate int `json:"EndEpochDate"`
	// MobileLink string `json:"MobileLink"`
	// Link string `json:"Link"`
}

type DailyForecasts struct {
	Temperature Temperature `json:"Temperature"`
	Day         Day         `json:"Day"`
	Night       Night       `json:"Night"`
}

type Temperature struct {
	MinTemperature `json:"Minimum"`
	MaxTemperature `json:"Maximum"`
}

type MinTemperature struct {
	Value           float64 `json:"Value"`
	TemperatureUnit string  `json:"Unit"`
	UnitType        int     `json:"UnitType"`
}

type MaxTemperature struct {
	Value           float64 `json:"Value"`
	TemperatureUnit string  `json:"Unit"`
	UnitType        int     `json:"UnitType"`
}

type Day struct {
	IconPhrase             string `json:"IconPhrase"`
	HasPrecipitation       bool   `json:"HasPrecipitation"`
	PrecipitationType      string `json:"PrecipitationType"`
	PrecipitationIntensity string `json:"PrecipitationIntensity"`
}

type Night struct {
	IconPhrase             string `json:"IconPhrase"`
	HasPrecipitation       bool   `json:"HasPrecipitation"`
	PrecipitationType      string `json:"PrecipitationType"`
	PrecipitationIntensity string `json:"PrecipitationIntensity"`
}

type DailyWeather struct {
	EffectiveDate               string
	Severity                    int
	Text                        string
	MinTemperature              float64 //in Farenheit
	MaxTemperature              float64 //in Farenheit
	DayPrecipitation            bool
	DayPrecipitationType        string
	DayPrecipitationIntensity   string
	NightPrecipitation          bool
	NightPrecipitationType      string
	NightPrecipitationIntensity string
}

func (w Weather) ConvertWeatherToGeneralizedStructWeather() []DailyWeather {
	var dailyWeathers []DailyWeather
	for i := 0; i < len(w.DailyForecasts); i++ {
		dailyWeather := DailyWeather{
			EffectiveDate:               w.Headline.EffectiveDate,
			Severity:                    w.Headline.Severity,
			Text:                        w.Headline.Text,
			MinTemperature:              (w.DailyForecasts[i].Temperature.MinTemperature.Value - 32) * 5 / 9,
			MaxTemperature:              (w.DailyForecasts[i].Temperature.MaxTemperature.Value - 32) * 5 / 9,
			DayPrecipitation:            w.DailyForecasts[i].Day.HasPrecipitation,
			DayPrecipitationType:        w.DailyForecasts[i].Day.PrecipitationType,
			DayPrecipitationIntensity:   w.DailyForecasts[i].Day.PrecipitationIntensity,
			NightPrecipitation:          w.DailyForecasts[i].Night.HasPrecipitation,
			NightPrecipitationType:      w.DailyForecasts[i].Night.PrecipitationType,
			NightPrecipitationIntensity: w.DailyForecasts[i].Night.PrecipitationIntensity,
		}
		dailyWeathers = append(dailyWeathers, dailyWeather)
	}
	return dailyWeathers

}

func (d DailyWeather) MessageFormat() string {
	message := fmt.Sprintf("Weather forcast for %v, with the minimum temperature of %.2f celcius, and maximum of %.2f celcius",
		d.EffectiveDate, d.MinTemperature, d.MaxTemperature)
	if d.DayPrecipitation {
		message = fmt.Sprintf(message+", "+"during the day, it's going to %v, wht the intensity of %v", d.DayPrecipitationType,
		d.DayPrecipitationIntensity)
	}
	if d.NightPrecipitation {
		message = fmt.Sprintf(message+", "+"during the night, it's going to %v, wht the intensity of %v", d.NightPrecipitationType,
		d.NightPrecipitationIntensity)
	}
	return message
}

