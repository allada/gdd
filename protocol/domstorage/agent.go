package domstorage


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type DOMStorageAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        Clear []chan<- ClearCommand
        GetDOMStorageItems []chan<- GetDOMStorageItemsCommand
        SetDOMStorageItem []chan<- SetDOMStorageItemCommand
        RemoveDOMStorageItem []chan<- RemoveDOMStorageItemCommand
    }
}
func NewAgent(conn *shared.Connection) *DOMStorageAgent {
    agent := &DOMStorageAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DOMStorageAgent) Name() string {
    return "DOMStorage"
}

func (agent *DOMStorageAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "enable":
            var out EnableCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Enable {
                c <-out
            }
        case "disable":
            var out DisableCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Disable {
                c <-out
            }
        case "clear":
            var out ClearCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Clear {
                c <-out
            }
        case "getDOMStorageItems":
            var out GetDOMStorageItemsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetDOMStorageItems {
                c <-out
            }
        case "setDOMStorageItem":
            var out SetDOMStorageItemCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDOMStorageItem {
                c <-out
            }
        case "removeDOMStorageItem":
            var out RemoveDOMStorageItemCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveDOMStorageItem {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *DOMStorageAgent) FireDomStorageItemsCleared(event DomStorageItemsClearedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemsCleared",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemsClearedOnTarget(targetId string, event DomStorageItemsClearedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemsCleared",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemRemoved(event DomStorageItemRemovedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemRemoved",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemRemovedOnTarget(targetId string, event DomStorageItemRemovedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemRemoved",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemAdded(event DomStorageItemAddedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemAdded",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemAddedOnTarget(targetId string, event DomStorageItemAddedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemAdded",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemUpdated(event DomStorageItemUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemUpdated",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemUpdatedOnTarget(targetId string, event DomStorageItemUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemUpdated",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *DOMStorageAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *DOMStorageAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *DOMStorageAgent) ClearNotify() <-chan ClearCommand {
    outChan := make(chan ClearCommand)
    agent.commandChans.Clear = append(agent.commandChans.Clear, outChan)
    return outChan
}
func (agent *DOMStorageAgent) GetDOMStorageItemsNotify() <-chan GetDOMStorageItemsCommand {
    outChan := make(chan GetDOMStorageItemsCommand)
    agent.commandChans.GetDOMStorageItems = append(agent.commandChans.GetDOMStorageItems, outChan)
    return outChan
}
func (agent *DOMStorageAgent) SetDOMStorageItemNotify() <-chan SetDOMStorageItemCommand {
    outChan := make(chan SetDOMStorageItemCommand)
    agent.commandChans.SetDOMStorageItem = append(agent.commandChans.SetDOMStorageItem, outChan)
    return outChan
}
func (agent *DOMStorageAgent) RemoveDOMStorageItemNotify() <-chan RemoveDOMStorageItemCommand {
    outChan := make(chan RemoveDOMStorageItemCommand)
    agent.commandChans.RemoveDOMStorageItem = append(agent.commandChans.RemoveDOMStorageItem, outChan)
    return outChan
}
func init() {

}
