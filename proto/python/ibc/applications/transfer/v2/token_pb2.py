# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ibc/applications/transfer/v2/token.proto
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
    'ibc/applications/transfer/v2/token.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ibc.applications.transfer.v1 import transfer_pb2 as ibc_dot_applications_dot_transfer_dot_v1_dot_transfer__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(ibc/applications/transfer/v2/token.proto\x12\x1cibc.applications.transfer.v2\x1a+ibc/applications/transfer/v1/transfer.proto\x1a\x14gogoproto/gogo.proto\"`\n\x05Token\x12?\n\x05\x64\x65nom\x18\x01 \x01(\x0b\x32#.ibc.applications.transfer.v2.DenomB\x04\xc8\xde\x1f\x00R\x05\x64\x65nom\x12\x16\n\x06\x61mount\x18\x02 \x01(\tR\x06\x61mount\"Z\n\x05\x44\x65nom\x12\x12\n\x04\x62\x61se\x18\x01 \x01(\tR\x04\x62\x61se\x12=\n\x05trace\x18\x03 \x03(\x0b\x32!.ibc.applications.transfer.v1.HopB\x04\xc8\xde\x1f\x00R\x05traceB9Z7github.com/cosmos/ibc-go/v9/modules/apps/transfer/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ibc.applications.transfer.v2.token_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z7github.com/cosmos/ibc-go/v9/modules/apps/transfer/types'
  _globals['_TOKEN'].fields_by_name['denom']._loaded_options = None
  _globals['_TOKEN'].fields_by_name['denom']._serialized_options = b'\310\336\037\000'
  _globals['_DENOM'].fields_by_name['trace']._loaded_options = None
  _globals['_DENOM'].fields_by_name['trace']._serialized_options = b'\310\336\037\000'
  _globals['_TOKEN']._serialized_start=141
  _globals['_TOKEN']._serialized_end=237
  _globals['_DENOM']._serialized_start=239
  _globals['_DENOM']._serialized_end=329
# @@protoc_insertion_point(module_scope)
