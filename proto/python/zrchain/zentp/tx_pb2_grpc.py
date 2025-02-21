# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from zrchain.zentp import tx_pb2 as zrchain_dot_zentp_dot_tx__pb2


class MsgStub(object):
    """Msg defines the Msg service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UpdateParams = channel.unary_unary(
                '/zrchain.zentp.Msg/UpdateParams',
                request_serializer=zrchain_dot_zentp_dot_tx__pb2.MsgUpdateParams.SerializeToString,
                response_deserializer=zrchain_dot_zentp_dot_tx__pb2.MsgUpdateParamsResponse.FromString,
                )
        self.BridgeRock = channel.unary_unary(
                '/zrchain.zentp.Msg/BridgeRock',
                request_serializer=zrchain_dot_zentp_dot_tx__pb2.MsgBridgeRock.SerializeToString,
                response_deserializer=zrchain_dot_zentp_dot_tx__pb2.MsgBridgeRockResponse.FromString,
                )
        self.Burn = channel.unary_unary(
                '/zrchain.zentp.Msg/Burn',
                request_serializer=zrchain_dot_zentp_dot_tx__pb2.MsgBurn.SerializeToString,
                response_deserializer=zrchain_dot_zentp_dot_tx__pb2.MsgBurnResponse.FromString,
                )


class MsgServicer(object):
    """Msg defines the Msg service.
    """

    def UpdateParams(self, request, context):
        """UpdateParams defines a (governance) operation for updating the module
        parameters. The authority defaults to the x/gov module account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BridgeRock(self, request, context):
        """MintRock defines an operation for creating a mint request of Rock
        on a destination chain
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Burn(self, request, context):
        """Burn defines an operation for burning Rock for a module account
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MsgServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UpdateParams': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateParams,
                    request_deserializer=zrchain_dot_zentp_dot_tx__pb2.MsgUpdateParams.FromString,
                    response_serializer=zrchain_dot_zentp_dot_tx__pb2.MsgUpdateParamsResponse.SerializeToString,
            ),
            'BridgeRock': grpc.unary_unary_rpc_method_handler(
                    servicer.BridgeRock,
                    request_deserializer=zrchain_dot_zentp_dot_tx__pb2.MsgBridgeRock.FromString,
                    response_serializer=zrchain_dot_zentp_dot_tx__pb2.MsgBridgeRockResponse.SerializeToString,
            ),
            'Burn': grpc.unary_unary_rpc_method_handler(
                    servicer.Burn,
                    request_deserializer=zrchain_dot_zentp_dot_tx__pb2.MsgBurn.FromString,
                    response_serializer=zrchain_dot_zentp_dot_tx__pb2.MsgBurnResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'zrchain.zentp.Msg', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Msg(object):
    """Msg defines the Msg service.
    """

    @staticmethod
    def UpdateParams(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zentp.Msg/UpdateParams',
            zrchain_dot_zentp_dot_tx__pb2.MsgUpdateParams.SerializeToString,
            zrchain_dot_zentp_dot_tx__pb2.MsgUpdateParamsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BridgeRock(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zentp.Msg/BridgeRock',
            zrchain_dot_zentp_dot_tx__pb2.MsgBridgeRock.SerializeToString,
            zrchain_dot_zentp_dot_tx__pb2.MsgBridgeRockResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Burn(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zentp.Msg/Burn',
            zrchain_dot_zentp_dot_tx__pb2.MsgBurn.SerializeToString,
            zrchain_dot_zentp_dot_tx__pb2.MsgBurnResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
