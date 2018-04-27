package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"strconv"
)

var col = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO",
	"AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ"}

func create() {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("会计科目")
	// Set value of a cell.
	xlsx.SetCellValue("会计科目", "1", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("E:\\Accounting\\total.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func read() {
	xlsx, err := excelize.OpenFile("E:\\Accounting\\1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("会计科目")
	for _, row := range rows {
		for index, colCell := range row {
			fmt.Print(colCell, index, "\t")
		}
		fmt.Println()
	}
}

func ReadAndWrite() {
	total := excelize.NewFile()
	// Create a new sheet.
	index := total.NewSheet("会计科目")
	// Set value of a cell.
	//total.SetCellValue("会计科目", "B2", 100)
	// Set active sheet of the workbook.
	total.SetActiveSheet(index)

	array, num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 0
	//array, num := []int{1, 2}, 0
	for i, n := range array {
		filename := "E:\\Accounting\\" + strconv.Itoa(n) + ".xlsx"
		xlsx, err := excelize.OpenFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cell := xlsx.GetCellValue("会计科目", "B2")
		fmt.Println(cell)
		// Get all the rows in the Sheet1.
		rows := xlsx.GetRows("会计科目")
		if num > 0 {
			num -= 3
		}
		for j, row := range rows {
			if i > 0 && j < 4 {
				continue
			}
			for index, colCell := range row {
				axis := col[index] + strconv.Itoa(num+j+1)
				//fmt.Println(axis)
				total.SetCellValue("会计科目", axis, colCell)
			}
		}
		num += len(rows)
	}
	err := total.SaveAs("E:\\Accounting\\totalWithBlank.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

//func main() {
//	//read()
//	//create()
//	ReadAndWrite()
//}
