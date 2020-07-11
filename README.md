# sweep

Delete all branches except the specified branch.

## Installation

```sh
go get github.com/yuzoiwasaki/sweep
```

or

```sh
git clone https://github.com/yuzoiwasaki/sweep
mv sweep/bin/sweep /usr/local/bin/
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

Delete all branches except master, foo.

```sh
sweep -exclude master -exclude foo
```
