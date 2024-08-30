// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: nxpod/v1/configuration.proto
//
package io.nxpod.publicapi.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class ConfigurationServiceClient(
  private val client: ProtocolClientInterface,
) : ConfigurationServiceClientInterface {
  /**
   *  Creates a new configuration.
   */
  override suspend
      fun createConfiguration(request: ConfigurationOuterClass.CreateConfigurationRequest,
      headers: Headers): ResponseMessage<ConfigurationOuterClass.CreateConfigurationResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "nxpod.v1.ConfigurationService/CreateConfiguration",
      io.nxpod.publicapi.v1.ConfigurationOuterClass.CreateConfigurationRequest::class,
      io.nxpod.publicapi.v1.ConfigurationOuterClass.CreateConfigurationResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  Retrieves a configuration.
   */
  override suspend fun getConfiguration(request: ConfigurationOuterClass.GetConfigurationRequest,
      headers: Headers): ResponseMessage<ConfigurationOuterClass.GetConfigurationResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "nxpod.v1.ConfigurationService/GetConfiguration",
      io.nxpod.publicapi.v1.ConfigurationOuterClass.GetConfigurationRequest::class,
      io.nxpod.publicapi.v1.ConfigurationOuterClass.GetConfigurationResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  Lists configurations.
   */
  override suspend
      fun listConfigurations(request: ConfigurationOuterClass.ListConfigurationsRequest,
      headers: Headers): ResponseMessage<ConfigurationOuterClass.ListConfigurationsResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "nxpod.v1.ConfigurationService/ListConfigurations",
      io.nxpod.publicapi.v1.ConfigurationOuterClass.ListConfigurationsRequest::class,
      io.nxpod.publicapi.v1.ConfigurationOuterClass.ListConfigurationsResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  Updates a configuration.
   */
  override suspend
      fun updateConfiguration(request: ConfigurationOuterClass.UpdateConfigurationRequest,
      headers: Headers): ResponseMessage<ConfigurationOuterClass.UpdateConfigurationResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "nxpod.v1.ConfigurationService/UpdateConfiguration",
      io.nxpod.publicapi.v1.ConfigurationOuterClass.UpdateConfigurationRequest::class,
      io.nxpod.publicapi.v1.ConfigurationOuterClass.UpdateConfigurationResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  Deletes a configuration.
   */
  override suspend
      fun deleteConfiguration(request: ConfigurationOuterClass.DeleteConfigurationRequest,
      headers: Headers): ResponseMessage<ConfigurationOuterClass.DeleteConfigurationResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "nxpod.v1.ConfigurationService/DeleteConfiguration",
      io.nxpod.publicapi.v1.ConfigurationOuterClass.DeleteConfigurationRequest::class,
      io.nxpod.publicapi.v1.ConfigurationOuterClass.DeleteConfigurationResponse::class,
      StreamType.UNARY,
    ),
  )

}
