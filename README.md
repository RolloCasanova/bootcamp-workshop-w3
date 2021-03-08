# Wizeline Go Bootcamp - Workshop 3

We are going to put in practice what you learned in `workshop #3`:

- `Pointers`
- `Methods`
- `Interfaces`
- `Packages`

## The challenge

You should do all the following steps:

1. Create the following structs types inside the `shape` package; each one of them with the required values to calculate both `area` and `perimeter`:
   1. `Square`
   2. `Circle`
   3. `Hexagon`

2. Add the following methods to the `Shape interface`; you are not allowed to modify them in any way:
   - `Area() float64`
      - Will return shape's area
   - `Perimeter() float64`
      - Will return shape's perimeter
   - `Info() string`
      - Will display a message with the shape's `area` and it's `perimeter`

3. Add methods to all shapes types from step #1 so all of them implements the `Shape interface`:
   - Add the required logic to these methods' body
   - Make use of already defined `constants` and `methods` from useful package (e.g. from `math`, `math.Pi`, `math.Pow()`, etc.)

4. Initialize a variable of each shape from the step #1 with a positive side/radius value, and print shape's info (`Info() method`) by using the `printShape function`, which again you are not allowed to modify it's signature

## Extra stuff (optional)

- You have in mind to add new functionality to let you change any shape size, how can you do that?
  - You are not allowed to create a new shape to replace the current one

- Let's say you wanted to add other values (e.g. Shape type, Nickname, etc.) to all your shapes; the first idea to attack this problem might be to add each one of these fields in each of the shape's structures, but you will be repeating yourself over and over in all shapes.
These will be values that all shapes share, do you have something in mind you could use to only have a single place to add these values and avoid repetition?
  - You also want to add a method to allow change these values, how can you do that?
  - Don't forget to update the Info() method to display the new information
