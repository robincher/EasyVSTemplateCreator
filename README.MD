## Easy Visual Studio Code Template Creator

My little VS GoLang/Rust project generator. Got tired of manually cloning and renaming files.
this program pulls the VSCode template and removes the git repo settings and then renames it.

It can be done with a bash script, but I'll like a chance to practice Go

### Start

#### Compile and Build the source code 

* (Option 1) Building in Visual Studio Code

   * Cmd + Shift + B (macOS)
   * Ctr + Shift + B (Windows)

* (Option 2) Command Promt / Bash

```
go build -o ./bin/output.exe ./src
```

#### Running the Visual Studio code Template Creator

1. Locate the executable, by default it should be in the root ./bin folder

2. Run the following command in Bash or cmd prompt

```
    ./output.exe <<projectName>> <<projecType>>
```

* **projectName** : The project name of the template. The executable will pull the template and renamed it to the projectName you specify. 

* **projectTyupe** : Project type, currently supported values are '**golang**', '**rust**', '**ts**' and '**js**'

