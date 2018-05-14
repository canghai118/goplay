package main

type Channel interface {
	send(user User, notification *Notification) error
}
