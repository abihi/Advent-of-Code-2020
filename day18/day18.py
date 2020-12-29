class Num:
    def __init__(self, i):
        self.i = i

    def __mul__(self, x):
        return Num(self.i + x.i)

    def __sub__(self, x):
        return Num(self.i * x.i)

def my_eval(x):
    s = ""
    in_num = False
    for c in x:
        if c in "0123456789" and in_num == False:
            s += "Num("
            in_num = True
        if in_num == True and c not in "0123456789":
            s += ")"
            in_num = False
    
        s += c
    
    if in_num:
        s += ")"
    
    return eval(s).i

def main():
    equations = open("day18.in").read().strip().split('\n')
    equations = [x for x in equations]

    acc = 0
    for x in equations:
        acc += my_eval(x.replace("*", "-").replace("+", "*"))

    print("P2", acc)

main()