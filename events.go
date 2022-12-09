// AUTOGENERATED - DO NOT EDIT

package tdlib

type EventFunctions interface {
	GetRawUpdatesChannel(capacity int) chan UpdateMsg
	AddEventReceiver(msgInstance TdMessage, filterFunc EventFilterFunc, channelCapacity int) EventReceiver
}

// EventFilterFunc used to filter out unwanted messages in receiver channels
type EventFilterFunc func(msg *TdMessage) bool

// EventReceiver used to retreive updates from tdlib to user
type EventReceiver struct {
	Instance   TdMessage
	Chan       chan TdMessage
	FilterFunc EventFilterFunc
}

// GetRawUpdatesChannel creates a general channel that fetches every update comming from tdlib
func (client *ClientImpl) GetRawUpdatesChannel(capacity int) chan UpdateMsg {
	client.rawUpdates = make(chan UpdateMsg, capacity)
	return client.rawUpdates
}

// AddEventReceiver adds a new receiver to be subscribed in receiver channels
func (client *ClientImpl) AddEventReceiver(msgInstance TdMessage, filterFunc EventFilterFunc, channelCapacity int) EventReceiver {
	receiver := EventReceiver{
		Instance:   msgInstance,
		Chan:       make(chan TdMessage, channelCapacity),
		FilterFunc: filterFunc,
	}

	client.receiverLock.Lock()
	defer client.receiverLock.Unlock()
	client.receivers = append(client.receivers, receiver)

	return receiver
}
