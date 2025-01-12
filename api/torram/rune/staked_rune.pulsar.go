// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package rune

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
	md_StakedRune             protoreflect.MessageDescriptor
	fd_StakedRune_runeId      protoreflect.FieldDescriptor
	fd_StakedRune_owner       protoreflect.FieldDescriptor
	fd_StakedRune_amount      protoreflect.FieldDescriptor
	fd_StakedRune_stakingTime protoreflect.FieldDescriptor
	fd_StakedRune_status      protoreflect.FieldDescriptor
)

func init() {
	file_torram_rune_staked_rune_proto_init()
	md_StakedRune = File_torram_rune_staked_rune_proto.Messages().ByName("StakedRune")
	fd_StakedRune_runeId = md_StakedRune.Fields().ByName("runeId")
	fd_StakedRune_owner = md_StakedRune.Fields().ByName("owner")
	fd_StakedRune_amount = md_StakedRune.Fields().ByName("amount")
	fd_StakedRune_stakingTime = md_StakedRune.Fields().ByName("stakingTime")
	fd_StakedRune_status = md_StakedRune.Fields().ByName("status")
}

var _ protoreflect.Message = (*fastReflection_StakedRune)(nil)

type fastReflection_StakedRune StakedRune

func (x *StakedRune) ProtoReflect() protoreflect.Message {
	return (*fastReflection_StakedRune)(x)
}

