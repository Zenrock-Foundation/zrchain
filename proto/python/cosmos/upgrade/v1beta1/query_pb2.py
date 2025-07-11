# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/upgrade/v1beta1/query.proto
# Protobuf Python Version: 6.31.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    31,
    1,
    '',
    'cosmos/upgrade/v1beta1/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from cosmos.upgrade.v1beta1 import upgrade_pb2 as cosmos_dot_upgrade_dot_v1beta1_dot_upgrade__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"cosmos/upgrade/v1beta1/query.proto\x12\x16\x63osmos.upgrade.v1beta1\x1a\x1cgoogle/api/annotations.proto\x1a$cosmos/upgrade/v1beta1/upgrade.proto\x1a\x19\x63osmos_proto/cosmos.proto\"\x19\n\x17QueryCurrentPlanRequest\"L\n\x18QueryCurrentPlanResponse\x12\x30\n\x04plan\x18\x01 \x01(\x0b\x32\x1c.cosmos.upgrade.v1beta1.PlanR\x04plan\"-\n\x17QueryAppliedPlanRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"2\n\x18QueryAppliedPlanResponse\x12\x16\n\x06height\x18\x01 \x01(\x03R\x06height\"I\n\"QueryUpgradedConsensusStateRequest\x12\x1f\n\x0blast_height\x18\x01 \x01(\x03R\nlastHeight:\x02\x18\x01\"~\n#QueryUpgradedConsensusStateResponse\x12M\n\x18upgraded_consensus_state\x18\x02 \x01(\x0c\x42\x13\xda\xb4-\x0f\x63osmos-sdk 0.43R\x16upgradedConsensusState:\x02\x18\x01J\x04\x08\x01\x10\x02\"R\n\x1aQueryModuleVersionsRequest\x12\x1f\n\x0bmodule_name\x18\x01 \x01(\tR\nmoduleName:\x13\xd2\xb4-\x0f\x63osmos-sdk 0.43\"\x82\x01\n\x1bQueryModuleVersionsResponse\x12N\n\x0fmodule_versions\x18\x01 \x03(\x0b\x32%.cosmos.upgrade.v1beta1.ModuleVersionR\x0emoduleVersions:\x13\xd2\xb4-\x0f\x63osmos-sdk 0.43\",\n\x15QueryAuthorityRequest:\x13\xd2\xb4-\x0f\x63osmos-sdk 0.46\"G\n\x16QueryAuthorityResponse\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress:\x13\xd2\xb4-\x0f\x63osmos-sdk 0.462\x9a\x07\n\x05Query\x12\x9e\x01\n\x0b\x43urrentPlan\x12/.cosmos.upgrade.v1beta1.QueryCurrentPlanRequest\x1a\x30.cosmos.upgrade.v1beta1.QueryCurrentPlanResponse\",\x82\xd3\xe4\x93\x02&\x12$/cosmos/upgrade/v1beta1/current_plan\x12\xa5\x01\n\x0b\x41ppliedPlan\x12/.cosmos.upgrade.v1beta1.QueryAppliedPlanRequest\x1a\x30.cosmos.upgrade.v1beta1.QueryAppliedPlanResponse\"3\x82\xd3\xe4\x93\x02-\x12+/cosmos/upgrade/v1beta1/applied_plan/{name}\x12\xdc\x01\n\x16UpgradedConsensusState\x12:.cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest\x1a;.cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse\"I\x88\x02\x01\x82\xd3\xe4\x93\x02@\x12>/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}\x12\xbd\x01\n\x0eModuleVersions\x12\x32.cosmos.upgrade.v1beta1.QueryModuleVersionsRequest\x1a\x33.cosmos.upgrade.v1beta1.QueryModuleVersionsResponse\"B\xca\xb4-\x0f\x63osmos-sdk 0.43\x82\xd3\xe4\x93\x02)\x12\'/cosmos/upgrade/v1beta1/module_versions\x12\xa8\x01\n\tAuthority\x12-.cosmos.upgrade.v1beta1.QueryAuthorityRequest\x1a..cosmos.upgrade.v1beta1.QueryAuthorityResponse\"<\xca\xb4-\x0f\x63osmos-sdk 0.46\x82\xd3\xe4\x93\x02#\x12!/cosmos/upgrade/v1beta1/authorityB.Z,github.com/cosmos/cosmos-sdk/x/upgrade/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.upgrade.v1beta1.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,github.com/cosmos/cosmos-sdk/x/upgrade/types'
  _globals['_QUERYUPGRADEDCONSENSUSSTATEREQUEST']._loaded_options = None
  _globals['_QUERYUPGRADEDCONSENSUSSTATEREQUEST']._serialized_options = b'\030\001'
  _globals['_QUERYUPGRADEDCONSENSUSSTATERESPONSE'].fields_by_name['upgraded_consensus_state']._loaded_options = None
  _globals['_QUERYUPGRADEDCONSENSUSSTATERESPONSE'].fields_by_name['upgraded_consensus_state']._serialized_options = b'\332\264-\017cosmos-sdk 0.43'
  _globals['_QUERYUPGRADEDCONSENSUSSTATERESPONSE']._loaded_options = None
  _globals['_QUERYUPGRADEDCONSENSUSSTATERESPONSE']._serialized_options = b'\030\001'
  _globals['_QUERYMODULEVERSIONSREQUEST']._loaded_options = None
  _globals['_QUERYMODULEVERSIONSREQUEST']._serialized_options = b'\322\264-\017cosmos-sdk 0.43'
  _globals['_QUERYMODULEVERSIONSRESPONSE']._loaded_options = None
  _globals['_QUERYMODULEVERSIONSRESPONSE']._serialized_options = b'\322\264-\017cosmos-sdk 0.43'
  _globals['_QUERYAUTHORITYREQUEST']._loaded_options = None
  _globals['_QUERYAUTHORITYREQUEST']._serialized_options = b'\322\264-\017cosmos-sdk 0.46'
  _globals['_QUERYAUTHORITYRESPONSE']._loaded_options = None
  _globals['_QUERYAUTHORITYRESPONSE']._serialized_options = b'\322\264-\017cosmos-sdk 0.46'
  _globals['_QUERY'].methods_by_name['CurrentPlan']._loaded_options = None
  _globals['_QUERY'].methods_by_name['CurrentPlan']._serialized_options = b'\202\323\344\223\002&\022$/cosmos/upgrade/v1beta1/current_plan'
  _globals['_QUERY'].methods_by_name['AppliedPlan']._loaded_options = None
  _globals['_QUERY'].methods_by_name['AppliedPlan']._serialized_options = b'\202\323\344\223\002-\022+/cosmos/upgrade/v1beta1/applied_plan/{name}'
  _globals['_QUERY'].methods_by_name['UpgradedConsensusState']._loaded_options = None
  _globals['_QUERY'].methods_by_name['UpgradedConsensusState']._serialized_options = b'\210\002\001\202\323\344\223\002@\022>/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}'
  _globals['_QUERY'].methods_by_name['ModuleVersions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['ModuleVersions']._serialized_options = b'\312\264-\017cosmos-sdk 0.43\202\323\344\223\002)\022\'/cosmos/upgrade/v1beta1/module_versions'
  _globals['_QUERY'].methods_by_name['Authority']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Authority']._serialized_options = b'\312\264-\017cosmos-sdk 0.46\202\323\344\223\002#\022!/cosmos/upgrade/v1beta1/authority'
  _globals['_QUERYCURRENTPLANREQUEST']._serialized_start=157
  _globals['_QUERYCURRENTPLANREQUEST']._serialized_end=182
  _globals['_QUERYCURRENTPLANRESPONSE']._serialized_start=184
  _globals['_QUERYCURRENTPLANRESPONSE']._serialized_end=260
  _globals['_QUERYAPPLIEDPLANREQUEST']._serialized_start=262
  _globals['_QUERYAPPLIEDPLANREQUEST']._serialized_end=307
  _globals['_QUERYAPPLIEDPLANRESPONSE']._serialized_start=309
  _globals['_QUERYAPPLIEDPLANRESPONSE']._serialized_end=359
  _globals['_QUERYUPGRADEDCONSENSUSSTATEREQUEST']._serialized_start=361
  _globals['_QUERYUPGRADEDCONSENSUSSTATEREQUEST']._serialized_end=434
  _globals['_QUERYUPGRADEDCONSENSUSSTATERESPONSE']._serialized_start=436
  _globals['_QUERYUPGRADEDCONSENSUSSTATERESPONSE']._serialized_end=562
  _globals['_QUERYMODULEVERSIONSREQUEST']._serialized_start=564
  _globals['_QUERYMODULEVERSIONSREQUEST']._serialized_end=646
  _globals['_QUERYMODULEVERSIONSRESPONSE']._serialized_start=649
  _globals['_QUERYMODULEVERSIONSRESPONSE']._serialized_end=779
  _globals['_QUERYAUTHORITYREQUEST']._serialized_start=781
  _globals['_QUERYAUTHORITYREQUEST']._serialized_end=825
  _globals['_QUERYAUTHORITYRESPONSE']._serialized_start=827
  _globals['_QUERYAUTHORITYRESPONSE']._serialized_end=898
  _globals['_QUERY']._serialized_start=901
  _globals['_QUERY']._serialized_end=1823
# @@protoc_insertion_point(module_scope)
