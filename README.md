# EcoAloha agent integrations

Official integration resources for [EcoAloha](https://ecoaloha.com).
Read the [developer portal](https://ecoaloha.com/developers) and [OpenAPI specification](https://ecoaloha.com/openapi.json).

## Connect with MCP

Use Streamable HTTP at https://ecoaloha.com/mcp. No API key is required.
Find the official [EcoAloha listing on Smithery](https://smithery.ai/servers/ecoaloha/ecoaloha).
The separate documentation server is https://ecoaloha.com/mcp/docs.
The [server card](https://ecoaloha.com/.well-known/mcp/server-card.json) lists the product tools.
Verified listing: [official MCP Registry](https://registry.modelcontextprotocol.io/v0.1/servers/io.github.ArneFfm%2Fecoaloha/versions/latest).

## Install skills

```sh
npx skills add ArneFfm/ecoaloha-agents
```

- [ecoaloha-trip-planning](https://skills.sh/arneffm/ecoaloha-agents/ecoaloha-trip-planning): find and compare outdoor experiences.
- [ecoaloha-api-integration](https://skills.sh/arneffm/ecoaloha-agents/ecoaloha-api-integration): integrate and test the HTTP API or MCP server.

The root `plugin.json` and `mcp.json` also form an Agent Plugin.

## SDKs and CLI

Download installable packages from [release 0.1.0](https://github.com/ArneFfm/ecoaloha-agents/releases/tag/v0.1.0).
The JavaScript SDK and CLI are published as `@ecoaloha/sdk` on npm.

```sh
npm install @ecoaloha/sdk
npx --package=@ecoaloha/sdk ecoaloha --sandbox destinations
```

The Python SDK is published as `ecoaloha` on PyPI. Python 3.10 or later is required.

```sh
python3 -m pip install ecoaloha
```

You can also install from this checkout:

```sh
npm install ./packages/sdk-js
python3 -m pip install ./packages/sdk-python
node packages/sdk-js/cli.js --sandbox destinations
```

See the [JavaScript guide](packages/sdk-js/README.md) and [Python guide](packages/sdk-python/README.md).

## Safety and scope

The API is public and keyless. Respect rate limits and Retry-After.
Live data requires partner entitlement. Do not retry or bypass AGENT_ACCESS_NOT_APPROVED.
Use fictional fixtures at https://ecoaloha.com/api/sandbox/v1 for read-only testing.
Booking and payment happen on Viator. Prices are indicative. Disclose affiliate links.

## Publishing

Manual workflows publish the MCP listing and SDK packages separately.
Package versions are immutable after publication. Publisher accounts must be configured first.
For PyPI, register a pending trusted publisher for project `ecoaloha`, owner `ArneFfm`, repository `ecoaloha-agents`, workflow `publish-pypi.yml`.
For npm, publish the first release with an account that owns `@ecoaloha`, then configure workflow `publish-npm.yml` as a trusted publisher.
