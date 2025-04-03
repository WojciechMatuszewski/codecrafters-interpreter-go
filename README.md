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
