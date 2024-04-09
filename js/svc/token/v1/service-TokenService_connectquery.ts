// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/token/v1/service.proto (package svc.token.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { GenerateAccountTokenRequest, GenerateAccountTokenResponse, GetAccountTokenRequest, GetAccountTokenResponse, ListAccountTokenRequest, ListAccountTokenResponse, RevokeAccountTokenRequest, RevokeAccountTokenResponse } from "./service_pb";

/**
 * @generated from rpc svc.token.v1.TokenService.GenerateAccountToken
 */
export const generateAccountToken = {
  localName: "generateAccountToken",
  name: "GenerateAccountToken",
  kind: MethodKind.Unary,
  I: GenerateAccountTokenRequest,
  O: GenerateAccountTokenResponse,
  service: {
    typeName: "svc.token.v1.TokenService"
  }
} as const;

/**
 * @generated from rpc svc.token.v1.TokenService.RevokeAccountToken
 */
export const revokeAccountToken = {
  localName: "revokeAccountToken",
  name: "RevokeAccountToken",
  kind: MethodKind.Unary,
  I: RevokeAccountTokenRequest,
  O: RevokeAccountTokenResponse,
  service: {
    typeName: "svc.token.v1.TokenService"
  }
} as const;

/**
 * @generated from rpc svc.token.v1.TokenService.GetAccountToken
 */
export const getAccountToken = {
  localName: "getAccountToken",
  name: "GetAccountToken",
  kind: MethodKind.Unary,
  I: GetAccountTokenRequest,
  O: GetAccountTokenResponse,
  service: {
    typeName: "svc.token.v1.TokenService"
  }
} as const;

/**
 * @generated from rpc svc.token.v1.TokenService.ListAccountToken
 */
export const listAccountToken = {
  localName: "listAccountToken",
  name: "ListAccountToken",
  kind: MethodKind.Unary,
  I: ListAccountTokenRequest,
  O: ListAccountTokenResponse,
  service: {
    typeName: "svc.token.v1.TokenService"
  }
} as const;
