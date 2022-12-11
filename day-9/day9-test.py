import unittest

from day9 import *


class TestDay9(unittest.TestCase):

    def test_part1(self):
        test_file = 'test-data.txt'

        expected = 13
        actual = runPart1(test_file)

        self.assertEqual(expected, actual)


if __name__ == '__main__':
    unittest.main()
