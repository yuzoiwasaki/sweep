# sweep

Delete all branches except the specified branch.

## Requirements

- Unix-like operating system (Linux or macOS)
- Git

## Installation

```sh
go install github.com/yuzoiwasaki/sweep@latest
```

or

```sh
# For Linux (amd64)
curl -fLo /usr/local/bin/sweep \
    https://github.com/yuzoiwasaki/sweep/blob/master/bin/sweep

# For macOS Intel
curl -fLo /usr/local/bin/sweep \
    https://github.com/yuzoiwasaki/sweep/blob/master/bin/sweep-darwin-amd64

# For macOS Apple Silicon
curl -fLo /usr/local/bin/sweep \
    https://github.com/yuzoiwasaki/sweep/blob/master/bin/sweep-darwin-arm64
```

## Usage

Delete all branches except main.

```sh
sweep
```

Delete all branches except foo.

```sh
sweep -v foo
````