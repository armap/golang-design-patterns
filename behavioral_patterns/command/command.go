package main

import "fmt"

/*
INVOKER BUTTON
 */
type invokerButton struct {
	command Command
}
func NewInvokerButton(command Command) *invokerButton {
	return &invokerButton{command}
}
func (ib *invokerButton) executeCommand() {
	ib.command.execute()
}
func (ib *invokerButton) undoCommand() {
	ib.command.undo()
}

/*
COMMAND INTERFACE
 */
type Command interface {
	execute()
	undo()
}
type command struct {
	receiver Receiver
}
/*
CONCRETE COMMAND SWITCH ON
 */
type commandSwitchOn struct {
	command
}
func NewCommandSwitchOn(receiver Receiver) *commandSwitchOn {
	return &commandSwitchOn{command{receiver}}
}
func (co *commandSwitchOn) execute() {
	co.receiver.switchOn()
}
func (co *commandSwitchOn) undo() {
	co.receiver.undoSwitchAction()
}
/*
CONCRETE COMMAND SWITCH OFF
*/
type commandSwitchOff struct {
	command
}
func NewCommandSwitchOff(receiver Receiver) *commandSwitchOff {
	return &commandSwitchOff{command{receiver}}
}
func (ct *commandSwitchOff) execute() {
	ct.receiver.switchOff()
}
func (ct *commandSwitchOff) undo() {
	ct.receiver.undoSwitchAction()
}

/*
RECEIVER INTERFACE
 */
type Receiver interface {
	switchOn()
	switchOff()
	undoSwitchAction()
}
/*
CONCRETE RECEIVER DEVICE
 */
type receiverDevice struct {
	name string
	isRunning bool
}
func NewReceiverDevice(name string) *receiverDevice {
	return &receiverDevice{
		name,
		false,
	}
}
func (rd *receiverDevice) switchOn() {
	rd.isRunning = true
	fmt.Printf("%s switched on\n", rd.name)
}
func (rd *receiverDevice) switchOff() {
	rd.isRunning = false
	fmt.Printf("%s switched off\n", rd.name)
}
func (rd *receiverDevice) undoSwitchAction() {
	rd.isRunning = !rd.isRunning
	var onOff string
	if rd.isRunning {onOff = "on"} else {onOff = "off"}
	fmt.Printf("%s is back %s\n", rd.name, onOff)
}
/*
CONCRETE RECEIVER APPLIANCE
*/
type receiverAppliance struct {
	name string
	powerConsumption int
	isRunning bool
}


/*
CLIENT TV & RADIO
 */
func main() {
	/*
	Set UI: Receivers, Commands & Invokers
	 */
	tv := NewReceiverDevice("tv")

	switchOnTv := NewCommandSwitchOn(tv)
	switchOffTv := NewCommandSwitchOff(tv)

	buttonSwitchOnTv := NewInvokerButton(switchOnTv)
	buttonSwitchOffTv := NewInvokerButton(switchOffTv)


	radio := NewReceiverDevice("radio")

	switchOnRadio := NewCommandSwitchOn(radio)
	switchOffRadio := NewCommandSwitchOff(radio)

	buttonSwitchOnRadio:= NewInvokerButton(switchOnRadio)
	buttonSwitchOffRadio := NewInvokerButton(switchOffRadio)


	/*
	User triggers command execution from invokers
	 */
	buttonSwitchOnTv.executeCommand()
	buttonSwitchOffTv.executeCommand()

	buttonSwitchOnRadio.executeCommand()
	buttonSwitchOffRadio.executeCommand()

	buttonSwitchOffRadio.undoCommand()
}