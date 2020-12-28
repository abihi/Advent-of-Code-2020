def consume(x, rule_number, rules_map):
    rule = rules_map[rule_number]
    if rule[0] == '"':
        rule = rule.strip('"')
        if x.startswith(rule):
            return [len(rule)]
        else:
            return []
    
    acc_chains = []
    for opt in rule.split(" | "):
        acc_chain = [0]
        for rule_number in opt.split(" "):
            acc = []
            rule_number = int(rule_number)
            for ac in acc_chain:
                ret = consume(x[ac:], rule_number, rules_map)
                for c in ret:
                    acc.append(c + ac)
            acc_chain = acc
        acc_chains += acc_chain
    return acc_chains            

def main():
    rules, tests = open("day19.in").read().strip().split('\n\n')

    rules_map = {}
    for r in rules.split("\n"):
        rule_number, val = r.split(": ")
        if rule_number == "8":
            val = "42 | 42 8"
        if rule_number == "11":
            val = "42 31 | 42 11 31"
        rules_map[int(rule_number)] = val

    p2 = 0
    for x in tests.split('\n'):
        p2 += len(x) in consume(x, 0, rules_map)
    print("P2:", p2)

main()