# Coursera - Intro to Golang - Practice Task 1

This repository documents my progress on the first practice task from the Coursera "Introduction to Golang" course. The objective of this task was to take a non-functional Go program, identify the errors, and implement the necessary fixes to make it run as intended.

---

## 🎯 Project Goal

The original program was a simple command-line application designed to:
1.  Prompt the user to enter their name.
2.  Read the user's input.
3.  Print a personalized welcome message to the console.

---

## 🐞 The Errors & The Fixes

The initial code contained several issues that prevented it from compiling and running correctly. This exercise was a practical application of basic Go syntax and project setup.

The issues I identified and fixed were:

1.  **`package greeter` vs. `package main`**: The entry point of an executable Go program must be in the `main` package. I corrected the package declaration from `greeter` to `main`.
2.  **Incorrect `fmt.Scanln()` Usage**: The `fmt.Scanln` function requires a pointer to the variable where it should store the input. The original code passed `userName` by value. I fixed this by passing a pointer: `&userName`.
3.  **Incorrect Main Function Name**: The function that serves as the entry point for the program must be named `main` (lowercase). The original function was named `Main`, which I corrected.
4.  **Go Version Update**: Updated the `go.mod` file to use Go version `1.25` to ensure the project uses up-to-date language features and dependencies.

---

## 🎓 Course Context

This project is a supplementary practice task from a Coursera course, designed to test problem-solving and debugging skills. It is separate from the main course project, which I am tracking in another repository.

- **Course Link:** [Golang for Beginners: Data Types, Functions, and Packages](https://www.coursera.org/learn/golang-beginners-data-types-functions-packages/)

This repository serves as a personal record of my coursework and a demonstration of my ability to debug and understand fundamental Go concepts.
