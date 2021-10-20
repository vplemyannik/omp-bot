package card

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/payment"
)

type CardService interface {
	Describe(cardID uint64) (*payment.Card, error)
	List(cursor uint64, limit uint64) ([]payment.Card, error)
	Create(payment.Card) (uint64, error)
	Update(cardID uint64, card payment.Card) error
	Remove(cardID uint64) (bool, error)
}

type DummyCardService struct{}

var array = make([]payment.Card, 0)

func (d DummyCardService) Describe(cardID uint64) (*payment.Card, error) {
	if inBound(cardID) {
		return &array[cardID], nil
	}
	return nil, fmt.Errorf("There is no such card wiht id %v", cardID)
}

func (d DummyCardService) List(cursor uint64, limit uint64) ([]payment.Card, error) {
	if inBound(cursor) {
		return array[cursor:min(cursor+limit, uint64(len(array)))], nil
	}
	return nil, fmt.Errorf("Cursor is out of range")
}

func (d DummyCardService) Create(card payment.Card) (uint64, error) {
	array = append(array, card)
	return uint64(len(array) - 1), nil
}

func (d DummyCardService) Update(cardID uint64, card payment.Card) error {
	if inBound(cardID) {
		array[cardID] = card
		return nil
	}
	return fmt.Errorf("There is no such card with id %v", cardID)
}

func (d DummyCardService) Remove(cardID uint64) (bool, error) {
	if inBound(cardID) {
		removeIndex(array, cardID)
		return true, nil
	}
	return false, fmt.Errorf("There is no such card with id %v", cardID)
}

func NewDummyCardService() *DummyCardService {
	return &DummyCardService{}
}

func inBound(cardID uint64) bool {
	if cardID >= 0 && cardID < uint64(len(array)) {
		return true
	}
	return false
}

func removeIndex(arr []payment.Card, index uint64) []payment.Card {
	return append(arr[:index], arr[index+1:]...)
}

func min(first uint64, second uint64) uint64 {
	if first < second {
		return first
	}

	return second
}
