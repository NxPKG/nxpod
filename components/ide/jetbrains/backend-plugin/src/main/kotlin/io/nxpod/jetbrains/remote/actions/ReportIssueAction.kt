// Copyright (c) 2022 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.gitpod.jetbrains.remote.actions

import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.components.service
import io.gitpod.jetbrains.remote.NxpodManager

class ReportIssueAction : AnAction() {
    private val manager = service<NxpodManager>()

    override fun actionPerformed(event: AnActionEvent) {
       manager.openUrlFromAction("https://github.com/nxpkg/nxpod/issues/new/choose")
    }
}
