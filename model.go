package ts3_webquery

type ServerInfos struct {
	Body   []ServerInfo `json:"body"`
	Status Status       `json:"status"`
}
type ServerInfo struct {
	ServerName      string `json:"virtualserver_name"`
	ServerUpTime    string `json:"virtualserver_uptime"`
	ServerMessage   string `json:"virtualserver_welcomemessage"`
	ServerMaxClient string `json:"virtualserver_maxclients"`
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
type ChannelInfos struct {
	Body   []ChannelInfo `json:"body"`
	Status Status        `json:"status"`
}
type ChannelInfo struct {
	ChannelMaxClient string `json:"channel_maxclients"`
}
type ClientList struct {
	Body   []ClientListInfo `json:"body"`
	Status Status           `json:"status"`
}
type ClientListInfo struct {
	ClientNickname string `json:"client_nickname"`
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
