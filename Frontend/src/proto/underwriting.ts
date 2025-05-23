// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v4.24.4
// source: underwriting.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "loan_decision";

/** CreateLoanDecisionRequest message for creating a loan decision */
export interface CreateLoanDecisionRequest {
  /** Loan application ID */
  loanApplicationId: number;
  /** Decision from AI (approved/rejected) */
  aiDecision: string;
  /** Reasoning for the decision */
  reasoning: string;
}

/** CreateLoanDecisionResponse message for response after creating a loan decision */
export interface CreateLoanDecisionResponse {
  /** Unique ID of the loan decision */
  loanDecisionId: number;
  /** Status message (e.g., "Created successfully") */
  status: string;
}

/** GetLoanDecisionRequest message to request an existing loan decision */
export interface GetLoanDecisionRequest {
  /** ID of the loan decision to fetch */
  loanDecisionId: number;
}

/** GetLoanDecisionResponse message containing the loan decision details */
export interface GetLoanDecisionResponse {
  /** Unique ID of the loan decision */
  loanDecisionId: number;
  /** Associated loan application ID */
  loanApplicationId: number;
  /** Decision from AI (approved/rejected) */
  aiDecision: string;
  /** Reasoning for the decision */
  reasoning: string;
  /** Timestamp of decision creation */
  createdAt: Date | undefined;
}

function createBaseCreateLoanDecisionRequest(): CreateLoanDecisionRequest {
  return { loanApplicationId: 0, aiDecision: "", reasoning: "" };
}

