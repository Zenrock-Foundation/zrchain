# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/validation/hybrid_validation.proto
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
    'zrchain/validation/hybrid_validation.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from tendermint.types import types_pb2 as tendermint_dot_types_dot_types__pb2
from zrchain.validation import staking_pb2 as zrchain_dot_validation_dot_staking__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n*zrchain/validation/hybrid_validation.proto\x12\x12zrchain.validation\x1a\x11\x61mino/amino.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x14gogoproto/gogo.proto\x1a\x19google/protobuf/any.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1ctendermint/types/types.proto\x1a zrchain/validation/staking.proto\"\xd7\x07\n\x0bValidatorHV\x12\x43\n\x10operator_address\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x0foperatorAddress\x12Y\n\x10\x63onsensus_pubkey\x18\x02 \x01(\x0b\x32\x14.google.protobuf.AnyB\x18\xca\xb4-\x14\x63osmos.crypto.PubKeyR\x0f\x63onsensusPubkey\x12\x16\n\x06jailed\x18\x03 \x01(\x08R\x06jailed\x12\x36\n\x06status\x18\x04 \x01(\x0e\x32\x1e.zrchain.validation.BondStatusR\x06status\x12O\n\x0ctokensNative\x18\x05 \x01(\tB+\xc8\xde\x1f\x00\xda\xde\x1f\x15\x63osmossdk.io/math.Int\xd2\xb4-\ncosmos.IntR\x0ctokensNative\x12I\n\ttokensAVS\x18\x06 \x01(\tB+\xc8\xde\x1f\x00\xda\xde\x1f\x15\x63osmossdk.io/math.Int\xd2\xb4-\ncosmos.IntR\ttokensAVS\x12\\\n\x10\x64\x65legator_shares\x18\x07 \x01(\tB1\xc8\xde\x1f\x00\xda\xde\x1f\x1b\x63osmossdk.io/math.LegacyDec\xd2\xb4-\ncosmos.DecR\x0f\x64\x65legatorShares\x12L\n\x0b\x64\x65scription\x18\x08 \x01(\x0b\x32\x1f.zrchain.validation.DescriptionB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x0b\x64\x65scription\x12)\n\x10unbonding_height\x18\t \x01(\x03R\x0funbondingHeight\x12P\n\x0eunbonding_time\x18\n \x01(\x0b\x32\x1a.google.protobuf.TimestampB\r\xc8\xde\x1f\x00\x90\xdf\x1f\x01\xa8\xe7\xb0*\x01R\runbondingTime\x12I\n\ncommission\x18\x0b \x01(\x0b\x32\x1e.zrchain.validation.CommissionB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\ncommission\x12[\n\x13min_self_delegation\x18\x0c \x01(\tB+\xc8\xde\x1f\x00\xda\xde\x1f\x15\x63osmossdk.io/math.Int\xd2\xb4-\ncosmos.IntR\x11minSelfDelegation\x12<\n\x1bunbonding_on_hold_ref_count\x18\r \x01(\x03R\x17unbondingOnHoldRefCount\x12#\n\runbonding_ids\x18\x0e \x03(\x04R\x0cunbondingIds:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\x93\x01\n\x10HistoricalInfoHV\x12;\n\x06header\x18\x01 \x01(\x0b\x32\x18.tendermint.types.HeaderB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06header\x12\x42\n\x06valset\x18\x02 \x03(\x0b\x32\x1f.zrchain.validation.ValidatorHVB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06valset\"\xed\x02\n\nSlashEvent\x12 \n\x0b\x62lockHeight\x18\x01 \x01(\x03R\x0b\x62lockHeight\x12$\n\rvalidatorAddr\x18\x02 \x01(\tR\rvalidatorAddr\x12_\n\x11percentageSlashed\x18\x03 \x01(\tB1\xc8\xde\x1f\x00\xda\xde\x1f\x1b\x63osmossdk.io/math.LegacyDec\xd2\xb4-\ncosmos.DecR\x11percentageSlashed\x12]\n\x13tokensSlashedNative\x18\x04 \x01(\tB+\xc8\xde\x1f\x00\xda\xde\x1f\x15\x63osmossdk.io/math.Int\xd2\xb4-\ncosmos.IntR\x13tokensSlashedNative\x12W\n\x10tokensSlashedAVS\x18\x05 \x01(\tB+\xc8\xde\x1f\x00\xda\xde\x1f\x15\x63osmossdk.io/math.Int\xd2\xb4-\ncosmos.IntR\x10tokensSlashedAVS\"\xc9\x01\n\x08HVParams\x12Y\n\x0e\x41VSRewardsRate\x18\x01 \x01(\tB1\xc8\xde\x1f\x00\xda\xde\x1f\x1b\x63osmossdk.io/math.LegacyDec\xd2\xb4-\ncosmos.DecR\x0e\x41VSRewardsRate\x12\x1c\n\tBlockTime\x18\x02 \x01(\x03R\tBlockTime\x12\x44\n\x0cZenBTCParams\x18\x03 \x01(\x0b\x32 .zrchain.validation.ZenBTCParamsR\x0cZenBTCParams\"\xae\x01\n\x0cZenBTCParams\x12\x34\n\x15zenBTCEthContractAddr\x18\x01 \x01(\tR\x15zenBTCEthContractAddr\x12:\n\x18zenBTCDepositKeyringAddr\x18\x02 \x01(\tR\x18zenBTCDepositKeyringAddr\x12,\n\x11zenBTCMinterKeyID\x18\x03 \x01(\x04R\x11zenBTCMinterKeyID\"\x82\x01\n\x0eValidationInfo\x12\x32\n\x15non_voting_validators\x18\x01 \x03(\tR\x13nonVotingValidators\x12<\n\x1amismatched_vote_extensions\x18\x02 \x03(\tR\x18mismatchedVoteExtensions\"x\n\x0eWithdrawalInfo\x12-\n\x12withdrawal_address\x18\x01 \x01(\tR\x11withdrawalAddress\x12\x16\n\x06\x61mount\x18\x02 \x01(\x04R\x06\x61mount\x12\x1f\n\x0bretry_count\x18\x03 \x01(\rR\nretryCountB>Z<github.com/zenrocklabs/zenrock/zrchain/v4/x/validation/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.validation.hybrid_validation_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z<github.com/zenrocklabs/zenrock/zrchain/v4/x/validation/types'
  _globals['_VALIDATORHV'].fields_by_name['operator_address']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['operator_address']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_VALIDATORHV'].fields_by_name['consensus_pubkey']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['consensus_pubkey']._serialized_options = b'\312\264-\024cosmos.crypto.PubKey'
  _globals['_VALIDATORHV'].fields_by_name['tokensNative']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['tokensNative']._serialized_options = b'\310\336\037\000\332\336\037\025cosmossdk.io/math.Int\322\264-\ncosmos.Int'
  _globals['_VALIDATORHV'].fields_by_name['tokensAVS']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['tokensAVS']._serialized_options = b'\310\336\037\000\332\336\037\025cosmossdk.io/math.Int\322\264-\ncosmos.Int'
  _globals['_VALIDATORHV'].fields_by_name['delegator_shares']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['delegator_shares']._serialized_options = b'\310\336\037\000\332\336\037\033cosmossdk.io/math.LegacyDec\322\264-\ncosmos.Dec'
  _globals['_VALIDATORHV'].fields_by_name['description']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['description']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_VALIDATORHV'].fields_by_name['unbonding_time']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['unbonding_time']._serialized_options = b'\310\336\037\000\220\337\037\001\250\347\260*\001'
  _globals['_VALIDATORHV'].fields_by_name['commission']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['commission']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_VALIDATORHV'].fields_by_name['min_self_delegation']._loaded_options = None
  _globals['_VALIDATORHV'].fields_by_name['min_self_delegation']._serialized_options = b'\310\336\037\000\332\336\037\025cosmossdk.io/math.Int\322\264-\ncosmos.Int'
  _globals['_VALIDATORHV']._loaded_options = None
  _globals['_VALIDATORHV']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_HISTORICALINFOHV'].fields_by_name['header']._loaded_options = None
  _globals['_HISTORICALINFOHV'].fields_by_name['header']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_HISTORICALINFOHV'].fields_by_name['valset']._loaded_options = None
  _globals['_HISTORICALINFOHV'].fields_by_name['valset']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_SLASHEVENT'].fields_by_name['percentageSlashed']._loaded_options = None
  _globals['_SLASHEVENT'].fields_by_name['percentageSlashed']._serialized_options = b'\310\336\037\000\332\336\037\033cosmossdk.io/math.LegacyDec\322\264-\ncosmos.Dec'
  _globals['_SLASHEVENT'].fields_by_name['tokensSlashedNative']._loaded_options = None
  _globals['_SLASHEVENT'].fields_by_name['tokensSlashedNative']._serialized_options = b'\310\336\037\000\332\336\037\025cosmossdk.io/math.Int\322\264-\ncosmos.Int'
  _globals['_SLASHEVENT'].fields_by_name['tokensSlashedAVS']._loaded_options = None
  _globals['_SLASHEVENT'].fields_by_name['tokensSlashedAVS']._serialized_options = b'\310\336\037\000\332\336\037\025cosmossdk.io/math.Int\322\264-\ncosmos.Int'
  _globals['_HVPARAMS'].fields_by_name['AVSRewardsRate']._loaded_options = None
  _globals['_HVPARAMS'].fields_by_name['AVSRewardsRate']._serialized_options = b'\310\336\037\000\332\336\037\033cosmossdk.io/math.LegacyDec\322\264-\ncosmos.Dec'
  _globals['_VALIDATORHV']._serialized_start=259
  _globals['_VALIDATORHV']._serialized_end=1242
  _globals['_HISTORICALINFOHV']._serialized_start=1245
  _globals['_HISTORICALINFOHV']._serialized_end=1392
  _globals['_SLASHEVENT']._serialized_start=1395
  _globals['_SLASHEVENT']._serialized_end=1760
  _globals['_HVPARAMS']._serialized_start=1763
  _globals['_HVPARAMS']._serialized_end=1964
  _globals['_ZENBTCPARAMS']._serialized_start=1967
  _globals['_ZENBTCPARAMS']._serialized_end=2141
  _globals['_VALIDATIONINFO']._serialized_start=2144
  _globals['_VALIDATIONINFO']._serialized_end=2274
  _globals['_WITHDRAWALINFO']._serialized_start=2276
  _globals['_WITHDRAWALINFO']._serialized_end=2396
# @@protoc_insertion_point(module_scope)
