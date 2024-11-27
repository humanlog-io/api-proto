// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/feature/v1/service.proto (package svc.feature.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { HasFeatureRequest, HasFeatureResponse } from "./service_pb";

/**
 * @generated from rpc svc.feature.v1.FeatureService.HasFeature
 */
export const hasFeature = {
  localName: "hasFeature",
  name: "HasFeature",
  kind: MethodKind.Unary,
  I: HasFeatureRequest,
  O: HasFeatureResponse,
  service: {
    typeName: "svc.feature.v1.FeatureService"
  }
} as const;
