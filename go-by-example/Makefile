SHELL := /bin/bash
# 你可以定义一个变量来传递提交信息
COMMIT_MESSAGE ?= "feature:Update by sth"

.PHONY: ci


ci:
	@echo "Pulling latest changes..."
	@git pull
	@echo "Adding changes..."
	@git add .
	@echo "Committing changes..."
	@git commit -m '$(COMMIT_MESSAGE)'
	@echo "Pushing changes to origin..."
	@git push

