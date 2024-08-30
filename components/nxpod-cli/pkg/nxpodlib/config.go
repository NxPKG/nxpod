// Copyright (c) 2020 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package gitpodlib

type NxpodImage struct {
	File    string
	Context string `yaml:"context,omitempty"`
}

type gitpodPort struct {
	Number int32 `yaml:"port"`
}

type gitpodTask struct {
	Command string `yaml:"command,omitempty"`
	Init    string `yaml:"init,omitempty"`
}

type NxpodFile struct {
	Image             interface{}  `yaml:"image,omitempty"`
	Ports             []gitpodPort `yaml:"ports,omitempty"`
	Tasks             []gitpodTask `yaml:"tasks,omitempty"`
	CheckoutLocation  string       `yaml:"checkoutLocation,omitempty"`
	WorkspaceLocation string       `yaml:"workspaceLocation,omitempty"`
}

// SetImageName configures a pre-built docker image by name
func (cfg *NxpodFile) SetImageName(name string) {
	if name == "" {
		return
	}
	cfg.Image = name
}

// SetImage configures a Dockerfile as workspace image
func (cfg *NxpodFile) SetImage(img NxpodImage) {
	cfg.Image = img
}

// AddPort adds a port to the list of exposed ports
func (cfg *NxpodFile) AddPort(port int32) {
	cfg.Ports = append(cfg.Ports, gitpodPort{
		Number: port,
	})
}

// AddTask adds a workspace startup task
func (cfg *NxpodFile) AddTask(task ...string) {
	if len(task) > 1 {
		cfg.Tasks = append(cfg.Tasks, gitpodTask{
			Command: task[0],
			Init:    task[1],
		})
	} else {
		cfg.Tasks = append(cfg.Tasks, gitpodTask{
			Command: task[0],
		})
	}
}
