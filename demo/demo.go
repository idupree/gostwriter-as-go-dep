//    Copyright 2014, Raphael Estrada
//    Author email:  <galaktor@gmx.de>
//    Project home:  <https://github.com/idupree/gostwriter-as-go-dep>
//    Licensed under The GPL v3 License (see README and LICENSE files)
package main

// a simple demo tool showing how gostwriter can be used

import (
	"log"
	"time"

	"github.com/idupree/gostwriter-as-go-dep"
	"github.com/idupree/gostwriter-as-go-dep/key"
)

// uses the 't', 'e' and 's' keys to write 'test' to the
// console ten times. then it uses the 'ctrl' and 'c' keys
// to kill itself by emulating a 'CTRL+C' command
func main() {
	kb, err := gostwriter.New("foo")

	guard(err)

	t, err    := kb.Get(key.CODE_T);         guard(err);
	e, err    := kb.Get(key.CODE_E);         guard(err);
	s, err    := kb.Get(key.CODE_S);         guard(err);
	ret, err  := kb.Get(key.CODE_ENTER);     guard(err);

	ctrl, err := kb.Get(key.CODE_LEFTCTRL);  guard(err);
	c, err    := kb.Get(key.CODE_C);         guard(err);

	log.Println("this demo will type the word 'test' and a newline 10 times")
	log.Println("then it will terminate itself by pressing CTRL + C")

	cnt := 0
	for {
		<-time.After(time.Millisecond*100)
		push(t)
		<-time.After(time.Millisecond*100)
		push(e)
		<-time.After(time.Millisecond*100)
		push(s)
		<-time.After(time.Millisecond*100)
		push(t)
		<-time.After(time.Millisecond*500)
		push(ret)
		
		if cnt = cnt + 1; cnt == 10 {
			press(ctrl)
			press(c)
		}
	}

	kb.Destroy()
}

// presses and subsequently releases a key
func push(k *gostwriter.K) {
	err := k.Push(); guard(err);
}

// presses a key, if not already pressed. does not release
func press(k *gostwriter.K) {
	err := k.Press(); guard(err);
}

// releases a key, if not aready released.
func release(k *gostwriter.K) {
	err := k.Release(); guard(err);
}

// contains boilerplate error check. if error is present,
// prints it then exits the app
func guard(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
