# Organizational notes

## Purpose

The main purpose for writing `cliptool` is to have a suitable example to illustrate the [`testscript-cookbook`](https://github.com/datacharmer/testscript-cookbook). As such, the tool needs to be reasonably complex to help show both the basic and the advanced techniques available with testscript.

The second purpose is to document TDD techniques writing a CLI tool using both traditional testing and script-based tests.

## Build and automation

Since this is part of a testing exercise, we also need to integrate as much automation tasks as they fit into such an educational project. 

It would be nice to fit into it:

* `goreleaser` to handle automated builds and releases;
* `go-linter` to make sure from the beginning that we don't are leaving behind any technical debt.

## Dependencies

The minumim dependency for this project is a library that allows us to manage the clipboard (github.com/atotto/clipboard).
To minimize unfocused code, we could use a command and options manager such as [cobra](github.com/spf13/cobra). This has the side effect of increasing the dependencies considerably.

## Code organization

### main

The main package, which includes a testscript-based test.

### /cmd

Where the commands and options are parsed, their syntax proven, and the wanted target estalished, before passing the ball to `/ops`.

### /ops

This is where the operations are called, independent from the command line parameters and options. This organization allows the code to be used as a library, in addition to making internal testing easier.

### /defaults

Contains default and configurable data that can be changed either temporarily or permanently.