package ts3_webquery

type Response struct {
	Body   *[]Body `json:"body,omitempty"`
	Status Status  `json:"status"`
}
type Body struct {
	// /1/version
	Build    *string `json:"build,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Version  *string `json:"version,omitempty"`
	// /1/hostinfo
	ConnectionBandwidthReceivedLastMinuteTotal *string `json:"connection_bandwidth_received_last_minute_total,omitempty"`
	ConnectionBandwidthReceivedLastSecondTotal *string `json:"connection_bandwidth_received_last_second_total,omitempty"`
	ConnectionBandwidthSentLastMinuteTotal     *string `json:"connection_bandwidth_sent_last_minute_total,omitempty"`
	ConnectionBandwidthSentLastSecondTotal     *string `json:"connection_bandwidth_sent_last_second_total,omitempty"`
	ConnectionBytesReceivedTotal               *string `json:"connection_bytes_received_total,omitempty"`
	ConnectionBytesSentTotal                   *string `json:"connection_bytes_sent_total,omitempty"`
	ConnectionFiletransferBandwidthReceived    *string `json:"connection_filetransfer_bandwidth_received,omitempty"`
	ConnectionFiletransferBandwidthSent        *string `json:"connection_filetransfer_bandwidth_sent,omitempty"`
	ConnectionFiletransferBytesReceivedTotal   *string `json:"connection_filetransfer_bytes_received_total,omitempty"`
	ConnectionFiletransferBytesSentTotal       *string `json:"connection_filetransfer_bytes_sent_total,omitempty"`
	ConnectionPacketsReceivedTotal             *string `json:"connection_packets_received_total,omitempty"`
	ConnectionPacketsSentTotal                 *string `json:"connection_packets_sent_total,omitempty"`
	HostTimestampUTC                           *string `json:"host_timestamp_utc,omitempty"`
	InstanceUptime                             *string `json:"instance_uptime,omitempty"`
	VirtualserversRunningTotal                 *string `json:"virtualservers_running_total,omitempty"`
	VirtualserversTotalChannelsOnline          *string `json:"virtualservers_total_channels_online,omitempty"`
	VirtualserversTotalClientsOnline           *string `json:"virtualservers_total_clients_online,omitempty"`
	VirtualserversTotalMaxclients              *string `json:"virtualservers_total_maxclients,omitempty"`
	// /1/instanceinfo
	ServerInstanceDatabaseVersion                *string `json:"serverinstance_database_version,omitempty"`
	ServerInstanceFiletransferPort               *string `json:"serverinstance_filetransfer_port,omitempty"`
	ServerInstanceGuestServerqueryGroup          *string `json:"serverinstance_guest_serverquery_group,omitempty"`
	ServerInstanceMaxDownloadTotalBandwidth      *string `json:"serverinstance_max_download_total_bandwidth,omitempty"`
	ServerInstanceMaxUploadTotalBandwidth        *string `json:"serverinstance_max_upload_total_bandwidth,omitempty"`
	ServerInstancePendingConnectionsPerIP        *string `json:"serverinstance_pending_connections_per_ip,omitempty"`
	ServerInstancePermissionsVersion             *string `json:"serverinstance_permissions_version,omitempty"`
	ServerInstanceServerqueryBanTime             *string `json:"serverinstance_serverquery_ban_time,omitempty"`
	ServerInstanceServerqueryFloodCommands       *string `json:"serverinstance_serverquery_flood_commands,omitempty"`
	ServerInstanceServerqueryFloodTime           *string `json:"serverinstance_serverquery_flood_time,omitempty"`
	ServerInstanceServerqueryMaxConnectionsPerIP *string `json:"serverinstance_serverquery_max_connections_per_ip,omitempty"`
	ServerInstanceTemplateChanneladminGroup      *string `json:"serverinstance_template_channeladmin_group,omitempty"`
	ServerInstanceTemplateChanneldefaultGroup    *string `json:"serverinstance_template_channeldefault_group,omitempty"`
	ServerInstanceTemplateServeradminGroup       *string `json:"serverinstance_template_serveradmin_group,omitempty"`
	ServerInstanceTemplateServerdefaultGroup     *string `json:"serverinstance_template_serverdefault_group,omitempty"`
	// /1/bindinglist
	IP *string `json:"ip,omitempty"`
	// /1/serverlist
	VirtualServerAutostart          *string `json:"virtualserver_autostart,omitempty"`
	VirtualServerClientsOnline      *string `json:"virtualserver_clientsonline,omitempty"`
	VirtualServerID                 *string `json:"virtualserver_id,omitempty"`
	VirtualServerMachineID          *string `json:"virtualserver_machine_id,omitempty"`
	VirtualServerMaxClients         *string `json:"virtualserver_maxclients,omitempty"`
	VirtualServerName               *string `json:"virtualserver_name,omitempty"`
	VirtualServerPort               *string `json:"virtualserver_port,omitempty"`
	VirtualServerQueryClientsOnline *string `json:"virtualserver_queryclientsonline,omitempty"`
	VirtualServerStatus             *string `json:"virtualserver_status,omitempty"`
	VirtualServerUptime             *string `json:"virtualserver_uptime,omitempty"`
	// /1/serveridgetbyport
	ServerID *string `json:"server_id,omitempty"`
	// /1/servercreate
	SID   *string `json:"sid,omitempty"`
	Token *string `json:"token,omitempty"`
	// /1/serverinfo
	ConnectionBytesReceivedControl                      *string `json:"connection_bytes_received_control,omitempty"`
	ConnectionBytesReceivedKeepalive                    *string `json:"connection_bytes_received_keepalive,omitempty"`
	ConnectionBytesReceivedSpeech                       *string `json:"connection_bytes_received_speech,omitempty"`
	ConnectionBytesSentControl                          *string `json:"connection_bytes_sent_control,omitempty"`
	ConnectionBytesSentKeepalive                        *string `json:"connection_bytes_sent_keepalive,omitempty"`
	ConnectionBytesSentSpeech                           *string `json:"connection_bytes_sent_speech,omitempty"`
	ConnectionPacketsReceivedControl                    *string `json:"connection_packets_received_control,omitempty"`
	ConnectionPacketsReceivedKeepalive                  *string `json:"connection_packets_received_keepalive,omitempty"`
	ConnectionPacketsReceivedSpeech                     *string `json:"connection_packets_received_speech,omitempty"`
	ConnectionPacketsSentControl                        *string `json:"connection_packets_sent_control,omitempty"`
	ConnectionPacketsSentKeepalive                      *string `json:"connection_packets_sent_keepalive,omitempty"`
	ConnectionPacketsSentSpeech                         *string `json:"connection_packets_sent_speech,omitempty"`
	VirtualServerAntifloodPointsNeededCommandBlock      *string `json:"virtualserver_antiflood_points_needed_command_block,omitempty"`
	VirtualServerAntifloodPointsNeededIPBlock           *string `json:"virtualserver_antiflood_points_needed_ip_block,omitempty"`
	VirtualServerAntifloodPointsNeededPluginBlock       *string `json:"virtualserver_antiflood_points_needed_plugin_block,omitempty"`
	VirtualServerAntifloodPointsTickReduce              *string `json:"virtualserver_antiflood_points_tick_reduce,omitempty"`
	VirtualServerAskForPrivilegekey                     *string `json:"virtualserver_ask_for_privilegekey,omitempty"`
	VirtualServerCapabilityExtensions                   *string `json:"virtualserver_capability_extensions,omitempty"`
	VirtualServerChannelTempDeleteDelayDefault          *string `json:"virtualserver_channel_temp_delete_delay_default,omitempty"`
	VirtualServerChannelsOnline                         *string `json:"virtualserver_channelsonline,omitempty"`
	VirtualServerClientConnections                      *string `json:"virtualserver_client_connections,omitempty"`
	VirtualServerCodecEncryptionMode                    *string `json:"virtualserver_codec_encryption_mode,omitempty"`
	VirtualServerComplainAutobanCount                   *string `json:"virtualserver_complain_autoban_count,omitempty"`
	VirtualServerComplainAutobanTime                    *string `json:"virtualserver_complain_autoban_time,omitempty"`
	VirtualServerComplainRemoveTime                     *string `json:"virtualserver_complain_remove_time,omitempty"`
	VirtualServerCreated                                *string `json:"virtualserver_created,omitempty"`
	VirtualServerDefaultChannelAdminGroup               *string `json:"virtualserver_default_channel_admin_group,omitempty"`
	VirtualServerDefaultChannelGroup                    *string `json:"virtualserver_default_channel_group,omitempty"`
	VirtualServerDefaultServerGroup                     *string `json:"virtualserver_default_server_group,omitempty"`
	VirtualServerDownloadQuota                          *string `json:"virtualserver_download_quota,omitempty"`
	VirtualServerFileStorageClass                       *string `json:"virtualserver_file_storage_class,omitempty"`
	VirtualServerFilebase                               *string `json:"virtualserver_filebase,omitempty"`
	VirtualServerFlagPassword                           *string `json:"virtualserver_flag_password,omitempty"`
	VirtualServerHostbannerGfxInterval                  *string `json:"virtualserver_hostbanner_gfx_interval,omitempty"`
	VirtualServerHostbannerGfxURL                       *string `json:"virtualserver_hostbanner_gfx_url,omitempty"`
	VirtualServerHostbannerMode                         *string `json:"virtualserver_hostbanner_mode,omitempty"`
	VirtualServerHostbannerURL                          *string `json:"virtualserver_hostbanner_url,omitempty"`
	VirtualServerHostbuttonGfxURL                       *string `json:"virtualserver_hostbutton_gfx_url,omitempty"`
	VirtualServerHostbuttonTooltip                      *string `json:"virtualserver_hostbutton_tooltip,omitempty"`
	VirtualServerHostbuttonURL                          *string `json:"virtualserver_hostbutton_url,omitempty"`
	VirtualServerHostmessage                            *string `json:"virtualserver_hostmessage,omitempty"`
	VirtualServerHostmessageMode                        *string `json:"virtualserver_hostmessage_mode,omitempty"`
	VirtualServerIconID                                 *string `json:"virtualserver_icon_id,omitempty"`
	VirtualServerIP                                     *string `json:"virtualserver_ip,omitempty"`
	VirtualServerLogChannel                             *string `json:"virtualserver_log_channel,omitempty"`
	VirtualServerLogClient                              *string `json:"virtualserver_log_client,omitempty"`
	VirtualServerLogFiletransfer                        *string `json:"virtualserver_log_filetransfer,omitempty"`
	VirtualServerLogPermissions                         *string `json:"virtualserver_log_permissions,omitempty"`
	VirtualServerLogQuery                               *string `json:"virtualserver_log_query,omitempty"`
	VirtualServerLogServer                              *string `json:"virtualserver_log_server,omitempty"`
	VirtualServerMaxDownloadTotalBandwidth              *string `json:"virtualserver_max_download_total_bandwidth,omitempty"`
	VirtualServerMaxUploadTotalBandwidth                *string `json:"virtualserver_max_upload_total_bandwidth,omitempty"`
	VirtualServerMinAndroidVersion                      *string `json:"virtualserver_min_android_version,omitempty"`
	VirtualServerMinClientVersion                       *string `json:"virtualserver_min_client_version,omitempty"`
	VirtualServerMinClientsInChannelBeforeForcedSilence *string `json:"virtualserver_min_clients_in_channel_before_forced_silence,omitempty"`
	VirtualServerMinIOSVersion                          *string `json:"virtualserver_min_ios_version,omitempty"`
	VirtualServerMonthBytesDownloaded                   *string `json:"virtualserver_month_bytes_downloaded,omitempty"`
	VirtualServerMonthBytesUploaded                     *string `json:"virtualserver_month_bytes_uploaded,omitempty"`
	VirtualServerNamePhonetic                           *string `json:"virtualserver_name_phonetic,omitempty"`
	VirtualServerNeededIdentitySecurityLevel            *string `json:"virtualserver_needed_identity_security_level,omitempty"`
	VirtualServerNickname                               *string `json:"virtualserver_nickname,omitempty"`
	VirtualServerPassword                               *string `json:"virtualserver_password,omitempty"`
	VirtualServerPlatform                               *string `json:"virtualserver_platform,omitempty"`
	VirtualServerPrioritySpeakerDimmModificator         *string `json:"virtualserver_priority_speaker_dimm_modificator,omitempty"`
	VirtualServerQueryClientConnections                 *string `json:"virtualserver_query_client_connections,omitempty"`
	VirtualServerReservedSlots                          *string `json:"virtualserver_reserved_slots,omitempty"`
	VirtualServerTotalBytesDownloaded                   *string `json:"virtualserver_total_bytes_downloaded,omitempty"`
	VirtualServerTotalBytesUploaded                     *string `json:"virtualserver_total_bytes_uploaded,omitempty"`
	VirtualServerTotalPacketlossControl                 *string `json:"virtualserver_total_packetloss_control,omitempty"`
	VirtualServerTotalPacketlossKeepalive               *string `json:"virtualserver_total_packetloss_keepalive,omitempty"`
	VirtualServerTotalPacketlossSpeech                  *string `json:"virtualserver_total_packetloss_speech,omitempty"`
	VirtualServerTotalPacketlossTotal                   *string `json:"virtualserver_total_packetloss_total,omitempty"`
	VirtualServerTotalPing                              *string `json:"virtualserver_total_ping,omitempty"`
	VirtualServerUniqueIdentifier                       *string `json:"virtualserver_unique_identifier,omitempty"`
	VirtualServerUploadQuota                            *string `json:"virtualserver_upload_quota,omitempty"`
	VirtualServerVersion                                *string `json:"virtualserver_version,omitempty"`
	VirtualServerWeblistEnabled                         *string `json:"virtualserver_weblist_enabled,omitempty"`
	VirtualServerWelcomemessage                         *string `json:"virtualserver_welcomemessage,omitempty"`
	// /1/serverrequestconnectioninfo
	ConnectionConnectedTime   *string `json:"connection_connected_time,omitempty"`
	ConnectionPacketlossTotal *string `json:"connection_packetloss_total,omitempty"`
	ConnectionPing            *string `json:"connection_ping,omitempty"`
	// /1/servertemppasswordlist
	Nickname *string `json:"nickname,omitempty"`
	UID      *string `json:"uid,omitempty"`
	Desc     *string `json:"desc,omitempty"`
	PwClear  *string `json:"pw_clear,omitempty"`
	Start    *string `json:"start,omitempty"`
	End      *string `json:"end,omitempty"`
	Tcid     *string `json:"tcid,omitempty"`
	// /1/servergrouplist
	IconID         *string `json:"iconid,omitempty"`
	NMemberAddP    *string `json:"n_member_addp,omitempty"`
	NMemberRemoveP *string `json:"n_member_removep,omitempty"`
	NModifyP       *string `json:"n_modifyp,omitempty"`
	Name           *string `json:"name,omitempty"`
	NameMode       *string `json:"namemode,omitempty"`
	SaveDB         *string `json:"savedb,omitempty"`
	SGID           *string `json:"sgid,omitempty"`
	SortID         *string `json:"sortid,omitempty"`
	Type           *string `json:"type,omitempty"`
	// /1/servergrouppermlist
	PermID      *string `json:"permid,omitempty"`
	PermNegated *string `json:"permnegated,omitempty"`
	PermSkip    *string `json:"permskip,omitempty"`
	PermValue   *string `json:"permvalue,omitempty"`
	// /1/servergroupclientlist
	CLDBID *string `json:"cldbid,omitempty"`
	// /1/servergroupsbyclientid
	// ...
	// /1/serversnapshotcreate 异常
	Data *string `json:"data,omitempty"`
	// /1/logview
	FileSize *string `json:"file_size,omitempty"`
	L        *string `json:"l,omitempty"`
	LastPos  *string `json:"last_pos,omitempty"`
	// /1/channellist
	ChannelName                 *string `json:"channel_name,omitempty"`
	ChannelNeededSubscribePower *string `json:"channel_needed_subscribe_power,omitempty"`
	ChannelOrder                *string `json:"channel_order,omitempty"`
	Cid                         *string `json:"cid,omitempty"`
	Pid                         *string `json:"pid,omitempty"`
	TotalClients                *string `json:"total_clients,omitempty"`
	// /1/channelinfo
	ChannelBannerGfxURL                  *string `json:"channel_banner_gfx_url,omitempty"`
	ChannelBannerMode                    *string `json:"channel_banner_mode,omitempty"`
	ChannelCodec                         *string `json:"channel_codec,omitempty"`
	ChannelCodecIsUnencrypted            *string `json:"channel_codec_is_unencrypted,omitempty"`
	ChannelCodecLatencyFactor            *string `json:"channel_codec_latency_factor,omitempty"`
	ChannelCodecQuality                  *string `json:"channel_codec_quality,omitempty"`
	ChannelDeleteDelay                   *string `json:"channel_delete_delay,omitempty"`
	ChannelDescription                   *string `json:"channel_description,omitempty"`
	ChannelFilepath                      *string `json:"channel_filepath,omitempty"`
	ChannelFlagDefault                   *string `json:"channel_flag_default,omitempty"`
	ChannelFlagMaxclientsUnlimited       *string `json:"channel_flag_maxclients_unlimited,omitempty"`
	ChannelFlagMaxfamilyclientsInherited *string `json:"channel_flag_maxfamilyclients_inherited,omitempty"`
	ChannelFlagMaxfamilyclientsUnlimited *string `json:"channel_flag_maxfamilyclients_unlimited,omitempty"`
	ChannelFlagPassword                  *string `json:"channel_flag_password,omitempty"`
	ChannelFlagPermanent                 *string `json:"channel_flag_permanent,omitempty"`
	ChannelFlagSemiPermanent             *string `json:"channel_flag_semi_permanent,omitempty"`
	ChannelForcedSilence                 *string `json:"channel_forced_silence,omitempty"`
	ChannelIconID                        *string `json:"channel_icon_id,omitempty"`
	ChannelMaxclients                    *string `json:"channel_maxclients,omitempty"`
	ChannelMaxfamilyclients              *string `json:"channel_maxfamilyclients,omitempty"`
	ChannelNamePhonetic                  *string `json:"channel_name_phonetic,omitempty"`
	ChannelNeededTalkPower               *string `json:"channel_needed_talk_power,omitempty"`
	ChannelPassword                      *string `json:"channel_password,omitempty"`
	ChannelSecuritySalt                  *string `json:"channel_security_salt,omitempty"`
	ChannelTopic                         *string `json:"channel_topic,omitempty"`
	ChannelUniqueIdentifier              *string `json:"channel_unique_identifier,omitempty"`
	SecondsEmpty                         *string `json:"seconds_empty,omitempty"`
	// /1/channelfind
	// channel_name,cid
	// /1/channelcreate
	// cid
	// /1/channelgrouplist
	Cgid *string `json:"cgid,omitempty"`
	// /1/channelgroupclientlist
	// cid=2 cldbid=9 cgid=9
	// /1/tokenadd
	// token=eKnFZQ9EK7G7MhtuQB6+N2B1PNZZ6OZL3ycDp2OW
	// /1/tokenadd
	TokenCreated     *string `json:"token_created,omitempty"`
	TokenCustomSet   *string `json:"token_customset,omitempty"`
	TokenDescription *string `json:"token_description,omitempty"`
	TokenID1         *string `json:"token_id1,omitempty"`
	TokenID2         *string `json:"token_id2,omitempty"`
	TokenType        *string `json:"token_type,omitempty"`
	// /1/channelpermlist
	// cid=2 permid=4353 permvalue=1 permnegated=0 permskip=0|permid=17276 permvalue=50 …
	// /1/clientlist
	ClientDatabaseID *string `json:"client_database_id,omitempty"`
	ClientNickname   *string `json:"client_nickname,omitempty"`
	ClientType       *string `json:"client_type,omitempty"`
	// /1/clientinfo
	ClientAway                           *string `json:"client_away,omitempty"`
	ClientAwayMessage                    *string `json:"client_away_message,omitempty"`
	ClientBadges                         *string `json:"client_badges,omitempty"`
	ClientBase64HashClientUID            *string `json:"client_base64HashClientUID,omitempty"`
	ClientChannelGroupID                 *string `json:"client_channel_group_id,omitempty"`
	ClientChannelGroupInheritedChannelID *string `json:"client_channel_group_inherited_channel_id,omitempty"`
	ClientCountry                        *string `json:"client_country,omitempty"`
	ClientCreated                        *string `json:"client_created,omitempty"`
	ClientDefaultChannel                 *string `json:"client_default_channel,omitempty"`
	ClientDefaultToken                   *string `json:"client_default_token,omitempty"`
	ClientDescription                    *string `json:"client_description,omitempty"`
	ClientFlagAvatar                     *string `json:"client_flag_avatar,omitempty"`
	ClientIconID                         *string `json:"client_icon_id,omitempty"`
	ClientIdleTime                       *string `json:"client_idle_time,omitempty"`
	ClientInputHardware                  *string `json:"client_input_hardware,omitempty"`
	ClientInputMuted                     *string `json:"client_input_muted,omitempty"`
	ClientIntegrations                   *string `json:"client_integrations,omitempty"`
	ClientIsChannelCommander             *string `json:"client_is_channel_commander,omitempty"`
	ClientIsPrioritySpeaker              *string `json:"client_is_priority_speaker,omitempty"`
	ClientIsRecording                    *string `json:"client_is_recording,omitempty"`
	ClientIsTalker                       *string `json:"client_is_talker,omitempty"`
	ClientLastconnected                  *string `json:"client_lastconnected,omitempty"`
	ClientLoginName                      *string `json:"client_login_name,omitempty"`
	ClientMetaData                       *string `json:"client_meta_data,omitempty"`
	ClientMonthBytesDownloaded           *string `json:"client_month_bytes_downloaded,omitempty"`
	ClientMonthBytesUploaded             *string `json:"client_month_bytes_uploaded,omitempty"`
	ClientMyteamspeakAvatar              *string `json:"client_myteamspeak_avatar,omitempty"`
	ClientMyteamspeakID                  *string `json:"client_myteamspeak_id,omitempty"`
	ClientNeededServerqueryViewPower     *string `json:"client_needed_serverquery_view_power,omitempty"`
	ClientNicknamePhonetic               *string `json:"client_nickname_phonetic,omitempty"`
	ClientOutputHardware                 *string `json:"client_output_hardware,omitempty"`
	ClientOutputMuted                    *string `json:"client_output_muted,omitempty"`
	ClientOutputonlyMuted                *string `json:"client_outputonly_muted,omitempty"`
	ClientPlatform                       *string `json:"client_platform,omitempty"`
	ClientSecurityHash                   *string `json:"client_security_hash,omitempty"`
	ClientServergroups                   *string `json:"client_servergroups,omitempty"`
	ClientSignedBadges                   *string `json:"client_signed_badges,omitempty"`
	ClientTalkPower                      *string `json:"client_talk_power,omitempty"`
	ClientTalkRequest                    *string `json:"client_talk_request,omitempty"`
	ClientTalkRequestMsg                 *string `json:"client_talk_request_msg,omitempty"`
	ClientTotalBytesDownloaded           *string `json:"client_total_bytes_downloaded,omitempty"`
	ClientTotalBytesUploaded             *string `json:"client_total_bytes_uploaded,omitempty"`
	ClientTotalconnections               *string `json:"client_totalconnections,omitempty"`
	ClientUniqueIdentifier               *string `json:"client_unique_identifier,omitempty"`
	ClientUnreadMessages                 *string `json:"client_unread_messages,omitempty"`
	ClientVersion                        *string `json:"client_version,omitempty"`
	ClientVersionSign                    *string `json:"client_version_sign,omitempty"`
	ConnectionClientIP                   *string `json:"connection_client_ip,omitempty"`
	// /1/clientfind
	Clid *string `json:"clid,omitempty"`
	// /1/clientdblist
	ClientLastIP *string `json:"client_lastip,omitempty"`
	// /1/clientdbinfo
	// ...
	// /1/clientdbfind
	// cldbid=56
	// /1/clientgetids 异常
	// ... 许多待测试命令
	// /1/permissionlist
	PermDesc *string `json:"permdesc,omitempty"`
	PermName *string `json:"permname,omitempty"`
	// /1/permreset
	// token
	// /1/privilegekeylist
	// ...
	// /1/whoami
	ClientChannelID      *string `json:"client_channel_id,omitempty"`
	ClientID             *string `json:"client_id,omitempty"`
	ClientOriginServerID *string `json:"client_origin_server_id,omitempty"`
}
type Status struct {
	Code         int     `json:"code"`
	ExtraMessage *string `json:"extra_message,omitempty"`
	Message      string  `json:"message"`
}
