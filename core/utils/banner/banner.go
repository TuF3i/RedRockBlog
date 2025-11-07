package banner

import (
	"fmt"

	"github.com/fatih/color"
)

func GetBanner() *Banner {
	return &Banner{}
}

func (root *Banner) ShowBanner() {
	fmt.Println()
	color.Red(":::::::::  :::::::::: :::::::::  :::::::::   ::::::::   ::::::::  :::    ::: ")
	color.Red(":+:    :+: :+:        :+:    :+: :+:    :+: :+:    :+: :+:    :+: :+:   :+:  ")
	color.Red("+:+    +:+ +:+        +:+    +:+ +:+    +:+ +:+    +:+ +:+        +:+  +:+   ")
	color.Red("+#++:++#:  +#++:++#   +#+    +:+ +#++:++#:  +#+    +:+ +#+        +#++:++    ")
	color.Red("+#+    +#+ +#+        +#+    +#+ +#+    +#+ +#+    +#+ +#+        +#+  +#+   ")
	color.Red("#+#    #+# #+#        #+#    #+# #+#    #+# #+#    #+# #+#    #+# #+#   #+#  ")
	color.Red("###    ### ########## #########  ###    ###  ########   ########  ###    ### ")

	color.Magenta("[Hex] 52 65 64 52 6F 63 6B ")
	color.Green("=======================================================================")
	color.Green("ProjectName: RedRockBlog")
	color.Green("Author: TuF3i")
	color.Green("GitHub: https://github.com/TuF3i")
	color.Green("")
	color.Green("I want to become a RedRocker, so I must try my best, and if you also")
	color.Green("want what I want, just make sure you will NEVER GIVE UP !!!")
	color.Green("=======================================================================")
}
