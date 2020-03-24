#Concurrency

Concurrency: 
- Creating multiple processes that execute independently.
- _Dealing with lots_ of things at once.
- Challenges:
    - Coordinating tasks
    - Shared memory

Parallelism: 
- Simultaneous execution of (possibly related) computations.
- _Doing lots_ of things at once.
- GOMAXPROCES()

Channels:
- Don't communicate by sharing memory, share memory by communicating.
- A channel make a copy of a piece of memory and pass it by to another goroutine.
- Channels are blocking constructs: In order to receive/send a message from/to the channel 
a message needs to be available or it will block until a message is available.
- A channel *needs to have a match of senders or receivers otherwise it's blocked*.
- Golang doesn't use os threads for concurrency.

## Channel types
- Bidirectional: Created channels are always bidirectional.
    `ch := make(ch int)`

- Send only: `func foo(ch chan<- int){...}`

- Receive only: `func foo(ch <-chan int){...}`

## Creating channels

    ch := make(chan int) //type of message which is going to send/receive in that channel
    ch := make(chan int, 5) //buffered channel with a capacity of 5.

Concurrency models:
- Processor threads: 
    - process: how the os views a single instance of an application.
    - thread: 
        - sequence of programming instructions that execute in order. 
        - Have own execution stack
        - Fixed stack space (around 1MB)
        - This is what runs the application on the context of the application.
        - Mutex one of the programming constructs to coordinate between threads, use to control how many threads 
        are allowed to access a certain piece of code at one time and while accessing what they are allowed to do.
    - advatanges:
        - managed by os
        - performance.
    - disadvatanges:
        - poor performance due syncing communication between threads.
        - memory consumption(dedicated memory for each thread stack).
        - shared memory between threads.
        - race conditions when 2 or more threads are compiting for the same resource.
- Communicating Sequential Processes(CSP):
    - aka Actor Model.
    - Individual entities to work in a synchronous and isolated way only with data that they are provided 
    with and then passing the results on to be worked on by whatever is downstream from them. 
    - all the async logic and coordination are buried  in the infrastructure of the language or framework.
    - entities:
        - Actor: Responsible for receiving information in, applying some processing to it, and passing 
        the results to the next actor in the chain.
        - Message: Object that is passed between 2 actors. Once a message is being passed between one 
        actor to the other the information contained in that message is no longer accesible to the sender. 
        Only one actor can work on data at the time.
    - advantages:
        - Fully decoupled.
        - Multiple Handlers.
        - Memory isolation.
    - disadvatanges:
        - Complicated mental model.
        - Traceability is difficult.

Concurrency in Go:
- No thread primitives.
- Have own execution stack
- Variable stack space(starts at 2KB)
- Goroutines: all the concurrency of a go program happens on virtual thread based on OS threads.
    - Lighter weight than os threads(2kb).
    - Go runtime manages goroutines which handles how it's gonna be schedule on the processor thread which has made available to the application.
    - Less switching between threads because the gorutine is reused when other threads are idle. 
    - Faster start-up times.
    - Safe communication due channels.
- Channels: safe way to synchronize between actors.
    - unbuffered: sender blocks until receiver picks it up. 
        - `make(chan int)`
    - buffered: sender gets blocked only if channel is full.
        - `make(chan int, n)`
    - If there's no data in a channel, a receiver is blocked even if it's unbuffered or buffered channel.
    - closing channels:
        - `close()` closes a channel.
        - Cannot check a closed channel.
        - Cannot open a channel after it's closed.
        - Sending new messages to a closed channel triggers a panic.
        - Receiving messages from a closed channel is feasible:
            - If buffered, all buffered messages are available.
            - If ubuffered, or buffer empty, receive _zero-value_.
        - Closing needs to be always in the sending side of the operation
        - Even if closing can be possible in the receiver, it will error out if 
         close() is defined in a reciever function.
    - control flow:
        - if statements: 
            - `if msg, ok := <ch, ok {...}`
            - Returns `true` if channel is open.
            - Returns `false` if channel is closed.
        - for loop:
            - `for msg := range ch{..}` in the receiver.
            - `close()` needs to be defined in the sender to avoid an infinite loop in the receiver.
        - select():
            - It doesn't look for trustiness(as switch statements), but if it's possible to send or receive from a channel.
            - There's no predetermine order to choose a case; in other words, if more than one case fits the case, it will be chosen randomly.
            - `select()` will block until it can act in one of the cases - *blocking select statement*
            - `default` case is to avoid the blocking select statement.
- Sync package: Allow goroutines to coordinate their work.
    - `sync.WaitGroup`: waits for a collection of goroutines to finish.
    - `sync.Mutex`: A mutual exclusion lock. 
        - Protects memory access by locking the mutex, accessing the memory and unlocking the mutex.
