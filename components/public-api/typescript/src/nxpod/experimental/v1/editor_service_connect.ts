/**
 * Copyright (c) 2024 Nxpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file nxpod/experimental/v1/editor_service.proto (package nxpod.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ListEditorOptionsRequest, ListEditorOptionsResponse } from "./editor_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service nxpod.experimental.v1.EditorService
 */
export const EditorService = {
  typeName: "nxpod.experimental.v1.EditorService",
  methods: {
    /**
     * @generated from rpc nxpod.experimental.v1.EditorService.ListEditorOptions
     */
    listEditorOptions: {
      name: "ListEditorOptions",
      I: ListEditorOptionsRequest,
      O: ListEditorOptionsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
