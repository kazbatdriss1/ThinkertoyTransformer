# ThinkertoyTransformer

## Description

**ThinkertoyTransformer** is a simple yet powerful Go-based tool designed to read a text file, transform its content by replacing specific characters, and save the transformed text to a new file. This utility is useful for anyone looking to quickly and efficiently modify text files by applying a series of predefined character substitutions.

### Key Features:
- **File Reading and Writing:** Reads content from an input file and writes the transformed content to an output file.
- **Character Substitution:** Utilizes Go's `strings.NewReplacer` to define and apply a series of character replacements.
- **Error Handling:** Basic error handling to manage file read and write operations.

### How to Use:
1. **Input File:** Ensure you have an input file named `thinkertoy.txt` in the same directory as the program.
2. **Run the Program:** Execute the program using Go.
3. **Output File:** The transformed content will be saved to `resul.txt`.

### Example
If the content of `thinkertoy.txt` is:



      
 | |  
-O-O- 
 | |  
-O-O- 
 | |  
      
      


The content of resul.txt will be:


      
 I I  
-A-A- 
 I I  
-A-A- 
 I I  
      

