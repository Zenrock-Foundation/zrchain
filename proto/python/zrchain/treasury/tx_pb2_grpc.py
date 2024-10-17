# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from zrchain.treasury import tx_pb2 as zrchain_dot_treasury_dot_tx__pb2


class MsgStub(object):
    """Msg defines the Msg service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UpdateParams = channel.unary_unary(
                '/zrchain.treasury.Msg/UpdateParams',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateParams.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateParamsResponse.FromString,
                )
        self.NewKeyRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/NewKeyRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewKeyRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewKeyRequestResponse.FromString,
                )
        self.FulfilKeyRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/FulfilKeyRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilKeyRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilKeyRequestResponse.FromString,
                )
        self.NewSignatureRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/NewSignatureRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignatureRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignatureRequestResponse.FromString,
                )
        self.FulfilSignatureRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/FulfilSignatureRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilSignatureRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilSignatureRequestResponse.FromString,
                )
        self.NewSignTransactionRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/NewSignTransactionRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignTransactionRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignTransactionRequestResponse.FromString,
                )
        self.TransferFromKeyring = channel.unary_unary(
                '/zrchain.treasury.Msg/TransferFromKeyring',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgTransferFromKeyring.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgTransferFromKeyringResponse.FromString,
                )
        self.NewICATransactionRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/NewICATransactionRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewICATransactionRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewICATransactionRequestResponse.FromString,
                )
        self.FulfilICATransactionRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/FulfilICATransactionRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilICATransactionRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilICATransactionRequestResponse.FromString,
                )
        self.NewZrSignSignatureRequest = channel.unary_unary(
                '/zrchain.treasury.Msg/NewZrSignSignatureRequest',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewZrSignSignatureRequest.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewZrSignSignatureRequestResponse.FromString,
                )
        self.UpdateKeyPolicy = channel.unary_unary(
                '/zrchain.treasury.Msg/UpdateKeyPolicy',
                request_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateKeyPolicy.SerializeToString,
                response_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateKeyPolicyResponse.FromString,
                )


class MsgServicer(object):
    """Msg defines the Msg service.
    """

    def UpdateParams(self, request, context):
        """UpdateParams defines the operation for updating the module
        parameters. The authority defaults to the x/gov module account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NewKeyRequest(self, request, context):
        """NewKeyRequest defines an operation for creating a key request.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def FulfilKeyRequest(self, request, context):
        """FulfilKeyRequest defines an operation for responding to a key request
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NewSignatureRequest(self, request, context):
        """NewSignatureRequest defines an operation for creating a signature request
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def FulfilSignatureRequest(self, request, context):
        """FulfilSignatureRequest defines an operation for returning a signature
        response to a request
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NewSignTransactionRequest(self, request, context):
        """NewSignTransactionRequest defines an operation for creating a signature for
        transaction request
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TransferFromKeyring(self, request, context):
        """TransferFromKeyring defines an operation for transferring tokens from a
        keyring
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NewICATransactionRequest(self, request, context):
        """NewICATransactionRequest defines an operation for creating an interchain
        account transaction request
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def FulfilICATransactionRequest(self, request, context):
        """FulfilICATransactionRequest defines an operation for responding to an
        interchain account transaction request
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NewZrSignSignatureRequest(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateKeyPolicy(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MsgServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UpdateParams': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateParams,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateParams.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateParamsResponse.SerializeToString,
            ),
            'NewKeyRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.NewKeyRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewKeyRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewKeyRequestResponse.SerializeToString,
            ),
            'FulfilKeyRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.FulfilKeyRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilKeyRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilKeyRequestResponse.SerializeToString,
            ),
            'NewSignatureRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.NewSignatureRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignatureRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignatureRequestResponse.SerializeToString,
            ),
            'FulfilSignatureRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.FulfilSignatureRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilSignatureRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilSignatureRequestResponse.SerializeToString,
            ),
            'NewSignTransactionRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.NewSignTransactionRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignTransactionRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewSignTransactionRequestResponse.SerializeToString,
            ),
            'TransferFromKeyring': grpc.unary_unary_rpc_method_handler(
                    servicer.TransferFromKeyring,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgTransferFromKeyring.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgTransferFromKeyringResponse.SerializeToString,
            ),
            'NewICATransactionRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.NewICATransactionRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewICATransactionRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewICATransactionRequestResponse.SerializeToString,
            ),
            'FulfilICATransactionRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.FulfilICATransactionRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilICATransactionRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgFulfilICATransactionRequestResponse.SerializeToString,
            ),
            'NewZrSignSignatureRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.NewZrSignSignatureRequest,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewZrSignSignatureRequest.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgNewZrSignSignatureRequestResponse.SerializeToString,
            ),
            'UpdateKeyPolicy': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateKeyPolicy,
                    request_deserializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateKeyPolicy.FromString,
                    response_serializer=zrchain_dot_treasury_dot_tx__pb2.MsgUpdateKeyPolicyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'zrchain.treasury.Msg', rpc_method_handlers)
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
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/UpdateParams',
            zrchain_dot_treasury_dot_tx__pb2.MsgUpdateParams.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgUpdateParamsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NewKeyRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/NewKeyRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgNewKeyRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgNewKeyRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def FulfilKeyRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/FulfilKeyRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgFulfilKeyRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgFulfilKeyRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NewSignatureRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/NewSignatureRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgNewSignatureRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgNewSignatureRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def FulfilSignatureRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/FulfilSignatureRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgFulfilSignatureRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgFulfilSignatureRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NewSignTransactionRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/NewSignTransactionRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgNewSignTransactionRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgNewSignTransactionRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TransferFromKeyring(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/TransferFromKeyring',
            zrchain_dot_treasury_dot_tx__pb2.MsgTransferFromKeyring.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgTransferFromKeyringResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NewICATransactionRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/NewICATransactionRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgNewICATransactionRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgNewICATransactionRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def FulfilICATransactionRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/FulfilICATransactionRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgFulfilICATransactionRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgFulfilICATransactionRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NewZrSignSignatureRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/NewZrSignSignatureRequest',
            zrchain_dot_treasury_dot_tx__pb2.MsgNewZrSignSignatureRequest.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgNewZrSignSignatureRequestResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateKeyPolicy(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.treasury.Msg/UpdateKeyPolicy',
            zrchain_dot_treasury_dot_tx__pb2.MsgUpdateKeyPolicy.SerializeToString,
            zrchain_dot_treasury_dot_tx__pb2.MsgUpdateKeyPolicyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
