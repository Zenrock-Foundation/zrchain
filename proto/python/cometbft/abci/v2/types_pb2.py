# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cometbft/abci/v2/types.proto
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
    'cometbft/abci/v2/types.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from cometbft.crypto.v1 import proof_pb2 as cometbft_dot_crypto_dot_v1_dot_proof__pb2
from cometbft.types.v2 import params_pb2 as cometbft_dot_types_dot_v2_dot_params__pb2
from cometbft.types.v2 import validator_pb2 as cometbft_dot_types_dot_v2_dot_validator__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x63ometbft/abci/v2/types.proto\x12\x10\x63ometbft.abci.v2\x1a\x1e\x63ometbft/crypto/v1/proof.proto\x1a\x1e\x63ometbft/types/v2/params.proto\x1a!cometbft/types/v2/validator.proto\x1a\x14gogoproto/gogo.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"\xcf\t\n\x07Request\x12\x33\n\x04\x65\x63ho\x18\x01 \x01(\x0b\x32\x1d.cometbft.abci.v2.EchoRequestH\x00R\x04\x65\x63ho\x12\x36\n\x05\x66lush\x18\x02 \x01(\x0b\x32\x1e.cometbft.abci.v2.FlushRequestH\x00R\x05\x66lush\x12\x33\n\x04info\x18\x03 \x01(\x0b\x32\x1d.cometbft.abci.v2.InfoRequestH\x00R\x04info\x12\x43\n\ninit_chain\x18\x05 \x01(\x0b\x32\".cometbft.abci.v2.InitChainRequestH\x00R\tinitChain\x12\x36\n\x05query\x18\x06 \x01(\x0b\x32\x1e.cometbft.abci.v2.QueryRequestH\x00R\x05query\x12=\n\x08\x63heck_tx\x18\x08 \x01(\x0b\x32 .cometbft.abci.v2.CheckTxRequestH\x00R\x07\x63heckTx\x12\x39\n\x06\x63ommit\x18\x0b \x01(\x0b\x32\x1f.cometbft.abci.v2.CommitRequestH\x00R\x06\x63ommit\x12O\n\x0elist_snapshots\x18\x0c \x01(\x0b\x32&.cometbft.abci.v2.ListSnapshotsRequestH\x00R\rlistSnapshots\x12O\n\x0eoffer_snapshot\x18\r \x01(\x0b\x32&.cometbft.abci.v2.OfferSnapshotRequestH\x00R\rofferSnapshot\x12\\\n\x13load_snapshot_chunk\x18\x0e \x01(\x0b\x32*.cometbft.abci.v2.LoadSnapshotChunkRequestH\x00R\x11loadSnapshotChunk\x12_\n\x14\x61pply_snapshot_chunk\x18\x0f \x01(\x0b\x32+.cometbft.abci.v2.ApplySnapshotChunkRequestH\x00R\x12\x61pplySnapshotChunk\x12U\n\x10prepare_proposal\x18\x10 \x01(\x0b\x32(.cometbft.abci.v2.PrepareProposalRequestH\x00R\x0fprepareProposal\x12U\n\x10process_proposal\x18\x11 \x01(\x0b\x32(.cometbft.abci.v2.ProcessProposalRequestH\x00R\x0fprocessProposal\x12\x46\n\x0b\x65xtend_vote\x18\x12 \x01(\x0b\x32#.cometbft.abci.v2.ExtendVoteRequestH\x00R\nextendVote\x12\x62\n\x15verify_vote_extension\x18\x13 \x01(\x0b\x32,.cometbft.abci.v2.VerifyVoteExtensionRequestH\x00R\x13verifyVoteExtension\x12O\n\x0e\x66inalize_block\x18\x14 \x01(\x0b\x32&.cometbft.abci.v2.FinalizeBlockRequestH\x00R\rfinalizeBlockB\x07\n\x05valueJ\x04\x08\x04\x10\x05J\x04\x08\x07\x10\x08J\x04\x08\t\x10\nJ\x04\x08\n\x10\x0b\"\'\n\x0b\x45\x63hoRequest\x12\x18\n\x07message\x18\x01 \x01(\tR\x07message\"\x0e\n\x0c\x46lushRequest\"\x90\x01\n\x0bInfoRequest\x12\x18\n\x07version\x18\x01 \x01(\tR\x07version\x12#\n\rblock_version\x18\x02 \x01(\x04R\x0c\x62lockVersion\x12\x1f\n\x0bp2p_version\x18\x03 \x01(\x04R\np2pVersion\x12!\n\x0c\x61\x62\x63i_version\x18\x04 \x01(\tR\x0b\x61\x62\x63iVersion\"\xce\x02\n\x10InitChainRequest\x12\x38\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01R\x04time\x12\x19\n\x08\x63hain_id\x18\x02 \x01(\tR\x07\x63hainId\x12M\n\x10\x63onsensus_params\x18\x03 \x01(\x0b\x32\".cometbft.types.v2.ConsensusParamsR\x0f\x63onsensusParams\x12G\n\nvalidators\x18\x04 \x03(\x0b\x32!.cometbft.abci.v2.ValidatorUpdateB\x04\xc8\xde\x1f\x00R\nvalidators\x12&\n\x0f\x61pp_state_bytes\x18\x05 \x01(\x0cR\rappStateBytes\x12%\n\x0einitial_height\x18\x06 \x01(\x03R\rinitialHeight\"d\n\x0cQueryRequest\x12\x12\n\x04\x64\x61ta\x18\x01 \x01(\x0cR\x04\x64\x61ta\x12\x12\n\x04path\x18\x02 \x01(\tR\x04path\x12\x16\n\x06height\x18\x03 \x01(\x03R\x06height\x12\x14\n\x05prove\x18\x04 \x01(\x08R\x05prove\"Y\n\x0e\x43heckTxRequest\x12\x0e\n\x02tx\x18\x01 \x01(\x0cR\x02tx\x12\x31\n\x04type\x18\x03 \x01(\x0e\x32\x1d.cometbft.abci.v2.CheckTxTypeR\x04typeJ\x04\x08\x02\x10\x03\"\x0f\n\rCommitRequest\"\x16\n\x14ListSnapshotsRequest\"i\n\x14OfferSnapshotRequest\x12\x36\n\x08snapshot\x18\x01 \x01(\x0b\x32\x1a.cometbft.abci.v2.SnapshotR\x08snapshot\x12\x19\n\x08\x61pp_hash\x18\x02 \x01(\x0cR\x07\x61ppHash\"`\n\x18LoadSnapshotChunkRequest\x12\x16\n\x06height\x18\x01 \x01(\x04R\x06height\x12\x16\n\x06\x66ormat\x18\x02 \x01(\rR\x06\x66ormat\x12\x14\n\x05\x63hunk\x18\x03 \x01(\rR\x05\x63hunk\"_\n\x19\x41pplySnapshotChunkRequest\x12\x14\n\x05index\x18\x01 \x01(\rR\x05index\x12\x14\n\x05\x63hunk\x18\x02 \x01(\x0cR\x05\x63hunk\x12\x16\n\x06sender\x18\x03 \x01(\tR\x06sender\"\x9a\x03\n\x16PrepareProposalRequest\x12 \n\x0cmax_tx_bytes\x18\x01 \x01(\x03R\nmaxTxBytes\x12\x10\n\x03txs\x18\x02 \x03(\x0cR\x03txs\x12V\n\x11local_last_commit\x18\x03 \x01(\x0b\x32$.cometbft.abci.v2.ExtendedCommitInfoB\x04\xc8\xde\x1f\x00R\x0flocalLastCommit\x12\x45\n\x0bmisbehavior\x18\x04 \x03(\x0b\x32\x1d.cometbft.abci.v2.MisbehaviorB\x04\xc8\xde\x1f\x00R\x0bmisbehavior\x12\x16\n\x06height\x18\x05 \x01(\x03R\x06height\x12\x38\n\x04time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01R\x04time\x12\x30\n\x14next_validators_hash\x18\x07 \x01(\x0cR\x12nextValidatorsHash\x12)\n\x10proposer_address\x18\x08 \x01(\x0cR\x0fproposerAddress\"\x8a\x03\n\x16ProcessProposalRequest\x12\x10\n\x03txs\x18\x01 \x03(\x0cR\x03txs\x12T\n\x14proposed_last_commit\x18\x02 \x01(\x0b\x32\x1c.cometbft.abci.v2.CommitInfoB\x04\xc8\xde\x1f\x00R\x12proposedLastCommit\x12\x45\n\x0bmisbehavior\x18\x03 \x03(\x0b\x32\x1d.cometbft.abci.v2.MisbehaviorB\x04\xc8\xde\x1f\x00R\x0bmisbehavior\x12\x12\n\x04hash\x18\x04 \x01(\x0cR\x04hash\x12\x16\n\x06height\x18\x05 \x01(\x03R\x06height\x12\x38\n\x04time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01R\x04time\x12\x30\n\x14next_validators_hash\x18\x07 \x01(\x0cR\x12nextValidatorsHash\x12)\n\x10proposer_address\x18\x08 \x01(\x0cR\x0fproposerAddress\"\x85\x03\n\x11\x45xtendVoteRequest\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\x12\x16\n\x06height\x18\x02 \x01(\x03R\x06height\x12\x38\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01R\x04time\x12\x10\n\x03txs\x18\x04 \x03(\x0cR\x03txs\x12T\n\x14proposed_last_commit\x18\x05 \x01(\x0b\x32\x1c.cometbft.abci.v2.CommitInfoB\x04\xc8\xde\x1f\x00R\x12proposedLastCommit\x12\x45\n\x0bmisbehavior\x18\x06 \x03(\x0b\x32\x1d.cometbft.abci.v2.MisbehaviorB\x04\xc8\xde\x1f\x00R\x0bmisbehavior\x12\x30\n\x14next_validators_hash\x18\x07 \x01(\x0cR\x12nextValidatorsHash\x12)\n\x10proposer_address\x18\x08 \x01(\x0cR\x0fproposerAddress\"\xcf\x01\n\x1aVerifyVoteExtensionRequest\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\x12+\n\x11validator_address\x18\x02 \x01(\x0cR\x10validatorAddress\x12\x16\n\x06height\x18\x03 \x01(\x03R\x06height\x12%\n\x0evote_extension\x18\x04 \x01(\x0cR\rvoteExtension\x12\x31\n\x15non_rp_vote_extension\x18\x05 \x01(\x0cR\x12nonRpVoteExtension\"\xb2\x03\n\x14\x46inalizeBlockRequest\x12\x10\n\x03txs\x18\x01 \x03(\x0cR\x03txs\x12R\n\x13\x64\x65\x63ided_last_commit\x18\x02 \x01(\x0b\x32\x1c.cometbft.abci.v2.CommitInfoB\x04\xc8\xde\x1f\x00R\x11\x64\x65\x63idedLastCommit\x12\x45\n\x0bmisbehavior\x18\x03 \x03(\x0b\x32\x1d.cometbft.abci.v2.MisbehaviorB\x04\xc8\xde\x1f\x00R\x0bmisbehavior\x12\x12\n\x04hash\x18\x04 \x01(\x0cR\x04hash\x12\x16\n\x06height\x18\x05 \x01(\x03R\x06height\x12\x38\n\x04time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01R\x04time\x12\x30\n\x14next_validators_hash\x18\x07 \x01(\x0cR\x12nextValidatorsHash\x12)\n\x10proposer_address\x18\x08 \x01(\x0cR\x0fproposerAddress\x12*\n\x11syncing_to_height\x18\t \x01(\x03R\x0fsyncingToHeight\"\xa5\n\n\x08Response\x12\x43\n\texception\x18\x01 \x01(\x0b\x32#.cometbft.abci.v2.ExceptionResponseH\x00R\texception\x12\x34\n\x04\x65\x63ho\x18\x02 \x01(\x0b\x32\x1e.cometbft.abci.v2.EchoResponseH\x00R\x04\x65\x63ho\x12\x37\n\x05\x66lush\x18\x03 \x01(\x0b\x32\x1f.cometbft.abci.v2.FlushResponseH\x00R\x05\x66lush\x12\x34\n\x04info\x18\x04 \x01(\x0b\x32\x1e.cometbft.abci.v2.InfoResponseH\x00R\x04info\x12\x44\n\ninit_chain\x18\x06 \x01(\x0b\x32#.cometbft.abci.v2.InitChainResponseH\x00R\tinitChain\x12\x37\n\x05query\x18\x07 \x01(\x0b\x32\x1f.cometbft.abci.v2.QueryResponseH\x00R\x05query\x12>\n\x08\x63heck_tx\x18\t \x01(\x0b\x32!.cometbft.abci.v2.CheckTxResponseH\x00R\x07\x63heckTx\x12:\n\x06\x63ommit\x18\x0c \x01(\x0b\x32 .cometbft.abci.v2.CommitResponseH\x00R\x06\x63ommit\x12P\n\x0elist_snapshots\x18\r \x01(\x0b\x32\'.cometbft.abci.v2.ListSnapshotsResponseH\x00R\rlistSnapshots\x12P\n\x0eoffer_snapshot\x18\x0e \x01(\x0b\x32\'.cometbft.abci.v2.OfferSnapshotResponseH\x00R\rofferSnapshot\x12]\n\x13load_snapshot_chunk\x18\x0f \x01(\x0b\x32+.cometbft.abci.v2.LoadSnapshotChunkResponseH\x00R\x11loadSnapshotChunk\x12`\n\x14\x61pply_snapshot_chunk\x18\x10 \x01(\x0b\x32,.cometbft.abci.v2.ApplySnapshotChunkResponseH\x00R\x12\x61pplySnapshotChunk\x12V\n\x10prepare_proposal\x18\x11 \x01(\x0b\x32).cometbft.abci.v2.PrepareProposalResponseH\x00R\x0fprepareProposal\x12V\n\x10process_proposal\x18\x12 \x01(\x0b\x32).cometbft.abci.v2.ProcessProposalResponseH\x00R\x0fprocessProposal\x12G\n\x0b\x65xtend_vote\x18\x13 \x01(\x0b\x32$.cometbft.abci.v2.ExtendVoteResponseH\x00R\nextendVote\x12\x63\n\x15verify_vote_extension\x18\x14 \x01(\x0b\x32-.cometbft.abci.v2.VerifyVoteExtensionResponseH\x00R\x13verifyVoteExtension\x12P\n\x0e\x66inalize_block\x18\x15 \x01(\x0b\x32\'.cometbft.abci.v2.FinalizeBlockResponseH\x00R\rfinalizeBlockB\x07\n\x05valueJ\x04\x08\x05\x10\x06J\x04\x08\x08\x10\tJ\x04\x08\n\x10\x0bJ\x04\x08\x0b\x10\x0c\")\n\x11\x45xceptionResponse\x12\x14\n\x05\x65rror\x18\x01 \x01(\tR\x05\x65rror\"(\n\x0c\x45\x63hoResponse\x12\x18\n\x07message\x18\x01 \x01(\tR\x07message\"\x0f\n\rFlushResponse\"\xfb\x02\n\x0cInfoResponse\x12\x12\n\x04\x64\x61ta\x18\x01 \x01(\tR\x04\x64\x61ta\x12\x18\n\x07version\x18\x02 \x01(\tR\x07version\x12\x1f\n\x0b\x61pp_version\x18\x03 \x01(\x04R\nappVersion\x12*\n\x11last_block_height\x18\x04 \x01(\x03R\x0flastBlockHeight\x12-\n\x13last_block_app_hash\x18\x05 \x01(\x0cR\x10lastBlockAppHash\x12[\n\x0flane_priorities\x18\x06 \x03(\x0b\x32\x32.cometbft.abci.v2.InfoResponse.LanePrioritiesEntryR\x0elanePriorities\x12!\n\x0c\x64\x65\x66\x61ult_lane\x18\x07 \x01(\tR\x0b\x64\x65\x66\x61ultLane\x1a\x41\n\x13LanePrioritiesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\rR\x05value:\x02\x38\x01\"\xc6\x01\n\x11InitChainResponse\x12M\n\x10\x63onsensus_params\x18\x01 \x01(\x0b\x32\".cometbft.types.v2.ConsensusParamsR\x0f\x63onsensusParams\x12G\n\nvalidators\x18\x02 \x03(\x0b\x32!.cometbft.abci.v2.ValidatorUpdateB\x04\xc8\xde\x1f\x00R\nvalidators\x12\x19\n\x08\x61pp_hash\x18\x03 \x01(\x0cR\x07\x61ppHash\"\xf8\x01\n\rQueryResponse\x12\x12\n\x04\x63ode\x18\x01 \x01(\rR\x04\x63ode\x12\x10\n\x03log\x18\x03 \x01(\tR\x03log\x12\x12\n\x04info\x18\x04 \x01(\tR\x04info\x12\x14\n\x05index\x18\x05 \x01(\x03R\x05index\x12\x10\n\x03key\x18\x06 \x01(\x0cR\x03key\x12\x14\n\x05value\x18\x07 \x01(\x0cR\x05value\x12\x39\n\tproof_ops\x18\x08 \x01(\x0b\x32\x1c.cometbft.crypto.v1.ProofOpsR\x08proofOps\x12\x16\n\x06height\x18\t \x01(\x03R\x06height\x12\x1c\n\tcodespace\x18\n \x01(\tR\tcodespace\"\xc4\x02\n\x0f\x43heckTxResponse\x12\x12\n\x04\x63ode\x18\x01 \x01(\rR\x04\x63ode\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\x0cR\x04\x64\x61ta\x12\x10\n\x03log\x18\x03 \x01(\tR\x03log\x12\x12\n\x04info\x18\x04 \x01(\tR\x04info\x12\x1e\n\ngas_wanted\x18\x05 \x01(\x03R\ngas_wanted\x12\x1a\n\x08gas_used\x18\x06 \x01(\x03R\x08gas_used\x12I\n\x06\x65vents\x18\x07 \x03(\x0b\x32\x17.cometbft.abci.v2.EventB\x18\xc8\xde\x1f\x00\xea\xde\x1f\x10\x65vents,omitemptyR\x06\x65vents\x12\x1c\n\tcodespace\x18\x08 \x01(\tR\tcodespace\x12\x17\n\x07lane_id\x18\x0c \x01(\tR\x06laneIdJ\x04\x08\t\x10\x0cR\x06senderR\x08priorityR\rmempool_error\"A\n\x0e\x43ommitResponse\x12#\n\rretain_height\x18\x03 \x01(\x03R\x0cretainHeightJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\"Q\n\x15ListSnapshotsResponse\x12\x38\n\tsnapshots\x18\x01 \x03(\x0b\x32\x1a.cometbft.abci.v2.SnapshotR\tsnapshots\"V\n\x15OfferSnapshotResponse\x12=\n\x06result\x18\x01 \x01(\x0e\x32%.cometbft.abci.v2.OfferSnapshotResultR\x06result\"1\n\x19LoadSnapshotChunkResponse\x12\x14\n\x05\x63hunk\x18\x01 \x01(\x0cR\x05\x63hunk\"\xae\x01\n\x1a\x41pplySnapshotChunkResponse\x12\x42\n\x06result\x18\x01 \x01(\x0e\x32*.cometbft.abci.v2.ApplySnapshotChunkResultR\x06result\x12%\n\x0erefetch_chunks\x18\x02 \x03(\rR\rrefetchChunks\x12%\n\x0ereject_senders\x18\x03 \x03(\tR\rrejectSenders\"+\n\x17PrepareProposalResponse\x12\x10\n\x03txs\x18\x01 \x03(\x0cR\x03txs\"Z\n\x17ProcessProposalResponse\x12?\n\x06status\x18\x01 \x01(\x0e\x32\'.cometbft.abci.v2.ProcessProposalStatusR\x06status\"e\n\x12\x45xtendVoteResponse\x12%\n\x0evote_extension\x18\x01 \x01(\x0cR\rvoteExtension\x12(\n\x10non_rp_extension\x18\x02 \x01(\x0cR\x0enonRpExtension\"b\n\x1bVerifyVoteExtensionResponse\x12\x43\n\x06status\x18\x01 \x01(\x0e\x32+.cometbft.abci.v2.VerifyVoteExtensionStatusR\x06status\"\xbd\x03\n\x15\x46inalizeBlockResponse\x12I\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x17.cometbft.abci.v2.EventB\x18\xc8\xde\x1f\x00\xea\xde\x1f\x10\x65vents,omitemptyR\x06\x65vents\x12=\n\ntx_results\x18\x02 \x03(\x0b\x32\x1e.cometbft.abci.v2.ExecTxResultR\ttxResults\x12T\n\x11validator_updates\x18\x03 \x03(\x0b\x32!.cometbft.abci.v2.ValidatorUpdateB\x04\xc8\xde\x1f\x00R\x10validatorUpdates\x12Z\n\x17\x63onsensus_param_updates\x18\x04 \x01(\x0b\x32\".cometbft.types.v2.ConsensusParamsR\x15\x63onsensusParamUpdates\x12\x19\n\x08\x61pp_hash\x18\x05 \x01(\x0cR\x07\x61ppHash\x12M\n\x10next_block_delay\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationB\x08\xc8\xde\x1f\x00\x98\xdf\x1f\x01R\x0enextBlockDelay\"Z\n\nCommitInfo\x12\x14\n\x05round\x18\x01 \x01(\x05R\x05round\x12\x36\n\x05votes\x18\x02 \x03(\x0b\x32\x1a.cometbft.abci.v2.VoteInfoB\x04\xc8\xde\x1f\x00R\x05votes\"j\n\x12\x45xtendedCommitInfo\x12\x14\n\x05round\x18\x01 \x01(\x05R\x05round\x12>\n\x05votes\x18\x02 \x03(\x0b\x32\".cometbft.abci.v2.ExtendedVoteInfoB\x04\xc8\xde\x1f\x00R\x05votes\"{\n\x05\x45vent\x12\x12\n\x04type\x18\x01 \x01(\tR\x04type\x12^\n\nattributes\x18\x02 \x03(\x0b\x32 .cometbft.abci.v2.EventAttributeB\x1c\xc8\xde\x1f\x00\xea\xde\x1f\x14\x61ttributes,omitemptyR\nattributes\"N\n\x0e\x45ventAttribute\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value\x12\x14\n\x05index\x18\x03 \x01(\x08R\x05index\"\x81\x02\n\x0c\x45xecTxResult\x12\x12\n\x04\x63ode\x18\x01 \x01(\rR\x04\x63ode\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\x0cR\x04\x64\x61ta\x12\x10\n\x03log\x18\x03 \x01(\tR\x03log\x12\x12\n\x04info\x18\x04 \x01(\tR\x04info\x12\x1e\n\ngas_wanted\x18\x05 \x01(\x03R\ngas_wanted\x12\x1a\n\x08gas_used\x18\x06 \x01(\x03R\x08gas_used\x12I\n\x06\x65vents\x18\x07 \x03(\x0b\x32\x17.cometbft.abci.v2.EventB\x18\xc8\xde\x1f\x00\xea\xde\x1f\x10\x65vents,omitemptyR\x06\x65vents\x12\x1c\n\tcodespace\x18\x08 \x01(\tR\tcodespace\"\x86\x01\n\x08TxResult\x12\x16\n\x06height\x18\x01 \x01(\x03R\x06height\x12\x14\n\x05index\x18\x02 \x01(\rR\x05index\x12\x0e\n\x02tx\x18\x03 \x01(\x0cR\x02tx\x12<\n\x06result\x18\x04 \x01(\x0b\x32\x1e.cometbft.abci.v2.ExecTxResultB\x04\xc8\xde\x1f\x00R\x06result\";\n\tValidator\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0cR\x07\x61\x64\x64ress\x12\x14\n\x05power\x18\x03 \x01(\x03R\x05power\"s\n\x0fValidatorUpdate\x12\x14\n\x05power\x18\x02 \x01(\x03R\x05power\x12\"\n\rpub_key_bytes\x18\x03 \x01(\x0cR\x0bpubKeyBytes\x12 \n\x0cpub_key_type\x18\x04 \x01(\tR\npubKeyTypeJ\x04\x08\x01\x10\x02\"\x95\x01\n\x08VoteInfo\x12?\n\tvalidator\x18\x01 \x01(\x0b\x32\x1b.cometbft.abci.v2.ValidatorB\x04\xc8\xde\x1f\x00R\tvalidator\x12\x42\n\rblock_id_flag\x18\x03 \x01(\x0e\x32\x1e.cometbft.types.v2.BlockIDFlagR\x0b\x62lockIdFlagJ\x04\x08\x02\x10\x03\"\xe5\x02\n\x10\x45xtendedVoteInfo\x12?\n\tvalidator\x18\x01 \x01(\x0b\x32\x1b.cometbft.abci.v2.ValidatorB\x04\xc8\xde\x1f\x00R\tvalidator\x12%\n\x0evote_extension\x18\x03 \x01(\x0cR\rvoteExtension\x12/\n\x13\x65xtension_signature\x18\x04 \x01(\x0cR\x12\x65xtensionSignature\x12\x42\n\rblock_id_flag\x18\x05 \x01(\x0e\x32\x1e.cometbft.types.v2.BlockIDFlagR\x0b\x62lockIdFlag\x12\x31\n\x15non_rp_vote_extension\x18\x06 \x01(\x0cR\x12nonRpVoteExtension\x12;\n\x1anon_rp_extension_signature\x18\x07 \x01(\x0cR\x17nonRpExtensionSignatureJ\x04\x08\x02\x10\x03\"\x85\x02\n\x0bMisbehavior\x12\x35\n\x04type\x18\x01 \x01(\x0e\x32!.cometbft.abci.v2.MisbehaviorTypeR\x04type\x12?\n\tvalidator\x18\x02 \x01(\x0b\x32\x1b.cometbft.abci.v2.ValidatorB\x04\xc8\xde\x1f\x00R\tvalidator\x12\x16\n\x06height\x18\x03 \x01(\x03R\x06height\x12\x38\n\x04time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01R\x04time\x12,\n\x12total_voting_power\x18\x05 \x01(\x03R\x10totalVotingPower\"\x82\x01\n\x08Snapshot\x12\x16\n\x06height\x18\x01 \x01(\x04R\x06height\x12\x16\n\x06\x66ormat\x18\x02 \x01(\rR\x06\x66ormat\x12\x16\n\x06\x63hunks\x18\x03 \x01(\rR\x06\x63hunks\x12\x12\n\x04hash\x18\x04 \x01(\x0cR\x04hash\x12\x1a\n\x08metadata\x18\x05 \x01(\x0cR\x08metadata*b\n\x0b\x43heckTxType\x12\x19\n\x15\x43HECK_TX_TYPE_UNKNOWN\x10\x00\x12\x19\n\x15\x43HECK_TX_TYPE_RECHECK\x10\x01\x12\x17\n\x13\x43HECK_TX_TYPE_CHECK\x10\x02\x1a\x04\x88\xa3\x1e\x00*\xf5\x01\n\x13OfferSnapshotResult\x12!\n\x1dOFFER_SNAPSHOT_RESULT_UNKNOWN\x10\x00\x12 \n\x1cOFFER_SNAPSHOT_RESULT_ACCEPT\x10\x01\x12\x1f\n\x1bOFFER_SNAPSHOT_RESULT_ABORT\x10\x02\x12 \n\x1cOFFER_SNAPSHOT_RESULT_REJECT\x10\x03\x12\'\n#OFFER_SNAPSHOT_RESULT_REJECT_FORMAT\x10\x04\x12\'\n#OFFER_SNAPSHOT_RESULT_REJECT_SENDER\x10\x05\x1a\x04\x88\xa3\x1e\x00*\xa0\x02\n\x18\x41pplySnapshotChunkResult\x12\'\n#APPLY_SNAPSHOT_CHUNK_RESULT_UNKNOWN\x10\x00\x12&\n\"APPLY_SNAPSHOT_CHUNK_RESULT_ACCEPT\x10\x01\x12%\n!APPLY_SNAPSHOT_CHUNK_RESULT_ABORT\x10\x02\x12%\n!APPLY_SNAPSHOT_CHUNK_RESULT_RETRY\x10\x03\x12.\n*APPLY_SNAPSHOT_CHUNK_RESULT_RETRY_SNAPSHOT\x10\x04\x12/\n+APPLY_SNAPSHOT_CHUNK_RESULT_REJECT_SNAPSHOT\x10\x05\x1a\x04\x88\xa3\x1e\x00*\x8a\x01\n\x15ProcessProposalStatus\x12#\n\x1fPROCESS_PROPOSAL_STATUS_UNKNOWN\x10\x00\x12\"\n\x1ePROCESS_PROPOSAL_STATUS_ACCEPT\x10\x01\x12\"\n\x1ePROCESS_PROPOSAL_STATUS_REJECT\x10\x02\x1a\x04\x88\xa3\x1e\x00*\x9d\x01\n\x19VerifyVoteExtensionStatus\x12(\n$VERIFY_VOTE_EXTENSION_STATUS_UNKNOWN\x10\x00\x12\'\n#VERIFY_VOTE_EXTENSION_STATUS_ACCEPT\x10\x01\x12\'\n#VERIFY_VOTE_EXTENSION_STATUS_REJECT\x10\x02\x1a\x04\x88\xa3\x1e\x00*\x84\x01\n\x0fMisbehaviorType\x12\x1c\n\x18MISBEHAVIOR_TYPE_UNKNOWN\x10\x00\x12#\n\x1fMISBEHAVIOR_TYPE_DUPLICATE_VOTE\x10\x01\x12(\n$MISBEHAVIOR_TYPE_LIGHT_CLIENT_ATTACK\x10\x02\x1a\x04\x88\xa3\x1e\x00\x42\x33Z1github.com/cometbft/cometbft/api/cometbft/abci/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cometbft.abci.v2.types_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z1github.com/cometbft/cometbft/api/cometbft/abci/v2'
  _globals['_CHECKTXTYPE']._loaded_options = None
  _globals['_CHECKTXTYPE']._serialized_options = b'\210\243\036\000'
  _globals['_OFFERSNAPSHOTRESULT']._loaded_options = None
  _globals['_OFFERSNAPSHOTRESULT']._serialized_options = b'\210\243\036\000'
  _globals['_APPLYSNAPSHOTCHUNKRESULT']._loaded_options = None
  _globals['_APPLYSNAPSHOTCHUNKRESULT']._serialized_options = b'\210\243\036\000'
  _globals['_PROCESSPROPOSALSTATUS']._loaded_options = None
  _globals['_PROCESSPROPOSALSTATUS']._serialized_options = b'\210\243\036\000'
  _globals['_VERIFYVOTEEXTENSIONSTATUS']._loaded_options = None
  _globals['_VERIFYVOTEEXTENSIONSTATUS']._serialized_options = b'\210\243\036\000'
  _globals['_MISBEHAVIORTYPE']._loaded_options = None
  _globals['_MISBEHAVIORTYPE']._serialized_options = b'\210\243\036\000'
  _globals['_INITCHAINREQUEST'].fields_by_name['time']._loaded_options = None
  _globals['_INITCHAINREQUEST'].fields_by_name['time']._serialized_options = b'\310\336\037\000\220\337\037\001'
  _globals['_INITCHAINREQUEST'].fields_by_name['validators']._loaded_options = None
  _globals['_INITCHAINREQUEST'].fields_by_name['validators']._serialized_options = b'\310\336\037\000'
  _globals['_PREPAREPROPOSALREQUEST'].fields_by_name['local_last_commit']._loaded_options = None
  _globals['_PREPAREPROPOSALREQUEST'].fields_by_name['local_last_commit']._serialized_options = b'\310\336\037\000'
  _globals['_PREPAREPROPOSALREQUEST'].fields_by_name['misbehavior']._loaded_options = None
  _globals['_PREPAREPROPOSALREQUEST'].fields_by_name['misbehavior']._serialized_options = b'\310\336\037\000'
  _globals['_PREPAREPROPOSALREQUEST'].fields_by_name['time']._loaded_options = None
  _globals['_PREPAREPROPOSALREQUEST'].fields_by_name['time']._serialized_options = b'\310\336\037\000\220\337\037\001'
  _globals['_PROCESSPROPOSALREQUEST'].fields_by_name['proposed_last_commit']._loaded_options = None
  _globals['_PROCESSPROPOSALREQUEST'].fields_by_name['proposed_last_commit']._serialized_options = b'\310\336\037\000'
  _globals['_PROCESSPROPOSALREQUEST'].fields_by_name['misbehavior']._loaded_options = None
  _globals['_PROCESSPROPOSALREQUEST'].fields_by_name['misbehavior']._serialized_options = b'\310\336\037\000'
  _globals['_PROCESSPROPOSALREQUEST'].fields_by_name['time']._loaded_options = None
  _globals['_PROCESSPROPOSALREQUEST'].fields_by_name['time']._serialized_options = b'\310\336\037\000\220\337\037\001'
  _globals['_EXTENDVOTEREQUEST'].fields_by_name['time']._loaded_options = None
  _globals['_EXTENDVOTEREQUEST'].fields_by_name['time']._serialized_options = b'\310\336\037\000\220\337\037\001'
  _globals['_EXTENDVOTEREQUEST'].fields_by_name['proposed_last_commit']._loaded_options = None
  _globals['_EXTENDVOTEREQUEST'].fields_by_name['proposed_last_commit']._serialized_options = b'\310\336\037\000'
  _globals['_EXTENDVOTEREQUEST'].fields_by_name['misbehavior']._loaded_options = None
  _globals['_EXTENDVOTEREQUEST'].fields_by_name['misbehavior']._serialized_options = b'\310\336\037\000'
  _globals['_FINALIZEBLOCKREQUEST'].fields_by_name['decided_last_commit']._loaded_options = None
  _globals['_FINALIZEBLOCKREQUEST'].fields_by_name['decided_last_commit']._serialized_options = b'\310\336\037\000'
  _globals['_FINALIZEBLOCKREQUEST'].fields_by_name['misbehavior']._loaded_options = None
  _globals['_FINALIZEBLOCKREQUEST'].fields_by_name['misbehavior']._serialized_options = b'\310\336\037\000'
  _globals['_FINALIZEBLOCKREQUEST'].fields_by_name['time']._loaded_options = None
  _globals['_FINALIZEBLOCKREQUEST'].fields_by_name['time']._serialized_options = b'\310\336\037\000\220\337\037\001'
  _globals['_INFORESPONSE_LANEPRIORITIESENTRY']._loaded_options = None
  _globals['_INFORESPONSE_LANEPRIORITIESENTRY']._serialized_options = b'8\001'
  _globals['_INITCHAINRESPONSE'].fields_by_name['validators']._loaded_options = None
  _globals['_INITCHAINRESPONSE'].fields_by_name['validators']._serialized_options = b'\310\336\037\000'
  _globals['_CHECKTXRESPONSE'].fields_by_name['events']._loaded_options = None
  _globals['_CHECKTXRESPONSE'].fields_by_name['events']._serialized_options = b'\310\336\037\000\352\336\037\020events,omitempty'
  _globals['_FINALIZEBLOCKRESPONSE'].fields_by_name['events']._loaded_options = None
  _globals['_FINALIZEBLOCKRESPONSE'].fields_by_name['events']._serialized_options = b'\310\336\037\000\352\336\037\020events,omitempty'
  _globals['_FINALIZEBLOCKRESPONSE'].fields_by_name['validator_updates']._loaded_options = None
  _globals['_FINALIZEBLOCKRESPONSE'].fields_by_name['validator_updates']._serialized_options = b'\310\336\037\000'
  _globals['_FINALIZEBLOCKRESPONSE'].fields_by_name['next_block_delay']._loaded_options = None
  _globals['_FINALIZEBLOCKRESPONSE'].fields_by_name['next_block_delay']._serialized_options = b'\310\336\037\000\230\337\037\001'
  _globals['_COMMITINFO'].fields_by_name['votes']._loaded_options = None
  _globals['_COMMITINFO'].fields_by_name['votes']._serialized_options = b'\310\336\037\000'
  _globals['_EXTENDEDCOMMITINFO'].fields_by_name['votes']._loaded_options = None
  _globals['_EXTENDEDCOMMITINFO'].fields_by_name['votes']._serialized_options = b'\310\336\037\000'
  _globals['_EVENT'].fields_by_name['attributes']._loaded_options = None
  _globals['_EVENT'].fields_by_name['attributes']._serialized_options = b'\310\336\037\000\352\336\037\024attributes,omitempty'
  _globals['_EXECTXRESULT'].fields_by_name['events']._loaded_options = None
  _globals['_EXECTXRESULT'].fields_by_name['events']._serialized_options = b'\310\336\037\000\352\336\037\020events,omitempty'
  _globals['_TXRESULT'].fields_by_name['result']._loaded_options = None
  _globals['_TXRESULT'].fields_by_name['result']._serialized_options = b'\310\336\037\000'
  _globals['_VOTEINFO'].fields_by_name['validator']._loaded_options = None
  _globals['_VOTEINFO'].fields_by_name['validator']._serialized_options = b'\310\336\037\000'
  _globals['_EXTENDEDVOTEINFO'].fields_by_name['validator']._loaded_options = None
  _globals['_EXTENDEDVOTEINFO'].fields_by_name['validator']._serialized_options = b'\310\336\037\000'
  _globals['_MISBEHAVIOR'].fields_by_name['validator']._loaded_options = None
  _globals['_MISBEHAVIOR'].fields_by_name['validator']._serialized_options = b'\310\336\037\000'
  _globals['_MISBEHAVIOR'].fields_by_name['time']._loaded_options = None
  _globals['_MISBEHAVIOR'].fields_by_name['time']._serialized_options = b'\310\336\037\000\220\337\037\001'
  _globals['_CHECKTXTYPE']._serialized_start=10122
  _globals['_CHECKTXTYPE']._serialized_end=10220
  _globals['_OFFERSNAPSHOTRESULT']._serialized_start=10223
  _globals['_OFFERSNAPSHOTRESULT']._serialized_end=10468
  _globals['_APPLYSNAPSHOTCHUNKRESULT']._serialized_start=10471
  _globals['_APPLYSNAPSHOTCHUNKRESULT']._serialized_end=10759
  _globals['_PROCESSPROPOSALSTATUS']._serialized_start=10762
  _globals['_PROCESSPROPOSALSTATUS']._serialized_end=10900
  _globals['_VERIFYVOTEEXTENSIONSTATUS']._serialized_start=10903
  _globals['_VERIFYVOTEEXTENSIONSTATUS']._serialized_end=11060
  _globals['_MISBEHAVIORTYPE']._serialized_start=11063
  _globals['_MISBEHAVIORTYPE']._serialized_end=11195
  _globals['_REQUEST']._serialized_start=237
  _globals['_REQUEST']._serialized_end=1468
  _globals['_ECHOREQUEST']._serialized_start=1470
  _globals['_ECHOREQUEST']._serialized_end=1509
  _globals['_FLUSHREQUEST']._serialized_start=1511
  _globals['_FLUSHREQUEST']._serialized_end=1525
  _globals['_INFOREQUEST']._serialized_start=1528
  _globals['_INFOREQUEST']._serialized_end=1672
  _globals['_INITCHAINREQUEST']._serialized_start=1675
  _globals['_INITCHAINREQUEST']._serialized_end=2009
  _globals['_QUERYREQUEST']._serialized_start=2011
  _globals['_QUERYREQUEST']._serialized_end=2111
  _globals['_CHECKTXREQUEST']._serialized_start=2113
  _globals['_CHECKTXREQUEST']._serialized_end=2202
  _globals['_COMMITREQUEST']._serialized_start=2204
  _globals['_COMMITREQUEST']._serialized_end=2219
  _globals['_LISTSNAPSHOTSREQUEST']._serialized_start=2221
  _globals['_LISTSNAPSHOTSREQUEST']._serialized_end=2243
  _globals['_OFFERSNAPSHOTREQUEST']._serialized_start=2245
  _globals['_OFFERSNAPSHOTREQUEST']._serialized_end=2350
  _globals['_LOADSNAPSHOTCHUNKREQUEST']._serialized_start=2352
  _globals['_LOADSNAPSHOTCHUNKREQUEST']._serialized_end=2448
  _globals['_APPLYSNAPSHOTCHUNKREQUEST']._serialized_start=2450
  _globals['_APPLYSNAPSHOTCHUNKREQUEST']._serialized_end=2545
  _globals['_PREPAREPROPOSALREQUEST']._serialized_start=2548
  _globals['_PREPAREPROPOSALREQUEST']._serialized_end=2958
  _globals['_PROCESSPROPOSALREQUEST']._serialized_start=2961
  _globals['_PROCESSPROPOSALREQUEST']._serialized_end=3355
  _globals['_EXTENDVOTEREQUEST']._serialized_start=3358
  _globals['_EXTENDVOTEREQUEST']._serialized_end=3747
  _globals['_VERIFYVOTEEXTENSIONREQUEST']._serialized_start=3750
  _globals['_VERIFYVOTEEXTENSIONREQUEST']._serialized_end=3957
  _globals['_FINALIZEBLOCKREQUEST']._serialized_start=3960
  _globals['_FINALIZEBLOCKREQUEST']._serialized_end=4394
  _globals['_RESPONSE']._serialized_start=4397
  _globals['_RESPONSE']._serialized_end=5714
  _globals['_EXCEPTIONRESPONSE']._serialized_start=5716
  _globals['_EXCEPTIONRESPONSE']._serialized_end=5757
  _globals['_ECHORESPONSE']._serialized_start=5759
  _globals['_ECHORESPONSE']._serialized_end=5799
  _globals['_FLUSHRESPONSE']._serialized_start=5801
  _globals['_FLUSHRESPONSE']._serialized_end=5816
  _globals['_INFORESPONSE']._serialized_start=5819
  _globals['_INFORESPONSE']._serialized_end=6198
  _globals['_INFORESPONSE_LANEPRIORITIESENTRY']._serialized_start=6133
  _globals['_INFORESPONSE_LANEPRIORITIESENTRY']._serialized_end=6198
  _globals['_INITCHAINRESPONSE']._serialized_start=6201
  _globals['_INITCHAINRESPONSE']._serialized_end=6399
  _globals['_QUERYRESPONSE']._serialized_start=6402
  _globals['_QUERYRESPONSE']._serialized_end=6650
  _globals['_CHECKTXRESPONSE']._serialized_start=6653
  _globals['_CHECKTXRESPONSE']._serialized_end=6977
  _globals['_COMMITRESPONSE']._serialized_start=6979
  _globals['_COMMITRESPONSE']._serialized_end=7044
  _globals['_LISTSNAPSHOTSRESPONSE']._serialized_start=7046
  _globals['_LISTSNAPSHOTSRESPONSE']._serialized_end=7127
  _globals['_OFFERSNAPSHOTRESPONSE']._serialized_start=7129
  _globals['_OFFERSNAPSHOTRESPONSE']._serialized_end=7215
  _globals['_LOADSNAPSHOTCHUNKRESPONSE']._serialized_start=7217
  _globals['_LOADSNAPSHOTCHUNKRESPONSE']._serialized_end=7266
  _globals['_APPLYSNAPSHOTCHUNKRESPONSE']._serialized_start=7269
  _globals['_APPLYSNAPSHOTCHUNKRESPONSE']._serialized_end=7443
  _globals['_PREPAREPROPOSALRESPONSE']._serialized_start=7445
  _globals['_PREPAREPROPOSALRESPONSE']._serialized_end=7488
  _globals['_PROCESSPROPOSALRESPONSE']._serialized_start=7490
  _globals['_PROCESSPROPOSALRESPONSE']._serialized_end=7580
  _globals['_EXTENDVOTERESPONSE']._serialized_start=7582
  _globals['_EXTENDVOTERESPONSE']._serialized_end=7683
  _globals['_VERIFYVOTEEXTENSIONRESPONSE']._serialized_start=7685
  _globals['_VERIFYVOTEEXTENSIONRESPONSE']._serialized_end=7783
  _globals['_FINALIZEBLOCKRESPONSE']._serialized_start=7786
  _globals['_FINALIZEBLOCKRESPONSE']._serialized_end=8231
  _globals['_COMMITINFO']._serialized_start=8233
  _globals['_COMMITINFO']._serialized_end=8323
  _globals['_EXTENDEDCOMMITINFO']._serialized_start=8325
  _globals['_EXTENDEDCOMMITINFO']._serialized_end=8431
  _globals['_EVENT']._serialized_start=8433
  _globals['_EVENT']._serialized_end=8556
  _globals['_EVENTATTRIBUTE']._serialized_start=8558
  _globals['_EVENTATTRIBUTE']._serialized_end=8636
  _globals['_EXECTXRESULT']._serialized_start=8639
  _globals['_EXECTXRESULT']._serialized_end=8896
  _globals['_TXRESULT']._serialized_start=8899
  _globals['_TXRESULT']._serialized_end=9033
  _globals['_VALIDATOR']._serialized_start=9035
  _globals['_VALIDATOR']._serialized_end=9094
  _globals['_VALIDATORUPDATE']._serialized_start=9096
  _globals['_VALIDATORUPDATE']._serialized_end=9211
  _globals['_VOTEINFO']._serialized_start=9214
  _globals['_VOTEINFO']._serialized_end=9363
  _globals['_EXTENDEDVOTEINFO']._serialized_start=9366
  _globals['_EXTENDEDVOTEINFO']._serialized_end=9723
  _globals['_MISBEHAVIOR']._serialized_start=9726
  _globals['_MISBEHAVIOR']._serialized_end=9987
  _globals['_SNAPSHOT']._serialized_start=9990
  _globals['_SNAPSHOT']._serialized_end=10120
# @@protoc_insertion_point(module_scope)
