example_input = [int(x) for x in "1721 979 366 299 675 1456".split(" ")]

def find_2020(number_list):
    for x in number_list:
        for y in number_list:
            for z in number_list:
                if x + y + z == 2020:
                    return x*y*z

def process_input(filename):
    number_list = list()
    with open(filename) as input_file:
        for line in input_file:
            number_list.append(int(line))
    return number_list

number_list = process_input('day1.in')
print(find_2020(number_list))