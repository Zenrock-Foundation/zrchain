# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/zenbtc/query.proto
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
    'zrchain/zenbtc/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from cosmos.base.query.v1beta1 import pagination_pb2 as cosmos_dot_base_dot_query_dot_v1beta1_dot_pagination__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from zrchain.zenbtc import params_pb2 as zrchain_dot_zenbtc_dot_params__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1azrchain/zenbtc/query.proto\x12\x0ezrchain.zenbtc\x1a\x11\x61mino/amino.proto\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a\x14gogoproto/gogo.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1bzrchain/zenbtc/params.proto\"\x14\n\x12QueryParamsRequest\"P\n\x13QueryParamsResponse\x12\x39\n\x06params\x18\x01 \x01(\x0b\x32\x16.zrchain.zenbtc.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06params\"\x1e\n\x1cQueryLockTransactionsRequest\"L\n\x1dQueryLockTransactionsResponse\x12+\n\x11lock_transactions\x18\x01 \x03(\tR\x10lockTransactions\"?\n\'QueryConfirmedUnlockTransactionsRequest\x12\x14\n\x05\x63hain\x18\x01 \x01(\tR\x05\x63hain\"[\n(QueryConfirmedUnlockTransactionsResponse\x12/\n\x13unlock_transactions\x18\x01 \x03(\tR\x12unlockTransactions2\xf1\x03\n\x05Query\x12t\n\x06Params\x12\".zrchain.zenbtc.QueryParamsRequest\x1a#.zrchain.zenbtc.QueryParamsResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/zrchain/v4/zenbtc/params\x12\x9d\x01\n\x10LockTransactions\x12,.zrchain.zenbtc.QueryLockTransactionsRequest\x1a-.zrchain.zenbtc.QueryLockTransactionsResponse\",\x82\xd3\xe4\x93\x02&\x12$/zrchain/v4/zenbtc/lock_transactions\x12\xd1\x01\n\x1b\x43onfirmedUnlockTransactions\x12\x37.zrchain.zenbtc.QueryConfirmedUnlockTransactionsRequest\x1a\x38.zrchain.zenbtc.QueryConfirmedUnlockTransactionsResponse\"?\x82\xd3\xe4\x93\x02\x39\x12\x37/zrchain/v4/zenbtc/confirmed_solana_unlock_transactionsB9Z7github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.zenbtc.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z7github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types'
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['params']._loaded_options = None
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['params']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERY'].methods_by_name['Params']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Params']._serialized_options = b'\202\323\344\223\002\033\022\031/zrchain/v4/zenbtc/params'
  _globals['_QUERY'].methods_by_name['LockTransactions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['LockTransactions']._serialized_options = b'\202\323\344\223\002&\022$/zrchain/v4/zenbtc/lock_transactions'
  _globals['_QUERY'].methods_by_name['ConfirmedUnlockTransactions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['ConfirmedUnlockTransactions']._serialized_options = b'\202\323\344\223\0029\0227/zrchain/v4/zenbtc/confirmed_solana_unlock_transactions'
  _globals['_QUERYPARAMSREQUEST']._serialized_start=190
  _globals['_QUERYPARAMSREQUEST']._serialized_end=210
  _globals['_QUERYPARAMSRESPONSE']._serialized_start=212
  _globals['_QUERYPARAMSRESPONSE']._serialized_end=292
  _globals['_QUERYLOCKTRANSACTIONSREQUEST']._serialized_start=294
  _globals['_QUERYLOCKTRANSACTIONSREQUEST']._serialized_end=324
  _globals['_QUERYLOCKTRANSACTIONSRESPONSE']._serialized_start=326
  _globals['_QUERYLOCKTRANSACTIONSRESPONSE']._serialized_end=402
  _globals['_QUERYCONFIRMEDUNLOCKTRANSACTIONSREQUEST']._serialized_start=404
  _globals['_QUERYCONFIRMEDUNLOCKTRANSACTIONSREQUEST']._serialized_end=467
  _globals['_QUERYCONFIRMEDUNLOCKTRANSACTIONSRESPONSE']._serialized_start=469
  _globals['_QUERYCONFIRMEDUNLOCKTRANSACTIONSRESPONSE']._serialized_end=560
  _globals['_QUERY']._serialized_start=563
  _globals['_QUERY']._serialized_end=1060
# @@protoc_insertion_point(module_scope)
