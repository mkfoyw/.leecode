msg := $(shell date)

up:
	git add *
	git commit -m '$(msg)''
	git push origin 