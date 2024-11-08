package ts3_webquery

type Response struct {
	Body   []Body `json:"body,omitempty"`
	Status Result `json:"status"`
}

// login serveradmin Ox3uLCud

type Body struct {
	// apikeyadd
	APIKey    string `json:"apikey,omitempty"`
	ClDBID    string `json:"cldbid,omitempty"`
	CreatedAt string `json:"createdat,omitempty"`
	ExpiresAt string `json:"expiresat,omitempty"`
	ID        string `json:"id,omitempty"`
	Scope     string `json:"scope,omitempty"`
	SID       string `json:"sid,omitempty"`
	TimeLeft  string `json:"timeleft,omitempty"`
	// 	apikeydel
	// 	apikeylist
	Count string `json:"count,omitempty"`
	/*
		CLDBID    string `json:"cldbid,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		ExpiresAt string `json:"expires_at,omitempty"`
		ID        string `json:"id,omitempty"`
		Scope     string `json:"scope,omitempty"`
		SID       string `json:"sid,omitempty"`
		TimeLeft  string `json:"time_left,omitempty"`
	*/

	//	banadd
	BanID string `json:"banid,omitempty"`

	//	banclient
	//	bandel
	// 	bandelall

	// 	banlist
	Created       string `json:"created,omitempty"`
	Duration      string `json:"duration,omitempty"`
	Enforcements  string `json:"enforcements,omitempty"`
	InvokerCLDBID string `json:"invokercldbid,omitempty"`
	InvokerName   string `json:"invokername,omitempty"`
	InvokerUID    string `json:"invokeruid,omitempty"`
	IP            string `json:"ip,omitempty"`
	LastNickname  string `json:"lastnickname,omitempty"`
	MyTSID        string `json:"mytsid,omitempty"`
	Name          string `json:"name,omitempty"`
	Reason        string `json:"reason,omitempty"`
	UID           string `json:"uid,omitempty"`
	/*
		BanID	string 	`json:"banid,omitempty"`
		Count	int    	`json:"count,omitempty"`
	*/

	//	bindinglist
	//	channeladdperm
	//	channelclientaddperm
	//	channelclientdelperm

	//	channelclientpermlist
	CID         string `json:"cid,omitempty"`
	PermID      string `json:"permid,omitempty"`
	PermValue   string `json:"permvalue,omitempty"`
	PermNegated string `json:"permnegated,omitempty"`
	/*
		ClDBID string `json:"cldbid,omitempty"`
	*/

	//	channelcreate

	//	channeldelete
	//	channeldelperm
	//	channeledit

	//	channelfind
	ChannelName string `json:"channel_name,omitempty"`
	/*
		CID string `json:"cid,omitempty"`
	*/

	//	channelgroupadd
	CGID string `json:"cgid,omitempty"`

	//	channelgroupaddperm

	//	channelgroupclientlist
	/*
		CID    string `json:"cid,omitempty"`
		ClDBID string `json:"cldbid,omitempty"`
		CGID   string `json:"cgid,omitempty"`
	*/

	//	channelgroupcopy
	/*
		CGID string `json:"cgid,omitempty"`
	*/

	//	channelgroupdel
	//	channelgroupdelperm

	//	channelgrouplist
	Type   string `json:"type,omitempty"`
	IconID string `json:"iconid,omitempty"`
	SaveDB string `json:"savedb,omitempty"`
	/*
		CGID string `json:"cgid,omitempty"`
		Name string `json:"name,omitempty"`
	*/

	//	channelgrouppermlist
	PremSkip string `json:"permskip,omitempty"`
	/*
		PermID      string `json:"permid,omitempty"`
		PermValue   string `json:"permvalue,omitempty"`
		PermNegated string `json:"permnegated,omitempty"`
	*/

	//	channelgrouprename

	//	channelinfo
	ChannelBannerGfxURL                  string `json:"channel_banner_gfx_url,omitempty"`
	ChannelBannerMode                    string `json:"channel_banner_mode,omitempty"`
	ChannelCodec                         string `json:"channel_codec,omitempty"`
	ChannelCodecIsUnencrypted            string `json:"channel_codec_is_unencrypted,omitempty"`
	ChannelCodecLatencyFactor            string `json:"channel_codec_latency_factor,omitempty"`
	ChannelCodecQuality                  string `json:"channel_codec_quality,omitempty"`
	ChannelDeleteDelay                   string `json:"channel_delete_delay,omitempty"`
	ChannelDescription                   string `json:"channel_description,omitempty"`
	ChannelFilepath                      string `json:"channel_filepath,omitempty"`
	ChannelFlagDefault                   string `json:"channel_flag_default,omitempty"`
	ChannelFlagMaxclientsUnlimited       string `json:"channel_flag_maxclients_unlimited,omitempty"`
	ChannelFlagMaxfamilyclientsInherited string `json:"channel_flag_maxfamilyclients_inherited,omitempty"`
	ChannelFlagMaxfamilyclientsUnlimited string `json:"channel_flag_maxfamilyclients_unlimited,omitempty"`
	ChannelFlagPassword                  string `json:"channel_flag_password,omitempty"`
	ChannelFlagPermanent                 string `json:"channel_flag_permanent,omitempty"`
	ChannelFlagSemiPermanent             string `json:"channel_flag_semi_permanent,omitempty"`
	ChannelForcedSilence                 string `json:"channel_forced_silence,omitempty"`
	ChannelIconID                        string `json:"channel_icon_id,omitempty"`
	ChannelMaxclients                    string `json:"channel_maxclients,omitempty"`
	ChannelMaxfamilyclients              string `json:"channel_maxfamilyclients,omitempty"`
	ChannelNamePhonetic                  string `json:"channel_name_phonetic,omitempty"`
	ChannelNeededTalkPower               string `json:"channel_needed_talk_power,omitempty"`
	ChannelOrder                         string `json:"channel_order,omitempty"`
	ChannelPassword                      string `json:"channel_password,omitempty"`
	ChannelSecuritySalt                  string `json:"channel_security_salt,omitempty"`
	ChannelTopic                         string `json:"channel_topic,omitempty"`
	ChannelUniqueIdentifier              string `json:"channel_unique_identifier,omitempty"`
	PID                                  string `json:"pid,omitempty"`
	SecondsEmpty                         string `json:"seconds_empty,omitempty"`
	/*
		ChannelName string `json:"channel_name,omitempty"`
	*/

	//	channellist
	ChannelNeededSubscribePower string `json:"channel_needed_subscribe_power,omitempty"`
	TotalClients                string `json:"total_clients,omitempty"`
	TotalClientsFamily          string `json:"total_clients_family,omitempty"`
	/*
		ChannelBannerGfxURL      string `json:"channel_banner_gfx_url,omitempty"`
		ChannelBannerMode        string `json:"channel_banner_mode,omitempty"`
		ChannelCodec             string `json:"channel_codec,omitempty"`
		ChannelCodecQuality      string `json:"channel_codec_quality,omitempty"`
		ChannelFlagDefault       string `json:"channel_flag_default,omitempty"`
		ChannelFlagPassword      string `json:"channel_flag_password,omitempty"`
		ChannelFlagPermanent     string `json:"channel_flag_permanent,omitempty"`
		ChannelFlagSemiPermanent string `json:"channel_flag_semi_permanent,omitempty"`
		ChannelIconID            string `json:"channel_icon_id,omitempty"`
		ChannelMaxclients        string `json:"channel_maxclients,omitempty"`
		ChannelMaxfamilyclients  string `json:"channel_maxfamilyclients,omitempty"`
		ChannelName              string `json:"channel_name,omitempty"`
		ChannelNeededTalkPower   string `json:"channel_needed_talk_power,omitempty"`
		ChannelOrder             string `json:"channel_order,omitempty"`
		ChannelTopic             string `json:"channel_topic,omitempty"`
		CID                      string `json:"cid,omitempty"`
		PID                      string `json:"pid,omitempty"`
		SecondsEmpty             string `json:"seconds_empty,omitempty"`
	*/

	//	channelmove

	//	channelpermlist
	/*
		CID         string `json:"cid,omitempty"`
		PermID      string `json:"permid,omitempty"`
		PermValue   string `json:"permvalue,omitempty"`
		PermNegated string `json:"permnegated,omitempty"`
		PremSkip    string `json:"permskip,omitempty"`
	*/

	//	clientaddperm

	//	clientdbdelete

	//	clientdbedit

	//	clientdbfind
	/*
		CLDBID string `json:"cldbid,omitempty"`
	*/

	//	clientdbinfo
	ClientBase64HashClientUID  string `json:"client_base64HashClientUID,omitempty"`
	ClientCreated              string `json:"client_created,omitempty"`
	ClientDatabaseID           string `json:"client_database_id,omitempty"`
	ClientDescription          string `json:"client_description,omitempty"`
	ClientFlagAvatar           string `json:"client_flag_avatar,omitempty"`
	ClientLastConnected        string `json:"client_lastconnected,omitempty"`
	ClientLastIP               string `json:"client_lastip,omitempty"`
	ClientMonthBytesDownloaded string `json:"client_month_bytes_downloaded,omitempty"`
	ClientMonthBytesUploaded   string `json:"client_month_bytes_uploaded,omitempty"`
	ClientNickname             string `json:"client_nickname,omitempty"`
	ClientTotalBytesDownloaded string `json:"client_total_bytes_downloaded,omitempty"`
	ClientTotalBytesUploaded   string `json:"client_total_bytes_uploaded,omitempty"`
	ClientTotalConnections     string `json:"client_totalconnections,omitempty"`
	ClientUniqueIdentifier     string `json:"client_unique_identifier,omitempty"`

	//	clientdblist
	ClientLoginName string `json:"client_login_name,omitempty"`
	/*
		ClDBID                 string `json:"cldbid,omitempty"`
		ClientCreated          string `json:"client_created,omitempty"`
		ClientDescription      string `json:"client_description,omitempty"`
		ClientLastConnected    string `json:"client_lastconnected,omitempty"`
		ClientLastIP           string `json:"client_lastip,omitempty"`
		ClientNickname         string `json:"client_nickname,omitempty"`
		ClientTotalConnections string `json:"client_totalconnections,omitempty"`
		ClientUniqueIdentifier string `json:"client_unique_identifier,omitempty"`
	*/

	//	clientdelperm

	//	clientedit

	//	clientfind
	Clid string `json:"clid,omitempty"`
	/*
		ClientNickname string `json:"client_nickname,omitempty"`
	*/

	//	clientgetdbidfromuid
	ClUID string `json:"cluid,omitempty"`
	/*
		ClDBID string `json:"cldbid,omitempty"`
	*/

	//	clientgetids
	/*
		ClUID string `json:"cluid,omitempty"`
		Clid  string `json:"clid,omitempty"`
		Name  string `json:"name,omitempty"`
	*/

	//	clientgetnamefromdbid
	/*
		ClDBID  string `json:"cldbid,omitempty"`
		ClDBID string `json:"cldbid,omitempty"`
		Name   string `json:"name,omitempty"`
	*/

	//	clientgetnamefromuid
	/*
		ClUID  string `json:"cluid,omitempty"`
		ClDBID string `json:"cldbid,omitempty"`
		Name   string `json:"name,omitempty"`
	*/

	//	clientgetuidfromclid
	/*
		ClID  string `json:"clid,omitempty"`
		ClDBID string `json:"cldbid,omitempty"`
		Name   string `json:"name,omitempty"`
	*/

	//	clientinfo
	ClientAway                                 string `json:"client_away,omitempty"`
	ClientAwayMessage                          string `json:"client_away_message,omitempty"`
	ClientBadges                               string `json:"client_badges,omitempty"`
	ClientChannelGroupID                       string `json:"client_channel_group_id,omitempty"`
	ClientChannelGroupInheritedChannelID       string `json:"client_channel_group_inherited_channel_id,omitempty"`
	ClientCountry                              string `json:"client_country,omitempty"`
	ClientDefaultChannel                       string `json:"client_default_channel,omitempty"`
	ClientDefaultToken                         string `json:"client_default_token,omitempty"`
	ClientIconID                               string `json:"client_icon_id,omitempty"`
	ClientIdleTime                             string `json:"client_idle_time,omitempty"`
	ClientInputHardware                        string `json:"client_input_hardware,omitempty"`
	ClientInputMuted                           string `json:"client_input_muted,omitempty"`
	ClientIntegrations                         string `json:"client_integrations,omitempty"`
	ClientIsChannelCommander                   string `json:"client_is_channel_commander,omitempty"`
	ClientIsPrioritySpeaker                    string `json:"client_is_priority_speaker,omitempty"`
	ClientIsRecording                          string `json:"client_is_recording,omitempty"`
	ClientIsTalker                             string `json:"client_is_talker,omitempty"`
	ClientMetaData                             string `json:"client_meta_data,omitempty"`
	ClientMyteamspeakAvatar                    string `json:"client_myteamspeak_avatar,omitempty"`
	ClientMyteamspeakID                        string `json:"client_myteamspeak_id,omitempty"`
	ClientNeededServerqueryViewPower           string `json:"client_needed_serverquery_view_power,omitempty"`
	ClientNicknamePhonetic                     string `json:"client_nickname_phonetic,omitempty"`
	ClientOutputHardware                       string `json:"client_output_hardware,omitempty"`
	ClientOutputMuted                          string `json:"client_output_muted,omitempty"`
	ClientOutputonlyMuted                      string `json:"client_outputonly_muted,omitempty"`
	ClientPlatform                             string `json:"client_platform,omitempty"`
	ClientSecurityHash                         string `json:"client_security_hash,omitempty"`
	ClientServergroups                         string `json:"client_servergroups,omitempty"`
	ClientSignedBadges                         string `json:"client_signed_badges,omitempty"`
	ClientTalkPower                            string `json:"client_talk_power,omitempty"`
	ClientTalkRequest                          string `json:"client_talk_request,omitempty"`
	ClientTalkRequestMsg                       string `json:"client_talk_request_msg,omitempty"`
	ClientType                                 string `json:"client_type,omitempty"`
	ClientUnreadMessages                       string `json:"client_unread_messages,omitempty"`
	ClientVersion                              string `json:"client_version,omitempty"`
	ClientVersionSign                          string `json:"client_version_sign,omitempty"`
	ConnectionBandwidthReceivedLastMinuteTotal string `json:"connection_bandwidth_received_last_minute_total,omitempty"`
	ConnectionBandwidthReceivedLastSecondTotal string `json:"connection_bandwidth_received_last_second_total,omitempty"`
	ConnectionBandwidthSentLastMinuteTotal     string `json:"connection_bandwidth_sent_last_minute_total,omitempty"`
	ConnectionBandwidthSentLastSecondTotal     string `json:"connection_bandwidth_sent_last_second_total,omitempty"`
	ConnectionBytesReceivedTotal               string `json:"connection_bytes_received_total,omitempty"`
	ConnectionBytesSentTotal                   string `json:"connection_bytes_sent_total,omitempty"`
	ConnectionClientIP                         string `json:"connection_client_ip,omitempty"`
	ConnectionConnectedTime                    string `json:"connection_connected_time,omitempty"`
	ConnectionFiletransferBandwidthReceived    string `json:"connection_filetransfer_bandwidth_received,omitempty"`
	ConnectionFiletransferBandwidthSent        string `json:"connection_filetransfer_bandwidth_sent,omitempty"`
	ConnectionPacketsReceivedTotal             string `json:"connection_packets_received_total,omitempty"`
	ConnectionPacketsSentTotal                 string `json:"connection_packets_sent_total,omitempty"`
	/*
		CID                        string `json:"cid,omitempty"`
		ClientBase64HashClientUID  string `json:"client_base64HashClientUID,omitempty"`
		ClientCreated              string `json:"client_created,omitempty"`
		ClientDatabaseID           string `json:"client_database_id,omitempty"`
		ClientDescription          string `json:"client_description,omitempty"`
		ClientFlagAvatar           string `json:"client_flag_avatar,omitempty"`
		ClientLastConnected        string `json:"client_lastconnected,omitempty"`
		ClientLoginName            string `json:"client_login_name,omitempty"`
		ClientMonthBytesDownloaded string `json:"client_month_bytes_downloaded,omitempty"`
		ClientMonthBytesUploaded   string `json:"client_month_bytes_uploaded,omitempty"`
		ClientNickname             string `json:"client_nickname,omitempty"`
		ClientTotalBytesDownloaded string `json:"client_total_bytes_downloaded,omitempty"`
		ClientTotalBytesUploaded   string `json:"client_total_bytes_uploaded,omitempty"`
		ClientUniqueIdentifier     string `json:"client_unique_identifier,omitempty"`
		ClientTotalconnections     string `json:"client_totalconnections,omitempty"`
	*/

	//	clientkick

	//	clientlist
	/*
		CID                                  string `json:"cid,omitempty"`
		ClID                                 string `json:"clid,omitempty"`
		ClientAway                           string `json:"client_away,omitempty"`
		ClientAwayMessage                    string `json:"client_away_message,omitempty"`
		ClientBadges                         string `json:"client_badges,omitempty"`
		ClientChannelGroupID                 string `json:"client_channel_group_id,omitempty"`
		ClientChannelGroupInheritedChannelID string `json:"client_channel_group_inherited_channel_id,omitempty"`
		ClientCountry                        string `json:"client_country,omitempty"`
		ClientCreated                        string `json:"client_created,omitempty"`
		ClientDatabaseID                     string `json:"client_database_id,omitempty"`
		ClientFlagTalking                    string `json:"client_flag_talking,omitempty"`
		ClientIdleTime                       string `json:"client_idle_time,omitempty"`
		ClientInputHardware                  string `json:"client_input_hardware,omitempty"`
		ClientInputMuted                     string `json:"client_input_muted,omitempty"`
		ClientIsChannelCommander             string `json:"client_is_channel_commander,omitempty"`
		ClientIsPrioritySpeaker              string `json:"client_is_priority_speaker,omitempty"`
		ClientIsRecording                    string `json:"client_is_recording,omitempty"`
		ClientIsTalker                       string `json:"client_is_talker,omitempty"`
		ClientLastConnected                  string `json:"client_lastconnected,omitempty"`
		ClientNickname                       string `json:"client_nickname,omitempty"`
		ClientOutputHardware                 string `json:"client_output_hardware,omitempty"`
		ClientOutputMuted                    string `json:"client_output_muted,omitempty"`
		ClientPlatform                       string `json:"client_platform,omitempty"`
		ClientServerGroups                   string `json:"client_servergroups,omitempty"`
		ClientTalkPower                      string `json:"client_talk_power,omitempty"`
		ClientType                           string `json:"client_type,omitempty"`
		ClientUniqueIdentifier               string `json:"client_unique_identifier,omitempty"`
		ClientVersion                        string `json:"client_version,omitempty"`
		ConnectionClientIP                   string `json:"connection_client_ip,omitempty"`
	*/

	//	clientmove

	//	clientpermlist
	/*
		ClDBID      string `json:"cldbid,omitempty"`
		CID         string `json:"cid,omitempty"`
		PermID      string `json:"permid,omitempty"`
		PermValue   string `json:"permvalue,omitempty"`
		PermNegated string `json:"permnegated,omitempty"`
		PremSkip    string `json:"permskip,omitempty"`
	*/

	//	clientpoke

	//	clientsetserverquerylogin
	ClientLoginPassword string `json:"client_login_password,omitempty"`

	//	clientupdate

	//	complainadd

	//	complaindel

	//	complaindelall

	//	complainlist
	FromClDBID string `json:"fcldbid,omitempty"`
	FromName   string `json:"fname,omitempty"`
	Message    string `json:"message,omitempty"`
	ToClDBID   string `json:"tcldbid,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	ToName     string `json:"tname,omitempty"`
	//	custominfo
	Ident string `json:"ident,omitempty"`
	Value string `json:"value,omitempty"`
	/*
		ClDBID string `json:"cldbid,omitempty"`
	*/
	//	customsearch

	//	customset

	//	customdelete

	//	ftcreatedir

	//	ftdeletefile

	//	ftgetfileinfo
	Size string `json:"size,omitempty"`
	/*
		CID      string `json:"cid,omitempty"`
		Name     string `json:"name,omitempty"`
		DataTime string `json:"datetime,omitempty"`
		Type     string `json:"type,omitempty"`
	*/

	//	gm

	//	hostinfo

	ConnectionFiletransferBytesReceivedTotal string `json:"connection_filetransfer_bytes_received_total,omitempty"`
	ConnectionFiletransferBytesSentTotal     string `json:"connection_filetransfer_bytes_sent_total,omitempty"`
	HostTimestampUTC                         string `json:"host_timestamp_utc,omitempty"`
	InstanceUptime                           string `json:"instance_uptime,omitempty"`
	VirtualserversRunningTotal               string `json:"virtualservers_running_total,omitempty"`
	VirtualserversTotalChannelsOnline        string `json:"virtualservers_total_channels_online,omitempty"`
	VirtualserversTotalClientsOnline         string `json:"virtualservers_total_clients_online,omitempty"`
	VirtualserversTotalMaxclients            string `json:"virtualservers_total_maxclients,omitempty"`
	/*
		ConnectionBandwidthReceivedLastMinuteTotal string `json:"connection_bandwidth_received_last_minute_total,omitempty"`
		ConnectionBandwidthReceivedLastSecondTotal string `json:"connection_bandwidth_received_last_second_total,omitempty"`
		ConnectionBandwidthSentLastMinuteTotal     string `json:"connection_bandwidth_sent_last_minute_total,omitempty"`
		ConnectionBandwidthSentLastSecondTotal     string `json:"connection_bandwidth_sent_last_second_total,omitempty"`
		ConnectionBytesReceivedTotal               string `json:"connection_bytes_received_total,omitempty"`
		ConnectionBytesSentTotal                   string `json:"connection_bytes_sent_total,omitempty"`
		ConnectionFiletransferBandwidthReceived    string `json:"connection_filetransfer_bandwidth_received,omitempty"`
		ConnectionFiletransferBandwidthSent        string `json:"connection_filetransfer_bandwidth_sent,omitempty"`
		ConnectionPacketsReceivedTotal             string `json:"connection_packets_received_total,omitempty"`
		ConnectionPacketsSentTotal                 string `json:"connection_packets_sent_total,omitempty"`
	*/

	//	instanceedit

	//	instanceinfo
	ServerInstanceDatabaseVersion                string `json:"serverinstance_database_version,omitempty"`
	ServerInstanceFiletransferPort               string `json:"serverinstance_filetransfer_port,omitempty"`
	ServerInstanceGuestServerqueryGroup          string `json:"serverinstance_guest_serverquery_group,omitempty"`
	ServerInstanceMaxDownloadTotalBandwidth      string `json:"serverinstance_max_download_total_bandwidth,omitempty"`
	ServerInstanceMaxUploadTotalBandwidth        string `json:"serverinstance_max_upload_total_bandwidth,omitempty"`
	ServerInstancePendingConnectionsPerIP        string `json:"serverinstance_pending_connections_per_ip,omitempty"`
	ServerInstancePermissionsVersion             string `json:"serverinstance_permissions_version,omitempty"`
	ServerInstanceServerqueryBanTime             string `json:"serverinstance_serverquery_ban_time,omitempty"`
	ServerInstanceServerqueryFloodCommands       string `json:"serverinstance_serverquery_flood_commands,omitempty"`
	ServerInstanceServerqueryFloodTime           string `json:"serverinstance_serverquery_flood_time,omitempty"`
	ServerInstanceServerqueryMaxConnectionsPerIP string `json:"serverinstance_serverquery_max_connections_per_ip,omitempty"`
	ServerInstanceTemplateChanneladminGroup      string `json:"serverinstance_template_channeladmin_group,omitempty"`
	ServerInstanceTemplateChanneldefaultGroup    string `json:"serverinstance_template_channeldefault_group,omitempty"`
	ServerInstanceTemplateServeradminGroup       string `json:"serverinstance_template_serveradmin_group,omitempty"`
	ServerInstanceTemplateServerdefaultGroup     string `json:"serverinstance_template_serverdefault_group,omitempty"`

	//	logview
	FileSize string `json:"file_size,omitempty"`
	LogEntry string `json:"l,omitempty"`
	LastPos  string `json:"last_pos,omitempty"`
	//	messageadd

	//	messagedel

	//	messageget
	MsgID   string `json:"msgid,omitempty"`
	Subject string `json:"subject,omitempty"`
	/*
		ClUID   string `json:"cluid,omitempty"`
		Message string `json:"message,omitempty"`
	*/

	//	messagelist

	//	messageupdateflag

	//	permfind

	//	permget
	PermsID string `json:"permsid,omitempty"`
	/*
		PermID    string `json:"permid,omitempty"`
		PermValue string `json:"permvalue,omitempty"`
	*/

	//	permidgetbyname
	/*
		PermsID string `json:"permsid,omitempty"`
		PermID  string `json:"permid,omitempty"`
	*/

	//	permissionlist
	PermName string `json:"permname,omitempty"`
	PermDesc string `json:"permdesc,omitempty"`
	/*
		PermID string `json:"permid,omitempty"`
	*/
	//	permoverview

	//	permreset
	Token string `json:"token,omitempty"`

	//	privilegekeyadd
	/*
		Token string `json:"token,omitempty"`
	*/
	//	privilegekeydelete

	//	privilegekeylist
	TokenCreated     string `json:"token_created,omitempty"`
	TokenCustomSet   string `json:"token_customset,omitempty"`
	TokenDescription string `json:"token_description,omitempty"`
	TokenID1         string `json:"token_id1,omitempty"`
	TokenID2         string `json:"token_id2,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	/*
		Token string `json:"token,omitempty"`
	*/

	//	privilegekeyuse

	//	queryloginadd
	/*
		ClDBID              string `json:"cldbid,omitempty"`
		ClientLoginName     string `json:"client_login_name,omitempty"`
		ClientLoginPassword string `json:"client_login_password,omitempty"`
	*/

	//	querylogindel

	//	queryloginlist
	/*
		ClDBID          string `json:"cldbid,omitempty"`
		SID             string `json:"sid,omitempty"`
		ClientLoginName string `json:"client_login_name,omitempty"`
	*/

	//	sendtextmessage

	//	servercreate
	VirtualServerPort string `json:"virtualserver_port,omitempty"`
	/*
		SID   string `json:"sid,omitempty"`
		Token string `json:"token,omitempty"`
	*/

	//	serverdelete

	//	serveredit

	//	servergroupadd
	SGID string `json:"sgid,omitempty"`

	//	servergroupaddclient

	//	servergroupaddperm

	//	servergroupautoaddperm

	//	servergroupautodelperm

	//	servergroupclientlist
	/*
		ClDBID string `json:"cldbid,omitempty"`
	*/

	//	servergroupcopy
	/*
		SGID string `json:"sgid,omitempty"`
	*/

	//	servergroupdel

	//	servergroupdelclient

	//	servergroupdelperm

	//	servergrouplist
	NMemberAddP    string `json:"n_member_addp,omitempty"`
	NMemberRemoveP string `json:"n_member_removep,omitempty"`
	NModifyP       string `json:"n_modifyp,omitempty"`
	NameMode       string `json:"namemode,omitempty"`
	SortID         string `json:"sortid,omitempty"`
	/*
		SGID   string `json:"sgid,omitempty"`
		IconID string `json:"iconid,omitempty"`
		SaveDB string `json:"savedb,omitempty"`
		Name   string `json:"name,omitempty"`
		Type   string `json:"type,omitempty"`
	*/

	//	servergrouppermlist
	/*
		CID         string `json:"cid,omitempty"`
		PermID      string `json:"permid,omitempty"`
		PermValue   string `json:"permvalue,omitempty"`
		PermNegated string `json:"permnegated,omitempty"`
		PremSkip    string `json:"permskip,omitempty"`
	*/

	//	servergrouprename

	//	servergroupsbyclientid
	/*
		SGID string `json:"sgid,omitempty"`
		Name   string `json:"name,omitempty"`
		ClDBID string `json:"cldbid,omitempty"`
	*/

	//	serveridgetbyport
	ServerID string `json:"server_id,omitempty"`

	//	serverinfo

	ConnectionBytesReceivedControl                 string `json:"connection_bytes_received_control,omitempty"`
	ConnectionBytesReceivedSpeech                  string `json:"connection_bytes_received_speech,omitempty"`
	ConnectionBytesSentControl                     string `json:"connection_bytes_sent_control,omitempty"`
	ConnectionBytesReceivedKeepalive               string `json:"connection_bytes_received_keepalive,omitempty"`
	ConnectionBytesSentKeepalive                   string `json:"connection_bytes_sent_keepalive,omitempty"`
	ConnectionBytesSentSpeech                      string `json:"connection_bytes_sent_speech,omitempty"`
	ConnectionPacketsReceivedControl               string `json:"connection_packets_received_control,omitempty"`
	ConnectionPacketsReceivedKeepalive             string `json:"connection_packets_received_keepalive,omitempty"`
	ConnectionPacketsReceivedSpeech                string `json:"connection_packets_received_speech,omitempty"`
	ConnectionPacketsSentControl                   string `json:"connection_packets_sent_control,omitempty"`
	ConnectionPacketsSentKeepalive                 string `json:"connection_packets_sent_keepalive,omitempty"`
	ConnectionPacketsSentSpeech                    string `json:"connection_packets_sent_speech,omitempty"`
	VirtualServerAntifloodPointsNeededCommandBlock string `json:"virtualserver_antiflood_points_needed_command_block,omitempty"`
	VirtualServerAntifloodPointsNeededIPBlock      string `json:"virtualserver_antiflood_points_needed_ip_block,omitempty"`
	VirtualServerAntifloodPointsNeededPluginBlock  string `json:"virtualserver_antiflood_points_needed_plugin_block,omitempty"`
	VirtualServerAntifloodPointsTickReduce         string `json:"virtualserver_antiflood_points_tick_reduce,omitempty"`
	VirtualServerAskForPrivilegekey                string `json:"virtualserver_ask_for_privilegekey,omitempty"`
	VirtualServerAutostart                         string `json:"virtualserver_autostart,omitempty"`
	VirtualServerCapabilityExtensions              string `json:"virtualserver_capability_extensions,omitempty"`
	VirtualServerChannelTempDeleteDelayDefault     string `json:"virtualserver_channel_temp_delete_delay_default,omitempty"`
	VirtualServerChannelsonline                    string `json:"virtualserver_channelsonline,omitempty"`
	VirtualServerClientConnections                 string `json:"virtualserver_client_connections,omitempty"`
	VirtualServerClientsonline                     string `json:"virtualserver_clientsonline,omitempty"`
	VirtualServerCodecEncryptionMode               string `json:"virtualserver_codec_encryption_mode,omitempty"`
	VirtualServerComplainAutobanCount              string `json:"virtualserver_complain_autoban_count,omitempty"`
	VirtualServerComplainAutobanTime               string `json:"virtualserver_complain_autoban_time,omitempty"`
	VirtualServerComplainRemoveTime                string `json:"virtualserver_complain_remove_time,omitempty"`
	VirtualServerCreated                           string `json:"virtualserver_created,omitempty"`
	VirtualServerDefaultChannelAdminGroup          string `json:"virtualserver_default_channel_admin_group,omitempty"`
	VirtualServerDefaultChannelGroup               string `json:"virtualserver_default_channel_group,omitempty"`
	VirtualServerDefaultServerGroup                string `json:"virtualserver_default_server_group,omitempty"`
	VirtualServerDownloadQuota                     string `json:"virtualserver_download_quota,omitempty"`
	VirtualServerFileStorageClass                  string `json:"virtualserver_file_storage_class,omitempty"`
	VirtualServerFilebase                          string `json:"virtualserver_filebase,omitempty"`
	VirtualServerFlagPassword                      string `json:"virtualserver_flag_password,omitempty"`
	VirtualServerHostbannerGfxInterval             string `json:"virtualserver_hostbanner_gfx_interval,omitempty"`
	VirtualServerHostbannerGfxURL                  string `json:"virtualserver_hostbanner_gfx_url,omitempty"`
	VirtualServerHostbannerMode                    string `json:"virtualserver_hostbanner_mode,omitempty"`
	VirtualServerHostbannerURL                     string `json:"virtualserver_hostbanner_url,omitempty"`
	VirtualServerHostbuttonGfxURL                  string `json:"virtualserver_hostbutton_gfx_url,omitempty"`
	VirtualServerHostbuttonTooltip                 string `json:"virtualserver_hostbutton_tooltip,omitempty"`
	VirtualServerHostbuttonURL                     string `json:"virtualserver_hostbutton_url,omitempty"`
	VirtualServerHostmessage                       string `json:"virtualserver_hostmessage,omitempty"`
	VirtualServerHostmessageMode                   string `json:"virtualserver_hostmessage_mode,omitempty"`
	VirtualServerIconID                            string `json:"virtualserver_icon_id,omitempty"`
	VirtualServerID                                string `json:"virtualserver_id,omitempty"`
	VirtualServerIP                                string `json:"virtualserver_ip,omitempty"`
	VirtualServerLogChannel                        string `json:"virtualserver_log_channel,omitempty"`
	VirtualServerLogClient                         string `json:"virtualserver_log_client,omitempty"`
	VirtualServerLogFiletransfer                   string `json:"virtualserver_log_filetransfer,omitempty"`
	VirtualServerLogPermissions                    string `json:"virtualserver_log_permissions,omitempty"`
	VirtualServerLogQuery                          string `json:"virtualserver_log_query,omitempty"`
	VirtualServerLogServer                         string `json:"virtualserver_log_server,omitempty"`
	VirtualServerMachineID                         string `json:"virtualserver_machine_id,omitempty"`
	VirtualServerMaxDownloadTotalBandwidth         string `json:"virtualserver_max_download_total_bandwidth,omitempty"`
	VirtualServerMaxUploadTotalBandwidth           string `json:"virtualserver_max_upload_total_bandwidth,omitempty"`
	VirtualServerMaxclients                        string `json:"virtualserver_maxclients,omitempty"`
	VirtualServerMinAndroidVersion                 string `json:"virtualserver_min_android_version,omitempty"`
	VirtualServerMinClientVersion                  string `json:"virtualserver_min_client_version,omitempty"`
	VirtualServerMinClientsInChannelBeforeSilence  string `json:"virtualserver_min_clients_in_channel_before_forced_silence,omitempty"`
	VirtualServerMinIOSVersion                     string `json:"virtualserver_min_ios_version,omitempty"`
	VirtualServerMonthBytesDownloaded              string `json:"virtualserver_month_bytes_downloaded,omitempty"`
	VirtualServerMonthBytesUploaded                string `json:"virtualserver_month_bytes_uploaded,omitempty"`
	VirtualServerName                              string `json:"virtualserver_name,omitempty"`
	VirtualServerNamePhonetic                      string `json:"virtualserver_name_phonetic,omitempty"`
	VirtualServerNeededIdentitySecurityLevel       string `json:"virtualserver_needed_identity_security_level,omitempty"`
	VirtualServerNickname                          string `json:"virtualserver_nickname,omitempty"`
	VirtualServerPassword                          string `json:"virtualserver_password,omitempty"`
	VirtualServerPlatform                          string `json:"virtualserver_platform,omitempty"`
	VirtualServerPrioritySpeakerDimmModificator    string `json:"virtualserver_priority_speaker_dimm_modificator,omitempty"`
	VirtualServerQueryClientConnections            string `json:"virtualserver_query_client_connections,omitempty"`
	VirtualServerQueryclientsonline                string `json:"virtualserver_queryclientsonline,omitempty"`
	VirtualServerReservedSlots                     string `json:"virtualserver_reserved_slots,omitempty"`
	VirtualServerStatus                            string `json:"virtualserver_status,omitempty"`
	VirtualServerTotalBytesDownloaded              string `json:"virtualserver_total_bytes_downloaded,omitempty"`
	VirtualServerTotalBytesUploaded                string `json:"virtualserver_total_bytes_uploaded,omitempty"`
	VirtualServerTotalPacketlossControl            string `json:"virtualserver_total_packetloss_control,omitempty"`
	VirtualServerTotalPacketlossKeepalive          string `json:"virtualserver_total_packetloss_keepalive,omitempty"`
	VirtualServerTotalPacketlossSpeech             string `json:"virtualserver_total_packetloss_speech,omitempty"`
	VirtualServerTotalPacketlossTotal              string `json:"virtualserver_total_packetloss_total,omitempty"`
	VirtualServerTotalPing                         string `json:"virtualserver_total_ping,omitempty"`
	VirtualServerUniqueIdentifier                  string `json:"virtualserver_unique_identifier,omitempty"`
	VirtualServerUploadQuota                       string `json:"virtualserver_upload_quota,omitempty"`
	VirtualServerUptime                            string `json:"virtualserver_uptime,omitempty"`
	VirtualServerVersion                           string `json:"virtualserver_version,omitempty"`
	VirtualServerWeblistEnabled                    string `json:"virtualserver_weblist_enabled,omitempty"`
	VirtualServerWelcomemessage                    string `json:"virtualserver_welcomemessage,omitempty"`
	/*
		ConnectionPacketsSentTotal                 string `json:"connection_packets_sent_total,omitempty"`
		ConnectionBandwidthReceivedLastMinuteTotal string `json:"connection_bandwidth_received_last_minute_total,omitempty"`
		ConnectionBandwidthReceivedLastSecondTotal string `json:"connection_bandwidth_received_last_second_total,omitempty"`
		ConnectionBandwidthSentLastMinuteTotal     string `json:"connection_bandwidth_sent_last_minute_total,omitempty"`
		ConnectionBandwidthSentLastSecondTotal     string `json:"connection_bandwidth_sent_last_second_total,omitempty"`
		ConnectionPacketsReceivedTotal             string `json:"connection_packets_received_total,omitempty"`
		ConnectionBytesSentTotal                   string `json:"connection_bytes_sent_total,omitempty"`
		ConnectionFiletransferBandwidthReceived    string `json:"connection_filetransfer_bandwidth_received,omitempty"`
		ConnectionFiletransferBandwidthSent        string `json:"connection_filetransfer_bandwidth_sent,omitempty"`
		ConnectionFiletransferBytesReceivedTotal   string `json:"connection_filetransfer_bytes_received_total,omitempty"`
		ConnectionFiletransferBytesSentTotal       string `json:"connection_filetransfer_bytes_sent_total,omitempty"`
		ConnectionBytesReceivedTotal               string `json:"connection_bytes_received_total,omitempty"`
		VirtualServerPort                          string `json:"virtualserver_port,omitempty"`
	*/

	//	serverlist
	/*
		VirtualServerClientsOnline      string `json:"virtualserver_clientsonline,omitempty"`
		VirtualServerAutostart          string `json:"virtualserver_autostart,omitempty"`
		VirtualServerID                 string `json:"virtualserver_id,omitempty"`
		VirtualServerMachineID          string `json:"virtualserver_machine_id,omitempty"`
		VirtualServerMaxclients         string `json:"virtualserver_maxclients,omitempty"`
		VirtualServerName               string `json:"virtualserver_name,omitempty"`
		VirtualServerPort               string `json:"virtualserver_port,omitempty"`
		VirtualServerQueryclientsonline string `json:"virtualserver_queryclientsonline,omitempty"`
		VirtualServerStatus             string `json:"virtualserver_status,omitempty"`
		VirtualServerUniqueIdentifier   string `json:"virtualserver_unique_identifier,omitempty"`
		VirtualServerUptime             string `json:"virtualserver_uptime,omitempty"`
	*/

	//	servernotifyregister

	//	servernotifyunregister

	//	serverprocessstop

	//	serverrequestconnectioninfo
	ConnectionPacketlossTotal string `json:"connection_packetloss_total,omitempty"`
	ConnectionPing            string `json:"connection_ping,omitempty"`
	/*
		ConnectionBandwidthReceivedLastMinuteTotal string `json:"connection_bandwidth_received_last_minute_total,omitempty"`
		ConnectionBandwidthReceivedLastSecondTotal string `json:"connection_bandwidth_received_last_second_total,omitempty"`
		ConnectionBandwidthSentLastMinuteTotal     string `json:"connection_bandwidth_sent_last_minute_total,omitempty"`
		ConnectionBandwidthSentLastSecondTotal     string `json:"connection_bandwidth_sent_last_second_total,omitempty"`
		ConnectionBytesReceivedTotal               string `json:"connection_bytes_received_total,omitempty"`
		ConnectionBytesSentTotal                   string `json:"connection_bytes_sent_total,omitempty"`
		ConnectionConnectedTime                    string `json:"connection_connected_time,omitempty"`
		ConnectionFiletransferBandwidthReceived    string `json:"connection_filetransfer_bandwidth_received,omitempty"`
		ConnectionFiletransferBandwidthSent        string `json:"connection_filetransfer_bandwidth_sent,omitempty"`
		ConnectionFiletransferBytesReceivedTotal   string `json:"connection_filetransfer_bytes_received_total,omitempty"`
		ConnectionFiletransferBytesSentTotal       string `json:"connection_filetransfer_bytes_sent_total,omitempty"`
		ConnectionPacketsReceivedTotal             string `json:"connection_packets_received_total,omitempty"`
		ConnectionPacketsSentTotal                 string `json:"connection_packets_sent_total,omitempty"`
	*/

	//	serversnapshotcreate
	Data    string `json:"data,omitempty"`
	Version string `json:"version,omitempty"`

	//	serversnapshotdeploy

	//	serverstart

	//	serverstop

	//	servertemppasswordadd

	//	servertemppassworddel

	//	servertemppasswordlist
	NickName string `json:"nickname,omitempty"`
	Desc     string `json:"desc,omitempty"`
	PWClear  string `json:"pw_clear,omitempty"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
	TCid     string `json:"tcid,omitempty"`
	/*
		UID string `json:"uid,omitempty"`
	*/

	//	setclientchannelgroup

	//	tokenadd
	/*
		Token string `json:"token,omitempty"`
	*/

	//	tokendelete

	//	tokenlist
	/*
		Token            string `json:"token,omitempty"`
		TokenCreated     string  `json:"token_created,omitempty"`
		TokenCustomSet   string `json:"token_customset,omitempty"`
		TokenDescription string `json:"token_description,omitempty"`
		TokenID1         string  `json:"token_id1,omitempty"`
		TokenID2         string  `json:"token_id2,omitempty"`
		TokenType        string  `json:"token_type,omitempty"`
	*/

	//	tokenuse

	//	version
	Build    string `json:"build,omitempty"`
	Platform string `json:"platform,omitempty"`
	/*
		Version  string `json:"version,omitempty"`
	*/

	//	whoami
	ClientChannelID      string `json:"client_channel_id,omitempty"`
	ClientID             string `json:"client_id,omitempty"`
	ClientOriginServerID string `json:"client_origin_server_id,omitempty"`
	/*
		ClientUniqueIdentifier        string `json:"client_unique_identifier,omitempty"`
		VirtualServerID               string `json:"virtualserver_id,omitempty"`
		VirtualServerPort             string `json:"virtualserver_port,omitempty"`
		ClientNickname                string `json:"client_nickname,omitempty"`
		ClientLoginName               string `json:"client_login_name,omitempty"`
		VirtualServerStatus           string `json:"virtualserver_status,omitempty"`
		VirtualServerUniqueIdentifier string `json:"virtualserver_unique_identifier,omitempty"`
		ClientDatabaseID              string `json:"client_database_id,omitempty"`
	*/
}

type Status struct {
	Status Result `json:"status"`
}

type Result struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ExtraMessage string `json:"extra_message,omitempty"`
}
