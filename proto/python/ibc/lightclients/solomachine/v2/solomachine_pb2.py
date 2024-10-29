# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ibc/lightclients/solomachine/v2/solomachine.proto
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
    'ibc/lightclients/solomachine/v2/solomachine.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ibc.core.connection.v1 import connection_pb2 as ibc_dot_core_dot_connection_dot_v1_dot_connection__pb2
from ibc.core.channel.v1 import channel_pb2 as ibc_dot_core_dot_channel_dot_v1_dot_channel__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1ibc/lightclients/solomachine/v2/solomachine.proto\x12\x1fibc.lightclients.solomachine.v2\x1a\'ibc/core/connection/v1/connection.proto\x1a!ibc/core/channel/v1/channel.proto\x1a\x14gogoproto/gogo.proto\x1a\x19google/protobuf/any.proto\"\xe5\x01\n\x0b\x43lientState\x12\x1a\n\x08sequence\x18\x01 \x01(\x04R\x08sequence\x12\x1b\n\tis_frozen\x18\x02 \x01(\x08R\x08isFrozen\x12X\n\x0f\x63onsensus_state\x18\x03 \x01(\x0b\x32/.ibc.lightclients.solomachine.v2.ConsensusStateR\x0e\x63onsensusState\x12=\n\x1b\x61llow_update_after_proposal\x18\x04 \x01(\x08R\x18\x61llowUpdateAfterProposal:\x04\x88\xa0\x1f\x00\"\x8b\x01\n\x0e\x43onsensusState\x12\x33\n\npublic_key\x18\x01 \x01(\x0b\x32\x14.google.protobuf.AnyR\tpublicKey\x12 \n\x0b\x64iversifier\x18\x02 \x01(\tR\x0b\x64iversifier\x12\x1c\n\ttimestamp\x18\x03 \x01(\x04R\ttimestamp:\x04\x88\xa0\x1f\x00\"\xcb\x01\n\x06Header\x12\x1a\n\x08sequence\x18\x01 \x01(\x04R\x08sequence\x12\x1c\n\ttimestamp\x18\x02 \x01(\x04R\ttimestamp\x12\x1c\n\tsignature\x18\x03 \x01(\x0cR\tsignature\x12:\n\x0enew_public_key\x18\x04 \x01(\x0b\x32\x14.google.protobuf.AnyR\x0cnewPublicKey\x12\'\n\x0fnew_diversifier\x18\x05 \x01(\tR\x0enewDiversifier:\x04\x88\xa0\x1f\x00\"\xfd\x01\n\x0cMisbehaviour\x12\x1b\n\tclient_id\x18\x01 \x01(\tR\x08\x63lientId\x12\x1a\n\x08sequence\x18\x02 \x01(\x04R\x08sequence\x12V\n\rsignature_one\x18\x03 \x01(\x0b\x32\x31.ibc.lightclients.solomachine.v2.SignatureAndDataR\x0csignatureOne\x12V\n\rsignature_two\x18\x04 \x01(\x0b\x32\x31.ibc.lightclients.solomachine.v2.SignatureAndDataR\x0csignatureTwo:\x04\x88\xa0\x1f\x00\"\xb0\x01\n\x10SignatureAndData\x12\x1c\n\tsignature\x18\x01 \x01(\x0cR\tsignature\x12\x46\n\tdata_type\x18\x02 \x01(\x0e\x32).ibc.lightclients.solomachine.v2.DataTypeR\x08\x64\x61taType\x12\x12\n\x04\x64\x61ta\x18\x03 \x01(\x0cR\x04\x64\x61ta\x12\x1c\n\ttimestamp\x18\x04 \x01(\x04R\ttimestamp:\x04\x88\xa0\x1f\x00\"e\n\x18TimestampedSignatureData\x12%\n\x0esignature_data\x18\x01 \x01(\x0cR\rsignatureData\x12\x1c\n\ttimestamp\x18\x02 \x01(\x04R\ttimestamp:\x04\x88\xa0\x1f\x00\"\xc9\x01\n\tSignBytes\x12\x1a\n\x08sequence\x18\x01 \x01(\x04R\x08sequence\x12\x1c\n\ttimestamp\x18\x02 \x01(\x04R\ttimestamp\x12 \n\x0b\x64iversifier\x18\x03 \x01(\tR\x0b\x64iversifier\x12\x46\n\tdata_type\x18\x04 \x01(\x0e\x32).ibc.lightclients.solomachine.v2.DataTypeR\x08\x64\x61taType\x12\x12\n\x04\x64\x61ta\x18\x05 \x01(\x0cR\x04\x64\x61ta:\x04\x88\xa0\x1f\x00\"q\n\nHeaderData\x12\x34\n\x0bnew_pub_key\x18\x01 \x01(\x0b\x32\x14.google.protobuf.AnyR\tnewPubKey\x12\'\n\x0fnew_diversifier\x18\x02 \x01(\tR\x0enewDiversifier:\x04\x88\xa0\x1f\x00\"d\n\x0f\x43lientStateData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\x12\x37\n\x0c\x63lient_state\x18\x02 \x01(\x0b\x32\x14.google.protobuf.AnyR\x0b\x63lientState:\x04\x88\xa0\x1f\x00\"m\n\x12\x43onsensusStateData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\x12=\n\x0f\x63onsensus_state\x18\x02 \x01(\x0b\x32\x14.google.protobuf.AnyR\x0e\x63onsensusState:\x04\x88\xa0\x1f\x00\"v\n\x13\x43onnectionStateData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\x12\x45\n\nconnection\x18\x02 \x01(\x0b\x32%.ibc.core.connection.v1.ConnectionEndR\nconnection:\x04\x88\xa0\x1f\x00\"d\n\x10\x43hannelStateData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\x12\x36\n\x07\x63hannel\x18\x02 \x01(\x0b\x32\x1c.ibc.core.channel.v1.ChannelR\x07\x63hannel:\x04\x88\xa0\x1f\x00\"J\n\x14PacketCommitmentData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\x12\x1e\n\ncommitment\x18\x02 \x01(\x0cR\ncommitment\"Y\n\x19PacketAcknowledgementData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\x12(\n\x0f\x61\x63knowledgement\x18\x02 \x01(\x0cR\x0f\x61\x63knowledgement\".\n\x18PacketReceiptAbsenceData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\"N\n\x14NextSequenceRecvData\x12\x12\n\x04path\x18\x01 \x01(\x0cR\x04path\x12\"\n\rnext_seq_recv\x18\x02 \x01(\x04R\x0bnextSeqRecv*\x8c\x04\n\x08\x44\x61taType\x12\x38\n#DATA_TYPE_UNINITIALIZED_UNSPECIFIED\x10\x00\x1a\x0f\x8a\x9d \x0bUNSPECIFIED\x12&\n\x16\x44\x41TA_TYPE_CLIENT_STATE\x10\x01\x1a\n\x8a\x9d \x06\x43LIENT\x12,\n\x19\x44\x41TA_TYPE_CONSENSUS_STATE\x10\x02\x1a\r\x8a\x9d \tCONSENSUS\x12.\n\x1a\x44\x41TA_TYPE_CONNECTION_STATE\x10\x03\x1a\x0e\x8a\x9d \nCONNECTION\x12(\n\x17\x44\x41TA_TYPE_CHANNEL_STATE\x10\x04\x1a\x0b\x8a\x9d \x07\x43HANNEL\x12\x35\n\x1b\x44\x41TA_TYPE_PACKET_COMMITMENT\x10\x05\x1a\x14\x8a\x9d \x10PACKETCOMMITMENT\x12?\n DATA_TYPE_PACKET_ACKNOWLEDGEMENT\x10\x06\x1a\x19\x8a\x9d \x15PACKETACKNOWLEDGEMENT\x12>\n DATA_TYPE_PACKET_RECEIPT_ABSENCE\x10\x07\x1a\x18\x8a\x9d \x14PACKETRECEIPTABSENCE\x12\x36\n\x1c\x44\x41TA_TYPE_NEXT_SEQUENCE_RECV\x10\x08\x1a\x14\x8a\x9d \x10NEXTSEQUENCERECV\x12 \n\x10\x44\x41TA_TYPE_HEADER\x10\t\x1a\n\x8a\x9d \x06HEADER\x1a\x04\x88\xa3\x1e\x00\x42\x42Z@github.com/cosmos/ibc-go/v9/modules/core/02-client/migrations/v7b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ibc.lightclients.solomachine.v2.solomachine_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z@github.com/cosmos/ibc-go/v9/modules/core/02-client/migrations/v7'
  _globals['_DATATYPE']._loaded_options = None
  _globals['_DATATYPE']._serialized_options = b'\210\243\036\000'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_UNINITIALIZED_UNSPECIFIED"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_UNINITIALIZED_UNSPECIFIED"]._serialized_options = b'\212\235 \013UNSPECIFIED'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CLIENT_STATE"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CLIENT_STATE"]._serialized_options = b'\212\235 \006CLIENT'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CONSENSUS_STATE"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CONSENSUS_STATE"]._serialized_options = b'\212\235 \tCONSENSUS'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CONNECTION_STATE"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CONNECTION_STATE"]._serialized_options = b'\212\235 \nCONNECTION'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CHANNEL_STATE"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_CHANNEL_STATE"]._serialized_options = b'\212\235 \007CHANNEL'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_PACKET_COMMITMENT"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_PACKET_COMMITMENT"]._serialized_options = b'\212\235 \020PACKETCOMMITMENT'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_PACKET_ACKNOWLEDGEMENT"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_PACKET_ACKNOWLEDGEMENT"]._serialized_options = b'\212\235 \025PACKETACKNOWLEDGEMENT'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_PACKET_RECEIPT_ABSENCE"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_PACKET_RECEIPT_ABSENCE"]._serialized_options = b'\212\235 \024PACKETRECEIPTABSENCE'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_NEXT_SEQUENCE_RECV"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_NEXT_SEQUENCE_RECV"]._serialized_options = b'\212\235 \020NEXTSEQUENCERECV'
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_HEADER"]._loaded_options = None
  _globals['_DATATYPE'].values_by_name["DATA_TYPE_HEADER"]._serialized_options = b'\212\235 \006HEADER'
  _globals['_CLIENTSTATE']._loaded_options = None
  _globals['_CLIENTSTATE']._serialized_options = b'\210\240\037\000'
  _globals['_CONSENSUSSTATE']._loaded_options = None
  _globals['_CONSENSUSSTATE']._serialized_options = b'\210\240\037\000'
  _globals['_HEADER']._loaded_options = None
  _globals['_HEADER']._serialized_options = b'\210\240\037\000'
  _globals['_MISBEHAVIOUR']._loaded_options = None
  _globals['_MISBEHAVIOUR']._serialized_options = b'\210\240\037\000'
  _globals['_SIGNATUREANDDATA']._loaded_options = None
  _globals['_SIGNATUREANDDATA']._serialized_options = b'\210\240\037\000'
  _globals['_TIMESTAMPEDSIGNATUREDATA']._loaded_options = None
  _globals['_TIMESTAMPEDSIGNATUREDATA']._serialized_options = b'\210\240\037\000'
  _globals['_SIGNBYTES']._loaded_options = None
  _globals['_SIGNBYTES']._serialized_options = b'\210\240\037\000'
  _globals['_HEADERDATA']._loaded_options = None
  _globals['_HEADERDATA']._serialized_options = b'\210\240\037\000'
  _globals['_CLIENTSTATEDATA']._loaded_options = None
  _globals['_CLIENTSTATEDATA']._serialized_options = b'\210\240\037\000'
  _globals['_CONSENSUSSTATEDATA']._loaded_options = None
  _globals['_CONSENSUSSTATEDATA']._serialized_options = b'\210\240\037\000'
  _globals['_CONNECTIONSTATEDATA']._loaded_options = None
  _globals['_CONNECTIONSTATEDATA']._serialized_options = b'\210\240\037\000'
  _globals['_CHANNELSTATEDATA']._loaded_options = None
  _globals['_CHANNELSTATEDATA']._serialized_options = b'\210\240\037\000'
  _globals['_DATATYPE']._serialized_start=2379
  _globals['_DATATYPE']._serialized_end=2903
  _globals['_CLIENTSTATE']._serialized_start=212
  _globals['_CLIENTSTATE']._serialized_end=441
  _globals['_CONSENSUSSTATE']._serialized_start=444
  _globals['_CONSENSUSSTATE']._serialized_end=583
  _globals['_HEADER']._serialized_start=586
  _globals['_HEADER']._serialized_end=789
  _globals['_MISBEHAVIOUR']._serialized_start=792
  _globals['_MISBEHAVIOUR']._serialized_end=1045
  _globals['_SIGNATUREANDDATA']._serialized_start=1048
  _globals['_SIGNATUREANDDATA']._serialized_end=1224
  _globals['_TIMESTAMPEDSIGNATUREDATA']._serialized_start=1226
  _globals['_TIMESTAMPEDSIGNATUREDATA']._serialized_end=1327
  _globals['_SIGNBYTES']._serialized_start=1330
  _globals['_SIGNBYTES']._serialized_end=1531
  _globals['_HEADERDATA']._serialized_start=1533
  _globals['_HEADERDATA']._serialized_end=1646
  _globals['_CLIENTSTATEDATA']._serialized_start=1648
  _globals['_CLIENTSTATEDATA']._serialized_end=1748
  _globals['_CONSENSUSSTATEDATA']._serialized_start=1750
  _globals['_CONSENSUSSTATEDATA']._serialized_end=1859
  _globals['_CONNECTIONSTATEDATA']._serialized_start=1861
  _globals['_CONNECTIONSTATEDATA']._serialized_end=1979
  _globals['_CHANNELSTATEDATA']._serialized_start=1981
  _globals['_CHANNELSTATEDATA']._serialized_end=2081
  _globals['_PACKETCOMMITMENTDATA']._serialized_start=2083
  _globals['_PACKETCOMMITMENTDATA']._serialized_end=2157
  _globals['_PACKETACKNOWLEDGEMENTDATA']._serialized_start=2159
  _globals['_PACKETACKNOWLEDGEMENTDATA']._serialized_end=2248
  _globals['_PACKETRECEIPTABSENCEDATA']._serialized_start=2250
  _globals['_PACKETRECEIPTABSENCEDATA']._serialized_end=2296
  _globals['_NEXTSEQUENCERECVDATA']._serialized_start=2298
  _globals['_NEXTSEQUENCERECVDATA']._serialized_end=2376
# @@protoc_insertion_point(module_scope)
