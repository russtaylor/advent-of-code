import unittest

from day12 import *


class TestDay12(unittest.TestCase):

    def test_part1(self):
        expected = 31
        actual = runPart1('test-data.txt')

        self.assertEqual(expected, actual)

    def test_part2(self):
        expected = 0
        actual = runPart2('test-data.txt')

        self.assertEqual(expected, actual)


if __name__ == '__main__':
    unittest.main()
