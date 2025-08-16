Allowing yourself to do multiple tasks but not at the same time - Concurrency
All tasks at the same time not one by one - Parallelism
Eg. of instagram :-
Suppose you are eating rice and seeing the insta reels, now a notification came, and you want to open the notification, and also want to turn the A.C on while eating the rice
Concurrency - You will first click on the notification see it, then turn on the A.C and then eat the rice
Parallelism - You will simulataneously eat rice, see the notification and turn the A.C on. 
<img src="./concurrencyvsparallel.png">


Go routines - Lightweight threads managed by the Go runtime. - Flexible stack - 2KB
Thread - Managed by the OS. - Fixed stack - 1MB

 GO ROUTINES IS THE WAY WE ACHIEVE PARALLELSIM IN GOLANG 
 Go runtime can fire up more threads without getting permisison from OS. More control is with go runtime.


MOTO OF GOLANG - "DO NOT COMMUNICATE BY SHARING MEMORY; INSTEAD SHARE MEMORY BY COMMUNICATING".


when using go routines in go you have to also use wait groups to wait for all goroutines to finish.
wait groups are basically an advanced version of time.Sleep().
3 functions of wait groups :-
1. Add(int): Increments the WaitGroup counter by the specified number.
2. Done(): Decrements the WaitGroup counter by one.
3. Wait(): Blocks until the WaitGroup counter is zero.



Mutex is a mutual exclusion lock. The zero value for a mutex is an unlocked mutex. A mutex must not be copied after first use. 
Lock/Unlock mutex :- It basically provides you a lock over memory when one go routine is working and till then it will not allow any other to use this memory. 

 Read write mutex :- Allow multiple go routines for sharing resource for reading purposes but does not allow to write on memory when one go routine is working.


 Use go run --race . for checking for race condition errors in go file -> Gives exit status 66
 It can be solved using mutex

 Channels - Channels are a way by which go routines talk to each other by passing value
 Channels allow you to pass them value only if somebody is listening to it otherwise it will give a deadlock error
 Have to use wait groups for channels 