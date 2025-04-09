package m5_download_plugin

import (
	"fmt"
	"plugin"
)

// LoadAndInvokePlugin loads the plugin and invokes the specified function
func LoadAndInvokePlugin(pluginPath string, pluginFunc string) error {
	// Load the plugin dynamically
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return fmt.Errorf("error loading plugin: %w", err)
	}

	// Look up the symbol (function) from the loaded plugin
	sym, err := p.Lookup(pluginFunc)
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
