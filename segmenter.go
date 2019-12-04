package segmenter

import "math"

// Segmenter is used for loading strings in segment
type Segmenter struct {
	collections   []interface{}
	startIdx      int
	endIdx        int
	segmentLength int
	numOps        int
	opsCounter    int
}

// Configs holds configs for StringsSegmenter
type Configs struct {
	Collections   []interface{}
	SegmentLength int
}

// NewSegmenter returns new instance of StringsSegmenter
func NewSegmenter(configs Configs) *Segmenter {
	segmenter := &Segmenter{
		collections:   configs.Collections,
		segmentLength: configs.SegmentLength,
		startIdx:      0,
		endIdx:        configs.SegmentLength,
		numOps:        int(math.Ceil(float64(len(configs.Collections)) / float64(configs.SegmentLength))),
		opsCounter:    0,
	}
	if segmenter.endIdx > len(configs.Collections) {
		segmenter.endIdx = len(configs.Collections)
	}
	return segmenter
}

// Next returns next strings segment, when there is no next
// segment the returned value is nil
func (s *Segmenter) Next() []interface{} {
	if s.opsCounter == s.numOps {
		return nil
	}
	defer func() {
		s.opsCounter++
		s.startIdx = s.opsCounter * s.segmentLength
		s.endIdx = s.startIdx + s.segmentLength
		if s.endIdx > len(s.collections) {
			s.endIdx = len(s.collections)
		}
	}()
	return s.collections[s.startIdx:s.endIdx]
}

// HasNext returns true when strings still has next segment
func (s *Segmenter) HasNext() bool {
	return s.opsCounter < s.numOps
}
