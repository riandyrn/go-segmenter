# Go Segmenter

This is library to break collections ([]interface{}) into segments iteratively. This project is actually generalization of [Go Strings Segmenter](https://github.com/riandyrn/go-strings-segmenter).

For example we have following collections:

```golang
strs := []interface{}{"0", "1", "2", "3", "4"}
```

We want to break this slice of strings into segment that has maximum length of `2`. So we would have following segments iteratively:

```golang
[]interface{}{"0", "1"} // first iteration
[]interface{}{"2", "3"} // second iteration
[]interface{}{"4"} // third iteration
```

We could use this library to do just that.

## Sample Usage

```golang
func main() {
    sgmntr := segmenter.NewSegmenter(segmenter.Configs{
        Strings:       []interface{}{"0", "1", "2", "3", "4"},
        SegmentLength: 2
    })
    for sgmntr.HasNext() {
        log.Println(sgmntr.Next())
    }
    // The result would be:
    //
    // []interface{}{"0", "1"}
    // []interface{}{"2", "3"}
    // []interface{}{"4"}
}
```

For more examples on usage please check `segmenter_test.go`.

[Back to Top](#go-strings-segmenter)

---
