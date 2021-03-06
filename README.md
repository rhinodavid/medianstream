# Median Stream

## Overview

This application takes a series of integers. With each new integer, the app computes the median
of all the integers seen so far. Once the list is complete, the application displays
a total of these medians.

## Input

The integers must be in a text file, each on their own line. See `sample_dataset.txt`
for an example.

## Usage

Run `go run main.go <input_file.txt>`. With `sample_dataset.txt`, you should see:

```
The sum of all the medians is:
26
```

once the application completes.


## Under the hood

Median stream uses a max heap and a min heap for median maintenance. This allows new
integers to be added with a time complexity relative to the logarithm of half the total
number of inputs, and it allows the median to be found in constant time.
