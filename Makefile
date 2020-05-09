include .env

init:
	go get github.com/fatih/color
	go build src/train/train.go
	go build src/predict/predict.go

clean: 
	rm $(TRAINNAME)
	rm $(RESNAME)

all: init