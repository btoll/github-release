CC      	= go
PROGRAM		= github-release
prefix		= /usr

.PHONY: build clean distclean install package uninstall

build: $(PROGRAM)

$(PROGRAM):
	$(CC) build

clean:
	rm -f $(PROGRAM)

distclean: clean

# https://www.gnu.org/software/make/manual/html_node/DESTDIR.html
install:
	install -D -m 0755 $(PROGRAM) $(DESTDIR)$(prefix)/bin/$(PROGRAM)

package:
	git package-and-release --create --tag 1.0.0

uninstall:
	-rm -f $(DESTDIR)$(prefix)/bin/$(PROGRAM)

