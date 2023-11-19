allowed = set(['+', '-', '>', '<', '.', '[', ']', ','])

def analyze(tokens):
    stack = []
    data, data_pointer = [0], 0

    inst_pointer = 0

    while inst_pointer < len(tokens):
        token = tokens[inst_pointer]
        
        match token['char']:
            case '+':
                data[data_pointer] += 1
            case '-':
                data[data_pointer] -= 1

            case '>':
                data_pointer += 1
                if data_pointer >= len(data):
                    data.append(0)
            case '<':
                data_pointer -= 1
                if data_pointer < 0:
                    raise Exception('Negative index for data pointer.')

            case '.':
                print(chr(data[data_pointer]), end='')
            case ',':
                data[data_pointer] = int(input('Enter: ')) % 256

            case '[':
                if data[data_pointer] == 0:
                    depth = 1
                    while depth > 0:
                        inst_pointer += 1

                        brace = tokens[inst_pointer]['char']
                        if brace == '[':
                            depth += 1
                        elif brace == ']':
                            depth -= 1

            case ']':
                if data[data_pointer] != 0:
                    depth = 1
                    while depth > 0:
                        inst_pointer -= 1

                        brace = tokens[inst_pointer]['char']
                        if brace == '[':
                            depth -= 1
                        elif brace == ']':
                            depth += 1

        inst_pointer += 1;


def tokenize(text):
    result = []
    index = 0
    for position, char in enumerate(text):
        if char in allowed:
            result.append({
                'char': char,
                'index': index,
                'position': position
            })
            index += 1
    return result


def main():
    with open("sample.txt", "r") as file:
        text = file.read()

    tokens = tokenize(text)
    analyze(tokens)


if __name__ == "__main__":
    main()
