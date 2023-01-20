import {
  loadPackageDefinition,
  Server,
  ServerCredentials,
} from "@grpc/grpc-js";
import { loadSync } from "@grpc/proto-loader";
const TRANSFORMER_PROTO_PATH = "./proto/transformer.proto";
const SERVER_URL = "localhost:50051";

const OPTIONS = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};
const packageDefinition = loadSync(TRANSFORMER_PROTO_PATH, OPTIONS);
const transformer = loadPackageDefinition(packageDefinition).Transformer;

const server = new Server();

server.addService(transformer.service, {
  Uppercase: (call, callback) => {
    console.log(`\x1b[35mgRPC received:\u001b[0m ${JSON.stringify(call.request)}`)
    const { message } = call.request;
    const outputMessage = {
      message: message.toLocaleUpperCase(),
      action: true,
    };
    callback(null, outputMessage);
  },
});

server.bindAsync(SERVER_URL, ServerCredentials.createInsecure(), () => {
  console.log(`\x1b[32mgRPC server running at:\u001b[0m ${SERVER_URL}`);
  server.start();
});
