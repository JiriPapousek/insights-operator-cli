/*
Copyright © 2019, 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/c-bata/go-prompt"
	"testing"
)

// createDocumentWithCommand construct an instance of prompt.Document containing the command and cursor position.
func createDocumentWithCommand(t *testing.T, command string) prompt.Document {
	buffer := prompt.NewBuffer()
	if buffer == nil {
		t.Fatal("Error in prompt library - can not constructs new buffer")
	}
	buffer.InsertText(command, false, true)
	document := buffer.Document()
	if document == nil {
		t.Fatal("Error in prompt library - can not get document for a buffer")
	}
	return *document
}

// checkSuggestionCount checks the number of suggestions returned by suggester
func checkSuggestionCount(t *testing.T, suggests []prompt.Suggest, expected int) {
	if len(suggests) != expected {
		t.Fatal("Invalid suggestion returned by completer:", suggests)
	}
}

// checkSuggestionCount checks the suggestion text and description
func checkSuggestion(t *testing.T, suggest prompt.Suggest, command string, description string) {
	if suggest.Text != command {
		t.Fatal("Invalid suggestion command:", suggest.Text)
	}
	if suggest.Description != description {
		t.Fatal("Invalid suggestion description:", suggest.Description)
	}
}

// TestCompleterEmptyInput check which suggestions are returned for empty input
func TestCompleterEmptyInput(t *testing.T) {
	suggests := completer(createDocumentWithCommand(t, ""))
	checkSuggestionCount(t, suggests, 0)
}

// TestCompleterHelpCommand check which suggestions are returned for 'help' input
func TestCompleterHelpCommand(t *testing.T) {
	suggests := completer(createDocumentWithCommand(t, "help"))
	checkSuggestionCount(t, suggests, 1)
	checkSuggestion(t, suggests[0], "help", "show help with all commands")
}
