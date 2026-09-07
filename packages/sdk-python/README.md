# EcoAloha Python SDK

Official source: https://github.com/ArneFfm/ecoaloha-agents.
API documentation: https://ecoaloha.com/developers.
Python 3.10 or later is required. Install with `python3 -m pip install ecoaloha`.
From a checkout, use `python3 -m pip install ./packages/sdk-python`.

```python
from ecoaloha import EcoAloha
client = EcoAloha(base_url="https://ecoaloha.com/api/sandbox/v1")
print(client.experiences("paris"))
```

Responses preserve the data envelope and pagination cursor.
Use `request(path, method=..., query=..., body=...)` for other documented operations.
`EcoAlohaError` exposes `status`, `body`, and `retry_after`. Calls time out after 30 seconds.
The client does not retry writes. No API key is required.
