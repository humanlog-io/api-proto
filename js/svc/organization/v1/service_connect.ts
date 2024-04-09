// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts,import_extension=none"
// @generated from file svc/organization/v1/service.proto (package svc.organization.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateAccountRequest, CreateAccountResponse, InviteUserRequest, InviteUserResponse, ListAccountRequest, ListAccountResponse, ListUserRequest, ListUserResponse, RevokeUserRequest, RevokeUserResponse } from "./service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service svc.organization.v1.OrganizationService
 */
export const OrganizationService = {
  typeName: "svc.organization.v1.OrganizationService",
  methods: {
    /**
     * @generated from rpc svc.organization.v1.OrganizationService.CreateAccount
     */
    createAccount: {
      name: "CreateAccount",
      I: CreateAccountRequest,
      O: CreateAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.organization.v1.OrganizationService.ListAccount
     */
    listAccount: {
      name: "ListAccount",
      I: ListAccountRequest,
      O: ListAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.organization.v1.OrganizationService.ListUser
     */
    listUser: {
      name: "ListUser",
      I: ListUserRequest,
      O: ListUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.organization.v1.OrganizationService.InviteUser
     */
    inviteUser: {
      name: "InviteUser",
      I: InviteUserRequest,
      O: InviteUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.organization.v1.OrganizationService.RevokeUser
     */
    revokeUser: {
      name: "RevokeUser",
      I: RevokeUserRequest,
      O: RevokeUserResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
