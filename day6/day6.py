def groupScoreP1(group):
	scores = [0] * 26
	score = 0
	for c in group:
		asciiC = ord(c)-97
		if scores[asciiC] == 0:
			scores[asciiC] = 1
			score += 1
	return score

def groupScoreP2(group, people):
	scores = [0] * 26
	score = 0
	for c in group:
		asciiC = ord(c)-97
		scores[asciiC] += 1
	for i in range(len(scores)):
		if scores[i] == people:
			score += 1
	return score

def main():
	totalScoreP1 = 0
	totalScoreP2 = 0
	group = ""
	people = 0
	with open("day6.in") as file:
		for line in file:
			if line != "\n":
				group += line.replace("\n", "")
				people += 1
			else:
				totalScoreP1 += groupScoreP1(group)
				totalScoreP2 += groupScoreP2(group, people)
				group = ""
				people = 0
	print("Answer P1:", totalScoreP1)
	print("Answer P2:", totalScoreP2)

main()