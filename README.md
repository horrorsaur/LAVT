# LAVT

(WIP - but not actually in progress atm)

this application is built with [Wails](https://wails.io/docs/howdoesitwork), which is essentially a Go app with a webkit frontend (like electron)

### Background

I often find myself taking a look at some peoples ranks and it gets annoying having to constantly alt tab out of the game. 
ive tried the tracker stuff out there, but ads yadayada. this gave me a chance to pick up some knowledge about the Riot LCU API & test out some
cross-platform desktop application dev w/ golang (which is still relatively new to me)

### Getting started
WIP - still working on this thing

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
