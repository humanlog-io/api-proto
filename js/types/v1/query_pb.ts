// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/query.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { LogQuery } from "./logquery_pb";
import { Obj, Scalar, Table, VarType } from "./types_pb";
import { LogEventGroup } from "./logevent_pb";

/**
 * @generated from message types.v1.Data
 */
export class Data extends Message<Data> {
  /**
   * Schema schema = 1;
   *
   * @generated from oneof types.v1.Data.shape
   */
  shape: {
    /**
     * subqueries means that the result must be obtained by re-issuing
     * more queries.
     *
     * @generated from field: types.v1.SubQueries subqueries = 201;
     */
    value: SubQueries;
    case: "subqueries";
  } | {
    /**
     * tabular is data that has a table header, followed by rows that match the
     * header
     *
     * @generated from field: types.v1.Tabular tabular = 301;
     */
    value: Tabular;
    case: "tabular";
  } | {
    /**
     * scalar timeseries data is a list of [ts, scalar] where scalar is usually
     * a numeric value
     *
     * @generated from field: types.v1.ScalarTimeseries scalar_timeseries = 401;
     */
    value: ScalarTimeseries;
    case: "scalarTimeseries";
  } | {
    /**
     * vector timeseries data is a list of [ts, [scalar0, ..., scalarN]] where
     * scalar is usually a numeric value
     *
     * @generated from field: types.v1.VectorTimeseries vector_timeseries = 402;
     */
    value: VectorTimeseries;
    case: "vectorTimeseries";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Data>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Data";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 201, name: "subqueries", kind: "message", T: SubQueries, oneof: "shape" },
    { no: 301, name: "tabular", kind: "message", T: Tabular, oneof: "shape" },
    { no: 401, name: "scalar_timeseries", kind: "message", T: ScalarTimeseries, oneof: "shape" },
    { no: 402, name: "vector_timeseries", kind: "message", T: VectorTimeseries, oneof: "shape" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Data {
    return new Data().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Data {
    return new Data().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Data {
    return new Data().fromJsonString(jsonString, options);
  }

  static equals(a: Data | PlainMessage<Data> | undefined, b: Data | PlainMessage<Data> | undefined): boolean {
    return proto3.util.equals(Data, a, b);
  }
}

/**
 * @generated from message types.v1.SubQueries
 */
export class SubQueries extends Message<SubQueries> {
  /**
   * @generated from field: repeated types.v1.LogQuery queries = 1;
   */
  queries: LogQuery[] = [];

  constructor(data?: PartialMessage<SubQueries>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.SubQueries";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "queries", kind: "message", T: LogQuery, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubQueries {
    return new SubQueries().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubQueries {
    return new SubQueries().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubQueries {
    return new SubQueries().fromJsonString(jsonString, options);
  }

  static equals(a: SubQueries | PlainMessage<SubQueries> | undefined, b: SubQueries | PlainMessage<SubQueries> | undefined): boolean {
    return proto3.util.equals(SubQueries, a, b);
  }
}

/**
 * @generated from message types.v1.Tabular
 */
export class Tabular extends Message<Tabular> {
  /**
   * @generated from field: types.v1.Obj context = 1;
   */
  context?: Obj;

  /**
   * @generated from oneof types.v1.Tabular.shape
   */
  shape: {
    /**
     * @generated from field: types.v1.LogEventGroup log_events = 201;
     */
    value: LogEventGroup;
    case: "logEvents";
  } | {
    /**
     * @generated from field: types.v1.Table free_form = 202;
     */
    value: Table;
    case: "freeForm";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Tabular>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Tabular";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "context", kind: "message", T: Obj },
    { no: 201, name: "log_events", kind: "message", T: LogEventGroup, oneof: "shape" },
    { no: 202, name: "free_form", kind: "message", T: Table, oneof: "shape" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Tabular {
    return new Tabular().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Tabular {
    return new Tabular().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Tabular {
    return new Tabular().fromJsonString(jsonString, options);
  }

  static equals(a: Tabular | PlainMessage<Tabular> | undefined, b: Tabular | PlainMessage<Tabular> | undefined): boolean {
    return proto3.util.equals(Tabular, a, b);
  }
}

/**
 * @generated from message types.v1.ScalarTimeseries
 */
export class ScalarTimeseries extends Message<ScalarTimeseries> {
  /**
   * @generated from field: types.v1.Obj context = 1;
   */
  context?: Obj;

  /**
   * scalars will contain only this type of value
   *
   * @generated from field: types.v1.VarType type = 2;
   */
  type = VarType.unknown;

  /**
   * @generated from field: repeated types.v1.ScalarTimestamp scalars = 3;
   */
  scalars: ScalarTimestamp[] = [];

  constructor(data?: PartialMessage<ScalarTimeseries>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ScalarTimeseries";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "context", kind: "message", T: Obj },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(VarType) },
    { no: 3, name: "scalars", kind: "message", T: ScalarTimestamp, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ScalarTimeseries {
    return new ScalarTimeseries().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ScalarTimeseries {
    return new ScalarTimeseries().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ScalarTimeseries {
    return new ScalarTimeseries().fromJsonString(jsonString, options);
  }

  static equals(a: ScalarTimeseries | PlainMessage<ScalarTimeseries> | undefined, b: ScalarTimeseries | PlainMessage<ScalarTimeseries> | undefined): boolean {
    return proto3.util.equals(ScalarTimeseries, a, b);
  }
}

/**
 * @generated from message types.v1.ScalarTimestamp
 */
export class ScalarTimestamp extends Message<ScalarTimestamp> {
  /**
   * @generated from field: google.protobuf.Timestamp ts = 1;
   */
  ts?: Timestamp;

  /**
   * @generated from field: types.v1.Scalar scalar = 2;
   */
  scalar?: Scalar;

  constructor(data?: PartialMessage<ScalarTimestamp>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ScalarTimestamp";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ts", kind: "message", T: Timestamp },
    { no: 2, name: "scalar", kind: "message", T: Scalar },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ScalarTimestamp {
    return new ScalarTimestamp().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ScalarTimestamp {
    return new ScalarTimestamp().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ScalarTimestamp {
    return new ScalarTimestamp().fromJsonString(jsonString, options);
  }

  static equals(a: ScalarTimestamp | PlainMessage<ScalarTimestamp> | undefined, b: ScalarTimestamp | PlainMessage<ScalarTimestamp> | undefined): boolean {
    return proto3.util.equals(ScalarTimestamp, a, b);
  }
}

/**
 * @generated from message types.v1.VectorTimeseries
 */
export class VectorTimeseries extends Message<VectorTimeseries> {
  /**
   * @generated from field: types.v1.Obj context = 1;
   */
  context?: Obj;

  /**
   * vectors will contain only this type of value
   *
   * @generated from field: types.v1.VarType type = 2;
   */
  type = VarType.unknown;

  /**
   * @generated from field: repeated types.v1.VectorTimestamp vectors = 3;
   */
  vectors: VectorTimestamp[] = [];

  constructor(data?: PartialMessage<VectorTimeseries>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.VectorTimeseries";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "context", kind: "message", T: Obj },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(VarType) },
    { no: 3, name: "vectors", kind: "message", T: VectorTimestamp, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VectorTimeseries {
    return new VectorTimeseries().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VectorTimeseries {
    return new VectorTimeseries().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VectorTimeseries {
    return new VectorTimeseries().fromJsonString(jsonString, options);
  }

  static equals(a: VectorTimeseries | PlainMessage<VectorTimeseries> | undefined, b: VectorTimeseries | PlainMessage<VectorTimeseries> | undefined): boolean {
    return proto3.util.equals(VectorTimeseries, a, b);
  }
}

/**
 * @generated from message types.v1.VectorTimestamp
 */
export class VectorTimestamp extends Message<VectorTimestamp> {
  /**
   * @generated from field: google.protobuf.Timestamp ts = 1;
   */
  ts?: Timestamp;

  /**
   * @generated from field: repeated types.v1.Scalar vector = 2;
   */
  vector: Scalar[] = [];

  constructor(data?: PartialMessage<VectorTimestamp>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.VectorTimestamp";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ts", kind: "message", T: Timestamp },
    { no: 2, name: "vector", kind: "message", T: Scalar, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VectorTimestamp {
    return new VectorTimestamp().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VectorTimestamp {
    return new VectorTimestamp().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VectorTimestamp {
    return new VectorTimestamp().fromJsonString(jsonString, options);
  }

  static equals(a: VectorTimestamp | PlainMessage<VectorTimestamp> | undefined, b: VectorTimestamp | PlainMessage<VectorTimestamp> | undefined): boolean {
    return proto3.util.equals(VectorTimestamp, a, b);
  }
}

