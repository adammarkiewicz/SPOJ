from bisect import bisect_left


def main():
    input_lines_count = int(input())
    for _ in range(0, input_lines_count):
        inpt = map(int, input().split())
        numbers_count = next(inpt)
        numbers = list(inpt)
        numbers.sort()
        avarage = sum(numbers) / numbers_count
        result = closest(numbers, avarage)
        print(result)


def closest(myList, myNumber):
    pos = bisect_left(myList, myNumber)
    if pos == 0:
        return myList[0]
    if pos == len(myList):
        return myList[-1]
    before = myList[pos - 1]
    after = myList[pos]
    if after - myNumber < myNumber - before:
        return after
    else:
        return before


if __name__ == "__main__":
    main()
