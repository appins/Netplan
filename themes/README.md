## We need more themes!

If you create a theme, please submit a pull request for that theme.

To add a theme to the main.go file (`../main.go`), do the following.

Look inside, near the bottom of the code for an if statement like this

```go
if journal_url == "theme.css" {
  //...
}
```

Inside there should be a switch statement, add your theme as part of that switch statement.
