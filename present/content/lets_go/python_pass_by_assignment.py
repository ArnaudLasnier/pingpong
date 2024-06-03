def double(x):
    x *= 2

a = 1
b = double(a)


def uppercase(names: list[str]):
    for name in names:
        name = name.upper()

n = ["Arnaud", "Lasnier"]
uppercase(n)


def change_first_element(list_):
    list_[0] = 42

l = [1, 2, 3]
change_first_element(l)
print(l)

t = (1, 2, 3)
change_first_element(t)
print(l)
