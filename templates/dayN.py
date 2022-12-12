def parse_file_to_lines(filepath):
    with open(filepath) as file:
        lines = file.read().strip().split("\n")

    return lines


def runPart1(filename):
    return 0


def runPart2(filename):
    return 0


def main():
    print("Part 1: ")
    print(runPart1('full-data.txt'))

    print("Part 2: ")
    print(runPart2('full-data.txt'))


if __name__ == '__main__':
    main()
