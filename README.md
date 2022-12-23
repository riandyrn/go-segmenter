# Go Segmenter

This is helper library to help break slice into segments.

For example we have following slice:

```golang
strs := []string{"0", "1", "2", "3", "4"}
```

We want to break it into each segments that has maximum length of `2`. So we would have following segments iteratively:

```golang
[]string{"0", "1"} // first iteration
[]string{"2", "3"} // second iteration
[]string{"4"} // third iteration
```

We could use this library to do that.

## Sample Usage

```golang
func main() {
    sgmntr, _ := segmenter.New(segmenter.Config[string]{
        Slice:       []string{"0", "1", "2", "3", "4"},
        SegmentLength: 2
    })
    for sgmntr.HasNext() {
        log.Println(sgmntr.Next())
    }
    // The result would be:
    //
    // []string{"0", "1"}
    // []string{"2", "3"}
    // []string{"4"}
}
```

For more examples on usage please check `segmenter_test.go`.