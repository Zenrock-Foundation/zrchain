# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: cosmos/nft/v1beta1/query.proto
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
    'cosmos/nft/v1beta1/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from cosmos.base.query.v1beta1 import pagination_pb2 as cosmos_dot_base_dot_query_dot_v1beta1_dot_pagination__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from cosmos.nft.v1beta1 import nft_pb2 as cosmos_dot_nft_dot_v1beta1_dot_nft__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1e\x63osmos/nft/v1beta1/query.proto\x12\x12\x63osmos.nft.v1beta1\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1c\x63osmos/nft/v1beta1/nft.proto\"F\n\x13QueryBalanceRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\x12\x14\n\x05owner\x18\x02 \x01(\tR\x05owner\"S\n QueryBalanceByQueryStringRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\x12\x14\n\x05owner\x18\x02 \x01(\tR\x05owner\".\n\x14QueryBalanceResponse\x12\x16\n\x06\x61mount\x18\x01 \x01(\x04R\x06\x61mount\";\n!QueryBalanceByQueryStringResponse\x12\x16\n\x06\x61mount\x18\x01 \x01(\x04R\x06\x61mount\">\n\x11QueryOwnerRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\"K\n\x1eQueryOwnerByQueryStringRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\"*\n\x12QueryOwnerResponse\x12\x14\n\x05owner\x18\x01 \x01(\tR\x05owner\"7\n\x1fQueryOwnerByQueryStringResponse\x12\x14\n\x05owner\x18\x01 \x01(\tR\x05owner\"/\n\x12QuerySupplyRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\"<\n\x1fQuerySupplyByQueryStringRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\"-\n\x13QuerySupplyResponse\x12\x16\n\x06\x61mount\x18\x01 \x01(\x04R\x06\x61mount\":\n QuerySupplyByQueryStringResponse\x12\x16\n\x06\x61mount\x18\x01 \x01(\x04R\x06\x61mount\"\x8b\x01\n\x10QueryNFTsRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\x12\x14\n\x05owner\x18\x02 \x01(\tR\x05owner\x12\x46\n\npagination\x18\x03 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\x89\x01\n\x11QueryNFTsResponse\x12+\n\x04nfts\x18\x01 \x03(\x0b\x32\x17.cosmos.nft.v1beta1.NFTR\x04nfts\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"<\n\x0fQueryNFTRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\"I\n\x1cQueryNFTByQueryStringRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\"=\n\x10QueryNFTResponse\x12)\n\x03nft\x18\x01 \x01(\x0b\x32\x17.cosmos.nft.v1beta1.NFTR\x03nft\"J\n\x1dQueryNFTByQueryStringResponse\x12)\n\x03nft\x18\x01 \x01(\x0b\x32\x17.cosmos.nft.v1beta1.NFTR\x03nft\".\n\x11QueryClassRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\";\n\x1eQueryClassByQueryStringRequest\x12\x19\n\x08\x63lass_id\x18\x01 \x01(\tR\x07\x63lassId\"E\n\x12QueryClassResponse\x12/\n\x05\x63lass\x18\x01 \x01(\x0b\x32\x19.cosmos.nft.v1beta1.ClassR\x05\x63lass\"R\n\x1fQueryClassByQueryStringResponse\x12/\n\x05\x63lass\x18\x01 \x01(\x0b\x32\x19.cosmos.nft.v1beta1.ClassR\x05\x63lass\"]\n\x13QueryClassesRequest\x12\x46\n\npagination\x18\x01 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\"\x94\x01\n\x14QueryClassesResponse\x12\x33\n\x07\x63lasses\x18\x01 \x03(\x0b\x32\x19.cosmos.nft.v1beta1.ClassR\x07\x63lasses\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination2\xf1\r\n\x05Query\x12\x94\x01\n\x07\x42\x61lance\x12\'.cosmos.nft.v1beta1.QueryBalanceRequest\x1a(.cosmos.nft.v1beta1.QueryBalanceResponse\"6\x82\xd3\xe4\x93\x02\x30\x12./cosmos/nft/v1beta1/balance/{owner}/{class_id}\x12\xa8\x01\n\x14\x42\x61lanceByQueryString\x12\x34.cosmos.nft.v1beta1.QueryBalanceByQueryStringRequest\x1a\x35.cosmos.nft.v1beta1.QueryBalanceByQueryStringResponse\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/cosmos/nft/v1beta1/balance\x12\x89\x01\n\x05Owner\x12%.cosmos.nft.v1beta1.QueryOwnerRequest\x1a&.cosmos.nft.v1beta1.QueryOwnerResponse\"1\x82\xd3\xe4\x93\x02+\x12)/cosmos/nft/v1beta1/owner/{class_id}/{id}\x12\xa0\x01\n\x12OwnerByQueryString\x12\x32.cosmos.nft.v1beta1.QueryOwnerByQueryStringRequest\x1a\x33.cosmos.nft.v1beta1.QueryOwnerByQueryStringResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/cosmos/nft/v1beta1/owner\x12\x88\x01\n\x06Supply\x12&.cosmos.nft.v1beta1.QuerySupplyRequest\x1a\'.cosmos.nft.v1beta1.QuerySupplyResponse\"-\x82\xd3\xe4\x93\x02\'\x12%/cosmos/nft/v1beta1/supply/{class_id}\x12\xa4\x01\n\x13SupplyByQueryString\x12\x33.cosmos.nft.v1beta1.QuerySupplyByQueryStringRequest\x1a\x34.cosmos.nft.v1beta1.QuerySupplyByQueryStringResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/cosmos/nft/v1beta1/supply\x12u\n\x04NFTs\x12$.cosmos.nft.v1beta1.QueryNFTsRequest\x1a%.cosmos.nft.v1beta1.QueryNFTsResponse\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/cosmos/nft/v1beta1/nfts\x12\x82\x01\n\x03NFT\x12#.cosmos.nft.v1beta1.QueryNFTRequest\x1a$.cosmos.nft.v1beta1.QueryNFTResponse\"0\x82\xd3\xe4\x93\x02*\x12(/cosmos/nft/v1beta1/nfts/{class_id}/{id}\x12\x98\x01\n\x10NFTByQueryString\x12\x30.cosmos.nft.v1beta1.QueryNFTByQueryStringRequest\x1a\x31.cosmos.nft.v1beta1.QueryNFTByQueryStringResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/cosmos/nft/v1beta1/nft\x12\x86\x01\n\x05\x43lass\x12%.cosmos.nft.v1beta1.QueryClassRequest\x1a&.cosmos.nft.v1beta1.QueryClassResponse\".\x82\xd3\xe4\x93\x02(\x12&/cosmos/nft/v1beta1/classes/{class_id}\x12\xa0\x01\n\x12\x43lassByQueryString\x12\x32.cosmos.nft.v1beta1.QueryClassByQueryStringRequest\x1a\x33.cosmos.nft.v1beta1.QueryClassByQueryStringResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/cosmos/nft/v1beta1/class\x12\x81\x01\n\x07\x43lasses\x12\'.cosmos.nft.v1beta1.QueryClassesRequest\x1a(.cosmos.nft.v1beta1.QueryClassesResponse\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/cosmos/nft/v1beta1/classesB\x14Z\x12\x63osmossdk.io/x/nftb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cosmos.nft.v1beta1.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\022cosmossdk.io/x/nft'
  _globals['_QUERY'].methods_by_name['Balance']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Balance']._serialized_options = b'\202\323\344\223\0020\022./cosmos/nft/v1beta1/balance/{owner}/{class_id}'
  _globals['_QUERY'].methods_by_name['BalanceByQueryString']._loaded_options = None
  _globals['_QUERY'].methods_by_name['BalanceByQueryString']._serialized_options = b'\202\323\344\223\002\035\022\033/cosmos/nft/v1beta1/balance'
  _globals['_QUERY'].methods_by_name['Owner']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Owner']._serialized_options = b'\202\323\344\223\002+\022)/cosmos/nft/v1beta1/owner/{class_id}/{id}'
  _globals['_QUERY'].methods_by_name['OwnerByQueryString']._loaded_options = None
  _globals['_QUERY'].methods_by_name['OwnerByQueryString']._serialized_options = b'\202\323\344\223\002\033\022\031/cosmos/nft/v1beta1/owner'
  _globals['_QUERY'].methods_by_name['Supply']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Supply']._serialized_options = b'\202\323\344\223\002\'\022%/cosmos/nft/v1beta1/supply/{class_id}'
  _globals['_QUERY'].methods_by_name['SupplyByQueryString']._loaded_options = None
  _globals['_QUERY'].methods_by_name['SupplyByQueryString']._serialized_options = b'\202\323\344\223\002\034\022\032/cosmos/nft/v1beta1/supply'
  _globals['_QUERY'].methods_by_name['NFTs']._loaded_options = None
  _globals['_QUERY'].methods_by_name['NFTs']._serialized_options = b'\202\323\344\223\002\032\022\030/cosmos/nft/v1beta1/nfts'
  _globals['_QUERY'].methods_by_name['NFT']._loaded_options = None
  _globals['_QUERY'].methods_by_name['NFT']._serialized_options = b'\202\323\344\223\002*\022(/cosmos/nft/v1beta1/nfts/{class_id}/{id}'
  _globals['_QUERY'].methods_by_name['NFTByQueryString']._loaded_options = None
  _globals['_QUERY'].methods_by_name['NFTByQueryString']._serialized_options = b'\202\323\344\223\002\031\022\027/cosmos/nft/v1beta1/nft'
  _globals['_QUERY'].methods_by_name['Class']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Class']._serialized_options = b'\202\323\344\223\002(\022&/cosmos/nft/v1beta1/classes/{class_id}'
  _globals['_QUERY'].methods_by_name['ClassByQueryString']._loaded_options = None
  _globals['_QUERY'].methods_by_name['ClassByQueryString']._serialized_options = b'\202\323\344\223\002\033\022\031/cosmos/nft/v1beta1/class'
  _globals['_QUERY'].methods_by_name['Classes']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Classes']._serialized_options = b'\202\323\344\223\002\035\022\033/cosmos/nft/v1beta1/classes'
  _globals['_QUERYBALANCEREQUEST']._serialized_start=158
  _globals['_QUERYBALANCEREQUEST']._serialized_end=228
  _globals['_QUERYBALANCEBYQUERYSTRINGREQUEST']._serialized_start=230
  _globals['_QUERYBALANCEBYQUERYSTRINGREQUEST']._serialized_end=313
  _globals['_QUERYBALANCERESPONSE']._serialized_start=315
  _globals['_QUERYBALANCERESPONSE']._serialized_end=361
  _globals['_QUERYBALANCEBYQUERYSTRINGRESPONSE']._serialized_start=363
  _globals['_QUERYBALANCEBYQUERYSTRINGRESPONSE']._serialized_end=422
  _globals['_QUERYOWNERREQUEST']._serialized_start=424
  _globals['_QUERYOWNERREQUEST']._serialized_end=486
  _globals['_QUERYOWNERBYQUERYSTRINGREQUEST']._serialized_start=488
  _globals['_QUERYOWNERBYQUERYSTRINGREQUEST']._serialized_end=563
  _globals['_QUERYOWNERRESPONSE']._serialized_start=565
  _globals['_QUERYOWNERRESPONSE']._serialized_end=607
  _globals['_QUERYOWNERBYQUERYSTRINGRESPONSE']._serialized_start=609
  _globals['_QUERYOWNERBYQUERYSTRINGRESPONSE']._serialized_end=664
  _globals['_QUERYSUPPLYREQUEST']._serialized_start=666
  _globals['_QUERYSUPPLYREQUEST']._serialized_end=713
  _globals['_QUERYSUPPLYBYQUERYSTRINGREQUEST']._serialized_start=715
  _globals['_QUERYSUPPLYBYQUERYSTRINGREQUEST']._serialized_end=775
  _globals['_QUERYSUPPLYRESPONSE']._serialized_start=777
  _globals['_QUERYSUPPLYRESPONSE']._serialized_end=822
  _globals['_QUERYSUPPLYBYQUERYSTRINGRESPONSE']._serialized_start=824
  _globals['_QUERYSUPPLYBYQUERYSTRINGRESPONSE']._serialized_end=882
  _globals['_QUERYNFTSREQUEST']._serialized_start=885
  _globals['_QUERYNFTSREQUEST']._serialized_end=1024
  _globals['_QUERYNFTSRESPONSE']._serialized_start=1027
  _globals['_QUERYNFTSRESPONSE']._serialized_end=1164
  _globals['_QUERYNFTREQUEST']._serialized_start=1166
  _globals['_QUERYNFTREQUEST']._serialized_end=1226
  _globals['_QUERYNFTBYQUERYSTRINGREQUEST']._serialized_start=1228
  _globals['_QUERYNFTBYQUERYSTRINGREQUEST']._serialized_end=1301
  _globals['_QUERYNFTRESPONSE']._serialized_start=1303
  _globals['_QUERYNFTRESPONSE']._serialized_end=1364
  _globals['_QUERYNFTBYQUERYSTRINGRESPONSE']._serialized_start=1366
  _globals['_QUERYNFTBYQUERYSTRINGRESPONSE']._serialized_end=1440
  _globals['_QUERYCLASSREQUEST']._serialized_start=1442
  _globals['_QUERYCLASSREQUEST']._serialized_end=1488
  _globals['_QUERYCLASSBYQUERYSTRINGREQUEST']._serialized_start=1490
  _globals['_QUERYCLASSBYQUERYSTRINGREQUEST']._serialized_end=1549
  _globals['_QUERYCLASSRESPONSE']._serialized_start=1551
  _globals['_QUERYCLASSRESPONSE']._serialized_end=1620
  _globals['_QUERYCLASSBYQUERYSTRINGRESPONSE']._serialized_start=1622
  _globals['_QUERYCLASSBYQUERYSTRINGRESPONSE']._serialized_end=1704
  _globals['_QUERYCLASSESREQUEST']._serialized_start=1706
  _globals['_QUERYCLASSESREQUEST']._serialized_end=1799
  _globals['_QUERYCLASSESRESPONSE']._serialized_start=1802
  _globals['_QUERYCLASSESRESPONSE']._serialized_end=1950
  _globals['_QUERY']._serialized_start=1953
  _globals['_QUERY']._serialized_end=3730
# @@protoc_insertion_point(module_scope)