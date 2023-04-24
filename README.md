# Tic Tac Go!
A simple implementation of tic-tac-toe in go.
There is a random bot and a minimax box.

# Building
There are two options of building:
- With go
```
# build a bin
go build
# run directly
go run .
```
- With docker
```
# build image
docker build -t tic_tac_toe .
# run the shell
docker run -it tic_tac_toe:latest
```

# Usage
```
❯ ./tic_tac_toe -h
Usage of ./tic_tac_toe:
  -p1 string
        Player 1 type (default "human")
  -p2 string
        Player 2 type (default "bot")
```
