# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: tendermint/version/types.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'tendermint/version/types.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1etendermint/version/types.proto\x12\x12tendermint.version\x1a\x14gogoproto/gogo.proto\"=\n\x03\x41pp\x12\x1a\n\x08protocol\x18\x01 \x01(\x04R\x08protocol\x12\x1a\n\x08software\x18\x02 \x01(\tR\x08software\"9\n\tConsensus\x12\x14\n\x05\x62lock\x18\x01 \x01(\x04R\x05\x62lock\x12\x10\n\x03\x61pp\x18\x02 \x01(\x04R\x03\x61pp:\x04\xe8\xa0\x1f\x01\x42\x37Z5github.com/cometbft/cometbft/proto/tendermint/versionb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'tendermint.version.types_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z5github.com/cometbft/cometbft/proto/tendermint/version'
  _globals['_CONSENSUS']._loaded_options = None
  _globals['_CONSENSUS']._serialized_options = b'\350\240\037\001'
  _globals['_APP']._serialized_start=76
  _globals['_APP']._serialized_end=137
  _globals['_CONSENSUS']._serialized_start=139
  _globals['_CONSENSUS']._serialized_end=196
# @@protoc_insertion_point(module_scope)
