// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/user/v1/service.proto (package svc.user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { CreateOrganizationRequest, CreateOrganizationResponse, GetLogoutURLRequest, GetLogoutURLResponse, ListOrganizationRequest, ListOrganizationResponse, WhoamiRequest, WhoamiResponse } from "./service_pb";

/**
 * @generated from rpc svc.user.v1.UserService.Whoami
 */
export const whoami = {
  localName: "whoami",
  name: "Whoami",
  kind: MethodKind.Unary,
  I: WhoamiRequest,
  O: WhoamiResponse,
  service: {
    typeName: "svc.user.v1.UserService"
  }
} as const;

/**
 * @generated from rpc svc.user.v1.UserService.GetLogoutURL
 */
export const getLogoutURL = {
  localName: "getLogoutURL",
  name: "GetLogoutURL",
  kind: MethodKind.Unary,
  I: GetLogoutURLRequest,
  O: GetLogoutURLResponse,
  service: {
    typeName: "svc.user.v1.UserService"
  }
} as const;

/**
 * @generated from rpc svc.user.v1.UserService.CreateOrganization
 */
export const createOrganization = {
  localName: "createOrganization",
  name: "CreateOrganization",
  kind: MethodKind.Unary,
  I: CreateOrganizationRequest,
  O: CreateOrganizationResponse,
  service: {
    typeName: "svc.user.v1.UserService"
  }
} as const;

/**
 * @generated from rpc svc.user.v1.UserService.ListOrganization
 */
export const listOrganization = {
  localName: "listOrganization",
  name: "ListOrganization",
  kind: MethodKind.Unary,
  I: ListOrganizationRequest,
  O: ListOrganizationResponse,
  service: {
    typeName: "svc.user.v1.UserService"
  }
} as const;
