// Package nautilus is Nautilus scripts client.
// By using this package you can get the variables of the Nautilus scripts.
//
//		1.NAUTILUS_SCRIPT_SELECTED_FILE_PATHS
//				selected files (only if local)
//		2.NAUTILUS_SCRIPT_SELECTED_URIS
//				URIs for selected files
//		3.NAUTILUS_SCRIPT_CURRENT_URI
//				current location
//		4.NAUTILUS_SCRIPT_WINDOW_GEOMETRY
//				position and size of current window
//
// NautilusScriptsHowto https://help.ubuntu.com/community/NautilusScriptsHowto
package nautilus

import (
	"os/exec"
	"strings"
)

// GetSelectedFilePaths get values by parsing the NAUTILUS_SCRIPT_SELECTED_FILE_PATHS
func GetSelectedFilePaths() []string {
	output, _ := exec.Command("sh", "-c", "echo $NAUTILUS_SCRIPT_SELECTED_FILE_PATHS").Output()
	return strings.Split(string(output)[:len(output)-1], " ")
}

// GetSelectedURIs get values by parsing the NAUTILUS_SCRIPT_SELECTED_URIS
func GetSelectedURIs() []string {
	output, _ := exec.Command("sh", "-c", "echo $NAUTILUS_SCRIPT_SELECTED_URIS").Output()
	return strings.Split(string(output)[:len(output)-1], " ")
}

// GetCurrentURI get a value of the NAUTILUS_SCRIPT_CURRENT_URI
func GetCurrentURI() string {
	output, _ := exec.Command("sh", "-c", "echo $NAUTILUS_SCRIPT_CURRENT_URI").Output()
	return string(output[:len(output)-1])
}

// GetWindowGeometry get a value of the NAUTILUS_SCRIPT_WINDOW_GEOMETRY
func GetWindowGeometry() string {
	output, _ := exec.Command("sh", "-c", "echo $NAUTILUS_SCRIPT_WINDOW_GEOMETRY").Output()
	return string(output[:len(output)-1])
}
