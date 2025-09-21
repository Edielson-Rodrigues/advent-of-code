## \--- Day 3: Think about ---

"Our computers are having problems, so I have no idea if we have any chief historian in stock\! You can check the warehouse, however," says the slightly disturbed shopkeeper from the North Pole Tobgan Rental Shop. The historians go out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having problems again?"

The computer seems to be trying to run a program, but its memory (your puzzle input) is *corrupted*. All the instructions have been scrambled\!

It seems the program's goal is just to *multiply some numbers*. It does this with instructions like `mul(X,Y)`, where `X` and `Y` are 1 to 3 digit numbers. For example, `mul(44,46)` multiplies `44` by `46` to get a result of `2024`. Similarly, `mul(123,4)` would multiply `123` by `4`.

However, because the program's memory has been corrupted, there are also many invalid characters that must be *ignored*, even if they look like part of an instruction. Sequences like `mul`, `mul(4*`, `mul(6,9!`, `?(12,34)`, or `mul ( 2 , 4 )` do *nothing*.

For example, consider the following section of corrupted memory:

```
x*mul(2,4)*%&mul[3,7]!@^do_not_*mul(5,5)*+mul(32,64]then(*mul(11,8)mul(8,5)*)
```

Only the four highlighted sections are real `mul` instructions. The sum of the result of each instruction produces **161** (`2*4 + 5*5 + 11*8 + 8*5`).

Examine the corrupted memory for uncorrupted `mul` instructions. *What do you get if you add up all the multiplication results?*

-----

## \--- Part Two ---

As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact. If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more accurate result.

There are two new instructions you'll need to handle:

  * The `do()` instruction *enables* future `mul` instructions.
  * The `don't()` instruction *disables* future `mul` instructions.

Only the *most recent* `do()` or `don't()` instruction applies. At the beginning of the program, `mul` instructions are *enabled*.

For example:

```
x*mul(2,4)*&mul[3,7]!^*don't()*_mul(5,5)+mul(32,64](mul(11,8)un*do()*?*mul(8,5)*)
```

This corrupted memory is similar to the example from before, but this time the `mul(5,5)` and `mul(11,8)` instructions are *disabled* because there is a `don't()` instruction before them. The other `mul` instructions function normally, including the one at the end that gets re-*enabled* by a `do()` instruction.

This time, the sum of the results is `48` (`2*4 + 8*5`).

Handle the new instructions; *what do you get if you add up all of the results of just the enabled multiplications?*