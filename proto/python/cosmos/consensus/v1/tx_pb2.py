# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/consensus/v1/tx.proto
# Protobuf Python Version: 6.30.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    30,
    2,
    '',
    'cosmos/consensus/v1/tx.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from cosmos.msg.v1 import msg_pb2 as cosmos_dot_msg_dot_v1_dot_msg__pb2
from tendermint.types import params_pb2 as tendermint_dot_types_dot_params__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x63osmos/consensus/v1/tx.proto\x12\x13\x63osmos.consensus.v1\x1a\x11\x61mino/amino.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x17\x63osmos/msg/v1/msg.proto\x1a\x1dtendermint/types/params.proto\"\xea\x02\n\x0fMsgUpdateParams\x12\x36\n\tauthority\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\tauthority\x12\x33\n\x05\x62lock\x18\x02 \x01(\x0b\x32\x1d.tendermint.types.BlockParamsR\x05\x62lock\x12<\n\x08\x65vidence\x18\x03 \x01(\x0b\x32 .tendermint.types.EvidenceParamsR\x08\x65vidence\x12?\n\tvalidator\x18\x04 \x01(\x0b\x32!.tendermint.types.ValidatorParamsR\tvalidator\x12\x30\n\x04\x61\x62\x63i\x18\x05 \x01(\x0b\x32\x1c.tendermint.types.ABCIParamsR\x04\x61\x62\x63i:9\x82\xe7\xb0*\tauthority\x8a\xe7\xb0*&cosmos-sdk/x/consensus/MsgUpdateParams\"\x19\n\x17MsgUpdateParamsResponse2p\n\x03Msg\x12\x62\n\x0cUpdateParams\x12$.cosmos.consensus.v1.MsgUpdateParams\x1a,.cosmos.consensus.v1.MsgUpdateParamsResponse\x1a\x05\x80\xe7\xb0*\x01\x42\x30Z.github.com/cosmos/cosmos-sdk/x/consensus/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.consensus.v1.tx_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z.github.com/cosmos/cosmos-sdk/x/consensus/types'
  _globals['_MSGUPDATEPARAMS'].fields_by_name['authority']._loaded_options = None
  _globals['_MSGUPDATEPARAMS'].fields_by_name['authority']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_MSGUPDATEPARAMS']._loaded_options = None
  _globals['_MSGUPDATEPARAMS']._serialized_options = b'\202\347\260*\tauthority\212\347\260*&cosmos-sdk/x/consensus/MsgUpdateParams'
  _globals['_MSG']._loaded_options = None
  _globals['_MSG']._serialized_options = b'\200\347\260*\001'
  _globals['_MSGUPDATEPARAMS']._serialized_start=156
  _globals['_MSGUPDATEPARAMS']._serialized_end=518
  _globals['_MSGUPDATEPARAMSRESPONSE']._serialized_start=520
  _globals['_MSGUPDATEPARAMSRESPONSE']._serialized_end=545
  _globals['_MSG']._serialized_start=547
  _globals['_MSG']._serialized_end=659
# @@protoc_insertion_point(module_scope)
