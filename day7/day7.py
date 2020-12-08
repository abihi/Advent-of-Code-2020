class Bag():
    def __init__(self, name, holds):
        self.name = name
        self.holds = holds

def fill_bags_dict(filename):
    bags_dict = {}
    with open(filename) as file:
        for line in file:
            contain_split = line.replace(".", "").replace("\n", "").split('contain')
            capacity_split = contain_split[1].split(',')
            bag_name = contain_split[0][:-1]
            if "no other bags" in capacity_split[0]:
                bags_dict[bag_name] = Bag(bag_name, [])
                continue
            holds = list()
            for c in capacity_split:
                holds.append([int(c[1]), c[3:]])
            bags_dict[bag_name] = Bag(bag_name, holds)
    return bags_dict

def can_hold_shiny_bag(key, bags_dict):
    for bag in bags_dict[key].holds:
        if "shiny gold" in bag[1]:
            return True
        if can_hold_shiny_bag(bag[1], bags_dict):
            return True
    return False

def bags_in_bag(key, bags_dict):
    count = 0
    for bag in bags_dict[key].holds:
        if bag[0]:
            count += bag[0] + bag[0] * bags_in_bag(bag[1], bags_dict)
    return count

def main():
    answer_p1 = 0
    bags_dict = fill_bags_dict('day7.in')

    for key in bags_dict.keys():
        if can_hold_shiny_bag(key, bags_dict):
            answer_p1 += 1
    
    answer_p2 = bags_in_bag("shiny gold bags", bags_dict)
    print("p1:", answer_p1)
    print("p2:", answer_p2)

main()