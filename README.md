# Go Out of Bounds Error Example
This repository demonstrates a common error in Go: accessing an array index beyond its bounds. The `bug.go` file contains the erroneous code, while `bugSolution.go` provides the corrected version.

The bug occurs because the loop does not check if `i+1` is a valid index within the `data` array before accessing `data[i+1]`.  This leads to a runtime panic when `i` is at the last index of the array.  The solution implements a check to prevent this.