package main

import (
	"github.com/hoisie/web"
	"github.com/xyproto/genericsite"
	"github.com/xyproto/permissions2"
	"github.com/xyproto/siteengines"
	"github.com/xyproto/webhandle"
)

// TODO: Norwegian everywhere
// TODO: Different Redis database than the other sites

const JQUERY_VERSION = "2.0.0"

func ServeEngines(userState *permissions.UserState, mainMenuEntries genericsite.MenuEntries) {
	// The user engine
	userEngine := siteengines.NewUserEngine(userState)
	userEngine.ServePages("mosebark.roboticoverlords.org")

	// The admin engine
	adminEngine := siteengines.NewAdminEngine(userState)
	adminEngine.ServePages(MosebarkBaseCP, mainMenuEntries)

	// Wiki engine
	wikiEngine := siteengines.NewWikiEngine(userState)
	wikiEngine.ServePages(MosebarkBaseCP, mainMenuEntries)
}

func main() {

	// UserState with a Redis Connection Pool
	userState := permissions.NewUserState(4, true, ":6379")
	defer userState.Close()

	// The archlinux.no webpage,
	mainMenuEntries := ServeMosebark(userState, "/js/jquery-"+JQUERY_VERSION+".min.js")

	ServeEngines(userState, mainMenuEntries)

	// Compilation errors, vim-compatible filename
	web.Get("/error", webhandle.GenerateErrorHandle("errors.err"))
	web.Get("/errors", webhandle.GenerateErrorHandle("errors.err"))

	// Various .php and .asp urls that showed up in the log
	genericsite.ServeForFun()

	// TODO: Incorporate this check into web.go, to only return
	// stuff in the header when the HEAD method is requested:
	// if ctx.Request.Method == "HEAD" { return }
	// See also: curl -I

	// Serve on port 3004 for the Nginx instance to use
	web.Run("0.0.0.0:3004")
}
