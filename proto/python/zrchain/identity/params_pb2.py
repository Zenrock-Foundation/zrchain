# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/identity/params.proto
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
    'zrchain/identity/params.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dzrchain/identity/params.proto\x12\x10zrchain.identity\x1a\x11\x61mino/amino.proto\x1a\x14gogoproto/gogo.proto\"\x7f\n\x06Params\x12\x30\n\x14keyring_creation_fee\x18\x01 \x01(\x04R\x12keyringCreationFee:C\xe8\xa0\x1f\x01\x8a\xe7\xb0*:github.com/Zenrock-Foundation/zrchain/v4/x/identity/ParamsB;Z9github.com/Zenrock-Foundation/zrchain/v4/x/identity/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.identity.params_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z9github.com/Zenrock-Foundation/zrchain/v4/x/identity/types'
  _globals['_PARAMS']._loaded_options = None
  _globals['_PARAMS']._serialized_options = b'\350\240\037\001\212\347\260*:github.com/Zenrock-Foundation/zrchain/v4/x/identity/Params'
  _globals['_PARAMS']._serialized_start=92
  _globals['_PARAMS']._serialized_end=219
# @@protoc_insertion_point(module_scope)
