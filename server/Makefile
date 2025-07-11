CERTS_DIR = .certs/
LOG_FILE =  app.log

run: $(CERTS_DIR)
	touch $(LOG_FILE)
	sudo docker compose up -d

stop:
	sudo docker compose down

$(CERTS_DIR):
	./scripts/certs.sh

clean:
	sudo docker compose down

fclean: clean
	rm -rf $(CERTS_DIR) $(LOG_FILE)

re: fclean run

.PHONY: run clean fclean re