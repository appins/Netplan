all:
	go run *.go

log:
	$(eval TIME := $(shell date +"%a%b%d%y%T"))
	mkdir -p logs
	touch logs/$(TIME).log
	touch logs/latest.log
	ln -f logs/$(TIME).log logs/latest.log
	go run *.go > logs/$(TIME).log&
	watch -n 10 -d cat logs/latest.log

killit:
	$(shell killall go)

reset:
	rm -rf logs
	sh ./deleteentries.sh

watch:
	mkdir -p logs
	touch logs/latest.log
	watch -n 1 -d cat logs/latest.log

entfix:
	chmod 777 -R entries
