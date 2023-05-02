var PROTO_PATH = './proto/vote.proto';

var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc\proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var vote_proto = grpc.loadPackageDefinition(packageDefinition).vote;

function AddVoto(call, callback){
    const query =`INSERT INTO Vote(sede, departamento, municipio, papeleta,partido) VALUES ('${call.request.sede}','${call.request.departamento}','${call.request.municipio}','${call.request.papeleta}','${call.request.partido}')`
    
    19/4/23 1:15:00
    // console.log("CALL",call);
}