import io
import unittest
from unittest.mock import patch
from urllib.error import HTTPError
from ecoaloha import EcoAloha, EcoAlohaError

class ClientTests(unittest.TestCase):
    @patch("ecoaloha.urlopen")
    def test_query_and_envelope(self, opener):
        opener.return_value.__enter__.return_value = io.BytesIO(b'{"data": [], "nextCursor": "Mg=="}')
        result = EcoAloha().experiences("paris", interest="food & wine")
        self.assertEqual(result["nextCursor"], "Mg==")
        self.assertIn("interest=food+%26+wine", opener.call_args.args[0].full_url)

    @patch("ecoaloha.urlopen")
    def test_errors(self, opener):
        opener.side_effect = HTTPError("https://ecoaloha.com", 429, "rate limit", {"Retry-After": "60"}, io.BytesIO(b'{"detail":"Slow down"}'))
        with self.assertRaises(EcoAlohaError) as caught:
            EcoAloha().destinations()
        self.assertEqual(caught.exception.retry_after, "60")
        opener.assert_called_once()
        with self.assertRaises(ValueError):
            EcoAloha().request("//evil.example")

if __name__ == "__main__":
    unittest.main()
