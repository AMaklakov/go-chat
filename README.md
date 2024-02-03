# TCP Chat using Golang

A simple TCP chat written in Golang.

It supports `room`s so that multiple users can join the same or different chats.

## Example

```console
❯ nc 0.0.0.0 5000
💬 Please, enter your name: 
user 2
💬 Type in room number: 
xxx
💬 *** A new user "user 2" joined the room "xxx" ***
💬 *** Now in room: user 2, user 1 ***
Hello!
💬 user 2 (127.0.0.1:52604): Hello!
💬 user 1 (127.0.0.1:52603): Hey!
How u doing?
💬 user 2 (127.0.0.1:52604): How u doing?
💬 user 1 (127.0.0.1:52603): Doin good!
^C
```

```console
❯ nc 0.0.0.0 5000
💬 Please, enter your name: 
user 1
💬 Type in room number: 
xxx
💬 *** A new user "user 1" joined the room "xxx" ***
💬 *** Now in room: user 1 ***
💬 *** A new user "user 2" joined the room "xxx" ***
💬 *** Now in room: user 2, user 1 ***
💬 user 2 (127.0.0.1:52604): Hello!
Hey!
💬 user 1 (127.0.0.1:52603): Hey!
💬 user 2 (127.0.0.1:52604): How u doing?
Doin good!
💬 user 1 (127.0.0.1:52603): Doin good!
💬 *** User "user 2" left the room "xxx" ***
^C
```

## Usage

Start the server using the following command in the app directory:

```sh
go run .
```

Use `nc` to connect to the server using TCP:

```sh
nc 0.0.0.0 5000
```
