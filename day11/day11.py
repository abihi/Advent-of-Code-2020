from pandas import DataFrame

def padded_row(line, pad_symbol='.'):
    row = [c for c in line if c != '\n']
    row.insert(0, pad_symbol)
    row.append(pad_symbol)
    return row

def generate_board(filename):
    board = []
    with open(filename) as file:
        line = file.readline()
        board.append(['.'] * (len(line)+2)) 
        board.append(padded_row(line))
        for line in file:
            board.append(padded_row(line))
        board.append(['.'] * (len(line)+2)) 
    return board

def get_neighbors(board, i, j):
    top = board[i-1][j-1] + board[i][j-1] + board[i+1][j-1]
    mid = board[i-1][j]   + board[i+1][j]   
    bot = board[i-1][j+1] + board[i][j+1] + board[i+1][j+1]
    return top + mid + bot

def rules(board):
    len_board = len(board)
    new_board = []
    new_board.append(['.'] * (len_board + 2))
    
    for i in range(1, len_board-1):
        row = []
        for j in range(1, len_board-1):
            neighbors = get_neighbors(board, i, j)
            if board[i][j] == 'L' and '#' not in neighbors: 
                row.append('#')
            elif board[i][j] == 'L' and '#' in neighbors: 
                row.append('L')
            elif board[i][j] == '#' and neighbors.count("#") >= 4:
                row.append('L')
            elif board[i][j] == '#' and neighbors.count("#") < 4:
                row.append('#')
            else:
                row.append('.')
        row.insert(0, '.')
        row.append('.')
        new_board.append(row)

    new_board.append(['.'] * (len_board + 2))
    return new_board

def isStable(prev_board, curr_board):
    for i in range(len(prev_board)):
        for j in range(len(prev_board)):
            if prev_board[i][j] != curr_board[i][j]:
                return False
    return True

def get_board(board):
    new_board = []
    for i in range(1, len(board)-1):
        row = []
        for j in range(1, len(board)-1):
            row.append(board[i][j])
        new_board.append(row)
    return new_board

def count_occupied_seats(board):
    occupied = 0
    for i in range(len(board)):
        for j in range(len(board)):
            if board[i][j] == "#":
                occupied += 1
    return occupied

def main():
    prev_board = generate_board("day11_ex.in")
    while True:
        curr_board = rules(prev_board)
        if isStable(get_board(prev_board), get_board(curr_board)):
            break
        print("testing")
        prev_board = curr_board
    
    c = count_occupied_seats(prev_board)
    print(c)

main()