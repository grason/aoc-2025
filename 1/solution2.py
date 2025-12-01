file_path = 'input.txt'

clicks = 0
dial = 50

with open(file_path, 'r', encoding='utf-8') as file:
    for line in file:
        mod = 0
        if line[0] == 'L':
            mod = -1
        else:
            mod = 1
        number = int(line[1:])
        while number > 0:
            dial += mod
            number -= 1
            dial = dial % 100
            if dial == 0:
                clicks += 1
print (clicks)