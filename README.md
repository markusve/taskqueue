# taskqueue


## build:

go build cmd/main.go

## copilot prompts used to create:

- Create in golang a program that implements task queue. It has API endpoint called "scheduleTask" that takes in two parameters: a context object and triggering time for the task. It runs scheduling loop which triggers task at given time. It should only print each context object at the given time.
- Make it an independent module usable from other modules, without http api.
- Add DeleteTask function

## final copilot result:

- taskqueue.go
- main.go

