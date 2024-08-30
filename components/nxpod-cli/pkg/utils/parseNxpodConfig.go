// Copyright (c) 2023 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package utils

import (
	"errors"
	"os"
	"path/filepath"

	gitpod "github.com/nxpkg/nxpod/gitpod-protocol"
	yaml "gopkg.in/yaml.v2"
)

func ParseNxpodConfig(repoRoot string) (*gitpod.NxpodConfig, error) {
	if repoRoot == "" {
		return nil, errors.New("repoRoot is empty")
	}
	data, err := os.ReadFile(filepath.Join(repoRoot, ".gitpod.yml"))
	if err != nil {
		// .gitpod.yml not exist is ok
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, errors.New("read .gitpod.yml file failed: " + err.Error())
	}
	var config *gitpod.NxpodConfig
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.New("unmarshal .gitpod.yml file failed" + err.Error())
	}
	return config, nil
}
