# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/zenbtc/tx.proto
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
    'zrchain/zenbtc/tx.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from cosmos.msg.v1 import msg_pb2 as cosmos_dot_msg_dot_v1_dot_msg__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from zrchain.zenbtc import params_pb2 as zrchain_dot_zenbtc_dot_params__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17zrchain/zenbtc/tx.proto\x12\x0ezrchain.zenbtc\x1a\x11\x61mino/amino.proto\x1a\x17\x63osmos/msg/v1/msg.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x14gogoproto/gogo.proto\x1a\x1bzrchain/zenbtc/params.proto\"\xb9\x01\n\x0fMsgUpdateParams\x12\x36\n\tauthority\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\tauthority\x12\x39\n\x06params\x18\x02 \x01(\x0b\x32\x16.zrchain.zenbtc.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06params:3\x82\xe7\xb0*\tauthority\x8a\xe7\xb0* zrchain/x/zenbtc/MsgUpdateParams\"\x19\n\x17MsgUpdateParamsResponse\"\x88\x02\n\x1eMsgVerifyDepositBlockInclusion\x12\x18\n\x07\x63reator\x18\x01 \x01(\tR\x07\x63reator\x12\x1d\n\nchain_name\x18\x02 \x01(\tR\tchainName\x12!\n\x0c\x62lock_height\x18\x03 \x01(\x03R\x0b\x62lockHeight\x12\x15\n\x06raw_tx\x18\x04 \x01(\tR\x05rawTx\x12\x14\n\x05index\x18\x05 \x01(\x05R\x05index\x12\x14\n\x05proof\x18\x06 \x03(\tR\x05proof\x12!\n\x0c\x64\x65posit_addr\x18\x07 \x01(\tR\x0b\x64\x65positAddr\x12\x16\n\x06\x61mount\x18\x08 \x01(\x04R\x06\x61mount:\x0c\x82\xe7\xb0*\x07\x63reator\"(\n&MsgVerifyDepositBlockInclusionResponse\"\xae\x01\n MsgSubmitSolanaUnlockTransaction\x12\x18\n\x07\x63reator\x18\x01 \x01(\tR\x07\x63reator\x12!\n\x0ctx_signature\x18\x02 \x01(\tR\x0btxSignature\x12\'\n\x0fwithdrawal_addr\x18\x03 \x01(\tR\x0ewithdrawalAddr\x12\x16\n\x06\x61mount\x18\x04 \x01(\x04R\x06\x61mount:\x0c\x82\xe7\xb0*\x07\x63reator\"*\n(MsgSubmitSolanaUnlockTransactionResponse2\xfc\x02\n\x03Msg\x12X\n\x0cUpdateParams\x12\x1f.zrchain.zenbtc.MsgUpdateParams\x1a\'.zrchain.zenbtc.MsgUpdateParamsResponse\x12\x85\x01\n\x1bVerifyDepositBlockInclusion\x12..zrchain.zenbtc.MsgVerifyDepositBlockInclusion\x1a\x36.zrchain.zenbtc.MsgVerifyDepositBlockInclusionResponse\x12\x8b\x01\n\x1dSubmitSolanaUnlockTransaction\x12\x30.zrchain.zenbtc.MsgSubmitSolanaUnlockTransaction\x1a\x38.zrchain.zenbtc.MsgSubmitSolanaUnlockTransactionResponse\x1a\x05\x80\xe7\xb0*\x01\x42:Z8github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.zenbtc.tx_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z8github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types'
  _globals['_MSGUPDATEPARAMS'].fields_by_name['authority']._loaded_options = None
  _globals['_MSGUPDATEPARAMS'].fields_by_name['authority']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_MSGUPDATEPARAMS'].fields_by_name['params']._loaded_options = None
  _globals['_MSGUPDATEPARAMS'].fields_by_name['params']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_MSGUPDATEPARAMS']._loaded_options = None
  _globals['_MSGUPDATEPARAMS']._serialized_options = b'\202\347\260*\tauthority\212\347\260* zrchain/x/zenbtc/MsgUpdateParams'
  _globals['_MSGVERIFYDEPOSITBLOCKINCLUSION']._loaded_options = None
  _globals['_MSGVERIFYDEPOSITBLOCKINCLUSION']._serialized_options = b'\202\347\260*\007creator'
  _globals['_MSGSUBMITSOLANAUNLOCKTRANSACTION']._loaded_options = None
  _globals['_MSGSUBMITSOLANAUNLOCKTRANSACTION']._serialized_options = b'\202\347\260*\007creator'
  _globals['_MSG']._loaded_options = None
  _globals['_MSG']._serialized_options = b'\200\347\260*\001'
  _globals['_MSGUPDATEPARAMS']._serialized_start=166
  _globals['_MSGUPDATEPARAMS']._serialized_end=351
  _globals['_MSGUPDATEPARAMSRESPONSE']._serialized_start=353
  _globals['_MSGUPDATEPARAMSRESPONSE']._serialized_end=378
  _globals['_MSGVERIFYDEPOSITBLOCKINCLUSION']._serialized_start=381
  _globals['_MSGVERIFYDEPOSITBLOCKINCLUSION']._serialized_end=645
  _globals['_MSGVERIFYDEPOSITBLOCKINCLUSIONRESPONSE']._serialized_start=647
  _globals['_MSGVERIFYDEPOSITBLOCKINCLUSIONRESPONSE']._serialized_end=687
  _globals['_MSGSUBMITSOLANAUNLOCKTRANSACTION']._serialized_start=690
  _globals['_MSGSUBMITSOLANAUNLOCKTRANSACTION']._serialized_end=864
  _globals['_MSGSUBMITSOLANAUNLOCKTRANSACTIONRESPONSE']._serialized_start=866
  _globals['_MSGSUBMITSOLANAUNLOCKTRANSACTIONRESPONSE']._serialized_end=908
  _globals['_MSG']._serialized_start=911
  _globals['_MSG']._serialized_end=1291
# @@protoc_insertion_point(module_scope)
