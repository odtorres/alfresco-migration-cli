package workflow

import (
	"alfmigcli/config"
	"alfmigcli/services/user"
	"alfmigcli/services/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//Stringer interface to be printed
type Stringer interface {
	String() string
}

type WfDefAllResponse struct {
	Data []Wfdefinition
}

type WfResponse struct {
	Data WfItem
}

// Item struct represents a item
type WfItem struct {
	Id                  string       `json:"id"`
	Url                 string       `json:"url"`
	Name                string       `json:"name"`
	Title               string       `json:"title"`
	Description         string       `json:"description"`
	Version             string       `json:"version"`
	IsActive            bool         `json:"isActive"`
	StartDate           time.Time    `json:"startDate"`
	Priority            int          `json:"priority"`
	Message             string       `json:"message"`
	EndDate             time.Time    `json:"endDate"`
	DueDate             time.Time    `json:"dueDate"`
	Context             string       `json:"context"`
	Package             string       `json:"package"`
	Initiator           user.User    `json:"initiator"`
	DefinitionUrl       string       `json:"definitionUrl"`
	DiagramUrl          string       `json:"diagramUrl"`
	StartTaskInstanceId string       `json:"startTaskInstanceId"`
	Definition          Wfdefinition `json:"definition"`
	Tasks               []Wftask     `json:"tasks"`
}

type Wfdefinition struct {
	Id                      string        `json:"id"`
	Url                     string        `json:"url"`
	Name                    string        `json:"name"`
	Title                   string        `json:"title"`
	Description             string        `json:"description"`
	Version                 string        `json:"version"`
	StartTaskDefinitionUrl  string        `json:"startTaskDefinitionUrl"`
	StartTaskDefinitionType string        `json:"startTaskDefinitionType"`
	TaskDefinitions         []interface{} `json:"taskDefinitions"`
}

type Wftask struct {
	Id             string                 `json:"id"`
	Url            string                 `json:"url"`
	Name           string                 `json:"name"`
	Title          string                 `json:"title"`
	Description    string                 `json:"description"`
	State          string                 `json:"state"`
	Path           string                 `json:"path"`
	IsPooled       bool                   `json:"isPooled"`
	IsEditable     bool                   `json:"isEditable"`
	IsReassignable bool                   `json:"isReassignable"`
	IsClaimable    bool                   `json:"isClaimable"`
	IsReleasable   bool                   `json:"isReleasable"`
	Outcome        string                 `json:"outcome"`
	Owner          user.User              `json:"owner"`
	Creator        user.User              `json:"creator"`
	Properties     map[string]interface{} `json:"properties"`
}

type ResponseEnvelopeCreateWF struct {
	Entry ResponseCreateWF `json:"entry"`
	Error ErrorCreateWF    `json:"error"`
}

type ResponseCreateWF struct {
	Id                        string `json:"id"`
	ProcessDefinitionId       string `json:"processDefinitionId"`
	ProcessDefinitionKey      string `json:"processDefinitionKey"`
	BusinessKey               string `json:"businessKey"`
	StartedAt                 string `json:"startedAt"`
	EndedAt                   string `json:"endedAt"`
	DurationInMs              int    `json:"durationInMs"`
	StartActivityDefinitionId string `json:"startActivityDefinitionId"`
	EndActivityDefinitionId   string `json:"endActivityDefinitionId"`
	StartUserId               string `json:"startUserId"`
	DeleteReason              string `json:"deleteReason"`
}
type ErrorCreateWF struct {
	ErrorKey       string `json:"errorKey"`
	StatusCode     int    `json:"statusCode"`
	BriefSummary   string `json:"id"`
	StackTrace     string `json:"stackTrace"`
	DescriptionURL string `json:"descriptionURL"`
	LogId          string `json:"logId"`
}

//List of workflows
type List []WfItem

//List of workflows
type ListVerify []WfItem

//Decode Json Response
func (l *ResponseEnvelopeCreateWF) Decode(text []byte) error {
	return json.Unmarshal(text, l)
}

func (l *WfDefAllResponse) Decode(text []byte) error {
	return json.Unmarshal(text, l)
}

func (l *WfResponse) Decode(text []byte) error {
	return json.Unmarshal(text, l)
}

//PrintToTable prints outs a formatted table
func (l *WfDefAllResponse) PrintToTable() {
	var data = [][]string{}
	for k, t := range *&l.Data {
		data = append(data, []string{fmt.Sprintf("%d", k+1), t.Id, t.Name, t.Version})
	}
	util.PrintToTable([]string{"#", "Id", "Name", "Version"}, data)
}

func (l *ResponseEnvelopeCreateWF) PrintToTable() {
	var data = [][]string{}

	if (l.Error == ErrorCreateWF{}) {
		data = append(data, []string{fmt.Sprintf("%d", 1), l.Entry.Id, l.Entry.ProcessDefinitionId, l.Entry.ProcessDefinitionKey})
	} else {
		data = append(data, []string{fmt.Sprintf("%d", 1), l.Error.ErrorKey, l.Error.BriefSummary, l.Error.StackTrace})
	}

	util.PrintToTable([]string{"#", "Id", "ProcessDefinition", "Business"}, data)
}

func (l *WfResponse) PrintToTable() {
	var data = [][]string{}

	data = append(data, []string{fmt.Sprintf("%d", 1), l.Data.Id, l.Data.Name, fmt.Sprintf("%d", len(l.Data.Tasks))})
	util.PrintToTable([]string{"#", "Id", "Name", "Tasks"}, data)
}

func (l *List) Add(t WfItem) {
	*l = append(*l, t)
}

/*func (l *List) Get(pos int) WfItem {
	return *l[pos]
}*/

func (l *List) Save(file_name string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file_name+".json", js, 0644)
}

func (l *List) GetFromFile() error {
	file, err := ioutil.ReadFile(config.FileNameWorkflow)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

func (l *List) GetByFileName(file_name string) error {
	file, err := ioutil.ReadFile(file_name + ".json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// Appends it to the list
func (l *ListVerify) Add(t WfItem) {
	*l = append(*l, t)
}

func (l *ListVerify) Save() error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.FileNameWorkflowVerify, js, 0644)
}

func (l *ListVerify) GetByFileName() error {
	file, err := ioutil.ReadFile(config.FileNameWorkflowVerify)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
