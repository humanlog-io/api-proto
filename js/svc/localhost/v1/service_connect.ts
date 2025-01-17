// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/localhost/v1/service.proto (package svc.localhost.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PingRequest, PingResponse } from "./service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service svc.localhost.v1.LocalhostService
 */
export const LocalhostService = {
  typeName: "svc.localhost.v1.LocalhostService",
  methods: {
    /**
     * @generated from rpc svc.localhost.v1.LocalhostService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

