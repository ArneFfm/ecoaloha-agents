# EcoAloha JavaScript SDK and CLI

Official source: https://github.com/ArneFfm/ecoaloha-agents.
API documentation: https://ecoaloha.com/developers.
Registry publication is pending. Node.js 20 or later is required.

From a checkout, run `npm install ./packages/sdk-js` in your integration project.

```js
import { EcoAloha } from "@ecoaloha/sdk";
const client = new EcoAloha({ baseUrl: "https://ecoaloha.com/api/sandbox/v1" });
console.log(await client.experiences({ destinationId: "paris" }));
```

Run `node packages/sdk-js/cli.js --sandbox destinations` from the repository root.
Run `ecoaloha --help` after package installation.

The client returns the complete response envelope, including `nextCursor`.
Use `request(path, options)` for other documented operations.
Errors expose `status`, `body`, and `retryAfter`. The client does not retry writes.
No key is required. Respect rate limits and affiliate disclosure rules.
