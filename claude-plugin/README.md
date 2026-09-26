# EcoAloha

EcoAloha finds outdoor, nature and adventure experiences near 133 gateway cities.
It also serves manually curated destination guides.

## Use it

Ask Claude for a hike, a bike tour, a kayak trip or a nature guide in a city.
Examples:

- "Find a hiking tour near Barcelona and show me the booking page."
- "Compare the outdoor tours EcoAloha has near Las Vegas."
- "Which EcoAloha guides cover Lisbon?"

Claude lists the destinations, searches the reviewed offers and compares them.
It gives the starting price and an affiliate link to the provider page.
EcoAloha does not book or charge. Booking and payment happen on Viator.

Offers cover a limited set of reviewed physical tours and admissions.
Some destinations have guides but no offers.

## Components

- MCP connector `ecoaloha` at `https://ecoaloha.com/mcp`. It needs no account and no key.
- Skill `ecoaloha-trip-planning`: the search, compare and disclosure workflow for travellers.
- Skill `ecoaloha-api-integration`: setup and tests for developers who integrate the REST API or MCP server.

## Data

The connector sends the destination, filters, product IDs and the chosen currency to `ecoaloha.com`.
EcoAloha fetches offers and reviews from the Viator partner API.
It stores no chat content and no search history.
Rate limits count requests per IP address.
A share list is created only on request. It is public by link and expires after 30 days.
The plugin runs no local code and sends data to no other destination.

- Privacy: https://ecoaloha.com/privacy
- Terms: https://ecoaloha.com/terms
- Contact: hello@ecoaloha.com
