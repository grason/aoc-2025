file_path = 'input.txt'

zero_count = 0
dial = 50

with open(file_path, 'r', encoding='utf-8') as file:
    for line in file:
        number = int(line[1:])
        if line[0] == 'L':
            dial = dial - number
        if line[0] == 'R':
            dial = dial + number
        dial = dial % 100
        if dial == 0:
            zero_count = zero_count + 1;
print (zero_count)