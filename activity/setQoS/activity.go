package setQoS


import (
    "github.com/TIBCOSoftware/flogo-lib/core/activity"
    "github.com/TIBCOSoftware/flogo-lib/logger"
    //"fmt"
    "os/exec"
	"os"
)
// Constants
const (
	ivScript = "script"
    ivDevice = "device"
    ivSpeed = "speed"
	ovResult = "result"
)

// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-setQoS")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
    metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
    return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
    return a.metadata
}

// THIS HAS CHANGED
// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {
    // Get the activity data from the context
    script := context.GetInput(ivScript).(string)
    device := context.GetInput(ivDevice).(string)
    speed := context.GetInput(ivSpeed).(string)

    script = "/Users/iain/setQoS.sh"
    log.Infof("The Flogo run script input: [%s],[%s],[%s]", script, device, speed)
	var cmdOut []byte

    // Check if the file exists
	_, err = os.Stat(script)

	if err != nil {
		// If the file doesn't exist return error
			context.SetOutput("result", err.Error())
            log.Infof("Error from setQoS activity: File [%s] does not exist", script)
			return true, err
	}


    //if cmdOut, err = exec.Command(cmdName, os.Args[1:]...).Output(); err != nil { 
    if cmdOut, err = exec.Command(script, device, speed).Output(); err != nil {       
		//fmt.Fprintln(os.Stderr, "There was an error runScript activity ", err)
        log.Infof("Error running Flogo setQoS activity: [%s]", err)
        context.SetOutput(ovResult, err.Error()) 
        return true, err
	} 
	rslt := string(cmdOut)
    // Set the result as part of the context
    context.SetOutput(ovResult, rslt)
    



    // Signal to the Flogo engine that the activity is completed
    return true, nil
}


