import itertools

parties = {'Smer': 33, 'Hlas': 32, 'PS': 23, 'Republika': 18, 'Sme_Rodina': 14, 'SaS': 10, 'KDH': 11, 'Demokrati':9}

exclusions = [('Smer', 'SaS'), ('Smer','PS'),('PS', 'Republika'),('SaS', 'Republika'),('KDH', 'Republika'),('Demokrati', 'Republika'),('Demokrati', 'Smer'),('Demokrati', 'Hlas')]

combinations = []
for r in range(1, len(parties) + 1):
    for c in itertools.combinations(parties.items(), r):
        if sum(party[1] for party in c) > 75 and not any((excl[0] in dict(c)) and (excl[1] in dict(c)) for excl in exclusions):
            combinations.append(c)

print(f"Number of combinations with sum greater than 75 and no exclusions: {len(combinations)}")

for c in combinations:
  print (sum(party[1] for party in c))
  #names = [party[0] for party in c]  # Extract the names from the combination tuple
  #print(names)
  print(c)