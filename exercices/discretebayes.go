package main

import (
	"fmt"
	"github.com/quentingls/go-filter"
	"math/rand"
	"time"
)

type Train struct {
	pos            int
	trackLen       int
	kernel         []float64 // uncertainty of movement / model
	sensorAccuracy float64
}

func NewTrain(trackLen int, kernel []float64, sensorAccuracy float64) *Train {
	return &Train{
		trackLen:       trackLen,
		pos:            0,
		kernel:         kernel,
		sensorAccuracy: sensorAccuracy,
	}
}

func (t *Train) sensePositionAndGetLikelihood() (int, []float64) {
	position := t.sense()
	likelihood := make([]float64, t.trackLen)
	factor := t.sensorAccuracy / (1 - t.sensorAccuracy)
	for i := 0; i < t.trackLen; i++ {
		likelihood[i] = 1
		if position == i {
			likelihood[i] *= factor
		}
	}
	return position, likelihood
}

func (t *Train) move(distance int) {
	r := rand.New(rand.NewSource(99))
	r.Seed(time.Now().UnixNano())
	number := r.Float64()
	var offset int
	acc := 0.0
	for i, p := range t.kernel {
		acc += p
		if acc > number {
			offset = i - ((len(t.kernel) - 1) / 2)
			break
		}
	}
	t.pos = (t.pos + distance + offset) % t.trackLen
}

func (t *Train) sense() int {
	offset := 0
	r := rand.New(rand.NewSource(99))
	r.Seed(time.Now().UnixNano())
	number := r.Float64()
	if number > t.sensorAccuracy {
		if r.Float64() > 0.5 {
			offset++
		} else {
			offset--
		}
	}
	return t.pos + offset
}

func main() {
	kernel := []float64{0.1, 0.9, 0.1}
	petitTrain := NewTrain(10, kernel, 0.95)
	fmt.Println(petitTrain.sense())

	prior := make([]float64, 10)
	for i, _ := range prior {
		prior[i] = 1.0 / float64(len(prior))
	}

	for i := 0; i < 100; i++ {
		petitTrain.move(1)
		_, likelihood := petitTrain.sensePositionAndGetLikelihood()
		posterior := gofilter.Update(likelihood, prior)
		fmt.Println(posterior)
		prior = gofilter.Predict(posterior, kernel, 1)
	}
	fmt.Println(petitTrain.pos)
}
