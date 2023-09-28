package postOffice

import "fmt"

// Створити інтерфейс «Посилка» й реалізувати його для двох класів — «Коробка» і «Конверт».
// Для кожної поштової посилки необхідно зберігати адресу отримувача й відправника.
// Додати сортувальний відділ, який залежно від типу посилки відправляє її тим або іншим шляхом.

type Parcel interface {
	GetSender() string
	GetRecipient() string
	GetType() string
}

type Box struct {
	Sender    string
	Recipient string
}

func (b *Box) GetSender() string {
	return b.Sender
}
func (b *Box) GetRecipient() string {
	return b.Recipient
}

func (b *Box) GetType() string {
	return "Box"
}

type Envelope struct {
	Sender    string
	Recipient string
}

func (e *Envelope) GetSender() string {
	return e.Sender
}
func (e *Envelope) GetRecipient() string {
	return e.Recipient
}
func (e *Envelope) GetType() string {
	return "Envelope"
}

type SortingOffice struct{}

func (so *SortingOffice) Send(p Parcel) {
	switch p.GetType() {
	case "Box":
		fmt.Printf("Sending box from %s to %s via courier\n", p.GetSender(), p.GetRecipient())
	case "Envelope":
		fmt.Printf("Sending envelope from %s to %s via post\n", p.GetSender(), p.GetRecipient())
	}
}

func Office() {
	so := SortingOffice{}

	box := Box{
		Sender:    "John Doe",
		Recipient: "Jane Doe",
	}

	envelope := Envelope{
		Sender:    "John Doe",
		Recipient: "Jane Doe",
	}

	so.Send(&box)
	so.Send(&envelope)
}
