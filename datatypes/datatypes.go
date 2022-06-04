package datatypes

import (
	"time"

	"github.com/gocql/gocql"
)

type Ticket struct {
	Ticketid    gocql.UUID `json:"ticketid"`
	Userid      gocql.UUID `json:"userid"`
	Serverid    gocql.UUID `json:"serverid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reward      string     `json:"reward"`
	Lifespan    int        `json:"lifespan"`
	Type        string     `json:"type"`
	Archived    bool       `json:"archived"`
}

type Server struct {
	Serverid     gocql.UUID   `json:"serverid"`
	Users        []gocql.UUID `json:"users"` //userids
	Name         string       `json:"name"`
	Queuepages   []Queuepage  `json:"queuepages"` //queuepageid
	Date_created time.Time    `json:"date_created"`
}
type Queuepage struct {
	Queuepageid  int         `json:"queuepageid"`
	Queues       []Queue     `json:"queues"` //[queueids],
	Section_name string      `json:"section_name"`
	Name         string      `json:"name"`
	Messengers   []Messenger `json:"messengers"` //[messenger, messenger],
}
type Queue struct {
	Queueid int          `json:"queueid"`
	Name    string       `json:"name"`
	Tickets []gocql.UUID `json:"tickets"` //[ticketids]
}
type Messenger struct {
	Messengerid int          `json:"messengerid"`
	Name        string       `json:"name"`
	Messages    []gocql.UUID `json:"messages"` //[message, message]
}
type Message struct {
	Messageid gocql.UUID `json:"messageid"`
	Userid    gocql.UUID `json:"userid"`
	Content   string     `json:"content"`
	Date      time.Time  `json:"date"`
}

type User struct {
	Userid       gocql.UUID   `json:"userid"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Date_created time.Time    `json:"date_created"`
	Subscription Subscription `json:"subscription"`
}

type Subscription struct {
	Active       bool      `json:"active"`
	Type         string    `json:"type"`
	Renewal_date time.Time `json:"renewal_date"`
}
