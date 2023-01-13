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

There are two ways of assigning local variables, using the `var` and `val` keywords. The basic rule is similar to
Kotlin. But there are some differences.

1. You can't specify the type of the variable.
2. It's possible to reassign the same declared identifier with different data types.
3. Each assignment statement must end with a semicolon(`;`).

```kotlin
var a = 5;
val b = 20;

val c = 10; // [ERROR] missing token: ;
```

The `var` keyword stands for `variable` assignment. Variables defined with the `var` keyword can be reassigned. Remember
to omit the `val` and `var` keyword for reassigning already declared identifiers.

```kotlin
var a = 5;
a = 10;
a = true;

var a = 20; // [ERROR] given identifier 'a' is already declared
```

Read-only local variables are defined using the `val` keyword, which stands for `value` assignment. If you try to
reassign the value-assigned variable with another value, the interpreter will throw an error during the evaluation
process.

```kotlin
val b = 10;
b = 20; // [ERROR] can not reassign variables declared with 'val'
```

## Data types

Currently, Yail has six data types: integer, boolean, string, array, hash map and null.

As mentioned above, variables defined with the `var` keyword can be reassigned with a different data type.

```kotlin
val x = 5;
val y = true;
var z = null;
z = 10;
z = false;
```

### Basic Operations

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

1 + 2 * 3; // 7
(1 + 2) * 3; // 9
```

### Strings

A string is a sequence of characters wrapped by quotation marks(`"`). Any ASCII characters can be used as content, even
if it can not be used for identifiers.

`+` operator can be used for concatenating multiple strings to return a new string. Also, builtin function `len` can be
used for counting the length of the string, which is the number of characters including all the whitespace.

```kotlin
val a = "Hello";
val b = "World!";
val c = a + " " + b;
c; // Hello World!

len(""); // 0
len("abc123#$%^"); // 10
```

### Arrays

An array is a list of elements wrapped by brackets(`[`, `]`). Any type of data can be used as an element and arrays in
Yail can contain elements with different data types.

Each element can be accessed based on its index. If the given index is out of range, null would be returned instead of
throwing an error.

```kotlin
val arr = [1, "two", 3 + 3];

arr[0]; // 1
arr[1]; // "two"
arr[2]; // 6
arr[3]; // null
arr[-1] // null
```

Many builtin functions are supported for arrays.

- `len` returns the number of elements in the array.
- `head` returns the first element in the array without changing the array.
- `tail` returns the last element in the array without changing the array.
- `push` and `pushleft` each adds a new element as the new last or first element in the array.
- `pop` and `popleft` each removes the last or first element in the array.

```kotlin
val arr = [1, 2, 3];
len(arr); // 3
head(arr); // 1
tail(arr); // 3
arr; // [1, 2, 3]

push(arr, 10); arr; // [1, 2, 3, 10]
pushleft(arr, 20); arr; // [20, 1, 2, 3, 10]
pop(arr); arr; // [20, 1, 2, 3]
popleft(arr); arr; // [1, 2, 3]
```

### Hash Map

A hash map consists of multiple key-value pairs wrapped by curly brackets(`{`, `}`). 
Only hashable data types can be used for keys, which are strings, integers, and booleans.
Any data type can be used for values including arrays, functions, and another hash maps.

Each value can be accessed based on the corresponding key. 
If the given key does not exist, null would be returned instead of throwing an error.

```kotlin
val map = { "one": 1, true: [10, 20, 30], 100: { "a": 1, "b": 2 }, "f": func(x) { x * x }};

map["one"]; // 1
map["two"]; // null;

map[true]; // [10, 20, 30]
map[true][0]; // 10

map[100]; // { a: 1, b: 2 }
map[100]["b"]; // 2

val square = map["f"];
square(10); // 100
```

## Conditional Expressions

Basic `if` and `else` keywords are supported. When the conditions are met, multiple statements inside a selected block
are consecutively executed.

```kotlin
if (true) { val x = 10; }
x; // 10

if (5 > 10) { val a = 10; } else { val b = 15; val c = 20; }
a; // [ERROR] identifier not found: a
b; // 15
c; // 20
```

It's important to understand that `if` and `if-else` statements are actually expressions because they always return a
value.

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

## Functions

To support functional programming, all functions are first-class citizens in Yail.

This means that they are expressions, so you must assign the function to an identifier to call them.

```kotlin
val f = func(x) { x + 3; };
f(5); // 8
f(6); // 9
```

This also means that it's possible to implement higher order functions, functions that take another functions as
arguments or return a function.

```kotlin
val callTwoTimes = func(x, f) { 
    f(f(x)); 
};

callTwoTimes(2, func(x) { x * x; }); // 16
callTwoTimes(3, func(x) { x * x; }); // 81
callTwoTimes(1, func(x) { x + 10; }); // 21
```

### Scopes

When you try to use an identifier inside a function body, evaluator looks up the identifier following these steps.

1. If it's the name of a parameter or a variable declared inside the function, the evaluator uses the bound value.
2. If it's not one of them, it searches the outer scope.
3. The 2nd step is repeated until it reaches the outermost scope.

```kotlin
var i = 5; 
val useLocalVariableI = func() {
    return i;
};
val useLocalVariableInsideFunction = func() {
    val i = 15;
    return i;
};
val returnParameterI = func(i) {
    return i;
};

useLocalVariableI(); // 5
useLocalVariableInsideFunction(); // 15
returnParameterI(10); // 10
i; // 5

i = 30;

useLocalVariableI(); // 30
useLocalVariableInsideFunction(); // 15
returnParameterI(10); // 10
i; // 30
```

However, you can not reassign the variables that was declared at the outer scope from inside the function.

```kotlin
var i = 5; 
val reassignFunc = func() {
    i = 10; 
}; 
reassignFunc(); // [ERROR] identifier not found: 'i'
```

### Closures

Yail also supports closures, functions that captures the environment on the moment it was declared.

For example, `closure` function captures the environment that binds the `x` identifier to the argument `2`
when `higherOrderFunc` function is called. After that, when `closure` function is called, it uses the captured
environment instead of using the environment of the scope it is being called. Therefore, `closure` function still
interprets `x` as `2`, even though `x` is bound `1000` on the scope it is being called.

```kotlin
val higherOrderFunc = func(x) {
    func(y) { x + y };
};

val closure = higherOrderFunc(2);
val x = 1000;
closure(5); // 7
x; // 1000
```
