# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ibc/applications/fee/v1/query.proto
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
    'ibc/applications/fee/v1/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from cosmos.base.v1beta1 import coin_pb2 as cosmos_dot_base_dot_v1beta1_dot_coin__pb2
from cosmos.base.query.v1beta1 import pagination_pb2 as cosmos_dot_base_dot_query_dot_v1beta1_dot_pagination__pb2
from ibc.applications.fee.v1 import fee_pb2 as ibc_dot_applications_dot_fee_dot_v1_dot_fee__pb2
from ibc.applications.fee.v1 import genesis_pb2 as ibc_dot_applications_dot_fee_dot_v1_dot_genesis__pb2
from ibc.core.channel.v1 import channel_pb2 as ibc_dot_core_dot_channel_dot_v1_dot_channel__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#ibc/applications/fee/v1/query.proto\x12\x17ibc.applications.fee.v1\x1a\x14gogoproto/gogo.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1e\x63osmos/base/v1beta1/coin.proto\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a!ibc/applications/fee/v1/fee.proto\x1a%ibc/applications/fee/v1/genesis.proto\x1a!ibc/core/channel/v1/channel.proto\"\x8c\x01\n\x1fQueryIncentivizedPacketsRequest\x12\x46\n\npagination\x18\x01 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\x12!\n\x0cquery_height\x18\x02 \x01(\x04R\x0bqueryHeight\"\xd3\x01\n QueryIncentivizedPacketsResponse\x12\x66\n\x14incentivized_packets\x18\x01 \x03(\x0b\x32-.ibc.applications.fee.v1.IdentifiedPacketFeesB\x04\xc8\xde\x1f\x00R\x13incentivizedPackets\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"\x85\x01\n\x1eQueryIncentivizedPacketRequest\x12@\n\tpacket_id\x18\x01 \x01(\x0b\x32\x1d.ibc.core.channel.v1.PacketIdB\x04\xc8\xde\x1f\x00R\x08packetId\x12!\n\x0cquery_height\x18\x02 \x01(\x04R\x0bqueryHeight\"\x87\x01\n\x1fQueryIncentivizedPacketResponse\x12\x64\n\x13incentivized_packet\x18\x01 \x01(\x0b\x32-.ibc.applications.fee.v1.IdentifiedPacketFeesB\x04\xc8\xde\x1f\x00R\x12incentivizedPacket\"\xce\x01\n)QueryIncentivizedPacketsForChannelRequest\x12\x46\n\npagination\x18\x01 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\x12\x17\n\x07port_id\x18\x02 \x01(\tR\x06portId\x12\x1d\n\nchannel_id\x18\x03 \x01(\tR\tchannelId\x12!\n\x0cquery_height\x18\x04 \x01(\x04R\x0bqueryHeight\"\xd7\x01\n*QueryIncentivizedPacketsForChannelResponse\x12`\n\x14incentivized_packets\x18\x01 \x03(\x0b\x32-.ibc.applications.fee.v1.IdentifiedPacketFeesR\x13incentivizedPackets\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"]\n\x19QueryTotalRecvFeesRequest\x12@\n\tpacket_id\x18\x01 \x01(\x0b\x32\x1d.ibc.core.channel.v1.PacketIdB\x04\xc8\xde\x1f\x00R\x08packetId\"\x86\x01\n\x1aQueryTotalRecvFeesResponse\x12h\n\trecv_fees\x18\x01 \x03(\x0b\x32\x19.cosmos.base.v1beta1.CoinB0\xc8\xde\x1f\x00\xaa\xdf\x1f(github.com/cosmos/cosmos-sdk/types.CoinsR\x08recvFees\"\\\n\x18QueryTotalAckFeesRequest\x12@\n\tpacket_id\x18\x01 \x01(\x0b\x32\x1d.ibc.core.channel.v1.PacketIdB\x04\xc8\xde\x1f\x00R\x08packetId\"\x83\x01\n\x19QueryTotalAckFeesResponse\x12\x66\n\x08\x61\x63k_fees\x18\x01 \x03(\x0b\x32\x19.cosmos.base.v1beta1.CoinB0\xc8\xde\x1f\x00\xaa\xdf\x1f(github.com/cosmos/cosmos-sdk/types.CoinsR\x07\x61\x63kFees\"`\n\x1cQueryTotalTimeoutFeesRequest\x12@\n\tpacket_id\x18\x01 \x01(\x0b\x32\x1d.ibc.core.channel.v1.PacketIdB\x04\xc8\xde\x1f\x00R\x08packetId\"\x8f\x01\n\x1dQueryTotalTimeoutFeesResponse\x12n\n\x0ctimeout_fees\x18\x01 \x03(\x0b\x32\x19.cosmos.base.v1beta1.CoinB0\xc8\xde\x1f\x00\xaa\xdf\x1f(github.com/cosmos/cosmos-sdk/types.CoinsR\x0btimeoutFees\"L\n\x11QueryPayeeRequest\x12\x1d\n\nchannel_id\x18\x01 \x01(\tR\tchannelId\x12\x18\n\x07relayer\x18\x02 \x01(\tR\x07relayer\"9\n\x12QueryPayeeResponse\x12#\n\rpayee_address\x18\x01 \x01(\tR\x0cpayeeAddress\"X\n\x1dQueryCounterpartyPayeeRequest\x12\x1d\n\nchannel_id\x18\x01 \x01(\tR\tchannelId\x12\x18\n\x07relayer\x18\x02 \x01(\tR\x07relayer\"O\n\x1eQueryCounterpartyPayeeResponse\x12-\n\x12\x63ounterparty_payee\x18\x01 \x01(\tR\x11\x63ounterpartyPayee\"\x8b\x01\n\x1eQueryFeeEnabledChannelsRequest\x12\x46\n\npagination\x18\x01 \x01(\x0b\x32&.cosmos.base.query.v1beta1.PageRequestR\npagination\x12!\n\x0cquery_height\x18\x02 \x01(\x04R\x0bqueryHeight\"\xce\x01\n\x1fQueryFeeEnabledChannelsResponse\x12\x62\n\x14\x66\x65\x65_enabled_channels\x18\x01 \x03(\x0b\x32*.ibc.applications.fee.v1.FeeEnabledChannelB\x04\xc8\xde\x1f\x00R\x12\x66\x65\x65\x45nabledChannels\x12G\n\npagination\x18\x02 \x01(\x0b\x32\'.cosmos.base.query.v1beta1.PageResponseR\npagination\"W\n\x1dQueryFeeEnabledChannelRequest\x12\x17\n\x07port_id\x18\x01 \x01(\tR\x06portId\x12\x1d\n\nchannel_id\x18\x02 \x01(\tR\tchannelId\"A\n\x1eQueryFeeEnabledChannelResponse\x12\x1f\n\x0b\x66\x65\x65_enabled\x18\x01 \x01(\x08R\nfeeEnabled2\xe6\x11\n\x05Query\x12\xb9\x01\n\x13IncentivizedPackets\x12\x38.ibc.applications.fee.v1.QueryIncentivizedPacketsRequest\x1a\x39.ibc.applications.fee.v1.QueryIncentivizedPacketsResponse\"-\x82\xd3\xe4\x93\x02\'\x12%/ibc/apps/fee/v1/incentivized_packets\x12\x8f\x02\n\x12IncentivizedPacket\x12\x37.ibc.applications.fee.v1.QueryIncentivizedPacketRequest\x1a\x38.ibc.applications.fee.v1.QueryIncentivizedPacketResponse\"\x85\x01\x82\xd3\xe4\x93\x02\x7f\x12}/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/incentivized_packet\x12\xfd\x01\n\x1dIncentivizedPacketsForChannel\x12\x42.ibc.applications.fee.v1.QueryIncentivizedPacketsForChannelRequest\x1a\x43.ibc.applications.fee.v1.QueryIncentivizedPacketsForChannelResponse\"S\x82\xd3\xe4\x93\x02M\x12K/ibc/apps/fee/v1/channels/{channel_id}/ports/{port_id}/incentivized_packets\x12\xfc\x01\n\rTotalRecvFees\x12\x32.ibc.applications.fee.v1.QueryTotalRecvFeesRequest\x1a\x33.ibc.applications.fee.v1.QueryTotalRecvFeesResponse\"\x81\x01\x82\xd3\xe4\x93\x02{\x12y/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/total_recv_fees\x12\xf8\x01\n\x0cTotalAckFees\x12\x31.ibc.applications.fee.v1.QueryTotalAckFeesRequest\x1a\x32.ibc.applications.fee.v1.QueryTotalAckFeesResponse\"\x80\x01\x82\xd3\xe4\x93\x02z\x12x/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/total_ack_fees\x12\x88\x02\n\x10TotalTimeoutFees\x12\x35.ibc.applications.fee.v1.QueryTotalTimeoutFeesRequest\x1a\x36.ibc.applications.fee.v1.QueryTotalTimeoutFeesResponse\"\x84\x01\x82\xd3\xe4\x93\x02~\x12|/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/total_timeout_fees\x12\xa9\x01\n\x05Payee\x12*.ibc.applications.fee.v1.QueryPayeeRequest\x1a+.ibc.applications.fee.v1.QueryPayeeResponse\"G\x82\xd3\xe4\x93\x02\x41\x12?/ibc/apps/fee/v1/channels/{channel_id}/relayers/{relayer}/payee\x12\xda\x01\n\x11\x43ounterpartyPayee\x12\x36.ibc.applications.fee.v1.QueryCounterpartyPayeeRequest\x1a\x37.ibc.applications.fee.v1.QueryCounterpartyPayeeResponse\"T\x82\xd3\xe4\x93\x02N\x12L/ibc/apps/fee/v1/channels/{channel_id}/relayers/{relayer}/counterparty_payee\x12\xad\x01\n\x12\x46\x65\x65\x45nabledChannels\x12\x37.ibc.applications.fee.v1.QueryFeeEnabledChannelsRequest\x1a\x38.ibc.applications.fee.v1.QueryFeeEnabledChannelsResponse\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/ibc/apps/fee/v1/fee_enabled\x12\xd0\x01\n\x11\x46\x65\x65\x45nabledChannel\x12\x36.ibc.applications.fee.v1.QueryFeeEnabledChannelRequest\x1a\x37.ibc.applications.fee.v1.QueryFeeEnabledChannelResponse\"J\x82\xd3\xe4\x93\x02\x44\x12\x42/ibc/apps/fee/v1/channels/{channel_id}/ports/{port_id}/fee_enabledB8Z6github.com/cosmos/ibc-go/v10/modules/apps/29-fee/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ibc.applications.fee.v1.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z6github.com/cosmos/ibc-go/v10/modules/apps/29-fee/types'
  _globals['_QUERYINCENTIVIZEDPACKETSRESPONSE'].fields_by_name['incentivized_packets']._loaded_options = None
  _globals['_QUERYINCENTIVIZEDPACKETSRESPONSE'].fields_by_name['incentivized_packets']._serialized_options = b'\310\336\037\000'
  _globals['_QUERYINCENTIVIZEDPACKETREQUEST'].fields_by_name['packet_id']._loaded_options = None
  _globals['_QUERYINCENTIVIZEDPACKETREQUEST'].fields_by_name['packet_id']._serialized_options = b'\310\336\037\000'
  _globals['_QUERYINCENTIVIZEDPACKETRESPONSE'].fields_by_name['incentivized_packet']._loaded_options = None
  _globals['_QUERYINCENTIVIZEDPACKETRESPONSE'].fields_by_name['incentivized_packet']._serialized_options = b'\310\336\037\000'
  _globals['_QUERYTOTALRECVFEESREQUEST'].fields_by_name['packet_id']._loaded_options = None
  _globals['_QUERYTOTALRECVFEESREQUEST'].fields_by_name['packet_id']._serialized_options = b'\310\336\037\000'
  _globals['_QUERYTOTALRECVFEESRESPONSE'].fields_by_name['recv_fees']._loaded_options = None
  _globals['_QUERYTOTALRECVFEESRESPONSE'].fields_by_name['recv_fees']._serialized_options = b'\310\336\037\000\252\337\037(github.com/cosmos/cosmos-sdk/types.Coins'
  _globals['_QUERYTOTALACKFEESREQUEST'].fields_by_name['packet_id']._loaded_options = None
  _globals['_QUERYTOTALACKFEESREQUEST'].fields_by_name['packet_id']._serialized_options = b'\310\336\037\000'
  _globals['_QUERYTOTALACKFEESRESPONSE'].fields_by_name['ack_fees']._loaded_options = None
  _globals['_QUERYTOTALACKFEESRESPONSE'].fields_by_name['ack_fees']._serialized_options = b'\310\336\037\000\252\337\037(github.com/cosmos/cosmos-sdk/types.Coins'
  _globals['_QUERYTOTALTIMEOUTFEESREQUEST'].fields_by_name['packet_id']._loaded_options = None
  _globals['_QUERYTOTALTIMEOUTFEESREQUEST'].fields_by_name['packet_id']._serialized_options = b'\310\336\037\000'
  _globals['_QUERYTOTALTIMEOUTFEESRESPONSE'].fields_by_name['timeout_fees']._loaded_options = None
  _globals['_QUERYTOTALTIMEOUTFEESRESPONSE'].fields_by_name['timeout_fees']._serialized_options = b'\310\336\037\000\252\337\037(github.com/cosmos/cosmos-sdk/types.Coins'
  _globals['_QUERYFEEENABLEDCHANNELSRESPONSE'].fields_by_name['fee_enabled_channels']._loaded_options = None
  _globals['_QUERYFEEENABLEDCHANNELSRESPONSE'].fields_by_name['fee_enabled_channels']._serialized_options = b'\310\336\037\000'
  _globals['_QUERY'].methods_by_name['IncentivizedPackets']._loaded_options = None
  _globals['_QUERY'].methods_by_name['IncentivizedPackets']._serialized_options = b'\202\323\344\223\002\'\022%/ibc/apps/fee/v1/incentivized_packets'
  _globals['_QUERY'].methods_by_name['IncentivizedPacket']._loaded_options = None
  _globals['_QUERY'].methods_by_name['IncentivizedPacket']._serialized_options = b'\202\323\344\223\002\177\022}/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/incentivized_packet'
  _globals['_QUERY'].methods_by_name['IncentivizedPacketsForChannel']._loaded_options = None
  _globals['_QUERY'].methods_by_name['IncentivizedPacketsForChannel']._serialized_options = b'\202\323\344\223\002M\022K/ibc/apps/fee/v1/channels/{channel_id}/ports/{port_id}/incentivized_packets'
  _globals['_QUERY'].methods_by_name['TotalRecvFees']._loaded_options = None
  _globals['_QUERY'].methods_by_name['TotalRecvFees']._serialized_options = b'\202\323\344\223\002{\022y/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/total_recv_fees'
  _globals['_QUERY'].methods_by_name['TotalAckFees']._loaded_options = None
  _globals['_QUERY'].methods_by_name['TotalAckFees']._serialized_options = b'\202\323\344\223\002z\022x/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/total_ack_fees'
  _globals['_QUERY'].methods_by_name['TotalTimeoutFees']._loaded_options = None
  _globals['_QUERY'].methods_by_name['TotalTimeoutFees']._serialized_options = b'\202\323\344\223\002~\022|/ibc/apps/fee/v1/channels/{packet_id.channel_id}/ports/{packet_id.port_id}/sequences/{packet_id.sequence}/total_timeout_fees'
  _globals['_QUERY'].methods_by_name['Payee']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Payee']._serialized_options = b'\202\323\344\223\002A\022?/ibc/apps/fee/v1/channels/{channel_id}/relayers/{relayer}/payee'
  _globals['_QUERY'].methods_by_name['CounterpartyPayee']._loaded_options = None
  _globals['_QUERY'].methods_by_name['CounterpartyPayee']._serialized_options = b'\202\323\344\223\002N\022L/ibc/apps/fee/v1/channels/{channel_id}/relayers/{relayer}/counterparty_payee'
  _globals['_QUERY'].methods_by_name['FeeEnabledChannels']._loaded_options = None
  _globals['_QUERY'].methods_by_name['FeeEnabledChannels']._serialized_options = b'\202\323\344\223\002\036\022\034/ibc/apps/fee/v1/fee_enabled'
  _globals['_QUERY'].methods_by_name['FeeEnabledChannel']._loaded_options = None
  _globals['_QUERY'].methods_by_name['FeeEnabledChannel']._serialized_options = b'\202\323\344\223\002D\022B/ibc/apps/fee/v1/channels/{channel_id}/ports/{port_id}/fee_enabled'
  _globals['_QUERYINCENTIVIZEDPACKETSREQUEST']._serialized_start=302
  _globals['_QUERYINCENTIVIZEDPACKETSREQUEST']._serialized_end=442
  _globals['_QUERYINCENTIVIZEDPACKETSRESPONSE']._serialized_start=445
  _globals['_QUERYINCENTIVIZEDPACKETSRESPONSE']._serialized_end=656
  _globals['_QUERYINCENTIVIZEDPACKETREQUEST']._serialized_start=659
  _globals['_QUERYINCENTIVIZEDPACKETREQUEST']._serialized_end=792
  _globals['_QUERYINCENTIVIZEDPACKETRESPONSE']._serialized_start=795
  _globals['_QUERYINCENTIVIZEDPACKETRESPONSE']._serialized_end=930
  _globals['_QUERYINCENTIVIZEDPACKETSFORCHANNELREQUEST']._serialized_start=933
  _globals['_QUERYINCENTIVIZEDPACKETSFORCHANNELREQUEST']._serialized_end=1139
  _globals['_QUERYINCENTIVIZEDPACKETSFORCHANNELRESPONSE']._serialized_start=1142
  _globals['_QUERYINCENTIVIZEDPACKETSFORCHANNELRESPONSE']._serialized_end=1357
  _globals['_QUERYTOTALRECVFEESREQUEST']._serialized_start=1359
  _globals['_QUERYTOTALRECVFEESREQUEST']._serialized_end=1452
  _globals['_QUERYTOTALRECVFEESRESPONSE']._serialized_start=1455
  _globals['_QUERYTOTALRECVFEESRESPONSE']._serialized_end=1589
  _globals['_QUERYTOTALACKFEESREQUEST']._serialized_start=1591
  _globals['_QUERYTOTALACKFEESREQUEST']._serialized_end=1683
  _globals['_QUERYTOTALACKFEESRESPONSE']._serialized_start=1686
  _globals['_QUERYTOTALACKFEESRESPONSE']._serialized_end=1817
  _globals['_QUERYTOTALTIMEOUTFEESREQUEST']._serialized_start=1819
  _globals['_QUERYTOTALTIMEOUTFEESREQUEST']._serialized_end=1915
  _globals['_QUERYTOTALTIMEOUTFEESRESPONSE']._serialized_start=1918
  _globals['_QUERYTOTALTIMEOUTFEESRESPONSE']._serialized_end=2061
  _globals['_QUERYPAYEEREQUEST']._serialized_start=2063
  _globals['_QUERYPAYEEREQUEST']._serialized_end=2139
  _globals['_QUERYPAYEERESPONSE']._serialized_start=2141
  _globals['_QUERYPAYEERESPONSE']._serialized_end=2198
  _globals['_QUERYCOUNTERPARTYPAYEEREQUEST']._serialized_start=2200
  _globals['_QUERYCOUNTERPARTYPAYEEREQUEST']._serialized_end=2288
  _globals['_QUERYCOUNTERPARTYPAYEERESPONSE']._serialized_start=2290
  _globals['_QUERYCOUNTERPARTYPAYEERESPONSE']._serialized_end=2369
  _globals['_QUERYFEEENABLEDCHANNELSREQUEST']._serialized_start=2372
  _globals['_QUERYFEEENABLEDCHANNELSREQUEST']._serialized_end=2511
  _globals['_QUERYFEEENABLEDCHANNELSRESPONSE']._serialized_start=2514
  _globals['_QUERYFEEENABLEDCHANNELSRESPONSE']._serialized_end=2720
  _globals['_QUERYFEEENABLEDCHANNELREQUEST']._serialized_start=2722
  _globals['_QUERYFEEENABLEDCHANNELREQUEST']._serialized_end=2809
  _globals['_QUERYFEEENABLEDCHANNELRESPONSE']._serialized_start=2811
  _globals['_QUERYFEEENABLEDCHANNELRESPONSE']._serialized_end=2876
  _globals['_QUERY']._serialized_start=2879
  _globals['_QUERY']._serialized_end=5157
# @@protoc_insertion_point(module_scope)
