# Execute a shell script
This activity provides your Flogo application with the ability to execute shell or bash scripts.

# Installation
flogo install github.com/iainharfield/flogo/activity/runshellscript



# Schema
## Inputs and Outputs:
}
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



## Configuration Example
{
            "id": "setQoS_3",
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

