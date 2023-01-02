# YAIL: Yet Another Interpreted Language

YAIL is a simple interpreted language built with Go.

## Variables

### Identifier format

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

### Variable binding

There are two ways of assigning local variables, using the `var` and `val` keywords. The basic rule is similar to Kotlin. But there are some differences.

1. You can't specify the type of the variable. 
2. It's possible to reassign the same declared identifier with different data types.
3. Each assignment statement must end with a semicolon(`;`).

```kotlin
var a = 5;
val b = 20;

val c = 10; // [ERROR] missing token: ;
```

The `var` keyword stands for `variable` assignment. Variables defined with the `var` keyword can be reassigned. Remember to omit the `val` and `var` keyword for reassigning already declared identifiers.

```kotlin
var a = 5;
a = 10;
a = true;

var a = 20; // [ERROR] given identifier 'a' is already declared
```

Read-only local variables are defined using the `val` keyword, which stands for `value` assignment. If you try to reassign the value-assigned variable with another value, the interpreter will throw an error during the evaluation process.

```kotlin
val b = 10;
b = 20; // [ERROR] can not reassign variables declared with 'val'
```

## Data types

Currently, Yail has only three data types: integer, boolean, and null. As mentioned above, variables defined with the `var` keyword can be reassigned with a different data type.

```kotlin
val x = 5;
val y = true;
var z = null;
z = 10;
z = false;
```

## Operations

All the following operations are supported.

- basic arithmetic operations(`+`, `-`, `*`, `/`, `%`) for integers
- negative prefix(`-`) for integers
- basic comparison operations(`==`, `!=`, `<=`, `>=`, `<`, `>`)
- not prefix(`!`) for reversing a boolean
- grouping(`()`) for changing the priority of operations

```kotlin
5 + 5; // 10
5 - 5; // 0
5 * 5; // 25
5 / 5; // 1
5 % 5; // 0

5 == 5; // true
5 != 5 // false
5 <= 5 // true
5 >= 5 // true
10 > 5 // true
10 < 5 // false

!false; // true

(1 + 2) * -3; // -9
```

## Conditional Expressions

Basic `if` and `else` keywords are supported. When the conditions are met, multiple statements inside a selected block are consecutively executed.

```kotlin
if (true) { val x = 10; }
x; // 10

if (5 > 10) { val a = 10; } else { val b = 15; val c = 20; }
a; // [ERROR] identifier not found: a
b; // 15
c; // 20
```

It's important to understand that `if` and `if-else` statements are actually expressions because they always return a value.

1. If a block is executed, the value of the last expression is returned.
2. If the executed block ends with a statement, it returns `null`.
3. If no block is executed, the if expression returns `null`

```kotlin
val x = if (false) { 10 } else { 20 };
x; // 20

var y = if (false) { 10 };
y; // null

val z = if (true) { y = 15; };
z; // null
```
