package ts3_webquery

import (
	"encoding/json"
	"fmt"
)

/*
   channeladdperm              | assign permission to channel
   channelclientaddperm        | assign permission to channel-client combi
   channelclientdelperm        | remove permission from channel-client combi
   channelclientpermlist       | list channel-client specific permissions
   channelcreate               | create a channel
   channeldelete               | delete a channel
   channeldelperm              | remove permission from channel
   channeledit                 | change channel properties
   channelfind                 | find channel by name
   channelgroupadd             | create a channel group
   channelgroupaddperm         | assign permission to channel group
   channelgroupclientlist      | find channel groups by client ID
   channelgroupcopy            | copy a channel group
   channelgroupdel             | delete a channel group
   channelgroupdelperm         | remove permission from channel group
   channelgrouplist            | list channel groups
   channelgrouppermlist        | list channel group permissions
   channelgrouprename          | rename a channel group
   channelinfo                 | display channel properties
   channellist                 | list channels on a virtual server
   channelmove                 | move channel to new parent
   channelpermlist             | list channel specific permissions
*/

//	channeladdperm	            | 'assign permission to channel'
//
// Description:
//
//	Adds a set of specified permissions to a channel. Multiple permissions can be
//	added by providing the two parameters of each permission. A permission can be
//	specified by `permid` or `permsid`.
//
// Parameters:
//
//	'-continueonerror' : ignore some errors
//	'cid' : integer : id of channel
//	'permid' : integer : id of permission
//	'permsid' : server id of permission
//	'permvalue' : value of the permission
func (c *Client) ChannelAddPerm(data map[string]string) (*Status, error) {
	url := c.WebQuery + "/" + CheckServerID(c.Sid) + "/channeladdperm?api-key=" + c.APIKey
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelclientaddperm	    | 'assign permission to channel-client combi'
//
// Description:
//
//	Adds a set of specified permissions to a client in a specific channel.
//	Multiple permissions can be added by providing the three parameters of each
//	permission. A permission can be specified by permid or permsid.
//
// Parameters:
//
//	'-continueonerror' : continue processing on errors
//	'cid' : integer : id of channel
//	'permid' : integer : id of permission
//	'permsid' : server id of permission
//	'permvalue' : value of the permission
func (c *Client) ChannelClientAddPerm(data map[string]string) (*Status, error) {
	url := c.WebQuery + "/" + CheckServerID(c.Sid) + "/channelclientaddperm?api-key=" + c.APIKey
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelclientdelperm        | 'remove permission from channel-client combi'
//
// Description:
//
//	Removes a set of specified permissions from a client in a specific channel.
//	Multiple permissions can be removed at once. A permission can be specified by
//	 permid or permsid.
//
// Parameters:
//
//	'cid' : integer : id of channel
//	'cldbid' : integer : id of client
//	'permid' : integer : id of permission
//	'permsid' : server id of permission
func (c *Client) ChannelClientDelPerm(data map[string]string) (*Status, error) {
	url := c.WebQuery + "/" + CheckServerID(c.Sid) + "/channelclientdelperm?api-key=" + c.APIKey
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelclientpermlist       | 'list channel-client specific permissions'
//
// Description:
//
//	Displays a list of permissions defined for a client in a specific channel.
//
// Parameters:
//
//	'cid' : integer : id of channel
//	'cldbid' : integer : id of client
//	'-permsid' : include `permsid`
func (c *Client) ChannelClientPermList(data map[string]string) (*Response, error) {
	url := c.WebQuery + "/" + CheckServerID(c.Sid) + "/channelclientpermlist?api-key=" + c.APIKey
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channelcreate               | 'create a channel'
//
// Description:
//
//	Creates a new channel using the given properties and displays its ID. Note
//	that this command accepts multiple properties which means that you are able to
//	specifiy all settings of the new channel at once.
//	For detailed information, see Channel Properties.
//
// Parameters:
//
//	'channel_name' : Name of the channel
//	'channel_propertie'....
func (c *Client) ChannelCreate(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelcreate?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channeldelete               | 'delete a channel'
//
// Description:
//
//	Deletes an existing channel by ID. If force is set to 1, the channel will be
//	deleted even if there are clients within. The clients will be kicked to the
//	default channel with an appropriate reason message.
//
// Parameters:
//
//	'cid' : integer : id of channel
//	'force' : integer : set to 1 to force deletion
func (c *Client) ChannelDelete(cid string, force string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channeldelete?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	data := map[string]string{
		"cid":   cid,
		"force": force,
	}
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channeldelperm              | 'remove permission from channel'
//
// Description:
//
//	Removes a set of specified permissions from a channel. Multiple permissions
//	can be removed at once. A permission can be specified by permid or permsid.
//
// Parameters:
//
//	'cid' : integer : id of channel
//	'permid' : integer : id of permission
//	'permsid' : server id of permission
func (c *Client) ChannelDelPerm(data map[string]string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channeldelperm?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channeledit                 | 'change channel properties'
//
// Description:
//
//	Changes a channels configuration using given properties. Note that this
//	command accepts multiple properties which means that you are able to change
//	all settings of the channel specified with cid at once.
//	For detailed information, see Channel Properties.
//
// Parameters:
//
//	'cid' : integer : id of channel
//	'channel_properties'...
func (c *Client) ChannelEdit(data map[string]string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channeledit?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelfind                 | 'find channel by name'
//
// Description:
//
//	Displays a list of channels matching a given name pattern.
//
// Parameters:
//
//	'pattern' : pattern must be part of the channels name (case insensitive)
func (c *Client) ChannelFind(pattern string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channelfind?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	data := map[string]string{
		"pattern": pattern,
	}
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelgroupadd             | 'create a channel group'
//
// Description:
//
//	Creates a new channel group using a given name and displays its ID. The
//	optional type parameter can be used to create template groups.
//	For detailed information, see Definitions.
//
// Parameters:
//
//	'name' : name of the channel group
//	'type' : integer:`0` (template), `1` (regular), `2` (query); Default: `1`
func (c *Client) ChannelGroupAdd(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelgroupadd?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channelgroupaddperm         | 'assign permission to channel group'
//
// Description:
//
//	Adds a set of specified permissions to a channel group. Multiple permissions
//	can be added by providing the two parameters of each permission. A permission
//	can be specified by permid or permsid.
//
// Parameters:
//
//	'-continueonerror' : continue on some errors
//	'cgid' : integer : id of channel group
//	'permid' : integer : id of permission
//	'permsid' : server id of permission
//	'permvalue' : value of the permission
func (c *Client) ChannelGroupAddPerm(data map[string]string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channelgroupaddperm?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelgroupclientlist      | 'find channel groups by client ID'
//
// Description:
//
//	Displays all the client and/or channel IDs currently assigned to channel
//	groups. All three parameters are optional so you're free to choose the most
//	suitable combination for your requirements.
//
// Parameters:
//
//	'cid' : integer : only list entries of this channel
//	'cldbid' : integer : only list entries of this client
//	'cgid' : integer : only list entries of this channel group
func (c *Client) ChannelGroupClientList(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelgroupclientlist?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channelgroupcopy            | 'copy a channel group'
//
// Description:
//
//	Creates a copy of the channel group specified with scgid. If tcgid is set to
//	0, the server will create a new group. To overwrite an existing group, simply
//	set tcgid to the ID of a designated target group. If a target group is set,
//	the name parameter will be ignored. The type parameter can be used to create
//	template groups.
//	For detailed information, see Definitions.
//
// Parameters:
//
//	'scgid' : integer : the channel group to copy
//	'tcgid' : integer : the target group
//	'name' : string : the name of the new group (ignored if `tcgid` is set)
//	'type' : integer : group type `0` (template), `1` (regular), `2` (query)
func (c *Client) ChannelGroupCopy(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelgroupcopy?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channelgroupdel             | 'delete a channel group'
//
// Description:
//
//	Deletes a channel group by ID. If force is set to 1, the channel group will
//	be deleted even if there are clients within.
//
// Parameters:
//
//	'cgid' : integer : channel group id
//	'force' : integer : set to 1 to force deletion even if clients are present
func (c *Client) ChannelGroupDel(cgid string, force string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channelgroupdel?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	data := map[string]string{
		"cgid":  cgid,
		"force": force,
	}
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelgroupdelperm         | 'remove permission from channel group'
//
// Description:
// Parameters:
func (c *Client) ChannelGroupDelPerm(data map[string]string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channelgroupdelperm?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelgrouplist            | 'list channel groups'
//
// Description:
//
//	Removes a set of specified permissions from the channel group. Multiple
//	permissions can be removed at once. A permission can be specified by permid
//	or permsid.
//
// Parameters:
//
//	'-continueonerror' : continue on some errors
//	'cgid' : integer : id of channel group
//	'permid' : integer : id of permission
//	'permsid' : server id of permission
func (c *Client) ChannelGroupList() (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelgrouplist?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Get(url, CheckTimeOut(c.TimeOut))
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channelgrouppermlist        | 'list channel group permissions'
//
// Description:
//
//	Displays a list of permissions assigned to the channel group specified with
//	cgid. If the -permsid option is specified, the output will contain the
//	permission names instead of the internal IDs.
//
// Parameters:
//
//	'cgid' : integer : id of channel group
//	'-permsid' : include `permsid
func (c *Client) ChannelGroupPermList(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelgrouppermlist?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channelgrouprename          | 'rename a channel group'
//
// Description:
//
//	Changes the name of a specified channel group.
//
// Parameters:
//
//	'cgid' : integer:id of channel group
//	'name' : new name of the channel group
func (c *Client) ChannelGroupRename(cgid string, name string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channelgrouprename?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	data := map[string]string{
		"cgid": cgid,
		"name": name,
	}
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelinfo                 | 'display channel properties'
//
// Description:
//
//	Displays detailed configuration information about a channel including ID,
//	topic, description, etc.
//	For detailed information, see Channel Properties.
//
// Parameters:
//
//	'cid' : integer : The channel you're interested in
func (c *Client) ChannelInfo(cid string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelinfo?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	data := map[string]string{
		"cid": cid,
	}
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channellist                 | 'list channels on a virtual server'
//
// Description:
//
//	Displays a list of channels created on a virtual server including their ID,
//	order, name, etc. The output can be modified using several command options.
//
// Parameters:
//
//	'-voice' : include `channel_codec`, `channel_codec_quality` and `channel_needed_talk_power`
//	'-limits' : include `total_clients_family`
//	'-icon' : include `channel_icon_id`
//	'-secondsempty' : include `seconds_empty`
//	'-banners' : include `channel_banner_gfx_url` and `channel_banner_mode`
func (c *Client) ChannelList(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channellist?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	channelmove                 | 'move channel to new parent'
//
// Description:
//
//	Moves a channel to a new parent channel with the ID cpid. If order is
//	specified, the channel will be sorted right under the channel with the
//	specified ID. If order is set to 0, the channel will be sorted right below
//	the new parent.
//
// Parameters:
//
//	'cid' : integer : id of channel
//	'cpid' : integer : id of parent channel
//	'order' : integer : id of upper sibling channel
func (c *Client) ChannelMove(data map[string]string) (*Status, error) {
	url := fmt.Sprintf("%s/%s/channelmove?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

//	channelpermlist             | 'list channel specific permissions'
//
// Description:
//
//	Displays a list of permissions defined for a channel.
//
// Parameters:
//
//	'cid' : integer : id of channel
//	'-permsid' : include `permsid`
func (c *Client) ChannelPermList(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/%s/channelpermlist?api-key=%s", c.WebQuery, CheckServerID(c.Sid), c.APIKey)
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
