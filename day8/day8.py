def r_inst(q):
    acc = 0
    i = 0
    while i < len(q):
        inst = q[i]
        if inst[0] == "nop":
            inst[2] += 1
            i += 1
        elif inst[0] == "acc":
            inst[2] += 1
            if inst[2] > 1:
                return acc
            acc += inst[1]
            i += 1
        elif inst[0] == "jmp":
            inst[2] += 1
            i += inst[1]
        if inst[2] > 1:
            return acc
    return acc

def main():
    q = []
    with open('day8.in') as file:
        for line in file:
            inst = line.split(" ")
            op   = inst[0]
            arg  = int(inst[1].rstrip("\n"))
            q.append([op, arg, 0])
    print(r_inst(q))

main()