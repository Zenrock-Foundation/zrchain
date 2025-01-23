from amino import amino_pb2 as _amino_pb2
from gogoproto import gogo_pb2 as _gogo_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Params(_message.Message):
    __slots__ = ("keyring_creation_fee", "mpc_minimum_btl")
    KEYRING_CREATION_FEE_FIELD_NUMBER: _ClassVar[int]
    MPC_MINIMUM_BTL_FIELD_NUMBER: _ClassVar[int]
    keyring_creation_fee: int
    mpc_minimum_btl: int
    def __init__(self, keyring_creation_fee: _Optional[int] = ..., mpc_minimum_btl: _Optional[int] = ...) -> None: ...
