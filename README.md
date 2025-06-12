# Task Tracker CLI

This is the realization of project from [roadmap.sh](https://roadmap.sh)

Details about the project's specification you can find here - [task-tracker](https://roadmap.sh/projects/task-tracker)

To build program you have to [download GO](https://go.dev/dl/). And run next command in terminal after cloning the repo (in the project's directory) for:

- Win: ```go build -o task-cli.exe .\cmd\task-cli\task.cli.go```
- Linux\MacOS: ```go build -o task-cli /cmd/task-cli/task-cli.go```

Note: you can also use ```run``` instead of ```build```

The actual version of GO required to run and build app you can see in ```go.mod``` file

After build a program you can use it just like in [examples described in spec](https://roadmap.sh/projects/task-tracker#example). Also you can use next command to get more info about all allowed commands:
```task-cli help```

Before adding or changing a features run tests with following command:
```go test .\...```

The structure of project was implemented from open GitHub's repo [golang-standards/project-layout/](https://github.com/golang-standards/project-layout/)

The project was developed with educational purposes

