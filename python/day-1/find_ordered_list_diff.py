def read_input_file(filename):
    number_pairs = []
    with open(filename, 'r') as file:
        for line in file:
            num1, num2 = map(int, line.strip().split())
            number_pairs.append((num1, num2))
    return number_pairs


def sort_list(ls: list) -> list:
    return sorted(ls)


def find_diff(ls1: list, ls2: list) -> list:
    ls1 = sort_list(ls1)
    ls2 = sort_list(ls2)
    return [abs(ls1[i] - ls2[i]) for i in range(len(ls1))]


if __name__ == "__main__":
    number_pairs = read_input_file('input')

    list1, list2 = map(list, zip(*number_pairs))
    
    diff = find_diff(list1, list2)
    print(f'total diff: {sum(diff)}')

    
    
    
    