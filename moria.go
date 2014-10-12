package main

import "github.com/towski/grubby"
import "github.com/towski/go_dfhack"
import "fmt"
import "sync"

func main(){
    var wg sync.WaitGroup
    fmt.Println("hey")
    dfhack.Connect()
    dfhack.Update()
    wg.Add(1)
    go func(){
        channel := grubby.Start("rb/file.rb")
        channel.Send(dfhack.GetDwarf(1))
        wg.Done()
    }()
    wg.Wait()
    //_ = foo
}
