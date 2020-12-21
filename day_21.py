import re

with open('day_21_1_input.txt', 'r') as f:
    possible_allergens = {}
    all_ingredients = set()

    for line in f:
        i, a = line.strip().strip(')').split(' (contains ')
        i = set(i.split(' '))
        all_ingredients = all_ingredients.union(i)
        for allergen in a.split(', '):
            if allergen in possible_allergens:
                possible_allergens[allergen] = possible_allergens[allergen].intersection(i)
            else:
                possible_allergens[allergen] = i
    print(possible_allergens)
    print(all_ingredients)
    allergens = {}
    while len(allergens) < len(possible_allergens):
        for allergen, ingredient_set in possible_allergens.items():
            if len(ingredient_set) == 1:
                # Get element from set
                for ingredient in ingredient_set:
                    break
                allergens[allergen] = ingredient
                for other in possible_allergens:
                    if ingredient in possible_allergens[other]:
                        possible_allergens[other].remove(ingredient)
    print(allergens)

    safe_ingredients = all_ingredients.difference(allergens.values())
    print(safe_ingredients)

    f.seek(0)
    text = f.read()
    safe_count = sum(len(re.findall(r"\b{}\b".format(ingredient), text)) for ingredient in safe_ingredients)
    print('Answer 1: {}'.format(safe_count))
    print('Answer 2: {}'.format(','.join(allergens[k] for k in sorted(allergens))))
