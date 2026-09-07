---
name: ecoaloha-api-integration
description: Integrate and test the EcoAloha REST API or MCP server. Use for developer setup, response handling, and safe fixture tests.
---

# ecoaloha-api-integration

1. Read https://ecoaloha.com/developers.md and https://ecoaloha.com/openapi.json for current operations and schemas.
2. Use https://ecoaloha.com/api/v1 for production. No registration or bearer token is required.
3. Start with GET /destinations. REST responses wrap their payload in data.
4. Test fixture reads at https://ecoaloha.com/api/sandbox/v1/destinations and /experiences?destinationId=paris.
5. The sandbox has fictional data and rejects writes. Never interpret fixtures as bookable offers.
6. Handle REST application/problem+json errors. For MCP, check JSON-RPC errors and result.isError before consuming structuredContent. Respect Retry-After and preserve unknown fields.
7. Use the product MCP server at https://ecoaloha.com/mcp for travel tools.
8. Use https://ecoaloha.com/mcp/docs for search_docs. Its tool list differs from the product server.

Do not add OAuth to this keyless API. Do not bypass partner-access gates.
Share-list inputs contain experienceIds and optional currency. Do not invent a title field or send personal information. Never log deletion tokens.
Example: verify fixture search, then change the base URL to production and handle a possible 503.
Source and SDKs: https://github.com/ArneFfm/ecoaloha-agents.
