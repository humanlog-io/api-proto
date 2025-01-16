// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/query.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Obj, Scalar, Table, Val, VarType } from "./types_pb";
import { LogQuery } from "./logquery_pb";
import { IngestedLogEvent } from "./logevent_pb";

/**
 * @generated from message types.v1.Data
 */
export class Data extends Message<Data> {
  /**
   * @generated from oneof types.v1.Data.shape
   */
  shape: {
    /**
     * SubQueries need to be further queried to obtain actual data.
     *
     * @generated from field: types.v1.Data.SubQueries subqueries = 101;
     */
    value: Data_SubQueries;
    case: "subqueries";
  } | {
    /**
     * tabular is data that has a table header, followed by rows that match the
     * header
     *
     * @generated from field: types.v1.Tabular tabular = 201;
     */
    value: Tabular;
    case: "tabular";
  } | {
    /**
     * single_value is a single value
     *
     * @generated from field: types.v1.Val single_value = 301;
     */
    value: Val;
    case: "singleValue";
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
    { no: 101, name: "subqueries", kind: "message", T: Data_SubQueries, oneof: "shape" },
    { no: 201, name: "tabular", kind: "message", T: Tabular, oneof: "shape" },
    { no: 301, name: "single_value", kind: "message", T: Val, oneof: "shape" },
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
 * @generated from message types.v1.Data.SubQueries
 */
export class Data_SubQueries extends Message<Data_SubQueries> {
  /**
   * @generated from field: repeated types.v1.LogQuery queries = 1;
   */
  queries: LogQuery[] = [];

  constructor(data?: PartialMessage<Data_SubQueries>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Data.SubQueries";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "queries", kind: "message", T: LogQuery, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Data_SubQueries {
    return new Data_SubQueries().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Data_SubQueries {
    return new Data_SubQueries().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Data_SubQueries {
    return new Data_SubQueries().fromJsonString(jsonString, options);
  }

  static equals(a: Data_SubQueries | PlainMessage<Data_SubQueries> | undefined, b: Data_SubQueries | PlainMessage<Data_SubQueries> | undefined): boolean {
    return proto3.util.equals(Data_SubQueries, a, b);
  }
}

/**
 * @generated from message types.v1.Tabular
 */
export class Tabular extends Message<Tabular> {
  /**
   * @generated from oneof types.v1.Tabular.shape
   */
  shape: {
    /**
     * @generated from field: types.v1.LogEvents log_events = 201;
     */
    value: LogEvents;
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
    { no: 201, name: "log_events", kind: "message", T: LogEvents, oneof: "shape" },
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
   * scalars will contain only this type of value
   *
   * @generated from field: types.v1.VarType type = 2;
   */
  type?: VarType;

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
    { no: 2, name: "type", kind: "message", T: VarType },
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
 * @generated from message types.v1.LogEvents
 */
export class LogEvents extends Message<LogEvents> {
  /**
   * @generated from field: repeated types.v1.IngestedLogEvent events = 1;
   */
  events: IngestedLogEvent[] = [];

  constructor(data?: PartialMessage<LogEvents>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.LogEvents";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "events", kind: "message", T: IngestedLogEvent, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogEvents {
    return new LogEvents().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogEvents {
    return new LogEvents().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogEvents {
    return new LogEvents().fromJsonString(jsonString, options);
  }

  static equals(a: LogEvents | PlainMessage<LogEvents> | undefined, b: LogEvents | PlainMessage<LogEvents> | undefined): boolean {
    return proto3.util.equals(LogEvents, a, b);
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
  type?: VarType;

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
    { no: 2, name: "type", kind: "message", T: VarType },
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

