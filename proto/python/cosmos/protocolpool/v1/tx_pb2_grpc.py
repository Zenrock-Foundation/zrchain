# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from cosmos.protocolpool.v1 import tx_pb2 as cosmos_dot_protocolpool_dot_v1_dot_tx__pb2


class MsgStub(object):
    """Msg defines the pool Msg service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FundCommunityPool = channel.unary_unary(
                '/cosmos.protocolpool.v1.Msg/FundCommunityPool',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgFundCommunityPool.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgFundCommunityPoolResponse.FromString,
                )
        self.CommunityPoolSpend = channel.unary_unary(
                '/cosmos.protocolpool.v1.Msg/CommunityPoolSpend',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCommunityPoolSpend.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCommunityPoolSpendResponse.FromString,
                )
        self.SubmitBudgetProposal = channel.unary_unary(
                '/cosmos.protocolpool.v1.Msg/SubmitBudgetProposal',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgSubmitBudgetProposal.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgSubmitBudgetProposalResponse.FromString,
                )
        self.ClaimBudget = channel.unary_unary(
                '/cosmos.protocolpool.v1.Msg/ClaimBudget',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgClaimBudget.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgClaimBudgetResponse.FromString,
                )
        self.CreateContinuousFund = channel.unary_unary(
                '/cosmos.protocolpool.v1.Msg/CreateContinuousFund',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCreateContinuousFund.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCreateContinuousFundResponse.FromString,
                )
        self.WithdrawContinuousFund = channel.unary_unary(
                '/cosmos.protocolpool.v1.Msg/WithdrawContinuousFund',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgWithdrawContinuousFund.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgWithdrawContinuousFundResponse.FromString,
                )
        self.CancelContinuousFund = channel.unary_unary(
                '/cosmos.protocolpool.v1.Msg/CancelContinuousFund',
                request_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCancelContinuousFund.SerializeToString,
                response_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCancelContinuousFundResponse.FromString,
                )


class MsgServicer(object):
    """Msg defines the pool Msg service.
    """

    def FundCommunityPool(self, request, context):
        """FundCommunityPool defines a method to allow an account to directly
        fund the community pool.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CommunityPoolSpend(self, request, context):
        """CommunityPoolSpend defines a governance operation for sending tokens from
        the community pool in the x/protocolpool module to another account, which
        could be the governance module itself. The authority is defined in the
        keeper.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubmitBudgetProposal(self, request, context):
        """SubmitBudgetProposal defines a method to set a budget proposal.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ClaimBudget(self, request, context):
        """ClaimBudget defines a method to claim the distributed budget.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateContinuousFund(self, request, context):
        """CreateContinuousFund defines a method to add funds continuously.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def WithdrawContinuousFund(self, request, context):
        """WithdrawContinuousFund defines a method to withdraw continuous fund allocated.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CancelContinuousFund(self, request, context):
        """CancelContinuousFund defines a method for cancelling continuous fund.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MsgServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FundCommunityPool': grpc.unary_unary_rpc_method_handler(
                    servicer.FundCommunityPool,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgFundCommunityPool.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgFundCommunityPoolResponse.SerializeToString,
            ),
            'CommunityPoolSpend': grpc.unary_unary_rpc_method_handler(
                    servicer.CommunityPoolSpend,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCommunityPoolSpend.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCommunityPoolSpendResponse.SerializeToString,
            ),
            'SubmitBudgetProposal': grpc.unary_unary_rpc_method_handler(
                    servicer.SubmitBudgetProposal,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgSubmitBudgetProposal.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgSubmitBudgetProposalResponse.SerializeToString,
            ),
            'ClaimBudget': grpc.unary_unary_rpc_method_handler(
                    servicer.ClaimBudget,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgClaimBudget.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgClaimBudgetResponse.SerializeToString,
            ),
            'CreateContinuousFund': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateContinuousFund,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCreateContinuousFund.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCreateContinuousFundResponse.SerializeToString,
            ),
            'WithdrawContinuousFund': grpc.unary_unary_rpc_method_handler(
                    servicer.WithdrawContinuousFund,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgWithdrawContinuousFund.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgWithdrawContinuousFundResponse.SerializeToString,
            ),
            'CancelContinuousFund': grpc.unary_unary_rpc_method_handler(
                    servicer.CancelContinuousFund,
                    request_deserializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCancelContinuousFund.FromString,
                    response_serializer=cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCancelContinuousFundResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'cosmos.protocolpool.v1.Msg', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Msg(object):
    """Msg defines the pool Msg service.
    """

    @staticmethod
    def FundCommunityPool(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Msg/FundCommunityPool',
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgFundCommunityPool.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgFundCommunityPoolResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CommunityPoolSpend(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Msg/CommunityPoolSpend',
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCommunityPoolSpend.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCommunityPoolSpendResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SubmitBudgetProposal(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Msg/SubmitBudgetProposal',
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgSubmitBudgetProposal.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgSubmitBudgetProposalResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ClaimBudget(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Msg/ClaimBudget',
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgClaimBudget.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgClaimBudgetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateContinuousFund(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Msg/CreateContinuousFund',
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCreateContinuousFund.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCreateContinuousFundResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def WithdrawContinuousFund(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Msg/WithdrawContinuousFund',
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgWithdrawContinuousFund.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgWithdrawContinuousFundResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CancelContinuousFund(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/cosmos.protocolpool.v1.Msg/CancelContinuousFund',
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCancelContinuousFund.SerializeToString,
            cosmos_dot_protocolpool_dot_v1_dot_tx__pb2.MsgCancelContinuousFundResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
