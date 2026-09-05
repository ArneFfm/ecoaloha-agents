export interface ApiResponse<T = unknown> {
  data: T;
  nextCursor?: string | null;
}
export interface RequestOptions {
  method?: "GET" | "POST" | "DELETE";
  query?: Record<string, string | number | boolean | undefined>;
  body?: unknown;
  signal?: AbortSignal;
}
export class EcoAlohaError extends Error {
  status: number;
  body: unknown;
  retryAfter: string | null;
}
export class EcoAloha {
  constructor(options?: { baseUrl?: string; fetch?: typeof fetch });
  request<T = unknown>(path: string, options?: RequestOptions): Promise<ApiResponse<T>>;
  destinations(): Promise<ApiResponse<Array<{ id: string; name: string; country: string }>>>;
  experiences(
    query: { destinationId: string } & NonNullable<RequestOptions["query"]>,
  ): Promise<ApiResponse<unknown[]>>;
  compare(
    experienceIds: string[],
    currency?: "EUR" | "USD" | "GBP",
  ): Promise<ApiResponse<unknown[]>>;
}
