# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/counter/v1/query.proto
# Protobuf Python Version: 5.28.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    2,
    '',
    'cosmos/counter/v1/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1d\x63osmos/counter/v1/query.proto\x12\x11\x63osmos.counter.v1\"\x16\n\x14QueryGetCountRequest\"8\n\x15QueryGetCountResponse\x12\x1f\n\x0btotal_count\x18\x01 \x01(\x03R\ntotalCount2f\n\x05Query\x12]\n\x08GetCount\x12\'.cosmos.counter.v1.QueryGetCountRequest\x1a(.cosmos.counter.v1.QueryGetCountResponseB.Z,github.com/cosmos/cosmos-sdk/x/counter/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.counter.v1.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,github.com/cosmos/cosmos-sdk/x/counter/types'
  _globals['_QUERYGETCOUNTREQUEST']._serialized_start=52
  _globals['_QUERYGETCOUNTREQUEST']._serialized_end=74
  _globals['_QUERYGETCOUNTRESPONSE']._serialized_start=76
  _globals['_QUERYGETCOUNTRESPONSE']._serialized_end=132
  _globals['_QUERY']._serialized_start=134
  _globals['_QUERY']._serialized_end=236
# @@protoc_insertion_point(module_scope)
