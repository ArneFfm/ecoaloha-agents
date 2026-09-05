import assert from "node:assert/strict";
import { test } from "node:test";
import { EcoAloha, EcoAlohaError } from "../index.js";
test("encodes queries and preserves pagination", async () => {
  const client = new EcoAloha({
    fetch: async (url) => {
      assert.equal(url.searchParams.get("interest"), "food & wine");
      return Response.json({ data: [], nextCursor: "Mg==" });
    },
  });
  assert.deepEqual(await client.experiences({ destinationId: "paris", interest: "food & wine" }), {
    data: [],
    nextCursor: "Mg==",
  });
});
test("preserves errors without retrying writes", async () => {
  let calls = 0;
  const client = new EcoAloha({
    fetch: async (_url, options) => {
      calls++;
      assert.equal(options.method, "POST");
      assert.deepEqual(JSON.parse(options.body), { experienceIds: ["A", "B"], currency: "EUR" });
      return Response.json(
        { detail: "Slow down" },
        { status: 429, headers: { "Retry-After": "60" } },
      );
    },
  });
  await assert.rejects(
    client.compare(["A", "B"]),
    (error) => error instanceof EcoAlohaError && error.status === 429 && error.retryAfter === "60",
  );
  assert.equal(calls, 1);
  await assert.rejects(client.request("//evil.example"), TypeError);
});
