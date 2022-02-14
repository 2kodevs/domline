/*
Copyright © 2022 2kodevs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"strings"
	"text/template"

	"github.com/2kodevs/domline/configs"
	"github.com/2kodevs/domline/internal/utils"
	"github.com/2kodevs/domline/templates"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	pathToConfigFiles string
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		players, err := configs.GetConfigs(pathToConfigFiles)
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range players {
			splitted := strings.Split(p.URL, "/")
			id := splitted[len(splitted)-1]

			checkData := templates.CheckData{
				Repo:   p.URL,
				Branch: p.Branch,
				Dir:    id,
				Tag:    id,
			}

			rawScript, err := templates.Templates.ReadFile("check.tmpl")
			if err != nil {
				log.Fatal(err)
			}
			tmp, err := template.New("script").Parse(string(rawScript))
			if err != nil {
				return
			}

			if _, err := utils.ExecuteScript(tmp, false, checkData); err != nil {
				log.Fatal(err)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVarP(&pathToConfigFiles, "path", "p", "", "path to config files")
	if err := checkCmd.MarkFlagRequired("path"); err != nil {
		log.Fatal(err)
	}
}
