package main

import (
	"errors"
	"fmt"
)

type account struct {
	id string
}

func newAccount(id string) *account {
	return &account{id: id}
}

func (a *account) checkAccount(accountID string) bool {
	fmt.Println("check accountID")
	return true
}

/*******************************************************/

type wallet struct {
	total int
}

func (w *wallet) increment(amount int) {
	fmt.Println("increment wallet total")
	w.total += amount
}

func (w *wallet) decrement(amount int) {
	fmt.Println("decrement wallet total")
	w.total -= amount
}

/******************************************************/

type security struct {
	code int
}

func newSecurityCode(code int) *security {
	return &security{code: code}
}

func (s *security) checkCode(code int) bool{
	fmt.Println("check security code")
	return true
}

/******************************************************/

type notification struct {
}

func (n *notification) sendNotif(deviceID string) {
	fmt.Println("sending notification")
}

func (n *notification) getDeviceID(accountID string) string{
	fmt.Println("getting device id")
	return "deviceID"
}

/******************************************************/


type Bank struct {
	account  *account
	wallet   *wallet
	security *security
	notif    *notification
}

func NewBankFacade(accountID string, code int) *Bank {
	return &Bank{
		account:  newAccount(accountID),
		wallet:   &wallet{},
		security: newSecurityCode(code),
		notif:    &notification{},
	}
}

/******************************************************/


func (b *Bank) addMoney(accountID string, code int, amount int) error {
  ok := b.account.checkAccount(accountID)
  if !ok {
  	return errors.New("wrong account id")
  }
  ok = b.security.checkCode(code)
  if !ok {
  	return errors.New("wrong code")
  }
  b.wallet.increment(amount)
  deviceID := b.notif.getDeviceID(accountID)
  b.notif.sendNotif(deviceID)
  return nil
}
