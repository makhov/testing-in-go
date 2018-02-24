// GOOKTEST OMIT
➜ go test ./pkg
ok  	_/Users/amakhov/Documents/testing/testing_in_go/pkg	0.006s
// END OMIT

// GOOKVTEST OMIT
➜ go test -v ./pkg
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestSum2
--- PASS: TestSum2 (0.00s)
=== RUN   TestSum3
--- PASS: TestSum3 (0.00s)
=== RUN   TestSum4
--- PASS: TestSum4 (0.00s)
=== RUN   TestSum5
--- PASS: TestSum5 (0.00s)
=== RUN   TestSum6
--- PASS: TestSum6 (0.00s)
=== RUN   TestSum7
--- PASS: TestSum7 (0.00s)
PASS
ok 	_/Users/amakhov/Documents/testing/testing_in_go/pkg	0.006s
// END OMIT

// GOFAILVTEST OMIT
➜ go test -v ./pkg
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestSum2
--- FAIL: TestSum2 (0.00s)
=== RUN   TestSum3
--- PASS: TestSum3 (0.00s)
=== RUN   TestSum4
--- PASS: TestSum4 (0.00s)
=== RUN   TestSum5
--- PASS: TestSum5 (0.00s)
=== RUN   TestSum6
--- PASS: TestSum6 (0.00s)
=== RUN   TestSum7
--- PASS: TestSum7 (0.00s)
FAIL
exit status 1
FAIL 	_/Users/amakhov/Documents/testing/testing_in_go/pkg	0.005s
// END OMIT

// RICHGOFAILVTEST OMIT
➜ richgo test -v ./pkg
START| Sum
PASS | Sum (0.00s)
START| Sum2
FAIL | --- FAIL: TestSum2 (0.00s)
START| Sum3
PASS | Sum3 (0.00s)
START| Sum4
PASS | Sum4 (0.00s)
START| Sum5
PASS | Sum5 (0.00s)
START| Sum6
PASS | Sum6 (0.00s)
START| Sum7
PASS | Sum7 (0.00s)
     | exit status 1
FAIL 	_/Users/amakhov/Documents/testing/testing_in_go/pkg	0.005s
// END OMIT

// GOTESTLIST OMIT
➜ go test ./pkg -list .+
TestSum
TestSum2
TestSum3
TestSum4
TestSum5
TestSum6
TestSum7
ok      _/Users/amakhov/Documents/testing/testing_in_go/pkg     0.005s
// END OMIT

// JUNIT OMIT
$ go test -v 2>&1 | go-junit-report > report.xml
// END OMIT

// SHORT OMIT
if testing.Short() {
	t.Skip(…)
}
// END OMIT