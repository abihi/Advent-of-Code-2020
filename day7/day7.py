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
                bag = Bag(bag_name, [])
                bags_dict[bag_name] = bag
                continue
            holds = list()
            capacity1 = [int(capacity_split[0][1]), capacity_split[0][3:]] 
            holds.append(capacity1)
            if len(capacity_split) > 1:
                capacity2 = [int(capacity_split[1][1]), capacity_split[1][3:]]
                holds.append(capacity2)
            bag = Bag(bag_name, holds)
            bags_dict[bag_name] = bag
    return bags_dict

def can_hold_shiny_bag(key, bags_dict):
    if "shiny gold" in bags_dict[key].holds:
        return False

    for bag in bags_dict[key].holds:
        if "shiny gold" in bag[1]:
            return True

        if can_hold_shiny_bag(bag[1], bags_dict):
            return True
    return False

def main():
    bags_dict = fill_bags_dict('day7.in')
    answer = 0

    for key in bags_dict.keys():
        if can_hold_shiny_bag(key, bags_dict):
            answer += 1
    
    print(answer)
   
main()