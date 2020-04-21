def main():
    while True:
        try:
            a, operator, b = input().split()
            a = int(a)
            b = int(b)

            if operator == '==':
                print(int(a == b))
            elif operator == '!=':
                print(int(a != b))
            elif operator == '>=':
                print(int(a >= b))
            elif operator == '<=':
                print(int(a <= b))
        except:
            break


if __name__ == '__main__':
    main()
