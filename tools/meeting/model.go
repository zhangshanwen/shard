package meeting

type (
	MessageReqBody struct {
		MsgType string    `json:"msg_type"`
		Data    MediaBody `json:"data"`
	}
	MediaBody struct {
		Data   string `json:"data"`
		Offset int    `json:"offset"`
	}
	MessageRespBody struct {
		MsgType string    `json:"msg_type"`
		SendId  int64     `json:"send_id"`
		Data    MediaBody `json:"data"`
	}
)

const (
	Capture = "capture"
	Camera  = "camera"
	Voice   = "voice"
)
