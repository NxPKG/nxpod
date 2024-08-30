// Copyright (c) 2023 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.nxpod.jetbrains.gateway

import com.intellij.openapi.components.Service
import com.jetbrains.gateway.api.CustomConnectionFrameComponentProvider
import com.jetbrains.gateway.api.CustomConnectionFrameContext
import com.jetbrains.gateway.api.GatewayConnectionHandle
import com.jetbrains.gateway.ssh.ClientOverSshTunnelConnector
import com.jetbrains.gateway.ssh.HostTunnelConnector
import com.jetbrains.gateway.thinClientLink.ThinClientHandle
import com.jetbrains.rd.util.lifetime.Lifetime
import io.nxpod.jetbrains.gateway.NxpodConnectionProvider.ConnectParams
import java.net.URI
import javax.swing.JComponent

class NxpodConnectionHandle(
        lifetime: Lifetime,
        private val component: JComponent,
        private val params: ConnectParams
) : GatewayConnectionHandle(lifetime) {
    override fun customComponentProvider(lifetime: Lifetime) = object : CustomConnectionFrameComponentProvider {
        override val closeConfirmationText = "Disconnect from ${getTitle()}?"
        override fun createComponent(context: CustomConnectionFrameContext) = component
    }

    override fun getTitle(): String {
        return params.title
    }

    override fun hideToTrayOnStart(): Boolean {
        return false
    }
}
@Suppress("UnstableApiUsage")
@Service
class NxpodConnectionHandleFactory {
    fun createNxpodConnectionHandle(
            lifetime: Lifetime,
            component: JComponent,
            params: ConnectParams
    ): GatewayConnectionHandle {
        return NxpodConnectionHandle(lifetime, component, params)
    }

    suspend fun connect(lifetime: Lifetime, connector: HostTunnelConnector, tcpJoinLink: URI): ThinClientHandle {
        return ClientOverSshTunnelConnector(
            lifetime,
            connector
        ).connect(tcpJoinLink)
    }
}
