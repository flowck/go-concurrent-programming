# Concurrent Programming with Go

## Source

- Author: Mike Van Sickle
- URL: https://app.pluralsight.com/library/courses/concurrent-programming-go/table-of-contents

## Objectives

- How to use goroutines
- How to use `waitgroups` and `mutexes`
- How to use channels

## Outline

- Concurrency and Parallelism
- Working with `goroutines`

## Course Overview

- `goroutines`: Enable concurrent programming;
- The sync package: Allow goroutines to coordinate their work;
- Channels: A safe way for goroutines to communicate;

## Questions

- What's the difference between a Thread and a Goroutine?
  - Both have their own execution stack
  - Threads have a fixed stack (~ 1 MB)
  - Goroutines have a variable stack that starts around ~2KB
  - Threads are managed by the OS
  - Goroutines are managed by Go's runtime
