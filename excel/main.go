package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
    "github.com/go-gota/gota/dataframe"
)

func main() {
    // Open the Excel file
    f, err := excelize.OpenFile("popul.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // defer func() {
    //     // Close the spreadsheet.
    //     if err := f.Close(); err != nil {
    //         fmt.Println(err)
    //     }
    // }
    
    // Get the name of the first sheet in the Excel file
    sheetName := f.GetSheetName(0)
    
    // Get all the rows in the first sheet
    rows, err := f.GetRows(sheetName)
    if err != nil {
        fmt.Println(err)
        return
    }

     // Create a slice of slices to hold the data
     var data [][]string

     // Loop through each row in the sheet
     for _, row := range rows {
         var rowData []string
         // Loop through each cell in the row
         for _, colCell := range row {
             rowData = append(rowData, colCell)
         }
         data = append(data, rowData)
     }

     // Create a Gota DataFrame from the slice of slices
     df := dataframe.LoadRecords(data)
 
     // Print the Gota DataFrame
     fmt.Println(df)
}


