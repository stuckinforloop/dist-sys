# Distributed Systems Challenge

This repository contains my solutions for the [Fly.io Distributed Systems Challenge](https://fly.io/dist-sys/) 

## Prerequisites

- Golang (v 1.23.3?)
- Make

## Steps to run the solution:

- `make init`
    - Installs maelstrom test framework
    - Runs go mod tidy
    - builds the binary

- `make run <challenge_name>`
    - e.g. `make run echo` runs the echo challenge