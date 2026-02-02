# PSI

## /proc/pressure/io

some avg10=0.00 avg60=0.02 avg300=0.00 total=823527604
full avg10=0.00 avg60=0.02 avg300=0.00 total=673073799

It is strictly about I/O block contention. It tracks the amount time we
are stalled waiting for read or write operations.

It is calculated via a D state or TASK_UNINTERRUPTIBLE state. PSI tracks exactly
how long it spends in that state.

WHY MATTERS: High I/O pressure means your storage subsystem is the bottleneck.
The CPU is ready to do work but it's sitting idle waiting for data from the disk.

## /proc/pressure/cpu

some avg10=0.03 avg60=0.29 avg300=0.20 total=5529708213
full avg10=0.00 avg60=0.00 avg300=0.00 total=0

This is a temporal metric, just like all others. It measure the amount of stalling for scheduling of the CPU cycles we are seeing.

## /proc/pressure/memory

some avg10=0.00 avg60=0.08 avg300=0.02 total=611806358
full avg10=0.00 avg60=0.08 avg300=0.02 total=587730525

Memory pressure, total displays the amount of time requests were stalled, because
the kernel had to perform work to free up memory before it could hand it over.

When a process requests memory and none is free. The kernel must perform
reclaim operations.

WHY MATTERS: If the percentage is high, your system is spending more time managing memory
rather than running the application.

## some and full

[SOME] At least one task is stalled on this resource.
INTERPRETATION: Throughput impact. The CPU could be doing more work if this resource wasn't holding it back.
It measure lost potential, if you see 20, it means 20% of the time the cpu was stack waiting for this resource.

[FULL] All non-idle tasks are stalled on this resource simultaneously. If it's at 100, your system was doing no work what so ever, all processes froze, inlcuding SSH or typing, anything.
INTERPRETATION: Latency impact. The system is effectively totally unresponsive or freezing during this time.
