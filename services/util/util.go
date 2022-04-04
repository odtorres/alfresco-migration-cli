package util

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//GetParams get parameters size from Terminal
func GetParams(size int, r io.Reader, args ...string) ([]string, error) {
	if len(args) > 0 {
		return args, nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return nil, err
	}
	if len(s.Text()) == 0 {
		return nil, fmt.Errorf("Command cannot be blank")
	}

	resp := strings.Split(s.Text(), " ")
	if len(resp) < size {
		return nil, fmt.Errorf("Command needs %d arguments", size)
	}
	return resp, nil
}

//PrintIdentJson []byte to ident json
func PrintIdentJson(jsonBytes []byte) {
	var f interface{}
	err := json.Unmarshal(jsonBytes, &f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var out bytes.Buffer
	json.Indent(&out, jsonBytes, "", "\t")
	out.WriteTo(os.Stdout)
}

//PrintToTable print Table
func PrintToTable(title []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(title)
	//table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

//PrintToTableFooter print Table
func PrintToTableFooter(title []string, footer []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(title)
	table.SetFooter(footer)
	//table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

/// OS

//IsWindows returns true if windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

//IsMac returns true if MacOs
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

//IsLinux returns true if linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

//IfWindowsElse returns first parameter if windows otherwise the second parameter
func IfWindowsElse(textWindows, textLinux string) string {
	if IsWindows() {
		return textWindows
	}
	return textLinux
}
