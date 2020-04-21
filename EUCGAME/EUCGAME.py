def main():
    input_lines_count = int(input())
    for _ in range(0, input_lines_count):
        a, b = map(int, input().split())
        while a != b:
            if a < b:
                b = b - a
            else:
                a = a - b
        print(a + b)


if __name__ == "__main__":
    main()
