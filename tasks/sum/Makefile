SRC = $(CURDIR)/source
BLD = $(CURDIR)/build
RES = $(CURDIR)/built

$(shell mkdir -p $(BLD))
$(shell mkdir -p $(RES))

TESTS = $(RES)/tests
CHECKER = $(RES)/checker

all: $(TESTS) $(CHECKER)
	rm -r $(BLD)

$(TESTS): $(BLD)/generator
	$(shell mkdir -p $(TESTS))
	cd $(TESTS); $(BLD)/generator

$(BLD)/generator: $(SRC)/generator.cpp
	g++ -o $@ $^

$(CHECKER): $(SRC)/checker.cpp
	g++ -o $@ $^