func (x *StakedRune) slowProtoReflect() protoreflect.Message {
	mi := &file_torram_rune_staked_rune_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_StakedRune_messageType fastReflection_StakedRune_messageType
var _ protoreflect.MessageType = fastReflection_StakedRune_messageType{}

type fastReflection_StakedRune_messageType struct{}

func (x fastReflection_StakedRune_messageType) Zero() protoreflect.Message {
	return (*fastReflection_StakedRune)(nil)
}
func (x fastReflection_StakedRune_messageType) New() protoreflect.Message {
	return new(fastReflection_StakedRune)
}
func (x fastReflection_StakedRune_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_StakedRune
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_StakedRune) Descriptor() protoreflect.MessageDescriptor {
	return md_StakedRune
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_StakedRune) Type() protoreflect.MessageType {
	return _fastReflection_StakedRune_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_StakedRune) New() protoreflect.Message {
	return new(fastReflection_StakedRune)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_StakedRune) Interface() protoreflect.ProtoMessage {
	return (*StakedRune)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_StakedRune) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.RuneId != "" {
		value := protoreflect.ValueOfString(x.RuneId)
		if !f(fd_StakedRune_runeId, value) {
			return
		}
	}
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_StakedRune_owner, value) {
			return
		}
	}
	if x.Amount != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Amount)
		if !f(fd_StakedRune_amount, value) {
			return
		}
	}
	if x.StakingTime != uint64(0) {
		value := protoreflect.ValueOfUint64(x.StakingTime)
		if !f(fd_StakedRune_stakingTime, value) {
			return
		}
	}
	if x.Status != "" {
		value := protoreflect.ValueOfString(x.Status)
		if !f(fd_StakedRune_status, value) {
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
func (x *fastReflection_StakedRune) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "torram.rune.StakedRune.runeId":
		return x.RuneId != ""
	case "torram.rune.StakedRune.owner":
		return x.Owner != ""
	case "torram.rune.StakedRune.amount":
		return x.Amount != uint64(0)
	case "torram.rune.StakedRune.stakingTime":
		return x.StakingTime != uint64(0)
	case "torram.rune.StakedRune.status":
		return x.Status != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: torram.rune.StakedRune"))
		}
		panic(fmt.Errorf("message torram.rune.StakedRune does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StakedRune) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "torram.rune.StakedRune.runeId":
		x.RuneId = ""
	case "torram.rune.StakedRune.owner":
		x.Owner = ""
	case "torram.rune.StakedRune.amount":
		x.Amount = uint64(0)
	case "torram.rune.StakedRune.stakingTime":
		x.StakingTime = uint64(0)
	case "torram.rune.StakedRune.status":
		x.Status = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: torram.rune.StakedRune"))
		}
		panic(fmt.Errorf("message torram.rune.StakedRune does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_StakedRune) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "torram.rune.StakedRune.runeId":
		value := x.RuneId
		return protoreflect.ValueOfString(value)
	case "torram.rune.StakedRune.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "torram.rune.StakedRune.amount":
		value := x.Amount
		return protoreflect.ValueOfUint64(value)
	case "torram.rune.StakedRune.stakingTime":
		value := x.StakingTime
		return protoreflect.ValueOfUint64(value)
	case "torram.rune.StakedRune.status":
		value := x.Status
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: torram.rune.StakedRune"))
		}
		panic(fmt.Errorf("message torram.rune.StakedRune does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_StakedRune) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "torram.rune.StakedRune.runeId":
		x.RuneId = value.Interface().(string)
	case "torram.rune.StakedRune.owner":
		x.Owner = value.Interface().(string)
	case "torram.rune.StakedRune.amount":
		x.Amount = value.Uint()
	case "torram.rune.StakedRune.stakingTime":
		x.StakingTime = value.Uint()
	case "torram.rune.StakedRune.status":
		x.Status = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: torram.rune.StakedRune"))
		}
		panic(fmt.Errorf("message torram.rune.StakedRune does not contain field %s", fd.FullName()))
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
func (x *fastReflection_StakedRune) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "torram.rune.StakedRune.runeId":
		panic(fmt.Errorf("field runeId of message torram.rune.StakedRune is not mutable"))
	case "torram.rune.StakedRune.owner":
		panic(fmt.Errorf("field owner of message torram.rune.StakedRune is not mutable"))
	case "torram.rune.StakedRune.amount":
		panic(fmt.Errorf("field amount of message torram.rune.StakedRune is not mutable"))
	case "torram.rune.StakedRune.stakingTime":
		panic(fmt.Errorf("field stakingTime of message torram.rune.StakedRune is not mutable"))
	case "torram.rune.StakedRune.status":
		panic(fmt.Errorf("field status of message torram.rune.StakedRune is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: torram.rune.StakedRune"))
		}
		panic(fmt.Errorf("message torram.rune.StakedRune does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_StakedRune) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "torram.rune.StakedRune.runeId":
		return protoreflect.ValueOfString("")
	case "torram.rune.StakedRune.owner":
		return protoreflect.ValueOfString("")
	case "torram.rune.StakedRune.amount":
		return protoreflect.ValueOfUint64(uint64(0))
	case "torram.rune.StakedRune.stakingTime":
		return protoreflect.ValueOfUint64(uint64(0))
	case "torram.rune.StakedRune.status":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: torram.rune.StakedRune"))
		}
		panic(fmt.Errorf("message torram.rune.StakedRune does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_StakedRune) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in torram.rune.StakedRune", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_StakedRune) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StakedRune) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_StakedRune) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_StakedRune) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*StakedRune)
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
		l = len(x.RuneId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Amount != 0 {
			n += 1 + runtime.Sov(uint64(x.Amount))
		}
		if x.StakingTime != 0 {
			n += 1 + runtime.Sov(uint64(x.StakingTime))
		}
		l = len(x.Status)
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
		x := input.Message.Interface().(*StakedRune)
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
		if len(x.Status) > 0 {
			i -= len(x.Status)
			copy(dAtA[i:], x.Status)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Status)))
			i--
			dAtA[i] = 0x2a
		}
		if x.StakingTime != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StakingTime))
			i--
			dAtA[i] = 0x20
		}
		if x.Amount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Amount))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.RuneId) > 0 {
			i -= len(x.RuneId)
			copy(dAtA[i:], x.RuneId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RuneId)))
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
		x := input.Message.Interface().(*StakedRune)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StakedRune: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StakedRune: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RuneId", wireType)
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
				x.RuneId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				x.Amount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Amount |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StakingTime", wireType)
				}
				x.StakingTime = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StakingTime |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
				x.Status = string(dAtA[iNdEx:postIndex])
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
// source: torram/rune/staked_rune.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StakedRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuneId      string `protobuf:"bytes,1,opt,name=runeId,proto3" json:"runeId,omitempty"`
	Owner       string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Amount      uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	StakingTime uint64 `protobuf:"varint,4,opt,name=stakingTime,proto3" json:"stakingTime,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StakedRune) Reset() {
	*x = StakedRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torram_rune_staked_rune_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakedRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakedRune) ProtoMessage() {}

// Deprecated: Use StakedRune.ProtoReflect.Descriptor instead.
func (*StakedRune) Descriptor() ([]byte, []int) {
	return file_torram_rune_staked_rune_proto_rawDescGZIP(), []int{0}
}

func (x *StakedRune) GetRuneId() string {
	if x != nil {
		return x.RuneId
	}
	return ""
}

func (x *StakedRune) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *StakedRune) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *StakedRune) GetStakingTime() uint64 {
	if x != nil {
		return x.StakingTime
	}
	return 0
}

func (x *StakedRune) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_torram_rune_staked_rune_proto protoreflect.FileDescriptor

var file_torram_rune_staked_rune_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x6f, 0x72, 0x72, 0x61, 0x6d, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x2f, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x74, 0x6f, 0x72, 0x72, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x22, 0x8c, 0x01, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x75, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x87, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x72, 0x72, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x6e, 0x65, 0x42,
	0x0f, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x16, 0x74, 0x6f, 0x72, 0x72, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x6f, 0x72, 0x72, 0x61, 0x6d, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0xa2, 0x02, 0x03, 0x54, 0x52, 0x58,
	0xaa, 0x02, 0x0b, 0x54, 0x6f, 0x72, 0x72, 0x61, 0x6d, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0xca, 0x02,
	0x0b, 0x54, 0x6f, 0x72, 0x72, 0x61, 0x6d, 0x5c, 0x52, 0x75, 0x6e, 0x65, 0xe2, 0x02, 0x17, 0x54,
	0x6f, 0x72, 0x72, 0x61, 0x6d, 0x5c, 0x52, 0x75, 0x6e, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x54, 0x6f, 0x72, 0x72, 0x61, 0x6d, 0x3a,
	0x3a, 0x52, 0x75, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_torram_rune_staked_rune_proto_rawDescOnce sync.Once
	file_torram_rune_staked_rune_proto_rawDescData = file_torram_rune_staked_rune_proto_rawDesc
)

func file_torram_rune_staked_rune_proto_rawDescGZIP() []byte {
	file_torram_rune_staked_rune_proto_rawDescOnce.Do(func() {
		file_torram_rune_staked_rune_proto_rawDescData = protoimpl.X.CompressGZIP(file_torram_rune_staked_rune_proto_rawDescData)
	})
	return file_torram_rune_staked_rune_proto_rawDescData
}

var file_torram_rune_staked_rune_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_torram_rune_staked_rune_proto_goTypes = []interface{}{
	(*StakedRune)(nil), // 0: torram.rune.StakedRune
}
var file_torram_rune_staked_rune_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_torram_rune_staked_rune_proto_init() }
func file_torram_rune_staked_rune_proto_init() {
	if File_torram_rune_staked_rune_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_torram_rune_staked_rune_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakedRune); i {
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
			RawDescriptor: file_torram_rune_staked_rune_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_torram_rune_staked_rune_proto_goTypes,
		DependencyIndexes: file_torram_rune_staked_rune_proto_depIdxs,
		MessageInfos:      file_torram_rune_staked_rune_proto_msgTypes,
	}.Build()
	File_torram_rune_staked_rune_proto = out.File
	file_torram_rune_staked_rune_proto_rawDesc = nil
	file_torram_rune_staked_rune_proto_goTypes = nil
	file_torram_rune_staked_rune_proto_depIdxs = nil
}
