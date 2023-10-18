package wechat

type (
	UserInfo struct {
		HideInputBarFlag  int    `json:"hide_input_bar_flag"`
		StarFriend        int    `json:"star_friend"`
		Sex               int    `json:"sex"`
		AppAccountFlag    int    `json:"app_account_flag"`
		VerifyFlag        int    `json:"verify_flag"`
		ContactFlag       int    `json:"contact_flag"`
		WebWxPluginSwitch int    `json:"web_wx_plugin_switch"`
		HeadImgFlag       int    `json:"head_img_flag"`
		SnsFlag           int    `json:"sns_flag"`
		IsOwner           int    `json:"is_owner"`
		MemberCount       int    `json:"member_count"`
		ChatRoomId        int    `json:"chat_room_id"`
		UniFriend         int    `json:"uni_friend"`
		OwnerUin          int    `json:"owner_uin"`
		Statues           int    `json:"statues"`
		AttrStatus        int64  `json:"attr_status"`
		Uin               int64  `json:"uin"`
		Province          string `json:"province"`
		City              string `json:"city"`
		Alias             string `json:"alias"`
		DisplayName       string `json:"display_name"`
		KeyWord           string `json:"key_word"`
		EncryChatRoomId   string `json:"encry_chat_room_id"`
		UserName          string `json:"user_name"`
		NickName          string `json:"nick_name"`
		HeadImgUrl        string `json:"head_img_url"`
		RemarkName        string `json:"remark_name"`
		PYInitial         string `json:"py_initial"`
		PYQuanPin         string `json:"py_quan_pin"`
		RemarkPYInitial   string `json:"remark_py_initial"`
		RemarkPYQuanPin   string `json:"remark_py_quan_pin"`
		Signature         string `json:"signature"`
	}
)
