package console


type MessageAddedEvent struct {
    Message ConsoleMessage `json:"message"`// Console message that has been added.
}