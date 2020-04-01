package main

// this is the contract to implement
// each entity has different algorithm to implement this contract
type messagingHandler interface {
	send(msg string, to string)
}

/*****************************************/
// sms is an implementation of messagingHandler
type sms struct {}

func newSms() *sms{
	return &sms{}
}

func (s *sms) send(msg string, to string) {
	// buisness logic for messaging via sms
}

/********************************************/
// pushNotif is an implementation of messagingHandler
type pushNotif struct {}

func newPushNotif() *pushNotif{
	return &pushNotif{}
}

func (s *pushNotif) send(msg string, to string) {
	// buisness logic for messaging via push notification
}

/********************************************/

// createMessagingHandler returns an implementation of the contract for the given messaging type
func createMessagingHandler(mesType string) messagingHandler{
	switch mesType {
	case "sms": return newSms()
	case "push" : return newPushNotif()
	default:
		return nil
	}
}
