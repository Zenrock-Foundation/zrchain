# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: google/api/resource.proto
# Protobuf Python Version: 5.29.1
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
    1,
    '',
    'google/api/resource.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import descriptor_pb2 as google_dot_protobuf_dot_descriptor__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19google/api/resource.proto\x12\ngoogle.api\x1a google/protobuf/descriptor.proto\"\xaa\x03\n\x12ResourceDescriptor\x12\x12\n\x04type\x18\x01 \x01(\tR\x04type\x12\x18\n\x07pattern\x18\x02 \x03(\tR\x07pattern\x12\x1d\n\nname_field\x18\x03 \x01(\tR\tnameField\x12@\n\x07history\x18\x04 \x01(\x0e\x32&.google.api.ResourceDescriptor.HistoryR\x07history\x12\x16\n\x06plural\x18\x05 \x01(\tR\x06plural\x12\x1a\n\x08singular\x18\x06 \x01(\tR\x08singular\x12:\n\x05style\x18\n \x03(\x0e\x32$.google.api.ResourceDescriptor.StyleR\x05style\"[\n\x07History\x12\x17\n\x13HISTORY_UNSPECIFIED\x10\x00\x12\x1d\n\x19ORIGINALLY_SINGLE_PATTERN\x10\x01\x12\x18\n\x14\x46UTURE_MULTI_PATTERN\x10\x02\"8\n\x05Style\x12\x15\n\x11STYLE_UNSPECIFIED\x10\x00\x12\x18\n\x14\x44\x45\x43LARATIVE_FRIENDLY\x10\x01\"F\n\x11ResourceReference\x12\x12\n\x04type\x18\x01 \x01(\tR\x04type\x12\x1d\n\nchild_type\x18\x02 \x01(\tR\tchildType:l\n\x12resource_reference\x12\x1d.google.protobuf.FieldOptions\x18\x9f\x08 \x01(\x0b\x32\x1d.google.api.ResourceReferenceR\x11resourceReference:n\n\x13resource_definition\x12\x1c.google.protobuf.FileOptions\x18\x9d\x08 \x03(\x0b\x32\x1e.google.api.ResourceDescriptorR\x12resourceDefinition:\\\n\x08resource\x12\x1f.google.protobuf.MessageOptions\x18\x9d\x08 \x01(\x0b\x32\x1e.google.api.ResourceDescriptorR\x08resourceBn\n\x0e\x63om.google.apiB\rResourceProtoP\x01ZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations\xf8\x01\x01\xa2\x02\x04GAPIb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'google.api.resource_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\016com.google.apiB\rResourceProtoP\001ZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations\370\001\001\242\002\004GAPI'
  _globals['_RESOURCEDESCRIPTOR']._serialized_start=76
  _globals['_RESOURCEDESCRIPTOR']._serialized_end=502
  _globals['_RESOURCEDESCRIPTOR_HISTORY']._serialized_start=353
  _globals['_RESOURCEDESCRIPTOR_HISTORY']._serialized_end=444
  _globals['_RESOURCEDESCRIPTOR_STYLE']._serialized_start=446
  _globals['_RESOURCEDESCRIPTOR_STYLE']._serialized_end=502
  _globals['_RESOURCEREFERENCE']._serialized_start=504
  _globals['_RESOURCEREFERENCE']._serialized_end=574
# @@protoc_insertion_point(module_scope)
