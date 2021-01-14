// kern.go example application
// (c)copyright 2021 by Gerald Wodni <gerald.wodni@hmail.com>
package main

import (
    "boolshit.net/kern"
    "boolshit.net/kern/login"
    "boolshit.net/kern/logout"
    "boolshit.net/kern/view"
)

func main() {
    // create new kern instance on port 5000
    app := kern.New(":5000")

    // extend globals (available to all view-templates)
    view.Globals["AppName"] = "kern.go demo app"

    infoRouter := app.Router.NewMounted( "/info" )
    infoRouter.StaticText( "/about",   "this is kern.go" )
    infoRouter.StaticText( "/version", "0.0.0 - as fresh as they come" )

    // block all requests which are not authenticated against "view" permission
    app.Router.Mount( logout.Logout("/logout") )
    app.Router.Mount( login.PermissionReqired( "/", "view" ) )

    // add credential checkers
    login.Register( login.NewEnvironmentCredentialChecker() )
    login.Register( login.NewStaticCredentials( "tester", "mc testface", "read,write,drink,awesome" ) )

    // mount index.gohtml on "/"
    app.Router.Get("/", view.NewHandler( "index.gohtml" ) )

    // start server
    app.Run()
}
