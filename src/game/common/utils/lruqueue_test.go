package utils

import (
	"fmt"
	"testing"
)

func Test_AddQueue(t *testing.T) {
	fmt.Print("Test_AddQueue =============================================== \n")
	queue := NewQueue(5)
	item1 := CreateDataItem("1", 1111)
	item2 := CreateDataItem("2", 1111)
	item3 := CreateDataItem("3", 1111)
	item4 := CreateDataItem("4", 1111)
	item5 := CreateDataItem("5", 1111)
	item6 := CreateDataItem("6", 1111)
	err := queue.AddElement(item1)
	if err != nil {
		t.Fatal()
	}
	err = queue.AddElement(item2)
	if err != nil {
		t.Fatal()
	}
	err = queue.AddElement(item3)
	if err != nil {
		t.Fatal()
	}
	err = queue.AddElement(item4)
	if err != nil {
		t.Fatal()
	}
	err = queue.AddElement(item5)
	if err != nil {
		t.Fatal()
	}
	err = queue.AddElement(item6)
	if err != nil {
		t.Fatal()
	}
	if queue.Size() != queue.max {
		t.Fatal()
	}
}

func Test_DeleteQueue(t *testing.T) {
	fmt.Print("Test_DeleteQueue =============================================== \n")
	queue := NewQueue(5)
	item1 := CreateDataItem("1", 1111)
	item2 := CreateDataItem("2", 1111)
	item3 := CreateDataItem("3", 1111)
	item4 := CreateDataItem("4", 1111)
	item5 := CreateDataItem("5", 1111)
	item6 := CreateDataItem("6", 1111)
	queue.AddElement(item1)
	queue.AddElement(item2)
	queue.AddElement(item3)
	queue.AddElement(item4)
	queue.AddElement(item5)
	queue.AddElement(item6)

	if queue.Size() != queue.max {
		fmt.Printf("Size :%d ==== max:%d \n", queue.Size(), queue.max)
		t.Fatal()
	}

	if err := queue.DelElement("1"); err == nil {
		fmt.Print("delete error" + "1 " + "\n")
		t.Fatal()
	}

	if err := queue.DelElement("10"); err == nil {
		fmt.Print("delete error" + "10 \n")
		t.Fatal()
	}

	if err := queue.DelElement("2"); err != nil {
		fmt.Print("delete error" + "2 \n")
		t.Fatal()
	}

	if queue.Size() == queue.max {
		fmt.Printf("Size :%d ==== max:%d \n", queue.Size(), queue.max)
		t.Fatal()
	}
}

func Test_UpdateQueue(t *testing.T) {
	fmt.Print("Test_UpdateQueue =============================================== \n")
	queue := NewQueue(5)
	item1 := CreateDataItem("1", 1111)
	item2 := CreateDataItem("2", 1111)
	item3 := CreateDataItem("3", 1111)
	item4 := CreateDataItem("4", 1111)
	item5 := CreateDataItem("5", 1111)
	item6 := CreateDataItem("6", 1111)
	queue.AddElement(item1)
	queue.AddElement(item2)
	queue.AddElement(item3)
	queue.AddElement(item4)
	queue.AddElement(item5)
	queue.AddElement(item6)

	if queue.Size() != queue.max {
		fmt.Printf("Size :%d ==== max:%d \n", queue.Size(), queue.max)
		t.Fatal()
	}

	item5.Data = 1225554
	if err := queue.UpdateElement(item5); err != nil {
		fmt.Print("delete error" + "10 \n")
		t.Fatal()
	}

	if item := queue.GetElement(item5.Key); item == nil || item.Data != 1225554 {
		fmt.Print("UpdateElement error" + "5 \n")
		t.Fatal()
	}

	item7 := CreateDataItem("7", 1111)
	if err := queue.UpdateElement(item7); err == nil {
		fmt.Print("UpdateElement error \n")
		t.Fatal()
	}

}

