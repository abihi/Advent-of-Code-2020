def has_sum(number, numbers):
    for x in numbers:
        for y in numbers:
            if x+y == number:
                return True
    return False

def find_first_weakness(numbers, preamble):
    i = preamble
    while i < len(numbers):
        if not has_sum(numbers[i], numbers[i-preamble:i]):
            return numbers[i]
        i += 1

def contiguous_set(p1, numbers):
    for i in range(len(numbers)):
        set_sum = 0
        c_set = list()
        for num in numbers[i:]:
            set_sum += num
            c_set.append(num)
            if set_sum == p1:
                return min(c_set) + max(c_set)

def main():
    numbers = list()
    with open('day9.in') as file:
        for line in file:
            numbers.append(int(line))
    p1 = find_first_weakness(numbers, 25)
    print("p1", p1)
    p2 = contiguous_set(p1, numbers)
    print("p2", p2)

main()