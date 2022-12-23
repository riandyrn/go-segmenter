package segmenter

import (
	"fmt"
	"math"

	"gopkg.in/validator.v2"
)

// Segmenter is used for loading strings in segment
type Segmenter[T comparable] struct {
	items         []T
	startIdx      int
	endIdx        int
	segmentLength int
	numOps        int
	opsCounter    int
}

// Config holds configs for Segmenter
type Config[T comparable] struct {
	Items         []T `validate:"min=1"`
	SegmentLength int `validate:"min=1"`
}

// New returns new instance of Segmenter
func New[T comparable](cfg Config[T]) (*Segmenter[T], error) {
	// validate config
	err := validator.Validate(cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	// construct segmenter
	segmenter := &Segmenter[T]{
		items:         cfg.Items,
		segmentLength: cfg.SegmentLength,
		startIdx:      0,
		endIdx:        cfg.SegmentLength,
		numOps:        int(math.Ceil(float64(len(cfg.Items)) / float64(cfg.SegmentLength))),
		opsCounter:    0,
	}
	if segmenter.endIdx > len(cfg.Items) {
		segmenter.endIdx = len(cfg.Items)
	}
	return segmenter, nil
}

// Next returns next strings segment, when there is no next
// segment the returned value is nil
func (s *Segmenter[T]) Next() []T {
	if s.opsCounter == s.numOps {
		return nil
	}
	defer func() {
		s.opsCounter++
		s.startIdx = s.opsCounter * s.segmentLength
		s.endIdx = s.startIdx + s.segmentLength
		if s.endIdx > len(s.items) {
			s.endIdx = len(s.items)
		}
	}()
	return s.items[s.startIdx:s.endIdx]
}

// HasNext returns true when strings still has next segment
func (s *Segmenter[_]) HasNext() bool {
	return s.opsCounter < s.numOps
}
