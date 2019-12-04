package segmenter_test

import (
	"strings"
	"testing"

	"github.com/riandyrn/go-segmenter"
)

// We are testing this using list of strings, since I think it is the
// simplest one to test
func TestNext(t *testing.T) {
	testCases := []struct {
		Name          string
		Collections   []interface{}
		SegmentLength int
		ExpResults    [][]interface{}
	}{
		{
			Name:          "Segment Length Evenly Divisible With List Length",
			Collections:   []interface{}{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			SegmentLength: 5,
			ExpResults: [][]interface{}{
				[]interface{}{"0", "1", "2", "3", "4"},
				[]interface{}{"5", "6", "7", "8", "9"},
			},
		},
		{
			Name:          "Segment Length Not Evenly Divisible With List Length",
			Collections:   []interface{}{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			SegmentLength: 3,
			ExpResults: [][]interface{}{
				[]interface{}{"0", "1", "2"},
				[]interface{}{"3", "4", "5"},
				[]interface{}{"6", "7", "8"},
				[]interface{}{"9"},
			},
		},
		{
			Name:          "Segment Length Larger Than List Length",
			Collections:   []interface{}{"0", "1", "2", "3"},
			SegmentLength: 5,
			ExpResults: [][]interface{}{
				[]interface{}{"0", "1", "2", "3"},
			},
		},
		{
			Name:          "Zero Segment Length",
			Collections:   []interface{}{"0", "1", "2", "3"},
			SegmentLength: 0,
			ExpResults:    nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			sgmntr := segmenter.NewSegmenter(segmenter.Configs{
				Collections:   testCase.Collections,
				SegmentLength: testCase.SegmentLength,
			})
			var results [][]interface{}
			for sgmntr.HasNext() {
				results = append(results, sgmntr.Next())
			}
			if len(results) != len(testCase.ExpResults) {
				t.Fatalf("unexpected results, expected: %v, got: %v", testCase.ExpResults, results)
			}
			for i := 0; i < len(results); i++ {
				resultStr := strings.Join(toStrings(results[i]), ",")
				expResultStr := strings.Join(toStrings(testCase.ExpResults[i]), ",")
				if resultStr != expResultStr {
					t.Fatalf("unexpected result, expected: %v, got: %v", expResultStr, resultStr)
				}
			}
		})
	}
}

func toStrings(collections []interface{}) []string {
	result := []string{}
	for _, e := range collections {
		v := e.(string)
		result = append(result, v)
	}
	return result
}
