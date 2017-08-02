all:
	go run *.go

log:
	mkdir -p logs
	go run *.go > logs/$(shell date +"%a%b%d%y%T").log

reset:
	rm -rf logs
	sh ./deleteentries.sh
