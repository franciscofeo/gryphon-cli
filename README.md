<h1 align="center">Gryphon CLI :gear: </h1>

<img style="margin-top: 2em" align="right" src="./.docs/gryphon.png" height="269"> 

<br>

<p style="margin-top: 3em">
A Command Line Interface application to help improve development time by doing status
checking of the most used APIs in the world for web development and automating some 
stuff like opening essential software using just 1 command! </p>

<br>
<br>

## Current Modules
#### API Status Verification

#### Open Essentials Web Development Software

### Commands

- Help Command
```shellscript
    gryphon -h
```

This command will show a general information about the CLI and the main commands
or palettes to start using the application.

- Run all modules

```shellscript
    gryphon run -s DELAY_DURATION
```

`DELAY_DURATION` is a variable that represents the delay between the execution 
of the modules. The default value is 400 milliseconds and the maximum is 
600 milliseconds.

This command will execute all modules shown in "Current Modules" section.

#### API Verification Commands

- Verify All APIs

```shellscript
    gryphon api run
```

Check all APIs status in the official status page.

- Show Available APIs

```shellscript
    gryphon api -l
```

List all available APIs to check their status.

#### Open Software

- Open All Programs

```shellscript
    gryphon app run
```

Open all programs and URLs present in `urls.txt` file.

- List Available Software

```shellscript
    gryphon app -l
```

List all available programs to open automatically.

- Open a Specific Software

```shellscript
    gryphon app run -n SOFTWARE_NAME
```

`SOFTWARE_NAME` is a variable that represents the software name that will be
opened. The possible values can be found using the command to show the list of
available softwares in the topic above.