# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from cosmos.protocolpool.v1 import query_pb2 as cosmos_dot_protocolpool_dot_v1_dot_query__pb2


class QueryStub(object):
    """Query defines the gRPC querier service for community pool module.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CommunityPool = channel.unary_unary(
                '/cosmos.protocolpool.v1.Query/CommunityPool',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryCommunityPoolRequest.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryCommunityPoolResponse.FromString,
                )
        self.UnclaimedBudget = channel.unary_unary(
                '/cosmos.protocolpool.v1.Query/UnclaimedBudget',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryUnclaimedBudgetRequest.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryUnclaimedBudgetResponse.FromString,
                )


class QueryServicer(object):
    """Query defines the gRPC querier service for community pool module.
    """

    def CommunityPool(self, request, context):
        """CommunityPool queries the community pool coins.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UnclaimedBudget(self, request, context):
        """UnclaimedBudget queries the remaining budget left to be claimed and it gives overall budget allocation view.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_QueryServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CommunityPool': grpc.unary_unary_rpc_method_handler(
                    servicer.CommunityPool,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryCommunityPoolRequest.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryCommunityPoolResponse.SerializeToString,
            ),
            'UnclaimedBudget': grpc.unary_unary_rpc_method_handler(
                    servicer.UnclaimedBudget,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryUnclaimedBudgetRequest.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryUnclaimedBudgetResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'cosmos.protocolpool.v1.Query', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Query(object):
    """Query defines the gRPC querier service for community pool module.
    """

    @staticmethod
    def CommunityPool(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Query/CommunityPool',
            cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryCommunityPoolRequest.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryCommunityPoolResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UnclaimedBudget(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Query/UnclaimedBudget',
            cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryUnclaimedBudgetRequest.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_query__pb2.QueryUnclaimedBudgetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
