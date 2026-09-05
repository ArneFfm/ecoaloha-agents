"""Public EcoAloha API client. No authentication is required."""
import json
import re
from urllib.error import HTTPError
from urllib.parse import urlencode
from urllib.request import Request, urlopen


class EcoAlohaError(Exception):
    def __init__(self, status, body, retry_after):
        super().__init__(body.get("detail", f"EcoAloha request failed ({status})"))
        self.status = status
        self.body = body
        self.retry_after = retry_after


class EcoAloha:
    def __init__(self, base_url="https://ecoaloha.com/api/v1", timeout=30):
        self.base_url = base_url.rstrip("/")
        self.timeout = timeout

    def request(self, path, *, method="GET", query=None, body=None):
        if not re.fullmatch(r"/[a-zA-Z0-9_/-]*", path) or "//" in path:
            raise ValueError("Use an API-relative path, such as /destinations")
        url = self.base_url + path
        if query:
            url += "?" + urlencode({k: v for k, v in query.items() if v is not None})
        headers = {"Accept": "application/json"}
        data = None
        if body is not None:
            data = json.dumps(body).encode()
            headers["Content-Type"] = "application/json"
        request = Request(url, data=data, headers=headers, method=method)
        try:
            with urlopen(request, timeout=self.timeout) as response:
                return json.load(response)
        except HTTPError as error:
            with error:
                payload = json.load(error)
            raise EcoAlohaError(error.code, payload, error.headers.get("Retry-After")) from error

    def destinations(self):
        return self.request("/destinations")

    def experiences(self, destination_id, **filters):
        return self.request("/experiences", query={**filters, "destinationId": destination_id})

    def compare(self, experience_ids, currency="EUR"):
        return self.request("/compare", method="POST", body={"experienceIds": experience_ids, "currency": currency})
