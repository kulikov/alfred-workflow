package workflows

func makeItem(title, subtitle string) Item {
	return Item{title, subtitle, title, Icon{"default", "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/Clock.icns"}}
}

type Item struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     Icon   `json:"icon"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}
