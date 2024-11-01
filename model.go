package ts3_webquery

type ServerInfo struct {
	Body []struct {
		// serverinfo
		ServerName      string `json:"virtualserver_name"`
		ServerUpTime    string `json:"virtualserver_uptime"`
		ServerMessage   string `json:"virtualserver_welcomemessage"`
		ServerMaxClient string `json:"virtualserver_maxclients"`
	} `json:"body"`
	Status Status `json:"status"`
}

type ChannelOptions struct {
	ChannelName                          string `json:"channel_name"`
	ChannelPassword                      string `json:"channel_password"`
	ChannelCodec                         string `json:"channel_codec"`
	ChannelCodecQuality                  string `json:"channel_codec_quality"`
	ChannelFlagPermanent                 string `json:"channel_flag_permanent"`
	ChannelMaxClients                    string `json:"channel_maxclients"`
	ChannelFlagMaxClientsUnlimited       string `json:"channel_flag_maxclients_unlimited"`
	ChannelFlagMaxFamilyClientsInherited string `json:"channel_flag_maxfamilyclients_inherited"`
	ChannelFlagMaxFamilyClientsUnlimited string `json:"channel_flag_maxfamilyclients_unlimited"`
}

type ChannelList struct {
	Body   []ChannelListInfo `json:"body"`
	Status Status            `json:"status"`
}
type ChannelListInfo struct {
	ChannelName      string `json:"channel_name"`
	ChannelID        string `json:"cid"`
	ChannelClient    string `json:"total_clients"`
	ChannelMaxClient string `json:"max_client"`
	ChannelURL       string `json:"url"`
}
type ChannelInfo struct {
	Body []struct {
		ChannelMaxClient string `json:"channel_maxclients"`
	} `json:"body"`
	Status Status `json:"status"`
}
type ClientList struct {
	Body []struct {
		ClientNickname string `json:"client_nickname"`
	} `json:"body"`
	Status Status `json:"status"`
}
type ChannelCreate struct {
	Body []struct {
		CID string `json:"cid"`
	} `json:"body"`
	Status Status `json:"status"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
