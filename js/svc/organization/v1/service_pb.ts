// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file svc/organization/v1/service.proto (package svc.organization.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Account } from "../../../types/v1/account_pb";
import { Cursor } from "../../../types/v1/cursor_pb";
import { User } from "../../../types/v1/user_pb";

/**
 * @generated from message svc.organization.v1.CreateAccountRequest
 */
export class CreateAccountRequest extends Message<CreateAccountRequest> {
  /**
   * @generated from field: int64 organization_id = 1;
   */
  organizationId = protoInt64.zero;

  /**
   * @generated from field: string account_name = 2;
   */
  accountName = "";

  constructor(data?: PartialMessage<CreateAccountRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.CreateAccountRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "organization_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "account_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAccountRequest {
    return new CreateAccountRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAccountRequest {
    return new CreateAccountRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAccountRequest {
    return new CreateAccountRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateAccountRequest | PlainMessage<CreateAccountRequest> | undefined, b: CreateAccountRequest | PlainMessage<CreateAccountRequest> | undefined): boolean {
    return proto3.util.equals(CreateAccountRequest, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.CreateAccountResponse
 */
export class CreateAccountResponse extends Message<CreateAccountResponse> {
  /**
   * @generated from field: types.v1.Account account = 1;
   */
  account?: Account;

  constructor(data?: PartialMessage<CreateAccountResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.CreateAccountResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "message", T: Account },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAccountResponse {
    return new CreateAccountResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAccountResponse {
    return new CreateAccountResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAccountResponse {
    return new CreateAccountResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateAccountResponse | PlainMessage<CreateAccountResponse> | undefined, b: CreateAccountResponse | PlainMessage<CreateAccountResponse> | undefined): boolean {
    return proto3.util.equals(CreateAccountResponse, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.ListAccountRequest
 */
export class ListAccountRequest extends Message<ListAccountRequest> {
  /**
   * @generated from field: types.v1.Cursor cursor = 1;
   */
  cursor?: Cursor;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  /**
   * @generated from field: int64 organization_id = 3;
   */
  organizationId = protoInt64.zero;

  constructor(data?: PartialMessage<ListAccountRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.ListAccountRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cursor", kind: "message", T: Cursor },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "organization_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAccountRequest {
    return new ListAccountRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAccountRequest {
    return new ListAccountRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAccountRequest {
    return new ListAccountRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListAccountRequest | PlainMessage<ListAccountRequest> | undefined, b: ListAccountRequest | PlainMessage<ListAccountRequest> | undefined): boolean {
    return proto3.util.equals(ListAccountRequest, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.ListAccountResponse
 */
export class ListAccountResponse extends Message<ListAccountResponse> {
  /**
   * @generated from field: types.v1.Cursor next = 1;
   */
  next?: Cursor;

  /**
   * @generated from field: repeated svc.organization.v1.ListAccountResponse.ListItem items = 2;
   */
  items: ListAccountResponse_ListItem[] = [];

  constructor(data?: PartialMessage<ListAccountResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.ListAccountResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "message", T: Cursor },
    { no: 2, name: "items", kind: "message", T: ListAccountResponse_ListItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAccountResponse {
    return new ListAccountResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAccountResponse {
    return new ListAccountResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAccountResponse {
    return new ListAccountResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListAccountResponse | PlainMessage<ListAccountResponse> | undefined, b: ListAccountResponse | PlainMessage<ListAccountResponse> | undefined): boolean {
    return proto3.util.equals(ListAccountResponse, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.ListAccountResponse.ListItem
 */
export class ListAccountResponse_ListItem extends Message<ListAccountResponse_ListItem> {
  /**
   * @generated from field: types.v1.Account account = 1;
   */
  account?: Account;

  constructor(data?: PartialMessage<ListAccountResponse_ListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.ListAccountResponse.ListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "message", T: Account },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAccountResponse_ListItem {
    return new ListAccountResponse_ListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAccountResponse_ListItem {
    return new ListAccountResponse_ListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAccountResponse_ListItem {
    return new ListAccountResponse_ListItem().fromJsonString(jsonString, options);
  }

  static equals(a: ListAccountResponse_ListItem | PlainMessage<ListAccountResponse_ListItem> | undefined, b: ListAccountResponse_ListItem | PlainMessage<ListAccountResponse_ListItem> | undefined): boolean {
    return proto3.util.equals(ListAccountResponse_ListItem, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.ListUserRequest
 */
export class ListUserRequest extends Message<ListUserRequest> {
  /**
   * @generated from field: types.v1.Cursor cursor = 1;
   */
  cursor?: Cursor;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  /**
   * @generated from field: int64 organization_id = 3;
   */
  organizationId = protoInt64.zero;

  constructor(data?: PartialMessage<ListUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.ListUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cursor", kind: "message", T: Cursor },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "organization_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUserRequest {
    return new ListUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUserRequest {
    return new ListUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUserRequest {
    return new ListUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListUserRequest | PlainMessage<ListUserRequest> | undefined, b: ListUserRequest | PlainMessage<ListUserRequest> | undefined): boolean {
    return proto3.util.equals(ListUserRequest, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.ListUserResponse
 */
export class ListUserResponse extends Message<ListUserResponse> {
  /**
   * @generated from field: types.v1.Cursor next = 1;
   */
  next?: Cursor;

  /**
   * @generated from field: repeated svc.organization.v1.ListUserResponse.ListItem items = 2;
   */
  items: ListUserResponse_ListItem[] = [];

  constructor(data?: PartialMessage<ListUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.ListUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "message", T: Cursor },
    { no: 2, name: "items", kind: "message", T: ListUserResponse_ListItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUserResponse {
    return new ListUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUserResponse {
    return new ListUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUserResponse {
    return new ListUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListUserResponse | PlainMessage<ListUserResponse> | undefined, b: ListUserResponse | PlainMessage<ListUserResponse> | undefined): boolean {
    return proto3.util.equals(ListUserResponse, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.ListUserResponse.ListItem
 */
export class ListUserResponse_ListItem extends Message<ListUserResponse_ListItem> {
  /**
   * @generated from field: types.v1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<ListUserResponse_ListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.ListUserResponse.ListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUserResponse_ListItem {
    return new ListUserResponse_ListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUserResponse_ListItem {
    return new ListUserResponse_ListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUserResponse_ListItem {
    return new ListUserResponse_ListItem().fromJsonString(jsonString, options);
  }

  static equals(a: ListUserResponse_ListItem | PlainMessage<ListUserResponse_ListItem> | undefined, b: ListUserResponse_ListItem | PlainMessage<ListUserResponse_ListItem> | undefined): boolean {
    return proto3.util.equals(ListUserResponse_ListItem, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.InviteUserRequest
 */
export class InviteUserRequest extends Message<InviteUserRequest> {
  /**
   * @generated from field: int64 organization_id = 1;
   */
  organizationId = protoInt64.zero;

  /**
   * @generated from field: int64 user_email = 2;
   */
  userEmail = protoInt64.zero;

  constructor(data?: PartialMessage<InviteUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.InviteUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "organization_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "user_email", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InviteUserRequest {
    return new InviteUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InviteUserRequest {
    return new InviteUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InviteUserRequest {
    return new InviteUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: InviteUserRequest | PlainMessage<InviteUserRequest> | undefined, b: InviteUserRequest | PlainMessage<InviteUserRequest> | undefined): boolean {
    return proto3.util.equals(InviteUserRequest, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.InviteUserResponse
 */
export class InviteUserResponse extends Message<InviteUserResponse> {
  constructor(data?: PartialMessage<InviteUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.InviteUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InviteUserResponse {
    return new InviteUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InviteUserResponse {
    return new InviteUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InviteUserResponse {
    return new InviteUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: InviteUserResponse | PlainMessage<InviteUserResponse> | undefined, b: InviteUserResponse | PlainMessage<InviteUserResponse> | undefined): boolean {
    return proto3.util.equals(InviteUserResponse, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.RevokeUserRequest
 */
export class RevokeUserRequest extends Message<RevokeUserRequest> {
  /**
   * @generated from field: int64 organization_id = 1;
   */
  organizationId = protoInt64.zero;

  /**
   * @generated from field: int64 user_id = 2;
   */
  userId = protoInt64.zero;

  constructor(data?: PartialMessage<RevokeUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.RevokeUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "organization_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "user_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RevokeUserRequest {
    return new RevokeUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RevokeUserRequest {
    return new RevokeUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RevokeUserRequest {
    return new RevokeUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RevokeUserRequest | PlainMessage<RevokeUserRequest> | undefined, b: RevokeUserRequest | PlainMessage<RevokeUserRequest> | undefined): boolean {
    return proto3.util.equals(RevokeUserRequest, a, b);
  }
}

/**
 * @generated from message svc.organization.v1.RevokeUserResponse
 */
export class RevokeUserResponse extends Message<RevokeUserResponse> {
  constructor(data?: PartialMessage<RevokeUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.organization.v1.RevokeUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RevokeUserResponse {
    return new RevokeUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RevokeUserResponse {
    return new RevokeUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RevokeUserResponse {
    return new RevokeUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RevokeUserResponse | PlainMessage<RevokeUserResponse> | undefined, b: RevokeUserResponse | PlainMessage<RevokeUserResponse> | undefined): boolean {
    return proto3.util.equals(RevokeUserResponse, a, b);
  }
}
