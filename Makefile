PREFIX=/usr/local

MAKEFLAGS += --no-print-directory

test:
	@echo Running tests
	@$(MAKE) -C t

install:
	@echo install age $(PREFIX)/bin

