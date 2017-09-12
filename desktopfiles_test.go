package desktopfiles

import (
	"fmt"
	"testing"
)

func TestDesktopfiles(t *testing.T) {
	fmt.Println(&Desktopfile{
		Exec: "/usr/bin/GroupMatcher",
	})
}
