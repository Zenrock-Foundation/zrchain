# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/distribution/v1beta1/genesis.proto
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
    'cosmos/distribution/v1beta1/genesis.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from cosmos.base.v1beta1 import coin_pb2 as cosmos_dot_base_dot_v1beta1_dot_coin__pb2
from cosmos.distribution.v1beta1 import distribution_pb2 as cosmos_dot_distribution_dot_v1beta1_dot_distribution__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from amino import amino_pb2 as amino_dot_amino__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)cosmos/distribution/v1beta1/genesis.proto\x12\x1b\x63osmos.distribution.v1beta1\x1a\x14gogoproto/gogo.proto\x1a\x1e\x63osmos/base/v1beta1/coin.proto\x1a.cosmos/distribution/v1beta1/distribution.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x11\x61mino/amino.proto\"\xad\x01\n\x15\x44\x65legatorWithdrawInfo\x12\x45\n\x11\x64\x65legator_address\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10\x64\x65legatorAddress\x12\x43\n\x10withdraw_address\x18\x02 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x0fwithdrawAddress:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\x87\x02\n!ValidatorOutstandingRewardsRecord\x12N\n\x11validator_address\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\x10validatorAddress\x12\x87\x01\n\x13outstanding_rewards\x18\x02 \x03(\x0b\x32\x1c.cosmos.base.v1beta1.DecCoinB8\xc8\xde\x1f\x00\xaa\xdf\x1f+github.com/cosmos/cosmos-sdk/types.DecCoins\xa8\xe7\xb0*\x01R\x12outstandingRewards:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xea\x01\n$ValidatorAccumulatedCommissionRecord\x12N\n\x11validator_address\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\x10validatorAddress\x12h\n\x0b\x61\x63\x63umulated\x18\x02 \x01(\x0b\x32;.cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x0b\x61\x63\x63umulated:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xf2\x01\n ValidatorHistoricalRewardsRecord\x12N\n\x11validator_address\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\x10validatorAddress\x12\x16\n\x06period\x18\x02 \x01(\x04R\x06period\x12\\\n\x07rewards\x18\x03 \x01(\x0b\x32\x37.cosmos.distribution.v1beta1.ValidatorHistoricalRewardsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x07rewards:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xd4\x01\n\x1dValidatorCurrentRewardsRecord\x12N\n\x11validator_address\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\x10validatorAddress\x12Y\n\x07rewards\x18\x02 \x01(\x0b\x32\x34.cosmos.distribution.v1beta1.ValidatorCurrentRewardsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x07rewards:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xa2\x02\n\x1b\x44\x65legatorStartingInfoRecord\x12\x45\n\x11\x64\x65legator_address\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10\x64\x65legatorAddress\x12N\n\x11validator_address\x18\x02 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\x10validatorAddress\x12\x62\n\rstarting_info\x18\x03 \x01(\x0b\x32\x32.cosmos.distribution.v1beta1.DelegatorStartingInfoB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x0cstartingInfo:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\x96\x02\n\x19ValidatorSlashEventRecord\x12N\n\x11validator_address\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\x10validatorAddress\x12\x16\n\x06height\x18\x02 \x01(\x04R\x06height\x12\x16\n\x06period\x18\x03 \x01(\x04R\x06period\x12o\n\x15validator_slash_event\x18\x04 \x01(\x0b\x32\x30.cosmos.distribution.v1beta1.ValidatorSlashEventB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x13validatorSlashEvent:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\x8c\t\n\x0cGenesisState\x12\x46\n\x06params\x18\x01 \x01(\x0b\x32#.cosmos.distribution.v1beta1.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06params\x12J\n\x08\x66\x65\x65_pool\x18\x02 \x01(\x0b\x32$.cosmos.distribution.v1beta1.FeePoolB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x07\x66\x65\x65Pool\x12w\n\x18\x64\x65legator_withdraw_infos\x18\x03 \x03(\x0b\x32\x32.cosmos.distribution.v1beta1.DelegatorWithdrawInfoB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x16\x64\x65legatorWithdrawInfos\x12\x45\n\x11previous_proposer\x18\x04 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10previousProposer\x12z\n\x13outstanding_rewards\x18\x05 \x03(\x0b\x32>.cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecordB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x12outstandingRewards\x12\x98\x01\n!validator_accumulated_commissions\x18\x06 \x03(\x0b\x32\x41.cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecordB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x1fvalidatorAccumulatedCommissions\x12\x8a\x01\n\x1cvalidator_historical_rewards\x18\x07 \x03(\x0b\x32=.cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecordB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x1avalidatorHistoricalRewards\x12\x81\x01\n\x19validator_current_rewards\x18\x08 \x03(\x0b\x32:.cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecordB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x17validatorCurrentRewards\x12}\n\x18\x64\x65legator_starting_infos\x18\t \x03(\x0b\x32\x38.cosmos.distribution.v1beta1.DelegatorStartingInfoRecordB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x16\x64\x65legatorStartingInfos\x12w\n\x16validator_slash_events\x18\n \x03(\x0b\x32\x36.cosmos.distribution.v1beta1.ValidatorSlashEventRecordB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x14validatorSlashEvents:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\x42\'Z!cosmossdk.io/x/distribution/types\xa8\xe2\x1e\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.distribution.v1beta1.genesis_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z!cosmossdk.io/x/distribution/types\250\342\036\001'
  _globals['_DELEGATORWITHDRAWINFO'].fields_by_name['delegator_address']._loaded_options = None
  _globals['_DELEGATORWITHDRAWINFO'].fields_by_name['delegator_address']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_DELEGATORWITHDRAWINFO'].fields_by_name['withdraw_address']._loaded_options = None
  _globals['_DELEGATORWITHDRAWINFO'].fields_by_name['withdraw_address']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_DELEGATORWITHDRAWINFO']._loaded_options = None
  _globals['_DELEGATORWITHDRAWINFO']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD'].fields_by_name['validator_address']._loaded_options = None
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD'].fields_by_name['validator_address']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD'].fields_by_name['outstanding_rewards']._loaded_options = None
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD'].fields_by_name['outstanding_rewards']._serialized_options = b'\310\336\037\000\252\337\037+github.com/cosmos/cosmos-sdk/types.DecCoins\250\347\260*\001'
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD']._loaded_options = None
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD'].fields_by_name['validator_address']._loaded_options = None
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD'].fields_by_name['validator_address']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD'].fields_by_name['accumulated']._loaded_options = None
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD'].fields_by_name['accumulated']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD']._loaded_options = None
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_VALIDATORHISTORICALREWARDSRECORD'].fields_by_name['validator_address']._loaded_options = None
  _globals['_VALIDATORHISTORICALREWARDSRECORD'].fields_by_name['validator_address']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_VALIDATORHISTORICALREWARDSRECORD'].fields_by_name['rewards']._loaded_options = None
  _globals['_VALIDATORHISTORICALREWARDSRECORD'].fields_by_name['rewards']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_VALIDATORHISTORICALREWARDSRECORD']._loaded_options = None
  _globals['_VALIDATORHISTORICALREWARDSRECORD']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_VALIDATORCURRENTREWARDSRECORD'].fields_by_name['validator_address']._loaded_options = None
  _globals['_VALIDATORCURRENTREWARDSRECORD'].fields_by_name['validator_address']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_VALIDATORCURRENTREWARDSRECORD'].fields_by_name['rewards']._loaded_options = None
  _globals['_VALIDATORCURRENTREWARDSRECORD'].fields_by_name['rewards']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_VALIDATORCURRENTREWARDSRECORD']._loaded_options = None
  _globals['_VALIDATORCURRENTREWARDSRECORD']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_DELEGATORSTARTINGINFORECORD'].fields_by_name['delegator_address']._loaded_options = None
  _globals['_DELEGATORSTARTINGINFORECORD'].fields_by_name['delegator_address']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_DELEGATORSTARTINGINFORECORD'].fields_by_name['validator_address']._loaded_options = None
  _globals['_DELEGATORSTARTINGINFORECORD'].fields_by_name['validator_address']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_DELEGATORSTARTINGINFORECORD'].fields_by_name['starting_info']._loaded_options = None
  _globals['_DELEGATORSTARTINGINFORECORD'].fields_by_name['starting_info']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_DELEGATORSTARTINGINFORECORD']._loaded_options = None
  _globals['_DELEGATORSTARTINGINFORECORD']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_VALIDATORSLASHEVENTRECORD'].fields_by_name['validator_address']._loaded_options = None
  _globals['_VALIDATORSLASHEVENTRECORD'].fields_by_name['validator_address']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_VALIDATORSLASHEVENTRECORD'].fields_by_name['validator_slash_event']._loaded_options = None
  _globals['_VALIDATORSLASHEVENTRECORD'].fields_by_name['validator_slash_event']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_VALIDATORSLASHEVENTRECORD']._loaded_options = None
  _globals['_VALIDATORSLASHEVENTRECORD']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_GENESISSTATE'].fields_by_name['params']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['params']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['fee_pool']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['fee_pool']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['delegator_withdraw_infos']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['delegator_withdraw_infos']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['previous_proposer']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['previous_proposer']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_GENESISSTATE'].fields_by_name['outstanding_rewards']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['outstanding_rewards']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['validator_accumulated_commissions']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['validator_accumulated_commissions']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['validator_historical_rewards']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['validator_historical_rewards']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['validator_current_rewards']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['validator_current_rewards']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['delegator_starting_infos']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['delegator_starting_infos']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE'].fields_by_name['validator_slash_events']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['validator_slash_events']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_GENESISSTATE']._loaded_options = None
  _globals['_GENESISSTATE']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_DELEGATORWITHDRAWINFO']._serialized_start=223
  _globals['_DELEGATORWITHDRAWINFO']._serialized_end=396
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD']._serialized_start=399
  _globals['_VALIDATOROUTSTANDINGREWARDSRECORD']._serialized_end=662
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD']._serialized_start=665
  _globals['_VALIDATORACCUMULATEDCOMMISSIONRECORD']._serialized_end=899
  _globals['_VALIDATORHISTORICALREWARDSRECORD']._serialized_start=902
  _globals['_VALIDATORHISTORICALREWARDSRECORD']._serialized_end=1144
  _globals['_VALIDATORCURRENTREWARDSRECORD']._serialized_start=1147
  _globals['_VALIDATORCURRENTREWARDSRECORD']._serialized_end=1359
  _globals['_DELEGATORSTARTINGINFORECORD']._serialized_start=1362
  _globals['_DELEGATORSTARTINGINFORECORD']._serialized_end=1652
  _globals['_VALIDATORSLASHEVENTRECORD']._serialized_start=1655
  _globals['_VALIDATORSLASHEVENTRECORD']._serialized_end=1933
  _globals['_GENESISSTATE']._serialized_start=1936
  _globals['_GENESISSTATE']._serialized_end=3100
# @@protoc_insertion_point(module_scope)
