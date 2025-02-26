package plugin

import (
	"fmt"
	"plugin"
)

func LoadAndInvokePlugin() error {
	// Load the plugin dynamically
	p, err := plugin.Open("./plugin.so")
	if err != nil {
		return fmt.Errorf("error loading plugin: %w", err)
	}

	// Look up the symbol (function) from the loaded plugin
	sym, err := p.Lookup("PluginFunc")
	if err != nil {
		return fmt.Errorf("error looking up symbol: %w", err)
	}

	// Assert and call the function if found
	if fn, ok := sym.(func()); ok {
		fn()
	} else {
		return fmt.Errorf("function %s has unexpected type", "PluginFunc")
	}

	return nil
}
