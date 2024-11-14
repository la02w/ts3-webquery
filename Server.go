package ts3_webquery

import (
	"encoding/json"
)

/*
   servercreate                | create a virtual server
   serverdelete                | delete a virtual server
   serveredit                  | change virtual server properties
   servergroupadd              | create a server group
   servergroupaddclient        | add client to server group
   servergroupaddperm          | assign permissions to server group
   servergroupautoaddperm      | globally assign permissions to server groups
   servergroupautodelperm      | globally remove permissions from server group
   servergroupclientlist       | list clients in a server group
   servergroupcopy             | create a copy of an existing server group
   servergroupdel              | delete a server group
   servergroupdelclient        | remove client from server group
   servergroupdelperm          | remove permissions from server group
   servergrouplist             | list server groups
   servergrouppermlist         | list server group permissions
   servergrouprename           | rename a server group
   servergroupsbyclientid      | get all server groups of specified client
   serveridgetbyport           | find database ID by virtual server port
   serverinfo                  | display virtual server properties
   serverlist                  | list virtual servers
   servernotifyregister        | register for event notifications
   servernotifyunregister      | unregister from event notifications
   serverprocessstop           | shutdown server process
   serverrequestconnectioninfo | display virtual server connection info
   serversnapshotcreate        | create snapshot of a virtual server
   serversnapshotdeploy        | deploy snapshot of a virtual server
   serverstart                 | start a virtual server
   serverstop                  | stop a virtual server
   servertemppasswordadd       | create a new temporary server password
   servertemppassworddel       | delete an existing temporary server password
   servertemppasswordlist      | list all existing temporary server passwords
*/

//	servercreate                | 'create a virtual server'
//
// Description:
// Parameters:
func (c *Client) ServerCreate() {}

//	serverdelete                | 'delete a virtual server'
//
// Description:
// Parameters:
func (c *Client) ServerDelete() {}

//	serveredit                  | 'change virtual server properties'
//
// Description:
// Parameters:
func (c *Client) ServerEdit() {}

//	servergroupadd              | 'create a server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupAdd() {}

//	servergroupaddclient        | 'add client to server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupAddClient() {}

//	servergroupaddperm          | 'assign permissions to server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupAddPerm() {}

//	servergroupautoaddperm      | 'globally assign permissions to server groups'
//
// Description:
// Parameters:
func (c *Client) ServerGroupAutoAddPerm() {}

//	servergroupautodelperm      | 'globally remove permissions from server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupAutoDelPerm() {}

//	servergroupclientlist       | 'list clients in a server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupClientList() {}

//	servergroupcopy             | 'create a copy of an existing server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupCopy() {}

//	servergroupdel              | 'delete a server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupDel() {}

//	servergroupdelclient        | 'remove client from server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupDelClient() {}

//	servergroupdelperm          | 'remove permissions from server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupDelPerm() {}

//	servergrouplist             | 'list server groups'
//
// Description:
// Parameters:
func (c *Client) ServerGroupList() {}

//	servergrouppermlist         | 'list server group permissions'
//
// Description:
// Parameters:
func (c *Client) ServerGroupPermList() {}

//	servergrouprename           | 'rename a server group'
//
// Description:
// Parameters:
func (c *Client) ServerGroupRename() {}

//	servergroupsbyclientid      | 'get all server groups of specified client'
//
// Description:
// Parameters:
func (c *Client) ServerGroupByClientID() {}

//	serveridgetbyport           | 'find database ID by virtual server port'
//
// Description:
// Parameters:
func (c *Client) ServerIDGetByPort() {}

//	serverinfo                  | 'display virtual server properties'
//
// Description:
//
//	Displays detailed configuration information about the selected virtual server
//	including unique ID, number of clients online, configuration, etc.
//	For detailed information, see Virtual Server Properties.
//
// Parameters:
func (c *Client) ServerInfo() (*Response, error) {
	url := c.WebQuery + "/" + CheckServerID(c.Sid) + "/serverinfo?api-key=" + c.APIKey
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

//	serverlist                  | 'list virtual servers'
//
// Description:
// Parameters:
func (c *Client) ServerList() {}

//	servernotifyregister        | 'register for event notifications'
//
// Description:
// Parameters:
func (c *Client) ServerNotifyRegister() {}

//	servernotifyunregister      | 'unregister from event notifications'
//
// Description:
// Parameters:
func (c *Client) ServerNotifyUnregister() {}

//	serverprocessstop           | 'shutdown server process'
//
// Description:
// Parameters:
func (c *Client) ServerProcessTop() {}

//	serverrequestconnectioninfo | 'display virtual server connection info'
//
// Description:
// Parameters:
func (c *Client) ServerRequestConnectionInfo() {}

//	serversnapshotcreate        | 'create snapshot of a virtual server'
//
// Description:
// Parameters:
func (c *Client) ServerSnapshotCreate() {}

//	serversnapshotdeploy        | 'deploy snapshot of a virtual server'
//
// Description:
// Parameters:
func (c *Client) ServerSnapshotDeploy() {}

//	serverstart                 | 'start a virtual server'
//
// Description:
// Parameters:
func (c *Client) ServerStart() {}

//	serverstop                  | 'stop a virtual server'
//
// Description:
// Parameters:
func (c *Client) SaverStop() {}

//	servertemppasswordadd       | 'create a new temporary server password'
//
// Description:
// Parameters:
func (c *Client) ServerTempPasswordAdd() {}

//	servertemppassworddel       | 'delete an existing temporary server password'
//
// Description:
// Parameters:
func (c *Client) ServerTempPasswordDel() {}

//	servertemppasswordlist      | 'list all existing temporary server passwords'
//
// Description:
// Parameters:
func (c *Client) ServerTempPasswordList() {}