export const CreateLoanDecisionRequest: MessageFns<CreateLoanDecisionRequest> = {
  encode(message: CreateLoanDecisionRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.loanApplicationId !== 0) {
      writer.uint32(8).uint64(message.loanApplicationId);
    }
    if (message.aiDecision !== "") {
      writer.uint32(18).string(message.aiDecision);
    }
    if (message.reasoning !== "") {
      writer.uint32(26).string(message.reasoning);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateLoanDecisionRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateLoanDecisionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.loanApplicationId = longToNumber(reader.uint64());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.aiDecision = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.reasoning = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateLoanDecisionRequest {
    return {
      loanApplicationId: isSet(object.loanApplicationId) ? globalThis.Number(object.loanApplicationId) : 0,
      aiDecision: isSet(object.aiDecision) ? globalThis.String(object.aiDecision) : "",
      reasoning: isSet(object.reasoning) ? globalThis.String(object.reasoning) : "",
    };
  },

  toJSON(message: CreateLoanDecisionRequest): unknown {
    const obj: any = {};
    if (message.loanApplicationId !== 0) {
      obj.loanApplicationId = Math.round(message.loanApplicationId);
    }
    if (message.aiDecision !== "") {
      obj.aiDecision = message.aiDecision;
    }
    if (message.reasoning !== "") {
      obj.reasoning = message.reasoning;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateLoanDecisionRequest>, I>>(base?: I): CreateLoanDecisionRequest {
    return CreateLoanDecisionRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateLoanDecisionRequest>, I>>(object: I): CreateLoanDecisionRequest {
    const message = createBaseCreateLoanDecisionRequest();
    message.loanApplicationId = object.loanApplicationId ?? 0;
    message.aiDecision = object.aiDecision ?? "";
    message.reasoning = object.reasoning ?? "";
    return message;
  },
};

function createBaseCreateLoanDecisionResponse(): CreateLoanDecisionResponse {
  return { loanDecisionId: 0, status: "" };
}

export const CreateLoanDecisionResponse: MessageFns<CreateLoanDecisionResponse> = {
  encode(message: CreateLoanDecisionResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.loanDecisionId !== 0) {
      writer.uint32(8).uint64(message.loanDecisionId);
    }
    if (message.status !== "") {
      writer.uint32(18).string(message.status);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateLoanDecisionResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateLoanDecisionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.loanDecisionId = longToNumber(reader.uint64());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.status = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateLoanDecisionResponse {
    return {
      loanDecisionId: isSet(object.loanDecisionId) ? globalThis.Number(object.loanDecisionId) : 0,
      status: isSet(object.status) ? globalThis.String(object.status) : "",
    };
  },

  toJSON(message: CreateLoanDecisionResponse): unknown {
    const obj: any = {};
    if (message.loanDecisionId !== 0) {
      obj.loanDecisionId = Math.round(message.loanDecisionId);
    }
    if (message.status !== "") {
      obj.status = message.status;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateLoanDecisionResponse>, I>>(base?: I): CreateLoanDecisionResponse {
    return CreateLoanDecisionResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateLoanDecisionResponse>, I>>(object: I): CreateLoanDecisionResponse {
    const message = createBaseCreateLoanDecisionResponse();
    message.loanDecisionId = object.loanDecisionId ?? 0;
    message.status = object.status ?? "";
    return message;
  },
};

function createBaseGetLoanDecisionRequest(): GetLoanDecisionRequest {
  return { loanDecisionId: 0 };
}

export const GetLoanDecisionRequest: MessageFns<GetLoanDecisionRequest> = {
  encode(message: GetLoanDecisionRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.loanDecisionId !== 0) {
      writer.uint32(8).uint64(message.loanDecisionId);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetLoanDecisionRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLoanDecisionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.loanDecisionId = longToNumber(reader.uint64());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLoanDecisionRequest {
    return { loanDecisionId: isSet(object.loanDecisionId) ? globalThis.Number(object.loanDecisionId) : 0 };
  },

  toJSON(message: GetLoanDecisionRequest): unknown {
    const obj: any = {};
    if (message.loanDecisionId !== 0) {
      obj.loanDecisionId = Math.round(message.loanDecisionId);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLoanDecisionRequest>, I>>(base?: I): GetLoanDecisionRequest {
    return GetLoanDecisionRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetLoanDecisionRequest>, I>>(object: I): GetLoanDecisionRequest {
    const message = createBaseGetLoanDecisionRequest();
    message.loanDecisionId = object.loanDecisionId ?? 0;
    return message;
  },
};

function createBaseGetLoanDecisionResponse(): GetLoanDecisionResponse {
  return { loanDecisionId: 0, loanApplicationId: 0, aiDecision: "", reasoning: "", createdAt: undefined };
}

export const GetLoanDecisionResponse: MessageFns<GetLoanDecisionResponse> = {
  encode(message: GetLoanDecisionResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.loanDecisionId !== 0) {
      writer.uint32(8).uint64(message.loanDecisionId);
    }
    if (message.loanApplicationId !== 0) {
      writer.uint32(16).uint64(message.loanApplicationId);
    }
    if (message.aiDecision !== "") {
      writer.uint32(26).string(message.aiDecision);
    }
    if (message.reasoning !== "") {
      writer.uint32(34).string(message.reasoning);
    }
    if (message.createdAt !== undefined) {
      Timestamp.encode(toTimestamp(message.createdAt), writer.uint32(42).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetLoanDecisionResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLoanDecisionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.loanDecisionId = longToNumber(reader.uint64());
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.loanApplicationId = longToNumber(reader.uint64());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.aiDecision = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.reasoning = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.createdAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLoanDecisionResponse {
    return {
      loanDecisionId: isSet(object.loanDecisionId) ? globalThis.Number(object.loanDecisionId) : 0,
      loanApplicationId: isSet(object.loanApplicationId) ? globalThis.Number(object.loanApplicationId) : 0,
      aiDecision: isSet(object.aiDecision) ? globalThis.String(object.aiDecision) : "",
      reasoning: isSet(object.reasoning) ? globalThis.String(object.reasoning) : "",
      createdAt: isSet(object.createdAt) ? fromJsonTimestamp(object.createdAt) : undefined,
    };
  },

  toJSON(message: GetLoanDecisionResponse): unknown {
    const obj: any = {};
    if (message.loanDecisionId !== 0) {
      obj.loanDecisionId = Math.round(message.loanDecisionId);
    }
    if (message.loanApplicationId !== 0) {
      obj.loanApplicationId = Math.round(message.loanApplicationId);
    }
    if (message.aiDecision !== "") {
      obj.aiDecision = message.aiDecision;
    }
    if (message.reasoning !== "") {
      obj.reasoning = message.reasoning;
    }
    if (message.createdAt !== undefined) {
      obj.createdAt = message.createdAt.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLoanDecisionResponse>, I>>(base?: I): GetLoanDecisionResponse {
    return GetLoanDecisionResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetLoanDecisionResponse>, I>>(object: I): GetLoanDecisionResponse {
    const message = createBaseGetLoanDecisionResponse();
    message.loanDecisionId = object.loanDecisionId ?? 0;
    message.loanApplicationId = object.loanApplicationId ?? 0;
    message.aiDecision = object.aiDecision ?? "";
    message.reasoning = object.reasoning ?? "";
    message.createdAt = object.createdAt ?? undefined;
    return message;
  },
};

/** LoanDecision service for interacting with loan decisions */
export interface LoanDecisionService {
  CreateLoanDecision(
    request: DeepPartial<CreateLoanDecisionRequest>,
    metadata?: grpc.Metadata,
  ): Promise<CreateLoanDecisionResponse>;
  GetLoanDecision(
    request: DeepPartial<GetLoanDecisionRequest>,
    metadata?: grpc.Metadata,
  ): Promise<GetLoanDecisionResponse>;
}

export class LoanDecisionServiceClientImpl implements LoanDecisionService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateLoanDecision = this.CreateLoanDecision.bind(this);
    this.GetLoanDecision = this.GetLoanDecision.bind(this);
  }

  CreateLoanDecision(
    request: DeepPartial<CreateLoanDecisionRequest>,
    metadata?: grpc.Metadata,
  ): Promise<CreateLoanDecisionResponse> {
    return this.rpc.unary(
      LoanDecisionServiceCreateLoanDecisionDesc,
      CreateLoanDecisionRequest.fromPartial(request),
      metadata,
    );
  }

  GetLoanDecision(
    request: DeepPartial<GetLoanDecisionRequest>,
    metadata?: grpc.Metadata,
  ): Promise<GetLoanDecisionResponse> {
    return this.rpc.unary(
      LoanDecisionServiceGetLoanDecisionDesc,
      GetLoanDecisionRequest.fromPartial(request),
      metadata,
    );
  }
}

export const LoanDecisionServiceDesc = { serviceName: "loan_decision.LoanDecisionService" };

export const LoanDecisionServiceCreateLoanDecisionDesc: UnaryMethodDefinitionish = {
  methodName: "CreateLoanDecision",
  service: LoanDecisionServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return CreateLoanDecisionRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = CreateLoanDecisionResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const LoanDecisionServiceGetLoanDecisionDesc: UnaryMethodDefinitionish = {
  methodName: "GetLoanDecision",
  service: LoanDecisionServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetLoanDecisionRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetLoanDecisionResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

interface UnaryMethodDefinitionishR extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any;
  responseStream: any;
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR;

interface Rpc {
  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined,
  ): Promise<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;

    debug?: boolean;
    metadata?: grpc.Metadata;
    upStreamRetryCodes?: number[];
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;

      debug?: boolean;
      metadata?: grpc.Metadata;
      upStreamRetryCodes?: number[];
    },
  ) {
    this.host = host;
    this.options = options;
  }

  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined,
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata = metadata && this.options.metadata
      ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata ?? this.options.metadata;
    return new Promise((resolve, reject) => {
      grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata ?? {},
        ...(this.options.transport !== undefined ? { transport: this.options.transport } : {}),
        debug: this.options.debug ?? false,
        onEnd: function (response) {
          if (response.status === grpc.Code.OK) {
            resolve(response.message!.toObject());
          } else {
            const err = new GrpcWebError(response.statusMessage, response.status, response.trailers);
            reject(err);
          }
        },
      });
    });
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = Math.trunc(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof globalThis.Date) {
    return o;
  } else if (typeof o === "string") {
    return new globalThis.Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function longToNumber(int64: { toString(): string }): number {
  const num = globalThis.Number(int64.toString());
  if (num > globalThis.Number.MAX_SAFE_INTEGER) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  if (num < globalThis.Number.MIN_SAFE_INTEGER) {
    throw new globalThis.Error("Value is smaller than Number.MIN_SAFE_INTEGER");
  }
  return num;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends globalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
