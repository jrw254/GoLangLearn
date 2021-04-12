package main 
import 
(   
	  "fmt"
    "log"    
    "syscall"     
    "time"     
    "unsafe" 
) 

var ( 
    user32 = syscall.NewLazyDLL("user32.dll")     
    sendInputProc = user32.NewProc("SendInput") ) 
    
func keyPress() {     
    type keyboardInput struct {         
        wVk uint16         
        wScan uint16         
        dwFlags uint32         
        time uint32         
        dwExtraInfo uint64     
    }     

    type input struct {         
        inputType uint32         
        ki keyboardInput         
        padding uint64     
    }     
    
    var i input     
    i.inputType = 1 // INPUTKEYBOARD     
    i.ki.wVk = 0x87 // virtual key code for F24
    lop, _, err := sendInputProc.Call(         
        uintptr(1),
        uintptr(unsafe.Pointer(&i)),
        uintptr(unsafe.Sizeof(i)),     
        )
		log.Printf("%v %v", lop, err)
    } 

    func counter() {
	for i := 0; i >= 0;  {
	fmt.Println("Rotations:", i)
	time.Sleep(time.Second * 464) // Time set to repeat the function
	i++
	}
     }
	
    func main() {
		allDone := make(chan bool)  // Ends the timer
		tick := time.NewTicker(462 * time.Second) // --7.7 minutes or 462 Second, each loop length
		tn := time.Now() // Starts off with current time
			fmt.Printf("%s", tn)
			fmt.Println("\n.....Initializing.....")
	    go counter() 
			go func() {
				for {
					select {
					case <-allDone:
							return
					case  <-tick.C: 
						fmt.Println("\n In Session: Your Timer:")
						keyPress()
						fmt.Println("")
					case <-time.After(1 * time.Minute): // Running Log Here. Change Time to whatever
						log.Printf("Running...")
					}
				}	
			}()
			time.Sleep(4 * time.Hour)  // Length of time till it stops/ends
			allDone <- true
			fmt.Println("Done for the day!!!!!!!!!")
    }     
