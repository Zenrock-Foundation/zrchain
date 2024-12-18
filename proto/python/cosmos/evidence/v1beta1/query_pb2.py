# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/evidence/v1beta1/query.proto
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
    'cosmos/evidence/v1beta1/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from cosmos.base.query.v1beta1 import pagination_pb2 as cosmos_dot_base_dot_query_dot_v1beta1_dot_pagination__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#cosmos/evidence/v1beta1/query.proto\x12\x17\x63osmos.evidence.v1beta1\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a\x19google/protobuf/any.proto\x1a\x1cgoogle/api/annotations.proto\"S\n\x14QueryEvidenceRequest\x12\'\n\revidence_hash\x18\x01 \x01(\x0c\x42\x02\x18\x01R\x0c\x65videnceHash\x12\x12\n\x04hash\x18\x02 \x01(\tR\x04hash\"I\n\x15QueryEvidenceResponse\x12\x30\n\x08\x65vidence\x18\x01 \x01(\x0b\x32\x14.google.protobuf.AnyR\x08\x65vidence\"a\n\x17QueryAllEvidenceRequest\x12\x46\n\npagination\x18\x01 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\x95\x01\n\x18QueryAllEvidenceResponse\x12\x30\n\x08\x65vidence\x18\x01 \x03(\x0b\x32\x14.google.protobuf.AnyR\x08\x65vidence\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination2\xc5\x02\n\x05Query\x12\x9b\x01\n\x08\x45vidence\x12-.cosmos.evidence.v1beta1.QueryEvidenceRequest\x1a..cosmos.evidence.v1beta1.QueryEvidenceResponse\"0\x82\xd3\xe4\x93\x02*\x12(/cosmos/evidence/v1beta1/evidence/{hash}\x12\x9d\x01\n\x0b\x41llEvidence\x12\x30.cosmos.evidence.v1beta1.QueryAllEvidenceRequest\x1a\x31.cosmos.evidence.v1beta1.QueryAllEvidenceResponse\")\x82\xd3\xe4\x93\x02#\x12!/cosmos/evidence/v1beta1/evidenceB\x1fZ\x1d\x63osmossdk.io/x/evidence/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.evidence.v1beta1.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\035cosmossdk.io/x/evidence/types'
  _globals['_QUERYEVIDENCEREQUEST'].fields_by_name['evidence_hash']._loaded_options = None
  _globals['_QUERYEVIDENCEREQUEST'].fields_by_name['evidence_hash']._serialized_options = b'\030\001'
  _globals['_QUERY'].methods_by_name['Evidence']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Evidence']._serialized_options = b'\202\323\344\223\002*\022(/cosmos/evidence/v1beta1/evidence/{hash}'
  _globals['_QUERY'].methods_by_name['AllEvidence']._loaded_options = None
  _globals['_QUERY'].methods_by_name['AllEvidence']._serialized_options = b'\202\323\344\223\002#\022!/cosmos/evidence/v1beta1/evidence'
  _globals['_QUERYEVIDENCEREQUEST']._serialized_start=165
  _globals['_QUERYEVIDENCEREQUEST']._serialized_end=248
  _globals['_QUERYEVIDENCERESPONSE']._serialized_start=250
  _globals['_QUERYEVIDENCERESPONSE']._serialized_end=323
  _globals['_QUERYALLEVIDENCEREQUEST']._serialized_start=325
  _globals['_QUERYALLEVIDENCEREQUEST']._serialized_end=422
  _globals['_QUERYALLEVIDENCERESPONSE']._serialized_start=425
  _globals['_QUERYALLEVIDENCERESPONSE']._serialized_end=574
  _globals['_QUERY']._serialized_start=577
  _globals['_QUERY']._serialized_end=902
# @@protoc_insertion_point(module_scope)
