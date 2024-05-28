# Alphamirror

This is program that takes a string as argument and displays the string after replacing each alphabetical character with the opposite alphabetical character.

* The case of the letter remains unchanged, for example :

* 'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

* The final result will be followed by a newline ('\n').

* If the number of arguments is different from 1, the program prints a new line.

# Usage

* Input 1

```bash
go run . "abc" | cat -e
```

* Output 1

```bash
zyx$
```

* Input 2

```bash
go run . "My horse is Amazing." | cat -e
```

* Output 2

```bash
Nb slihv rh Znzarmt.$
```