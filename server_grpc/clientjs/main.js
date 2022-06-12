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
  Title: "newest loaded from javascript!",
  Description: "javascript description!",
  Reward: "my reward is this working!",
  Lifespan: "2046-10-26",
  Type: "servicesssssssssssssssssssssss!!",
  Archived: false,
  Status: "incompleteeeeeeeeeeeeeeeeeeeeeeeeeeeee",
  Claimed: true,
}

function ticketMessageToTicket (data) {
    return {
    Ticketid: data.Ticketid, 
    Serverid: data.Serverid, 
    Userid: data.Userid,
    Title: data.Title,
    Description: data.Description,
    Reward: data.Reward,
    Lifespan: data.Lifespan,
    Type: data.Type,
    Archived: data.Archived,
    Status: data.Status,
    Claimed: data.Claimed,
    }
}

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

// function getTicket(){
//     let promise = new Promise(function getCallback(call, callback) {
//         if(callback)
//             callback(call)
//         console.log("callback finished!")
//     })
// }
function goGetTicket(ticketid){
    return new Promise( resolve => {
        client.GetTicket({Result: ticketid}, (error, ticket) => {
            if (!error) {console.log(error)}
            resolve(ticket)
        });
    })
}

// function goGetTicket(ticketid) {
//     client.GetTicket({Result: ticketid}, (error, ticket) => {
//         if (!error) {console.log(error)}
//         return ticket
//       });
// }

function goReadTickets() {
    client.ReadTickets({}).on('data', function(ticket){
        console.log(ticket)
    })
    
    // client.ReadTickets({}).on('end',function(){
    //     console.log('All tickets have been read');
    //   });
}

function goUpdateTicket(ticket) {
    client.UpdateTicket(ticket, (error, tickets) => {
        if (!error) {console.log(error)}
        console.log(tickets);
      });
}
var newticket;
// function main() {
//     // goCreateTicket(myticket)
//     // console.log("ticket created!", uuid)
//     // goDeleteTicket(myticket)
//     // newticket, error = goGetTicket(myticket.Ticketid, newticket)
//     newticket = await goGetTicket(myticket.Ticketid)
//     console.log(newticket)
//     // goReadTickets()
//     // goCreateTicket(myticket)
//     // goUpdateTicket(myticket)
// }

async function callback() {
    newticket = await goGetTicket(myticket.Ticketid)
    console.log(newticket)
}

callback()
