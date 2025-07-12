CERTS_DIR = .certs/
LOG_FILE =  app.log

run: $(CERTS_DIR) $(LOG_FILE)
	sudo docker compose up -d

$(CERTS_DIR):
	./scripts/certs.sh

$(LOG_FILE):
	touch $(LOG_FILE)

stop:
	sudo docker compose down

clean: stop
	rm -rf $(CERTS_DIR) $(LOG_FILE)

rebuild: stop
	sudo docker compose up --build -d

.PHONY: run stop clean rebuild
