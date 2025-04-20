# CSV to JSON Converter for Assignment 3

This Go application has been created to read CSV data and convert it into JSON (line) format. Each line of the json file will correspond to its row in the csv file.


## Installation and Execution

1. Clone the repository to your local machine:

   git clone https://github.com/<yourusername>/csv2jsonlines.git (in my case, username is 'carolinemaciag')
   cd csv2jsonlines

2. If you do not have Go's dependency management tool, you can install it with the command:
    
    go mod tidy

3. Prepare your CSV file, for this project, we used "housesInput.csv"

4. Run the command below, which will:
    - Read the input file (housesInput.csv.)
    - Convert its data into JSON Lines format.
    - Save the output as 'housesOutput.jl'

    go run main.go housesInput.csv housesOutput.jl

5. The output JSON lines file should look like the following:
{"value":452600,"income":8.3252,"age":41,"rooms":880,"bedrooms":129,"pop":322,"hh":126}
{"value":358500,"income":8.3014,"age":21,"rooms":7099,"bedrooms":1106,"pop":2401,"hh":1138}
{"value":352100,"income":7.2574,"age":52,"rooms":1467,"bedrooms":190,"pop":496,"hh":177}

### Running Unit Tests:

To ensure everything works as it should, run the following test command:

    go test ./...

If the test ran successfully, you should see an output similar to the follwing:

    ok  	csv2jsonlines/converter	0.214s

### AI Assistant used for the assignment:
ChatGPT 4.0 turbo model was used mainly in troubleshooting syntax for this assignment. The GPT model also greatly assisted in optimizing certain sections of code, such as enhancing the error handling in the 'rowtohouse' functions. 

The AI model also provided useful insights on how to improve and structure the unit tests for this project, and again provided root causes for syntax errors within this portion of the assignment. 
