package ts3_webquery

import "encoding/json"

/*
   clientaddperm               | assign permission to client
   clientdbdelete              | delete client database properties
   clientdbedit                | change client database properties
   clientdbfind                | find client database ID by nickname or UID
   clientdbinfo                | display client database properties
   clientdblist                | list known client UIDs
   clientdelperm               | remove permission from client
   clientedit                  | change client properties
   clientfind                  | find client by nickname
   clientgetdbidfromuid        | find client database ID by UID
   clientgetids                | find client IDs by UID
   clientgetnamefromdbid       | find client nickname by database ID
   clientgetnamefromuid        | find client nickname by UID
   clientgetuidfromclid        | find client UID by client ID
   clientinfo                  | display client properties
   clientkick                  | kick a client
   clientlist                  | list clients online on a virtual server
   clientmove                  | move a client
   clientpermlist              | list client specific permissions
   clientpoke                  | poke a client
   clientsetserverquerylogin   | set own login credentials
   clientupdate                | set own properties
*/

// clientaddperm               | assign permission to client
// Description:
// Parameters:
func (c *Client) ClientAddPerm() {}

// clientdbdelete              | delete client database properties
// Description:
// Parameters:
func (c *Client) ClientDBDelete() {}

// clientdbedit                | change client database properties
// Description:
// Parameters:
func (c *Client) ClientDBEdit() {}

// clientdbfind                | find client database ID by nickname or UID
// Description:
// Parameters:
func (c *Client) ClientDBFind() {}

// clientdbinfo                | display client database properties
// Description:
// Parameters:
func (c *Client) ClientDBInfo() {}

// clientdblist                | list known client UIDs
// Description:
// Parameters:
func (c *Client) ClientDBList() {}

// clientdelperm               | remove permission from client
// Description:
// Parameters:
func (c *Client) ClientDelPerm() {}

// clientedit                  | change client properties
// Description:
// Parameters:
func (c *Client) ClientEdit() {}

// clientfind                  | find client by nickname
// Description:
// Parameters:
func (c *Client) ClientFind() {}

// clientgetdbidfromuid        | find client database ID by UID
// Description:
// Parameters:
func (c *Client) ClientGetDBIDFormUID() {}

// clientgetids                | find client IDs by UID
// Description:
// Parameters:
func (c *Client) ClientGetIDS() {}

// clientgetnamefromdbid       | find client nickname by database ID
// Description:
// Parameters:
func (c *Client) ClientGetNameFormDBID() {}

// clientgetnamefromuid        | find client nickname by UID
// Description:
// Parameters:
func (c *Client) ClientGetNameFormUID() {}

// clientgetuidfromclid        | find client UID by client ID
// Description:
// Parameters:
func (c *Client) ClientGetUIDFromClID() {}

// clientinfo                  | display client properties
// Description:
// Parameters:
func (c *Client) ClientInfo() {}

// clientkick                  | kick a client
// Description:
// Parameters:
func (c *Client) ClientKick() {}

//	clientlist                  | 'list clients online on a virtual server'
//
// Description:
//
//	Displays a list of clients online on a virtual server including their ID,
//	nickname, status flags, etc. The output can be modified using several command
//	options.
//	Please note that the output will only contain clients which are currently in
//	channels you are able to subscribe to.
//
// Parameters:
//
//	'-uid' : Reports the client unique identifier
//	'-away' : Adds client away status and message
//	'-voice' : Adds information on whether the user is talking, their mute/hardware status and if they are recording/priority speaker
//	'-times' : Add idle, created and last connected time.
//	'-groups' : Includes info on server and channel groups.
//	'-info' : Adds client version and platform
//	'-country' : Reports the client country code
//	'-ip' : Adds the client IP, depending on the callers **PERMISSION_b_client_remoteaddress_view** permission
//	'-icon' : Adds the client icon id
//	'-badges' : Reports on the clients badges
func (c *Client) ClientList(data map[string]string) (*Response, error) {
	url := c.WebQuery + "/" + c.Sid + "/clientlist?api-key=" + c.APIKey
	body, err := Post(url, c.TimeOut, data)
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

// clientmove                  | move a client
// Description:
// Parameters:
func (c *Client) ClientMove() {}

// clientpermlist              | list client specific permissions
// Description:
// Parameters:
func (c *Client) ClientPermList() {}

// clientpoke                  | poke a client
// Description:
// Parameters:
func (c *Client) ClientPoke() {}

// clientsetserverquerylogin   | set own login credentials
// Description:
// Parameters:
func (c *Client) ClientSetServerQueryLogin() {}

// clientupdate                | set own properties
// Description:
// Parameters:
func (c *Client) ClientUpdate() {}
