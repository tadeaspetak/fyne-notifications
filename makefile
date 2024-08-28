.PHONY: package bundle

package:
	fyne package

# use as a.SetIcon(res.ResourceIconPng)
bundleIcon:
	mkdir -p res && fyne bundle --package=res --prefix=Resource icon.png >> res/icon.go 