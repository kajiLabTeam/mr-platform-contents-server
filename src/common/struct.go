package common

type RequestPublicSpace struct {
	OrganizationId string `json:"organizationId"`
}

type PublicSpace struct {
	PublicSpaceId  string `json:"id"`
	OrganizationId string `json:"organizationId"`
}

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Rotation struct {
	Row   float64 `json:"row"`
	Pitch float64 `json:"pitch"`
	Yaw   float64 `json:"yaw"`
}

type RequestRelayServer struct {
	ContentIds []string `json:"contentIds"`
}

type RequestContent struct {
	ContentType string `json:"contentType"`
	Domain      string `json:"domain"`
}

type Content struct {
	ContentId   string `json:"contentId"`
	ContentType string `json:"contentType"`
	Domain      string `json:"domain"`
}

type RequestHtml2d struct {
	Location Location `json:"location"`
	Rotation Rotation `json:"rotation"`
	TextType string   `json:"textType"`
	TextURL  string   `json:"textUrl"`
	StyleURL string   `json:"styleUrl"`
}

type ResponseHtml2d struct {
	ContentId string   `json:"contentId"`
	Location  Location `json:"location"`
	Rotation  Rotation `json:"rotation"`
	TextType  string   `json:"textType"`
	TextURL   string   `json:"textUrl"`
	StyleURL  string   `json:"styleUrl"`
}

type Html2d struct {
	Id        string
	ContentId string
	Location  Location
	Rotation  Rotation
	TextType  string
	TextURL   string
	StyleURL  string
}

type RequestModel3d struct {
	Location Location `json:"location"`
	Rotation Rotation `json:"rotation"`
}

type ResponseModel3d struct {
	ContentId    string   `json:"contentId"`
	Location     Location `json:"location"`
	Rotation     Rotation `json:"rotation"`
	PresignedURL string   `json:"presignedUrl"`
}

type Model3d struct {
	Id           string
	ContentId    string
	Location     Location
	Rotation     Rotation
	PresignedURL string
}

type ResponseContents struct {
	ResponseHtml2ds  []ResponseHtml2d  `json:"html2d"`
	ResponseModel3ds []ResponseModel3d `json:"model3d"`
}
