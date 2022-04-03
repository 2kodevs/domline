package utils

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/2kodevs/domline/configs"
	"github.com/2kodevs/domline/templates"
	log "github.com/sirupsen/logrus"
)

func Workflow(players configs.Players, tmplName string) error {
	log.Debugf("Players: %+v", players)
	if len(players) == 0 {
		return fmt.Errorf(configs.NotPlayersFoundError)
	}

	for _, p := range players {
		splitted := strings.Split(p.URL, "/")
		id := splitted[len(splitted)-1]

		log.Debugf(configs.Player, p)

		checkData := templates.CheckData{
			Repo:        p.URL,
			Branch:      p.Branch,
			Dir:         id,
			Tag:         id,
			ManagerRepo: os.Getenv(configs.ManagerEnvV),
		}

		rawBase, err := templates.Templates.ReadFile("build_base.tmpl")
		if err != nil {
			return err
		}

		tmp, err := template.New("build").Parse(string(rawBase))
		if err != nil {
			return err
		}

		rawScript, err := templates.Templates.ReadFile(tmplName)
		if err != nil {
			return err
		}
		_, err = tmp.Parse(string(rawScript))
		if err != nil {
			return err
		}

		script := Script{
			Tmp:       tmp,
			GetOutput: false,
			Data:      checkData,
		}
		if _, err := ExecuteScript(script); err != nil {
			return err
		}

		// Uncomment for printing the embedded script
		// err = tmp.Execute(os.Stdout, checkData)
		// if err != nil {
		// 	return fmt.Errorf("Printing file error: ", err)
		// }
	}

	return nil
}
