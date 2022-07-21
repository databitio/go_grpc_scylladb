import uuid from "uuid";
import { goCreateTicket, goGetTicket, goUpdateTicket, goDeleteTicket } from '../api/create_ticket';

var ticketid = uuid.v4()
var serverid = uuid.v4();
var userid = uuid.v4();

let ticket = {
  Ticketid: ticketid, 
  Serverid: serverid, 
  Userid: userid,
  Title: "example title",
  Description: "example description",
  Reward: "example reward",
  Lifespan: "example lifespan",
  Type: "example type",
  Archived: false,
  Status: "example status",
  Claimed: true,
}

export const TestEndpoints = () => {
    console.log("Creating ticket...")
  await ticketCallback(ticket, goCreateTicket).then(value => {
    console.log("Ticket created!", value)
  })

  console.log("Getting ticket...")
  await ticketCallback(ticket.Ticketid, goGetTicket).then(value => {
    console.log("Ticket found in database!", value)
  })

  ticket.Title = "This is the updated title!"

  console.log("Updating ticket...")
  await ticketCallback(ticket, goUpdateTicket).then(value => {
    console.log("Ticket updated!", value)
  })

  console.log("Deleting ticket...")
  await ticketCallback(ticket, goDeleteTicket).then(value => {
    console.log("Ticket deleted!", value)
  })
}