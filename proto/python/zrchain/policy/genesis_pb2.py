# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/policy/genesis.proto
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
    'zrchain/policy/genesis.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from zrchain.policy import params_pb2 as zrchain_dot_policy_dot_params__pb2
from zrchain.policy import action_pb2 as zrchain_dot_policy_dot_action__pb2
from zrchain.policy import policy_pb2 as zrchain_dot_policy_dot_policy__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1czrchain/policy/genesis.proto\x12\x0ezrchain.policy\x1a\x14gogoproto/gogo.proto\x1a\x1bzrchain/policy/params.proto\x1a\x1bzrchain/policy/action.proto\x1a\x1bzrchain/policy/policy.proto\x1a\x19google/protobuf/any.proto\"\x9b\x02\n\x0cGenesisState\x12\x34\n\x06params\x18\x01 \x01(\x0b\x32\x16.zrchain.policy.ParamsB\x04\xc8\xde\x1f\x00R\x06params\x12\x17\n\x07port_id\x18\x02 \x01(\tR\x06portId\x12\x38\n\x08policies\x18\x03 \x03(\x0b\x32\x16.zrchain.policy.PolicyB\x04\xc8\xde\x1f\x00R\x08policies\x12\x36\n\x07\x61\x63tions\x18\x04 \x03(\x0b\x32\x16.zrchain.policy.ActionB\x04\xc8\xde\x1f\x00R\x07\x61\x63tions\x12J\n\x0csign_methods\x18\x05 \x03(\x0b\x32!.zrchain.policy.GenesisSignMethodB\x04\xc8\xde\x1f\x00R\x0bsignMethods\"g\n\x11GenesisSignMethod\x12\x14\n\x05owner\x18\x01 \x01(\tR\x05owner\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\x12,\n\x06\x63onfig\x18\x03 \x01(\x0b\x32\x14.google.protobuf.AnyR\x06\x63onfigB9Z7github.com/Zenrock-Foundation/zrchain/v4/x/policy/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.policy.genesis_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z7github.com/Zenrock-Foundation/zrchain/v4/x/policy/types'
  _globals['_GENESISSTATE'].fields_by_name['params']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['params']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE'].fields_by_name['policies']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['policies']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE'].fields_by_name['actions']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['actions']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE'].fields_by_name['sign_methods']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['sign_methods']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE']._serialized_start=185
  _globals['_GENESISSTATE']._serialized_end=468
  _globals['_GENESISSIGNMETHOD']._serialized_start=470
  _globals['_GENESISSIGNMETHOD']._serialized_end=573
# @@protoc_insertion_point(module_scope)
