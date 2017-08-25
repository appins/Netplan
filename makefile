all:
	mkdir -p entries
	go run *.go

log:
	mkdir -p entries
	$(eval TIME := $(shell date +"%a%b%d%y%T"))
	mkdir -p logs
	touch logs/$(TIME).log
	touch logs/latest.log
	ln -f logs/$(TIME).log logs/latest.log
	go run *.go 2> logs/$(TIME).log&
	tail -f logs/latest.log

killit:
	$(shell killall go)

reset:
	rm -rf logs
	sh ./deleteentries.sh

watch:
	mkdir -p logs
	touch logs/latest.log
	watch -n 5 -d cat logs/latest.log

entfix:
	chmod 777 -R entries

backup:
	mkdir -p entries
	zip -r entries.zip entries

postupdate:
	unzip entries.zip

convert-old:
	sh ./converttojson.sh

check:
	sh ./checkversion.sh

updatepush:
	zip -r public/netplan.zip .
	make
