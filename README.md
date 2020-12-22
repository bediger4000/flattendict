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
that the dicts are string-key/integer-value.
I'm going to proceed with that assumption.
