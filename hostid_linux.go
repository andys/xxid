// +build linux

package xxid

import "io/ioutil"

func readPlatformMachineID() (string, error) {
	b, err := ioutil.ReadFile("/var/lib/dbus/machine")
	if err != nil {
		// try fallback path
		b, err = readFile("/etc/machine-id")
	}

	return string(b), err
}
