// Copyright (c) 2021 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package minio

import (
	"github.com/nxpkg/nxpod/installer/pkg/common"
)

const Component = "minio"

var Objects = common.CompositeRenderFunc(
	rolebinding,
)
