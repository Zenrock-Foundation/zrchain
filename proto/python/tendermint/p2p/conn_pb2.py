# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: tendermint/p2p/conn.proto
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
    'tendermint/p2p/conn.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from tendermint.crypto import keys_pb2 as tendermint_dot_crypto_dot_keys__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19tendermint/p2p/conn.proto\x12\x0etendermint.p2p\x1a\x14gogoproto/gogo.proto\x1a\x1ctendermint/crypto/keys.proto\"\x0c\n\nPacketPing\"\x0c\n\nPacketPong\"h\n\tPacketMsg\x12,\n\nchannel_id\x18\x01 \x01(\x05\x42\r\xe2\xde\x1f\tChannelIDR\tchannelId\x12\x19\n\x03\x65of\x18\x02 \x01(\x08\x42\x07\xe2\xde\x1f\x03\x45OFR\x03\x65of\x12\x12\n\x04\x64\x61ta\x18\x03 \x01(\x0cR\x04\x64\x61ta\"\xc9\x01\n\x06Packet\x12=\n\x0bpacket_ping\x18\x01 \x01(\x0b\x32\x1a.tendermint.p2p.PacketPingH\x00R\npacketPing\x12=\n\x0bpacket_pong\x18\x02 \x01(\x0b\x32\x1a.tendermint.p2p.PacketPongH\x00R\npacketPong\x12:\n\npacket_msg\x18\x03 \x01(\x0b\x32\x19.tendermint.p2p.PacketMsgH\x00R\tpacketMsgB\x05\n\x03sum\"_\n\x0e\x41uthSigMessage\x12;\n\x07pub_key\x18\x01 \x01(\x0b\x32\x1c.tendermint.crypto.PublicKeyB\x04\xc8\xde\x1f\x00R\x06pubKey\x12\x10\n\x03sig\x18\x02 \x01(\x0cR\x03sigB3Z1github.com/cometbft/cometbft/proto/tendermint/p2pb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'tendermint.p2p.conn_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z1github.com/cometbft/cometbft/proto/tendermint/p2p'
  _globals['_PACKETMSG'].fields_by_name['channel_id']._loaded_options = None
  _globals['_PACKETMSG'].fields_by_name['channel_id']._serialized_options = b'\342\336\037\tChannelID'
  _globals['_PACKETMSG'].fields_by_name['eof']._loaded_options = None
  _globals['_PACKETMSG'].fields_by_name['eof']._serialized_options = b'\342\336\037\003EOF'
  _globals['_AUTHSIGMESSAGE'].fields_by_name['pub_key']._loaded_options = None
  _globals['_AUTHSIGMESSAGE'].fields_by_name['pub_key']._serialized_options = b'\310\336\037\000'
  _globals['_PACKETPING']._serialized_start=97
  _globals['_PACKETPING']._serialized_end=109
  _globals['_PACKETPONG']._serialized_start=111
  _globals['_PACKETPONG']._serialized_end=123
  _globals['_PACKETMSG']._serialized_start=125
  _globals['_PACKETMSG']._serialized_end=229
  _globals['_PACKET']._serialized_start=232
  _globals['_PACKET']._serialized_end=433
  _globals['_AUTHSIGMESSAGE']._serialized_start=435
  _globals['_AUTHSIGMESSAGE']._serialized_end=530
# @@protoc_insertion_point(module_scope)
