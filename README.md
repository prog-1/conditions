# Conditions (test 2)

## Programs

### Sorting (2 points)

Write a program in the directory `sorting`.

Sort three numbers that are provided as the input.

Example:

```
The program sorts three numbers.
Enter the numbers:
5 -2 0
Numbers in a sorted order: -2 0 5
```

### Speed conversions (4 points)

Write a program in the directory `speed`.

Write a program that performs `km/h <-> m/s` speed conversion. The program
should show a selection menu, e.g.

```
This program performs km/h <-> m/s speed conversion. Please enter
(1) if you want to convert km/h to m/s;
(2) if you want to convert m/s to km/h.
```

Example:

```
This program performs km/h <-> m/s speed conversion. Please enter
(1) if you want to convert km/h to m/s;
(2) if you want to convert m/s to km/h.
2
Enter the speed in m/s:
100
The speed in km/h is 360
```

### Linear inequality (6 points)

Write a program in the directory `linear`.

Write a program that solves a linear inequality: `Ax + B > 0` (`A` and `B` are
provided as the input).

Example:

```
The program solves a linear inequality Ax + B > 0.
Enter A and B:
2 -6
x > 3
```

### Quadratic equation (6 points)

Write a program in the directory `quadratic`.

Write a program that solves a quadratic equation: `Ax^2 + Bx + C = 0`. The
program should correctly handle cases when there are two solutions, one
solution, zero solutions, A = 0.

Example:

```
The program finds solutions of quadratic equation.
Enter the coefficients A, B and C:
1 -4 4
x1 = x2 = 2
```

### Leap year (6 points)

Write a program in the directory `leap`.

Write a program that determines whether a given year is a leap year. All years
which are perfectly divisible by 4 are leap years except for century years
(years ending with 00) which is leap year only it is perfectly divisible by
400.

Examples:

```
The program determines if a year is a leap year.
Enter the year:
1900
1900 is a not leap year.
```

```
The program determines if a year is a leap year.
Enter the year:
2012
2012 is a leap year.
```

```
The program determines if a year is a leap year.
Enter the year:
2021
2021 is a not leap year.
```

### Triangle Area (6 points)

Write a program in the directory `area`.

Given three numbers, which denote triangle side lengths, as an input, write a
program that determines if a triangle with such side lengths exists. If yes,
calculate the area of the triangle (you may use [Heron’s
formula](https://en.wikipedia.org/wiki/Heron%27s_formula)).

Examples:

```
Enter three triangle sides:
4 5 10
A triangle with such sides doesn't exist.
```

```
Enter three triangle sides:
4 5 6
Triangle area: 9.921567416492215
```

### Brick (10 points)

Write a program in the directory `brick`.

You are given a hole of a rectangular shape defined by dimensions `A x B` (`A`
and `B1 are provided as an input), and a brick with dimensions `P x Q x R`
(`P`, `Q` and `R` are provided as an input). The brick can be rotated to fit in
the hole, but it must stay perpendicular to the hole surface. Write a program
that determines if the brick can be inserted in the hole.

Examples:

```
Enter the hole dimensions (2 numbers):
5 8
Enter the brick dimensions (3 numbers):
9 5 7
The brick can be inserted in the hole!
```

```
Enter the hole dimensions (2 numbers):
5 8
Enter the brick dimensions (3 numbers):
9 6 8
The brick cannot be inserted in the hole.
```

