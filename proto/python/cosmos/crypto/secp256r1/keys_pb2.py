# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/crypto/secp256r1/keys.proto
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
    'cosmos/crypto/secp256r1/keys.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"cosmos/crypto/secp256r1/keys.proto\x12\x17\x63osmos.crypto.secp256r1\x1a\x14gogoproto/gogo.proto\"\'\n\x06PubKey\x12\x1d\n\x03key\x18\x01 \x01(\x0c\x42\x0b\xda\xde\x1f\x07\x65\x63\x64saPKR\x03key\".\n\x07PrivKey\x12#\n\x06secret\x18\x01 \x01(\x0c\x42\x0b\xda\xde\x1f\x07\x65\x63\x64saSKR\x06secretB@Z2github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1\xc8\xe1\x1e\x00\xd8\xe1\x1e\x00\xc8\xe3\x1e\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.crypto.secp256r1.keys_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z2github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1\310\341\036\000\330\341\036\000\310\343\036\001'
  _globals['_PUBKEY'].fields_by_name['key']._loaded_options = None
  _globals['_PUBKEY'].fields_by_name['key']._serialized_options = b'\332\336\037\007ecdsaPK'
  _globals['_PRIVKEY'].fields_by_name['secret']._loaded_options = None
  _globals['_PRIVKEY'].fields_by_name['secret']._serialized_options = b'\332\336\037\007ecdsaSK'
  _globals['_PUBKEY']._serialized_start=85
  _globals['_PUBKEY']._serialized_end=124
  _globals['_PRIVKEY']._serialized_start=126
  _globals['_PRIVKEY']._serialized_end=172
# @@protoc_insertion_point(module_scope)
