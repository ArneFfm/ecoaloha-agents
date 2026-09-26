---
name: ecoaloha-api-integration
description: Integrate and test the EcoAloha REST API or MCP server. Use for developer setup, response handling, and safe fixture tests.
---

# ecoaloha-api-integration

Use https://ecoaloha.com/api/v1 for production. No registration or bearer token is required.
Read https://ecoaloha.com/openapi.json for the affected operation and schema.
Use https://ecoaloha.com/developers.md for setup or response-handling details when needed.
REST responses wrap their payload in data. Handle application/problem+json errors and preserve unknown fields.
Respect Retry-After. A production 503 can indicate unavailable partner access; do not bypass that gate.

For fixture tests, use https://ecoaloha.com/api/sandbox/v1/destinations or
https://ecoaloha.com/api/sandbox/v1/experiences?destinationId=paris.
The sandbox has fictional data and rejects writes. Never interpret fixtures as bookable offers.
Verify the requested integration with fixture reads and error handling before reporting it complete.

For MCP integration, use https://ecoaloha.com/mcp for travel tools.
Check JSON-RPC errors and result.isError before reading structuredContent.
Use https://ecoaloha.com/mcp/docs for search_docs. Discover the selected server's tool list; the lists differ.

Do not add OAuth to this keyless API. Do not bypass partner-access gates.
Share-list inputs contain experienceIds and optional currency. Do not invent a title field or send personal information. Never log deletion tokens.
Example: verify fixture search, then change the base URL to production and handle a possible 503.
Source and SDKs: https://github.com/ArneFfm/ecoaloha-agents.

## Destination and page discovery

Use GET /page-references?type=collection&id=monteverde for parent, child, and related pages.
MCP exposes the same result through get_page_references. Index pages use id=index.

GET /destinations returns all enabled destinations. The website directory uses q, country, and page with 24 results per page.
GET /destinations/nearest requires both lat and lng. It returns the closest enabled destination without a distance limit.
Check distanceKm before describing a result as nearby. The website limits approximate suggestions to 200 kilometres.
A traveller's explicit destination takes priority over an inferred location.
