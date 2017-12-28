package main

type Kotak struct {
	Panjang int
	Lebar   int
}

func (p *Kotak) Luas() int {
	return p.Panjang * p.Lebar
}

func (p *Kotak) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (p *Kotak) IsPersegi() bool {
	if p.Panjang == p.Lebar {
		return true
	}
	return false
}

func NewKotak(p, l int) *Kotak {
	return &Kotak{p, l}
}

// func main() {
// 	p := NewKotak(4, 8)
// 	fmt.Printf("Panjang= %d, Lebar= %d\n", p.Panjang, p.Lebar)
// 	fmt.Printf("Luas Persegi = %d\n", p.Luas())
// 	fmt.Printf("Keliling = %d\n", p.Keliling())
// 	fmt.Printf("Apakah persegi ? %v\n", p.IsPersegi())
// }
