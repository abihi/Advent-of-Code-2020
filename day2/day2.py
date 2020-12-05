example_input = ["1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"]

def preprocess_line(line):
    split_line = line.split(" ")
    low, high = map(int, split_line[0].split("-"))
    letter = split_line[1][0]
    password = split_line[2]
    return low, high, letter, password

def is_valid_password_part2(low, high, letter, password):
    valid = False
    if password[low-1] == letter: 
        valid = True
    if password[high-1] == letter: 
        valid = not valid
    return valid

def main():
    valid_passwords_p1 = 0
    valid_passwords_p2 = 0
    with open("day2.in") as file:
        for line in file:
            low, high, letter, password = preprocess_line(line)
            valid_passwords_p1 += low <= password.count(letter) <= high
            if is_valid_password_part2(low, high, letter, password):
                valid_passwords_p2 += 1
    print(valid_passwords_p1)
    print(valid_passwords_p2)
main()