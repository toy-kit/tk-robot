package internal

type Config struct {
	DisplayId int `json:"display_id"`
}

type Result struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type Params struct {
	Tm        int      `json:"tm"`
	X         int      `json:"x"`
	Y         int      `json:"y"`
	W         int      `json:"w"`
	H         int      `json:"h"`
	Name      string   `json:"name"`
	Button    string   `json:"button"`
	Double    bool     `json:"double"`
	Str       string   `json:"str"`
	Speed     int      `json:"speed"`
	Pid       int      `json:"pid"`
	Direction string   `json:"direction"`
	Path      string   `json:"path"`
	Color     string   `json:"color"`
	Keys      []string `json:"keys"`
}
