import unittest

from day11 import *


class TestDay11(unittest.TestCase):

    def test_part1(self):
        expected = 10605
        actual = runPart1('test-data.txt')

        self.assertEqual(expected, actual)

    def test_part2(self):
        expected = 2713310158
        actual = runPart2('test-data.txt')

        self.assertEqual(expected, actual)


if __name__ == '__main__':
    unittest.main()
