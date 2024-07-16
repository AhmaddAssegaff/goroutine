package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	NameProduct string
	Price       int
}

func calculateTotal(items []Item, quantityMap map[string]int, ch chan int) {
	localTotal := 0
	for _, item := range items {
		if quantity, exists := quantityMap[item.NameProduct]; exists {
			itemTotal := item.Price * quantity
			fmt.Printf("Product: %s, Quantity: %d, Price per unit: %d, Total: %d\n", item.NameProduct, quantity, item.Price, itemTotal)
			localTotal += itemTotal
		}
	}
	ch <- localTotal // Mengirim hasil total lokal melalui channel
}

func main() {
	items := []Item{
		{NameProduct: "hp", Price: 100},
		{NameProduct: "laptop", Price: 200},
		{NameProduct: "tablet", Price: 300},
		{NameProduct: "monitor", Price: 400},
		{NameProduct: "keyboard", Price: 500},
		{NameProduct: "mouse1", Price: 600},
		{NameProduct: "mouse2", Price: 600},
		{NameProduct: "mouse3", Price: 600},
		{NameProduct: "mouse4", Price: 600},
	}

	// Specify the quantity of each product to purchase
	quantityMap := map[string]int{
		"hp":       30320,
		"laptop":   10032,
		"mouse1":   2200,
		"mouse2":   2200,
		"mouse3":   2200,
		"mouse4":   2200,
		"tablet":   30230,
		"monitor":  10900,
		"keyboard": 2121,
	}

	ch := make(chan int, len(items)) // Channel dengan buffer untuk menerima hasil total

	// Memulai goroutine untuk setiap item
	var wg sync.WaitGroup
	for _, item := range items {
		wg.Add(1)
		go func(item Item, quantityMap map[string]int) {
			defer wg.Done()
			calculateTotal([]Item{item}, quantityMap, ch)
		}(item, quantityMap)
	}

	// Menunggu sampai semua goroutine selesai
	wg.Wait()

	// Menutup channel setelah semua goroutine selesai
	close(ch)

	start := time.Now()

	total := 0
	// Menerima hasil dari channel dan menghitung total akhir
	for localTotal := range ch {
		total += localTotal
	}

	elapsed := time.Since(start)
	fmt.Printf("Total Price: %d, Waktu eksekusi: %v\n", total, elapsed.Milliseconds())
}
