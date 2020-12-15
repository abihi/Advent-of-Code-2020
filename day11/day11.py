from copy import deepcopy

def padded_row(line, pad_symbol='.'):
    row = [c for c in line if c != '\n']
    row.insert(0, pad_symbol)
    row.append(pad_symbol)
    return row

def generate_board(filename):
    board = []
    with open(filename) as file:
        line = file.readline()
        board.append(['.'] * (len(line)+1)) 
        board.append(padded_row(line))
        for line in file:
            board.append(padded_row(line))
        board.append(['.'] * len(board[0]))
    return board

def get_neighbors(board, i, j):
    top = board[i-1][j-1] + board[i-1][j] + board[i-1][j+1]
    mid = board[i][j-1]   + board[i][j+1]   
    bot = board[i+1][j-1] + board[i+1][j] + board[i+1][j+1]
    return top + mid + bot

def rules(board, next_board):
    changed = False
    
    for i in range(1, len(board)-1):
        for j in range(1, len(board[0])-1):
            neighbors = get_neighbors(board, i, j)
            if board[i][j] == 'L' and '#' not in neighbors: 
                next_board[i][j] = '#'
                changed = True
            elif board[i][j] == '#' and neighbors.count("#") >= 4:
                next_board[i][j] = 'L'
                changed = True

    return next_board, changed

def count_occupied_seats(board):
    occupied = 0
    for i in range(len(board)):
        for j in range(len(board[0])):
            if board[i][j] == "#":
                occupied += 1
    return occupied

def main():
    board = generate_board("day11.in")

    while True:
        next_board = deepcopy(board)
        board, changed = rules(board, next_board)
        if not changed:
            break
    
    print("P1:", count_occupied_seats(next_board))

main()