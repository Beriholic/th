package handler

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

func newTable() table.Writer {
	t := table.NewWriter()
	t.SetTitle("Trash List")
	header := table.Row{"ID", "Name", "Path", "DeletionDate"}
	t.AppendHeader(header)
	return t
}
func ShowTable(sortType, sortOrder string) ([]Info, error) {
	info, err := GetTrashList(sortType, sortOrder)

	if err != nil {
		return nil, err
	}
	t := newTable()
	if len(info) == 0 {
		row := table.Row{"Trash is empty", "X", "X", "X"}
		t.AppendRow(row)
		fmt.Println(t.Render())
		return nil, nil
	}

	for id, v := range info {
		rows := table.Row{id, v.fileName, v.fromPath, v.trashTime}
		t.AppendRow(rows)
	}

	fmt.Println(t.Render())
	return info, nil
}
