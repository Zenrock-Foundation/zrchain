# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ibc/core/channel/v2/genesis.proto
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
    'ibc/core/channel/v2/genesis.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!ibc/core/channel/v2/genesis.proto\x12\x13ibc.core.channel.v2\x1a\x14gogoproto/gogo.proto\"\x8f\x03\n\x0cGenesisState\x12R\n\x10\x61\x63knowledgements\x18\x02 \x03(\x0b\x32 .ibc.core.channel.v2.PacketStateB\x04\xc8\xde\x1f\x00R\x10\x61\x63knowledgements\x12H\n\x0b\x63ommitments\x18\x03 \x03(\x0b\x32 .ibc.core.channel.v2.PacketStateB\x04\xc8\xde\x1f\x00R\x0b\x63ommitments\x12\x42\n\x08receipts\x18\x04 \x03(\x0b\x32 .ibc.core.channel.v2.PacketStateB\x04\xc8\xde\x1f\x00R\x08receipts\x12K\n\rasync_packets\x18\x05 \x03(\x0b\x32 .ibc.core.channel.v2.PacketStateB\x04\xc8\xde\x1f\x00R\x0c\x61syncPackets\x12P\n\x0esend_sequences\x18\x06 \x03(\x0b\x32#.ibc.core.channel.v2.PacketSequenceB\x04\xc8\xde\x1f\x00R\rsendSequences\"`\n\x0bPacketState\x12\x1b\n\tclient_id\x18\x01 \x01(\tR\x08\x63lientId\x12\x1a\n\x08sequence\x18\x02 \x01(\x04R\x08sequence\x12\x12\n\x04\x64\x61ta\x18\x03 \x01(\x0cR\x04\x64\x61ta:\x04\x88\xa0\x1f\x00\"I\n\x0ePacketSequence\x12\x1b\n\tclient_id\x18\x01 \x01(\tR\x08\x63lientId\x12\x1a\n\x08sequence\x18\x02 \x01(\x04R\x08sequenceB?Z=github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ibc.core.channel.v2.genesis_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types'
  _globals['_GENESISSTATE'].fields_by_name['acknowledgements']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['acknowledgements']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE'].fields_by_name['commitments']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['commitments']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE'].fields_by_name['receipts']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['receipts']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE'].fields_by_name['async_packets']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['async_packets']._serialized_options = b'\310\336\037\000'
  _globals['_GENESISSTATE'].fields_by_name['send_sequences']._loaded_options = None
  _globals['_GENESISSTATE'].fields_by_name['send_sequences']._serialized_options = b'\310\336\037\000'
  _globals['_PACKETSTATE']._loaded_options = None
  _globals['_PACKETSTATE']._serialized_options = b'\210\240\037\000'
  _globals['_GENESISSTATE']._serialized_start=81
  _globals['_GENESISSTATE']._serialized_end=480
  _globals['_PACKETSTATE']._serialized_start=482
  _globals['_PACKETSTATE']._serialized_end=578
  _globals['_PACKETSEQUENCE']._serialized_start=580
  _globals['_PACKETSEQUENCE']._serialized_end=653
# @@protoc_insertion_point(module_scope)
