# solutions

## task 01

replace `110001` with `110000` + team id

```
grep -F '"110001",' file.csv | cut -d',' -f10,11 | tr -d '"' | tr ',' ' ' | head -n 1
```

## task 02

```
grep '<div class="110027' index.html
```

## task 03

```

```
