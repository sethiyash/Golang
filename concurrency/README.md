# Concurrency


Looks into concurrency pattern codes in flowing order - 

Taken from - https://youtu.be/f6kdp27TYZs?t=866 (Go Concurrency Patterns Rob Pike)
1. Generator: function that returns a channel: generator.go
2. Channels as a handle on service: channel_service.go
3. Multiplexing/ fanIn: fan_in.go
4. Restoring sequence: restoring_sequence.go
    send a channel on a channel, making goroutine wait it turn
    