# GOCalculator

- [GOCalculator](#gocalculator)
  - [Build the calculator](#build-the-calculator)
  - [Use the calculator](#use-the-calculator)


## Build the calculator
You either can run the calculator with
```bash
go run calculator.go
```
or you can build it with this instruction.

First init the repo.
```bash
go mod init calculator
``` 
and build it after with
```bash
go build
```
Now we should have a file named `calculator`. To execute it type
```bash
./calculator
```


## Use the calculator

After executing the calculator you can enter the first number.
```bash
Enter first number: 5 # 5 is the input
```
Then choose between the calculation functions.
```bash
Choose between:
1)Addition
2)Subtraction
3)Multiplication
4)Division
1 # < This is our input
```
We now can enter the second number
```bash
Enter second number: 2 # 2 is the input
```
Now it should give the right output.
```bash
7 # 7 is the return value
```