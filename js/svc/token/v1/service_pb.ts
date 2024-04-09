// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file svc/token/v1/service.proto (package svc.token.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";
import { AccountRole, AccountToken } from "../../../types/v1/account_token_pb";
import { Cursor } from "../../../types/v1/cursor_pb";

/**
 * @generated from message svc.token.v1.GenerateAccountTokenRequest
 */
export class GenerateAccountTokenRequest extends Message<GenerateAccountTokenRequest> {
  /**
   * @generated from field: int64 account_id = 1;
   */
  accountId = protoInt64.zero;

  /**
   * @generated from field: google.protobuf.Timestamp expires_at = 2;
   */
  expiresAt?: Timestamp;

  /**
   * @generated from field: repeated types.v1.AccountRole roles = 3;
   */
  roles: AccountRole[] = [];

  constructor(data?: PartialMessage<GenerateAccountTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.GenerateAccountTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "expires_at", kind: "message", T: Timestamp },
    { no: 3, name: "roles", kind: "enum", T: proto3.getEnumType(AccountRole), repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenerateAccountTokenRequest {
    return new GenerateAccountTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenerateAccountTokenRequest {
    return new GenerateAccountTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenerateAccountTokenRequest {
    return new GenerateAccountTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GenerateAccountTokenRequest | PlainMessage<GenerateAccountTokenRequest> | undefined, b: GenerateAccountTokenRequest | PlainMessage<GenerateAccountTokenRequest> | undefined): boolean {
    return proto3.util.equals(GenerateAccountTokenRequest, a, b);
  }
}

/**
 * @generated from message svc.token.v1.GenerateAccountTokenResponse
 */
export class GenerateAccountTokenResponse extends Message<GenerateAccountTokenResponse> {
  /**
   * @generated from field: types.v1.AccountToken token = 1;
   */
  token?: AccountToken;

  constructor(data?: PartialMessage<GenerateAccountTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.GenerateAccountTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "message", T: AccountToken },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenerateAccountTokenResponse {
    return new GenerateAccountTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenerateAccountTokenResponse {
    return new GenerateAccountTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenerateAccountTokenResponse {
    return new GenerateAccountTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GenerateAccountTokenResponse | PlainMessage<GenerateAccountTokenResponse> | undefined, b: GenerateAccountTokenResponse | PlainMessage<GenerateAccountTokenResponse> | undefined): boolean {
    return proto3.util.equals(GenerateAccountTokenResponse, a, b);
  }
}

/**
 * @generated from message svc.token.v1.RevokeAccountTokenRequest
 */
export class RevokeAccountTokenRequest extends Message<RevokeAccountTokenRequest> {
  /**
   * @generated from field: int64 account_id = 1;
   */
  accountId = protoInt64.zero;

  /**
   * @generated from field: int64 token_id = 2;
   */
  tokenId = protoInt64.zero;

  constructor(data?: PartialMessage<RevokeAccountTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.RevokeAccountTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "token_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RevokeAccountTokenRequest {
    return new RevokeAccountTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RevokeAccountTokenRequest {
    return new RevokeAccountTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RevokeAccountTokenRequest {
    return new RevokeAccountTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RevokeAccountTokenRequest | PlainMessage<RevokeAccountTokenRequest> | undefined, b: RevokeAccountTokenRequest | PlainMessage<RevokeAccountTokenRequest> | undefined): boolean {
    return proto3.util.equals(RevokeAccountTokenRequest, a, b);
  }
}

/**
 * @generated from message svc.token.v1.RevokeAccountTokenResponse
 */
export class RevokeAccountTokenResponse extends Message<RevokeAccountTokenResponse> {
  constructor(data?: PartialMessage<RevokeAccountTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.RevokeAccountTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RevokeAccountTokenResponse {
    return new RevokeAccountTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RevokeAccountTokenResponse {
    return new RevokeAccountTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RevokeAccountTokenResponse {
    return new RevokeAccountTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RevokeAccountTokenResponse | PlainMessage<RevokeAccountTokenResponse> | undefined, b: RevokeAccountTokenResponse | PlainMessage<RevokeAccountTokenResponse> | undefined): boolean {
    return proto3.util.equals(RevokeAccountTokenResponse, a, b);
  }
}

/**
 * @generated from message svc.token.v1.GetAccountTokenRequest
 */
export class GetAccountTokenRequest extends Message<GetAccountTokenRequest> {
  /**
   * @generated from field: int64 account_id = 1;
   */
  accountId = protoInt64.zero;

  /**
   * @generated from field: int64 token_id = 2;
   */
  tokenId = protoInt64.zero;

  constructor(data?: PartialMessage<GetAccountTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.GetAccountTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "token_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAccountTokenRequest {
    return new GetAccountTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAccountTokenRequest {
    return new GetAccountTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAccountTokenRequest {
    return new GetAccountTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetAccountTokenRequest | PlainMessage<GetAccountTokenRequest> | undefined, b: GetAccountTokenRequest | PlainMessage<GetAccountTokenRequest> | undefined): boolean {
    return proto3.util.equals(GetAccountTokenRequest, a, b);
  }
}

/**
 * @generated from message svc.token.v1.GetAccountTokenResponse
 */
export class GetAccountTokenResponse extends Message<GetAccountTokenResponse> {
  /**
   * @generated from field: types.v1.AccountToken token = 1;
   */
  token?: AccountToken;

  constructor(data?: PartialMessage<GetAccountTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.GetAccountTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "message", T: AccountToken },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAccountTokenResponse {
    return new GetAccountTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAccountTokenResponse {
    return new GetAccountTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAccountTokenResponse {
    return new GetAccountTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetAccountTokenResponse | PlainMessage<GetAccountTokenResponse> | undefined, b: GetAccountTokenResponse | PlainMessage<GetAccountTokenResponse> | undefined): boolean {
    return proto3.util.equals(GetAccountTokenResponse, a, b);
  }
}

/**
 * @generated from message svc.token.v1.ListAccountTokenRequest
 */
export class ListAccountTokenRequest extends Message<ListAccountTokenRequest> {
  /**
   * @generated from field: types.v1.Cursor cursor = 1;
   */
  cursor?: Cursor;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  /**
   * @generated from field: int64 account_id = 3;
   */
  accountId = protoInt64.zero;

  constructor(data?: PartialMessage<ListAccountTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.ListAccountTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cursor", kind: "message", T: Cursor },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "account_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAccountTokenRequest {
    return new ListAccountTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAccountTokenRequest {
    return new ListAccountTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAccountTokenRequest {
    return new ListAccountTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListAccountTokenRequest | PlainMessage<ListAccountTokenRequest> | undefined, b: ListAccountTokenRequest | PlainMessage<ListAccountTokenRequest> | undefined): boolean {
    return proto3.util.equals(ListAccountTokenRequest, a, b);
  }
}

/**
 * @generated from message svc.token.v1.ListAccountTokenResponse
 */
export class ListAccountTokenResponse extends Message<ListAccountTokenResponse> {
  /**
   * @generated from field: types.v1.Cursor next = 1;
   */
  next?: Cursor;

  /**
   * @generated from field: repeated svc.token.v1.ListAccountTokenResponse.ListItem items = 2;
   */
  items: ListAccountTokenResponse_ListItem[] = [];

  constructor(data?: PartialMessage<ListAccountTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.ListAccountTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "message", T: Cursor },
    { no: 2, name: "items", kind: "message", T: ListAccountTokenResponse_ListItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAccountTokenResponse {
    return new ListAccountTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAccountTokenResponse {
    return new ListAccountTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAccountTokenResponse {
    return new ListAccountTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListAccountTokenResponse | PlainMessage<ListAccountTokenResponse> | undefined, b: ListAccountTokenResponse | PlainMessage<ListAccountTokenResponse> | undefined): boolean {
    return proto3.util.equals(ListAccountTokenResponse, a, b);
  }
}

/**
 * @generated from message svc.token.v1.ListAccountTokenResponse.ListItem
 */
export class ListAccountTokenResponse_ListItem extends Message<ListAccountTokenResponse_ListItem> {
  /**
   * @generated from field: types.v1.AccountToken token = 1;
   */
  token?: AccountToken;

  constructor(data?: PartialMessage<ListAccountTokenResponse_ListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.token.v1.ListAccountTokenResponse.ListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "message", T: AccountToken },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAccountTokenResponse_ListItem {
    return new ListAccountTokenResponse_ListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAccountTokenResponse_ListItem {
    return new ListAccountTokenResponse_ListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAccountTokenResponse_ListItem {
    return new ListAccountTokenResponse_ListItem().fromJsonString(jsonString, options);
  }

  static equals(a: ListAccountTokenResponse_ListItem | PlainMessage<ListAccountTokenResponse_ListItem> | undefined, b: ListAccountTokenResponse_ListItem | PlainMessage<ListAccountTokenResponse_ListItem> | undefined): boolean {
    return proto3.util.equals(ListAccountTokenResponse_ListItem, a, b);
  }
}
