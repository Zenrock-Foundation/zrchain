# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/mint/v1beta1/genesis.proto
# Protobuf Python Version: 5.29.1
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
    1,
    '',
    'zrchain/mint/v1beta1/genesis.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from zrchain.mint.v1beta1 import mint_pb2 as zrchain_dot_mint_dot_v1beta1_dot_mint__pb2
from amino import amino_pb2 as amino_dot_amino__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"zrchain/mint/v1beta1/genesis.proto\x12\x14zrchain.mint.v1beta1\x1a\x14gogoproto/gogo.proto\x1a\x1fzrchain/mint/v1beta1/mint.proto\x1a\x11\x61mino/amino.proto\"\x90\x01\n\x0cGenesisState\x12?\n\x06minter\x18\x01 \x01(\x0b\x32\x1c.zrchain.mint.v1beta1.MinterB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06minter\x12?\n\x06params\x18\x02 \x01(\x0b\x32\x1c.zrchain.mint.v1beta1.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06paramsB7Z5github.com/Zenrock-Foundation/zrchain/v5/x/mint/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.mint.v1beta1.genesis_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z5github.com/Zenrock-Foundation/zrchain/v5/x/mint/types'
  _globals['_GENESISSTATE'].fields_by_name['minter']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['minter']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['params']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['params']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE']._serialized_start=135
  _globals['_GENESISSTATE']._serialized_end=279
# @@protoc_insertion_point(module_scope)
