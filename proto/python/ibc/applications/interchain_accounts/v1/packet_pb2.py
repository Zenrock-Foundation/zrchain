# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ibc/applications/interchain_accounts/v1/packet.proto
# Protobuf Python Version: 5.28.3
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
    3,
    '',
    'ibc/applications/interchain_accounts/v1/packet.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n4ibc/applications/interchain_accounts/v1/packet.proto\x12\'ibc.applications.interchain_accounts.v1\x1a\x19google/protobuf/any.proto\x1a\x14gogoproto/gogo.proto\"\x88\x01\n\x1bInterchainAccountPacketData\x12\x41\n\x04type\x18\x01 \x01(\x0e\x32-.ibc.applications.interchain_accounts.v1.TypeR\x04type\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\x0cR\x04\x64\x61ta\x12\x12\n\x04memo\x18\x03 \x01(\tR\x04memo\"<\n\x08\x43osmosTx\x12\x30\n\x08messages\x18\x01 \x03(\x0b\x32\x14.google.protobuf.AnyR\x08messages*X\n\x04Type\x12%\n\x10TYPE_UNSPECIFIED\x10\x00\x1a\x0f\x8a\x9d \x0bUNSPECIFIED\x12#\n\x0fTYPE_EXECUTE_TX\x10\x01\x1a\x0e\x8a\x9d \nEXECUTE_TX\x1a\x04\x88\xa3\x1e\x00\x42GZEgithub.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ibc.applications.interchain_accounts.v1.packet_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZEgithub.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/types'
  _globals['_TYPE']._loaded_options = None
  _globals['_TYPE']._serialized_options = b'\210\243\036\000'
  _globals['_TYPE'].values_by_name["TYPE_UNSPECIFIED"]._loaded_options = None
  _globals['_TYPE'].values_by_name["TYPE_UNSPECIFIED"]._serialized_options = b'\212\235 \013UNSPECIFIED'
  _globals['_TYPE'].values_by_name["TYPE_EXECUTE_TX"]._loaded_options = None
  _globals['_TYPE'].values_by_name["TYPE_EXECUTE_TX"]._serialized_options = b'\212\235 \nEXECUTE_TX'
  _globals['_TYPE']._serialized_start=347
  _globals['_TYPE']._serialized_end=435
  _globals['_INTERCHAINACCOUNTPACKETDATA']._serialized_start=147
  _globals['_INTERCHAINACCOUNTPACKETDATA']._serialized_end=283
  _globals['_COSMOSTX']._serialized_start=285
  _globals['_COSMOSTX']._serialized_end=345
# @@protoc_insertion_point(module_scope)