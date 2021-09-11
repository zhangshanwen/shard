package response

type (
	BannerResponse struct {
		Banners []Banner `json:"banners"`
	}
	Banner struct {
		Img string `json:"img"`
	}
)
