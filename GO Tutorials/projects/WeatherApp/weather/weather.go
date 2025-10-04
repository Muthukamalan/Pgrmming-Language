package weather

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}
type Clouds struct {
	All int `json:"all"`
}
type Sys struct {
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type DetailedJSON struct {
	Coord      `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       `json:"main"`
	Visibility int `json:"visibility"`
	Wind       `json:"wind"`
	Clouds     `json:"clouds"`
	Dt         int `json:"dt"`
	Sys        `json:"sys"`
	Timezone   int    `json:"timezone"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Cod        int    `json:"cod"`
}
