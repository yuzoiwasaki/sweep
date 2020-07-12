# sweep

Delete all branches except the specified branch.

## Installation

```sh
go get github.com/yuzoiwasaki/sweep
```

or

```sh
curl -o /usr/local/bin/sweep https://github.com/yuzoiwasaki/sweep/blob/master/bin/sweep
```

## Usage

Delete all branches except master.

```sh
sweep
```

Delete all branches except foo.

```sh
sweep -v foo
```

Delete all branches except master, foo.

```sh
sweep -v master -v foo
```
