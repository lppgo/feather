// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CategoryMetadata struct {
	_tab flatbuffers.Table
}

func GetRootAsCategoryMetadata(buf []byte, offset flatbuffers.UOffsetT) *CategoryMetadata {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CategoryMetadata{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *CategoryMetadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CategoryMetadata) Table() flatbuffers.Table {
	return rcv._tab
}

/// The category codes are presumed to be integers that are valid indexes into
/// the levels array
func (rcv *CategoryMetadata) Levels(obj *PrimitiveArray) *PrimitiveArray {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PrimitiveArray)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The category codes are presumed to be integers that are valid indexes into
/// the levels array
func (rcv *CategoryMetadata) Ordered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *CategoryMetadata) MutateOrdered(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func CategoryMetadataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CategoryMetadataAddLevels(builder *flatbuffers.Builder, levels flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(levels), 0)
}
func CategoryMetadataAddOrdered(builder *flatbuffers.Builder, ordered bool) {
	builder.PrependBoolSlot(1, ordered, false)
}
func CategoryMetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
