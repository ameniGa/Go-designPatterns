package main

import "fmt"

// window represents the receiver of command
type window struct {
}

func (w *window) open() {
	fmt.Println("window is opened")
}

func (w *window) close() {
	fmt.Println("window is closed")
}

/*****************************************************************/

// command is the contract to implement by any type of command
type command interface {
	execute()
	unexecute()
}
/*****************************************************************/

// closeCmd is a concrete command
// implements the command interface
type closeCmd struct {
	receiver *window
}

func (c *closeCmd) execute() {
	c.receiver.close()
}

func (c *closeCmd) unexecute() {
	c.receiver.open()
}

/*****************************************************************/

// openCmd is a concrete command
// implements the command interface
type openCmd struct {
	receiver window
}

func (o *openCmd) execute() {
	o.receiver.open()
}

func (o *openCmd) unexecute() {
	o.receiver.close()
}

/*****************************************************************/

// remoteControl represents the invoker
type remoteControl struct {
	openCmd command
	closeCmd command
}

func (u remoteControl) openWindow() {
	u.openCmd.execute()
}

func (u remoteControl) closeWindow() {
	u.closeCmd.execute()
}

func NewRemoteControl(openCmd, closeCmd command) remoteControl {
	return remoteControl{
		openCmd:  openCmd,
		closeCmd: closeCmd,
	}
}

/*****************************************************************/

func main() {
	remoteControl := NewRemoteControl(&openCmd{}, &closeCmd{})
	remoteControl.openWindow()
	remoteControl.closeWindow()
}
