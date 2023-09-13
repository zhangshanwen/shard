package wechat

type (
	FriendsGroups struct {
		IsAllFriends   bool     // 是否全部朋友
		ExcludeFriends []string // 排除掉的朋友(优先级高于选中的朋友)
		Friends        []string // 选中的朋友
		IsAllGroups    bool     // 是否全部组群
		ExcludeGroups  []string // 排除掉的群聊(优先级高于选中的群聊)
		Groups         []string // 选中的群聊
	}
	TimerReply struct {
		FriendsGroups
		Msg     string
		Spec    string
		Times   int
		EntryId int
	}
)
