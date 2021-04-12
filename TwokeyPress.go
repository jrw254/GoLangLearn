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

// First Key Function Below	
	
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
    i.inputType = 1 //INPUTKEYBOARD     
    i.ki.wVk = 0x12 // virtual key code for Alt 
    lop, _, err := sendInputProc.Call(         
        uintptr(1),
        uintptr(unsafe.Pointer(&i)),
        uintptr(unsafe.Sizeof(i)),     
        )     
	//log.Printf("KeyBoard: %v MSG: %v", lop, err) 
		log.Printf("%v %v", lop, err)
    } 

// Second Key Function Below	
    
func keyPress1() {     
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
    i.inputType = 1 //INPUTKEYBOARD     
    i.ki.wVk = 0x87 // virtual key code for F24 ; 0x41; 0x90(Insert)
    lop, _, err := sendInputProc.Call(         
        uintptr(1),
        uintptr(unsafe.Pointer(&i)),
        uintptr(unsafe.Sizeof(i)),     
        )     
	//log.Printf("KeyBoard: %v MSG: %v", lop, err) 
		log.Printf("%v %v", lop, err)
    } 

// Main Script Below

func main() {
		allDone := make(chan bool)  // Ends the script
		tick := time.NewTicker(10 * time.Second) // Time the keyPress function are pressed
		go func() {
		for {
		select {
		case <- allDone:
		return
		case <- tick.C:
			keyPress()
			keyPress1()
			}
		}
	}()
		time.Sleep(30 * time.Second)  // Length of time till it stops
			allDone <- true
		fmt.Println("Done for the day!!!!!!!!!")
	}
