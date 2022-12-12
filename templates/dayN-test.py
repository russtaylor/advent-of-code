import unittest

from dayN import *


class TestDayN(unittest.TestCase):

    def test_part1(self):
        test_file = 'test-data.txt'

        expected = 0
        actual = runPart1(test_file)

        self.assertEqual(expected, actual)

    def test_part2(self):
        test_file = 'test-data.txt'

        expected = 0
        actual = runPart2(test_file)

        self.assertEqual(expected, actual)


if __name__ == '__main__':
    unittest.main()
