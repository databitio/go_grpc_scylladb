

function goCreateTicket(callback) {

    var rectangle = {
      lo: {
        latitude: 400000000,
        longitude: -750000000
      },
      hi: {
        latitude: 420000000,
        longitude: -730000000
      }
    };


    var call = client.CreateTicket(rectangle);
    call.on('data', function(feature) {
        console.log('Found feature called "' + feature.name + '" at ' +
            feature.location.latitude/COORD_FACTOR + ', ' +
            feature.location.longitude/COORD_FACTOR);
    });
    call.on('end', callback);
  }