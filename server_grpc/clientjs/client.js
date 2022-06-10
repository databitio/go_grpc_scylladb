const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
var PROTO_PATH = "../proto/server_grpc.proto";

var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

var protoDescriptor = grpc.loadPackageDefinition(packageDefinition).server_grpc;

const client = new protoDescriptor.TicketService(
    "localhost:50051",
    grpc.credentials.createInsecure()
  );

module.exports = client;

//   client.ReadTickets({}, (error, news) => {
//     if (!error) throw error
//       console.log(news);
//   });

// function getServer() {
//     var server = new grpc.Server();
//     server.addService(routeguide.RouteGuide.service, {
//       getFeature: getFeature,
//       listFeatures: listFeatures,
//       recordRoute: recordRoute,
//       routeChat: routeChat
//     });
//     return server;
//   }
//   var routeServer = getServer();
//   routeServer.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
//     routeServer.start();
//   });