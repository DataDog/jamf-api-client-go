package classic

import "encoding/xml"

// ScriptsList holds a list of all the scripts available in Jamf
type ScriptsList struct {
	Scripts []BasicScriptInfo `json:"scripts"`
}

// BasicScriptInfo holds the most basic information about the scripts available in Jamf
type BasicScriptInfo struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name"`
}

// Script holds the details to a specific script queried by ID
type Script struct {
	Content *ScriptContents `json:"script" xml:"script,omitempty"`
}

// ScriptContents holds the inner content of a script in Jamf
type ScriptContents struct {
	XMLName         xml.Name    `json:"-" xml:"script,omitempty"`
	ID              int         `json:"id,omitempty" xml:"id,omitempty"`
	Name            string      `json:"name" xml:"name,omitempty"`
	Category        string      `json:"category" xml:"category,omitempty"`
	Filename        string      `json:"filename" xml:"filename,omitempty"`
	Info            string      `json:"info" xml:"info,omitempty"`
	Notes           string      `json:"notes" xml:"notes,omitempty"`
	Priority        string      `json:"priority" xml:"priority,omitempty"`
	Parameters      interface{} `json:"parameters" xml:"parameters,omitempty"`
	Requirements    string      `json:"os_requirements" xml:"os_requirements,omitempty"`
	Contents        string      `json:"script_contents" xml:"script_contents,omitempty"`
	EncodedContents string      `json:"script_contents_encoded" xml:"script_contents_encoded,omitempty"`
}

// ParametersList holds the potential parameters that can be specified for a script in Jamf
type ParametersList struct {
	Parameter4  string `json:"parameter4" xml:"parameter4"`
	Parameter5  string `json:"parameter5" xml:"parameter5"`
	Parameter6  string `json:"parameter6" xml:"parameter6"`
	Parameter7  string `json:"parameter7" xml:"parameter7"`
	Parameter8  string `json:"parameter8" xml:"parameter8"`
	Parameter9  string `json:"parameter9," xml:"parameter9"`
	Parameter10 string `json:"parameter10" xml:"parameter10"`
	Parameter11 string `json:"parameter11" xml:"parameter11"`
}
