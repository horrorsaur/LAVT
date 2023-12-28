package local

// type PresenceResponse = {
//     presences: {
//         actor?: unknown | null;
//         basic: string;
//         details?: unknown | null;
//         game_name: string;
//         game_tag: string;
//         location?: unknown | null;
//         msg?: unknown | null;
//         name: string;
//         patchline?: unknown | null;
//         pid: string;
//         platform?: unknown | null;
//         private: string | null;
//         privateJwt?: unknown | null;
//         product: "valorant" | "league_of_legends";
//         /** Player UUID */
//         puuid: string;
//         region: string;
//         resource: string;
//         state: "mobile" | "dnd" | "away" | "chat";
//         summary: string;
//         /** Milliseconds since epoch */
//         time: number;
//     }[];
// };

type (
	// friend stuff, just grabbing data i care for
	Presence struct {
		PID      string `json:"pid"`
		Actor    string `json:"actor"`
		GameName string `json:"game_name"`
		GameTag  string `json:"game_tag"`
		State    string `json:"state"`
		Product  string `json:"product"`
	}
)
