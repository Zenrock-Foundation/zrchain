# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/authz/v1beta1/authz.proto
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
    'cosmos/authz/v1beta1/authz.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from cosmos_proto import cosmos_pb2 as cosmos__proto_dot_cosmos__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n cosmos/authz/v1beta1/authz.proto\x12\x14\x63osmos.authz.v1beta1\x1a\x11\x61mino/amino.proto\x1a\x19\x63osmos_proto/cosmos.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x14gogoproto/gogo.proto\x1a\x19google/protobuf/any.proto\"t\n\x14GenericAuthorization\x12\x10\n\x03msg\x18\x01 \x01(\tR\x03msg:J\xca\xb4-\"cosmos.authz.v1beta1.Authorization\x8a\xe7\xb0*\x1f\x63osmos-sdk/GenericAuthorization\"\xb1\x01\n\x05Grant\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32\x14.google.protobuf.AnyB&\xca\xb4-\"cosmos.authz.v1beta1.AuthorizationR\rauthorization\x12\x44\n\nexpiration\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x01\x90\xdf\x1f\x01R\nexpiration\"\xa2\x02\n\x12GrantAuthorization\x12\x32\n\x07granter\x18\x01 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x07granter\x12\x32\n\x07grantee\x18\x02 \x01(\tB\x18\xd2\xb4-\x14\x63osmos.AddressStringR\x07grantee\x12\x62\n\rauthorization\x18\x03 \x01(\x0b\x32\x14.google.protobuf.AnyB&\xca\xb4-\"cosmos.authz.v1beta1.AuthorizationR\rauthorization\x12@\n\nexpiration\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x04\x90\xdf\x1f\x01R\nexpiration\"4\n\x0eGrantQueueItem\x12\"\n\rmsg_type_urls\x18\x01 \x03(\tR\x0bmsgTypeUrlsB\x1aZ\x14\x63osmossdk.io/x/authz\xc8\xe1\x1e\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.authz.v1beta1.authz_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\024cosmossdk.io/x/authz\310\341\036\000'
  _globals['_GENERICAUTHORIZATION']._loaded_options = None
  _globals['_GENERICAUTHORIZATION']._serialized_options = b'\312\264-\"cosmos.authz.v1beta1.Authorization\212\347\260*\037cosmos-sdk/GenericAuthorization'
  _globals['_GRANT'].fields_by_name['authorization']._loaded_options = None
  _globals['_GRANT'].fields_by_name['authorization']._serialized_options = b'\312\264-\"cosmos.authz.v1beta1.Authorization'
  _globals['_GRANT'].fields_by_name['expiration']._loaded_options = None
  _globals['_GRANT'].fields_by_name['expiration']._serialized_options = b'\310\336\037\001\220\337\037\001'
  _globals['_GRANTAUTHORIZATION'].fields_by_name['granter']._loaded_options = None
  _globals['_GRANTAUTHORIZATION'].fields_by_name['granter']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_GRANTAUTHORIZATION'].fields_by_name['grantee']._loaded_options = None
  _globals['_GRANTAUTHORIZATION'].fields_by_name['grantee']._serialized_options = b'\322\264-\024cosmos.AddressString'
  _globals['_GRANTAUTHORIZATION'].fields_by_name['authorization']._loaded_options = None
  _globals['_GRANTAUTHORIZATION'].fields_by_name['authorization']._serialized_options = b'\312\264-\"cosmos.authz.v1beta1.Authorization'
  _globals['_GRANTAUTHORIZATION'].fields_by_name['expiration']._loaded_options = None
  _globals['_GRANTAUTHORIZATION'].fields_by_name['expiration']._serialized_options = b'\220\337\037\001'
  _globals['_GENERICAUTHORIZATION']._serialized_start=186
  _globals['_GENERICAUTHORIZATION']._serialized_end=302
  _globals['_GRANT']._serialized_start=305
  _globals['_GRANT']._serialized_end=482
  _globals['_GRANTAUTHORIZATION']._serialized_start=485
  _globals['_GRANTAUTHORIZATION']._serialized_end=775
  _globals['_GRANTQUEUEITEM']._serialized_start=777
  _globals['_GRANTQUEUEITEM']._serialized_end=829
# @@protoc_insertion_point(module_scope)