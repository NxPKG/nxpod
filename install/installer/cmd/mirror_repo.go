// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"

	"github.com/nxpkg/nxpod/installer/pkg/config"
	"github.com/spf13/cobra"
)

// mirrorRepoCmd represents the mirror list command
var mirrorRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Get original image repo for this installer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf(config.NxpodContainerRegistry)
		return nil
	},
}

func init() {
	mirrorCmd.AddCommand(mirrorRepoCmd)
}
