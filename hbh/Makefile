BUILD_DIR = _build_

all: _init_ timed1 timed2
	@rm -f $(BUILD_DIR)/_go_.6

# --- [ cmd ] ------------------------------------------------------------------
timed1: hbh parse str
	6g -I $(BUILD_DIR)/pkg/ -o $(BUILD_DIR)/_go_.6 src/cmd/$@/*.go
	6l -L $(BUILD_DIR)/pkg/ -o $@ $(BUILD_DIR)/_go_.6

timed2: hbh parse str
	6g -I $(BUILD_DIR)/pkg/ -o $(BUILD_DIR)/_go_.6 src/cmd/$@/*.go
	6l -L $(BUILD_DIR)/pkg/ -o $@ $(BUILD_DIR)/_go_.6

# --- [ pkg ] ------------------------------------------------------------------
hbh:
	6g -p $@ -I $(BUILD_DIR)/pkg/ -o $(BUILD_DIR)/_go_.6 src/pkg/$@/*.go
	@rm -f $(BUILD_DIR)/pkg/$@.a
	gopack grc $(BUILD_DIR)/pkg/$@.a $(BUILD_DIR)/_go_.6

parse:
	6g -p $@ -I $(BUILD_DIR)/pkg/ -o $(BUILD_DIR)/_go_.6 src/pkg/$@/*.go
	@rm -f $(BUILD_DIR)/pkg/$@.a
	gopack grc $(BUILD_DIR)/pkg/$@.a $(BUILD_DIR)/_go_.6

str:
	6g -p $@ -I $(BUILD_DIR)/pkg/ -o $(BUILD_DIR)/_go_.6 src/pkg/$@/*.go
	@rm -f $(BUILD_DIR)/pkg/$@.a
	gopack grc $(BUILD_DIR)/pkg/$@.a $(BUILD_DIR)/_go_.6

# --- [ general ] --------------------------------------------------------------
_init_:
	@mkdir -p $(BUILD_DIR)/pkg/

clean:
	rm -rf $(BUILD_DIR)/
	rm -f timed1
	rm -f timed2
