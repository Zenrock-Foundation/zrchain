# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/evidence/v1beta1/tx.proto
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
    'cosmos/evidence/v1beta1/tx.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from cosmos.msg.v1 import msg_pb2 as cosmos_dot_msg_dot_v1_dot_msg__pb2
from amino import amino_pb2 as amino_dot_amino__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n cosmos/evidence/v1beta1/tx.proto\x12\x17\x63osmos.evidence.v1beta1\x1a\x14gogoproto/gogo.proto\x1a\x19google/protobuf/any.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x17\x63osmos/msg/v1/msg.proto\x1a\x11\x61mino/amino.proto\"\xdc\x01\n\x11MsgSubmitEvidence\x12\x36\n\tsubmitter\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\tsubmitter\x12V\n\x08\x65vidence\x18\x02 \x01(\x0b\x32\x14.google.protobuf.AnyB$\xca\xb4- cosmos.evidence.v1beta1.EvidenceR\x08\x65vidence:7\x88\xa0\x1f\x00\xe8\xa0\x1f\x00\x82\xe7\xb0*\tsubmitter\x8a\xe7\xb0*\x1c\x63osmos-sdk/MsgSubmitEvidence\"/\n\x19MsgSubmitEvidenceResponse\x12\x12\n\x04hash\x18\x04 \x01(\x0cR\x04hash2~\n\x03Msg\x12p\n\x0eSubmitEvidence\x12*.cosmos.evidence.v1beta1.MsgSubmitEvidence\x1a\x32.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse\x1a\x05\x80\xe7\xb0*\x01\x42#Z\x1d\x63osmossdk.io/x/evidence/types\xa8\xe2\x1e\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.evidence.v1beta1.tx_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\035cosmossdk.io/x/evidence/types\250\342\036\001'
  _globals['_MSGSUBMITEVIDENCE'].fields_by_name['submitter']._loaded_options = None
  _globals['_MSGSUBMITEVIDENCE'].fields_by_name['submitter']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_MSGSUBMITEVIDENCE'].fields_by_name['evidence']._loaded_options = None
  _globals['_MSGSUBMITEVIDENCE'].fields_by_name['evidence']._serialized_options = b'\312\264- cosmos.evidence.v1beta1.Evidence'
  _globals['_MSGSUBMITEVIDENCE']._loaded_options = None
  _globals['_MSGSUBMITEVIDENCE']._serialized_options = b'\210\240\037\000\350\240\037\000\202\347\260*\tsubmitter\212\347\260*\034cosmos-sdk/MsgSubmitEvidence'
  _globals['_MSG']._loaded_options = None
  _globals['_MSG']._serialized_options = b'\200\347\260*\001'
  _globals['_MSGSUBMITEVIDENCE']._serialized_start=182
  _globals['_MSGSUBMITEVIDENCE']._serialized_end=402
  _globals['_MSGSUBMITEVIDENCERESPONSE']._serialized_start=404
  _globals['_MSGSUBMITEVIDENCERESPONSE']._serialized_end=451
  _globals['_MSG']._serialized_start=453
  _globals['_MSG']._serialized_end=579
# @@protoc_insertion_point(module_scope)
