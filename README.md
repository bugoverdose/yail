# YAIL: Yet Another Interpreted Language

## Variables

### identifier format

Identifiers, or name of the variables, must follow the following rules.

1. The identifier should be a string with at least one character.
2. The first character should be an alphabet(`a`~`Z`) or an underscore(`_`).
3. The rest of the characters should be alphabets(`a`~`Z`), numbers(`0`~`9`), or an underscore(`_`).
4. The identifier should not be one of the reserved keywords used by Yail(`var`, `val`, `return`, etc).

The **lexer** will throw an error if it finds an identifer not following these rules.

```kotlin
// allowed
var a = 1;
var _ = 2;
var a_b_2 = 3;

// not allowed
var 1a = 1; // starts with a number
var a b = 2; // whitespace included
var val = 3; // val is a keyword
var 변수 = 4; // only ASCII allowed
```

### variable binding

There are two ways of assigning local variables, using the `var` and `val` keywords. The basic rule is similar to Kotlin. But there are a few differences.

1. You can't specify the type of the variable. Just like JavaScript, the type is always inferred.
2. Each assignment statement must end with a semicolon(`;`).

```kotlin
var a = 5;
val b = 10;
```

The `var` keyword stands for `variable` assignment. Variables defined with the `var` keyword can be reassigned.

```kotlin
var a = 5;
a = 10;
```

Read-only local variables are defined using the `val` keyword which stands for `value` assignment. The **parser** will throw an error if you try to reassign the value-assigned variable with another value.

```kotlin
val b = 10;
// b = 20; // error!
```
