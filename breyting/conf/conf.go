// Package conf implements functions for watching config files and retrieving
// relevant page download information.
package conf

import (
	"log"
	"time"

	"github.com/howeyc/fsnotify"
	"github.com/mewmew/playground/breyting/page"
	"github.com/mewpkg/goini"
)

// activeDict is the active config file.
var activeDict ini.Dict

// settingsSection is the name of the settings section.
const settingsSection = ""

// Reload reloads the config file and adds a watcher to all new pages.
func Reload(confPath string) (err error) {
	dict, err := ini.Load(confPath)
	if err != nil {
		return err
	}
	for section := range dict {
		if section == settingsSection {
			rawTimeout, found := dict.GetString(settingsSection, "timeout")
			if found {
				page.Timeout, err = time.ParseDuration(rawTimeout)
				if err != nil {
					log.Println(err)
					page.Timeout = page.DefaultTimeout
				}
			} else {
				page.Timeout = page.DefaultTimeout
			}
			continue
		}
		// Initially no pages are active, add watchers to them all.
		var active bool
		if activeDict != nil {
			// Only add watchers to new pages.
			_, active = activeDict[section]
		}
		if !active {
			page := getPage(dict, section)
			go page.Watch()
		}
	}
	activeDict = dict
	// All active page watchers will justify their existance once every timeout
	// interval. If no longer justified they will commit suicide.
	return nil
}

// getPage returns the page of a given section in dict, or nil if no valid page
// could be located.
func getPage(dict ini.Dict, section string) (p *page.Page) {
	if !isValidPageSection(dict, section) {
		return nil
	}
	rawSel, ok := dict.GetString(section, "sel")
	if !ok {
		/// ### [ todo ] ###
		///   - how to handle empty selector?
		/// ### [/ todo ] ###
		rawSel = ""
	}
	p, err := page.New(section, rawSel)
	if err != nil {
		log.Fatalln(err) /// ### tmp ###
	}
	return p
}

// isValidPageSection returns true if the section is present and it isn't the
// predeclared settings section.
func isValidPageSection(dict ini.Dict, section string) bool {
	if section == settingsSection {
		return false
	}
	_, ok := dict[section]
	if !ok {
		return false
	}
	return true
}

// Watch monitors the config file for changes and reloads it after modification
// events.
func Watch(confPath string) {
	for {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatalln(err)
		}
		defer watcher.Close()
		err = watcher.WatchFlags(confPath, fsnotify.FSN_MODIFY|fsnotify.FSN_DELETE)
		if err != nil {
			log.Fatalln(err)
		}
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsDelete() {
					// Some editors replace the file with a new one when saving. This
					// will generate a delete event and requires a new watcher for
					// the new file.
					time.Sleep(1 * time.Second)
					go Watch(confPath)
					return
				}
				if ev.IsModify() {
					Reload(confPath)
				}
			case err = <-watcher.Error:
				log.Println(err)
			}
		}
	}
}
