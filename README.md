# Diversity Index Calculator

[![Build, Test and Lint](https://github.com/at-silva/diversityindex/actions/workflows/build_test_and_lint.yml/badge.svg)](https://github.com/at-silva/diversityindex/actions/workflows/build_test_and_lint.yml)

This project provides an implementation of diversity indices, including Shannon's Diversity Index and Simpson's Diversity Index, allowing users to calculate diversity measures for both slices and maps of data.
Overview

Diversity indices are essential metrics in various fields, including ecology, biology, and information theory. They measure the diversity or richness of a dataset based on the distribution of different elements within that dataset.
Implemented Diversity Indices

    Shannon's Diversity Index
        Measures the uncertainty or entropy in a dataset's diversity.
    Simpson's Diversity Index
        Measures the probability that two elements randomly selected from a sample will belong to the same group or category.

## Usage

### Installation

To use this package, simply install it using go get:

```bash
go get github.com/at-silva/diversityindex
```

### Example Usage

#### Calculating Diversity Indices for Slices

```go
// Calculate Simpson's Diversity Index for a uint64 slice
simpsonsIndex, _ := diversityindex.CalculateSlice(diversityindex.DiversityIndexSimpsons, []uint64{1, 1, 1, 2, 2})

// Calculate Shannon's Diversity Index for a uint64 slice
shannonsIndex, _ := diversityindex.CalculateSlice(diversityindex.DiversityIndexShannons, []uint64{1, 1, 1, 2, 2})
```
#### Calculating Diversity Indices for Maps

```go
// Calculate Simpson's Diversity Index for a map with integer keys
simpsonsMapIndex, _ := diversityindex.CalculateMap(diversityindex.DiversityIndexSimpsons, map[int]uint{1: 3, 2: 2})

// Calculate Shannon's Diversity Index for a map with string keys
shannonsMapIndex, _ := diversityindex.CalculateMap(diversityindex.DiversityIndexShannons, map[string]uint{"1": 40, "2": 20, "3": 15})
```

## Testing

The package includes comprehensive test coverage to ensure the correctness of diversity index calculations. Run tests using:

```bash
go test github.com/at-silva/diversityindex
```

## Contribution

Contributions to this project are welcome! To contribute, follow these steps:

- Fork the repository
 - Create a new branch (git checkout -b feature/awesome-feature)
 - Make your changes
 - Commit your changes (git commit -am 'Add new feature')
 - Push to the branch (git push origin feature/awesome-feature)
 - Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.