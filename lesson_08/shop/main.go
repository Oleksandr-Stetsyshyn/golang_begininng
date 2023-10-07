package shop

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Product struct {
	Name  string
	Price int
}

type Order struct {
	Name     string
	Item     Product
	Quantity int
}

func (o *Order) String() string {
	return fmt.Sprintf("%s ordered %d %s for a total of %.f грн\n", o.Name, o.Quantity, o.Item.Name, float64(o.Quantity)*float64(o.Item.Price))
}

func generateRequests(ctx context.Context, requests chan<- Order, names []string) {
	defer close(requests)

	products := []Product{
		{Name: "гумова курка", Price: 100},
		{Name: "дерев'яна корова", Price: 50},
		{Name: "зoлoтe яйце продуцент", Price: 70},
		{Name: "самокошарка з інтернет-підключенням", Price: 30},
		{Name: "металева качка-нагрівач", Price: 80},
		{Name: "робот-селянин", Price: 200},
		{Name: "автоматичний трюфельовикопувач", Price: 150},
		{Name: "смарт-віник", Price: 300},
		{Name: "короводоїльник", Price: 300},
		{Name: "бензинова коса 'Дружба 2'", Price: 300},
	}

	clients := names

	for {
		select {
		case <-ctx.Done():
			return
		default:
			product := products[rand.Intn(len(products))]
			name := clients[rand.Intn(len(clients))]
			quantity := rand.Intn(5) + 1
			requests <- Order{Name: name, Item: product, Quantity: quantity}
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		}
	}
}

func processRequests(ctx context.Context, requests <-chan Order) {

	for {
		select {
		case <-ctx.Done():
			return
		case order := <-requests:
			orderString := order.String()
			fmt.Print(orderString)
		}

	}
}

func Shop() {
	// go run main.go -names="Акош Августа Бодоґ Чаба Чолт Ежайяш Ференц Гуґо Жиґмонд Йоужі Бачі" -duration=10
	var namesStr string
	flag.StringVar(&namesStr, "names", "Акош Августа Бодоґ Чаба Чолт Ежайяш Ференц Гуґо Жиґмонд Йоужі Бачі", "space-separated list of names")

	var duration int
	flag.IntVar(&duration, "duration", 10, "duration of the simulation in seconds")
	flag.Parse()

	names := strings.Split(namesStr, " ")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	requests := make(chan Order)

	go generateRequests(ctx, requests, names)
	go processRequests(ctx, requests)

	time.Sleep(time.Duration(duration) * time.Second)

	
	fmt.Println("Working day is over, go home!")
}
