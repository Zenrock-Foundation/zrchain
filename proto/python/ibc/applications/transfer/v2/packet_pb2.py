# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ibc/applications/transfer/v2/packet.proto
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
    'ibc/applications/transfer/v2/packet.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ibc.applications.transfer.v2 import token_pb2 as ibc_dot_applications_dot_transfer_dot_v2_dot_token__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)ibc/applications/transfer/v2/packet.proto\x12\x1cibc.applications.transfer.v2\x1a(ibc/applications/transfer/v2/token.proto\x1a\x14gogoproto/gogo.proto\"\x8f\x01\n\x17\x46ungibleTokenPacketData\x12\x14\n\x05\x64\x65nom\x18\x01 \x01(\tR\x05\x64\x65nom\x12\x16\n\x06\x61mount\x18\x02 \x01(\tR\x06\x61mount\x12\x16\n\x06sender\x18\x03 \x01(\tR\x06sender\x12\x1a\n\x08receiver\x18\x04 \x01(\tR\x08receiver\x12\x12\n\x04memo\x18\x05 \x01(\tR\x04memo\"\xa4\x01\n\x19\x46ungibleTokenPacketDataV2\x12?\n\x05token\x18\x01 \x01(\x0b\x32#.ibc.applications.transfer.v2.TokenB\x04\xc8\xde\x1f\x00R\x05token\x12\x16\n\x06sender\x18\x02 \x01(\tR\x06sender\x12\x1a\n\x08receiver\x18\x03 \x01(\tR\x08receiver\x12\x12\n\x04memo\x18\x04 \x01(\tR\x04memoB:Z8github.com/cosmos/ibc-go/v10/modules/apps/transfer/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ibc.applications.transfer.v2.packet_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z8github.com/cosmos/ibc-go/v10/modules/apps/transfer/types'
  _globals['_FUNGIBLETOKENPACKETDATAV2'].fields_by_name['token']._loaded_options = None
  _globals['_FUNGIBLETOKENPACKETDATAV2'].fields_by_name['token']._serialized_options = b'\310\336\037\000'
  _globals['_FUNGIBLETOKENPACKETDATA']._serialized_start=140
  _globals['_FUNGIBLETOKENPACKETDATA']._serialized_end=283
  _globals['_FUNGIBLETOKENPACKETDATAV2']._serialized_start=286
  _globals['_FUNGIBLETOKENPACKETDATAV2']._serialized_end=450
# @@protoc_insertion_point(module_scope)
