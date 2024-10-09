include .env
export $(shell sed 's/=.*//' .env)

DB_URL = database/logs.db


MIGRATE_PATH := database/migrations
MIGRATE_CMD := goose -dir=$(MIGRATE_PATH) sqlite3 $(DB_URL)

.PHONY: mig-up mig-down mig-force mig-create

mig-up:
	$(MIGRATE_CMD) up

mig-down:
	$(MIGRATE_CMD) down

mig-status:
	$(MIGRATE_CMD) status

mig-create:
	goose -dir=database create logs sql