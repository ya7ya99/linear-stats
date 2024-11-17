إليك محتوى ملف README.md لبرنامجك:

# Linear Regression and Pearson Correlation Coefficient Calculator

This program reads a dataset from a text file, calculates the **Linear Regression Line** and the **Pearson Correlation Coefficient**, and prints the results in a specific format.

## How It Works

The dataset represents a graph where:
- The x-axis values are the line numbers in the file (0, 1, 2, 3, ...).
- The y-axis values are the actual numbers in the file.

### Output Format

1. **Linear Regression Line:**  
   The equation of the form: `y = mx + b`, where `m` is the slope and `b` is the intercept. The values have 6 decimal places.

2. **Pearson Correlation Coefficient:**  
   A decimal value representing the correlation strength. The value has 10 decimal places.

### Example Output

Linear Regression Line: y = 2.345678x + 1.234567 Pearson Correlation Coefficient: 0.9876543210

---

## Requirements

- Go 1.20 or higher
- A valid dataset file with numerical values, one per line.

---

## Usage

1. Clone the repository or copy the program file.

2. Save the data in a file (e.g., `data.txt`):

189 113 121 114 145 110

3. Compile and run the program:
```bash
go run main.go data.txt

4. The program will print the results in the terminal.




---

Error Handling

The program includes error handling for:

Missing or invalid file input.

Empty files.

Non-numeric data in the file.

Division by zero during calculations.



---

Example

Input (data.txt)

10
20
30
40
50
60

Command

go run main.go data.txt

Output

Linear Regression Line: y = 10.000000x + 10.000000
Pearson Correlation Coefficient: 1.0000000000


---

License

This project is licensed under the MIT License.
# Linear Regression and Pearson Correlation Coefficient Calculator

This program reads a dataset from a text file, calculates the **Linear Regression Line** and the **Pearson Correlation Coefficient**, and prints the results in a specific format.

## How It Works

The dataset represents a graph where:
- The x-axis values are the line numbers in the file (0, 1, 2, 3, ...).
- The y-axis values are the actual numbers in the file.

### Output Format

1. **Linear Regression Line:**  
   The equation of the form: `y = mx + b`, where `m` is the slope and `b` is the intercept. The values have 6 decimal places.

2. **Pearson Correlation Coefficient:**  
   A decimal value representing the correlation strength. The value has 10 decimal places.

### Example Output

Linear Regression Line: y = 2.345678x + 1.234567 Pearson Correlation Coefficient: 0.9876543210

---

## Requirements

- Go 1.20 or higher
- A valid dataset file with numerical values, one per line.

---

## Usage

1. Clone the repository or copy the program file.

2. Save the data in a file (e.g., `data.txt`):

189 113 121 114 145 110

3. Compile and run the program:
```bash
go run main.go data.txt

4. The program will print the results in the terminal.




---

Error Handling

The program includes error handling for:

Missing or invalid file input.

Empty files.

Non-numeric data in the file.

Division by zero during calculations.



---

Example

Input (data.txt)

10
20
30
40
50
60

Command

go run main.go data.txt

Output

Linear Regression Line: y = 10.000000x + 10.000000
Pearson Correlation Coefficient: 1.0000000000


---

License

This project is licensed under the MIT License.