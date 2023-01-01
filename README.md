# YAIL: Yet Another Interpreted Language

## Variables

### identifier format

The name of the variables, or `identifiers`, must follow the following rules.

1. The identifier should be a string with at least one character.
2. The first character should be an alphabet(`a`-`Z`) or an underscore(`_`).
3. The rest of the characters should be alphabets(`a`-`Z`), numbers(`0`-`9`), or an underscore(`_`).
4. The reserved keywords used by Yail can not be used as an identifier(`var`, `return`, etc).

```kotlin
// allowed
var a = 1;
var _ = true;
var a_b_2 = 3;

// not allowed
var 1a = 1; // starts with a number
var a b = 2; // whitespace included
var val = 3; // val is a keyword
var 변수 = 4; // non-ASCII characters are not allowed
```

### variable binding

There are two ways of assigning local variables, using the `var` and `val` keywords. The basic rule is similar to Kotlin. But there are some differences.

1. You can't specify the type of the variable. 
2. It's possible to reassign the same declared identifier with different data types.
3. Each assignment statement must end with a semicolon(`;`).

```kotlin
var a = 5;
val b = 20;

val c = 10
// Failed to execute the given source code for following reasons.
//        [ERROR] missing token: ;
```

The `var` keyword stands for `variable` assignment. Variables defined with the `var` keyword can be reassigned. Remember to omit the `val` and `var` keyword for reassigning already declared identifiers.

```kotlin
var a = 5;
a = 10;
a = true;

var a = 20; // ERROR: given identifier 'a' is already declared
```

Read-only local variables are defined using the `val` keyword, which stands for `value` assignment. If you try to reassign the value-assigned variable with another value, the interpreter will throw an error during the evaluation process.

```kotlin
val b = 10;
b = 20; // ERROR: can not reassign variables declared with 'val'
```
