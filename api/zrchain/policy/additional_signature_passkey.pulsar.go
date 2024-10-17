// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package policy

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_AdditionalSignaturePasskey                    protoreflect.MessageDescriptor
	fd_AdditionalSignaturePasskey_raw_id             protoreflect.FieldDescriptor
	fd_AdditionalSignaturePasskey_authenticator_data protoreflect.FieldDescriptor
	fd_AdditionalSignaturePasskey_client_data_json   protoreflect.FieldDescriptor
	fd_AdditionalSignaturePasskey_signature          protoreflect.FieldDescriptor
)

func init() {
	file_zrchain_policy_additional_signature_passkey_proto_init()
	md_AdditionalSignaturePasskey = File_zrchain_policy_additional_signature_passkey_proto.Messages().ByName("AdditionalSignaturePasskey")
	fd_AdditionalSignaturePasskey_raw_id = md_AdditionalSignaturePasskey.Fields().ByName("raw_id")
	fd_AdditionalSignaturePasskey_authenticator_data = md_AdditionalSignaturePasskey.Fields().ByName("authenticator_data")
	fd_AdditionalSignaturePasskey_client_data_json = md_AdditionalSignaturePasskey.Fields().ByName("client_data_json")
	fd_AdditionalSignaturePasskey_signature = md_AdditionalSignaturePasskey.Fields().ByName("signature")
}

var _ protoreflect.Message = (*fastReflection_AdditionalSignaturePasskey)(nil)

type fastReflection_AdditionalSignaturePasskey AdditionalSignaturePasskey

func (x *AdditionalSignaturePasskey) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AdditionalSignaturePasskey)(x)
}

