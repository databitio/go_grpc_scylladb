package queries

import (
	"fmt"
	"log"

	"github.com/databitio/go_server/datatypes"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

func CreateTicket(session gocqlx.Session, newTicket *datatypes.Ticket) error {

	insertTicket := qb.Insert("meed.ticket").
		Columns("ticketid",
			"userid",
			"serverid",
			"title",
			"description",
			"reward",
			"lifespan",
			"type",
			"archived",
			"status",
			"claimed").Query(session)
	insertTicket.BindStruct(*newTicket)

	if err := insertTicket.ExecRelease(); err != nil {
		log.Fatal("ExecRelease() failed:", err)
		return err
	}
	return nil
}

func GetByID(session gocqlx.Session, uuid gocql.UUID) (datatypes.Ticket, error) {
	q := qb.Select("meed.ticket").Where(qb.Eq("ticketid")).Query(session).Bind(uuid.String())

	rows, err := q.Iter().SliceMap()
	if err != nil {
		fmt.Println("Query error: " + err.Error())
	}
	ticket, err := DBToTicket(rows[0])
	return ticket, err
}

func GetAllTickets(session gocqlx.Session) ([]datatypes.Ticket, error) {

	q := qb.Select("meed.ticket").Query(session)
	defer q.Release()

	var tickets []datatypes.Ticket
	if rows, err := q.Iter().SliceMap(); err == nil {
		for _, row := range rows {
			ticket, err := DBToTicket(row)
			if err == nil {
				tickets = append(tickets, ticket)
			}
		}
	}
	return tickets, nil
}

func UpdateTicket(session gocqlx.Session, ticket *datatypes.Ticket) error {
	q := qb.Update("meed.ticket").
		Set(
			"title",
			"description",
			"reward",
			"lifespan",
			"type",
			"archived",
			"status",
			"claimed").
		Where(qb.Eq("ticketid"), qb.Eq("serverid"), qb.Eq("userid")).Existing().
		Query(session).
		Bind(
			ticket.Title,
			ticket.Description,
			ticket.Reward,
			ticket.Lifespan,
			ticket.Type,
			ticket.Archived,
			ticket.Status,
			ticket.Claimed,
			ticket.Ticketid,
			ticket.Serverid,
			ticket.Userid,
		)

	defer q.Release()

	if err := q.Exec(); err != nil {
		log.Fatalf("Error updating ticket: %v\n", err)
		return err
	}

	fmt.Println("Updated successfully!")
	return nil
}

func DeleteTicket(session gocqlx.Session, uuid gocql.UUID) error {

	w := qb.EqNamed("ticketid", uuid.String())
	deleteTicket := qb.Delete("meed.ticket").Where(w).Query(session).Bind(uuid.String())

	if err := deleteTicket.ExecRelease(); err != nil {
		log.Fatal("ExecRelease() failed:", err)
		return err
	}

	fmt.Println("deleted successfully!")
	return nil
}

func CreateTicketTable(session gocqlx.Session) error {
	DeleteTicketTable(session)

	err := session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS meed WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 2}`)
	if err != nil {
		fmt.Println("create keyspace:", err)
		return err
	}

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS meed.ticket (
		ticketid uuid,
		userid uuid,
		serverid uuid,
		title text,
		description text,
		reward text,
		lifespan date,
		type text,
		archived boolean,
		status text,
		claimed boolean,
		PRIMARY KEY (ticketid, serverid, userid)
		)`)
	if err != nil {
		fmt.Println("create table:", err)
		return err
	}

	return nil
}

func DeleteTicketTable(session gocqlx.Session) error {
	err := session.ExecStmt(`DROP KEYSPACE meed`)
	if err != nil {
		log.Printf("Error deleting ticket table: %v\n", err)
	}
	return nil
}
