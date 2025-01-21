# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/treasury/mpcsign.proto
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
    'zrchain/treasury/mpcsign.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from zrchain.treasury import key_pb2 as zrchain_dot_treasury_dot_key__pb2
from zrchain.treasury import wallet_pb2 as zrchain_dot_treasury_dot_wallet__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ezrchain/treasury/mpcsign.proto\x12\x10zrchain.treasury\x1a\x19google/protobuf/any.proto\x1a\x1azrchain/treasury/key.proto\x1a\x1dzrchain/treasury/wallet.proto\"\xe1\x04\n\x0bSignRequest\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x18\n\x07\x63reator\x18\x02 \x01(\tR\x07\x63reator\x12\x15\n\x06key_id\x18\x03 \x01(\x04R\x05keyId\x12\x34\n\x08key_type\x18\x04 \x01(\x0e\x32\x19.zrchain.treasury.KeyTypeR\x07keyType\x12(\n\x10\x64\x61ta_for_signing\x18\x05 \x03(\x0cR\x0e\x64\x61taForSigning\x12;\n\x06status\x18\x06 \x01(\x0e\x32#.zrchain.treasury.SignRequestStatusR\x06status\x12\x43\n\x0bsigned_data\x18\x07 \x03(\x0b\x32\".zrchain.treasury.SignedDataWithIDR\nsignedData\x12\x38\n\x18keyring_party_signatures\x18\x08 \x03(\x0cR\x16keyringPartySignatures\x12#\n\rreject_reason\x18\t \x01(\tR\x0crejectReason\x12\x30\n\x08metadata\x18\n \x01(\x0b\x32\x14.google.protobuf.AnyR\x08metadata\x12\"\n\rparent_req_id\x18\x0b \x01(\x04R\x0bparentReqId\x12\"\n\rchild_req_ids\x18\x0c \x03(\x04R\x0b\x63hildReqIds\x12\x19\n\x08\x63\x61\x63he_id\x18\r \x01(\x0cR\x07\x63\x61\x63heId\x12\x17\n\x07key_ids\x18\x0e \x03(\x04R\x06keyIds\x12\x10\n\x03\x62tl\x18\x0f \x01(\x04R\x03\x62tl\x12\x10\n\x03\x66\x65\x65\x18\x10 \x01(\x04R\x03\x66\x65\x65\"[\n\x10SignedDataWithID\x12&\n\x0fsign_request_id\x18\x01 \x01(\x04R\rsignRequestId\x12\x1f\n\x0bsigned_data\x18\x02 \x01(\x0cR\nsignedData\"\x96\x02\n\x16SignTransactionRequest\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x18\n\x07\x63reator\x18\x02 \x01(\tR\x07\x63reator\x12\x15\n\x06key_id\x18\x03 \x01(\x04R\x05keyId\x12=\n\x0bwallet_type\x18\x04 \x01(\x0e\x32\x1c.zrchain.treasury.WalletTypeR\nwalletType\x12\x31\n\x14unsigned_transaction\x18\x05 \x01(\x0cR\x13unsignedTransaction\x12&\n\x0fsign_request_id\x18\x06 \x01(\x04R\rsignRequestId\x12!\n\x0cno_broadcast\x18\x07 \x01(\x08R\x0bnoBroadcast\"\x8e\x04\n\x0fSignReqResponse\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x18\n\x07\x63reator\x18\x02 \x01(\tR\x07\x63reator\x12\x17\n\x07key_ids\x18\x03 \x03(\x04R\x06keyIds\x12\x19\n\x08key_type\x18\x04 \x01(\tR\x07keyType\x12(\n\x10\x64\x61ta_for_signing\x18\x05 \x03(\x0cR\x0e\x64\x61taForSigning\x12\x16\n\x06status\x18\x06 \x01(\tR\x06status\x12\x43\n\x0bsigned_data\x18\x07 \x03(\x0b\x32\".zrchain.treasury.SignedDataWithIDR\nsignedData\x12\x38\n\x18keyring_party_signatures\x18\x08 \x03(\x0cR\x16keyringPartySignatures\x12#\n\rreject_reason\x18\t \x01(\tR\x0crejectReason\x12\x30\n\x08metadata\x18\n \x01(\x0b\x32\x14.google.protobuf.AnyR\x08metadata\x12\"\n\rparent_req_id\x18\x0b \x01(\x04R\x0bparentReqId\x12\"\n\rchild_req_ids\x18\x0c \x03(\x04R\x0b\x63hildReqIds\x12\x19\n\x08\x63\x61\x63he_id\x18\r \x01(\x0cR\x07\x63\x61\x63heId\x12\x10\n\x03\x62tl\x18\x0e \x01(\x04R\x03\x62tl\x12\x10\n\x03\x66\x65\x65\x18\x0f \x01(\x04R\x03\x66\x65\x65\"\x94\x02\n\x11SignTxReqResponse\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x18\n\x07\x63reator\x18\x02 \x01(\tR\x07\x63reator\x12\x15\n\x06key_id\x18\x03 \x01(\x04R\x05keyId\x12\x1f\n\x0bwallet_type\x18\x04 \x01(\tR\nwalletType\x12\x31\n\x14unsigned_transaction\x18\x05 \x01(\x0cR\x13unsignedTransaction\x12&\n\x0fsign_request_id\x18\x06 \x01(\x04R\rsignRequestId\x12!\n\x0cno_broadcast\x18\x07 \x01(\x08R\x0bnoBroadcast\x12\x1f\n\x0bmpc_timeout\x18\x08 \x01(\x04R\nmpcTimeout\"\xe8\x02\n\x15ICATransactionRequest\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x18\n\x07\x63reator\x18\x02 \x01(\tR\x07\x63reator\x12\x15\n\x06key_id\x18\x03 \x01(\x04R\x05keyId\x12\x34\n\x08key_type\x18\x04 \x01(\x0e\x32\x19.zrchain.treasury.KeyTypeR\x07keyType\x12\x1b\n\tinput_msg\x18\x05 \x01(\x0cR\x08inputMsg\x12;\n\x06status\x18\x06 \x01(\x0e\x32#.zrchain.treasury.SignRequestStatusR\x06status\x12\x1f\n\x0bsigned_data\x18\x07 \x03(\x0cR\nsignedData\x12\x38\n\x18keyring_party_signatures\x18\x08 \x03(\x0cR\x16keyringPartySignatures\x12#\n\rreject_reason\x18\t \x01(\tR\x0crejectReason*\xbf\x01\n\x11SignRequestStatus\x12#\n\x1fSIGN_REQUEST_STATUS_UNSPECIFIED\x10\x00\x12\x1f\n\x1bSIGN_REQUEST_STATUS_PENDING\x10\x01\x12\x1f\n\x1bSIGN_REQUEST_STATUS_PARTIAL\x10\x02\x12!\n\x1dSIGN_REQUEST_STATUS_FULFILLED\x10\x03\x12 \n\x1cSIGN_REQUEST_STATUS_REJECTED\x10\x04\x42;Z9github.com/Zenrock-Foundation/zrchain/v5/x/treasury/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.treasury.mpcsign_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z9github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types'
  _globals['_SIGNREQUESTSTATUS']._serialized_start=2296
  _globals['_SIGNREQUESTSTATUS']._serialized_end=2487
  _globals['_SIGNREQUEST']._serialized_start=139
  _globals['_SIGNREQUEST']._serialized_end=748
  _globals['_SIGNEDDATAWITHID']._serialized_start=750
  _globals['_SIGNEDDATAWITHID']._serialized_end=841
  _globals['_SIGNTRANSACTIONREQUEST']._serialized_start=844
  _globals['_SIGNTRANSACTIONREQUEST']._serialized_end=1122
  _globals['_SIGNREQRESPONSE']._serialized_start=1125
  _globals['_SIGNREQRESPONSE']._serialized_end=1651
  _globals['_SIGNTXREQRESPONSE']._serialized_start=1654
  _globals['_SIGNTXREQRESPONSE']._serialized_end=1930
  _globals['_ICATRANSACTIONREQUEST']._serialized_start=1933
  _globals['_ICATRANSACTIONREQUEST']._serialized_end=2293
# @@protoc_insertion_point(module_scope)
