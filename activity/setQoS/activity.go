package setQoS


import (
    "github.com/TIBCOSoftware/flogo-lib/core/activity"
    "github.com/TIBCOSoftware/flogo-lib/logger"
    //"fmt"
    "os/exec"
	"os"
    "encoding/json"
)
// Constants
const (
	script = "script"
    device = "device"
    speed = "speed"
	result = "result"
)

type params struct {
  name string
  value string
}

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
    //ivMsg := `{"pigeon":"xxx","eagle":"yyy","animals":"zzz"}`

    ivMsg := context.GetInput(script).(string)
    log.Infof("ivMsg = %s", ivMsg)
    var params map[string]interface{}
    json.Unmarshal([]byte(ivMsg), &params)
    
    for key, value := range params {
        log.Infof("PARAMS:[%s],[%s]",key,value.(string))
    }
    fname := params["script"].(string)
    log.Infof("%s",fname)

    //log.Infof("PARAMS:[%s],[%s],[%s]", params.script, params.device, params.speed)

    // Get the activity data from the context
    ivScript := context.GetInput(script).(string)
    ivDevice := context.GetInput(device).(string)
    ivSpeed := context.GetInput(speed).(string)

    //ivScript = "/Users/iain/setQoS.sh"
    log.Infof("The Flogo run script input: [%s],[%s],[%s]", ivScript, ivDevice, ivSpeed)
	var cmdOut []byte

    // Check if the file exists
	_, err = os.Stat(fname)

	if err != nil {
		// If the file doesn't exist return error
			context.SetOutput("result", err.Error())
            log.Infof("Error from setQoS activity: File [%s] does not exist", ivScript)
			return true, err
	}


    //if cmdOut, err = exec.Command(cmdName, os.Args[1:]...).Output(); err != nil { 
    if cmdOut, err = exec.Command(fname, ivDevice, ivSpeed).Output(); err != nil {       
		//fmt.Fprintln(os.Stderr, "There was an error runScript activity ", err)
        log.Infof("Error running Flogo setQoS activity: [%s]", err)
        context.SetOutput(result, err.Error()) 
        return true, err
	} 
	rslt := string(cmdOut)
    // Set the result as part of the context
    context.SetOutput(result, rslt)
    



    // Signal to the Flogo engine that the activity is completed
    return true, nil
}


