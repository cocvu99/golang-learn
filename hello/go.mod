module example.com/hello

go 1.17

// $ go mod edit -replace example.com/greetings=../'3 greetings'

replace example.com/greetings => "../3 greetings"

require example.com/greetings v1.1.0
