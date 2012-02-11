BUILD_DIR = _build_

all: _init_ streamdump
	@rm -f $(BUILD_DIR)/_go_.6

# --- [ cmd ] ------------------------------------------------------------------
streamdump: dump
	6g -I $(BUILD_DIR)/pkg/ -o $(BUILD_DIR)/_go_.6 src/cmd/$@/*.go
	6l -L $(BUILD_DIR)/pkg/ -o $@ $(BUILD_DIR)/_go_.6

# --- [ pkg ] ------------------------------------------------------------------
dump:
	6g -p $@ -I $(BUILD_DIR)/pkg/ -o $(BUILD_DIR)/_go_.6 src/pkg/$@/*.go
	@rm -f $(BUILD_DIR)/pkg/$@.a
	pack grc $(BUILD_DIR)/pkg/$@.a $(BUILD_DIR)/_go_.6

# --- [ general ] --------------------------------------------------------------
_init_:
	@mkdir -p $(BUILD_DIR)/pkg/

clean:
	rm -rf $(BUILD_DIR)/
	rm -f streamdump
