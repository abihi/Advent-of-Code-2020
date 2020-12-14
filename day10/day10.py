def sort(adapters):
    stack = []
    while adapters:
        max_adapter = max(adapters)
        stack.append(max_adapter)
        adapters.remove(max_adapter)
    ordered = []
    while stack:
        ordered.append(stack.pop())
    return ordered

def find_differences(sorted_adapters):
    diff1 = 0
    diff3 = 1
    for i in range(len(sorted_adapters) - 1):
        diff = abs(sorted_adapters[i] - sorted_adapters[i+1])
        if diff == 1:
            diff1 += 1
        elif diff == 3:
            diff3 += 1
    return diff1, diff3

def count_combinations(mem, sorted_adapters, current):
    count = 0

    for i in [1, 2, 3]:
        if sorted_adapters[current+i] == sorted_adapters[-1]:
            return 1

        diff = abs(sorted_adapters[current] - sorted_adapters[current+i])
        if diff in [1, 2, 3]:
            combinations = 0
            if mem[current] == 0:
                    combinations += count_combinations(mem, sorted_adapters, current+i)
                    mem[current] = combinations
            else:
                combinations = mem[current]
            count += combinations
        
    return count

def main():
    adapters = [0] 
    with open("day10.in") as file:
        for line in file:
            adapters.append(int(line))
    sorted_adapters = sort(adapters)
    diff1, diff3 = find_differences(sorted_adapters)
    print("p1", diff1 * diff3)

    sorted_adapters.append(sorted_adapters[-1] + 3)
    # Memoization, store computations outside recursive function
    mem = [0] * len(sorted_adapters)
    c = count_combinations(mem, sorted_adapters, 0)
    print(mem)
    print("p2", c)

main()