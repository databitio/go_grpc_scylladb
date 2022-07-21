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
	Lifespan    time.Time  `json:"lifespan"`
	Type        string     `json:"type"`
	Archived    bool       `json:"archived"`
	Status      string     `json:"status"` //in progress, new, etc.
	Claimed     bool       `json:"claimed"`
}
