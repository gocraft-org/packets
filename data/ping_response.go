package data

// PingResponseData is the JSON field within the Status Response packet (0x00) sent to the client from the server
type PingResponseData struct {
	Version     PingResponseVersionField     `json:"version"`
	Players     PingResponsePlayersField     `json:"players"`
	Description PingResponseDescriptionField `json:"description"`
	Favicon     string                       `json:"favicon,omitempty"`
}

// PingResponseVersionField is the version field within the ping response data
type PingResponseVersionField struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

// PingResponsePlayersField is the players field within the ping response data
type PingResponsePlayersField struct {
	Max    int                       `json:"max"`
	Online int                       `json:"online"`
	Sample PingResponsePlayersSample `json:"sample"`
}

// PingResponsePlayersSample is an item in the samples array within the players field of the ping response data
type PingResponsePlayersSample struct {
	Username string `json:"name"`
	UUID     string `json:"id"`
}

// PingResponseDescriptionField is the description field within the ping response data
type PingResponseDescriptionField struct {
	Text          string                         `json:"text"`
	Bold          string                         `json:"bold,omitempty"`
	Italic        string                         `json:"italic,omitempty"`
	Underlined    string                         `json:"underlined,omitempty"`
	Strikethrough string                         `json:"strikethrough,omitempty"`
	Ofbuscated    string                         `json:"obfuscated,omitempty"`
	Color         string                         `json:"color,omitempty"`
	Extra         []PingResponseDescriptionField `json:"extra,omitempty"`
}
