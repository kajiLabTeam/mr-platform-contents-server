package common

type Location struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Height float64 `json:"height"`
	Scale  string  `json:"scale"`
}

type Rotation struct {
	Roll  float64 `json:"roll"`
	Pitch float64 `json:"pitch"`
	Yaw   float64 `json:"yaw"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ResponseGetContentsのときにはこれの配列を渡す
type Content struct {
	ContentId   string      `json:"contentId"`
	ContentType string      `json:"contentType"`
	Location    Location    `json:"location"`
	Content     interface{} `json:"content"`
}

type RequestGetContents struct {
	ContentIds []string `json:"contentIds"`
}

type RequestCreateContent struct {
	LayerId     string      `json:"layerId"`
	ContentType string      `json:"contentType"`
	Location    Location    `json:"location"`
	Content     interface{} `json:"content"`
}

type ResponseCreateContent struct {
	ContentId   string      `json:"contentId"`
	ContentType string      `json:"contentType"`
	Location    Location    `json:"location"`
	Content     interface{} `json:"content"`
}

type RequestUpdateContent struct {
	ContentId   string      `json:"contentId"`
	ContentType string      `json:"contentType"`
	Location    Location    `json:"location"`
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
	Size     Size   `json:"size"`
	TextType string `json:"textType"`
	TextURL  string `json:"textURL"`
}

type ReturnHtml2d struct {
	Size     Size   `json:"size"`
	TextType string `json:"textType"`
	TextURL  string `json:"textURL"`
	ImgURL   string `json:"imgURL"`
}

type SQLHtml2d struct {
	ContentId string
	Size      Size
	TextType  string
	TextURL   string
}

type Neo4jConfiguration struct {
	URL      string
	Username string
	Password string
}
