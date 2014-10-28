package main

import "github.com/towski/grubby"
import "github.com/towski/go_dfhack"
import "fmt"
import "sync"
import "time"

func main(){
    var wg sync.WaitGroup
    fmt.Println("hey")
    dfhack.Connect()
    dfhack.Update()
    channel := grubby.Start("dwarf_data.rb")
    wg.Add(1)
    go func(){
        fmt.Println("starting")
        for {
            dfhack.Update()
            time.Sleep(5 * time.Second)
            x := 0
            fmt.Println(dfhack.Size())
            for x < dfhack.Size() {
                channel.Send(dfhack.GetDwarf(x))
                x += 1
            }
        }
    }()
    wg.Wait()
    //_ = foo
}
