# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/group/module/v1/module.proto
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
    'cosmos/group/module/v1/module.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from cosmos.app.v1alpha1 import module_pb2 as cosmos_dot_app_dot_v1alpha1_dot_module__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from amino import amino_pb2 as amino_dot_amino__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#cosmos/group/module/v1/module.proto\x12\x16\x63osmos.group.module.v1\x1a cosmos/app/v1alpha1/module.proto\x1a\x14gogoproto/gogo.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x11\x61mino/amino.proto\"\x9a\x02\n\x06Module\x12Z\n\x14max_execution_period\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationB\r\xc8\xde\x1f\x00\x98\xdf\x1f\x01\xa8\xe7\xb0*\x01R\x12maxExecutionPeriod\x12(\n\x10max_metadata_len\x18\x02 \x01(\x04R\x0emaxMetadataLen\x12\x33\n\x16max_proposal_title_len\x18\x03 \x01(\x04R\x13maxProposalTitleLen\x12\x37\n\x18max_proposal_summary_len\x18\x04 \x01(\x04R\x15maxProposalSummaryLen:\x1c\xba\xc0\x96\xda\x01\x16\n\x14\x63osmossdk.io/x/groupb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.group.module.v1.module_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_MODULE'].fields_by_name['max_execution_period']._loaded_options = None
  _globals['_MODULE'].fields_by_name['max_execution_period']._serialized_options = b'\310\336\037\000\230\337\037\001\250\347\260*\001'
  _globals['_MODULE']._loaded_options = None
  _globals['_MODULE']._serialized_options = b'\272\300\226\332\001\026\n\024cosmossdk.io/x/group'
  _globals['_MODULE']._serialized_start=171
  _globals['_MODULE']._serialized_end=453
# @@protoc_insertion_point(module_scope)
