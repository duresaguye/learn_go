# Project: Concurrent URL Checker

Check the status of multiple websites at the same time.

## How it Works

1.  We have a list of URLs.
2.  Instead of checking them one by one (which would be slow), we launch a **goroutine** for each URL.
3.  Each goroutine tries to access the website and sends the result to a **channel**.
4.  The main function waits to receive a message for every URL we checked.

## Why is this cool?

If checking one site takes 1 second, checking 5 sites sequentially takes **5 seconds**.
Checking 5 sites concurrently takes **~1 second** (the time of the slowest one).

## How to Run

```bash
go run main.go
```
