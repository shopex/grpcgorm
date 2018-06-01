OBJ := grpcgorm
NOW := $(shell date "+%Y/%m/%d/%H:%m:%S")

all: ${OBJ}

${OBJ}: *.go */*.go */*/*.go
	go build -ldflags "-X main.BuildTime=${NOW} -s" .

doc:
	swag init

start: all
	./${OBJ} server start

teststart: all
	mv ${OBJ} ${OBJ}t
	PORT=8082 ./${OBJ}t server start

clean:
	-rm ${OBJ} 

test:
	(cd testcase ; go test)
