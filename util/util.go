/*
Sniperkit-Bot
- Status: analyzed
*/

package util

// CheckError General Error Checking
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
