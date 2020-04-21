def main():
    input_lines_count = int(input())
    for _ in range(0, input_lines_count):
        hour, minute = map(int, input().split(":"))
        for _ in range(24):
            for _ in range(60):
                if minute == 59:
                    minute = 0

                    if hour == 23:
                        hour = 0
                    else:
                        hour += 1
                else:
                    minute += 1

                hour_str = str(hour).zfill(2)
                minute_str = str(minute).zfill(2)

                string = "".join((hour_str, minute_str)).lstrip("0")

                if is_palindrom(string):
                    result = ":".join((hour_str, minute_str))
                    print(result)
                    break
            else:
                continue  # executed if the loop ended normally (no break)
            break


def is_palindrom(string):
    reversed_string = string[::-1]
    if string == reversed_string:
        return True
    else:
        return False


if __name__ == "__main__":
    main()
