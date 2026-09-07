import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { mkdtempSync, rmSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";

test("packed package installs its SDK and CLI", () => {
  const directory = mkdtempSync(join(tmpdir(), "ecoaloha-sdk-"));
  const source = fileURLToPath(new URL("..", import.meta.url));
  const run = (command, args) => execFileSync(command, args, {
    cwd: directory,
    encoding: "utf8",
    stdio: ["ignore", "pipe", "pipe"],
  });
  try {
    const [packed] = JSON.parse(run("npm", ["pack", source, "--json", "--ignore-scripts"]));
    run("npm", [
      "install", join(directory, packed.filename), "--ignore-scripts", "--no-audit", "--no-fund",
    ]);
    const help = run(join(directory, "node_modules/.bin/ecoaloha"), ["--help"]);
    assert.match(help, /Usage: ecoaloha/);
    run(process.execPath, ["--input-type=module", "-e", `
      import assert from "node:assert/strict";
      import { EcoAloha } from "@ecoaloha/sdk";
      const client = new EcoAloha({ fetch: async () => Response.json({ data: [] }) });
      assert.deepEqual(await client.destinations(), { data: [] });
    `]);
  } finally {
    rmSync(directory, { recursive: true, force: true });
  }
});
