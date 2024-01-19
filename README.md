# Math Examples

A small utility to generate arithmetic examples for pupils

# Installation

```shell
go install github.com/mih-kopylov/math-examples@latest
```

Or download the [latest binary](https://github.com/mih-kopylov/math-examples/releases/latest).

# Usage

```shell
math-examples -p Имя
```

Or if running from sources,

```shell
go run . -p Имя
```

# Configuration

Once it's run first time, it generates a `math-examples.yaml` file in current directory with generator configuration. 
Example:

```yaml
profiles:
  Имя:                               # User profile name, used to differentiate users configuration 
    examplesCount: 10                # Number of exercises to generate
    minBoundary: 0                   # Minimal value that is calculated on each iteration
    maxBoundary: 100                 # Maximum value that is calculated on each iteration 
    operandsCount: 2                 # Number of operands in each exercise
    showCorrectAnswerAfter: each     # Use 'each' to show a correct answer after each exercise or 'all' to show summary after all exercises
    availableOperands:               # Available operands values used in plus and minus generation
      - 1:100
    availableMultiplicationOperands: # Available operands values used for multiplication and division
      - 1:9
    availableOperationTypes:         # Available operations used in exercises
      - plus
      - minus
      - multiply
      - divide
```
