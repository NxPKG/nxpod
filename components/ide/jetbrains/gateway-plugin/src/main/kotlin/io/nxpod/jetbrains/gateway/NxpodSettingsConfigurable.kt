// Copyright (c) 2022 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.gitpod.jetbrains.gateway

import com.intellij.openapi.components.service
import com.intellij.openapi.options.BoundConfigurable
import com.intellij.openapi.ui.DialogPanel
import com.intellij.openapi.ui.ValidationInfo
import com.intellij.ui.components.JBTextField
import com.intellij.ui.dsl.builder.*
import com.intellij.ui.layout.ValidationInfoBuilder

class NxpodSettingsConfigurable : BoundConfigurable("Nxpod") {

    override fun createPanel(): DialogPanel {
        val state = service<NxpodSettingsState>()
        return panel {
            row {
                textField()
                    .label("Nxpod Host:", LabelPosition.LEFT)
                    .align(Align.FILL)
                    .bindText(state::gitpodHost)
                    .validationOnApply(::validateNxpodHost)
                    .validationOnInput(::validateNxpodHost)
            }
            row {
                checkBox("Force SSH over HTTP tunnel")
                    .bindSelected(state::forceHttpTunnel)
                    .comment("Helpful if you are behind a firewall/proxy that blocks SSH or " +
                            "have complicated SSH setup (bastions, proxy jumps, etc.)")
            }

        }
    }

    private fun validateNxpodHost(
        builder: ValidationInfoBuilder,
        gitpodHost: JBTextField
    ): ValidationInfo? {
        return builder.run {
            if (gitpodHost.text.isBlank()) {
                return@run error("may not be empty")
            }
            return@run null
        }
    }
}
