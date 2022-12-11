def parse_file_to_lines(filepath):
    with open(filepath) as file:
        lines = file.read().strip().split("\n")

    return lines


def addoToSignalIfApplicable(cycle, x):
    if cycle == 20 or ((cycle - 20) % 40 == 0):
        addToSignal = cycle * x
        return addToSignal
    return 0


def runPart1(filename):
    lines = parse_file_to_lines(filename)
    cycle = 0
    x = 1
    signalSum = 0
    for line in lines:
        if 'noop' in line:
            cycle += 1
            signalSum += addoToSignalIfApplicable(cycle, x)

        else:
            op = int(line.split(" ")[1])
            cycle += 1
            signalSum += addoToSignalIfApplicable(cycle, x)
            cycle += 1
            signalSum += addoToSignalIfApplicable(cycle, x)
            x += op

    return signalSum


def drawPixelIfApplicable(cycle, x):
    pixel = '.'
    if (cycle % 40) >= x - 1 and (cycle % 40) <= x + 1:
        pixel = '#'
    if (cycle + 1) % 40 == 0:
        pixel += "\n"
    return pixel


def runPart2(filename):
    lines = parse_file_to_lines(filename)
    cycle = 0
    x = 1
    crtRendered = ""
    for line in lines:
        if 'noop' in line:
            crtRendered += drawPixelIfApplicable(cycle, x)
            cycle += 1
        else:
            op = int(line.split()[1])
            crtRendered += drawPixelIfApplicable(cycle, x)
            cycle += 1
            crtRendered += drawPixelIfApplicable(cycle, x)
            cycle += 1
            x += op

    return crtRendered[:-1]  # Trim the last '\n'


def main():
    print("Part 1: ")
    print(runPart1('full-data.txt'))
    print("Part 2: ")
    print(runPart2('full-data.txt'))


if __name__ == '__main__':
    main()
