# Reading in Full Sentences

We’ll use Go’s while loop equivalent of a for loop without any parameters to ensure our program continues on forever. In this example every time text is entered and then enter is pressed, we assign text to equal everything up to and including the \n special character. If we want to do a comparison on the string that has just been entered then we can use the strings.Replace method in order to remove this trailing \n character with nothing and then do the comparison.
