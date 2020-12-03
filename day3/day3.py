def count_trees(right, down):
    trees = 0
    with open('day3.in') as file:
        if down < 2:
            first_line = file.readline()
        current_pos = 0
        for line in file.read().split("\n")[::down]:
            current_pos = (current_pos + right) % len(line)
            if line[current_pos] == '#':
                trees += 1
    return trees

def main():
    trees11 = count_trees(1, 1)
    trees31 = count_trees(3, 1)
    trees51 = count_trees(5, 1)
    trees71 = count_trees(7, 1)
    trees12 = count_trees(1, 2)
    answer  = trees11 * trees31 * trees51 * trees71 * trees12
    print(answer)

main()