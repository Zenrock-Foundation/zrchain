# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: tendermint/blocksync/types.proto
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
    'tendermint/blocksync/types.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from tendermint.types import block_pb2 as tendermint_dot_types_dot_block__pb2
from tendermint.types import types_pb2 as tendermint_dot_types_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n tendermint/blocksync/types.proto\x12\x14tendermint.blocksync\x1a\x1ctendermint/types/block.proto\x1a\x1ctendermint/types/types.proto\"&\n\x0c\x42lockRequest\x12\x16\n\x06height\x18\x01 \x01(\x03R\x06height\")\n\x0fNoBlockResponse\x12\x16\n\x06height\x18\x01 \x01(\x03R\x06height\"\x7f\n\rBlockResponse\x12-\n\x05\x62lock\x18\x01 \x01(\x0b\x32\x17.tendermint.types.BlockR\x05\x62lock\x12?\n\next_commit\x18\x02 \x01(\x0b\x32 .tendermint.types.ExtendedCommitR\textCommit\"\x0f\n\rStatusRequest\"<\n\x0eStatusResponse\x12\x16\n\x06height\x18\x01 \x01(\x03R\x06height\x12\x12\n\x04\x62\x61se\x18\x02 \x01(\x03R\x04\x62\x61se\"\x9d\x03\n\x07Message\x12I\n\rblock_request\x18\x01 \x01(\x0b\x32\".tendermint.blocksync.BlockRequestH\x00R\x0c\x62lockRequest\x12S\n\x11no_block_response\x18\x02 \x01(\x0b\x32%.tendermint.blocksync.NoBlockResponseH\x00R\x0fnoBlockResponse\x12L\n\x0e\x62lock_response\x18\x03 \x01(\x0b\x32#.tendermint.blocksync.BlockResponseH\x00R\rblockResponse\x12L\n\x0estatus_request\x18\x04 \x01(\x0b\x32#.tendermint.blocksync.StatusRequestH\x00R\rstatusRequest\x12O\n\x0fstatus_response\x18\x05 \x01(\x0b\x32$.tendermint.blocksync.StatusResponseH\x00R\x0estatusResponseB\x05\n\x03sumB9Z7github.com/cometbft/cometbft/proto/tendermint/blocksyncb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'tendermint.blocksync.types_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z7github.com/cometbft/cometbft/proto/tendermint/blocksync'
  _globals['_BLOCKREQUEST']._serialized_start=118
  _globals['_BLOCKREQUEST']._serialized_end=156
  _globals['_NOBLOCKRESPONSE']._serialized_start=158
  _globals['_NOBLOCKRESPONSE']._serialized_end=199
  _globals['_BLOCKRESPONSE']._serialized_start=201
  _globals['_BLOCKRESPONSE']._serialized_end=328
  _globals['_STATUSREQUEST']._serialized_start=330
  _globals['_STATUSREQUEST']._serialized_end=345
  _globals['_STATUSRESPONSE']._serialized_start=347
  _globals['_STATUSRESPONSE']._serialized_end=407
  _globals['_MESSAGE']._serialized_start=410
  _globals['_MESSAGE']._serialized_end=823
# @@protoc_insertion_point(module_scope)
