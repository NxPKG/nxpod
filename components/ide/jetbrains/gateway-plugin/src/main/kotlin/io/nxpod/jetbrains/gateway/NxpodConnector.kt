// Copyright (c) 2022 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.gitpod.jetbrains.gateway

import com.jetbrains.gateway.api.GatewayConnector
import com.jetbrains.gateway.api.GatewayConnectorDocumentationPage
import com.jetbrains.rd.util.lifetime.Lifetime
import io.gitpod.jetbrains.icons.NxpodIcons
import java.awt.Component

class NxpodConnector : GatewayConnector {
    override val icon = NxpodIcons.Logo

    override fun createView(lifetime: Lifetime) = NxpodConnectorView(lifetime)

    override fun getActionText() = "Connect to Nxpod"

    override fun getDescription() = "Connect to Nxpod workspaces"

    override fun getDocumentationAction() = GatewayConnectorDocumentationPage("https://www.gitpod.io/docs/ides-and-editors/jetbrains-gateway")

    override fun getConnectorId() = "gitpod.connector"

    override fun getRecentConnections(setContentCallback: (Component) -> Unit) = NxpodRecentConnections()

    override fun getTitle() = "Nxpod"

    @Deprecated("Not used", ReplaceWith("null"))
    override fun getTitleAdornment() = null

    override fun initProcedure() {}
}
