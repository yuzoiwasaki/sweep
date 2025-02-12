# sweep

Delete all branches except the specified branch.

## Installation

```sh
go install github.com/yuzoiwasaki/sweep@latest
```

or

```sh
curl -fLo /usr/local/bin/sweep \
    https://github.com/yuzoiwasaki/sweep/blob/master/bin/sweep
```
Please use the appropriate binary for your OS

## Usage

Delete all branches except main.

```sh
sweep
```

Delete all branches except foo.

```sh
sweep -v foo
```

Delete all branches except main, foo.

```sh
sweep -v main -v foo
```
