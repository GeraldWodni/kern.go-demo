// kern.go example application
// (c)copyright 2021 by Gerald Wodni <gerald.wodni@hmail.com>
package main

import (
    "boolshit.net/kern"
    "boolshit.net/kern/view"
)

func main() {
    // create new kern instance on port 5000
    app := kern.New(":5000")

    // extend globals (available to all view-templates)
    view.Globals["AppName"] = "kern.go demo app"

    // mount index.gohtml on "/"
    app.Router.Get("/", view.NewHandler( "index.gohtml" ) )

    // start server
    app.Run()
}
