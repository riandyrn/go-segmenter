package segmenter

import (
	"fmt"
	"math"

	"gopkg.in/validator.v2"
)

// Segmenter is used for loading strings in segment
type Segmenter[T comparable] struct {
	collection    []T
	startIdx      int
	endIdx        int
	segmentLength int
	numOps        int
	opsCounter    int
}

// Config holds configuration for Segmenter
type Config[T comparable] struct {
	Collection    []T `validate:"min=1"`
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
		collection:    cfg.Collection,
		segmentLength: cfg.SegmentLength,
		startIdx:      0,
		endIdx:        cfg.SegmentLength,
		numOps:        int(math.Ceil(float64(len(cfg.Collection)) / float64(cfg.SegmentLength))),
		opsCounter:    0,
	}
	if segmenter.endIdx > len(cfg.Collection) {
		segmenter.endIdx = len(cfg.Collection)
	}
	return segmenter, nil
}

// Next returns next segment. In case there is no longer next segment, it returns nil
func (s *Segmenter[T]) Next() []T {
	if s.opsCounter == s.numOps {
		return nil
	}
	defer func() {
		s.opsCounter++
		s.startIdx = s.opsCounter * s.segmentLength
		s.endIdx = s.startIdx + s.segmentLength
		if s.endIdx > len(s.collection) {
			s.endIdx = len(s.collection)
		}
	}()
	return s.collection[s.startIdx:s.endIdx]
}

// HasNext returns true when Segmenter still has next segment
func (s *Segmenter[_]) HasNext() bool {
	return s.opsCounter < s.numOps
}
