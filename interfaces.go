package main

import "fmt"

// Dikdörtgen modeli
type Dikdörtgen struct {
	En  float64
	Boy float64
}

// Kare modeli
type Kare struct {
	Kenar float64
}

// Alan bir dikdörtgenin alanını hesaplayan bir fonksiyondur.
func (di Dikdörtgen) Alan() float64 {
	return di.En * di.Boy
}

// Çevre bir dikdörtgenin çevresini hesaplayan bir fonksiyondur.
func (di Dikdörtgen) Çevre() float64 {
	return (2 * di.En) + (2 * di.Boy)
}

// Alan bir karenin alanını hesaplayan bir fonksiyondur.
func (k Kare) Alan() float64 {
	return k.Kenar * k.Kenar
}

// Çevre bir karenin çevresini hesaplayan bir fonksiyondur.
func (k Kare) Çevre() float64 {
	return k.Kenar * 4
}

// Geometri Alan ve Çevre fonksiyonlarını barındıran bir interface'dir.
type Geometri interface {
	Alan() float64
	Çevre() float64
}

func hesapla(g Geometri) {
	fmt.Println(g)
	fmt.Println(g.Alan())
	fmt.Println(g.Çevre())
}
