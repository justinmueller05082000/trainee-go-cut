# trainee-cut

## About

**trainee-cut** is a from scratch developed program which cuts text files. There are
various parameters you must use to filter or individualize the output of your to cut file.
Seen like this it's an imitation of the original cut command line command.
Note: cut does not work without parameters.

## Parameters

| Parameter    | Usage | 
| ------------- |-------------|
| **-f**      | Choose a field or a range of fields to print.
| **-c**      | Choose a character or a range of characters to print.  
| **-d**  | Choose a separator that your file should be split to. Can only be used with **-f**    
| **-s**     | Suppress lines with no field delimiter characters. Unless specified, lines with no delimiters are passed through unmodified.

## Usage Examples

In the following you'll see a quick example of how to use the cut command. Here is what our text file looks like:

toCut.txt:

               one	two	three	four
               uno	deus	trois	quattre
               linux:amd64-example
               windows:amd-example-2

Now we can cut out the fields with `./cut.linux-amd64 -f 1 < toCut.txt `  

The output is:

                one
                uno
                linux:amd64-example
                windows:amd-example-2
               
Instead of using a single value you can also define a range for your parameters.
To achieve your command should look like this: `./cut.linux-amd64 -f 1-3 < toCut.txt ` 

The output is:

                one	two	three
                uno	deus	trois
                linux:amd64-example
                windows:amd-example-2

There is also the possibility to combine a range and a single value.
Example: `./cut.linux-amd64 -f 1,2-5 < toCut.txt ` 

The output is:

                one	two	three	four
                uno	deus	trois	quattre
                linux:amd64-example
                windows:amd-example-2
                
If you insert values that are out of range cut just does not print those values.