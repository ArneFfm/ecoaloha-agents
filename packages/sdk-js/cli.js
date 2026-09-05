#!/usr/bin/env node
import { EcoAloha } from "./index.js";

const args = process.argv.slice(2);
const sandbox = args[0] === "--sandbox";
if (sandbox) args.shift();
const [command, ...params] = args;
if (!command || command === "--help") {
  console.log("Usage: ecoaloha [--sandbox] destinations | experiences destinationId [interest]");
} else {
  try {
    const client = new EcoAloha(sandbox ? { baseUrl: "https://ecoaloha.com/api/sandbox/v1" } : {});
    let result;
    if (command === "destinations" && params.length === 0) result = await client.destinations();
    else if (command === "experiences" && params.length >= 1 && params.length <= 2) {
      result = await client.experiences({ destinationId: params[0], interest: params[1] });
    } else throw new Error("Unknown command or arguments. Run ecoaloha --help.");
    console.log(JSON.stringify(result, null, 2));
  } catch (error) {
    console.error(error instanceof Error ? error.message : "Request failed");
    process.exitCode = 1;
  }
}
