from gogoproto import gogo_pb2 as _gogo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Params(_message.Message):
    __slots__ = ("send_enabled", "receive_enabled")
    SEND_ENABLED_FIELD_NUMBER: _ClassVar[int]
    RECEIVE_ENABLED_FIELD_NUMBER: _ClassVar[int]
    send_enabled: bool
    receive_enabled: bool
    def __init__(self, send_enabled: bool = ..., receive_enabled: bool = ...) -> None: ...

class Forwarding(_message.Message):
    __slots__ = ("unwind", "hops")
    UNWIND_FIELD_NUMBER: _ClassVar[int]
    HOPS_FIELD_NUMBER: _ClassVar[int]
    unwind: bool
    hops: _containers.RepeatedCompositeFieldContainer[Hop]
    def __init__(self, unwind: bool = ..., hops: _Optional[_Iterable[_Union[Hop, _Mapping]]] = ...) -> None: ...

class Hop(_message.Message):
    __slots__ = ("port_id", "channel_id")
    PORT_ID_FIELD_NUMBER: _ClassVar[int]
    CHANNEL_ID_FIELD_NUMBER: _ClassVar[int]
    port_id: str
    channel_id: str
    def __init__(self, port_id: _Optional[str] = ..., channel_id: _Optional[str] = ...) -> None: ...
