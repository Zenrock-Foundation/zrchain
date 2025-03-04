// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package treasury

import (
	_ "cosmossdk.io/api/amino"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Params                                protoreflect.MessageDescriptor
	fd_Params_mpc_keyring                    protoreflect.FieldDescriptor
	fd_Params_zr_sign_address                protoreflect.FieldDescriptor
	fd_Params_keyring_commission             protoreflect.FieldDescriptor
	fd_Params_keyring_commission_destination protoreflect.FieldDescriptor
	fd_Params_min_gas_fee                    protoreflect.FieldDescriptor
	fd_Params_default_btl                    protoreflect.FieldDescriptor
)

func init() {
	file_zrchain_treasury_params_proto_init()
	md_Params = File_zrchain_treasury_params_proto.Messages().ByName("Params")
	fd_Params_mpc_keyring = md_Params.Fields().ByName("mpc_keyring")
	fd_Params_zr_sign_address = md_Params.Fields().ByName("zr_sign_address")
	fd_Params_keyring_commission = md_Params.Fields().ByName("keyring_commission")
	fd_Params_keyring_commission_destination = md_Params.Fields().ByName("keyring_commission_destination")
	fd_Params_min_gas_fee = md_Params.Fields().ByName("min_gas_fee")
	fd_Params_default_btl = md_Params.Fields().ByName("default_btl")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_zrchain_treasury_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.MpcKeyring != "" {
		value := protoreflect.ValueOfString(x.MpcKeyring)
		if !f(fd_Params_mpc_keyring, value) {
			return
		}
	}
	if x.ZrSignAddress != "" {
		value := protoreflect.ValueOfString(x.ZrSignAddress)
		if !f(fd_Params_zr_sign_address, value) {
			return
		}
	}
	if x.KeyringCommission != uint64(0) {
		value := protoreflect.ValueOfUint64(x.KeyringCommission)
		if !f(fd_Params_keyring_commission, value) {
			return
		}
	}
	if x.KeyringCommissionDestination != "" {
		value := protoreflect.ValueOfString(x.KeyringCommissionDestination)
		if !f(fd_Params_keyring_commission_destination, value) {
			return
		}
	}
	if x.MinGasFee != "" {
		value := protoreflect.ValueOfString(x.MinGasFee)
		if !f(fd_Params_min_gas_fee, value) {
			return
		}
	}
	if x.DefaultBtl != uint64(0) {
		value := protoreflect.ValueOfUint64(x.DefaultBtl)
		if !f(fd_Params_default_btl, value) {
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
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "zrchain.treasury.Params.mpc_keyring":
		return x.MpcKeyring != ""
	case "zrchain.treasury.Params.zr_sign_address":
		return x.ZrSignAddress != ""
	case "zrchain.treasury.Params.keyring_commission":
		return x.KeyringCommission != uint64(0)
	case "zrchain.treasury.Params.keyring_commission_destination":
		return x.KeyringCommissionDestination != ""
	case "zrchain.treasury.Params.min_gas_fee":
		return x.MinGasFee != ""
	case "zrchain.treasury.Params.default_btl":
		return x.DefaultBtl != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.treasury.Params"))
		}
		panic(fmt.Errorf("message zrchain.treasury.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "zrchain.treasury.Params.mpc_keyring":
		x.MpcKeyring = ""
	case "zrchain.treasury.Params.zr_sign_address":
		x.ZrSignAddress = ""
	case "zrchain.treasury.Params.keyring_commission":
		x.KeyringCommission = uint64(0)
	case "zrchain.treasury.Params.keyring_commission_destination":
		x.KeyringCommissionDestination = ""
	case "zrchain.treasury.Params.min_gas_fee":
		x.MinGasFee = ""
	case "zrchain.treasury.Params.default_btl":
		x.DefaultBtl = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.treasury.Params"))
		}
		panic(fmt.Errorf("message zrchain.treasury.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "zrchain.treasury.Params.mpc_keyring":
		value := x.MpcKeyring
		return protoreflect.ValueOfString(value)
	case "zrchain.treasury.Params.zr_sign_address":
		value := x.ZrSignAddress
		return protoreflect.ValueOfString(value)
	case "zrchain.treasury.Params.keyring_commission":
		value := x.KeyringCommission
		return protoreflect.ValueOfUint64(value)
	case "zrchain.treasury.Params.keyring_commission_destination":
		value := x.KeyringCommissionDestination
		return protoreflect.ValueOfString(value)
	case "zrchain.treasury.Params.min_gas_fee":
		value := x.MinGasFee
		return protoreflect.ValueOfString(value)
	case "zrchain.treasury.Params.default_btl":
		value := x.DefaultBtl
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.treasury.Params"))
		}
		panic(fmt.Errorf("message zrchain.treasury.Params does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "zrchain.treasury.Params.mpc_keyring":
		x.MpcKeyring = value.Interface().(string)
	case "zrchain.treasury.Params.zr_sign_address":
		x.ZrSignAddress = value.Interface().(string)
	case "zrchain.treasury.Params.keyring_commission":
		x.KeyringCommission = value.Uint()
	case "zrchain.treasury.Params.keyring_commission_destination":
		x.KeyringCommissionDestination = value.Interface().(string)
	case "zrchain.treasury.Params.min_gas_fee":
		x.MinGasFee = value.Interface().(string)
	case "zrchain.treasury.Params.default_btl":
		x.DefaultBtl = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.treasury.Params"))
		}
		panic(fmt.Errorf("message zrchain.treasury.Params does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "zrchain.treasury.Params.mpc_keyring":
		panic(fmt.Errorf("field mpc_keyring of message zrchain.treasury.Params is not mutable"))
	case "zrchain.treasury.Params.zr_sign_address":
		panic(fmt.Errorf("field zr_sign_address of message zrchain.treasury.Params is not mutable"))
	case "zrchain.treasury.Params.keyring_commission":
		panic(fmt.Errorf("field keyring_commission of message zrchain.treasury.Params is not mutable"))
	case "zrchain.treasury.Params.keyring_commission_destination":
		panic(fmt.Errorf("field keyring_commission_destination of message zrchain.treasury.Params is not mutable"))
	case "zrchain.treasury.Params.min_gas_fee":
		panic(fmt.Errorf("field min_gas_fee of message zrchain.treasury.Params is not mutable"))
	case "zrchain.treasury.Params.default_btl":
		panic(fmt.Errorf("field default_btl of message zrchain.treasury.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.treasury.Params"))
		}
		panic(fmt.Errorf("message zrchain.treasury.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "zrchain.treasury.Params.mpc_keyring":
		return protoreflect.ValueOfString("")
	case "zrchain.treasury.Params.zr_sign_address":
		return protoreflect.ValueOfString("")
	case "zrchain.treasury.Params.keyring_commission":
		return protoreflect.ValueOfUint64(uint64(0))
	case "zrchain.treasury.Params.keyring_commission_destination":
		return protoreflect.ValueOfString("")
	case "zrchain.treasury.Params.min_gas_fee":
		return protoreflect.ValueOfString("")
	case "zrchain.treasury.Params.default_btl":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: zrchain.treasury.Params"))
		}
		panic(fmt.Errorf("message zrchain.treasury.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in zrchain.treasury.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
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
		l = len(x.MpcKeyring)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ZrSignAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.KeyringCommission != 0 {
			n += 1 + runtime.Sov(uint64(x.KeyringCommission))
		}
		l = len(x.KeyringCommissionDestination)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MinGasFee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DefaultBtl != 0 {
			n += 1 + runtime.Sov(uint64(x.DefaultBtl))
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
		x := input.Message.Interface().(*Params)
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
		if x.DefaultBtl != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.DefaultBtl))
			i--
			dAtA[i] = 0x30
		}
		if len(x.MinGasFee) > 0 {
			i -= len(x.MinGasFee)
			copy(dAtA[i:], x.MinGasFee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MinGasFee)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.KeyringCommissionDestination) > 0 {
			i -= len(x.KeyringCommissionDestination)
			copy(dAtA[i:], x.KeyringCommissionDestination)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.KeyringCommissionDestination)))
			i--
			dAtA[i] = 0x22
		}
		if x.KeyringCommission != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.KeyringCommission))
			i--
			dAtA[i] = 0x18
		}
		if len(x.ZrSignAddress) > 0 {
			i -= len(x.ZrSignAddress)
			copy(dAtA[i:], x.ZrSignAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ZrSignAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.MpcKeyring) > 0 {
			i -= len(x.MpcKeyring)
			copy(dAtA[i:], x.MpcKeyring)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MpcKeyring)))
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
		x := input.Message.Interface().(*Params)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MpcKeyring", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MpcKeyring = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ZrSignAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ZrSignAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field KeyringCommission", wireType)
				}
				x.KeyringCommission = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.KeyringCommission |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field KeyringCommissionDestination", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.KeyringCommissionDestination = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinGasFee", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MinGasFee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DefaultBtl", wireType)
				}
				x.DefaultBtl = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.DefaultBtl |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: zrchain/treasury/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the parameters for the module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MpcKeyring                   string `protobuf:"bytes,1,opt,name=mpc_keyring,json=mpcKeyring,proto3" json:"mpc_keyring,omitempty"`
	ZrSignAddress                string `protobuf:"bytes,2,opt,name=zr_sign_address,json=zrSignAddress,proto3" json:"zr_sign_address,omitempty"`
	KeyringCommission            uint64 `protobuf:"varint,3,opt,name=keyring_commission,json=keyringCommission,proto3" json:"keyring_commission,omitempty"`
	KeyringCommissionDestination string `protobuf:"bytes,4,opt,name=keyring_commission_destination,json=keyringCommissionDestination,proto3" json:"keyring_commission_destination,omitempty"`
	MinGasFee                    string `protobuf:"bytes,5,opt,name=min_gas_fee,json=minGasFee,proto3" json:"min_gas_fee,omitempty"`
	DefaultBtl                   uint64 `protobuf:"varint,6,opt,name=default_btl,json=defaultBtl,proto3" json:"default_btl,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zrchain_treasury_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_zrchain_treasury_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetMpcKeyring() string {
	if x != nil {
		return x.MpcKeyring
	}
	return ""
}

func (x *Params) GetZrSignAddress() string {
	if x != nil {
		return x.ZrSignAddress
	}
	return ""
}

func (x *Params) GetKeyringCommission() uint64 {
	if x != nil {
		return x.KeyringCommission
	}
	return 0
}

func (x *Params) GetKeyringCommissionDestination() string {
	if x != nil {
		return x.KeyringCommissionDestination
	}
	return ""
}

func (x *Params) GetMinGasFee() string {
	if x != nil {
		return x.MinGasFee
	}
	return ""
}

func (x *Params) GetDefaultBtl() uint64 {
	if x != nil {
		return x.DefaultBtl
	}
	return 0
}

var File_zrchain_treasury_params_proto protoreflect.FileDescriptor

var file_zrchain_treasury_params_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x7a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x7a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x79, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x70, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x70, 0x63, 0x4b,
	0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x7a, 0x72, 0x5f, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x7a, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d,
	0x0a, 0x12, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6b, 0x65, 0x79, 0x72,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a,
	0x1e, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x61, 0x73,
	0x46, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x62,
	0x74, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x42, 0x74, 0x6c, 0x3a, 0x43, 0xe8, 0xa0, 0x1f, 0x01, 0x8a, 0xe7, 0xb0, 0x2a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x65, 0x6e, 0x72, 0x6f, 0x63,
	0x6b, 0x2d, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7a, 0x72, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x35, 0x2f, 0x78, 0x2f, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x2f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0xa7, 0x01, 0x0a, 0x14, 0x63, 0x6f,
	0x6d, 0x2e, 0x7a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x42, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x21, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x7a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x79, 0xa2, 0x02, 0x03, 0x5a, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x5a, 0x72, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0xca, 0x02, 0x10,
	0x5a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79,
	0xe2, 0x02, 0x1c, 0x5a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x54, 0x72, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x11, 0x5a, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x54, 0x72, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zrchain_treasury_params_proto_rawDescOnce sync.Once
	file_zrchain_treasury_params_proto_rawDescData = file_zrchain_treasury_params_proto_rawDesc
)

func file_zrchain_treasury_params_proto_rawDescGZIP() []byte {
	file_zrchain_treasury_params_proto_rawDescOnce.Do(func() {
		file_zrchain_treasury_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_zrchain_treasury_params_proto_rawDescData)
	})
	return file_zrchain_treasury_params_proto_rawDescData
}

var file_zrchain_treasury_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_zrchain_treasury_params_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: zrchain.treasury.Params
}
var file_zrchain_treasury_params_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zrchain_treasury_params_proto_init() }
func file_zrchain_treasury_params_proto_init() {
	if File_zrchain_treasury_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zrchain_treasury_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
			RawDescriptor: file_zrchain_treasury_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zrchain_treasury_params_proto_goTypes,
		DependencyIndexes: file_zrchain_treasury_params_proto_depIdxs,
		MessageInfos:      file_zrchain_treasury_params_proto_msgTypes,
	}.Build()
	File_zrchain_treasury_params_proto = out.File
	file_zrchain_treasury_params_proto_rawDesc = nil
	file_zrchain_treasury_params_proto_goTypes = nil
	file_zrchain_treasury_params_proto_depIdxs = nil
}
