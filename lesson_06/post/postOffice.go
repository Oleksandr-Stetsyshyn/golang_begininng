package postOffice

import "fmt"

type Parcel interface {
	GetSender() string
	GetRecipient() string
	GetType() string
	Send()
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
func (b *Box) Send() {
	fmt.Println("Box was sent by postal car", b.Recipient)
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

type Bomb struct {
	Sender    string
	Recipient string
}
func (b *Bomb) GetSender() string {
	return b.Sender
}
func (b *Bomb) GetRecipient() string {
	return b.Recipient
}
func (b *Bomb) GetType() string {
	return "Bomb"
}
func (b *Bomb) Send() {
	fmt.Println("Bomb was sent by rocket to ", b.Recipient)
}



func (e *Envelope) Send() {
	fmt.Println("Envelope was sent by postal pigeon to", e.Recipient)
}

type SortingOffice struct{}

func (so *SortingOffice) Sort(p Parcel) {
	p.Send()
}

func Office() {
	so := SortingOffice{}

	box := Box{
		Sender:    "Василь Петрович",
		Recipient: "Jane Doe",
	}

	envelope := Envelope{
		Sender:    "John Doe",
		Recipient: "Василь Петрович",
	}

	bomb := Bomb{
		Sender:    "Йоужі",
		Recipient: "Путін",
	}

	so.Sort(&box)
	so.Sort(&envelope)
	so.Sort(&bomb)
}
