# Testing plan

Ideally, the testing should include two components for each feature:

* One or more tests for each function
* One or more script tests for the same function

Principles:

1. Each public function should be tested independently of the command line commands/subcommands that are related to it.
2. Every function should be as independent as possible, to make it usable as a library.
3. Every feature should have at least one script test.
4. Both the Go test and the script tests should be written **before** the function that implements the feature.
5. There should be at least a set of script tests that illustrates the intended use of the tool. Such set should be kept separated from the many tests that check the limits and prod for bugs. We want to demonstrate that a well organized set of script tests can become effective documentation.
