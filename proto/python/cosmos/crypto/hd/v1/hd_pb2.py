# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/crypto/hd/v1/hd.proto
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
    'cosmos/crypto/hd/v1/hd.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x63osmos/crypto/hd/v1/hd.proto\x12\x13\x63osmos.crypto.hd.v1\x1a\x11\x61mino/amino.proto\x1a\x14gogoproto/gogo.proto\"\xc0\x01\n\x0b\x42IP44Params\x12\x18\n\x07purpose\x18\x01 \x01(\rR\x07purpose\x12\x1b\n\tcoin_type\x18\x02 \x01(\rR\x08\x63oinType\x12\x18\n\x07\x61\x63\x63ount\x18\x03 \x01(\rR\x07\x61\x63\x63ount\x12\x16\n\x06\x63hange\x18\x04 \x01(\x08R\x06\x63hange\x12#\n\raddress_index\x18\x05 \x01(\rR\x0c\x61\x64\x64ressIndex:#\x98\xa0\x1f\x00\x8a\xe7\xb0*\x1a\x63rypto/keys/hd/BIP44ParamsB,Z&github.com/cosmos/cosmos-sdk/crypto/hd\xc8\xe1\x1e\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.crypto.hd.v1.hd_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z&github.com/cosmos/cosmos-sdk/crypto/hd\310\341\036\000'
  _globals['_BIP44PARAMS']._loaded_options = None
  _globals['_BIP44PARAMS']._serialized_options = b'\230\240\037\000\212\347\260*\032crypto/keys/hd/BIP44Params'
  _globals['_BIP44PARAMS']._serialized_start=95
  _globals['_BIP44PARAMS']._serialized_end=287
# @@protoc_insertion_point(module_scope)
