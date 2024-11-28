// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/environment/v1/service.proto (package svc.environment.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { ListMachineRequest, ListMachineResponse } from "./service_pb";

/**
 * @generated from rpc svc.environment.v1.EnvironmentService.ListMachine
 */
export const listMachine = {
  localName: "listMachine",
  name: "ListMachine",
  kind: MethodKind.Unary,
  I: ListMachineRequest,
  O: ListMachineResponse,
  service: {
    typeName: "svc.environment.v1.EnvironmentService"
  }
} as const;