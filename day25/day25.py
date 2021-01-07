import sys

def transform(x, ls):
    return pow(x, ls, 20201227)

card_pubkey = 6270530
door_pubkey = 14540258

dlz = 0
while transform(7, dlz) != door_pubkey:
    dlz += 1

encryption = transform(card_pubkey, dlz)
print(encryption)