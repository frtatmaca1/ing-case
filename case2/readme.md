# Case 2 Api

## About

This project was developed by Go Language.

## Usage

You can directly run /main.go file to use program.

## Run

You can run the project via command:

```Shell
go run main.go
```
## Testing

This project has integration tests in files.

You can test the code via the files.

## About me

I am Fırat Atmaca.

I have been working on software projects since 2013.

You can contact with me via [Linkedin](https://www.linkedin.com/in/firat-atmaca-469b2769/)

## Task - Anagram matching
Write a program that takes as argument the path to a file containing one word per line, groups the words that are anagrams to each other, and writes to the standard output each of these groups.
The groups should be separated by new lines, and the words inside each group by commas.

Assume that the words in the file are ordered by size.

NOTES:

* The code should compute the anagrams __without the help of any library__
* You are allowed to use any library for any other aspect functional to the task (e.g.: handling the CLI, testing, I/O)
* If CLI is not available in your development stack (e.g.: mobile) you can rely on automated tests and/or UI to feed your program with the input data and verify the output.

Clarifications:

* The order of the groups in the output does not matter
* The words in the input file are sorted by size
* The files provided in the `data` folder are just sample input data to help you reason about the problem
* If a word has no anagram, don't print it

Example:

```bash
command_to_run_your_program task_2/data/example.txt
```

Output:
```text
abc,bac,cba
unf,fun
```