package node

import (
	"alfmigcli/services/util"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// AlbertoNode struct
type AlbertoNode struct {
	Updated_at       string
	Unique_id        string
	Type             string
	Title            string
	Tenant           string
	Secondary_parent []string
	Read_privileged  []string
	Pathfs           string
	Parent           string
	Opts             string
	Nodetype         string
	Name             string
	Modifier         string
	Inserted_at      string
	Description      string
	Data             interface{}
	Creator          string
	Content          bool
	Alter_privileged []string
}

// GeneralResponse struct
type GeneralResponse struct {
	Success  bool
	Response AlbertoNode
}

// MessageResponse struct
type MessageResponse struct {
	Success  bool
	Response map[string]AlbertoNode
}

//Decode Json Response
func (g *GeneralResponse) Decode(text []byte) error {
	return json.Unmarshal(text, g)
}

//Decode Json Response
func (g *MessageResponse) Decode(text []byte) error {
	return json.Unmarshal(text, g)
}

//DataPrintToTable prints outs a formatted table
func (a *AlbertoNode) DataPrintToTable() {
	data := [][]string{}
	var i = 0
	for k, v := range a.Data.(map[string]interface{}) {

		switch vv := v.(type) {
		case string:
			data = append(data, []string{k, v.(string)})
		case float64:
			data = append(data, []string{k, fmt.Sprintf("%f", v.(float64))})
		case []string:
			data = append(data, []string{k, strings.Join(v.([]string), ", ")})
		case []interface{}:
			var array []string
			for _, u := range vv {
				array = append(array, fmt.Sprintf("%v", u))
			}
			data = append(data, []string{k, fmt.Sprintf("%v", strings.Join(array, ", "))})
		default:
			fmt.Println(k, vv, "is of a type I don't know how to handle")
		}
		i++
	}

	util.PrintToTableFooter([]string{"Name", "Value"}, []string{"", "Data"}, data)
}

//PropertiesPrintToTable prints outs a formatted table
func (a *AlbertoNode) PropertiesPrintToTable() {
	data := [][]string{
		{"parent", a.Parent},
		{"secondary_parent", strings.Join(a.Secondary_parent, ", ")},
		{"updated_at", a.Updated_at},
		{"unique_id", a.Unique_id},
		{"type", a.Type},
		{"title", a.Title},
		{"tenant", a.Tenant},
		{"read_privileged", strings.Join(a.Read_privileged, ", ")},
		{"pathfs", a.Pathfs},
		{"opts", a.Opts},
		{"nodetype", a.Nodetype},
		{"name", a.Name},
		{"modifier", a.Modifier},
		{"inserted_at", a.Inserted_at},
		{"description", a.Description},
		{"creator", a.Creator},
		{"content", strconv.FormatBool(a.Content)},
		{"alter_privileged", strings.Join(a.Alter_privileged, ", ")},
	}

	util.PrintToTableFooter([]string{"Name", "Value"}, []string{"", "Properties"}, data)
}
