package utils

import "text/template"

// Script : Struct Represents a script to be executed
type Script struct {
	// Tmp : script template
	Tmp *template.Template
	// getOutput : True to get output of the script
	GetOutput bool
	// Data: template data object
	Data interface{}
}
