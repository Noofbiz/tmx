# tmx
A library for parsing TMX files

To use:

```go
f, err := os.Open("test1.tmx")
defer f.Close()
if err != nil {
  fmt.Println(err)
  return
}
m, err := tmx.Parse(f)
if err != nil {
  fmt.Println(err)
  return
}
// Do stuff with your tmx map...
```

If your tmx resources are in another folder, or are somewhere other than where
the binary is called, you can set it by setting `TMXURL` to the right path.
Doing this will allow you to use external tilesets and templates for the map.
