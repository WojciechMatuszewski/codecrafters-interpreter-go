[![progress-banner](https://backend.codecrafters.io/progress/interpreter/5acad291-04f4-4261-9721-6792657eae02)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

# About

This is a starting point for Go solutions to the
["Build your own Interpreter" Challenge](https://app.codecrafters.io/courses/interpreter/overview).

This challenge follows the book
[Crafting Interpreters](https://craftinginterpreters.com/) by Robert Nystrom.

# Learnings

- In Go, you can lexically compare strings, like so:

  ```go
  if foo >= "a" && foo <= "z" {}
  ```

  For some reason, I thought that this was not possible.

- The `rune` type is _very_ useful for deducing what a given character is.

  ```go
  someLetter := "a"
  r := rune(someLetter[0])
  unicode.IsLetter(r)
  ```

  If you process something character by character, using `rune` type might make your life easier.

- I had great amount of trouble to understand the _recursive descent parsing_ implementation.

  - The implementation is recursive, and it is similar to _depth-first traversal_.

    - Instead of visiting pre-existing nodes, we _construct_ them by "visiting" expressions.

- Its interesting to me that, in Go, you have different naming convention for "error types" and "error variables". [Direct link](https://go.dev/wiki/Errors#naming)

  > Error types end in "Error" and error variables start with "Err" or "err".

  Here is what it means in practice:

  ```go
    var ErrNotFound = errors.New("not found")

    type MyCustomError struct {
      line int
      message string
    }
  ```

- Go generics are not as powerful as the TypeScript type system.

  - _Type inference_ is quite powerful in TypeScript. In Go, I found myself often having to specify the type explicitly.

  - There are no "default type parameters" in Go. You can't do the following:

    ```ts
    type Foo<V = number> = V;
    ```

  Having said that, I like how _type constraints_ work in Go. In my humble opinion, the syntax is much more readable.

  ```go
  type additive interface {
    string | int | float32 | float64
  }

  func add[V additive](a, b V) V {
    return a + b
  }
  ```

  There is an [experimental 'constraints' package you can use as well!](https://pkg.go.dev/golang.org/x/exp/constraints)
