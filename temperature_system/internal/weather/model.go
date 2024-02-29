package weather

type WeatherDataPayload struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Latitude       float64 `json:"lat"`
		Longitude      float64 `json:"lon"`
		TimezoneID     string  `json:"tz_id"`
		LocaltimeEpoch int64   `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TemperatureC     float64 `json:"temp_c"`
		TemperatureF     float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindSpeedMph    float64 `json:"wind_mph"`
		WindSpeedKph    float64 `json:"wind_kph"`
		WindDegree      int     `json:"wind_degree"`
		WindDir         string  `json:"wind_dir"`
		PressureMb      float64 `json:"pressure_mb"`
		PressureIn      float64 `json:"pressure_in"`
		PrecipMm        float64 `json:"precip_mm"`
		PrecipIn        float64 `json:"precip_in"`
		Humidity        int     `json:"humidity"`
		Cloud           int     `json:"cloud"`
		FeelslikeC      float64 `json:"feelslike_c"`
		FeelslikeF      float64 `json:"feelslike_f"`
		VisibilityKm    float64 `json:"vis_km"`
		VisibilityMiles float64 `json:"vis_miles"`
		UV              float64 `json:"uv"`
		GustMph         float64 `json:"gust_mph"`
		GustKph         float64 `json:"gust_kph"`
	} `json:"current"`
}

type Weather struct {
	Fahrenheit float64
	Celsius    float64
	Kelvin     float64
}
