# sweep

Delete all branches except the specified branch.

## Installation

```sh
go get github.com/yuzoiwasaki/sweep
```

or

```sh
curl -fLo /usr/local/bin/sweep \
    https://github.com/yuzoiwasaki/sweep/blob/master/bin/sweep
```

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
