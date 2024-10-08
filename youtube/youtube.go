package youtube

import (
   "github.com/Duckduckgot/rosso/http"
   "net/url"
   "path"
   "strings"
)

func Video_ID(data string, v *string) error {
   ref, err := url.Parse(data)
   if err != nil {
      return err
   }
   *v = ref.Query().Get("v")
   if *v == "" {
      *v = path.Base(ref.Path)
   }
   return nil
}

const origin = "https://www.youtube.com"

var HTTP_Client = http.Default_Client

type Image struct {
   Crop bool
   Height int
   Name string
   Width int
}

var Images = []Image{
   {Width:120, Height:90, Name:"default.jpg"},
   {Width:120, Height:90, Name:"1.jpg"},
   {Width:120, Height:90, Name:"2.jpg"},
   {Width:120, Height:90, Name:"3.jpg"},
   {Width:120, Height:90, Name:"default.webp"},
   {Width:120, Height:90, Name:"1.webp"},
   {Width:120, Height:90, Name:"2.webp"},
   {Width:120, Height:90, Name:"3.webp"},
   {Width:320, Height:180, Name:"mq1.jpg", Crop:true},
   {Width:320, Height:180, Name:"mq2.jpg", Crop:true},
   {Width:320, Height:180, Name:"mq3.jpg", Crop:true},
   {Width:320, Height:180, Name:"mqdefault.jpg"},
   {Width:320, Height:180, Name:"mq1.webp", Crop:true},
   {Width:320, Height:180, Name:"mq2.webp", Crop:true},
   {Width:320, Height:180, Name:"mq3.webp", Crop:true},
   {Width:320, Height:180, Name:"mqdefault.webp"},
   {Width:480, Height:360, Name:"0.jpg"},
   {Width:480, Height:360, Name:"hqdefault.jpg"},
   {Width:480, Height:360, Name:"hq1.jpg"},
   {Width:480, Height:360, Name:"hq2.jpg"},
   {Width:480, Height:360, Name:"hq3.jpg"},
   {Width:480, Height:360, Name:"0.webp"},
   {Width:480, Height:360, Name:"hqdefault.webp"},
   {Width:480, Height:360, Name:"hq1.webp"},
   {Width:480, Height:360, Name:"hq2.webp"},
   {Width:480, Height:360, Name:"hq3.webp"},
   {Width:640, Height:480, Name:"sddefault.jpg"},
   {Width:640, Height:480, Name:"sd1.jpg"},
   {Width:640, Height:480, Name:"sd2.jpg"},
   {Width:640, Height:480, Name:"sd3.jpg"},
   {Width:640, Height:480, Name:"sddefault.webp"},
   {Width:640, Height:480, Name:"sd1.webp"},
   {Width:640, Height:480, Name:"sd2.webp"},
   {Width:640, Height:480, Name:"sd3.webp"},
   {Width:1280, Height:720, Name:"hq720.jpg"},
   {Width:1280, Height:720, Name:"maxresdefault.jpg"},
   {Width:1280, Height:720, Name:"maxres1.jpg"},
   {Width:1280, Height:720, Name:"maxres2.jpg"},
   {Width:1280, Height:720, Name:"maxres3.jpg"},
   {Width:1280, Height:720, Name:"hq720.webp"},
   {Width:1280, Height:720, Name:"maxresdefault.webp"},
   {Width:1280, Height:720, Name:"maxres1.webp"},
   {Width:1280, Height:720, Name:"maxres2.webp"},
   {Width:1280, Height:720, Name:"maxres3.webp"},
}

func (i Image) Address(id string) string {
   var buf strings.Builder
   buf.WriteString("http://i.ytimg.com/vi")
   if strings.HasSuffix(i.Name, ".webp") {
      buf.WriteString("_webp")
   }
   buf.WriteByte('/')
   buf.WriteString(id)
   buf.WriteByte('/')
   buf.WriteString(i.Name)
   return buf.String()
}

type Item struct {
   CompactVideoRenderer *struct {
      Title struct {
         Runs []struct {
            Text string
         }
      }
      LengthText struct {
         Runs []struct {
            Text string
         }
      }
      VideoId string
   }
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer *struct {
               Contents []Item
            }
         }
      }
   }
}

func (s Search) Items() []Item {
   var items []Item
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      if sect.ItemSectionRenderer != nil {
         for _, item := range sect.ItemSectionRenderer.Contents {
            if item.CompactVideoRenderer != nil {
               items = append(items, item)
            }
         }
      }
   }
   return items
}
