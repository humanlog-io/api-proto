// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/product.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Price } from "./price_pb";

/**
 * @generated from message types.v1.Product
 */
export class Product extends Message<Product> {
  /**
   * @generated from field: google.protobuf.Timestamp created_at = 100;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 101;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: string stripe_id = 1;
   */
  stripeId = "";

  /**
   * @generated from field: types.v1.Price default_price = 2;
   */
  defaultPrice?: Price;

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string description = 4;
   */
  description = "";

  /**
   * @generated from field: repeated types.v1.Product.MarketingFeature marketing_features = 5;
   */
  marketingFeatures: Product_MarketingFeature[] = [];

  /**
   * @generated from field: bool is_new_env = 6;
   */
  isNewEnv = false;

  /**
   * @generated from field: string cta_link = 7;
   */
  ctaLink = "";

  constructor(data?: PartialMessage<Product>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Product";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 100, name: "created_at", kind: "message", T: Timestamp },
    { no: 101, name: "updated_at", kind: "message", T: Timestamp },
    { no: 1, name: "stripe_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "default_price", kind: "message", T: Price },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "marketing_features", kind: "message", T: Product_MarketingFeature, repeated: true },
    { no: 6, name: "is_new_env", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "cta_link", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Product {
    return new Product().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Product {
    return new Product().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Product {
    return new Product().fromJsonString(jsonString, options);
  }

  static equals(a: Product | PlainMessage<Product> | undefined, b: Product | PlainMessage<Product> | undefined): boolean {
    return proto3.util.equals(Product, a, b);
  }
}

/**
 * @generated from message types.v1.Product.MarketingFeature
 */
export class Product_MarketingFeature extends Message<Product_MarketingFeature> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<Product_MarketingFeature>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Product.MarketingFeature";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Product_MarketingFeature {
    return new Product_MarketingFeature().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Product_MarketingFeature {
    return new Product_MarketingFeature().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Product_MarketingFeature {
    return new Product_MarketingFeature().fromJsonString(jsonString, options);
  }

  static equals(a: Product_MarketingFeature | PlainMessage<Product_MarketingFeature> | undefined, b: Product_MarketingFeature | PlainMessage<Product_MarketingFeature> | undefined): boolean {
    return proto3.util.equals(Product_MarketingFeature, a, b);
  }
}

