# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: google/protobuf/java_features.proto
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
    'google/protobuf/java_features.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import descriptor_pb2 as google_dot_protobuf_dot_descriptor__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#google/protobuf/java_features.proto\x12\x02pb\x1a google/protobuf/descriptor.proto\"\x8b\x06\n\x0cJavaFeatures\x12\x90\x02\n\x12legacy_closed_enum\x18\x01 \x01(\x08\x42\xe1\x01\x88\x01\x01\x98\x01\x04\x98\x01\x01\xa2\x01\t\x12\x04true\x18\x84\x07\xa2\x01\n\x12\x05\x66\x61lse\x18\xe7\x07\xb2\x01\xbb\x01\x08\xe8\x07\x10\xe8\x07\x1a\xb2\x01The legacy closed enum behavior in Java is deprecated and is scheduled to be removed in edition 2025.  See http://protobuf.dev/programming-guides/enum/#java for more information.R\x10legacyClosedEnum\x12\xaf\x02\n\x0futf8_validation\x18\x02 \x01(\x0e\x32\x1f.pb.JavaFeatures.Utf8ValidationB\xe4\x01\x88\x01\x01\x98\x01\x04\x98\x01\x01\xa2\x01\x0c\x12\x07\x44\x45\x46\x41ULT\x18\x84\x07\xb2\x01\xc8\x01\x08\xe8\x07\x10\xe9\x07\x1a\xbf\x01The Java-specific utf8 validation feature is deprecated and is scheduled to be removed in edition 2025.  Utf8 validation behavior should use the global cross-language utf8_validation feature.R\x0eutf8Validation\x12n\n\x1fuse_old_outer_classname_default\x18\x04 \x01(\x08\x42(\x88\x01\x01\x98\x01\x01\xa2\x01\t\x12\x04true\x18\x84\x07\xa2\x01\n\x12\x05\x66\x61lse\x18\xe9\x07\xb2\x01\x06\x08\xe9\x07 \xe9\x07R\x1buseOldOuterClassnameDefault\"F\n\x0eUtf8Validation\x12\x1b\n\x17UTF8_VALIDATION_UNKNOWN\x10\x00\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x01\x12\n\n\x06VERIFY\x10\x02:B\n\x04java\x12\x1b.google.protobuf.FeatureSet\x18\xe9\x07 \x01(\x0b\x32\x10.pb.JavaFeaturesR\x04javaB(\n\x13\x63om.google.protobufB\x11JavaFeaturesProto')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'google.protobuf.java_features_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\023com.google.protobufB\021JavaFeaturesProto'
  _globals['_JAVAFEATURES'].fields_by_name['legacy_closed_enum']._loaded_options = None
  _globals['_JAVAFEATURES'].fields_by_name['legacy_closed_enum']._serialized_options = b'\210\001\001\230\001\004\230\001\001\242\001\t\022\004true\030\204\007\242\001\n\022\005false\030\347\007\262\001\273\001\010\350\007\020\350\007\032\262\001The legacy closed enum behavior in Java is deprecated and is scheduled to be removed in edition 2025.  See http://protobuf.dev/programming-guides/enum/#java for more information.'
  _globals['_JAVAFEATURES'].fields_by_name['utf8_validation']._loaded_options = None
  _globals['_JAVAFEATURES'].fields_by_name['utf8_validation']._serialized_options = b'\210\001\001\230\001\004\230\001\001\242\001\014\022\007DEFAULT\030\204\007\262\001\310\001\010\350\007\020\351\007\032\277\001The Java-specific utf8 validation feature is deprecated and is scheduled to be removed in edition 2025.  Utf8 validation behavior should use the global cross-language utf8_validation feature.'
  _globals['_JAVAFEATURES'].fields_by_name['use_old_outer_classname_default']._loaded_options = None
  _globals['_JAVAFEATURES'].fields_by_name['use_old_outer_classname_default']._serialized_options = b'\210\001\001\230\001\001\242\001\t\022\004true\030\204\007\242\001\n\022\005false\030\351\007\262\001\006\010\351\007 \351\007'
  _globals['_JAVAFEATURES']._serialized_start=78
  _globals['_JAVAFEATURES']._serialized_end=857
  _globals['_JAVAFEATURES_UTF8VALIDATION']._serialized_start=787
  _globals['_JAVAFEATURES_UTF8VALIDATION']._serialized_end=857
# @@protoc_insertion_point(module_scope)
