def main():
    input_lines_count = int(input())
    for _ in range(0, input_lines_count):
        class_count, sweet_count = map(int, input().split())
        class_count -= 1
        if class_count == 0:
            if sweet_count > 0:
                print("TAK")
            else:
                print("NIE")
        elif sweet_count % class_count == 0:
            print("NIE")
        else:
            print("TAK")


if __name__ == "__main__":
    main()
