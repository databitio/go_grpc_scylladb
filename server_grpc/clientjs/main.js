const client = require("./client");
const { v4: uuidv4 } = require('uuid');

// var uuid = uuidv4();
var uuid = "75bfd13a-54da-4405-8594-0359a5fec418"

let myticket = {
  Ticketid: uuid, 
  Serverid: uuid, 
  Userid: uuid,
  Title: "This was loaded from javascript!",
  Description: "javascript description!",
  Reward: "my reward is this working!",
  Lifespan: "2046-10-26:04:05",
  Type: "service!!",
  Archived: false,
  Status: "incomplete",
  Claimed: true,
}
// let myticket = {
//   Ticketid: data.Ticketid, 
//   Serverid: data.Serverid, 
//   Userid: data.Userid,
//   Title: data.Title,
//   Description: data.Description,
//   Reward: data.Reward,
//   Lifespan: data.Lifespan,
//   Type: data.Type,
//   Archived: data.Archived,
//   Status: data.Status,
//   Claimed: data.Claimed,
// }

function goCreateTicket(ticket) {
    client.CreateTicket(ticket, (error, tickets) => {
    if (!error) {console.log(error)}
      console.log(tickets);
  });
}

function goDeleteTicket(ticket) {
    client.DeleteTicket({Result: ticket.Ticketid}, (error, tickets) => {
        if (!error) {console.log(error)}
          console.log(tickets);
      });
}

function goGetTicket(ticket) {
    client.GetTicket({Result: ticket.Ticketid}, (error, tickets) => {
        if (!error) {console.log(error)}
          console.log(tickets);
      });
}

function goReadTickets() {
    client.ReadTickets({}).on('data', function(ticket){
        console.log(ticket)
    })
        // if (!error) {console.log(error)}
    // let call = client.client.ReadTickets({}, error);

    //   call.on('data',function(response){
    //     console.log(response.message);
    //   });
    
    //   call.on('end',function(){
    //     console.log('All Salaries have been paid');
    //   });
}

// goCreateTicket(myticket)
// console.log("ticket created!", uuid)
// goDeleteTicket(myticket)
// goGetTicket(myticket)
goReadTickets()
