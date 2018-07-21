////////////////////////////////////////////////////////////////
// '$name'                                                     /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
////////////////////////////////////////////////////////////////

// 'docs.go'

// by skinn
// on July 11, 2018
// for Davsk℠ Universe 4.0 project gbase

// ${DESCRIPTION}
package docs

/*
You have your choice of
    conclient
    davsk_web_client
    davsk_gui_client
Which connect to a davsk_game_srv that provides
    web server
    rpc server
Which connects to
    local postgresql universe database
    remote business server with http/rpc connected to remote postgresql business database.
My test_server is
    game web server on one port
    business web server on another port
    game rpc server on another port
    business rpc on another port
    connects to postgresql with universe & business database
Configuration options are to be in toml files for each app with defaults if file is not found then it creates one.

My test_client is like the conclient but runs off the event log from conclient so that you can play the same game over and over for testing.


*/
