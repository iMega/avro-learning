test: acceptance

acceptance:
	@touch $(CURDIR)/mysql.log
	@docker-compose up -d --scale acceptance=0
	@docker-compose up --abort-on-container-exit acceptance

clean:
	@-rm $(CURDIR)/mysql.log
	@TAG=$(TAG) docker-compose rm -sfv

.PHONY: acceptance
