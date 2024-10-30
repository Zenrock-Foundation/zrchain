# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/protocolpool/v1/types.proto
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
    'cosmos/protocolpool/v1/types.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from cosmos.base.v1beta1 import coin_pb2 as cosmos_dot_base_dot_v1beta1_dot_coin__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"cosmos/protocolpool/v1/types.proto\x12\x16\x63osmos.protocolpool.v1\x1a\x14gogoproto/gogo.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x1e\x63osmos/base/v1beta1/coin.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"\xd4\x03\n\x06\x42udget\x12\x45\n\x11recipient_address\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10recipientAddress\x12<\n\x0ctotal_budget\x18\x02 \x01(\x0b\x32\x19.cosmos.base.v1beta1.CoinR\x0btotalBudget\x12@\n\x0e\x63laimed_amount\x18\x03 \x01(\x0b\x32\x19.cosmos.base.v1beta1.CoinR\rclaimedAmount\x12?\n\nstart_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x04\x90\xdf\x1f\x01R\tstartTime\x12H\n\x0fnext_claim_from\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x04\x90\xdf\x1f\x01R\rnextClaimFrom\x12\x1a\n\x08tranches\x18\x06 \x01(\x04R\x08tranches\x12#\n\rtranches_left\x18\x07 \x01(\x04R\x0ctranchesLeft\x12\x37\n\x06period\x18\x08 \x01(\x0b\x32\x19.google.protobuf.DurationB\x04\x98\xdf\x1f\x01R\x06period\"\xd5\x01\n\x0e\x43ontinuousFund\x12\x36\n\trecipient\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\trecipient\x12Q\n\npercentage\x18\x02 \x01(\tB1\xc8\xde\x1f\x00\xda\xde\x1f\x1b\x63osmossdk.io/math.LegacyDec\xd2\xb4-\ncosmos.DecR\npercentage\x12\x38\n\x06\x65xpiry\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x04\x90\xdf\x1f\x01R\x06\x65xpiryB#Z!cosmossdk.io/x/protocolpool/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.protocolpool.v1.types_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z!cosmossdk.io/x/protocolpool/types'
  _globals['_BUDGET'].fields_by_name['recipient_address']._loaded_options = None
  _globals['_BUDGET'].fields_by_name['recipient_address']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_BUDGET'].fields_by_name['start_time']._loaded_options = None
  _globals['_BUDGET'].fields_by_name['start_time']._serialized_options = b'\220\337\037\001'
  _globals['_BUDGET'].fields_by_name['next_claim_from']._loaded_options = None
  _globals['_BUDGET'].fields_by_name['next_claim_from']._serialized_options = b'\220\337\037\001'
  _globals['_BUDGET'].fields_by_name['period']._loaded_options = None
  _globals['_BUDGET'].fields_by_name['period']._serialized_options = b'\230\337\037\001'
  _globals['_CONTINUOUSFUND'].fields_by_name['recipient']._loaded_options = None
  _globals['_CONTINUOUSFUND'].fields_by_name['recipient']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_CONTINUOUSFUND'].fields_by_name['percentage']._loaded_options = None
  _globals['_CONTINUOUSFUND'].fields_by_name['percentage']._serialized_options = b'\310\336\037\000\332\336\037\033cosmossdk.io/math.LegacyDec\322\264-\ncosmos.Dec'
  _globals['_CONTINUOUSFUND'].fields_by_name['expiry']._loaded_options = None
  _globals['_CONTINUOUSFUND'].fields_by_name['expiry']._serialized_options = b'\220\337\037\001'
  _globals['_BUDGET']._serialized_start=209
  _globals['_BUDGET']._serialized_end=677
  _globals['_CONTINUOUSFUND']._serialized_start=680
  _globals['_CONTINUOUSFUND']._serialized_end=893
# @@protoc_insertion_point(module_scope)
