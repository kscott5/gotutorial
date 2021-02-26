module hello

go 1.14

replace github.com/kscott5/gotutorial/greetings => ../greetings

replace github.com/kscott5/gotutorial/painkiller => ../painkiller

require github.com/kscott5/gotutorial/greetings v0.0.0-20201211172331-9e577cff9703 // indirect
