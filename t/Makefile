T = $(wildcard t[0-9][0-9][0-9][0-9]-*.sh)
MAKEFLAGS += --no-print-directory

work_dir = ./tmp
export work_dir

JUNK = $(work_dir)

all: init $(T)
	@$(MAKE) clean

init:
	@mkdir -p $(work_dir)

clean:
	@$(RM) -r $(JUNK)

$(T):
	@./$@

.PHONY: all $(T)
