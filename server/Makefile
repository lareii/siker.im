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

clean:
	sudo docker compose down

fclean: clean
	rm -rf $(CERTS_DIR) $(LOG_FILE)

re: fclean run

.PHONY: run clean fclean re
