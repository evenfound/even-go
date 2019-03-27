package tool

func ExampleTR0() {
	TR("AAA")
	// Output: AAA
}

func ExampleTR1() {
	TR("BBB", 100)
	// Output: BBB 100 (int)
}

func ExampleTR2() {
	TR("CCC", 200, true)
	// Output: CCC 200 (int) true (bool)
}

func ExampleTR3() {
	TR("DDD", 300, false, 3.14159)
	// Output: DDD 300 (int) false (bool) 3.14159 (float64)
}
