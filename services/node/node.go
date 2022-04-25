package node

import (
	"alfmigcli/services/user"
	"encoding/json"
	"time"
)

// AlbertoNode struct
type AlfrescoNode struct {
	NodeRef        string
	Name           string //interface{}
	QnamePath      interface{}
	Label          string
	Description    string
	createdDate    time.Time
	createdDateISO string
	Creator        user.User
}

// MessageResponse struct
type GeneralResponse struct {
	NumResults int
	Results    []AlfrescoNode
}

type List []AlfrescoNode

//Decode Json Response
func (g *GeneralResponse) Decode(text []byte) error {
	return json.Unmarshal(text, g)
}

//DataPrintToTable prints outs a formatted table
