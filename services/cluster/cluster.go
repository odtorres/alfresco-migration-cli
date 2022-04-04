package cluster

import (
	"alfmigcli/config"
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

// Item struct represents a item
type Item struct {
	ClusterName   string
	ClusterURL    string
	ClusterTICKET string
	Visible       bool
	Current       bool
	CreatedAt     time.Time
}

//List of clusters
type List []Item

// Add creates a new todo item and appends it to the list
func (l *List) Add(ClusterName string, ClusterURL string) {
	t := Item{
		ClusterName:   ClusterName,
		ClusterURL:    ClusterURL,
		ClusterTICKET: "",
		Visible:       true,
		Current:       false,
		CreatedAt:     time.Now(),
	}
	*l = append(*l, t)
}

// Visible method marks a item as visible by
// setting Visible = true
func (l *List) Visible(i int, visible bool) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	// Adjusting index for 0 based index
	ls[i-1].Visible = visible
	return nil
}

// Current method marks a item as a current cluster by
// setting Current = true
func (l *List) Current(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	// Adjusting index for 0 based index

	for index := range *l {
		if index != i-1 {
			ls[index].Current = false
		}

	}
	ls[i-1].Current = true
	return nil
}

// Delete method deletes a item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	// Adjusting index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save() error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.FileNameConfig, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get() error {
	file, err := ioutil.ReadFile(config.FileNameConfig)
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

// GetCurrent get the current cluster
func (l *List) GetCurrent() (Item, error) {
	for _, element := range *l {
		if element.Current {
			return element, nil
		}
	}
	return Item{}, fmt.Errorf("there is no current cluster")
}

// SetClusterTICKET set the current cluster ticket
func (l *List) SetClusterTICKET(ticket string) error {
	ls := *l
	for i, element := range *l {
		if element.Current {
			ls[i].ClusterTICKET = ticket
			return nil
		}
	}
	return fmt.Errorf("there is no current cluster")
}

//String prints outs a formatted list
//Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "   "
		loged := "   "
		if t.Current {
			prefix = util.IfWindowsElse(" ✔ ", " 👉 ")
		} else {
			prefix = util.IfWindowsElse("   ", "    ")
		}
		if t.ClusterTICKET != "" {
			loged = " ✓ "
		}
		if t.Visible {
			//prefix = " 👁️ "
			// Adjust the item number k to print numbers starting from 1 instead of 0
			formatted += fmt.Sprintf("%s%d: (%s) %s  %s\n", prefix, k+1, t.ClusterURL, t.ClusterName, loged)
		}

	}
	return formatted
}

//PrintToTable prints outs a formatted table
func (l *List) PrintToTable() {
	formatted := ""
	var data = [][]string{}

	for k, t := range *l {
		current := "   "
		logged := "   "
		if t.Current {
			current = util.IfWindowsElse(" ✔ ", " ✔ ")
		}
		if t.ClusterTICKET != "" {
			logged = " ✓ "
		}
		if t.Visible {
			//prefix = " 👁️ "
			// Adjust the item number k to print numbers starting from 1 instead of 0
			formatted += fmt.Sprintf("%s%d: (%s) %s  %s\n", current, k+1, t.ClusterURL, t.ClusterName, logged)
			data = append(data, []string{fmt.Sprintf("%d", k+1), t.ClusterName, t.ClusterURL, current, logged})
		}

	}
	util.PrintToTable([]string{"#", "Name", "URL", "Current", "LoggedIn"}, data)
}
