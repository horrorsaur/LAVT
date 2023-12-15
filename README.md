# LAVT


- https://pkg.go.dev/runtime/pprof

### LSP settings

i set some local lsp settings so that my go env is set up correctly when swapping between platforms
```json
{
  "go.toolsEnvVars": { "GOOS": "windows" },
  "gopls": {
    "build.env": {
      "GOOS": "windows"
    }
  }
}
```


### DISCLAIMER 
LAVT isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.
