# branch-sweeper
Delete all branches except the specified branch.

## Usage

Delete all branches except master.

```sh
go run sweep.go
```

Delete all branches except foo.

```sh
go run sweep.go -except foo
```
