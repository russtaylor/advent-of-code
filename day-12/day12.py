from collections import deque


def parse_file_to_grid(filepath):
    with open(filepath) as file:
        data = file.read().strip()

    grid = [list(line) for line in data.split("\n")]

    start = None
    end = None
    a_list = list()

    for row_index, row in enumerate(grid):
        for char_index, char in enumerate(row):
            if char == "S":
                start = (row_index, char_index)
                grid[row_index][char_index] = "a"
            if char == "E":
                end = (row_index, char_index)
                grid[row_index][char_index] = "z"
            if char == "a":
                a_list.append((row_index, char_index))

    return grid, start, end, a_list


def search(grid, starts, end):
    seen = set()
    queue = deque([(start, 0) for start in starts])

    while queue:
        position, distance = queue.popleft()
        x, y = position
        if position == end:
            return distance
        if position in seen:
            continue
        seen.add(position)
        for dx, dy in ((0, 1), (1, 0), (0, -1), (-1, 0)):
            if (
                0 <= x + dx < len(grid)
                and 0 <= y + dy < len(grid[0])
                and ord(grid[x + dx][y + dy]) - ord(grid[x][y]) <= 1
            ):
                queue.append(((x + dx, y + dy), distance + 1))


def runPart1(filename):
    grid, start, end, _ = parse_file_to_grid(filename)

    result = search(grid, [start], end)

    return result


def runPart2(filename):
    grid, start, end, a_list = parse_file_to_grid(filename)

    result = search(grid, a_list, end)

    return result


def main():
    print("Part 1: ")
    print(runPart1('full-data.txt'))

    print("Part 2: ")
    print(runPart2('full-data.txt'))


if __name__ == '__main__':
    main()
