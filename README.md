# Plugins

This project exposes the plugins implementation. 
Plugins should be exposing a Validator so that, at the time of provisioning
a workflow, the API role can validate the provided configuration; and a Factory 
and a Feature so that, at the time of 
provisioning a workflow, the Engine role can provision a Feature and execute 
it in the context of workflow execution.


## Build

Plugins can be build using the following target in the Makefile. 
Notice that projects importing plugins, AKA api and engine, will need to set 
the file location for proper loading, otherwise it will fail. 

```$ make build```

## Notes

A few findings while working with plugins:

* Projects using plugins and plugins implementing some interfaces must be 
  build using the exact same lib version, otherwise it will return an error 
  indicating so. 
* We used -trimpath file to build the plugins and plugin users, engine and 
  api, had to be launched with the same flag. According to the documentation: 
    remove all file system paths from the resulting executable.
    Instead of absolute file system paths, the recorded file names
  will begin either a module path@version (when using modules),
  or a plain import path (when using the standard library, or GOPATH).
* A preliminary conclusion indicates that in our ultimate use case, where 
  users will be implementing their on features, users should be providing 
  the source code so that it can be built using the proper versions of 
  language and libraries. 



