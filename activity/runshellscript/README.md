# Execute a shell script
This activity provides your Flogo application with the ability to execute shell or bash scripts.

# Installation
```
flogo install github.com/iainharfield/flogo/activity/runshellscript
```

# Schema
## Inputs and Outputs:
```
{
  "inputs":[
    {
      "name": "cmd",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```

### Inputs: Schema of cmd parameter:
The activity expects the name of the script to execute and the script parameters be put into JSON schema. The schema contains the fully qualifies path and script name to execute which is manfdatory and the parameters to be passed in.  The activity is limited to 20 parameters. 

JSON schma example, execute test01.sh in local directory and passes in 3 parameters.  
```
{
  "cmd":"./test01.sh",
  "params":"aaa bbb ccc"
}
```
Note that the params field is a single string of space separated values and not an array or params.
```
{"cmd":"./test01.sh","params":"aaa bbb ccc"}`
```
### Output from the Activity:
The output is the text string returned from the launched script. See very simple example below.

## Configuration Example
```
{
    "id": "runScript_1",
    "name": "Execute a shell script",
    "description": "Run a script",
    "activity": {
    "ref": "github.com/iainharfield/flogo/activity/runshellscript",
    "mappings": {
    "input": [
        {
            "type": "assign",
            "value": "$flow.command",
            "mapTo": "cmd"
        }
     ]
}
```
## Example bash script
```
#!/bin/bash
printf "test01 output: $1 $2 $3"
```

### Notes
This activity has only been tested on Linux/MacOS only.  No testing on Windows.

