import sys
import unittest
import platform


class ConfigTest(unittest.TestCase):
    def test_flag_matches(self):
        """Test the flag matches as expected. It's always arm64 or x86_64 because we don't build Please for x86."""
        if platform.machine() == "arm64":
            self.assertEqual("--arch=arm64", sys.argv[-1])
        else:
            self.assertEqual("--arch=x86_64", sys.argv[-1])
