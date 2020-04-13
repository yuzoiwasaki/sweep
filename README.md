# sweep

Delete all branches except the specified branch.

## Installation

```sh
go get -u github.com/yuzoiwasaki/sweep
```

## Usage

Delete all branches except master.

```sh
sweep
```

Delete all branches except foo.

```sh
sweep -exclude foo
```
