package urbase_test

import "github.com/glauberratti/boost-lab/go/07/urbase"

func ExampleRun() {
	urbase.Run()
	//for i in {1..16}; do echo "  // $i"; done

	// Output:
	// All your base ...
	// Base   2   8  10  16
	//        0   0   0   0
	//        1   1   1   1
	//       10   2   2   2
	//       11   3   3   3
	//      100   4   4   4
	//      101   5   5   5
	//      110   6   6   6
	//      111   7   7   7
	//     1000  10   8   8
	//     1001  11   9   9
	//     1010  12  10   a
	//     1011  13  11   b
	//     1100  14  12   c
	//     1101  15  13   d
	//     1110  16  14   e
	//     1111  17  15   f
}
