# Concurrent Requests Playground

This repo hosts code related to this [blog post](https://jellis18.github.io/post/2023-03-04-concurrency-in-languages/).

## Running code

To run locally you will need to have Go and python installed. You will also need `curl` and `jq` installed. You can
also run in github codespaces or locally using the devcontainer.

### Go

```bash
go build -o get-todos get_todos.go
./get-todos 200
```

### Python

You will need httpx installed. You can install it with `pip install httpx`.

```bash
python3 get_todos.py 200
```

### Bash

```bash
bash get_todos.sh 200
```
