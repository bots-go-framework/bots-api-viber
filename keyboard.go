package viberbotapi

type Keyboard struct {
	Buttons []Button

	// Background color of the keyboard.	Valid color HEX value.	By default Viber keyboard background.
	BgColor     string `json:",omitempty"`

	// When true - the keyboard will always be displayed with the same height as the native keyboard.
	// When false - short keyboards will be displayed with the minimal possible height.
	// Maximal height will be native keyboard height.
	DefaultHeight bool `json:",omitempty"` // Default: false
}

type Button struct {
	// Button width in columns. Default is 6. Can be 1-6.
	Columns     int `json:"columns,omitempty"`

	// Button height in rows.
	// Can be 1 or 2.Default is 1.
	// Maximal number of keyboard rows is 12.
	Rows        int `json:"rows,omitempty"`

	// Background color of button. Valid color HEX value. Default: white
	BgColor     string `json:",omitempty"`

	// Type of the background media: "picture" or "gif".
	// For picture - jpeg and png files are supported.
	// Max size: 500 kb, Default to "picture"
	BgMediaType string `json:",omitempty"`

	// URL for background media content (picture or GIF).
	// Will be placed with aspect to fill logic.
	BgMedia     string `json:",omitempty"`

	// When true - animated background media (GIF) will loop continuously.
	// When false - animated background media will play once and stop.
	BgLoop bool `json:",omitempty"`
	ActionType string `json:",omitempty"`
	ActionBody string `json:",omitempty"`
	Image string `json:",omitempty"`
	Text string `json:",omitempty"`
	TextVAlign string `json:",omitempty"` // Vertical alignment of the text: "top", "middle", "bottom". Default "middle"
	TextHAlign string `json:",omitempty"` // Horizontal align of the text: "left", "center", "right". Default "center"
	TextOpacity int `json:",omitempty"`
	TextSize string `json:",omitempty"` // Text size out of 3 options: "small", "regular", "large". Default	"regular"
}

