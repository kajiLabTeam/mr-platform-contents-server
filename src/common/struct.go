package common

type AbsoluteAddress struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Rotation struct {
	Roll  float64 `json:"roll"`
	Pitch float64 `json:"pitch"`
	Yaw   float64 `json:"yaw"`
}

type Size struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

type PublicSpace struct {
	LayerId string `json:"id"`
}

// ResponseGetContentsのときにはこれの配列を渡す
type Content struct {
	ContentId   string      `json:"contentId"`
	ContentType string      `json:"contentType"`
	Content     interface{} `json:"content"`
}

type RequestGetContents struct {
	ContentIds []string `json:"contentIds"`
}

type RequestCreateContent struct {
	LayerId     string      `json:"layerId"`
	ContentType string      `json:"contentType"`
	Content     interface{} `json:"content"`
}

type ResponseCreateContent struct {
	ContentId string `json:"contentId"`
}

type RequestUpdateContent struct {
	ContentId   string      `json:"contentId"`
	ContentType string      `json:"contentType"`
	Content     interface{} `json:"content"`
}

type RequestCreateLayer struct {
	LayerId string `json:"layerId"`
}

type ResponseCreateLayer struct {
	LayerId string `json:"layerId"`
}

type ResponseGetLayerContentIds struct {
	ContentIds []string `json:"contentIds"`
}

type Html2d struct {
	Location Location `json:"location"`
	Rotation Rotation `json:"rotation"`
	Size     Size     `json:"size"`
	TextType string   `json:"textType"`
	TextURL  string   `json:"textURL"`
	StyleURL string   `json:"styleURL"`
}

type SQLHtml2d struct {
	ContentId string
	Location  Location
	Rotation  Rotation
	Size      Size
	TextType  string
	TextURL   string
	StyleURL  string
}
