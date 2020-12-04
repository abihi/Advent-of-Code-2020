import string

def preprocess_passports(filename):
    with open(filename) as file:
        passports = list()
        passport = ""
        for line in file:
            if line != "\n":
                passport += line.replace('\n', ' ')
            else:
                passports.append(passport)
                passport = ""
        passports.append(passport)
    return passports

def validate_data(passport_dict, keys):
    for key in keys:
        if key not in passport_dict:
            return False
    if int(passport_dict["byr"]) < 1920 or int(passport_dict["byr"]) > 2002:
        return False
    if int(passport_dict["iyr"]) < 2010 or int(passport_dict["iyr"]) > 2020:
        return False   
    if int(passport_dict["eyr"]) < 2020 or int(passport_dict["eyr"]) > 2030:
        return False
    metric = passport_dict["hgt"][-2:]
    height = passport_dict["hgt"][:-2]
    am = ['cm', 'in']
    if metric not in am:
        return False
    if metric == 'cm' and (int(height) < 150 or int(height) > 193):
        return False
    if metric == 'in' and (int(height) < 59 or int(height) > 76):
        return False
    if len(passport_dict["hcl"]) != 7 or passport_dict["hcl"][0] != '#': 
        return False

    hex_digits = string.hexdigits[:16]
    for c in passport_dict["hcl"][1:]:
        if c not in hex_digits:
            return False

    eye_colors = "amb blu brn gry grn hzl oth".split(' ')
    if passport_dict["ecl"] not in eye_colors: 
        return False
    if len(passport_dict["pid"]) != 9: 
        return False
    return True

def valid_passports(passports, keys):
    valid = 0
    for passport in passports:
        pd = {p.split(':')[0]:p.split(':')[1] for p in passport.split(' ') if p != ''}
        if validate_data(pd, keys):
            valid += 1
    return valid

def main():
    passport_keys = 'byr iyr eyr hgt hcl ecl pid'.split(' ')
    passports = preprocess_passports('day4.in')
    answer = valid_passports(passports, passport_keys)
    print(answer)

main()