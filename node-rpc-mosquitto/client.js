import { loadPackageDefinition, credentials } from "@grpc/grpc-js";
import { connect } from "mqtt";
import { loadSync } from "@grpc/proto-loader";

const TRANSFORMER_PROTO_PATH = "./proto/transformer.proto";
const mqttClient = connect("mqtt://127.0.0.1");
const topicName = "transformer-uppercase";

const OPTIONS = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

const packageDefinition = loadSync(TRANSFORMER_PROTO_PATH, OPTIONS);
const transformer = loadPackageDefinition(packageDefinition).Transformer;

const grpcClient = new transformer(
  "localhost:50051",
  credentials.createInsecure()
);

mqttClient.on("connect", () => {
  mqttClient.subscribe(topicName, (err, granted) => {
    if (err) console.error(err);
    if (granted)
      console.log(
        `\u001b[32mMosquitto successfully subscribe to:\u001b[0m ${topicName}`
      );
  });
});

mqttClient.on("message", (topic, input) => {
  if (topic === topicName) {
    const payload = { message: input.toString() };

    console.log(
      `______________________________________________________\n\x1b[35mMosquitto received:\u001b[0m ${input.toString()}`
    );
    grpcClient.Uppercase(payload, (error, result) => {
      if (error) console.error(error);
      console.log(
        `\x1b[33mgRPC result:\u001b[0m ${JSON.stringify(result)}`
      );
    });
  }
});
