package main

import "fmt"

/*
ORIGINATOR
 */
type documentOriginator struct {
	state string
}
func (e *documentOriginator) createMemento() *documentMemento {
	return &documentMemento{state: e.state}
}
func (e *documentOriginator) restoreMemento(m *documentMemento) {
	e.state = m.getSavedState()
}
func (e *documentOriginator) setState(state string) {
	e.state = state
}
func (e *documentOriginator) getState() string {
	return e.state
}
/*
MEMENTO = Backup
*/
type documentMemento struct {
	state string
}
func (m *documentMemento) getSavedState() string {
	return m.state
}
/*
CARETAKER = Backups LIST of a particular ORIGINATOR
*/
type documentsCaretaker struct {
	documentBackups []*documentMemento
}
func (c *documentsCaretaker) addMemento(m *documentMemento) {
	c.documentBackups = append(c.documentBackups, m)
}
func (c *documentsCaretaker) getMemento(index int) *documentMemento {
	return c.documentBackups[index]
}


/*
CLIENT
 */
func main() {
	document := &documentOriginator{}
	documentBackups := &documentsCaretaker{
		documentBackups: make([]*documentMemento, 0),
	}

	//Document (Originator) SAVING Backups (Mementos) to its Backup List (Caretaker)
	document.setState("A")
	fmt.Printf("Document-Originator Current State: %s\n", document.getState())
	documentBackups.addMemento(document.createMemento())

	document.setState("B")
	fmt.Printf("Document-Originator Current State: %s\n", document.getState())
	documentBackups.addMemento(document.createMemento())

	document.setState("C")
	fmt.Printf("Document-Originator Current State: %s\n", document.getState())
	documentBackups.addMemento(document.createMemento())


	//Document (Originator) RESTORING saved Backups (Mementos) from its Backup List (Caretaker)
	document.restoreMemento(documentBackups.getMemento(1))
	fmt.Printf("Restored to State: %s\n", document.getState())

	document.restoreMemento(documentBackups.getMemento(0))
	fmt.Printf("Restored to State: %s\n", document.getState())

}
