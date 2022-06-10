const client = require("./client");

function main() {
    // client.ReadTickets()
    client.CreateTicket({}, (tickets, error) => {
    if (!error) throw error
      console.log(tickets);
  });
  // err = goCreateTicket(c, ticketinfo)
//   stream, err := c.ReadTickets(context.Background(), &emptypb.Empty{})
}

main()