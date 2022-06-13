const client = require("./client");
const { v4: uuidv4 } = require('uuid');

// var uuid = uuidv4();
var ticketid = "1e47633c-5e08-2b3c-3e21-2940073f2255"
var serverid = "2b130e31-3f31-4f55-0605-121262523117";
var userid = "0a31055e-1641-1f15-0b59-484e28515d28";

let myticket = {
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

function goCreateTicket(ticket) {
  return new Promise( resolve => {
    client.CreateTicket(ticket, (error, tickets) => {
    if (!error) {console.log(error)}
    resolve(tickets);
    });
  })
}

function goDeleteTicket(ticket) {
  return new Promise( resolve => {
    client.DeleteTicket({Result: ticket.Ticketid}, (error, tickets) => {
        if (!error) {console.log(error)}
        resolve(tickets);
      });
    })
}

function goGetTicket(ticketid){
    return new Promise( resolve => {
        client.GetTicket({Result: ticketid}, (error, ticket) => {
            if (!error) {console.log(error)}
            resolve(ticket)
        });
    })
}

function goReadTickets() {
  return new Promise( resolve => {
    client.ReadTickets({}).on('data', function(ticket){
        if (!error) {console.log(error)}
        resolve(ticket)
    });
  })
}

function goUpdateTicket(ticket) {
  return new Promise( resolve => {
    client.UpdateTicket(ticket, (error, tickets) => {
        if (!error) {console.log(error)}
        resolve(tickets);
      });
    })
}

 async function main() {
  console.log("Creating ticket...")
  await ticketCallback(myticket, goCreateTicket).then(value => {
    console.log("Ticket created!", value)
  })
  
  console.log("Getting ticket...")
  await ticketCallback(myticket.Ticketid, goGetTicket).then(value => {
    console.log("Ticket found in database!", value)
  })

  myticket.Title = "This is the updated title!"

  console.log("Updating ticket...")
  await ticketCallback(myticket, goUpdateTicket).then(value => {
    console.log("Ticket updated!", value)
  })

  // console.log("Getting all tickets...")
  // ticketCallback(myticket.Ticketid, goReadTickets).then(value => {
  //   console.log("All tickets in database:", value)
  // })

  console.log("Deleting ticket...")
  await ticketCallback(myticket, goDeleteTicket).then(value => {
    console.log("Ticket deleted!", value)
  })
}

async function ticketCallback(call, callback) {
    response = await callback(call)
    return response
}

main()

