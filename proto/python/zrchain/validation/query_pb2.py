# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/validation/query.proto
# Protobuf Python Version: 5.29.2
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
    2,
    '',
    'zrchain/validation/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from cosmos.base.query.v1beta1 import pagination_pb2 as cosmos_dot_base_dot_query_dot_v1beta1_dot_pagination__pb2
from cosmos.query.v1 import query_pb2 as cosmos_dot_query_dot_v1_dot_query__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from zrchain.validation import hybrid_validation_pb2 as zrchain_dot_validation_dot_hybrid__validation__pb2
from zrchain.validation import staking_pb2 as zrchain_dot_validation_dot_staking__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ezrchain/validation/query.proto\x12\x12zrchain.validation\x1a\x11\x61mino/amino.proto\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a\x1b\x63osmos/query/v1/query.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x14gogoproto/gogo.proto\x1a\x1cgoogle/api/annotations.proto\x1a*zrchain/validation/hybrid_validation.proto\x1a zrchain/validation/staking.proto\"x\n\x16QueryValidatorsRequest\x12\x16\n\x06status\x18\x01 \x01(\tR\x06status\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\xae\x01\n\x17QueryValidatorsResponse\x12J\n\nvalidators\x18\x01 \x03(\x0b\x32\x1f.zrchain.validation.ValidatorHVB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\nvalidators\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"a\n\x15QueryValidatorRequest\x12H\n\x0evalidator_addr\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr\"b\n\x16QueryValidatorResponse\x12H\n\tvalidator\x18\x01 \x01(\x0b\x32\x1f.zrchain.validation.ValidatorHVB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\tvalidator\"\xb4\x01\n QueryValidatorDelegationsRequest\x12H\n\x0evalidator_addr\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\xe9\x01\n!QueryValidatorDelegationsResponse\x12{\n\x14\x64\x65legation_responses\x18\x01 \x03(\x0b\x32&.zrchain.validation.DelegationResponseB \xc8\xde\x1f\x00\xaa\xdf\x1f\x13\x44\x65legationResponses\xa8\xe7\xb0*\x01R\x13\x64\x65legationResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xbd\x01\n)QueryValidatorUnbondingDelegationsRequest\x12H\n\x0evalidator_addr\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\xda\x01\n*QueryValidatorUnbondingDelegationsResponse\x12\x63\n\x13unbonding_responses\x18\x01 \x03(\x0b\x32\'.zrchain.validation.UnbondingDelegationB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x12unbondingResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xad\x01\n\x16QueryDelegationRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12H\n\x0evalidator_addr\x18\x02 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"r\n\x17QueryDelegationResponse\x12W\n\x13\x64\x65legation_response\x18\x01 \x01(\x0b\x32&.zrchain.validation.DelegationResponseR\x12\x64\x65legationResponse\"\xb6\x01\n\x1fQueryUnbondingDelegationRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12H\n\x0evalidator_addr\x18\x02 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"n\n QueryUnbondingDelegationResponse\x12J\n\x06unbond\x18\x01 \x01(\x0b\x32\'.zrchain.validation.UnbondingDelegationB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06unbond\"\xb5\x01\n QueryDelegatorDelegationsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xd2\x01\n!QueryDelegatorDelegationsResponse\x12\x64\n\x14\x64\x65legation_responses\x18\x01 \x03(\x0b\x32&.zrchain.validation.DelegationResponseB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x13\x64\x65legationResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xbe\x01\n)QueryDelegatorUnbondingDelegationsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xda\x01\n*QueryDelegatorUnbondingDelegationsResponse\x12\x63\n\x13unbonding_responses\x18\x01 \x03(\x0b\x32\'.zrchain.validation.UnbondingDelegationB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x12unbondingResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xbe\x02\n\x19QueryRedelegationsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\x12src_validator_addr\x18\x02 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10srcValidatorAddr\x12\x46\n\x12\x64st_validator_addr\x18\x03 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10\x64stValidatorAddr\x12\x46\n\npagination\x18\x04 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xd1\x01\n\x1aQueryRedelegationsResponse\x12j\n\x16redelegation_responses\x18\x01 \x03(\x0b\x32(.zrchain.validation.RedelegationResponseB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x15redelegationResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xb4\x01\n\x1fQueryDelegatorValidatorsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xb7\x01\n QueryDelegatorValidatorsResponse\x12J\n\nvalidators\x18\x01 \x03(\x0b\x32\x1f.zrchain.validation.ValidatorHVB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\nvalidators\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xb5\x01\n\x1eQueryDelegatorValidatorRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12H\n\x0evalidator_addr\x18\x02 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"k\n\x1fQueryDelegatorValidatorResponse\x12H\n\tvalidator\x18\x01 \x01(\x0b\x32\x1f.zrchain.validation.ValidatorHVB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\tvalidator\"4\n\x1aQueryHistoricalInfoRequest\x12\x16\n\x06height\x18\x01 \x01(\x03R\x06height\"W\n\x1bQueryHistoricalInfoResponse\x12\x38\n\x04hist\x18\x01 \x01(\x0b\x32$.zrchain.validation.HistoricalInfoHVR\x04hist\"\x12\n\x10QueryPoolRequest\"L\n\x11QueryPoolResponse\x12\x37\n\x04pool\x18\x01 \x01(\x0b\x32\x18.zrchain.validation.PoolB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x04pool\"\x14\n\x12QueryParamsRequest\"\x99\x01\n\x13QueryParamsResponse\x12=\n\x06Params\x18\x01 \x01(\x0b\x32\x1a.zrchain.validation.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06Params\x12\x43\n\x08HVParams\x18\x02 \x01(\x0b\x32\x1c.zrchain.validation.HVParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x08HVParams\"\x13\n\x11QueryPowerRequest\"\xb8\x01\n\x0eValidatorPower\x12#\n\roperator_addr\x18\x01 \x01(\tR\x0coperatorAddr\x12\x1b\n\tcons_addr\x18\x02 \x01(\tR\x08\x63onsAddr\x12\x14\n\x05power\x18\x03 \x01(\x03R\x05power\x12\x16\n\x06jailed\x18\x04 \x01(\x08R\x06jailed\x12\x36\n\x06status\x18\x05 \x01(\x0e\x32\x1e.zrchain.validation.BondStatusR\x06status\"\x9a\x01\n\x12QueryPowerResponse\x12K\n\x0fvalidator_power\x18\x01 \x03(\x0b\x32\".zrchain.validation.ValidatorPowerR\x0evalidatorPower\x12\x1f\n\x0btotal_power\x18\x02 \x01(\x03R\ntotalPower\x12\x16\n\x06height\x18\x03 \x01(\x03R\x06height\"%\n#QueryPendingMintTransactionsRequest\"\x96\x01\n$QueryPendingMintTransactionsResponse\x12n\n\x19pending_mint_transactions\x18\x01 \x03(\x0b\x32\x32.zrchain.validation.PendingMintTransactionResponseR\x17pendingMintTransactions\"\xd0\x01\n\x1ePendingMintTransactionResponse\x12\x19\n\x08\x63hain_id\x18\x01 \x01(\x04R\x07\x63hainId\x12\x1d\n\nchain_type\x18\x02 \x01(\tR\tchainType\x12+\n\x11recipient_address\x18\x03 \x01(\tR\x10recipientAddress\x12\x16\n\x06\x61mount\x18\x04 \x01(\x04R\x06\x61mount\x12\x18\n\x07\x63reator\x18\x05 \x01(\tR\x07\x63reator\x12\x15\n\x06key_id\x18\x06 \x01(\x04R\x05keyId2\xa9\x18\n\x05Query\x12\x96\x01\n\nValidators\x12*.zrchain.validation.QueryValidatorsRequest\x1a+.zrchain.validation.QueryValidatorsResponse\"/\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02$\x12\"/cosmos/staking/v1beta1/validators\x12\xa4\x01\n\tValidator\x12).zrchain.validation.QueryValidatorRequest\x1a*.zrchain.validation.QueryValidatorResponse\"@\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x35\x12\x33/cosmos/staking/v1beta1/validators/{validator_addr}\x12\xd1\x01\n\x14ValidatorDelegations\x12\x34.zrchain.validation.QueryValidatorDelegationsRequest\x1a\x35.zrchain.validation.QueryValidatorDelegationsResponse\"L\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x41\x12?/cosmos/staking/v1beta1/validators/{validator_addr}/delegations\x12\xf6\x01\n\x1dValidatorUnbondingDelegations\x12=.zrchain.validation.QueryValidatorUnbondingDelegationsRequest\x1a>.zrchain.validation.QueryValidatorUnbondingDelegationsResponse\"V\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02K\x12I/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations\x12\xc4\x01\n\nDelegation\x12*.zrchain.validation.QueryDelegationRequest\x1a+.zrchain.validation.QueryDelegationResponse\"]\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02R\x12P/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}\x12\xf4\x01\n\x13UnbondingDelegation\x12\x33.zrchain.validation.QueryUnbondingDelegationRequest\x1a\x34.zrchain.validation.QueryUnbondingDelegationResponse\"r\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02g\x12\x65/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation\x12\xc6\x01\n\x14\x44\x65legatorDelegations\x12\x34.zrchain.validation.QueryDelegatorDelegationsRequest\x1a\x35.zrchain.validation.QueryDelegatorDelegationsResponse\"A\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x36\x12\x34/cosmos/staking/v1beta1/delegations/{delegator_addr}\x12\xf6\x01\n\x1d\x44\x65legatorUnbondingDelegations\x12=.zrchain.validation.QueryDelegatorUnbondingDelegationsRequest\x1a>.zrchain.validation.QueryDelegatorUnbondingDelegationsResponse\"V\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02K\x12I/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations\x12\xbe\x01\n\rRedelegations\x12-.zrchain.validation.QueryRedelegationsRequest\x1a..zrchain.validation.QueryRedelegationsResponse\"N\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x43\x12\x41/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations\x12\xcd\x01\n\x13\x44\x65legatorValidators\x12\x33.zrchain.validation.QueryDelegatorValidatorsRequest\x1a\x34.zrchain.validation.QueryDelegatorValidatorsResponse\"K\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02@\x12>/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators\x12\xdb\x01\n\x12\x44\x65legatorValidator\x12\x32.zrchain.validation.QueryDelegatorValidatorRequest\x1a\x33.zrchain.validation.QueryDelegatorValidatorResponse\"\\\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02Q\x12O/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}\x12\xb0\x01\n\x0eHistoricalInfo\x12..zrchain.validation.QueryHistoricalInfoRequest\x1a/.zrchain.validation.QueryHistoricalInfoResponse\"=\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x32\x12\x30/cosmos/staking/v1beta1/historical_info/{height}\x12~\n\x04Pool\x12$.zrchain.validation.QueryPoolRequest\x1a%.zrchain.validation.QueryPoolResponse\")\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x1e\x12\x1c/cosmos/staking/v1beta1/pool\x12\x86\x01\n\x06Params\x12&.zrchain.validation.QueryParamsRequest\x1a\'.zrchain.validation.QueryParamsResponse\"+\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02 \x12\x1e/cosmos/staking/v1beta1/params\x12\x95\x01\n\x0eValidatorPower\x12%.zrchain.validation.QueryPowerRequest\x1a&.zrchain.validation.QueryPowerResponse\"4\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02)\x12\'/cosmos/staking/v1beta1/validator_power\x12\xcf\x01\n\x1aGetPendingMintTransactions\x12\x37.zrchain.validation.QueryPendingMintTransactionsRequest\x1a\x38.zrchain.validation.QueryPendingMintTransactionsResponse\">\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x33\x12\x31/cosmos/staking/v1beta1/pending_mint_transactionsB=Z;github.com/Zenrock-Foundation/zrchain/v5/x/validation/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.validation.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z;github.com/Zenrock-Foundation/zrchain/v5/x/validation/types'
  _globals['_QUERYVALIDATORSRESPONSE'].fields_by_name['validators']._loaded_options = None
  _globals['_QUERYVALIDATORSRESPONSE'].fields_by_name['validators']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYVALIDATORREQUEST'].fields_by_name['validator_addr']._loaded_options = None
  _globals['_QUERYVALIDATORREQUEST'].fields_by_name['validator_addr']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_QUERYVALIDATORRESPONSE'].fields_by_name['validator']._loaded_options = None
  _globals['_QUERYVALIDATORRESPONSE'].fields_by_name['validator']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYVALIDATORDELEGATIONSREQUEST'].fields_by_name['validator_addr']._loaded_options = None
  _globals['_QUERYVALIDATORDELEGATIONSREQUEST'].fields_by_name['validator_addr']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_QUERYVALIDATORDELEGATIONSRESPONSE'].fields_by_name['delegation_responses']._loaded_options = None
  _globals['_QUERYVALIDATORDELEGATIONSRESPONSE'].fields_by_name['delegation_responses']._serialized_options = b'\310\336\037\000\252\337\037\023DelegationResponses\250\347\260*\001'
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSREQUEST'].fields_by_name['validator_addr']._loaded_options = None
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSREQUEST'].fields_by_name['validator_addr']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSRESPONSE'].fields_by_name['unbonding_responses']._loaded_options = None
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSRESPONSE'].fields_by_name['unbonding_responses']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYDELEGATIONREQUEST'].fields_by_name['delegator_addr']._loaded_options = None
  _globals['_QUERYDELEGATIONREQUEST'].fields_by_name['delegator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYDELEGATIONREQUEST'].fields_by_name['validator_addr']._loaded_options = None
  _globals['_QUERYDELEGATIONREQUEST'].fields_by_name['validator_addr']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_QUERYDELEGATIONREQUEST']._loaded_options = None
  _globals['_QUERYDELEGATIONREQUEST']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST'].fields_by_name['delegator_addr']._loaded_options = None
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST'].fields_by_name['delegator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST'].fields_by_name['validator_addr']._loaded_options = None
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST'].fields_by_name['validator_addr']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST']._loaded_options = None
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_QUERYUNBONDINGDELEGATIONRESPONSE'].fields_by_name['unbond']._loaded_options = None
  _globals['_QUERYUNBONDINGDELEGATIONRESPONSE'].fields_by_name['unbond']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST'].fields_by_name['delegator_addr']._loaded_options = None
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST'].fields_by_name['delegator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST']._loaded_options = None
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_QUERYDELEGATORDELEGATIONSRESPONSE'].fields_by_name['delegation_responses']._loaded_options = None
  _globals['_QUERYDELEGATORDELEGATIONSRESPONSE'].fields_by_name['delegation_responses']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST'].fields_by_name['delegator_addr']._loaded_options = None
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST'].fields_by_name['delegator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST']._loaded_options = None
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSRESPONSE'].fields_by_name['unbonding_responses']._loaded_options = None
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSRESPONSE'].fields_by_name['unbonding_responses']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYREDELEGATIONSREQUEST'].fields_by_name['delegator_addr']._loaded_options = None
  _globals['_QUERYREDELEGATIONSREQUEST'].fields_by_name['delegator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYREDELEGATIONSREQUEST'].fields_by_name['src_validator_addr']._loaded_options = None
  _globals['_QUERYREDELEGATIONSREQUEST'].fields_by_name['src_validator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYREDELEGATIONSREQUEST'].fields_by_name['dst_validator_addr']._loaded_options = None
  _globals['_QUERYREDELEGATIONSREQUEST'].fields_by_name['dst_validator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYREDELEGATIONSREQUEST']._loaded_options = None
  _globals['_QUERYREDELEGATIONSREQUEST']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_QUERYREDELEGATIONSRESPONSE'].fields_by_name['redelegation_responses']._loaded_options = None
  _globals['_QUERYREDELEGATIONSRESPONSE'].fields_by_name['redelegation_responses']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYDELEGATORVALIDATORSREQUEST'].fields_by_name['delegator_addr']._loaded_options = None
  _globals['_QUERYDELEGATORVALIDATORSREQUEST'].fields_by_name['delegator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYDELEGATORVALIDATORSREQUEST']._loaded_options = None
  _globals['_QUERYDELEGATORVALIDATORSREQUEST']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_QUERYDELEGATORVALIDATORSRESPONSE'].fields_by_name['validators']._loaded_options = None
  _globals['_QUERYDELEGATORVALIDATORSRESPONSE'].fields_by_name['validators']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYDELEGATORVALIDATORREQUEST'].fields_by_name['delegator_addr']._loaded_options = None
  _globals['_QUERYDELEGATORVALIDATORREQUEST'].fields_by_name['delegator_addr']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_QUERYDELEGATORVALIDATORREQUEST'].fields_by_name['validator_addr']._loaded_options = None
  _globals['_QUERYDELEGATORVALIDATORREQUEST'].fields_by_name['validator_addr']._serialized_options = b'\322\264-\035cosmos.ValidatorAddressString'
  _globals['_QUERYDELEGATORVALIDATORREQUEST']._loaded_options = None
  _globals['_QUERYDELEGATORVALIDATORREQUEST']._serialized_options = b'\210\240\037\000\350\240\037\000'
  _globals['_QUERYDELEGATORVALIDATORRESPONSE'].fields_by_name['validator']._loaded_options = None
  _globals['_QUERYDELEGATORVALIDATORRESPONSE'].fields_by_name['validator']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYPOOLRESPONSE'].fields_by_name['pool']._loaded_options = None
  _globals['_QUERYPOOLRESPONSE'].fields_by_name['pool']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['Params']._loaded_options = None
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['Params']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['HVParams']._loaded_options = None
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['HVParams']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERY'].methods_by_name['Validators']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Validators']._serialized_options = b'\210\347\260*\001\202\323\344\223\002$\022\"/cosmos/staking/v1beta1/validators'
  _globals['_QUERY'].methods_by_name['Validator']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Validator']._serialized_options = b'\210\347\260*\001\202\323\344\223\0025\0223/cosmos/staking/v1beta1/validators/{validator_addr}'
  _globals['_QUERY'].methods_by_name['ValidatorDelegations']._loaded_options = None
  _globals['_QUERY'].methods_by_name['ValidatorDelegations']._serialized_options = b'\210\347\260*\001\202\323\344\223\002A\022?/cosmos/staking/v1beta1/validators/{validator_addr}/delegations'
  _globals['_QUERY'].methods_by_name['ValidatorUnbondingDelegations']._loaded_options = None
  _globals['_QUERY'].methods_by_name['ValidatorUnbondingDelegations']._serialized_options = b'\210\347\260*\001\202\323\344\223\002K\022I/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations'
  _globals['_QUERY'].methods_by_name['Delegation']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Delegation']._serialized_options = b'\210\347\260*\001\202\323\344\223\002R\022P/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}'
  _globals['_QUERY'].methods_by_name['UnbondingDelegation']._loaded_options = None
  _globals['_QUERY'].methods_by_name['UnbondingDelegation']._serialized_options = b'\210\347\260*\001\202\323\344\223\002g\022e/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation'
  _globals['_QUERY'].methods_by_name['DelegatorDelegations']._loaded_options = None
  _globals['_QUERY'].methods_by_name['DelegatorDelegations']._serialized_options = b'\210\347\260*\001\202\323\344\223\0026\0224/cosmos/staking/v1beta1/delegations/{delegator_addr}'
  _globals['_QUERY'].methods_by_name['DelegatorUnbondingDelegations']._loaded_options = None
  _globals['_QUERY'].methods_by_name['DelegatorUnbondingDelegations']._serialized_options = b'\210\347\260*\001\202\323\344\223\002K\022I/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations'
  _globals['_QUERY'].methods_by_name['Redelegations']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Redelegations']._serialized_options = b'\210\347\260*\001\202\323\344\223\002C\022A/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations'
  _globals['_QUERY'].methods_by_name['DelegatorValidators']._loaded_options = None
  _globals['_QUERY'].methods_by_name['DelegatorValidators']._serialized_options = b'\210\347\260*\001\202\323\344\223\002@\022>/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators'
  _globals['_QUERY'].methods_by_name['DelegatorValidator']._loaded_options = None
  _globals['_QUERY'].methods_by_name['DelegatorValidator']._serialized_options = b'\210\347\260*\001\202\323\344\223\002Q\022O/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}'
  _globals['_QUERY'].methods_by_name['HistoricalInfo']._loaded_options = None
  _globals['_QUERY'].methods_by_name['HistoricalInfo']._serialized_options = b'\210\347\260*\001\202\323\344\223\0022\0220/cosmos/staking/v1beta1/historical_info/{height}'
  _globals['_QUERY'].methods_by_name['Pool']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Pool']._serialized_options = b'\210\347\260*\001\202\323\344\223\002\036\022\034/cosmos/staking/v1beta1/pool'
  _globals['_QUERY'].methods_by_name['Params']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Params']._serialized_options = b'\210\347\260*\001\202\323\344\223\002 \022\036/cosmos/staking/v1beta1/params'
  _globals['_QUERY'].methods_by_name['ValidatorPower']._loaded_options = None
  _globals['_QUERY'].methods_by_name['ValidatorPower']._serialized_options = b'\210\347\260*\001\202\323\344\223\002)\022\'/cosmos/staking/v1beta1/validator_power'
  _globals['_QUERY'].methods_by_name['GetPendingMintTransactions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['GetPendingMintTransactions']._serialized_options = b'\210\347\260*\001\202\323\344\223\0023\0221/cosmos/staking/v1beta1/pending_mint_transactions'
  _globals['_QUERYVALIDATORSREQUEST']._serialized_start=303
  _globals['_QUERYVALIDATORSREQUEST']._serialized_end=423
  _globals['_QUERYVALIDATORSRESPONSE']._serialized_start=426
  _globals['_QUERYVALIDATORSRESPONSE']._serialized_end=600
  _globals['_QUERYVALIDATORREQUEST']._serialized_start=602
  _globals['_QUERYVALIDATORREQUEST']._serialized_end=699
  _globals['_QUERYVALIDATORRESPONSE']._serialized_start=701
  _globals['_QUERYVALIDATORRESPONSE']._serialized_end=799
  _globals['_QUERYVALIDATORDELEGATIONSREQUEST']._serialized_start=802
  _globals['_QUERYVALIDATORDELEGATIONSREQUEST']._serialized_end=982
  _globals['_QUERYVALIDATORDELEGATIONSRESPONSE']._serialized_start=985
  _globals['_QUERYVALIDATORDELEGATIONSRESPONSE']._serialized_end=1218
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSREQUEST']._serialized_start=1221
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSREQUEST']._serialized_end=1410
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSRESPONSE']._serialized_start=1413
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSRESPONSE']._serialized_end=1631
  _globals['_QUERYDELEGATIONREQUEST']._serialized_start=1634
  _globals['_QUERYDELEGATIONREQUEST']._serialized_end=1807
  _globals['_QUERYDELEGATIONRESPONSE']._serialized_start=1809
  _globals['_QUERYDELEGATIONRESPONSE']._serialized_end=1923
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST']._serialized_start=1926
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST']._serialized_end=2108
  _globals['_QUERYUNBONDINGDELEGATIONRESPONSE']._serialized_start=2110
  _globals['_QUERYUNBONDINGDELEGATIONRESPONSE']._serialized_end=2220
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST']._serialized_start=2223
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST']._serialized_end=2404
  _globals['_QUERYDELEGATORDELEGATIONSRESPONSE']._serialized_start=2407
  _globals['_QUERYDELEGATORDELEGATIONSRESPONSE']._serialized_end=2617
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST']._serialized_start=2620
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST']._serialized_end=2810
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSRESPONSE']._serialized_start=2813
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSRESPONSE']._serialized_end=3031
  _globals['_QUERYREDELEGATIONSREQUEST']._serialized_start=3034
  _globals['_QUERYREDELEGATIONSREQUEST']._serialized_end=3352
  _globals['_QUERYREDELEGATIONSRESPONSE']._serialized_start=3355
  _globals['_QUERYREDELEGATIONSRESPONSE']._serialized_end=3564
  _globals['_QUERYDELEGATORVALIDATORSREQUEST']._serialized_start=3567
  _globals['_QUERYDELEGATORVALIDATORSREQUEST']._serialized_end=3747
  _globals['_QUERYDELEGATORVALIDATORSRESPONSE']._serialized_start=3750
  _globals['_QUERYDELEGATORVALIDATORSRESPONSE']._serialized_end=3933
  _globals['_QUERYDELEGATORVALIDATORREQUEST']._serialized_start=3936
  _globals['_QUERYDELEGATORVALIDATORREQUEST']._serialized_end=4117
  _globals['_QUERYDELEGATORVALIDATORRESPONSE']._serialized_start=4119
  _globals['_QUERYDELEGATORVALIDATORRESPONSE']._serialized_end=4226
  _globals['_QUERYHISTORICALINFOREQUEST']._serialized_start=4228
  _globals['_QUERYHISTORICALINFOREQUEST']._serialized_end=4280
  _globals['_QUERYHISTORICALINFORESPONSE']._serialized_start=4282
  _globals['_QUERYHISTORICALINFORESPONSE']._serialized_end=4369
  _globals['_QUERYPOOLREQUEST']._serialized_start=4371
  _globals['_QUERYPOOLREQUEST']._serialized_end=4389
  _globals['_QUERYPOOLRESPONSE']._serialized_start=4391
  _globals['_QUERYPOOLRESPONSE']._serialized_end=4467
  _globals['_QUERYPARAMSREQUEST']._serialized_start=4469
  _globals['_QUERYPARAMSREQUEST']._serialized_end=4489
  _globals['_QUERYPARAMSRESPONSE']._serialized_start=4492
  _globals['_QUERYPARAMSRESPONSE']._serialized_end=4645
  _globals['_QUERYPOWERREQUEST']._serialized_start=4647
  _globals['_QUERYPOWERREQUEST']._serialized_end=4666
  _globals['_VALIDATORPOWER']._serialized_start=4669
  _globals['_VALIDATORPOWER']._serialized_end=4853
  _globals['_QUERYPOWERRESPONSE']._serialized_start=4856
  _globals['_QUERYPOWERRESPONSE']._serialized_end=5010
  _globals['_QUERYPENDINGMINTTRANSACTIONSREQUEST']._serialized_start=5012
  _globals['_QUERYPENDINGMINTTRANSACTIONSREQUEST']._serialized_end=5049
  _globals['_QUERYPENDINGMINTTRANSACTIONSRESPONSE']._serialized_start=5052
  _globals['_QUERYPENDINGMINTTRANSACTIONSRESPONSE']._serialized_end=5202
  _globals['_PENDINGMINTTRANSACTIONRESPONSE']._serialized_start=5205
  _globals['_PENDINGMINTTRANSACTIONRESPONSE']._serialized_end=5413
  _globals['_QUERY']._serialized_start=5416
  _globals['_QUERY']._serialized_end=8529
# @@protoc_insertion_point(module_scope)
