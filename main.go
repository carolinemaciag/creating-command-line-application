package main

import (
    "fmt"
    "log"
    "os"

    "csv2jsonlines/converter"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: ./main <input.csv> <output.jsonl>")
        os.Exit(1)
    }

    inputPath := os.Args[1]
    outputPath := os.Args[2]

    err := converter.ConvertCSVtoJSONLines(inputPath, outputPath)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Println("Conversion successful!")
}
