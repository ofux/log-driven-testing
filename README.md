# Log Driven Testing

## The idea

The idea behind *Log Driven Testing* is to make some assertions on the log messages that *should* be output by the code you test.

Before reading further, it is important to understand that this idea originally comes from an observation: **log messages are often missing, or they don't provide enough information. Sometimes, they are even wrong and can misslead investigations in case of problems.** The aim of *Log Driven Testing* is to fix this, while at the same time providing a new way of writting automated tests.

## What is it?

Usually, when we write automated tests, we make assertions on the *results* of a function (what it returns), or on the *state* of some objects that may be affected by the function, or even on *calls* to other functions that may be called by the function we are testing.

*Log Driven Testing* has a different approach, in that it doesn't assert this usual stuff, but instead asserts *logs*. The purpose is not necesseraly to replace classical tests, but it can advantageously complement them.

Main advantages *could be*:

- You are obliged to write proper log messages to have good tests.
- Log messages are also tested.
- Assertions are easy to read as they are plain text (log messages).
- Assertions are very easy to write.
- Multiple level of testing can be achieved thanks to different log levels.

## Experimentation

Log Driven Testing must be tried in a real-world application in order to see how it behaves and if it brings any productivity (and/or robustness) gain.

For this purpose, I am implementing a simple Go library to make it simple to write such tests.
