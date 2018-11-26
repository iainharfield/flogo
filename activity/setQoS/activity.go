package setQoS


import (
    "github.com/TIBCOSoftware/flogo-lib/core/activity"
    "github.com/TIBCOSoftware/flogo-lib/logger"
    "os/exec"
	"os"
    "encoding/json"
)
// Constants
const (
    command = "command"
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

    ivMsg := context.GetInput(command).(string)
    log.Infof("ivMsg = %s", ivMsg)
    var params map[string]interface{}
    json.Unmarshal([]byte(ivMsg), &params)
    
    // need to copy the slice to make it usable in exec.command later not exatly sure why.
    paramsArray  := make([]string, len(params))
    i := 0
    for key, value := range params {
        log.Infof("PARAMS:[%s],[%s]",key,value.(string))
        paramsArray[i] = value.(string)
        i++
    }
    
    // the first element is the command to execute including path
    cmd := params["script"].(string) // this should be the command or script to execute
    //log.Infof("%s",fname)

    // Check if the file exists
	_, err = os.Stat(cmd)
	if err != nil {
		// If the file doesn't exist return error
		context.SetOutput("result", err.Error())
        log.Infof("Error from runShell activity: File [%s] does not exist", cmd)
		return true, err
	}

	var cmdOut []byte
    //if cmdOut, err = exec.Command(cmdName, os.Args[1:]...).Output(); err != nil { 
    if cmdOut, err = exec.Command(cmd, paramsArray[1:]...).Output(); err != nil {       
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


