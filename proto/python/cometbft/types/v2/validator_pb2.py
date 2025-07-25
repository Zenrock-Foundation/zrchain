# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cometbft/types/v2/validator.proto
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
    'cometbft/types/v2/validator.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from cometbft.crypto.v1 import keys_pb2 as cometbft_dot_crypto_dot_v1_dot_keys__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!cometbft/types/v2/validator.proto\x12\x11\x63ometbft.types.v2\x1a\x1d\x63ometbft/crypto/v1/keys.proto\x1a\x14gogoproto/gogo.proto\"\xb4\x01\n\x0cValidatorSet\x12<\n\nvalidators\x18\x01 \x03(\x0b\x32\x1c.cometbft.types.v2.ValidatorR\nvalidators\x12\x38\n\x08proposer\x18\x02 \x01(\x0b\x32\x1c.cometbft.types.v2.ValidatorR\x08proposer\x12,\n\x12total_voting_power\x18\x03 \x01(\x03R\x10totalVotingPower\"\xf7\x01\n\tValidator\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0cR\x07\x61\x64\x64ress\x12:\n\x07pub_key\x18\x02 \x01(\x0b\x32\x1d.cometbft.crypto.v1.PublicKeyB\x02\x18\x01R\x06pubKey\x12!\n\x0cvoting_power\x18\x03 \x01(\x03R\x0bvotingPower\x12+\n\x11proposer_priority\x18\x04 \x01(\x03R\x10proposerPriority\x12\"\n\rpub_key_bytes\x18\x05 \x01(\x0cR\x0bpubKeyBytes\x12 \n\x0cpub_key_type\x18\x06 \x01(\tR\npubKeyType\"l\n\x0fSimpleValidator\x12\x36\n\x07pub_key\x18\x01 \x01(\x0b\x32\x1d.cometbft.crypto.v1.PublicKeyR\x06pubKey\x12!\n\x0cvoting_power\x18\x02 \x01(\x03R\x0bvotingPower*\xd7\x01\n\x0b\x42lockIDFlag\x12\x31\n\x15\x42LOCK_ID_FLAG_UNKNOWN\x10\x00\x1a\x16\x8a\x9d \x12\x42lockIDFlagUnknown\x12/\n\x14\x42LOCK_ID_FLAG_ABSENT\x10\x01\x1a\x15\x8a\x9d \x11\x42lockIDFlagAbsent\x12/\n\x14\x42LOCK_ID_FLAG_COMMIT\x10\x02\x1a\x15\x8a\x9d \x11\x42lockIDFlagCommit\x12)\n\x11\x42LOCK_ID_FLAG_NIL\x10\x03\x1a\x12\x8a\x9d \x0e\x42lockIDFlagNil\x1a\x08\x88\xa3\x1e\x00\xa8\xa4\x1e\x01\x42\x34Z2github.com/cometbft/cometbft/api/cometbft/types/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cometbft.types.v2.validator_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z2github.com/cometbft/cometbft/api/cometbft/types/v2'
  _globals['_BLOCKIDFLAG']._loaded_options = None
  _globals['_BLOCKIDFLAG']._serialized_options = b'\210\243\036\000\250\244\036\001'
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_UNKNOWN"]._loaded_options = None
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_UNKNOWN"]._serialized_options = b'\212\235 \022BlockIDFlagUnknown'
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_ABSENT"]._loaded_options = None
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_ABSENT"]._serialized_options = b'\212\235 \021BlockIDFlagAbsent'
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_COMMIT"]._loaded_options = None
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_COMMIT"]._serialized_options = b'\212\235 \021BlockIDFlagCommit'
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_NIL"]._loaded_options = None
  _globals['_BLOCKIDFLAG'].values_by_name["BLOCK_ID_FLAG_NIL"]._serialized_options = b'\212\235 \016BlockIDFlagNil'
  _globals['_VALIDATOR'].fields_by_name['pub_key']._loaded_options = None
  _globals['_VALIDATOR'].fields_by_name['pub_key']._serialized_options = b'\030\001'
  _globals['_BLOCKIDFLAG']._serialized_start=653
  _globals['_BLOCKIDFLAG']._serialized_end=868
  _globals['_VALIDATORSET']._serialized_start=110
  _globals['_VALIDATORSET']._serialized_end=290
  _globals['_VALIDATOR']._serialized_start=293
  _globals['_VALIDATOR']._serialized_end=540
  _globals['_SIMPLEVALIDATOR']._serialized_start=542
  _globals['_SIMPLEVALIDATOR']._serialized_end=650
# @@protoc_insertion_point(module_scope)
