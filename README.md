# console-application

An application merges csv file into json file using Golang

### Requirement
```bash
go version go1.16.6 darwin/amd64
```

### Install
```bash
git clone https://github.com/jokernguyen921731/console-application.git
cd console-application
make install
```

### Test
Run all test case in project
```bash
make test
```

### Build
The result of building creates at example-build folder
```bash
make run
```

## Adapter
Command line
```bash
Usage: ./email-application COMMAND [Params]
Example: ./email-application help
Commands:
  start     [port]
  send      [input_json] [input_csv] [output_folder] [error_csv]
  version
```