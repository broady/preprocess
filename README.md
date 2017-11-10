# preprocess

A simple line-by-line preprocessor.

## Command-line

### Installation

```
go get -u github.com/broady/preprocess
```

### Process a file

```
preprocess < in > out
```

### Usage/Options

```
Usage of preprocess:
  -prefix string
      Prefix for pragma. (default "//#")
```

## Directives

### `if`

Takes a flag as an argument. Flags are set via command-line arguments (i.e., `preprocess foo`).

```
//# if foo
hello
//# end

also handles negations:
//# if !foo
foo flag is unset
//# end
```

### `omit`

Omits the entire line.

```
won't be in the output //# omit
will be in output
```

### `omit if`

Takes a flag as an argument. Omits the entire line if the flag is set.

```
foo flag is set   //# omit if !foo
foo flag is unset //# omit if foo
```

Trailing spaces are omitted if the line is printed.

### `def`

Takes a name as an argument to define a template. The template is defined as the lines up to the `enddef` directive.

```
//# def newclient
ctx := context.Background()
client, err := foo.NewClient(ctx)
//# enddef
```

Templates may include other directives, like `omit` and `if`.

### `replace`

Takes a template name as an argument. Replaces the line with the given template.

```
//# def foo
foo
//# enddef

this line will be replaced by the foo template //# replace newclient
```

