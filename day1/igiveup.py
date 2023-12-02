words = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


def checkForNumber(string):
    substring = string[0:]
    for i in range(len(string)):
        if substring in words:
            return words[substring]
        else:
            substring = substring[1:]
    return "a"

LOWEST = 'a'
def main():
    ans = ""
    with open("day1.txt") as file:
        content = file.read()
    lines = content.strip().split()

    ans = 0
    for line in lines:
        start = LOWEST
        end = LOWEST
        possible_string = ""

        for char in line:
            possible_string += char
            if char.isdigit():
                if start == LOWEST:
                    start = char
                    end = char
                else:
                    end = char
            parsedNumber = checkForNumber(possible_string)
            if parsedNumber != LOWEST:
                if start == LOWEST:
                    start = parsedNumber
                    end = parsedNumber
                else:
                    end = parsedNumber

        ans += int(start + end)

    print(ans)


if __name__ == "__main__":
    main()
