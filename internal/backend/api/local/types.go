package local

type (
	Presence struct {
		PID      string `json:"pid"`
		Actor    string `json:"actor"`
		GameName string `json:"game_name"`
		GameTag  string `json:"game_tag"`
		State    string `json:"state"`
		Product  string `json:"product"`
		Region   string `json:"region"`
	}
)
