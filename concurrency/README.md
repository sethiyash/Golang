

## Contains
- [Achieving concurrency in Go](#concurrency)
- [Anatomy of goroutines in Go](#goroutine)

<a name="concurrency"></a>
## Achieving concurrency in Go

*If I had to choose one great feature of Go, then it has to be in-built concurrency model. Not only it supports concurrency but makes it better. Go Concurrency Model (goroutines) to concurrency is what Docker is to virtualization.*

## ☛ What is concurrency?
In computer programming, concurrency is ability of a computer to deal with multiple things at once. For general example, if you are surfing internet in a browser, there might be lot of things happening at once. In a particular situation, you might be downloading some files while listening to some music on the page that you are scrolling, at the same time. Hence a browser needs to deal with lot of things at once. If browser couldn’t deal with them at once, you need to wait until all download finishes and then you can start browsing internet again. That would be frustrating.
A general purpose PC might have just one CPU core which does all the processing and computation. A CPU core can handle one thing at a time. When we talk about concurrency, we are doing one thing at a time but we divide CPU Time among things that need to be processed. Hence we get a sensation of multiple things happening at the same time in in reality, only one thing is happening at a time.
Let’s look at the diagram how a CPU manage web browser manages things like in the example we talked about.

![alt text](https://miro.medium.com/v2/resize:fit:1400/format:webp/0*X0pg_FAWAv93kpii.jpg)

So from above diagram, you can see that a single core processor pretty much divide the workload based on the priority of each task, for example, while page scrolling, listening to music may have a low priority, hence sometimes your music stops because of low internet speed but you can still scroll the page.

## ☛ What is parallelism?
But then the question arises, how about if my CPU has multiple cores? If a processor has multiple processors, then it called as multi-core processor. You might have heard of this term while purchasing a laptop, pc or a smart phone. A multi-core processor has ability to handle multiple things at once.
In previous web browsing example, our single core processor had to divide CPU time among different things. With multi-core processor, we can run separate things at once in different cores. Let’s evaluate that using below diagram.

![alt text](https://miro.medium.com/v2/resize:fit:1400/format:webp/0*8NBpRcm6HQ4tfxgs.jpg)

**Concept of running different things in parallel known as parallelism.** When our CPU have multiple cores, we can use different CPU cores to do multiple things at once. Hence we could say that we might be able to finish a job (consist lot of things) very quickly, but that’s not the case. I will come back to this point.

## ☛ concurrency vs parallelism
Go recommends to use goroutines on one core only but we can modify the Go program to run goroutines on different processor cores. For now, think goroutines as Go functions, because they are, but there is more to it.
There are several differences between concurrency and parallelism. While **concurrency is dealing with multiple things at once, parallelism is doing multiple things at once.** Parallelism is not always beneficial over concurrency, we will learn this in upcoming lessons.
At this point, there might be a lot of questions flying in your head and you might have got the idea of concurrency but you might be wondering how Go implements it and how you can use it. To understand Go’s architecture of concurrency and how to use it in your code, as well as when to use it in your application architecture, we need to understand what computer processes are.

## ☛ What is a computer process?
When you write a computer program in languages like C, java or Go, it is just a text file. But as your computer only understands binary instructions which are composed of 0s and 1s, you need to compile that code to machine language. This is where compiler comes in. In scripting languages like python and javascript, the interpreter does the same thing.
When a compiled program is sent to OS to handle, OS allocates different things like memory address space (where process’s heap and stacks will be located), a program counter, a PID (process id) and other very crucial things. A process has at least one thread known as primary thread, while primary thread can create multiple other threads. When the primary thread is done with its execution, process exits.
So we understood that process is a container that has compiled the code, memory, different OS resources and other things which can be provided to threads. In nutshell, a process is a program in the memory. But then what are threads, what’s their job?

## ☛ What is a thread?
A thread is a light-weight process inside a process. A thread is the actual executor of a piece of code. A thread has access to memory provided by the process, OS resources, and other things.
While executing code, thread store variables (data) inside the memory region is known as stack which scratch space where variables hold temporary space. A stack is created at runtime and is normally of a fixed size, preferably 1-2 MB. While the stack of a thread can be used by only that thread and will not be shared with other thread. A heap is a property of a process and it is available to use by any thread. Heap is a shared memory space where data from one thread can be access by other threads as well.
Now we got a general idea of the process and a thread. But what is their use?
When you start a web browser, there must be some code which instructs OS to do something. This means we are creating a process. That process may ask OS to create another process for a new tab. When a browser tab opens and you are doing your normal everyday things, that tab process will start creating different threads for different activities (like page scroll, downloading, listening to music, etc.) as we seen in previous diagrams before.
Below is the screen-grab of Chrome Browser application on macOS platform.

![](https://miro.medium.com/v2/resize:fit:1400/format:webp/0*hKiunytiOCFFZHFU.png)

Above screen-grab shows that Google Chrome browser is using different processes for opened tabs and internal services. As each process has least one thread, we can see that a Google Chrome process, in this case, has more than 10 threads.
In previous topics, we talked about dealing with multiple things or doing multiple things. A thing here is an activity performed by a thread. Hence when multiple things happen in concurrency or parallelism mode, there are multiple threads running in series or parallel, AKA multi-theading.

*In multi-threading, where multiple threads are spawned in a process, a thread with memory leak can exhaust resources of other thread and make process irresponsive. You might have seen this many times while using browser or any other program. You may have used activity monitor or task manager to see process which is irresponsive and kill it.*

## ☛ Thread scheduling
When multiple threads are running in series or in parallel, as multiple threads might share some data, threads need to work in coordination so that only one thread can access a particular data at one time. Execution of multiple threads in some order is called scheduling. Os threads are scheduled by the kernel and some threads are managed by runtime environment of the programming language, like JRE. When multiple threads trying to access the same data at the same time which cause data to be changed or results into unexpected outcome, then race condition occurs.
*While designing concurrent Go programs, we need to look for race conditions which we will talk about in upcoming lessons.*

![](https://miro.medium.com/v2/resize:fit:1400/format:webp/0*bZen0r9lH4jsFFCx.png)

## ☛ Concurrency in Go
Finally, we reached to a point where we will talk about how Go implements concurrency. Traditional languages like java has a thread class which can be used to create multiple threads in the current process. Since Go does not have traditional OOP syntaxes, it provides go keyword to create goroutines. When go keyword is placed before a function call, it becomes goroutines.
We will talk about goroutines in next lesson but in nutshell, goroutines behave like threads but technically; it is an abstraction over threads.
When we run a Go program, Go runtime will create few threads on a core on which all the goroutines are multiplexed (spawned). At any point in time, one thread will be executing one goroutine and if that goroutine is blocked, then it will be swapped out for another goroutine that will execute on that thread instead. This is like thread scheduling but handled by Go runtime and this is much faster.
It is advised in most of the cases, to run all your goroutines on one core but if you need to divide goroutines among available CPU cores of your system, you can use GOMAXPROCS environment variable or call to runtime using function runtime.GOMAXPROCS(n) where n is the number of cores to use. But you may sometime feel that setting GOMAXPROCS > 1 is making your program slower. It truly depends on the nature of your program but you can find a solution or explanation of your problem on the internet. In practical terms, programs that spend more time communicating on channels than doing computation will experience performance degradation when using multiple cores, OS threads, and processes.
Go has an M:N scheduler that can also utilize multiple processors. At any time, M goroutines need to be scheduled on N OS threads that run on at most on GOMAXPROCS numbers of processors. At any time, at most only one thread is allowed to run per core. But scheduler can create more threads if required, but that rarely happens. If your program doesn’t start any additional goroutines, it will naturally run in only one thread no matter how many cores you allow it to use.

## ☛ threads vs goroutines
Since there is an obvious difference between threads and goroutines as we have seen earlier but below differences will shed some light on why threads are more expensive than goroutines and why goroutines is a key solution for achieving the highest level of concurrency in your application.

|thread|goroutine|
|------|---------|
| OS threads are managed by kernal and has hardware dependencies. | goroutines are managed by go runtime and has no hardware dependencies. |
| OS threads generally have fixed stack size of 1-2MB| goroutines typically have 8KB (2KB since Go 1.4) of stack size in newer versions of go |
| Stack size is determined during compile time and can not grow | Stack size of go is managed in run-time and can grow up to 1GB which is possible by allocating and freeing heap storage |
| There is no easy communication medium between threads. There is huge latency between inter-thread communication. | goroutine use `channels` to communicate with other goroutines with low latency ([read more](https://blog.twitch.tv/gos-march-to-low-latency-gc-a6fa96f06eb7)). |
| Threads have identity. There is TID which identifies each thread in a process. | goroutine do not have any identity. go implemented this because go does not have TLS([Thread Local Storage](https://msdn.microsoft.com/en-us/library/windows/desktop/ms686749(v=vs.85).aspx)). |
| Threads have significant setup and teardown cost as a thread has to request lots of resources from OS and return once it's done. | goroutines are created and destoryed by the go's runtime. These operations are very cheap compared to threads as go runtime already maintain pool of threads for goroutines. In this case OS is not aware of goroutines. |
| Threads are preemptively scheduled ([read here](https://stackoverflow.com/questions/4147221/preemptive-threads-vs-non-preemptive-threads)). Switching cost between threads is high as scheduler needs to save/restore more than 50 registers and states. This can be quite significant when there is rapid switching between threads. | goroutines are coopertively scheduled ([read more](https://stackoverflow.com/questions/37469995/goroutines-are-cooperatively-scheduled-does-that-mean-that-goroutines-that-don)). When a goroutine switch occurs, only 3 registers need to be saved or restored. |

Above are the few important differences but if you dive deep, you will find out the amazing world of Go’s concurrency model. To highlight some of the power points of Go’s concurrency strength, imagine you have a web server where you are handling 1000 requests per minute. If you had to run each request concurrently, that means you need to create 1000 threads or divided them under different processes. That’s how Apache server manages incoming requests (read here). If an OS thread consumes 1MB stack size per thread, that means you will exhaust 1GB of RAM for that traffic. Apache provides ThreadStackSize directive to manage stack size per thread but still, you have no idea if you run into a problem because of this.

In the case of goroutines, since stack size can grow dynamically, you can spawn 1000 goroutines without a problem. As a goroutine starts with 8KB (2KB since Go 1.4) of stack space, most of them generally don’t grow bigger than that. But if there is a recursive operation that demands more memory, Go can increase stack size up to 1GB which I hardly think will ever happen except for {} which is obviously a bug.

Also, rapid switching between goroutines is possible and more efficient compared to threads as we saw earlier. Since one goroutine is running on one thread at a time and goroutines are cooperatively scheduled, another goroutine is not scheduled until current goroutine is blocked. If any Goroutine in that thread blocks say waiting for user input, then another goroutine is scheduled in its place. goroutine can block on one of the following conditions

- network input
- sleeping
- channel operation
- blocking on primitives in the sync package

If the goroutine does not block on one of these conditions, it can starve the thread on which it was multiplexed, killing other goroutines in the process. While there are some remedies, but if it does, then it is considered as bad programming.

**Channels** will play a great role while working with goroutines as a medium to share data in between them which we will learn in upcoming lessons. This will prevent race conditions and inappropriate access to shared data between them as opposed to accessing the shared memory in case of threads.



<a name="goroutine"></a>
## Anatomy of goroutines in Go -Concurrency in Go

* goroutine is a lightweight execution thread running in the background. goroutines are key ingredients to achieve concurrency in Go.*

As goroutines are lightweight compared to OS threads, it is very common for a Go application to have thousands of goroutines running concurrently. Concurrency can speed up application significantly as well as help us write code with separation of concerns (SoC).

## ☛ What is a goroutine?
We understood in theory that how goroutine works, but in code, what is it? Well, a goroutine is simply a function or method that is running in background concurrently with other goroutines. It’s not a function or method definition that determines if it is a goroutine, it is determined by how we call it.
Go provides a special keyword go to create a goroutine. When we call a function or a method with go prefix, that function or method executes in a goroutine. Let’s see a simple example.

```
package main

import "fmt"

func printHello(){
    fmt.Println("Hello World!")
}

func main() {
    fmt.Println("main execution started")

    printHello()

    fmt.Println("main execution stopped")
}
```

In the above program, we created a function printHello which prints Hello World! to the console. In main function, we called printHello() like a normal function call and we got the desired result.
Now let’s create **goroutine** from the same printHello function.

```
package main

import "fmt"

func printHello(){
    fmt.Println("Hello World!")
}

func main() {
    fmt.Println("main execution started")

    go printHello()

    fmt.Println("main execution stopped")
}
```

Well, as per goroutine syntax, we prefixed function call with go keyword and program executed well. It yielded the following result.

```
main execution started
main execution stopped
```

It is a bit strange that *Hello World* did not get printed. So what happened?

goroutines always run in the background. When a goroutine is executed, here, Go does not block the program execution, unlike normal function call as we have seen in the previous example. Instead, the control is returned immediately to the next line of code and any returned value from goroutine is ignored. **But even then, why we can’t see the function output?**

By default, every Go standalone application creates one goroutine. This is known as the **main goroutine** that the **main** function operates on. In the above case, the main goroutine spawns another goroutine of *printHello* function, let’s call it *printHello goroutine*. Hence when we execute the above program, there are two goroutines running concurrently. As we saw in the earlier program, goroutines are scheduled cooperatively.

Hence when the main goroutine starts executing, go scheduler dot not pass control to the *printHello* goroutine until the main goroutine does not execute completely. Unfortunately, when the main goroutine is done with execution, the program terminates immediately and scheduler did not get time to schedule *printHello* goroutine. But as we know from other lessons, using blocking condition, we can pass control to other goroutines manually AKA telling the scheduler to schedule other available goroutines. Let’s use **time.Sleep()** call to do it.

```
package main

import "fmt"

func printHello(){
    fmt.Println("Hello World!")
}

func main() {
    fmt.Println("main execution started")

    go printHello()

    // schedule another goroutine
    time.Sleep(10 * time.Millisecond)
    fmt.Println("main execution stopped")
}
```

We have modified program in such a way that before **main** goroutine pass control to the last line of code, we pass control to **printHello** goroutine using time.Sleep(10 * time.Millisecond) call. In this case, the main goroutine sleeps for 10 milli-seconds and won’t be scheduled again for another 10 milliseconds. Once **printHello** goroutine executes, it prints ‘Hello World!’ to the console and terminates, then the main goroutine is scheduled again (after 10 milliseconds) to execute the last line of code **where stack pointer is**. Hence the above program yields the following result.

```
main execution started
Hello World!
main execution stopped
```

If we add a sleep call inside the function which will tell goroutine to schedule another available goroutine, in this case, the main goroutine. But from the last lesson, we learned that only non-sleeping goroutines are considered for scheduling, main won’t be scheduled again for 10 milli-seconds while it’s sleeping.

Hence the main goroutine will print ‘main execution started’, spawning **printHello** goroutine but still actively running, then sleeping for 10 milli-seconds and passing control to **printHello** goroutine. printHello goroutine then will sleep for 1 milli-second telling the scheduler to schedule another goroutine but since there isn’t any available, waking up after 1 milli-second and printing **‘Hello World!’** and then dying. Then the main goroutine will wake up after a few milliseconds, printing ‘main execution stopped’ and exiting the program.

```
package main

import "fmt"

func printHello(){
    time.Sleep(time.Millisecond)
    fmt.Println("Hello World!")
}

func main() {
    fmt.Println("main execution started")

    go printHello()

    // schedule another goroutine
    time.Sleep(10 * time.Millisecond)
    fmt.Println("main execution stopped")
}
```

Above program will still print the same result

```
main execution started
Hello World!
main execution stopped
```

What if, instead of 1 milli-second, printHello goroutine sleeps for 15 milliseconds.

```
package main

import "fmt"

func printHello(){
    time.Sleep(15 * time.Millisecond)
    fmt.Println("Hello World!")
}

func main() {
    fmt.Println("main execution started")

    go printHello()

    // schedule another goroutine
    time.Sleep(10 * time.Millisecond)
    fmt.Println("main execution stopped")
}
```

In that case, the main goroutine will be available to schedule for scheduler before printHello goroutine wakes up, which will also terminate the program immediately before scheduler had time to schedule **printHello** goroutine again. Hence it will yield below program

```
main execution started
main execution stopped
```

## ☛ Working with multiple goroutines
As I said earlier, you can create as many goroutines as you can. Let’s define two simple functions, one prints characters of the string and another prints digit of the integer slice.

```
package main

import "fmt"

func getChars(s string){
    for _, c := range s {
        fmt.Printf("%s ", c)
    }
}

func getDigits(s []int){
    for _, d := range s {
        fmt.Printf("%d ", d)
    }
}

func main() {
    fmt.Println("main execution started")

    go getChars("Hello")

    go getDigits([]int{1,2,3,4,5})

    time.Sleep(time.Millisecond)

    fmt.Println("main execution stopped")
}
```

In the above program, we are creating 2 goroutines from 2 function calls in series. Then we are scheduling any of the two goroutines and which goroutines to schedule is determined by the scheduler. This will yield the following result

```
main execution started
H e l l o 1 2 3 4 5 
main execution stopped
```

Above result again proves that goroutines are cooperatively scheduled. Let’s add another *time.Sleep* call in-between print operation in the function definition to tell the scheduler to schedule other available goroutines.

```
package main

import "fmt"

var start time.Time

func init() {
    start = time.Now()
}

func getChars(s string){
    for _, c := range s {
        fmt.Printf("%c at time %v\n", c, time.Since(start))
        time.Sleep(10 * time.Millisecond)
    }
}

func getDigits(s []int){
    for _, d := range s {
        fmt.Printf("%d at time %v\n", d, time.Since(start))
        time.Sleep(30 * time.Millisecond)
    }
}

func main() {
    fmt.Println("main execution started at time", time.Since(start))

    go getChars("Hello")

    go getDigits([]int{1,2,3,4,5})

    time.Sleep(200 * time.Millisecond)

    fmt.Println("\n main execution stopped at time", time.Since(start))
}
```

In the above program, we printed extra information to see when a print statement is executing since the time of execution of the program. In theory, the main goroutine will sleep for 200 milliseconds, hence all other goroutines must do their job in 200 milliseconds before it wakes up and kills the program. *getChars* goroutine will print 1 character and sleep for 10 milli-second, passing control to *getDigits* goroutine which will print a digit and sleeping for 3 milli-seconds passing control to getChars goroutine again when it wakes up. Since *getChars* goroutine can **print and sleep** multiple times, at least 2 times while other goroutines are sleeping, we are hoping to see more characters printed in succession than digits.

We can see the pattern we talked about. This will be cleared to you once you see the program execution diagram. We will approximate that print command takes 1ms of CPU time, compared on the 200ms scale, that’s negligible.

![](https://miro.medium.com/v2/resize:fit:1400/format:webp/0*4_Z0LRvi_DJR1JEr.jpg)

Now we understood how to create goroutine and how to work with them. But using time.Sleep is just a hack to see the result. In production, we don’t know how much time a goroutine is going to take for the execution. Hence we can't just add random sleep call in the main function. We want our goroutines to tell when they finished the execution. Also at this point, we don’t know how we can get data back from other goroutines or pass data to them, simply, communicate with them. This is where **channels** comes in. Let’s talk about them in the next lesson.

## ☛ Anonymous goroutines
If an anonymous function can exist then anonymous goroutine can also exist. Please read *Immediately invoked function* from functions lesson to understand this section. Let’s modify our earlier example of printHello goroutine.

```
package main

import "fmt"

func main() {
    fmt.Println("main execution started")

    go func() {
        fmt.Println("Hello World!")
    }()

    // schedule another goroutine
    time.Sleep(10 * time.Millisecond)
    fmt.Println("main execution stopped")
}
```

The result is quite obvious as we defined the function and executed as goroutine in the same statement.

*All goroutines are anonymous as we learned from concurrency lesson as goroutine does not have an identity. But we are calling that in the sense that function from which it was created was anonymous.*


