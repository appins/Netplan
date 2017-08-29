all: remove_CN
	mkdir -p entries
	go run *.go

log: remove_CN
	mkdir -p entries
	$(eval TIME := $(shell date +"%a%b%d%y%T"))
	mkdir -p logs
	touch logs/$(TIME).log
	touch logs/latest.log
	ln -f logs/$(TIME).log logs/latest.log
	go run *.go 2> logs/$(TIME).log&
	tail -f logs/latest.log

killit:
	touch CLOSE_NETPLAN
	sh safekill.sh
	rm -f CLOSE_NETPLAN
	rm -f KILL_NETPLAN_NOW

remove_CN:
	rm -f CLOSE_NETPLAN
	rm -f KILL_NETPLAN_NOW

reset: remove_CN
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

convert-old:
	sh ./converttojson.sh

check:
	sh ./checkversion.sh

updatepush:
	zip -r public/netplan.zip .
	make
