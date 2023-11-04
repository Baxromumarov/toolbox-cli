package info

import "time"

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDat []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type LocationInfo struct {
	Query    string  `json:"query"`
	City     string  `json:"city"`
	Region   string  `json:"regionName"`
	Country  string  `json:"country"`
	ZipCode  string  `json:"zip"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	ISP      string  `json:"isp"`
	Timezone string  `json:"timezone"`
}

type InternetSpeedResult struct {
	DownloadSpeed float64       `json:"download_speed"`
	UploadSpeed   float64       `json:"upload_speed"`
	Latency       time.Duration `json:"latency"`
	Country       string        `json:"country"`
	Name          string        `json:"name"`
	Sponsor       string        `json:"sponsor"`
}
