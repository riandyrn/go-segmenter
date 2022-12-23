package segmenter_test

import (
	"testing"

	"github.com/riandyrn/go-segmenter"
	"github.com/stretchr/testify/require"
)

func TestNext(t *testing.T) {
	testCases := []struct {
		Name          string
		Items         []string
		SegmentLength int
		ExpResults    [][]string
	}{
		{
			Name:          "Segment Length Evenly Divisible With List Length",
			Items:         []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			SegmentLength: 5,
			ExpResults: [][]string{
				{"0", "1", "2", "3", "4"},
				{"5", "6", "7", "8", "9"},
			},
		},
		{
			Name:          "Segment Length Not Evenly Divisible With List Length",
			Items:         []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			SegmentLength: 3,
			ExpResults: [][]string{
				{"0", "1", "2"},
				{"3", "4", "5"},
				{"6", "7", "8"},
				{"9"},
			},
		},
		{
			Name:          "Segment Length Larger Than List Length",
			Items:         []string{"0", "1", "2", "3"},
			SegmentLength: 5,
			ExpResults: [][]string{
				{"0", "1", "2", "3"},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			// initialize segmenter
			sgmntr, err := segmenter.New(segmenter.Config[string]{
				Items:         testCase.Items,
				SegmentLength: testCase.SegmentLength,
			})
			require.NoError(t, err)
			// collect all results
			var results [][]string
			for sgmntr.HasNext() {
				results = append(results, sgmntr.Next())
			}
			// compare collected results
			require.Equal(t, testCase.ExpResults, results)
		})
	}
}
