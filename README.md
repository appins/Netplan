# Online planner

An online, free planner written in Go for school and other activities. Can
also be used like an online journal.


To start up the server for testing, go to a terminal and type
```
  make
```

To start up the server for production use, change the port in main.go to 80
```go
  // NOTE: This should be 80 for production use
  PORT := 80
```

Then open a terminal and type
```
  sudo make
```
