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
	"os"

	"github.com/2kodevs/domline/configs"
	"github.com/2kodevs/domline/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv(configs.ManagerEnvV) == "" {
			log.Fatalf(configs.EmptyEnvVariableError, configs.ManagerEnvV)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		players, err := configs.GetConfigs(pathToConfigFiles)
		if err != nil {
			log.Fatal(err)
		}

		if err = utils.Workflow(players, "random_play.tmpl"); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	playCmd.Flags().StringVarP(&pathToConfigFiles, "path", "p", "", "path to config files")
	if err := playCmd.MarkFlagRequired("path"); err != nil {
		log.Fatal(err)
	}
}
