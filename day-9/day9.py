from math import dist, copysign

directions = {"R": (1, 0), "L": (-1, 0), "U": (0, 1), "D": (0, -1)}


def parse_file_to_lines(filepath):
    with open(filepath) as file:
        lines = file.read().strip().split("\n")

    return lines


def get_sign(x):
    return (x > 0) - (x < 0)


def follow(knot, lead):
    dx = lead[0] - knot[0]
    dy = lead[1] - knot[1]

    if abs(dx) <= 1 and abs(dy) <= 1:
        return knot

    if abs(dx) > 1 or (abs(dx) + abs(dy)) > 2:
        knot[0] += get_sign(dx)

    if abs(dy) > 1 or (abs(dx) + abs(dy)) > 2:
        knot[1] += get_sign(dy)

    return knot


def runRope(rope, moves):

    visited = {(0, 0)}

    for move in moves:
        direction, distance = move.split()
        for step in range(int(distance)):
            rope[0] = [rope[0][0] + directions[direction]
                       [0], rope[0][1] + directions[direction][1]]

            for knot in range(1, len(rope)):
                updated_position = follow(rope[knot], rope[knot - 1])
                rope[knot] = updated_position

            visited.add(tuple(rope[-1]))

    return len(visited)


def runPart1(filename):
    rope = [[0, 0] for _ in range(2)]
    moves = parse_file_to_lines(filename)
    return runRope(rope, moves)


def runPart2(filename):
    rope = [[0, 0] for _ in range(10)]
    moves = parse_file_to_lines(filename)
    return runRope(rope, moves)


def main():
    print("Part 1: ")
    print(runPart1('full-data.txt'))

    print("Part 2: ")
    print(runPart2('full-data.txt'))


if __name__ == '__main__':
    main()
