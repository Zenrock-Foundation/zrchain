# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/gov/module/v1/module.proto
# Protobuf Python Version: 5.29.3
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
    3,
    '',
    'cosmos/gov/module/v1/module.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from cosmos.app.v1alpha1 import module_pb2 as cosmos_dot_app_dot_v1alpha1_dot_module__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!cosmos/gov/module/v1/module.proto\x12\x14\x63osmos.gov.module.v1\x1a cosmos/app/v1alpha1/module.proto\"\xb8\x01\n\x06Module\x12(\n\x10max_metadata_len\x18\x01 \x01(\x04R\x0emaxMetadataLen\x12\x1c\n\tauthority\x18\x02 \x01(\tR\tauthority\x12\"\n\rmax_title_len\x18\x03 \x01(\x04R\x0bmaxTitleLen\x12&\n\x0fmax_summary_len\x18\x04 \x01(\x04R\rmaxSummaryLen:\x1a\xba\xc0\x96\xda\x01\x14\n\x12\x63osmossdk.io/x/govb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.gov.module.v1.module_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_MODULE']._loaded_options = None
  _globals['_MODULE']._serialized_options = b'\272\300\226\332\001\024\n\022cosmossdk.io/x/gov'
  _globals['_MODULE']._serialized_start=94
  _globals['_MODULE']._serialized_end=278
# @@protoc_insertion_point(module_scope)
