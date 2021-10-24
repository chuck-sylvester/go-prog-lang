# go-prog-lang
The Go Programming Language (Donovan &amp; Kernighan)

## CPS Notes
Go code is organized into packages, which are similar to libraries or modules in other languages.

Package `main` is special. It defines a standalone executable program, not a library. Within package `main` the *function* `main` is also special -- it's where execution of the program begins.

> Whatever `main` does is what the program does.

We must tell the compiler what packages are needed by the source file; that's the role of the `import` declaration that follows the `package` declaration.

> You must import exactly the pacages you need.

