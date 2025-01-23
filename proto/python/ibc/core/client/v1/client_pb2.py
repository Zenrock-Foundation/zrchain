# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ibc/core/client/v1/client.proto
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
    'ibc/core/client/v1/client.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fibc/core/client/v1/client.proto\x12\x12ibc.core.client.v1\x1a\x14gogoproto/gogo.proto\x1a\x19google/protobuf/any.proto\"m\n\x15IdentifiedClientState\x12\x1b\n\tclient_id\x18\x01 \x01(\tR\x08\x63lientId\x12\x37\n\x0c\x63lient_state\x18\x02 \x01(\x0b\x32\x14.google.protobuf.AnyR\x0b\x63lientState\"\x93\x01\n\x18\x43onsensusStateWithHeight\x12\x38\n\x06height\x18\x01 \x01(\x0b\x32\x1a.ibc.core.client.v1.HeightB\x04\xc8\xde\x1f\x00R\x06height\x12=\n\x0f\x63onsensus_state\x18\x02 \x01(\x0b\x32\x14.google.protobuf.AnyR\x0e\x63onsensusState\"\x93\x01\n\x15\x43lientConsensusStates\x12\x1b\n\tclient_id\x18\x01 \x01(\tR\x08\x63lientId\x12]\n\x10\x63onsensus_states\x18\x02 \x03(\x0b\x32,.ibc.core.client.v1.ConsensusStateWithHeightB\x04\xc8\xde\x1f\x00R\x0f\x63onsensusStates\"\x8e\x01\n\x06Height\x12<\n\x0frevision_number\x18\x01 \x01(\x04\x42\x13\xea\xde\x1f\x0frevision_numberR\x0erevisionNumber\x12<\n\x0frevision_height\x18\x02 \x01(\x04\x42\x13\xea\xde\x1f\x0frevision_heightR\x0erevisionHeight:\x08\x88\xa0\x1f\x00\x98\xa0\x1f\x00\"1\n\x06Params\x12\'\n\x0f\x61llowed_clients\x18\x01 \x03(\tR\x0e\x61llowedClientsB:Z8github.com/cosmos/ibc-go/v9/modules/core/02-client/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ibc.core.client.v1.client_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z8github.com/cosmos/ibc-go/v9/modules/core/02-client/types'
  _globals['_CONSENSUSSTATEWITHHEIGHT'].fields_by_name['height']._loaded_options = None
  _globals['_CONSENSUSSTATEWITHHEIGHT'].fields_by_name['height']._serialized_options = b'\310\336\037\000'
  _globals['_CLIENTCONSENSUSSTATES'].fields_by_name['consensus_states']._loaded_options = None
  _globals['_CLIENTCONSENSUSSTATES'].fields_by_name['consensus_states']._serialized_options = b'\310\336\037\000'
  _globals['_HEIGHT'].fields_by_name['revision_number']._loaded_options = None
  _globals['_HEIGHT'].fields_by_name['revision_number']._serialized_options = b'\352\336\037\017revision_number'
  _globals['_HEIGHT'].fields_by_name['revision_height']._loaded_options = None
  _globals['_HEIGHT'].fields_by_name['revision_height']._serialized_options = b'\352\336\037\017revision_height'
  _globals['_HEIGHT']._loaded_options = None
  _globals['_HEIGHT']._serialized_options = b'\210\240\037\000\230\240\037\000'
  _globals['_IDENTIFIEDCLIENTSTATE']._serialized_start=104
  _globals['_IDENTIFIEDCLIENTSTATE']._serialized_end=213
  _globals['_CONSENSUSSTATEWITHHEIGHT']._serialized_start=216
  _globals['_CONSENSUSSTATEWITHHEIGHT']._serialized_end=363
  _globals['_CLIENTCONSENSUSSTATES']._serialized_start=366
  _globals['_CLIENTCONSENSUSSTATES']._serialized_end=513
  _globals['_HEIGHT']._serialized_start=516
  _globals['_HEIGHT']._serialized_end=658
  _globals['_PARAMS']._serialized_start=660
  _globals['_PARAMS']._serialized_end=709
# @@protoc_insertion_point(module_scope)
