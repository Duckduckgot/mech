package vimeo

import (
   "fmt"
   "testing"
   "time"
)

var clip_refs = []string{
   "https://vimeo.com/477957994/2282452868",
   "https://player.vimeo.com/video/412573977?h=f7f2d6fcb7",
   "https://player.vimeo.com/video/412573977?unlisted_hash=f7f2d6fcb7",
   "https://vimeo.com/477957994?unlisted_hash=2282452868",
   "https://vimeo.com/534685752",
}

func Test_Vimeo(t *testing.T) {
   for _, ref := range clip_refs {
      clip, err := New_Clip(ref)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", clip)
   }
}

var clips = []Clip{
   // vimeo.com/581039021/9603038895
   {581039021, "9603038895"},
}

func Test_Clip(t *testing.T) {
   web, err := New_JSON_Web()
   if err != nil {
      t.Fatal(err)
   }
   for _, clip := range clips {
      video, err := web.Video(&clip)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(video)
      time.Sleep(time.Second)
   }
}
