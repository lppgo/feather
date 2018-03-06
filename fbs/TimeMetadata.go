// automatically generated by the FlatBuffers compiler, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TimeMetadata struct {
	_tab flatbuffers.Table
}

func GetRootAsTimeMetadata(buf []byte, offset flatbuffers.UOffsetT) *TimeMetadata {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TimeMetadata{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TimeMetadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TimeMetadata) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TimeMetadata) Unit() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TimeMetadata) MutateUnit(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func TimeMetadataStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TimeMetadataAddUnit(builder *flatbuffers.Builder, unit int8) {
	builder.PrependInt8Slot(0, unit, 0)
}
func TimeMetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
