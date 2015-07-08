# num

The `num` tool displays the binary, octal, decimal and hexadecimal representation of numbers.

## Installation

    go get github.com/mewmew/playground/cmd/num

## Usage

    num BIN|OCT|DEC|HEX

## Examples

```bash
$ num 0b111101101
bin: 0b111101101
oct: 0o755
dec: 493
hex: 0x1ED
```

```bash
[~]$ num 0o755
bin: 0b111101101
oct: 0o755
dec: 493
hex: 0x1ED
```

```bash
$ num 493
bin: 0b111101101
oct: 0o755
dec: 493
hex: 0x1ED
```

```bash
$ num 0x1ED
bin: 0b111101101
oct: 0o755
dec: 493
hex: 0x1ED
```

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
