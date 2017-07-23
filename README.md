# Netplan

An online, free planner written in Go for school and other activities. Can
also be used as an online journal.


To start up the server for testing, go to a terminal and type
```bash
  make
  # This also works
  go run *.go
```

To start up the server for production use, change the port in main.go to 80
```go
  // NOTE: This should be 80 for production use
  PORT := 80
```

Then open a terminal and type
```bash
  sudo make
  # Or
  sudo go run *.go
```

To reset the entries in the folder, use the `reset` target in the makefile,
like this
```bash
  # You should be able to run this as a user and not need to use root
  make reset
```

Found a bug? Email me at AlexAndersonOne@gmail.com
