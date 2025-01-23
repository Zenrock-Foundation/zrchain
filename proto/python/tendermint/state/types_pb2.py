# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: tendermint/state/types.proto
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
    'tendermint/state/types.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from tendermint.abci import types_pb2 as tendermint_dot_abci_dot_types__pb2
from tendermint.types import types_pb2 as tendermint_dot_types_dot_types__pb2
from tendermint.types import validator_pb2 as tendermint_dot_types_dot_validator__pb2
from tendermint.types import params_pb2 as tendermint_dot_types_dot_params__pb2
from tendermint.version import types_pb2 as tendermint_dot_version_dot_types__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ctendermint/state/types.proto\x12\x10tendermint.state\x1a\x14gogoproto/gogo.proto\x1a\x1btendermint/abci/types.proto\x1a\x1ctendermint/types/types.proto\x1a tendermint/types/validator.proto\x1a\x1dtendermint/types/params.proto\x1a\x1etendermint/version/types.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xdd\x01\n\x13LegacyABCIResponses\x12>\n\x0b\x64\x65liver_txs\x18\x01 \x03(\x0b\x32\x1d.tendermint.abci.ExecTxResultR\ndeliverTxs\x12?\n\tend_block\x18\x02 \x01(\x0b\x32\".tendermint.state.ResponseEndBlockR\x08\x65ndBlock\x12\x45\n\x0b\x62\x65gin_block\x18\x03 \x01(\x0b\x32$.tendermint.state.ResponseBeginBlockR\nbeginBlock\"^\n\x12ResponseBeginBlock\x12H\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x16.tendermint.abci.EventB\x18\xc8\xde\x1f\x00\xea\xde\x1f\x10\x65vents,omitemptyR\x06\x65vents\"\x8c\x02\n\x10ResponseEndBlock\x12S\n\x11validator_updates\x18\x01 \x03(\x0b\x32 .tendermint.abci.ValidatorUpdateB\x04\xc8\xde\x1f\x00R\x10validatorUpdates\x12Y\n\x17\x63onsensus_param_updates\x18\x02 \x01(\x0b\x32!.tendermint.types.ConsensusParamsR\x15\x63onsensusParamUpdates\x12H\n\x06\x65vents\x18\x03 \x03(\x0b\x32\x16.tendermint.abci.EventB\x18\xc8\xde\x1f\x00\xea\xde\x1f\x10\x65vents,omitemptyR\x06\x65vents\"\x85\x01\n\x0eValidatorsInfo\x12\x43\n\rvalidator_set\x18\x01 \x01(\x0b\x32\x1e.tendermint.types.ValidatorSetR\x0cvalidatorSet\x12.\n\x13last_height_changed\x18\x02 \x01(\x03R\x11lastHeightChanged\"\x99\x01\n\x13\x43onsensusParamsInfo\x12R\n\x10\x63onsensus_params\x18\x01 \x01(\x0b\x32!.tendermint.types.ConsensusParamsB\x04\xc8\xde\x1f\x00R\x0f\x63onsensusParams\x12.\n\x13last_height_changed\x18\x02 \x01(\x03R\x11lastHeightChanged\"\xe6\x01\n\x11\x41\x42\x43IResponsesInfo\x12Y\n\x15legacy_abci_responses\x18\x01 \x01(\x0b\x32%.tendermint.state.LegacyABCIResponsesR\x13legacyAbciResponses\x12\x16\n\x06height\x18\x02 \x01(\x03R\x06height\x12^\n\x17response_finalize_block\x18\x03 \x01(\x0b\x32&.tendermint.abci.ResponseFinalizeBlockR\x15responseFinalizeBlock\"h\n\x07Version\x12\x41\n\tconsensus\x18\x01 \x01(\x0b\x32\x1d.tendermint.version.ConsensusB\x04\xc8\xde\x1f\x00R\tconsensus\x12\x1a\n\x08software\x18\x02 \x01(\tR\x08software\"\xe1\x06\n\x05State\x12\x39\n\x07version\x18\x01 \x01(\x0b\x32\x19.tendermint.state.VersionB\x04\xc8\xde\x1f\x00R\x07version\x12&\n\x08\x63hain_id\x18\x02 \x01(\tB\x0b\xe2\xde\x1f\x07\x43hainIDR\x07\x63hainId\x12%\n\x0einitial_height\x18\x0e \x01(\x03R\rinitialHeight\x12*\n\x11last_block_height\x18\x03 \x01(\x03R\x0flastBlockHeight\x12R\n\rlast_block_id\x18\x04 \x01(\x0b\x32\x19.tendermint.types.BlockIDB\x13\xc8\xde\x1f\x00\xe2\xde\x1f\x0bLastBlockIDR\x0blastBlockId\x12L\n\x0flast_block_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01R\rlastBlockTime\x12G\n\x0fnext_validators\x18\x06 \x01(\x0b\x32\x1e.tendermint.types.ValidatorSetR\x0enextValidators\x12>\n\nvalidators\x18\x07 \x01(\x0b\x32\x1e.tendermint.types.ValidatorSetR\nvalidators\x12G\n\x0flast_validators\x18\x08 \x01(\x0b\x32\x1e.tendermint.types.ValidatorSetR\x0elastValidators\x12\x43\n\x1elast_height_validators_changed\x18\t \x01(\x03R\x1blastHeightValidatorsChanged\x12R\n\x10\x63onsensus_params\x18\n \x01(\x0b\x32!.tendermint.types.ConsensusParamsB\x04\xc8\xde\x1f\x00R\x0f\x63onsensusParams\x12N\n$last_height_consensus_params_changed\x18\x0b \x01(\x03R lastHeightConsensusParamsChanged\x12*\n\x11last_results_hash\x18\x0c \x01(\x0cR\x0flastResultsHash\x12\x19\n\x08\x61pp_hash\x18\r \x01(\x0cR\x07\x61ppHashB5Z3github.com/cometbft/cometbft/proto/tendermint/stateb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'tendermint.state.types_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z3github.com/cometbft/cometbft/proto/tendermint/state'
  _globals['_RESPONSEBEGINBLOCK'].fields_by_name['events']._loaded_options = None
  _globals['_RESPONSEBEGINBLOCK'].fields_by_name['events']._serialized_options = b'\310\336\037\000\352\336\037\020events,omitempty'
  _globals['_RESPONSEENDBLOCK'].fields_by_name['validator_updates']._loaded_options = None
  _globals['_RESPONSEENDBLOCK'].fields_by_name['validator_updates']._serialized_options = b'\310\336\037\000'
  _globals['_RESPONSEENDBLOCK'].fields_by_name['events']._loaded_options = None
  _globals['_RESPONSEENDBLOCK'].fields_by_name['events']._serialized_options = b'\310\336\037\000\352\336\037\020events,omitempty'
  _globals['_CONSENSUSPARAMSINFO'].fields_by_name['consensus_params']._loaded_options = None
  _globals['_CONSENSUSPARAMSINFO'].fields_by_name['consensus_params']._serialized_options = b'\310\336\037\000'
  _globals['_VERSION'].fields_by_name['consensus']._loaded_options = None
  _globals['_VERSION'].fields_by_name['consensus']._serialized_options = b'\310\336\037\000'
  _globals['_STATE'].fields_by_name['version']._loaded_options = None
  _globals['_STATE'].fields_by_name['version']._serialized_options = b'\310\336\037\000'
  _globals['_STATE'].fields_by_name['chain_id']._loaded_options = None
  _globals['_STATE'].fields_by_name['chain_id']._serialized_options = b'\342\336\037\007ChainID'
  _globals['_STATE'].fields_by_name['last_block_id']._loaded_options = None
  _globals['_STATE'].fields_by_name['last_block_id']._serialized_options = b'\310\336\037\000\342\336\037\013LastBlockID'
  _globals['_STATE'].fields_by_name['last_block_time']._loaded_options = None
  _globals['_STATE'].fields_by_name['last_block_time']._serialized_options = b'\310\336\037\000\220\337\037\001'
  _globals['_STATE'].fields_by_name['consensus_params']._loaded_options = None
  _globals['_STATE'].fields_by_name['consensus_params']._serialized_options = b'\310\336\037\000'
  _globals['_LEGACYABCIRESPONSES']._serialized_start=262
  _globals['_LEGACYABCIRESPONSES']._serialized_end=483
  _globals['_RESPONSEBEGINBLOCK']._serialized_start=485
  _globals['_RESPONSEBEGINBLOCK']._serialized_end=579
  _globals['_RESPONSEENDBLOCK']._serialized_start=582
  _globals['_RESPONSEENDBLOCK']._serialized_end=850
  _globals['_VALIDATORSINFO']._serialized_start=853
  _globals['_VALIDATORSINFO']._serialized_end=986
  _globals['_CONSENSUSPARAMSINFO']._serialized_start=989
  _globals['_CONSENSUSPARAMSINFO']._serialized_end=1142
  _globals['_ABCIRESPONSESINFO']._serialized_start=1145
  _globals['_ABCIRESPONSESINFO']._serialized_end=1375
  _globals['_VERSION']._serialized_start=1377
  _globals['_VERSION']._serialized_end=1481
  _globals['_STATE']._serialized_start=1484
  _globals['_STATE']._serialized_end=2349
# @@protoc_insertion_point(module_scope)
