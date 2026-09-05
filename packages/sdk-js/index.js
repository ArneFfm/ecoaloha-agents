export class EcoAlohaError extends Error {
  constructor(status, body, retryAfter) {
    super(body?.detail ?? `EcoAloha request failed (${status})`);
    this.status = status;
    this.body = body;
    this.retryAfter = retryAfter;
  }
}

export class EcoAloha {
  constructor({ baseUrl = "https://ecoaloha.com/api/v1", fetch: fetcher = globalThis.fetch } = {}) {
    this.baseUrl = baseUrl.replace(/\/$/, "");
    this.fetch = fetcher;
  }

  async request(path, { method = "GET", query = {}, body, signal } = {}) {
    if (!/^\/[a-zA-Z0-9_/-]*$/.test(path) || path.includes("//")) {
      throw new TypeError("Use an API-relative path, such as /destinations");
    }
    const url = new URL(this.baseUrl + path);
    for (const [key, value] of Object.entries(query)) {
      if (value !== undefined) url.searchParams.set(key, String(value));
    }
    const response = await this.fetch(url, {
      method,
      signal,
      headers: {
        Accept: "application/json",
        ...(body === undefined ? {} : { "Content-Type": "application/json" }),
      },
      ...(body === undefined ? {} : { body: JSON.stringify(body) }),
    });
    const payload = await response.json();
    if (!response.ok)
      throw new EcoAlohaError(response.status, payload, response.headers.get("Retry-After"));
    return payload;
  }

  destinations() {
    return this.request("/destinations");
  }
  experiences(query) {
    return this.request("/experiences", { query });
  }
  compare(experienceIds, currency = "EUR") {
    return this.request("/compare", { method: "POST", body: { experienceIds, currency } });
  }
}
