package browser

import "os/exec"

func openBrowser(url string) error {
	return runCmd("uiopen", url)
}

func setFlags(cmd *exec.Cmd) {}
