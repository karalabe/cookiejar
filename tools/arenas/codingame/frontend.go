// CookieJar - A contestant's algorithm toolbox
// Copyright 2014 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The toolbox is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the CookieJar toolbox may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).

package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/inconshreveable/log15.v2"
	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/webengine"
)

// Arena QML to not require additional files
const arenaQml = `
import QtQuick 2.1
import QtWebEngine 1.0
import QtQuick.Controls 1.0

ApplicationWindow {
	title: "CodinGame Karalabe Arena"
    width: 1024
    height: 768
    
    WebEngineView {
    	objectName: "arena"
        anchors.fill: parent
        url: "http://codingame.com"
    }
}
`

// Command line flags for the arena
var repo = flag.String("repo", "challenges", "Challenge repository to work with")

// Application entry point, simple starts the arena QML app.
func main() {
	flag.Parse()
	if err := qml.Run(arena); err != nil {
		log15.Crit("failed to start arena", "error", err)
		os.Exit(-1)
	}
}

// QML application entry point to assemble the arena.
func arena() error {
	// Initialize the Chromium WebEngine
	webengine.Initialize()

	// Initialize the WebSocket backend
	port, err := backend()
	if err != nil {
		return err
	}
	// Create a new QML engine and terminate on quit
	engine := qml.NewEngine()
	engine.On("quit", func() { os.Exit(0) })

	// Load the main arena content from QML
	view, err := engine.LoadString("codingame.qml", arenaQml)
	if err != nil {
		return err
	}
	win := view.CreateWindow(nil)

	// Create the controller around the arena
	ctrl := &Control{
		arena:   win.Root().ObjectByName("arena"),
		backend: port,
	}
	ctrl.arena.On("loadingChanged", ctrl.InjectControl)

	// Start the application and wait till closure
	win.Call("showMaximized")
	win.Wait()
	return nil
}

// Arena controller to handle events.
type Control struct {
	arena   qml.Object
	backend int
}

// Invoked whenever a page is navigated.
func (c *Control) InjectControl(request qml.Object) {
	// Make sure the page loaded correctly
	if request.Int("status") != 2 {
		return
	}
	// Try and indefinitely find an editor, and fetch its contents into the title
	script := `
		// Dial back to the arena backend
		var ws = new WebSocket("ws://localhost:` + fmt.Sprint(c.backend) + `");
		ws.onmessage = function(msg) {
			var update = JSON.parse(msg.data);

			// Code update arrived, fetch enditor and ensure correct challenge
			var editor = document.getElementById('ideFrame').contentWindow.ace.edit('ace_edit');
			if (editor != 'undefined') {
				var title = document.getElementById('ideFrame').contentDocument.getElementsByClassName('challengeTitle')[0];
				if (title.textContent.trim() == update.name) {
					editor.setValue(update.source, 0);
					editor.navigateFileStart();
				}
			}
		};

		// Keep checking for editor presence
		setInterval(function() {
			var editor = document.getElementById('ideFrame').contentWindow.ace.edit('ace_edit');
			if (editor != 'undefined') {
				// Assemble a source report and send it to the backend
				var title = document.getElementById('ideFrame').contentDocument.getElementsByClassName('challengeTitle')[0];
				var challenge = {
					'name':   title.textContent,
					'source': editor.getValue()
				};
				ws.send(JSON.stringify(challenge));
			}
		}, 1000);
	`

	log15.Info("injecting frontend javascript")
	c.arena.Call("runJavaScript", script)
}
