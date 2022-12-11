import unittest

from day10 import *


class TestDay10(unittest.TestCase):

    def test_part1_tiny(self):
        expected = 0
        actual = runPart1('tiny-test-data.txt')

        self.assertEqual(expected, actual)

    def test_part1(self):
        expected = 13140
        actual = runPart1('test-data.txt')

        self.assertEqual(expected, actual)

    def test_part2(self):
        expected = """##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######....."""
        actual = runPart2('test-data.txt')

        self.assertEqual(expected, actual)


if __name__ == '__main__':
    unittest.main()
