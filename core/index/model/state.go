package model

import (
	"github.com/balzaczyy/golucene/core/store"
	"github.com/balzaczyy/golucene/core/util"
)

// index/SegmentWriteState.java

/* Holder class for common parameters used during write. */
type SegmentWriteState struct {
	infoStream        util.InfoStream
	Directory         store.Directory
	SegmentInfo       *SegmentInfo
	FieldInfos        FieldInfos
	DelCountOnFlush   int
	SegDeletes        interface{} //*BufferedDeletes
	LiveDocs          util.MutableBits
	SegmentSuffix     string
	termIndexInternal int
	Context           store.IOContext
}

func NewSegmentWriteState(infoStream util.InfoStream,
	dir store.Directory, segmentInfo *SegmentInfo,
	fieldInfos FieldInfos, termIndexInterval int,
	segDeletes interface{} /**BufferedDeletes*/, ctx store.IOContext) SegmentWriteState {

	return SegmentWriteState{
		infoStream, dir, segmentInfo, fieldInfos, 0, segDeletes, nil, "",
		termIndexInterval, ctx}
}

/* Create a shallow copy of SegmentWriteState with a new segment suffix. */
func NewSegmentWriteStateFrom(state SegmentWriteState,
	segmentSuffix string) SegmentWriteState {

	return SegmentWriteState{
		state.infoStream,
		state.Directory,
		state.SegmentInfo,
		state.FieldInfos,
		state.DelCountOnFlush,
		state.SegDeletes,
		nil,
		segmentSuffix,
		state.termIndexInternal,
		state.Context,
	}
}