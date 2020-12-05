def find_row(line):
    binary_string = line.replace('B', '1').replace('F', '0')
    return int(binary_string, 2)

def find_col(line):
    binary_string = line.replace('R', '1').replace('L', '0')
    return int(binary_string, 2)

def calculate_seat_ID(row, col):
    return row * 8 + col

def main():
    highest_seat_ID = 0
    max_seat_id = 127*8+7
    occupied_seats = [0] * max_seat_id
    with open("day5.in") as file:
        for line in file:
            row = find_row(line[:7])
            col = find_col(line[7:])
            seat_ID = calculate_seat_ID(row, col)
            occupied_seats[seat_ID] = seat_ID
            if seat_ID > highest_seat_ID:
                highest_seat_ID = seat_ID
    print("Highest seat ID:",highest_seat_ID)
    for i in range(1, len(occupied_seats[:highest_seat_ID+1])):
        if occupied_seats[i] == 0 and occupied_seats[i-1] > 0 and occupied_seats[i+1] > 0:
            print("My seat ID:", i)


main()