# Algorithm

## Introduction

This is a git repo to store the implementation of some algorithms that I learned. I will mainly use golang to write these codes.

## Content

### FPTree

An algorithm for frequent pattern mining.

Ref: *Mining Frequent Patterns without Candidate Generation*, Jiawei Han, Jian Pei, and Yiwen Yin, *MOD 2000*

Module: ```./FPTree```

Input: a two-dimensional byte with strings, each line corresponds to a transaction record, and each string with an item.

Output: a two-dimensional byte representing patterns.

### LIS

The LIS (Longest Increasing Subsequence, 最长递增子序列) Algorithm is a classical algorithm to find a longest subsequence of a given input sequence that the values of it are increasing monotonically. 

Here we implement this algorithm with the O(nlogn) time complex.

## Run algorithm

Add the algorithm that you would like to run in the main function in ```main.go```. The algorithm handlers are named as ```<algorithm_name>Test()```.
