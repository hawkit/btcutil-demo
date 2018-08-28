package btcutil


// TstAppDataDir makes the internal appDataDir function available to the test package
func TstAppDataDir(goos, appName string, roaming bool) string {
	return appDataDir(goos, appName, roaming)
}