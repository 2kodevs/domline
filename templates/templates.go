package templates

import "embed"

//go:embed *
var Templates embed.FS

// ConsensusEnv : Struct Data object to be applied to check template
type CheckData struct {
	Repo   string
	Branch string
	Dir    string
	Tag    string
}
