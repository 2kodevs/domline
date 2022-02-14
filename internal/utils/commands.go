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
package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"text/template"

	"github.com/2kodevs/domline/configs"
	log "github.com/sirupsen/logrus"
)

/*
RunCmd :
Executes a command and returns the output.

params :-
a. cmd string
The command to be executed.
b. bool output
True if the output is to be returned.
b. args []string
The arguments to be passed to the command.

returns :-
a. string
The output of the command.
b. error
Error if any
*/
func RunCmd(cmd string, output bool, args ...string) (string, error) {
	log.Info(configs.RunningCMD)
	fullCmd := fmt.Sprintf(cmd, args)
	tmp, err := template.New("script").Parse(fullCmd)
	if err != nil {
		return "", err
	}

	return ExecuteScript(tmp, output, struct{}{})
}

/*
executeScript :
Execute the script in the given template.

params :-
a. tmp *template.Template
Script template
b. output bool
True if the output is to be returned. False to print the output to stdout and stderr

returns :-
a. string
The output of the script.
b. error
Error if any
*/
func ExecuteScript(tmp *template.Template, output bool, data interface{}) (out string, err error) {
	var script, combinedOut bytes.Buffer
	if err = tmp.Execute(&script, data); err != nil {
		return
	}

	cmd := exec.Command("bash")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	wait := sync.WaitGroup{}

	errChans := make([]<-chan error, 0)
	errChans = append(errChans, goCopy(&wait, stdin, &script, true))
	if output {
		cmd.Stdout = &combinedOut
		cmd.Stderr = &combinedOut
	} else {
		errChans = append(errChans, goCopy(&wait, os.Stdout, stdout, false))
		errChans = append(errChans, goCopy(&wait, os.Stderr, stderr, false))
	}

	if err = cmd.Start(); err != nil {
		return
	}

	for _, errChan := range errChans {
		err = <-errChan
		if err != nil {
			return
		}
	}

	wait.Wait()

	if err = cmd.Wait(); err != nil {
		return
	}

	if output {
		out = combinedOut.String()
	}

	return out, nil
}

/*
goCopy :
Copy the content from reader(src) to writer(dst).

params :-
a. wait *sync.WaitGroup
Wait group to wait for copying to finish
b. dst io.Writer
Destination to write to
c. src io.Reader
Source to read from
d. isStdin bool
True if the destination is stdin, false otherwise

returns :-
a. chan error
Channel to where error will be sent
*/
func goCopy(wait *sync.WaitGroup, dst io.WriteCloser, src io.Reader, isStdin bool) <-chan error {
	errChan := make(chan error)
	wait.Add(1)
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			errChan <- err
			return
		}
		if isStdin {
			if err := dst.Close(); err != nil {
				errChan <- err
				return
			}
		}
		close(errChan)
		wait.Done()
	}()
	return errChan
}
