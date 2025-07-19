LOG_FILE = app.log

run: $(LOG_FILE)
	sudo docker compose up -d

$(LOG_FILE):
	touch $(LOG_FILE)

stop:
	sudo docker compose down

clean: stop
	rm -rf $(LOG_FILE)

rebuild: stop
	sudo docker compose up --build -d

.PHONY: run stop clean rebuild
