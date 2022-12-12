import math


class Monkey:
    def __init__(self, items: list, operator: str, test_value: int, target_true: int, target_false: int):
        self.item_inspected_count = 0
        self.items = items
        self.operator = operator
        self.test_value = int(test_value)
        self.target_true = int(target_true)
        self.target_false = int(target_false)


def parse_monkey_file(filepath):
    with open(filepath) as file:
        input_monkeys = file.read().strip().split("\n\n")

    monkeys = []

    for input_monkey in input_monkeys:
        monkey_lines = input_monkey.split("\n")
        items = [eval(i) for i in monkey_lines[1].split(":")[1].split(",")]
        operator = monkey_lines[2].split(":")[1].split("=")[1].strip()
        test_value = monkey_lines[3].split()[-1]
        target_true = monkey_lines[4].split()[-1]
        target_false = monkey_lines[5].split()[-1]
        monkey = Monkey(items, operator, test_value, target_true, target_false)
        monkeys.append(monkey)

    return monkeys


def eval_item(operator, old):
    return eval(operator)


def evaluate_monkeys(monkeys, rounds, worry_divides=True):
    modulo = 1
    for monkey in monkeys:
        modulo *= monkey.test_value
    for _ in range(rounds):
        for index, monkey in enumerate(monkeys):
            for item in monkey.items:
                new_level = eval_item(monkey.operator, item)
                if worry_divides:
                    relief_level = math.floor(new_level / 3)
                else:
                    relief_level = new_level % modulo
                if relief_level % monkey.test_value == 0:
                    monkeys[monkey.target_true].items.append(relief_level)
                else:
                    monkeys[monkey.target_false].items.append(relief_level)
                monkey.item_inspected_count += 1
            monkey.items = []

    items_inspected = []
    for monkey in monkeys:
        items_inspected.append(monkey.item_inspected_count)

    sorted_items = sorted(items_inspected)
    answer = sorted_items[-1] * sorted_items[-2]

    return answer


def runPart1(filename):
    monkeys = parse_monkey_file(filename)
    return evaluate_monkeys(monkeys, 20, True)


def runPart2(filename):
    monkeys = parse_monkey_file(filename)
    return evaluate_monkeys(monkeys, 10000, False)


def main():
    print("Part 1: ")
    print(runPart1('full-data.txt'))
    print("Part 2: ")
    print(runPart2('full-data.txt'))


if __name__ == '__main__':
    main()
