# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/staking/v1beta1/query.proto
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
    'cosmos/staking/v1beta1/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from cosmos.base.query.v1beta1 import pagination_pb2 as cosmos_dot_base_dot_query_dot_v1beta1_dot_pagination__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from cosmos.staking.v1beta1 import staking_pb2 as cosmos_dot_staking_dot_v1beta1_dot_staking__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from cosmos.query.v1 import query_pb2 as cosmos_dot_query_dot_v1_dot_query__pb2
from amino import amino_pb2 as amino_dot_amino__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"cosmos/staking/v1beta1/query.proto\x12\x16\x63osmos.staking.v1beta1\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a\x14gogoproto/gogo.proto\x1a\x1cgoogle/api/annotations.proto\x1a$cosmos/staking/v1beta1/staking.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x1b\x63osmos/query/v1/query.proto\x1a\x11\x61mino/amino.proto\"x\n\x16QueryValidatorsRequest\x12\x16\n\x06status\x18\x01 \x01(\tR\x06status\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\xb0\x01\n\x17QueryValidatorsResponse\x12L\n\nvalidators\x18\x01 \x03(\x0b\x32!.cosmos.staking.v1beta1.ValidatorB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\nvalidators\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"a\n\x15QueryValidatorRequest\x12H\n\x0evalidator_addr\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr\"d\n\x16QueryValidatorResponse\x12J\n\tvalidator\x18\x01 \x01(\x0b\x32!.cosmos.staking.v1beta1.ValidatorB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\tvalidator\"\xb4\x01\n QueryValidatorDelegationsRequest\x12H\n\x0evalidator_addr\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\xed\x01\n!QueryValidatorDelegationsResponse\x12\x7f\n\x14\x64\x65legation_responses\x18\x01 \x03(\x0b\x32*.cosmos.staking.v1beta1.DelegationResponseB \xc8\xde\x1f\x00\xaa\xdf\x1f\x13\x44\x65legationResponses\xa8\xe7\xb0*\x01R\x13\x64\x65legationResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xbd\x01\n)QueryValidatorUnbondingDelegationsRequest\x12H\n\x0evalidator_addr\x18\x01 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\xde\x01\n*QueryValidatorUnbondingDelegationsResponse\x12g\n\x13unbonding_responses\x18\x01 \x03(\x0b\x32+.cosmos.staking.v1beta1.UnbondingDelegationB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x12unbondingResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xad\x01\n\x16QueryDelegationRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12H\n\x0evalidator_addr\x18\x02 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"v\n\x17QueryDelegationResponse\x12[\n\x13\x64\x65legation_response\x18\x01 \x01(\x0b\x32*.cosmos.staking.v1beta1.DelegationResponseR\x12\x64\x65legationResponse\"\xb6\x01\n\x1fQueryUnbondingDelegationRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12H\n\x0evalidator_addr\x18\x02 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"r\n QueryUnbondingDelegationResponse\x12N\n\x06unbond\x18\x01 \x01(\x0b\x32+.cosmos.staking.v1beta1.UnbondingDelegationB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06unbond\"\xb5\x01\n QueryDelegatorDelegationsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xd6\x01\n!QueryDelegatorDelegationsResponse\x12h\n\x14\x64\x65legation_responses\x18\x01 \x03(\x0b\x32*.cosmos.staking.v1beta1.DelegationResponseB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x13\x64\x65legationResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xbe\x01\n)QueryDelegatorUnbondingDelegationsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xde\x01\n*QueryDelegatorUnbondingDelegationsResponse\x12g\n\x13unbonding_responses\x18\x01 \x03(\x0b\x32+.cosmos.staking.v1beta1.UnbondingDelegationB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x12unbondingResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xbe\x02\n\x19QueryRedelegationsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\x12src_validator_addr\x18\x02 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10srcValidatorAddr\x12\x46\n\x12\x64st_validator_addr\x18\x03 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x10\x64stValidatorAddr\x12\x46\n\npagination\x18\x04 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xd5\x01\n\x1aQueryRedelegationsResponse\x12n\n\x16redelegation_responses\x18\x01 \x03(\x0b\x32,.cosmos.staking.v1beta1.RedelegationResponseB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x15redelegationResponses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xb4\x01\n\x1fQueryDelegatorValidatorsRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12\x46\n\npagination\x18\x02 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"\xb9\x01\n QueryDelegatorValidatorsResponse\x12L\n\nvalidators\x18\x01 \x03(\x0b\x32!.cosmos.staking.v1beta1.ValidatorB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\nvalidators\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\xb5\x01\n\x1eQueryDelegatorValidatorRequest\x12?\n\x0e\x64\x65legator_addr\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\rdelegatorAddr\x12H\n\x0evalidator_addr\x18\x02 \x01(\tB!\xd2\xb4-\x1d\x63osmos.ValidatorAddressStringR\rvalidatorAddr:\x08\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\"m\n\x1fQueryDelegatorValidatorResponse\x12J\n\tvalidator\x18\x01 \x01(\x0b\x32!.cosmos.staking.v1beta1.ValidatorB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\tvalidator\"4\n\x1aQueryHistoricalInfoRequest\x12\x16\n\x06height\x18\x01 \x01(\x03R\x06height\"\xb4\x01\n\x1bQueryHistoricalInfoResponse\x12>\n\x04hist\x18\x01 \x01(\x0b\x32&.cosmos.staking.v1beta1.HistoricalInfoB\x02\x18\x01R\x04hist\x12U\n\x11historical_record\x18\x02 \x01(\x0b\x32(.cosmos.staking.v1beta1.HistoricalRecordR\x10historicalRecord\"\x12\n\x10QueryPoolRequest\"P\n\x11QueryPoolResponse\x12;\n\x04pool\x18\x01 \x01(\x0b\x32\x1c.cosmos.staking.v1beta1.PoolB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x04pool\"\x14\n\x12QueryParamsRequest\"X\n\x13QueryParamsResponse\x12\x41\n\x06params\x18\x01 \x01(\x0b\x32\x1e.cosmos.staking.v1beta1.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06params2\xb0\x16\n\x05Query\x12\x9e\x01\n\nValidators\x12..cosmos.staking.v1beta1.QueryValidatorsRequest\x1a/.cosmos.staking.v1beta1.QueryValidatorsResponse\"/\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02$\x12\"/cosmos/staking/v1beta1/validators\x12\xac\x01\n\tValidator\x12-.cosmos.staking.v1beta1.QueryValidatorRequest\x1a..cosmos.staking.v1beta1.QueryValidatorResponse\"@\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x35\x12\x33/cosmos/staking/v1beta1/validators/{validator_addr}\x12\xd9\x01\n\x14ValidatorDelegations\x12\x38.cosmos.staking.v1beta1.QueryValidatorDelegationsRequest\x1a\x39.cosmos.staking.v1beta1.QueryValidatorDelegationsResponse\"L\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x41\x12?/cosmos/staking/v1beta1/validators/{validator_addr}/delegations\x12\xfe\x01\n\x1dValidatorUnbondingDelegations\x12\x41.cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest\x1a\x42.cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse\"V\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02K\x12I/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations\x12\xcc\x01\n\nDelegation\x12..cosmos.staking.v1beta1.QueryDelegationRequest\x1a/.cosmos.staking.v1beta1.QueryDelegationResponse\"]\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02R\x12P/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}\x12\xfc\x01\n\x13UnbondingDelegation\x12\x37.cosmos.staking.v1beta1.QueryUnbondingDelegationRequest\x1a\x38.cosmos.staking.v1beta1.QueryUnbondingDelegationResponse\"r\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02g\x12\x65/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation\x12\xce\x01\n\x14\x44\x65legatorDelegations\x12\x38.cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest\x1a\x39.cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse\"A\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x36\x12\x34/cosmos/staking/v1beta1/delegations/{delegator_addr}\x12\xfe\x01\n\x1d\x44\x65legatorUnbondingDelegations\x12\x41.cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest\x1a\x42.cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse\"V\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02K\x12I/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations\x12\xc6\x01\n\rRedelegations\x12\x31.cosmos.staking.v1beta1.QueryRedelegationsRequest\x1a\x32.cosmos.staking.v1beta1.QueryRedelegationsResponse\"N\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x43\x12\x41/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations\x12\xd5\x01\n\x13\x44\x65legatorValidators\x12\x37.cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest\x1a\x38.cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse\"K\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02@\x12>/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators\x12\xe3\x01\n\x12\x44\x65legatorValidator\x12\x36.cosmos.staking.v1beta1.QueryDelegatorValidatorRequest\x1a\x37.cosmos.staking.v1beta1.QueryDelegatorValidatorResponse\"\\\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02Q\x12O/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}\x12\xb8\x01\n\x0eHistoricalInfo\x12\x32.cosmos.staking.v1beta1.QueryHistoricalInfoRequest\x1a\x33.cosmos.staking.v1beta1.QueryHistoricalInfoResponse\"=\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x32\x12\x30/cosmos/staking/v1beta1/historical_info/{height}\x12\x86\x01\n\x04Pool\x12(.cosmos.staking.v1beta1.QueryPoolRequest\x1a).cosmos.staking.v1beta1.QueryPoolResponse\")\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02\x1e\x12\x1c/cosmos/staking/v1beta1/pool\x12\x8e\x01\n\x06Params\x12*.cosmos.staking.v1beta1.QueryParamsRequest\x1a+.cosmos.staking.v1beta1.QueryParamsResponse\"+\x88\xe7\xb0*\x01\x82\xd3\xe4\x93\x02 \x12\x1e/cosmos/staking/v1beta1/paramsB\x1eZ\x1c\x63osmossdk.io/x/staking/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.staking.v1beta1.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\034cosmossdk.io/x/staking/types'
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
  _globals['_QUERYHISTORICALINFORESPONSE'].fields_by_name['hist']._loaded_options = None
  _globals['_QUERYHISTORICALINFORESPONSE'].fields_by_name['hist']._serialized_options = b'\030\001'
  _globals['_QUERYPOOLRESPONSE'].fields_by_name['pool']._loaded_options = None
  _globals['_QUERYPOOLRESPONSE'].fields_by_name['pool']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['params']._loaded_options = None
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['params']._serialized_options = b'\310\336\037\000\250\347\260*\001'
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
  _globals['_QUERYVALIDATORSREQUEST']._serialized_start=271
  _globals['_QUERYVALIDATORSREQUEST']._serialized_end=391
  _globals['_QUERYVALIDATORSRESPONSE']._serialized_start=394
  _globals['_QUERYVALIDATORSRESPONSE']._serialized_end=570
  _globals['_QUERYVALIDATORREQUEST']._serialized_start=572
  _globals['_QUERYVALIDATORREQUEST']._serialized_end=669
  _globals['_QUERYVALIDATORRESPONSE']._serialized_start=671
  _globals['_QUERYVALIDATORRESPONSE']._serialized_end=771
  _globals['_QUERYVALIDATORDELEGATIONSREQUEST']._serialized_start=774
  _globals['_QUERYVALIDATORDELEGATIONSREQUEST']._serialized_end=954
  _globals['_QUERYVALIDATORDELEGATIONSRESPONSE']._serialized_start=957
  _globals['_QUERYVALIDATORDELEGATIONSRESPONSE']._serialized_end=1194
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSREQUEST']._serialized_start=1197
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSREQUEST']._serialized_end=1386
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSRESPONSE']._serialized_start=1389
  _globals['_QUERYVALIDATORUNBONDINGDELEGATIONSRESPONSE']._serialized_end=1611
  _globals['_QUERYDELEGATIONREQUEST']._serialized_start=1614
  _globals['_QUERYDELEGATIONREQUEST']._serialized_end=1787
  _globals['_QUERYDELEGATIONRESPONSE']._serialized_start=1789
  _globals['_QUERYDELEGATIONRESPONSE']._serialized_end=1907
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST']._serialized_start=1910
  _globals['_QUERYUNBONDINGDELEGATIONREQUEST']._serialized_end=2092
  _globals['_QUERYUNBONDINGDELEGATIONRESPONSE']._serialized_start=2094
  _globals['_QUERYUNBONDINGDELEGATIONRESPONSE']._serialized_end=2208
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST']._serialized_start=2211
  _globals['_QUERYDELEGATORDELEGATIONSREQUEST']._serialized_end=2392
  _globals['_QUERYDELEGATORDELEGATIONSRESPONSE']._serialized_start=2395
  _globals['_QUERYDELEGATORDELEGATIONSRESPONSE']._serialized_end=2609
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST']._serialized_start=2612
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSREQUEST']._serialized_end=2802
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSRESPONSE']._serialized_start=2805
  _globals['_QUERYDELEGATORUNBONDINGDELEGATIONSRESPONSE']._serialized_end=3027
  _globals['_QUERYREDELEGATIONSREQUEST']._serialized_start=3030
  _globals['_QUERYREDELEGATIONSREQUEST']._serialized_end=3348
  _globals['_QUERYREDELEGATIONSRESPONSE']._serialized_start=3351
  _globals['_QUERYREDELEGATIONSRESPONSE']._serialized_end=3564
  _globals['_QUERYDELEGATORVALIDATORSREQUEST']._serialized_start=3567
  _globals['_QUERYDELEGATORVALIDATORSREQUEST']._serialized_end=3747
  _globals['_QUERYDELEGATORVALIDATORSRESPONSE']._serialized_start=3750
  _globals['_QUERYDELEGATORVALIDATORSRESPONSE']._serialized_end=3935
  _globals['_QUERYDELEGATORVALIDATORREQUEST']._serialized_start=3938
  _globals['_QUERYDELEGATORVALIDATORREQUEST']._serialized_end=4119
  _globals['_QUERYDELEGATORVALIDATORRESPONSE']._serialized_start=4121
  _globals['_QUERYDELEGATORVALIDATORRESPONSE']._serialized_end=4230
  _globals['_QUERYHISTORICALINFOREQUEST']._serialized_start=4232
  _globals['_QUERYHISTORICALINFOREQUEST']._serialized_end=4284
  _globals['_QUERYHISTORICALINFORESPONSE']._serialized_start=4287
  _globals['_QUERYHISTORICALINFORESPONSE']._serialized_end=4467
  _globals['_QUERYPOOLREQUEST']._serialized_start=4469
  _globals['_QUERYPOOLREQUEST']._serialized_end=4487
  _globals['_QUERYPOOLRESPONSE']._serialized_start=4489
  _globals['_QUERYPOOLRESPONSE']._serialized_end=4569
  _globals['_QUERYPARAMSREQUEST']._serialized_start=4571
  _globals['_QUERYPARAMSREQUEST']._serialized_end=4591
  _globals['_QUERYPARAMSRESPONSE']._serialized_start=4593
  _globals['_QUERYPARAMSRESPONSE']._serialized_end=4681
  _globals['_QUERY']._serialized_start=4684
  _globals['_QUERY']._serialized_end=7548
# @@protoc_insertion_point(module_scope)
