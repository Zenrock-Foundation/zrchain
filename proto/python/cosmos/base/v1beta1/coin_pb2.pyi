from gogoproto import gogo_pb2 as _gogo_pb2
from cosmos_proto import cosmos_pb2 as _cosmos_pb2
from amino import amino_pb2 as _amino_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Coin(_message.Message):
    __slots__ = ("denom", "amount")
    DENOM_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    denom: str
    amount: str
    def __init__(self, denom: _Optional[str] = ..., amount: _Optional[str] = ...) -> None: ...

class DecCoin(_message.Message):
    __slots__ = ("denom", "amount")
    DENOM_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    denom: str
    amount: str
    def __init__(self, denom: _Optional[str] = ..., amount: _Optional[str] = ...) -> None: ...
