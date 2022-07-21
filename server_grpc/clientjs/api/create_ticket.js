const client = require("./client");

export function goCreateTicket(ticket) {
  return new Promise( resolve => {
    client.CreateTicket(ticket, (error, tickets) => {
    if (!error) {console.log(error)}
      resolve(tickets);
    });
  })
}

export function goDeleteTicket(ticket) {
  return new Promise( resolve => {
    client.DeleteTicket({Result: ticket.Ticketid}, (error, tickets) => {
        if (!error) {console.log(error)}
          resolve(tickets);
      });
    })
}

export function goGetTicket(ticketid){
    return new Promise( resolve => {
        client.GetTicket({Result: ticketid}, (error, ticket) => {
            if (!error) {console.log(error)}
              resolve(ticket)
        });
    })
}

export function goReadTickets() {
  return new Promise( resolve => {
    client.ReadTickets({}).on('data', function(ticket){
        if (!error) {console.log(error)}
          resolve(ticket)
    });
  })
}

export function goUpdateTicket(ticket) {
  return new Promise( resolve => {
    client.UpdateTicket(ticket, (error, tickets) => {
        if (!error) {console.log(error)}
          resolve(tickets);
      });
    })
}