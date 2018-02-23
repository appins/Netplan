[![AppIns](https://imgur.com/W5HbyhH.png)](https://github.com/appins)
# Netplan

An online, free planner written in Go for school and other activities. Can
also be used as an online journal or diary.

Notice that this project uses caching, so killing the program at any time could
result in lost data for all users that recently accessed their journals.

Instead of using `^C` to kill it, use `sudo make killit`


To start up the server for testing, go to a terminal and type
```bash
  make
  # This also works
  go run *.go
```

Then go to http://localhost:8080 in a browser

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

Then go to http://localhost in a browser

To reset the entries in the folder, use the `reset` target in the makefile,
like this
```bash
  # You might need to use sudo make reset if you used sudo make to run the server
  make reset
```

If you want to log what is going on, use the `log` target, like this:
```bash
  make log
  # OR
  sudo make log
```

Found a bug? Email me at AlexAndersonOne@gmail.com
