# monotonic

Monotonic is a library that leverages time package to provide always monotonically increasing clock time. And yes, it is not affected by the wall time synchronizations.

## Why?

In some time sensitive operations monotonically increasing clock time is critically important.

As characteritic, these systems mostly

* relies on the time for sorting,

* and/or relies on the first or the most recent occurences of events.

## Proved?

Check the unit test with the name `using wall time might cause time shift`.

## How?

1. Once the package is imported it sets an initilization time value `time` package's `time.Now()` return value.

2. Use the initial time as reference time to calculate the elapsed time using `time` package's `time.Since` function which is using monotonic clock time for calculations.

3. Add the elapsed time value back to the initial reference time to return monotonic clock time.

## Cons

* On some systems the monotonic clock will stop if the computer goes to sleep. On such a system, t.Sub(u) may not accurately reflect the actual time that passed between t and u.

* Adds extra latency for 2 additional arithmetic operations.

NOTE: This should not be used as a uniq id. It can return exactly the same time twice as nanoseconds. If you need a monotonically increasing id generator then give it a try to [monoton](https://github.com/mustafaturan/monoton) package.

## Usage

```
// import package
import (
    "fmt"
    "github.com/mustafaturan/monotonic"
)

// use it in a function
func somefunction() {
    t := monotonic.Now() // actual usage
    fmt.Println(t)
}
```

## Test

To run unit tests you can use the following command:
```
go test -timeout 60s
```

## Contributing

All contributors should follow [Contributing Guidelines](CONTRIBUTING.md).

## Credits

[Mustafa Turan](https://github.com/mustafaturan)

## License

Apache License 2.0

Copyright (c) 2023 Mustafa Turan

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.