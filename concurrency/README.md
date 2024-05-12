# Concurrency

#### Look into concurrency pattern codes in following order - 

Taken from - https://youtu.be/f6kdp27TYZs?t=866 (Go Concurrency Patterns Rob Pike)
1. Generator: function that returns a channel: generator.go
2. Channels as a handle on service: channel_service.go
3. Multiplexing/ fanIn: fan_in.go
4. Restoring sequence: restoring_sequence.go
    send a channel on a channel, making goroutine wait it turn
5. Fan-in function using `select` statement in go for channels: fan_in_using_select.go
6. Timeout using select: timeout_using_select.go

#### Using Mutex Examples 
1. Getting rid of Data race condition while doing read/write on shared data using mutex: data_race_mutex.go

#### Good Reads
Why Go Programming:
- https://www.ardanlabs.com/blog/2013/05/why-go-programming.html

Scheduling in Go:
- https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html
- https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html
- https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html

