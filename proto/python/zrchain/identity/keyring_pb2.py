# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/identity/keyring.proto
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
    'zrchain/identity/keyring.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ezrchain/identity/keyring.proto\x12\x10zrchain.identity\"\xbc\x02\n\x07Keyring\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\x12\x18\n\x07\x63reator\x18\x02 \x01(\tR\x07\x63reator\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription\x12\x16\n\x06\x61\x64mins\x18\x04 \x03(\tR\x06\x61\x64mins\x12\x18\n\x07parties\x18\x05 \x03(\tR\x07parties\x12\'\n\x0fparty_threshold\x18\x06 \x01(\rR\x0epartyThreshold\x12\x1e\n\x0bkey_req_fee\x18\x07 \x01(\x04R\tkeyReqFee\x12\x1e\n\x0bsig_req_fee\x18\x08 \x01(\x04R\tsigReqFee\x12\x1b\n\tis_active\x18\t \x01(\x08R\x08isActive\x12#\n\rdelegate_fees\x18\n \x01(\x08R\x0c\x64\x65legateFeesB;Z9github.com/Zenrock-Foundation/zrchain/v5/x/identity/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.identity.keyring_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z9github.com/Zenrock-Foundation/zrchain/v5/x/identity/types'
  _globals['_KEYRING']._serialized_start=53
  _globals['_KEYRING']._serialized_end=369
# @@protoc_insertion_point(module_scope)
