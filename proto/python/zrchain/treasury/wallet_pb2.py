# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/treasury/wallet.proto
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
    'zrchain/treasury/wallet.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dzrchain/treasury/wallet.proto\x12\x10zrchain.treasury*\xc4\x01\n\nWalletType\x12\x1b\n\x17WALLET_TYPE_UNSPECIFIED\x10\x00\x12\x16\n\x12WALLET_TYPE_NATIVE\x10\x01\x12\x13\n\x0fWALLET_TYPE_EVM\x10\x02\x12\x1b\n\x17WALLET_TYPE_BTC_TESTNET\x10\x03\x12\x1b\n\x17WALLET_TYPE_BTC_MAINNET\x10\x04\x12\x1a\n\x16WALLET_TYPE_BTC_REGNET\x10\x05\x12\x16\n\x12WALLET_TYPE_SOLANA\x10\x06\x42;Z9github.com/Zenrock-Foundation/zrchain/v4/x/treasury/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.treasury.wallet_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z9github.com/Zenrock-Foundation/zrchain/v4/x/treasury/types'
  _globals['_WALLETTYPE']._serialized_start=52
  _globals['_WALLETTYPE']._serialized_end=248
# @@protoc_insertion_point(module_scope)