func (x *AdditionalSignaturePasskey) slowProtoReflect() protoreflect.Message {
	mi := &file_zrchain_policy_additional_signature_passkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AdditionalSignaturePasskey_messageType fastReflection_AdditionalSignaturePasskey_messageType
var _ protoreflect.MessageType = fastReflection_AdditionalSignaturePasskey_messageType{}

type fastReflection_AdditionalSignaturePasskey_messageType struct{}

func (x fastReflection_AdditionalSignaturePasskey_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AdditionalSignaturePasskey)(nil)
}
func (x fastReflection_AdditionalSignaturePasskey_messageType) New() protoreflect.Message {
	return new(fastReflection_AdditionalSignaturePasskey)
}
func (x fastReflection_AdditionalSignaturePasskey_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AdditionalSignaturePasskey
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AdditionalSignaturePasskey) Descriptor() protoreflect.MessageDescriptor {
	return md_AdditionalSignaturePasskey
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AdditionalSignaturePasskey) Type() protoreflect.MessageType {
	return _fastReflection_AdditionalSignaturePasskey_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AdditionalSignaturePasskey) New() protoreflect.Message {
	return new(fastReflection_AdditionalSignaturePasskey)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AdditionalSignaturePasskey) Interface() protoreflect.ProtoMessage {
	return (*AdditionalSignaturePasskey)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AdditionalSignaturePasskey) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.RawId) != 0 {
		value := protoreflect.ValueOfBytes(x.RawId)
		if !f(fd_AdditionalSignaturePasskey_raw_id, value) {
			return
		}
	}
	if len(x.AuthenticatorData) != 0 {
		value := protoreflect.ValueOfBytes(x.AuthenticatorData)
		if !f(fd_AdditionalSignaturePasskey_authenticator_data, value) {
			return
		}
	}
	if len(x.ClientDataJson) != 0 {
		value := protoreflect.ValueOfBytes(x.ClientDataJson)
		if !f(fd_AdditionalSignaturePasskey_client_data_json, value) {
			return
		}
	}
	if len(x.Signature) != 0 {
		value := protoreflect.ValueOfBytes(x.Signature)
		if !f(fd_AdditionalSignaturePasskey_signature, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_AdditionalSignaturePasskey) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "zrchain.policy.AdditionalSignaturePasskey.raw_id":
		return len(x.RawId) != 0
	case "zrchain.policy.AdditionalSignaturePasskey.authenticator_data":
		return len(x.AuthenticatorData) != 0
	case "zrchain.policy.AdditionalSignaturePasskey.client_data_json":
		return len(x.ClientDataJson) != 0
	case "zrchain.policy.AdditionalSignaturePasskey.signature":
		return len(x.Signature) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.policy.AdditionalSignaturePasskey"))
		}
		panic(fmt.Errorf("message zrchain.policy.AdditionalSignaturePasskey does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AdditionalSignaturePasskey) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "zrchain.policy.AdditionalSignaturePasskey.raw_id":
		x.RawId = nil
	case "zrchain.policy.AdditionalSignaturePasskey.authenticator_data":
		x.AuthenticatorData = nil
	case "zrchain.policy.AdditionalSignaturePasskey.client_data_json":
		x.ClientDataJson = nil
	case "zrchain.policy.AdditionalSignaturePasskey.signature":
		x.Signature = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.policy.AdditionalSignaturePasskey"))
		}
		panic(fmt.Errorf("message zrchain.policy.AdditionalSignaturePasskey does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AdditionalSignaturePasskey) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "zrchain.policy.AdditionalSignaturePasskey.raw_id":
		value := x.RawId
		return protoreflect.ValueOfBytes(value)
	case "zrchain.policy.AdditionalSignaturePasskey.authenticator_data":
		value := x.AuthenticatorData
		return protoreflect.ValueOfBytes(value)
	case "zrchain.policy.AdditionalSignaturePasskey.client_data_json":
		value := x.ClientDataJson
		return protoreflect.ValueOfBytes(value)
	case "zrchain.policy.AdditionalSignaturePasskey.signature":
		value := x.Signature
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.policy.AdditionalSignaturePasskey"))
		}
		panic(fmt.Errorf("message zrchain.policy.AdditionalSignaturePasskey does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AdditionalSignaturePasskey) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "zrchain.policy.AdditionalSignaturePasskey.raw_id":
		x.RawId = value.Bytes()
	case "zrchain.policy.AdditionalSignaturePasskey.authenticator_data":
		x.AuthenticatorData = value.Bytes()
	case "zrchain.policy.AdditionalSignaturePasskey.client_data_json":
		x.ClientDataJson = value.Bytes()
	case "zrchain.policy.AdditionalSignaturePasskey.signature":
		x.Signature = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.policy.AdditionalSignaturePasskey"))
		}
		panic(fmt.Errorf("message zrchain.policy.AdditionalSignaturePasskey does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AdditionalSignaturePasskey) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "zrchain.policy.AdditionalSignaturePasskey.raw_id":
		panic(fmt.Errorf("field raw_id of message zrchain.policy.AdditionalSignaturePasskey is not mutable"))
	case "zrchain.policy.AdditionalSignaturePasskey.authenticator_data":
		panic(fmt.Errorf("field authenticator_data of message zrchain.policy.AdditionalSignaturePasskey is not mutable"))
	case "zrchain.policy.AdditionalSignaturePasskey.client_data_json":
		panic(fmt.Errorf("field client_data_json of message zrchain.policy.AdditionalSignaturePasskey is not mutable"))
	case "zrchain.policy.AdditionalSignaturePasskey.signature":
		panic(fmt.Errorf("field signature of message zrchain.policy.AdditionalSignaturePasskey is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.policy.AdditionalSignaturePasskey"))
		}
		panic(fmt.Errorf("message zrchain.policy.AdditionalSignaturePasskey does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AdditionalSignaturePasskey) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "zrchain.policy.AdditionalSignaturePasskey.raw_id":
		return protoreflect.ValueOfBytes(nil)
	case "zrchain.policy.AdditionalSignaturePasskey.authenticator_data":
		return protoreflect.ValueOfBytes(nil)
	case "zrchain.policy.AdditionalSignaturePasskey.client_data_json":
		return protoreflect.ValueOfBytes(nil)
	case "zrchain.policy.AdditionalSignaturePasskey.signature":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.policy.AdditionalSignaturePasskey"))
		}
		panic(fmt.Errorf("message zrchain.policy.AdditionalSignaturePasskey does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AdditionalSignaturePasskey) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in zrchain.policy.AdditionalSignaturePasskey", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AdditionalSignaturePasskey) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AdditionalSignaturePasskey) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_AdditionalSignaturePasskey) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AdditionalSignaturePasskey) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AdditionalSignaturePasskey)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.RawId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AuthenticatorData)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClientDataJson)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Signature)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*AdditionalSignaturePasskey)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Signature) > 0 {
			i -= len(x.Signature)
			copy(dAtA[i:], x.Signature)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signature)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.ClientDataJson) > 0 {
			i -= len(x.ClientDataJson)
			copy(dAtA[i:], x.ClientDataJson)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClientDataJson)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.AuthenticatorData) > 0 {
			i -= len(x.AuthenticatorData)
			copy(dAtA[i:], x.AuthenticatorData)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AuthenticatorData)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.RawId) > 0 {
			i -= len(x.RawId)
			copy(dAtA[i:], x.RawId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RawId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*AdditionalSignaturePasskey)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AdditionalSignaturePasskey: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AdditionalSignaturePasskey: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RawId", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RawId = append(x.RawId[:0], dAtA[iNdEx:postIndex]...)
				if x.RawId == nil {
					x.RawId = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuthenticatorData", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AuthenticatorData = append(x.AuthenticatorData[:0], dAtA[iNdEx:postIndex]...)
				if x.AuthenticatorData == nil {
					x.AuthenticatorData = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClientDataJson", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ClientDataJson = append(x.ClientDataJson[:0], dAtA[iNdEx:postIndex]...)
				if x.ClientDataJson == nil {
					x.ClientDataJson = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signature = append(x.Signature[:0], dAtA[iNdEx:postIndex]...)
				if x.Signature == nil {
					x.Signature = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: zrchain/policy/additional_signature_passkey.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AdditionalSignaturePasskey is a message that contains passkey signature data
type AdditionalSignaturePasskey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawId             []byte `protobuf:"bytes,1,opt,name=raw_id,json=rawId,proto3" json:"raw_id,omitempty"`
	AuthenticatorData []byte `protobuf:"bytes,2,opt,name=authenticator_data,json=authenticatorData,proto3" json:"authenticator_data,omitempty"`
	ClientDataJson    []byte `protobuf:"bytes,3,opt,name=client_data_json,json=clientDataJson,proto3" json:"client_data_json,omitempty"`
	Signature         []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *AdditionalSignaturePasskey) Reset() {
	*x = AdditionalSignaturePasskey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zrchain_policy_additional_signature_passkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionalSignaturePasskey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionalSignaturePasskey) ProtoMessage() {}

// Deprecated: Use AdditionalSignaturePasskey.ProtoReflect.Descriptor instead.
func (*AdditionalSignaturePasskey) Descriptor() ([]byte, []int) {
	return file_zrchain_policy_additional_signature_passkey_proto_rawDescGZIP(), []int{0}
}

func (x *AdditionalSignaturePasskey) GetRawId() []byte {
	if x != nil {
		return x.RawId
	}
	return nil
}

func (x *AdditionalSignaturePasskey) GetAuthenticatorData() []byte {
	if x != nil {
		return x.AuthenticatorData
	}
	return nil
}

func (x *AdditionalSignaturePasskey) GetClientDataJson() []byte {
	if x != nil {
		return x.ClientDataJson
	}
	return nil
}

func (x *AdditionalSignaturePasskey) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_zrchain_policy_additional_signature_passkey_proto protoreflect.FileDescriptor

var file_zrchain_policy_additional_signature_passkey_proto_rawDesc = []byte{
	0x0a, 0x31, 0x7a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x7a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x22, 0xaa, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x6b,
	0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x73,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x42, 0xaf, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x1f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x6b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x72, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0xa2, 0x02, 0x03, 0x5a, 0x50,
	0x58, 0xaa, 0x02, 0x0e, 0x5a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0xca, 0x02, 0x0e, 0x5a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0xe2, 0x02, 0x1a, 0x5a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0f, 0x5a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zrchain_policy_additional_signature_passkey_proto_rawDescOnce sync.Once
	file_zrchain_policy_additional_signature_passkey_proto_rawDescData = file_zrchain_policy_additional_signature_passkey_proto_rawDesc
)

func file_zrchain_policy_additional_signature_passkey_proto_rawDescGZIP() []byte {
	file_zrchain_policy_additional_signature_passkey_proto_rawDescOnce.Do(func() {
		file_zrchain_policy_additional_signature_passkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_zrchain_policy_additional_signature_passkey_proto_rawDescData)
	})
	return file_zrchain_policy_additional_signature_passkey_proto_rawDescData
}

var file_zrchain_policy_additional_signature_passkey_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_zrchain_policy_additional_signature_passkey_proto_goTypes = []interface{}{
	(*AdditionalSignaturePasskey)(nil), // 0: zrchain.policy.AdditionalSignaturePasskey
}
var file_zrchain_policy_additional_signature_passkey_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zrchain_policy_additional_signature_passkey_proto_init() }
func file_zrchain_policy_additional_signature_passkey_proto_init() {
	if File_zrchain_policy_additional_signature_passkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zrchain_policy_additional_signature_passkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionalSignaturePasskey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zrchain_policy_additional_signature_passkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zrchain_policy_additional_signature_passkey_proto_goTypes,
		DependencyIndexes: file_zrchain_policy_additional_signature_passkey_proto_depIdxs,
		MessageInfos:      file_zrchain_policy_additional_signature_passkey_proto_msgTypes,
	}.Build()
	File_zrchain_policy_additional_signature_passkey_proto = out.File
	file_zrchain_policy_additional_signature_passkey_proto_rawDesc = nil
	file_zrchain_policy_additional_signature_passkey_proto_goTypes = nil
	file_zrchain_policy_additional_signature_passkey_proto_depIdxs = nil
}
