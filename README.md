# Math Examples

A small utility to generate arithmetic examples for pupils

# Usage

```shell
go run .
```

Once it's run first time, it generates a `math-examples.yaml` file in current directory with generator configuration. Example:

```yaml
examplesCount: 10        # Number of exercises to generate
minBoundary: 0           # Minimal value that is calculated on each iteration
maxBoundary: 9           # Maximum value that is calculated on each iteration 
operandsCount: 2         # Number of operands in each exercise
availableOperands:       # Available operands values used in generation
  - 1
  - 2
  - 3
  - 4
  - 5
  - 6
  - 7
  - 8
  - 9
availableOperationTypes: # Available operations used in exercises
  - plus
  - minus
```