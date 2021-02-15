# Daily Coding Problem: Problem #742 [Easy]

This problem was asked by Stripe.

Write a function to flatten a nested dictionary.
Namespace the keys with a period.

For example, given the following dictionary:

```
{
    "key": 3,
    "foo": {
        "a": 5,
        "bar": {
            "baz": 8
        }
    }
}
```

it should become:

```
{
    "key": 3,
    "foo.a": 5,
    "foo.bar.baz": 8
}
```

You can assume keys do not contain dots in them,
i.e. no clobbering will occur.

## Analysis

The problem statement implies, but does not explicity state,
that the dicts are string-key/numerical-value.
I'm going to proceed with that assumption.
Since I'm using JSON input text to get different
dictionary-of-dictionaries structures,
I'll be using Go `float64` as the numeric type.

### Design

I did this using Go's handy `interface{}` feature.
I used Go's `map` family of types as the dictionary.

```go
func flatten(comma string, name string, d interface{}) {
    switch d.(type) {
    case float64:
		// output name and floating point value
    case map[string]interface{}:
		// compose new '.' seperated namespaced name from key
		// call func flatten on with namespaced name,
		// and value of each key's value
// rest of func
```
When `func flatten` receives a float64,
it prints out a namespaced key and the float64.

When `func flatten` receives a map,
it iterates over key/value pairs in the map.
It composes a namespaced name from the `name` argument
passed in,
and calls itself with the namespaced name and each value
from the map.

The hardest part to get correct?
Comma placement.
The final element of any flattened map seems to be specified
to not have a comma.
That made it necessary to pass the `comma` argument to
`func flatten`.
I did not get this correct on my first try,
I had to experiment.

It turns out that the output format the problem statement
hints at is JSON, more or less,
so it's possible to flatten a dictionary,
then run the flattened dictionary's output
back through the flattening process.

#### Reconstitute flattened dictionaries

Is this transformation a bijective function on dictionaries?
That is, can you "flatten" a dictionary,
then make it 3-D again, and get back the original?

[Yes, you can](beef1.go), along with some [testing](runt2)

I'm not sure it's a bijection,
but you can flatten a nested dictionary of dictionaries,
and turn it back into an apparently identical
dictionary of dictionaries.

### Interview Analysis

This is really not an "easy" problem,
at least not in Go, and not in many other languages.
I'll agree that it's "easy" to sketch out a solution,
but one that may or may not handle corner cases correctly,
and one that doesn't do the output correctly.

Job candidates might note that the assumption that
non-flattened-dictionary keys do not contain dots
eliminates some work:
it's certainly possible to do "byte-stuffing"
or other encoding of '.' characters so that non-flattened keys
could contain dots.
This requires some experience or a good memory to know,
so it's just worth a comment.
