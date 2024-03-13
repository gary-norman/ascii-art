package main

import (
	"log"
	"os"
)

/*

strings:
--------
hello
HELLO
HeLlo HuMaN
1Hello 2There
Hello\nThere
Hello\n\nThere
{Hello & There #}
hello There 1 to 2!
MaD3IrA&LiSboN
1a\"#FdwHywR&/()=
{|}~
[\]^_ 'a
RGB
:;<=>?@
\!" #$%&'"'"'()*+,-./
ABCDEFGHIJKLMNOPQRSTUVWXYZ
abcdefghijklmnopqrstuvwxyz
<a random string> with at least four lower case letters and three upper case letters.
<a random string> with at least five lower case letters, a space and two numbers.
<a random string> with at least one upper case letters and 3 special characters.
<a random string> with at least two lower case letters, two spaces, one number, two special characters and three upper case letters.

+Does the project run quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)
+Does the code obey the good practices?
+Is there a test file for this code?
+Are the tests checking each possible case?
+Is the output of the program well structured? Are the characters displayed correctly in line?
https://github.com/01-edu/public/blob/master/subjects/ascii-art/audit/README.md


····················································································
:                                                                                  :
:          _               _         __         _      _   _  _     _              :
:    __ _ | |__    ___  __| |  ___  / _|  __ _ | |__  (_) (_)| | __| | _ __ ___    :
:   / _` || '_ \  / __|/ _` | / _ \| |_  / _` || '_ \ | | | || |/ /| || '_ ` _ \   :
:  | (_| || |_) || (__| (_| ||  __/|  _|| (_| || | | || | | ||   < | || | | | | |  :
:   \__,_||_.__/  \___|\__,_| \___||_|   \__, ||_| |_||_|_/ ||_|\_\|_||_| |_| |_|  :
:                                        |___/          |__/                       :
:   _ __    ___   _ __    __ _  _ __  ___ | |_  _   _ __   ____      ____  __      :
:  | '_ \  / _ \ | '_ \  / _` || '__|/ __|| __|| | | |\ \ / /\ \ /\ / /\ \/ /      :
:  | | | || (_) || |_) || (_| || |   \__ \| |_ | |_| | \ V /  \ V  V /  >  <       :
:  |_| |_| \___/ | .__/  \__, ||_|   |___/ \__| \__,_|  \_/    \_/\_/  /_/\_\      :
:                |_|     ___|_|  ____  _____  _____  ____  _   _  ___     _  _  __ :
:   _   _  ____   / \   | __ )  / ___|| ____||  ___|/ ___|| | | ||_ _|   | || |/ / :
:  | | | ||_  /  / _ \  |  _ \ | |    |  _|  | |_  | |  _ | |_| | | | _  | || ' /  :
:  | |_| | / /  / ___ \ | |_) || |___ | |___ |  _| | |_| ||  _  | | || |_| || . \  :
:   \__, |/___|/_/   \_\|____/  \____||_____||_|    \____||_| |_||___|\___/ |_|\_\ :
:   |___/  __  __  _   _   ___   ____    ___   ____   ____  _____  _   _ __     __ :
:  | |    |  \/  || \ | | / _ \ |  _ \  / _ \ |  _ \ / ___||_   _|| | | |\ \   / / :
:  | |    | |\/| ||  \| || | | || |_) || | | || |_) |\___ \  | |  | | | | \ \ / /  :
:  | |___ | |  | || |\  || |_| ||  __/ | |_| ||  _ <  ___) | | |  | |_| |  \ V /   :
:  |_____||_|__|_||_|_\_| \___/_|_|     \__\_\|_| \_\|____/  |_|   \___/    \_/    :
:  \ \      / /\ \/ /\ \ / /|__  /                                                 :
:   \ \ /\ / /  \  /  \ V /   / /                                                  :
:    \ V  V /   /  \   | |   / /_                                                  :
:     \_/\_/   /_/\_\  |_|  /____|                                                 :
:                                                                                  :
:                                                                                  :
····················································································
*/

func main() {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	source := string(file)
	print(source)
}
