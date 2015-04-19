# Makefile for
all: usage

PROGRAM=hub-search
EDITOR=vim

edit e:
	$(EDITOR) $(PROGRAM).go

readme md:
	$(EDITOR) README.md

build b:
	go build $(PROGRAM).go
	@ls -alF --color=auto

run r:
	./$(PROGRAM) --lang=go webrtc server --list=1 -d

test t:
	./$(PROGRAM) tugboat 
	./$(PROGRAM) tugboat --lang=go
	./$(PROGRAM) tugboat --lang=go --sort=stars
	./$(PROGRAM) tugboat --lang=go --sort=stars --order=asc --text --score=10
	./$(PROGRAM) tugboat --lang=go --down --list=1
	./$(PROGRAM) --lang=go webrtc server --down

rebuild:
	rm -f ./$(PROGRAM)
	go build $(PROGRAM).go
	@ls -alF --color=auto

install i:
	go install

clean:
	rm -f ./$(PROGRAM)

start:
	ssh -T git@github.com

git g:
	make clean
	git status
	git init
	git add README.md Makefile $(PROGRAM).go
	git commit -m "git test and update Makefile"
	git push -u https://sikang99@github.com/sikang99/$(PROGRAM) master
	git log --oneline -5

git2:
	git init
	git add README.md Makefile $(PROGRAM).go
	git commit -m "git test and update Makefile"
	git remote add origin https://sikang99@github.com/sikang99/$(PROGRAM)
	git push -u origin master


make m:
	$(EDITOR) Makefile

usage:
	@echo ""
	@echo "Makefile for '$(PROGRAM)', by Stoney Kang, 2015/04/12"
	@echo ""
	@echo "usage: make [edit|readme|build|run|test|rebuild|clean|git]"
	@echo "	edit    : edit source"
	@echo "	build   : compile source"
	@echo "	run     : execute $(PROGRAM)"
	@echo "	test    : test $(PROGRAM) options"
	@echo "	install : install $(PROGRAM) to $(GOPATH)/bin"
	@echo "	git     : upload $(PROGRAM) to github.com"
	@echo ""
