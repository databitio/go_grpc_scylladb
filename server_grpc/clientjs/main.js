const client = require("./client");
const { v4: uuidv4 } = require('uuid');

var uuid = uuidv4();

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

function main() {
    // client.ReadTickets()
    client.CreateTicket(myticket, (tickets, error) => {
    if (!error) {console.log(error)}
      console.log(tickets);
  });
  // err = goCreateTicket(c, ticketinfo)
//   stream, err := c.ReadTickets(context.Background(), &emptypb.Empty{})
}

main()

