package valorant

type (
	SessionResponse struct {
		Name     string `json:"name"`
		PID      string `json:"pid"`
		PlayerId string `json:"puuid"`
		State    string `json:"state"`
	}
)
